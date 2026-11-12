# Translation pack: milhamot-hashem perek-3-08 (Milchamot ha-Shem 3:13–3:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-08`.

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
- **עשתונות איש ירושלים** → "the thoughts of the man of Jerusalem" (note); no
  identification.
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

- **מדה / מדות / מידות**: by context, never "sefirot". Hermeneutical rules: "הי״ג
  מדות דר' ישמעאל ול״ג מדות דר' יוסי הגלילי" → "the thirteen rules of Rabbi
  Ishmael and the thirty-three rules of Rabbi Yose the Galilean" (the number
  as written; the usual thirty-two in a note; note: literally "measures").
  "שלש עשרה מדות של רחמים" (Exod 34:6–7) → "the thirteen attributes of mercy",
  a different thirteen. Of God or the sefirot: "attribute" ("מדת הדין" → "the
  attribute of judgment"). Also "measure", and "quality / trait / virtue" of
  character. מסכת מדות → "tractate Middot".
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
- **קצר אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
- **אין לנו להחמיר** → "we may not be stringent".
- **הגיע חתור** (infinitive after an adjective of motion): the grammar as
  written; the reviewer's reading may go in `readings`.
- **ס' / סימן**: ס' before a quoted verse → "verse"; סימן → "section" in his
  numbered sections; by context, the choice noted.

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


## Milchamot ha-Shem 3:13

Hebrew (Wikisource): יג) למדנו מדברי הגאון כי הרבה ספרים יש שמיחסין אותם על שם גדולי חכמי ישראל (בשקר) בשביל איזה תועלת המגיע להם, ולמדנו ג"כ שכל שאינו מפורש במשנה ובגמרא וראינו שהוא היפך משמעות התלמוד לאו קבלה הוא, ולא סמכינן עליה. ומטעם זה אין פוסקים מן המדרשות במקום שהם חולקים על התלמוד, ואפי' תלמוד ירושלמי אין חוששין לו במקום שהוא חולק על הבבלי אשר כל בית ישראל נשענים עליו בין להקל בין להחמיר כמבואר כל זה בדברי רבותינו הפוסקים בכמה מקומות ועיין לקמן (סימן מ"ב) מה שכתבתי שם בשם הרא"ש, וצא ולמד ממה שכתב ר"ת ן' יחיא הובא בס' רביד הזהב בהל' פסח וז"ל ובזמנינו זה אפס המדע וסתרי תורה נעלמו מעיני כל חי, לא ידע אנוש סדר החכמה הזאת, וכל המחזיק זה יקרה לו מההזק והסכנה חלק רב ( ר"ל הזק וסכנת מינות שאין לו חלק באלהי ישראל). ועל כיוצא בזה הזהירנו רז"ל במופלא ממך אל תדרוש ובמכוסה ממך אל תחקור. ודי לנו במה שהורונו חכמי הקבלה האמתית והמוסכמת מכל צד בפי' התורה ומצותיה. ע"פ המסורה איש מפי איש עד ימות משה והיא הנקראת תורה שבע"פ אשר בגללה עלה משה למרום ועמד שם ארבעים יום ללמוד כללותיה ופרטיה במדות הנמסרות לו מפי הגבורה. ובדברים אלו אנו מחוייבים להתעסק בהם כי הם חיינו ואורך ימינו וא' מהם לא נשגה. אחת דתו להמית אשר ימרה את פיהם וכל העובר על דברי חכמים חייב מיתה. אך יתר הרמזים אשר בהם לא כלו ימיהם חכמי הקבלה עם היותם מסתרי תורה ומעניני המושכלות לא נצטוינו עליהם. כי אלו נשארו להתעסק בהם יחידי סגולה בזמנים אשר היו לפנינו וכו'. ולא כל הרוצה ליטול את השם יטול. כי דברים שהם כבשונו שלעולם הם. ובזמנינו זה תמה זכות אבות, ואין מחזיק בצדק זאת החכמה: אדרבה שהם הורסים פינות התורה ועמודיה יתפלצון מקול מחצצים העברים העורים לא ידעו דרך ישכן אור, ימששו כעורים, וילאו למצוא הפתח. ובמקום שחשבו להרויח פתע ישברו ואין מרפא. ורבים מעמי הארץ מתיהרים באמרם סוד ה' לנו, מורשה לנו המים העליונים, לא לכם בעלי התלמוד ההולכים בחשך. ולבם בל עמם, טח מראות עיניהם, והדבר אין בהם, חכמים להרע לנפשם. כי החזיקו במה שקצרה יד שכלם. ורוצים לעלות במעלות אין ספק בידם להגיע אל התכלית. ויעלו חרס בידם להתגרד בו. ויהיו להם למוקש וכו' כי תצא מהם תקלה לפרוץ גדר התורה ולהרוס חומותיה. יתפאר במה שלא יבין. ויטעה בהתחלות התורה ושרשיה ויבוא לידי מינות. ויתרחק ממחוז חפצו והיא קרבת אלהים, נוח לו שלא נברא, כל שכן לשנות מדרש ולבטל הלכה עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 3:13.1 | יג | ) |
| 3:13.2 | למדנו |  |
| 3:13.3 | מדברי |  |
| 3:13.4 | הגאון |  |
| 3:13.5 | כי |  |
| 3:13.6 | הרבה |  |
| 3:13.7 | ספרים |  |
| 3:13.8 | יש |  |
| 3:13.9 | שמיחסין |  |
| 3:13.10 | אותם |  |
| 3:13.11 | על |  |
| 3:13.12 | שם |  |
| 3:13.13 | גדולי |  |
| 3:13.14 | חכמי |  |
| 3:13.15 | ישראל |  |
| 3:13.16 | בשקר | ) |
| 3:13.17 | בשביל |  |
| 3:13.18 | איזה |  |
| 3:13.19 | תועלת |  |
| 3:13.20 | המגיע |  |
| 3:13.21 | להם | , |
| 3:13.22 | ולמדנו |  |
| 3:13.23 | ג"כ |  |
| 3:13.24 | שכל |  |
| 3:13.25 | שאינו |  |
| 3:13.26 | מפורש |  |
| 3:13.27 | במשנה |  |
| 3:13.28 | ובגמרא |  |
| 3:13.29 | וראינו |  |
| 3:13.30 | שהוא |  |
| 3:13.31 | היפך |  |
| 3:13.32 | משמעות |  |
| 3:13.33 | התלמוד |  |
| 3:13.34 | לאו |  |
| 3:13.35 | קבלה |  |
| 3:13.36 | הוא | , |
| 3:13.37 | ולא |  |
| 3:13.38 | סמכינן |  |
| 3:13.39 | עליה | . |
| 3:13.40 | ומטעם |  |
| 3:13.41 | זה |  |
| 3:13.42 | אין |  |
| 3:13.43 | פוסקים |  |
| 3:13.44 | מן |  |
| 3:13.45 | המדרשות |  |
| 3:13.46 | במקום |  |
| 3:13.47 | שהם |  |
| 3:13.48 | חולקים |  |
| 3:13.49 | על |  |
| 3:13.50 | התלמוד | , |
| 3:13.51 | ואפי | ' |
| 3:13.52 | תלמוד |  |
| 3:13.53 | ירושלמי |  |
| 3:13.54 | אין |  |
| 3:13.55 | חוששין |  |
| 3:13.56 | לו |  |
| 3:13.57 | במקום |  |
| 3:13.58 | שהוא |  |
| 3:13.59 | חולק |  |
| 3:13.60 | על |  |
| 3:13.61 | הבבלי |  |
| 3:13.62 | אשר |  |
| 3:13.63 | כל |  |
| 3:13.64 | בית |  |
| 3:13.65 | ישראל |  |
| 3:13.66 | נשענים |  |
| 3:13.67 | עליו |  |
| 3:13.68 | בין |  |
| 3:13.69 | להקל |  |
| 3:13.70 | בין |  |
| 3:13.71 | להחמיר |  |
| 3:13.72 | כמבואר |  |
| 3:13.73 | כל |  |
| 3:13.74 | זה |  |
| 3:13.75 | בדברי |  |
| 3:13.76 | רבותינו |  |
| 3:13.77 | הפוסקים |  |
| 3:13.78 | בכמה |  |
| 3:13.79 | מקומות |  |
| 3:13.80 | ועיין |  |
| 3:13.81 | לקמן |  |
| 3:13.82 | סימן |  |
| 3:13.83 | מ"ב | ) |
| 3:13.84 | מה |  |
| 3:13.85 | שכתבתי |  |
| 3:13.86 | שם |  |
| 3:13.87 | בשם |  |
| 3:13.88 | הרא"ש | , |
| 3:13.89 | וצא |  |
| 3:13.90 | ולמד |  |
| 3:13.91 | ממה |  |
| 3:13.92 | שכתב |  |
| 3:13.93 | ר"ת |  |
| 3:13.94 | ן | ' |
| 3:13.95 | יחיא |  |
| 3:13.96 | הובא |  |
| 3:13.97 | בס | ' |
| 3:13.98 | רביד |  |
| 3:13.99 | הזהב |  |
| 3:13.100 | בהל | ' |
| 3:13.101 | פסח |  |
| 3:13.102 | וז"ל |  |
| 3:13.103 | ובזמנינו |  |
| 3:13.104 | זה |  |
| 3:13.105 | אפס |  |
| 3:13.106 | המדע |  |
| 3:13.107 | וסתרי |  |
| 3:13.108 | תורה |  |
| 3:13.109 | נעלמו |  |
| 3:13.110 | מעיני |  |
| 3:13.111 | כל |  |
| 3:13.112 | חי | , |
| 3:13.113 | לא |  |
| 3:13.114 | ידע |  |
| 3:13.115 | אנוש |  |
| 3:13.116 | סדר |  |
| 3:13.117 | החכמה |  |
| 3:13.118 | הזאת | , |
| 3:13.119 | וכל |  |
| 3:13.120 | המחזיק |  |
| 3:13.121 | זה |  |
| 3:13.122 | יקרה |  |
| 3:13.123 | לו |  |
| 3:13.124 | מההזק |  |
| 3:13.125 | והסכנה |  |
| 3:13.126 | חלק |  |
| 3:13.127 | רב | ( |
| 3:13.128 | ר"ל |  |
| 3:13.129 | הזק |  |
| 3:13.130 | וסכנת |  |
| 3:13.131 | מינות |  |
| 3:13.132 | שאין |  |
| 3:13.133 | לו |  |
| 3:13.134 | חלק |  |
| 3:13.135 | באלהי |  |
| 3:13.136 | ישראל | ). |
| 3:13.137 | ועל |  |
| 3:13.138 | כיוצא |  |
| 3:13.139 | בזה |  |
| 3:13.140 | הזהירנו |  |
| 3:13.141 | רז"ל |  |
| 3:13.142 | במופלא |  |
| 3:13.143 | ממך |  |
| 3:13.144 | אל |  |
| 3:13.145 | תדרוש |  |
| 3:13.146 | ובמכוסה |  |
| 3:13.147 | ממך |  |
| 3:13.148 | אל |  |
| 3:13.149 | תחקור | . |
| 3:13.150 | ודי |  |
| 3:13.151 | לנו |  |
| 3:13.152 | במה |  |
| 3:13.153 | שהורונו |  |
| 3:13.154 | חכמי |  |
| 3:13.155 | הקבלה |  |
| 3:13.156 | האמתית |  |
| 3:13.157 | והמוסכמת |  |
| 3:13.158 | מכל |  |
| 3:13.159 | צד |  |
| 3:13.160 | בפי | ' |
| 3:13.161 | התורה |  |
| 3:13.162 | ומצותיה | . |
| 3:13.163 | ע"פ |  |
| 3:13.164 | המסורה |  |
| 3:13.165 | איש |  |
| 3:13.166 | מפי |  |
| 3:13.167 | איש |  |
| 3:13.168 | עד |  |
| 3:13.169 | ימות |  |
| 3:13.170 | משה |  |
| 3:13.171 | והיא |  |
| 3:13.172 | הנקראת |  |
| 3:13.173 | תורה |  |
| 3:13.174 | שבע"פ |  |
| 3:13.175 | אשר |  |
| 3:13.176 | בגללה |  |
| 3:13.177 | עלה |  |
| 3:13.178 | משה |  |
| 3:13.179 | למרום |  |
| 3:13.180 | ועמד |  |
| 3:13.181 | שם |  |
| 3:13.182 | ארבעים |  |
| 3:13.183 | יום |  |
| 3:13.184 | ללמוד |  |
| 3:13.185 | כללותיה |  |
| 3:13.186 | ופרטיה |  |
| 3:13.187 | במדות |  |
| 3:13.188 | הנמסרות |  |
| 3:13.189 | לו |  |
| 3:13.190 | מפי |  |
| 3:13.191 | הגבורה | . |
| 3:13.192 | ובדברים |  |
| 3:13.193 | אלו |  |
| 3:13.194 | אנו |  |
| 3:13.195 | מחוייבים |  |
| 3:13.196 | להתעסק |  |
| 3:13.197 | בהם |  |
| 3:13.198 | כי |  |
| 3:13.199 | הם |  |
| 3:13.200 | חיינו |  |
| 3:13.201 | ואורך |  |
| 3:13.202 | ימינו |  |
| 3:13.203 | וא | ' |
| 3:13.204 | מהם |  |
| 3:13.205 | לא |  |
| 3:13.206 | נשגה | . |
| 3:13.207 | אחת |  |
| 3:13.208 | דתו |  |
| 3:13.209 | להמית |  |
| 3:13.210 | אשר |  |
| 3:13.211 | ימרה |  |
| 3:13.212 | את |  |
| 3:13.213 | פיהם |  |
| 3:13.214 | וכל |  |
| 3:13.215 | העובר |  |
| 3:13.216 | על |  |
| 3:13.217 | דברי |  |
| 3:13.218 | חכמים |  |
| 3:13.219 | חייב |  |
| 3:13.220 | מיתה | . |
| 3:13.221 | אך |  |
| 3:13.222 | יתר |  |
| 3:13.223 | הרמזים |  |
| 3:13.224 | אשר |  |
| 3:13.225 | בהם |  |
| 3:13.226 | לא |  |
| 3:13.227 | כלו |  |
| 3:13.228 | ימיהם |  |
| 3:13.229 | חכמי |  |
| 3:13.230 | הקבלה |  |
| 3:13.231 | עם |  |
| 3:13.232 | היותם |  |
| 3:13.233 | מסתרי |  |
| 3:13.234 | תורה |  |
| 3:13.235 | ומעניני |  |
| 3:13.236 | המושכלות |  |
| 3:13.237 | לא |  |
| 3:13.238 | נצטוינו |  |
| 3:13.239 | עליהם | . |
| 3:13.240 | כי |  |
| 3:13.241 | אלו |  |
| 3:13.242 | נשארו |  |
| 3:13.243 | להתעסק |  |
| 3:13.244 | בהם |  |
| 3:13.245 | יחידי |  |
| 3:13.246 | סגולה |  |
| 3:13.247 | בזמנים |  |
| 3:13.248 | אשר |  |
| 3:13.249 | היו |  |
| 3:13.250 | לפנינו |  |
| 3:13.251 | וכו | '. |
| 3:13.252 | ולא |  |
| 3:13.253 | כל |  |
| 3:13.254 | הרוצה |  |
| 3:13.255 | ליטול |  |
| 3:13.256 | את |  |
| 3:13.257 | השם |  |
| 3:13.258 | יטול | . |
| 3:13.259 | כי |  |
| 3:13.260 | דברים |  |
| 3:13.261 | שהם |  |
| 3:13.262 | כבשונו |  |
| 3:13.263 | שלעולם |  |
| 3:13.264 | הם | . |
| 3:13.265 | ובזמנינו |  |
| 3:13.266 | זה |  |
| 3:13.267 | תמה |  |
| 3:13.268 | זכות |  |
| 3:13.269 | אבות | , |
| 3:13.270 | ואין |  |
| 3:13.271 | מחזיק |  |
| 3:13.272 | בצדק |  |
| 3:13.273 | זאת |  |
| 3:13.274 | החכמה | : |
| 3:13.275 | אדרבה |  |
| 3:13.276 | שהם |  |
| 3:13.277 | הורסים |  |
| 3:13.278 | פינות |  |
| 3:13.279 | התורה |  |
| 3:13.280 | ועמודיה |  |
| 3:13.281 | יתפלצון |  |
| 3:13.282 | מקול |  |
| 3:13.283 | מחצצים |  |
| 3:13.284 | העברים |  |
| 3:13.285 | העורים |  |
| 3:13.286 | לא |  |
| 3:13.287 | ידעו |  |
| 3:13.288 | דרך |  |
| 3:13.289 | ישכן |  |
| 3:13.290 | אור | , |
| 3:13.291 | ימששו |  |
| 3:13.292 | כעורים | , |
| 3:13.293 | וילאו |  |
| 3:13.294 | למצוא |  |
| 3:13.295 | הפתח | . |
| 3:13.296 | ובמקום |  |
| 3:13.297 | שחשבו |  |
| 3:13.298 | להרויח |  |
| 3:13.299 | פתע |  |
| 3:13.300 | ישברו |  |
| 3:13.301 | ואין |  |
| 3:13.302 | מרפא | . |
| 3:13.303 | ורבים |  |
| 3:13.304 | מעמי |  |
| 3:13.305 | הארץ |  |
| 3:13.306 | מתיהרים |  |
| 3:13.307 | באמרם |  |
| 3:13.308 | סוד |  |
| 3:13.309 | ה | ' |
| 3:13.310 | לנו | , |
| 3:13.311 | מורשה |  |
| 3:13.312 | לנו |  |
| 3:13.313 | המים |  |
| 3:13.314 | העליונים | , |
| 3:13.315 | לא |  |
| 3:13.316 | לכם |  |
| 3:13.317 | בעלי |  |
| 3:13.318 | התלמוד |  |
| 3:13.319 | ההולכים |  |
| 3:13.320 | בחשך | . |
| 3:13.321 | ולבם |  |
| 3:13.322 | בל |  |
| 3:13.323 | עמם | , |
| 3:13.324 | טח |  |
| 3:13.325 | מראות |  |
| 3:13.326 | עיניהם | , |
| 3:13.327 | והדבר |  |
| 3:13.328 | אין |  |
| 3:13.329 | בהם | , |
| 3:13.330 | חכמים |  |
| 3:13.331 | להרע |  |
| 3:13.332 | לנפשם | . |
| 3:13.333 | כי |  |
| 3:13.334 | החזיקו |  |
| 3:13.335 | במה |  |
| 3:13.336 | שקצרה |  |
| 3:13.337 | יד |  |
| 3:13.338 | שכלם | . |
| 3:13.339 | ורוצים |  |
| 3:13.340 | לעלות |  |
| 3:13.341 | במעלות |  |
| 3:13.342 | אין |  |
| 3:13.343 | ספק |  |
| 3:13.344 | בידם |  |
| 3:13.345 | להגיע |  |
| 3:13.346 | אל |  |
| 3:13.347 | התכלית | . |
| 3:13.348 | ויעלו |  |
| 3:13.349 | חרס |  |
| 3:13.350 | בידם |  |
| 3:13.351 | להתגרד |  |
| 3:13.352 | בו | . |
| 3:13.353 | ויהיו |  |
| 3:13.354 | להם |  |
| 3:13.355 | למוקש |  |
| 3:13.356 | וכו | ' |
| 3:13.357 | כי |  |
| 3:13.358 | תצא |  |
| 3:13.359 | מהם |  |
| 3:13.360 | תקלה |  |
| 3:13.361 | לפרוץ |  |
| 3:13.362 | גדר |  |
| 3:13.363 | התורה |  |
| 3:13.364 | ולהרוס |  |
| 3:13.365 | חומותיה | . |
| 3:13.366 | יתפאר |  |
| 3:13.367 | במה |  |
| 3:13.368 | שלא |  |
| 3:13.369 | יבין | . |
| 3:13.370 | ויטעה |  |
| 3:13.371 | בהתחלות |  |
| 3:13.372 | התורה |  |
| 3:13.373 | ושרשיה |  |
| 3:13.374 | ויבוא |  |
| 3:13.375 | לידי |  |
| 3:13.376 | מינות | . |
| 3:13.377 | ויתרחק |  |
| 3:13.378 | ממחוז |  |
| 3:13.379 | חפצו |  |
| 3:13.380 | והיא |  |
| 3:13.381 | קרבת |  |
| 3:13.382 | אלהים | , |
| 3:13.383 | נוח |  |
| 3:13.384 | לו |  |
| 3:13.385 | שלא |  |
| 3:13.386 | נברא | , |
| 3:13.387 | כל |  |
| 3:13.388 | שכן |  |
| 3:13.389 | לשנות |  |
| 3:13.390 | מדרש |  |
| 3:13.391 | ולבטל |  |
| 3:13.392 | הלכה |  |
| 3:13.393 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

