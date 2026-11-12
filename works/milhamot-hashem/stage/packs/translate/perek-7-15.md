# Translation pack: milhamot-hashem perek-7-15 (Milchamot ha-Shem 7:20–7:20)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-15`.

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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
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


## Milchamot ha-Shem 7:20

Hebrew (Wikisource): צב) ולא זו בלבד. אלא אף גם זאת שייחסו לו החוש המגונה והפחות שבחמשה חושים, והוא המשגל והזיווג, כי ידוע מה שכתב הרמב"ם ז"ל בפרק מ"ז חלק א'. שמהחמשה חושים שהם הראייה, והשמע, והריח, וישמע ה', וירא ה', וירח ה', ולא הושאל לו חוש הטעם והמשוש לומר ויטעם ה', וימשש ה'. והסבה לזה הוא כמו שנתבאר לעיל שכל מה שידמוהו ההמון חסרון אי אפשר לציירו בבורא ית', ולכן לא השתמשו בזה בספרי הנבואה עכ"ל. וכן כתב הרב שבילי אמונה בנתיב הא'. ובדרך הזה דרכו רז"ל בעלי המשנה והתלמוד ומתקני התפלות והפייטנים כרבי' סעדיה גאון ור"י הלוי בפיוטיהם הנעימים המעוררים את האדם לאהבת השי"ת וליראתו ולדבקה בו. ומה מאד מתקו דברי רבי יהודה הלוי בסליחה המתחלת יה שמך ארוממך, אמר בסופה, השתונן והכונן, והתבונן בסודיך, והבטת מה אתה, ומאין יסודיך, ומי הבינך, ומי הכינך, וכח מי ינידיך, והבט אל גבורות אל, והעירה כבודיך. חקור פעליו, רק אליו אל תשלח ידיך, כי תדרוש בסוף וראש במופלה ובמכוסה ע"כ. אשר על כן אלו שמרו המקובלים החדשים אזהרת רז"ל ולא נפתו לדברי נביא השקר המסית מחבר הזהר ולא שלחו רעיונותיהם במופלא ובמכוסה מהם לא באו לכל זה. ובפרק ח' לשלישי כתב הר' המורה שחוש המשוש הוא חרפה וכלמה לנו, אשר בעבורו נתאוה למאכל ומשתה ומשגל, והאיש המשכיל מריך למעט בהם מה שאפשר לו (כמו שביאר הרב בהלכות דעות) ולא ירחיב הדבור בהם, וישים תכליתו תכלית האדם מאשר הוא אדם, והוא ציור המושכלות, אשר בהם ישיג ידיעת הבורא, אבל האכילה והמשגל הם חרפתינו וכלימתינו, כי הושמו באדם לצורך קיומו וקיום המין, והוא דומה בהם לבהמה, וכל שכן המשגל שאין ראוי לדבר בו יותר כמו שביארתי בפרקי אבות ממה שבא בתורתינו הקדושה ממאוס בו, ואסור לזכרו ולדבר בו כלל, וכבר ידעת אמרם על אלישע הנביא ע"ה שנקרא קדוש מפני שלא היה חושב בו עש שלא היה רואה קרי, וגם ידעת גודל האיסור שבא אצלינו בענין נבלות הפה (כדגרסינן בכתובות, על כן על בחוריו לא ישמח ה' ואת יתומיו ואת אלמנותיו לא ירחם כי כלו חנף ומרע וכל פה דובר נבלה בכל זאת לא שב אפו ועוד ידו נטויה, מאי ועוד ידו נטויה אמר רב חנן ב"א הכל יודעים למה נכנסה כלה לחופה אלא כל המנבל פיו ומוציא דבר נבלה מפיו אפילו נחתם לו גזר דין של שבעים שנה לטובה נהפך עליו לרעה ע"כ. ובפרק במה מדליקים גרסינן בעון נבלות הפה צרות רבות ודעות וגזרות קשות ותחדשות בעולם ובחורי ישראל מתים ואלמנות ויתומים צועקים ואינם נענים שנאמר על כן על בחוריו לא ישמח ה' וכו' כי כלו חנף ומרע כל פה דובר נבלה וכו' מאי ועוד ידו נטויה א"ר חנן בר אבא א"ר חסדא כל המנבל פיו מעמיקים לו גיהנום שנ' שוחה עמוקה פי זרות זעום ה' יפל שם ע"כ). והטעם לפי שהדיבור בלשון הוא מסגולת האדם וטוב השם אשר גמר עליו להבדילו בו משאר בעלי החיים שנאמר מי שם פה לאדם, ואמר הנביא ה' אלהים נתן לי לשון למודים, ואין ראוי שנשתמש בטובה אשר גמלנו השי"ת לשלמות ללמוד וללמד, בדבר פחות ומגונה ובחרפה הגמורה שיש בנו ונדמה לגוים הסכלים הזונים בשיריהם שירי ענבים, ופחיתות אשר יתפארו בהם לרוב סכלות ופחיתותם (כדאמור רבנן תפארת גוים עבירה. גוי מפערה מעי) לא כאלה חלק יעקב, שנאמר בהם ואתר תהיו לי ממלכת כהנים וגוי קדוש, וכל מי שמשתמש במחשבה או בדבריו מענייני המשגל אשר הוא חרפה לנו ביותר מן הצריך לו (כמאמר רז"ל מגיד לאדם מה שיחו אפילו שיחה יתירה שבין איש לאשתו מגידין לו) וכ"ש אם יאמר שירים, הנה נשתמש בטובה אשר גמלו השי"צ בדברי מרי וסכלות גמורה ועבר על מצות הבורא ית'. והוא כמי שנאמר בהם וכסף הרביתי להם וזהב עשו לבעל, זהו תורף דברי הר' המורה ז"ל. גם הרס"ג בס' האמונות מאמר עשירי פרק ו' דבר בגנות רבוי המשגל ולכלוכו וטנופו והחרפה והבוז אשר בו, ושאין התואה הזאת באדם כי אם להקים בה זרע כמצות השי"ת ואתם פרו ורבו וכו' ובהצנע בעת הראוי ויעצור בה מהרבות ע"כ. ובעיני ראיתי אחד מהמון העם שאל לתינוק קטן ואמר האם אביך אוהב את אמך ומנשקה, וגער בו הדיין כמה"ר סלי' צאלח ז"ל ואמר לו כלך מגבלות הפה כי ענשו חמור.

| id | word | punctuation after |
|---|---|---|
| 7:20.1 | צב | ) |
| 7:20.2 | ולא |  |
| 7:20.3 | זו |  |
| 7:20.4 | בלבד | . |
| 7:20.5 | אלא |  |
| 7:20.6 | אף |  |
| 7:20.7 | גם |  |
| 7:20.8 | זאת |  |
| 7:20.9 | שייחסו |  |
| 7:20.10 | לו |  |
| 7:20.11 | החוש |  |
| 7:20.12 | המגונה |  |
| 7:20.13 | והפחות |  |
| 7:20.14 | שבחמשה |  |
| 7:20.15 | חושים | , |
| 7:20.16 | והוא |  |
| 7:20.17 | המשגל |  |
| 7:20.18 | והזיווג | , |
| 7:20.19 | כי |  |
| 7:20.20 | ידוע |  |
| 7:20.21 | מה |  |
| 7:20.22 | שכתב |  |
| 7:20.23 | הרמב"ם |  |
| 7:20.24 | ז"ל |  |
| 7:20.25 | בפרק |  |
| 7:20.26 | מ"ז |  |
| 7:20.27 | חלק |  |
| 7:20.28 | א | '. |
| 7:20.29 | שמהחמשה |  |
| 7:20.30 | חושים |  |
| 7:20.31 | שהם |  |
| 7:20.32 | הראייה | , |
| 7:20.33 | והשמע | , |
| 7:20.34 | והריח | , |
| 7:20.35 | וישמע |  |
| 7:20.36 | ה | ', |
| 7:20.37 | וירא |  |
| 7:20.38 | ה | ', |
| 7:20.39 | וירח |  |
| 7:20.40 | ה | ', |
| 7:20.41 | ולא |  |
| 7:20.42 | הושאל |  |
| 7:20.43 | לו |  |
| 7:20.44 | חוש |  |
| 7:20.45 | הטעם |  |
| 7:20.46 | והמשוש |  |
| 7:20.47 | לומר |  |
| 7:20.48 | ויטעם |  |
| 7:20.49 | ה | ', |
| 7:20.50 | וימשש |  |
| 7:20.51 | ה | '. |
| 7:20.52 | והסבה |  |
| 7:20.53 | לזה |  |
| 7:20.54 | הוא |  |
| 7:20.55 | כמו |  |
| 7:20.56 | שנתבאר |  |
| 7:20.57 | לעיל |  |
| 7:20.58 | שכל |  |
| 7:20.59 | מה |  |
| 7:20.60 | שידמוהו |  |
| 7:20.61 | ההמון |  |
| 7:20.62 | חסרון |  |
| 7:20.63 | אי |  |
| 7:20.64 | אפשר |  |
| 7:20.65 | לציירו |  |
| 7:20.66 | בבורא |  |
| 7:20.67 | ית | ', |
| 7:20.68 | ולכן |  |
| 7:20.69 | לא |  |
| 7:20.70 | השתמשו |  |
| 7:20.71 | בזה |  |
| 7:20.72 | בספרי |  |
| 7:20.73 | הנבואה |  |
| 7:20.74 | עכ"ל | . |
| 7:20.75 | וכן |  |
| 7:20.76 | כתב |  |
| 7:20.77 | הרב |  |
| 7:20.78 | שבילי |  |
| 7:20.79 | אמונה |  |
| 7:20.80 | בנתיב |  |
| 7:20.81 | הא | '. |
| 7:20.82 | ובדרך |  |
| 7:20.83 | הזה |  |
| 7:20.84 | דרכו |  |
| 7:20.85 | רז"ל |  |
| 7:20.86 | בעלי |  |
| 7:20.87 | המשנה |  |
| 7:20.88 | והתלמוד |  |
| 7:20.89 | ומתקני |  |
| 7:20.90 | התפלות |  |
| 7:20.91 | והפייטנים |  |
| 7:20.92 | כרבי | ' |
| 7:20.93 | סעדיה |  |
| 7:20.94 | גאון |  |
| 7:20.95 | ור"י |  |
| 7:20.96 | הלוי |  |
| 7:20.97 | בפיוטיהם |  |
| 7:20.98 | הנעימים |  |
| 7:20.99 | המעוררים |  |
| 7:20.100 | את |  |
| 7:20.101 | האדם |  |
| 7:20.102 | לאהבת |  |
| 7:20.103 | השי"ת |  |
| 7:20.104 | וליראתו |  |
| 7:20.105 | ולדבקה |  |
| 7:20.106 | בו | . |
| 7:20.107 | ומה |  |
| 7:20.108 | מאד |  |
| 7:20.109 | מתקו |  |
| 7:20.110 | דברי |  |
| 7:20.111 | רבי |  |
| 7:20.112 | יהודה |  |
| 7:20.113 | הלוי |  |
| 7:20.114 | בסליחה |  |
| 7:20.115 | המתחלת |  |
| 7:20.116 | יה |  |
| 7:20.117 | שמך |  |
| 7:20.118 | ארוממך | , |
| 7:20.119 | אמר |  |
| 7:20.120 | בסופה | , |
| 7:20.121 | השתונן |  |
| 7:20.122 | והכונן | , |
| 7:20.123 | והתבונן |  |
| 7:20.124 | בסודיך | , |
| 7:20.125 | והבטת |  |
| 7:20.126 | מה |  |
| 7:20.127 | אתה | , |
| 7:20.128 | ומאין |  |
| 7:20.129 | יסודיך | , |
| 7:20.130 | ומי |  |
| 7:20.131 | הבינך | , |
| 7:20.132 | ומי |  |
| 7:20.133 | הכינך | , |
| 7:20.134 | וכח |  |
| 7:20.135 | מי |  |
| 7:20.136 | ינידיך | , |
| 7:20.137 | והבט |  |
| 7:20.138 | אל |  |
| 7:20.139 | גבורות |  |
| 7:20.140 | אל | , |
| 7:20.141 | והעירה |  |
| 7:20.142 | כבודיך | . |
| 7:20.143 | חקור |  |
| 7:20.144 | פעליו | , |
| 7:20.145 | רק |  |
| 7:20.146 | אליו |  |
| 7:20.147 | אל |  |
| 7:20.148 | תשלח |  |
| 7:20.149 | ידיך | , |
| 7:20.150 | כי |  |
| 7:20.151 | תדרוש |  |
| 7:20.152 | בסוף |  |
| 7:20.153 | וראש |  |
| 7:20.154 | במופלה |  |
| 7:20.155 | ובמכוסה |  |
| 7:20.156 | ע"כ | . |
| 7:20.157 | אשר |  |
| 7:20.158 | על |  |
| 7:20.159 | כן |  |
| 7:20.160 | אלו |  |
| 7:20.161 | שמרו |  |
| 7:20.162 | המקובלים |  |
| 7:20.163 | החדשים |  |
| 7:20.164 | אזהרת |  |
| 7:20.165 | רז"ל |  |
| 7:20.166 | ולא |  |
| 7:20.167 | נפתו |  |
| 7:20.168 | לדברי |  |
| 7:20.169 | נביא |  |
| 7:20.170 | השקר |  |
| 7:20.171 | המסית |  |
| 7:20.172 | מחבר |  |
| 7:20.173 | הזהר |  |
| 7:20.174 | ולא |  |
| 7:20.175 | שלחו |  |
| 7:20.176 | רעיונותיהם |  |
| 7:20.177 | במופלא |  |
| 7:20.178 | ובמכוסה |  |
| 7:20.179 | מהם |  |
| 7:20.180 | לא |  |
| 7:20.181 | באו |  |
| 7:20.182 | לכל |  |
| 7:20.183 | זה | . |
| 7:20.184 | ובפרק |  |
| 7:20.185 | ח | ' |
| 7:20.186 | לשלישי |  |
| 7:20.187 | כתב |  |
| 7:20.188 | הר | ' |
| 7:20.189 | המורה |  |
| 7:20.190 | שחוש |  |
| 7:20.191 | המשוש |  |
| 7:20.192 | הוא |  |
| 7:20.193 | חרפה |  |
| 7:20.194 | וכלמה |  |
| 7:20.195 | לנו | , |
| 7:20.196 | אשר |  |
| 7:20.197 | בעבורו |  |
| 7:20.198 | נתאוה |  |
| 7:20.199 | למאכל |  |
| 7:20.200 | ומשתה |  |
| 7:20.201 | ומשגל | , |
| 7:20.202 | והאיש |  |
| 7:20.203 | המשכיל |  |
| 7:20.204 | מריך |  |
| 7:20.205 | למעט |  |
| 7:20.206 | בהם |  |
| 7:20.207 | מה |  |
| 7:20.208 | שאפשר |  |
| 7:20.209 | לו |  |
| 7:20.210 | כמו |  |
| 7:20.211 | שביאר |  |
| 7:20.212 | הרב |  |
| 7:20.213 | בהלכות |  |
| 7:20.214 | דעות | ) |
| 7:20.215 | ולא |  |
| 7:20.216 | ירחיב |  |
| 7:20.217 | הדבור |  |
| 7:20.218 | בהם | , |
| 7:20.219 | וישים |  |
| 7:20.220 | תכליתו |  |
| 7:20.221 | תכלית |  |
| 7:20.222 | האדם |  |
| 7:20.223 | מאשר |  |
| 7:20.224 | הוא |  |
| 7:20.225 | אדם | , |
| 7:20.226 | והוא |  |
| 7:20.227 | ציור |  |
| 7:20.228 | המושכלות | , |
| 7:20.229 | אשר |  |
| 7:20.230 | בהם |  |
| 7:20.231 | ישיג |  |
| 7:20.232 | ידיעת |  |
| 7:20.233 | הבורא | , |
| 7:20.234 | אבל |  |
| 7:20.235 | האכילה |  |
| 7:20.236 | והמשגל |  |
| 7:20.237 | הם |  |
| 7:20.238 | חרפתינו |  |
| 7:20.239 | וכלימתינו | , |
| 7:20.240 | כי |  |
| 7:20.241 | הושמו |  |
| 7:20.242 | באדם |  |
| 7:20.243 | לצורך |  |
| 7:20.244 | קיומו |  |
| 7:20.245 | וקיום |  |
| 7:20.246 | המין | , |
| 7:20.247 | והוא |  |
| 7:20.248 | דומה |  |
| 7:20.249 | בהם |  |
| 7:20.250 | לבהמה | , |
| 7:20.251 | וכל |  |
| 7:20.252 | שכן |  |
| 7:20.253 | המשגל |  |
| 7:20.254 | שאין |  |
| 7:20.255 | ראוי |  |
| 7:20.256 | לדבר |  |
| 7:20.257 | בו |  |
| 7:20.258 | יותר |  |
| 7:20.259 | כמו |  |
| 7:20.260 | שביארתי |  |
| 7:20.261 | בפרקי |  |
| 7:20.262 | אבות |  |
| 7:20.263 | ממה |  |
| 7:20.264 | שבא |  |
| 7:20.265 | בתורתינו |  |
| 7:20.266 | הקדושה |  |
| 7:20.267 | ממאוס |  |
| 7:20.268 | בו | , |
| 7:20.269 | ואסור |  |
| 7:20.270 | לזכרו |  |
| 7:20.271 | ולדבר |  |
| 7:20.272 | בו |  |
| 7:20.273 | כלל | , |
| 7:20.274 | וכבר |  |
| 7:20.275 | ידעת |  |
| 7:20.276 | אמרם |  |
| 7:20.277 | על |  |
| 7:20.278 | אלישע |  |
| 7:20.279 | הנביא |  |
| 7:20.280 | ע"ה |  |
| 7:20.281 | שנקרא |  |
| 7:20.282 | קדוש |  |
| 7:20.283 | מפני |  |
| 7:20.284 | שלא |  |
| 7:20.285 | היה |  |
| 7:20.286 | חושב |  |
| 7:20.287 | בו |  |
| 7:20.288 | עש |  |
| 7:20.289 | שלא |  |
| 7:20.290 | היה |  |
| 7:20.291 | רואה |  |
| 7:20.292 | קרי | , |
| 7:20.293 | וגם |  |
| 7:20.294 | ידעת |  |
| 7:20.295 | גודל |  |
| 7:20.296 | האיסור |  |
| 7:20.297 | שבא |  |
| 7:20.298 | אצלינו |  |
| 7:20.299 | בענין |  |
| 7:20.300 | נבלות |  |
| 7:20.301 | הפה |  |
| 7:20.302 | כדגרסינן |  |
| 7:20.303 | בכתובות | , |
| 7:20.304 | על |  |
| 7:20.305 | כן |  |
| 7:20.306 | על |  |
| 7:20.307 | בחוריו |  |
| 7:20.308 | לא |  |
| 7:20.309 | ישמח |  |
| 7:20.310 | ה | ' |
| 7:20.311 | ואת |  |
| 7:20.312 | יתומיו |  |
| 7:20.313 | ואת |  |
| 7:20.314 | אלמנותיו |  |
| 7:20.315 | לא |  |
| 7:20.316 | ירחם |  |
| 7:20.317 | כי |  |
| 7:20.318 | כלו |  |
| 7:20.319 | חנף |  |
| 7:20.320 | ומרע |  |
| 7:20.321 | וכל |  |
| 7:20.322 | פה |  |
| 7:20.323 | דובר |  |
| 7:20.324 | נבלה |  |
| 7:20.325 | בכל |  |
| 7:20.326 | זאת |  |
| 7:20.327 | לא |  |
| 7:20.328 | שב |  |
| 7:20.329 | אפו |  |
| 7:20.330 | ועוד |  |
| 7:20.331 | ידו |  |
| 7:20.332 | נטויה | , |
| 7:20.333 | מאי |  |
| 7:20.334 | ועוד |  |
| 7:20.335 | ידו |  |
| 7:20.336 | נטויה |  |
| 7:20.337 | אמר |  |
| 7:20.338 | רב |  |
| 7:20.339 | חנן |  |
| 7:20.340 | ב"א |  |
| 7:20.341 | הכל |  |
| 7:20.342 | יודעים |  |
| 7:20.343 | למה |  |
| 7:20.344 | נכנסה |  |
| 7:20.345 | כלה |  |
| 7:20.346 | לחופה |  |
| 7:20.347 | אלא |  |
| 7:20.348 | כל |  |
| 7:20.349 | המנבל |  |
| 7:20.350 | פיו |  |
| 7:20.351 | ומוציא |  |
| 7:20.352 | דבר |  |
| 7:20.353 | נבלה |  |
| 7:20.354 | מפיו |  |
| 7:20.355 | אפילו |  |
| 7:20.356 | נחתם |  |
| 7:20.357 | לו |  |
| 7:20.358 | גזר |  |
| 7:20.359 | דין |  |
| 7:20.360 | של |  |
| 7:20.361 | שבעים |  |
| 7:20.362 | שנה |  |
| 7:20.363 | לטובה |  |
| 7:20.364 | נהפך |  |
| 7:20.365 | עליו |  |
| 7:20.366 | לרעה |  |
| 7:20.367 | ע"כ | . |
| 7:20.368 | ובפרק |  |
| 7:20.369 | במה |  |
| 7:20.370 | מדליקים |  |
| 7:20.371 | גרסינן |  |
| 7:20.372 | בעון |  |
| 7:20.373 | נבלות |  |
| 7:20.374 | הפה |  |
| 7:20.375 | צרות |  |
| 7:20.376 | רבות |  |
| 7:20.377 | ודעות |  |
| 7:20.378 | וגזרות |  |
| 7:20.379 | קשות |  |
| 7:20.380 | ותחדשות |  |
| 7:20.381 | בעולם |  |
| 7:20.382 | ובחורי |  |
| 7:20.383 | ישראל |  |
| 7:20.384 | מתים |  |
| 7:20.385 | ואלמנות |  |
| 7:20.386 | ויתומים |  |
| 7:20.387 | צועקים |  |
| 7:20.388 | ואינם |  |
| 7:20.389 | נענים |  |
| 7:20.390 | שנאמר |  |
| 7:20.391 | על |  |
| 7:20.392 | כן |  |
| 7:20.393 | על |  |
| 7:20.394 | בחוריו |  |
| 7:20.395 | לא |  |
| 7:20.396 | ישמח |  |
| 7:20.397 | ה | ' |
| 7:20.398 | וכו | ' |
| 7:20.399 | כי |  |
| 7:20.400 | כלו |  |
| 7:20.401 | חנף |  |
| 7:20.402 | ומרע |  |
| 7:20.403 | כל |  |
| 7:20.404 | פה |  |
| 7:20.405 | דובר |  |
| 7:20.406 | נבלה |  |
| 7:20.407 | וכו | ' |
| 7:20.408 | מאי |  |
| 7:20.409 | ועוד |  |
| 7:20.410 | ידו |  |
| 7:20.411 | נטויה |  |
| 7:20.412 | א"ר |  |
| 7:20.413 | חנן |  |
| 7:20.414 | בר |  |
| 7:20.415 | אבא |  |
| 7:20.416 | א"ר |  |
| 7:20.417 | חסדא |  |
| 7:20.418 | כל |  |
| 7:20.419 | המנבל |  |
| 7:20.420 | פיו |  |
| 7:20.421 | מעמיקים |  |
| 7:20.422 | לו |  |
| 7:20.423 | גיהנום |  |
| 7:20.424 | שנ | ' |
| 7:20.425 | שוחה |  |
| 7:20.426 | עמוקה |  |
| 7:20.427 | פי |  |
| 7:20.428 | זרות |  |
| 7:20.429 | זעום |  |
| 7:20.430 | ה | ' |
| 7:20.431 | יפל |  |
| 7:20.432 | שם |  |
| 7:20.433 | ע"כ | ). |
| 7:20.434 | והטעם |  |
| 7:20.435 | לפי |  |
| 7:20.436 | שהדיבור |  |
| 7:20.437 | בלשון |  |
| 7:20.438 | הוא |  |
| 7:20.439 | מסגולת |  |
| 7:20.440 | האדם |  |
| 7:20.441 | וטוב |  |
| 7:20.442 | השם |  |
| 7:20.443 | אשר |  |
| 7:20.444 | גמר |  |
| 7:20.445 | עליו |  |
| 7:20.446 | להבדילו |  |
| 7:20.447 | בו |  |
| 7:20.448 | משאר |  |
| 7:20.449 | בעלי |  |
| 7:20.450 | החיים |  |
| 7:20.451 | שנאמר |  |
| 7:20.452 | מי |  |
| 7:20.453 | שם |  |
| 7:20.454 | פה |  |
| 7:20.455 | לאדם | , |
| 7:20.456 | ואמר |  |
| 7:20.457 | הנביא |  |
| 7:20.458 | ה | ' |
| 7:20.459 | אלהים |  |
| 7:20.460 | נתן |  |
| 7:20.461 | לי |  |
| 7:20.462 | לשון |  |
| 7:20.463 | למודים | , |
| 7:20.464 | ואין |  |
| 7:20.465 | ראוי |  |
| 7:20.466 | שנשתמש |  |
| 7:20.467 | בטובה |  |
| 7:20.468 | אשר |  |
| 7:20.469 | גמלנו |  |
| 7:20.470 | השי"ת |  |
| 7:20.471 | לשלמות |  |
| 7:20.472 | ללמוד |  |
| 7:20.473 | וללמד | , |
| 7:20.474 | בדבר |  |
| 7:20.475 | פחות |  |
| 7:20.476 | ומגונה |  |
| 7:20.477 | ובחרפה |  |
| 7:20.478 | הגמורה |  |
| 7:20.479 | שיש |  |
| 7:20.480 | בנו |  |
| 7:20.481 | ונדמה |  |
| 7:20.482 | לגוים |  |
| 7:20.483 | הסכלים |  |
| 7:20.484 | הזונים |  |
| 7:20.485 | בשיריהם |  |
| 7:20.486 | שירי |  |
| 7:20.487 | ענבים | , |
| 7:20.488 | ופחיתות |  |
| 7:20.489 | אשר |  |
| 7:20.490 | יתפארו |  |
| 7:20.491 | בהם |  |
| 7:20.492 | לרוב |  |
| 7:20.493 | סכלות |  |
| 7:20.494 | ופחיתותם |  |
| 7:20.495 | כדאמור |  |
| 7:20.496 | רבנן |  |
| 7:20.497 | תפארת |  |
| 7:20.498 | גוים |  |
| 7:20.499 | עבירה | . |
| 7:20.500 | גוי |  |
| 7:20.501 | מפערה |  |
| 7:20.502 | מעי | ) |
| 7:20.503 | לא |  |
| 7:20.504 | כאלה |  |
| 7:20.505 | חלק |  |
| 7:20.506 | יעקב | , |
| 7:20.507 | שנאמר |  |
| 7:20.508 | בהם |  |
| 7:20.509 | ואתר |  |
| 7:20.510 | תהיו |  |
| 7:20.511 | לי |  |
| 7:20.512 | ממלכת |  |
| 7:20.513 | כהנים |  |
| 7:20.514 | וגוי |  |
| 7:20.515 | קדוש | , |
| 7:20.516 | וכל |  |
| 7:20.517 | מי |  |
| 7:20.518 | שמשתמש |  |
| 7:20.519 | במחשבה |  |
| 7:20.520 | או |  |
| 7:20.521 | בדבריו |  |
| 7:20.522 | מענייני |  |
| 7:20.523 | המשגל |  |
| 7:20.524 | אשר |  |
| 7:20.525 | הוא |  |
| 7:20.526 | חרפה |  |
| 7:20.527 | לנו |  |
| 7:20.528 | ביותר |  |
| 7:20.529 | מן |  |
| 7:20.530 | הצריך |  |
| 7:20.531 | לו |  |
| 7:20.532 | כמאמר |  |
| 7:20.533 | רז"ל |  |
| 7:20.534 | מגיד |  |
| 7:20.535 | לאדם |  |
| 7:20.536 | מה |  |
| 7:20.537 | שיחו |  |
| 7:20.538 | אפילו |  |
| 7:20.539 | שיחה |  |
| 7:20.540 | יתירה |  |
| 7:20.541 | שבין |  |
| 7:20.542 | איש |  |
| 7:20.543 | לאשתו |  |
| 7:20.544 | מגידין |  |
| 7:20.545 | לו | ) |
| 7:20.546 | וכ"ש |  |
| 7:20.547 | אם |  |
| 7:20.548 | יאמר |  |
| 7:20.549 | שירים | , |
| 7:20.550 | הנה |  |
| 7:20.551 | נשתמש |  |
| 7:20.552 | בטובה |  |
| 7:20.553 | אשר |  |
| 7:20.554 | גמלו |  |
| 7:20.555 | השי"צ |  |
| 7:20.556 | בדברי |  |
| 7:20.557 | מרי |  |
| 7:20.558 | וסכלות |  |
| 7:20.559 | גמורה |  |
| 7:20.560 | ועבר |  |
| 7:20.561 | על |  |
| 7:20.562 | מצות |  |
| 7:20.563 | הבורא |  |
| 7:20.564 | ית | '. |
| 7:20.565 | והוא |  |
| 7:20.566 | כמי |  |
| 7:20.567 | שנאמר |  |
| 7:20.568 | בהם |  |
| 7:20.569 | וכסף |  |
| 7:20.570 | הרביתי |  |
| 7:20.571 | להם |  |
| 7:20.572 | וזהב |  |
| 7:20.573 | עשו |  |
| 7:20.574 | לבעל | , |
| 7:20.575 | זהו |  |
| 7:20.576 | תורף |  |
| 7:20.577 | דברי |  |
| 7:20.578 | הר | ' |
| 7:20.579 | המורה |  |
| 7:20.580 | ז"ל | . |
| 7:20.581 | גם |  |
| 7:20.582 | הרס"ג |  |
| 7:20.583 | בס | ' |
| 7:20.584 | האמונות |  |
| 7:20.585 | מאמר |  |
| 7:20.586 | עשירי |  |
| 7:20.587 | פרק |  |
| 7:20.588 | ו | ' |
| 7:20.589 | דבר |  |
| 7:20.590 | בגנות |  |
| 7:20.591 | רבוי |  |
| 7:20.592 | המשגל |  |
| 7:20.593 | ולכלוכו |  |
| 7:20.594 | וטנופו |  |
| 7:20.595 | והחרפה |  |
| 7:20.596 | והבוז |  |
| 7:20.597 | אשר |  |
| 7:20.598 | בו | , |
| 7:20.599 | ושאין |  |
| 7:20.600 | התואה |  |
| 7:20.601 | הזאת |  |
| 7:20.602 | באדם |  |
| 7:20.603 | כי |  |
| 7:20.604 | אם |  |
| 7:20.605 | להקים |  |
| 7:20.606 | בה |  |
| 7:20.607 | זרע |  |
| 7:20.608 | כמצות |  |
| 7:20.609 | השי"ת |  |
| 7:20.610 | ואתם |  |
| 7:20.611 | פרו |  |
| 7:20.612 | ורבו |  |
| 7:20.613 | וכו | ' |
| 7:20.614 | ובהצנע |  |
| 7:20.615 | בעת |  |
| 7:20.616 | הראוי |  |
| 7:20.617 | ויעצור |  |
| 7:20.618 | בה |  |
| 7:20.619 | מהרבות |  |
| 7:20.620 | ע"כ | . |
| 7:20.621 | ובעיני |  |
| 7:20.622 | ראיתי |  |
| 7:20.623 | אחד |  |
| 7:20.624 | מהמון |  |
| 7:20.625 | העם |  |
| 7:20.626 | שאל |  |
| 7:20.627 | לתינוק |  |
| 7:20.628 | קטן |  |
| 7:20.629 | ואמר |  |
| 7:20.630 | האם |  |
| 7:20.631 | אביך |  |
| 7:20.632 | אוהב |  |
| 7:20.633 | את |  |
| 7:20.634 | אמך |  |
| 7:20.635 | ומנשקה | , |
| 7:20.636 | וגער |  |
| 7:20.637 | בו |  |
| 7:20.638 | הדיין |  |
| 7:20.639 | כמה"ר |  |
| 7:20.640 | סלי | ' |
| 7:20.641 | צאלח |  |
| 7:20.642 | ז"ל |  |
| 7:20.643 | ואמר |  |
| 7:20.644 | לו |  |
| 7:20.645 | כלך |  |
| 7:20.646 | מגבלות |  |
| 7:20.647 | הפה |  |
| 7:20.648 | כי |  |
| 7:20.649 | ענשו |  |
| 7:20.650 | חמור | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

