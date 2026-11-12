# Review pack: milhamot-hashem perek-3-08 (Milchamot ha-Shem 3:13–3:13)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-08`.

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


### Translation under review (3:13)

English: 13) We learned from the words of the Gaon that there are many books that are ascribed to the names of the great sages of Israel (falsely), for the sake of some benefit that accrues to them, and we learned also that whatever is not stated explicitly in the Mishnah and in the Gemara, and that we see is the opposite of the plain sense of the Talmud, is not a tradition, and we do not rely on it. And for this reason no decisors rule from the midrashim where they disagree with the Talmud, and even the Jerusalem Talmud we do not heed where it disagrees with the Babylonian, on which all the house of Israel rely, both for leniency and for stringency, as all this is made clear in the words of our rabbis the decisors in several places. And see below, §42, what I wrote there in the name of Rabbeinu Asher, And go and learn from what Rabbi Tam ben Yichya wrote, cited in the book Rivid ha-Zahav in the laws of Passover, and these are his words: "In our time knowledge has ceased, and the hidden matters of the Torah have been hidden from the eyes of every living thing; no man knows the order of this wisdom; and whoever holds to this, there befalls him a great share of the harm and the danger (that is to say, the harm and the danger of heresy, in that there is no share for him in the God of Israel). And concerning such things our Rabbis of blessed memory warned us, 'Do not seek what is too wondrous for you, and do not inquire into what is hidden from you.' And it is enough for us what the sages of the true tradition, agreed upon from every side, have taught us in the explanation of the Torah and its commandments. According to the tradition, from man to man, until the days of Moses; and it is called the Oral Torah, on account of which Moses went up on high and stood there forty days to learn its general principles and its particulars by the rules transmitted to him from the mouth of the Power. And with these matters we are bound to occupy ourselves, for they are our life and the length of our days, and of them not one shall we err. His one law is to put to death whoever defies their word, and whoever transgresses the words of the Sages is liable to death. But the remaining allusions in them, which the sages of Kabbalah did not exhaust in their lifetimes, although they are hidden matters of the Torah and matters of the intellect; we were not commanded concerning them. For these remained for the select few to be occupied with them, in the times that were before us, etc. And not everyone who wishes to take up the Name takes it. For these are matters that are like its furnace; they are so forever. And in our time the merit of the fathers has ended, and no one holds this wisdom justly: on the contrary, they destroy the corners of the Torah and its pillars; they shudder at the sound of the trumpeters; the blind, who pass over, do not know the way where light dwells; they grope like the blind, and weary themselves seeking the door. And where they thought to gain, suddenly they shall be broken, and there is no healing. And many of the unlearned boast, saying: 'The secret of the LORD is ours; the upper waters are our heritage, not yours, you masters of the Talmud who walk in darkness.' Their heart is not with their people; their eyes are plastered over, and there is no substance in them; they are wise to do harm to themselves. For they hold fast to what the hand of their intellect is too short to reach. And they want to climb the rungs; there is no doubt that it is in their power to reach the goal. And they bring up a potsherd in their hands to scrape with it. And they become a snare to them, etc. For there comes out from them a stumbling-block, to breach the fence of the Torah and to tear down its walls. One boasts of what he does not understand, and errs in the beginnings of the Torah and its roots, and comes to heresy. And he is distanced from the place of his desire, which is the nearness of God; it would have been better for him had he not been created, all the more so to change a midrash and to annul a halakhah." Thus far his words.

Literal: 13) list marker learned-we from-words-of the-Gaon that many books there-are that-ascribe them to the-name-of the-great sages-of Israel falsely for-the-sake-of some benefit that-accrues to-them and-learned-we also that-whatever is-not stated-explicitly in-the-Mishnah and-in-the-Gemara and-we-see that-it the-opposite-of the-plain-sense-of the-Talmud is-not tradition it and-not we-rely on-it and-for this-reason no decisors from the-midrashim where they disagree with the-Talmud and-even the-Talmud Jerusalem not-we heed it where it disagrees with the-Babylonian which all the-house-of Israel rely on-it both for-leniency and for-stringency as-is-made-clear all this in-the-words-of our-rabbis the-decisors in-several places and-see below section 42 what I-wrote there in-the-name-of Rabbeinu Asher and-go-and learn from-what wrote Rabbi-Tam ben Yichya cited in-the-book Rivid ha-Zahav in-the-laws-of Passover and-these are-his-words in-our-time this has-ceased knowledge and-hidden-matters-of the-Torah have-been-hidden from-the-eyes-of every living-thing no knows man the-order-of the-wisdom this and-whoever holds-to this befalls him of-the-harm and-the-danger a-share great that-is-to-say the-harm and-the-danger-of heresy in-that-there-is-no for-him share in-the-God-of Israel and-concerning such things warned-us our-Rabbis-of-blessed-memory what-is-too-wondrous for-you do-not seek and-what-is-hidden from-you do-not inquire and-enough for-us what that-taught-us the-sages-of tradition true agreed-upon from-every side in-the-explanation-of the-Torah and-its-commandments according-to the-tradition man from-the-mouth-of man until the-days-of Moses and-it is-called the Torah of-the-mouth which on-account-of-it Moses went-up on-high and-stood there forty days to-learn its-general-principles and-its-particulars by-the-measures transmitted to-him from-the-mouth-of the-Power and-with these we are-bound to-occupy-ourselves with-them for they our-life and-the-length-of our-days and-one of-them not we-shall-err his-one law to-put-to-death who defies the-mouth-of-them and-whoever transgresses against the-words-of the-Sages is-liable death but the-remaining allusions which in-them not exhausted their-days the-sages-of Kabbalah although they-are hidden-matters-of the-Torah and-matters-of the-intellect we were-not commanded concerning-them for these remained to-be-occupied with-them the-select few in-the-times that were before-us etc and-not everyone who-wishes to-take the-Name takes for matters that like-its-furnace forever they are and-in-our-time this has-ended the-merit of-the-fathers and-no-one holds justly this wisdom on-the-contrary they destroy the-corners-of the-Torah and-its-pillars they-shudder at-the-sound-of the-trumpeters the-passers-over the-blind not know the-way where dwells light they-grope like-the-blind and-they-weary seeking the-door and-where they-thought to-gain suddenly they-shall-be-broken and-there-is no healing and-many of-the-unlearned boast saying the-secret-of the-LORD for-us a-heritage for-us the-waters the-upper not for-you the-masters-of the-Talmud who walk in-darkness and-their-heart not with-their-people plastered-over the-sight of-their-eyes and-substance not in-them wise to-do-harm to-themselves for they hold-fast to-what is-short the-hand of-their-intellect and-they-want to-climb up-the-rungs there-is no doubt in-their-power to-reach to the-goal and-they-bring-up a-potsherd in-their-hands to-scrape with-it and-they-become to-them a-snare etc for comes-out from-them a-stumbling-block to-breach the-fence-of the-Torah and-to-tear-down its-walls one-boasts of-what that-he-not understands and-errs in-the-beginnings-of the-Torah and-its-roots and-comes to heresy and-he-is-distanced from-the-place-of his-desire which is the-nearness-of God it-would-have-been-better for-him that-he-not been-created all the-more-so to-change a-midrash and-to-annul a-halakhah thus far his words

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:13.1 | 13) |  |  | Printed list marker, kept as the author's own (decisions, wave 1 rulings). |
| 3:13.2 | we learned |  |  |  |
| 3:13.3 | from the words of |  |  |  |
| 3:13.4 | the Gaon |  |  | The paragraph does not name the Gaon, and the English does not identify him. |
| 3:13.5 | that |  |  |  |
| 3:13.6 | many |  |  |  |
| 3:13.7 | books |  |  |  |
| 3:13.8 | there are |  |  |  |
| 3:13.9 | that are ascribed |  |  |  |
| 3:13.10 | them |  |  |  |
| 3:13.11 | to |  |  |  |
| 3:13.12 | the name of |  |  |  |
| 3:13.13 | the great |  |  |  |
| 3:13.14 | sages of |  |  |  |
| 3:13.15 | Israel |  |  |  |
| 3:13.16 | falsely |  |  |  |
| 3:13.17 | for the sake of |  |  |  |
| 3:13.18 | some |  |  |  |
| 3:13.19 | benefit |  |  |  |
| 3:13.20 | that accrues |  |  |  |
| 3:13.21 | to them |  |  |  |
| 3:13.22 | and we learned |  |  |  |
| 3:13.23 | also |  |  |  |
| 3:13.24 | that whatever |  |  |  |
| 3:13.25 | is not |  |  |  |
| 3:13.26 | stated explicitly |  |  |  |
| 3:13.27 | in the Mishnah |  |  |  |
| 3:13.28 | and in the Gemara |  |  |  |
| 3:13.29 | and we see |  |  |  |
| 3:13.30 | that it |  |  |  |
| 3:13.31 | the opposite of |  |  |  |
| 3:13.32 | the plain sense of |  |  |  |
| 3:13.33 | the Talmud |  |  |  |
| 3:13.34 | is not |  |  |  |
| 3:13.35 | a tradition |  |  | The Hebrew is קבלה, the word for Kabbalah. The author claims it for the Mishnah and Talmud, and here he denies that a teaching outside them is a tradition. |
| 3:13.36 | it |  |  |  |
| 3:13.37 | and we do not |  |  |  |
| 3:13.38 | rely |  |  |  |
| 3:13.39 | on it |  |  |  |
| 3:13.40 | And for |  |  |  |
| 3:13.41 | this reason |  |  |  |
| 3:13.42 | no |  |  |  |
| 3:13.43 | decisors |  |  |  |
| 3:13.44 | from |  |  |  |
| 3:13.45 | the midrashim |  |  |  |
| 3:13.46 | where |  |  |  |
| 3:13.47 | they |  |  |  |
| 3:13.48 | disagree |  |  |  |
| 3:13.49 | with |  |  |  |
| 3:13.50 | the Talmud |  |  |  |
| 3:13.51 | and even |  |  |  |
| 3:13.52 | the Talmud |  |  |  |
| 3:13.53 | Jerusalem |  |  |  |
| 3:13.54 | we do not |  |  |  |
| 3:13.55 | heed |  |  |  |
| 3:13.56 | it |  |  |  |
| 3:13.57 | where |  |  |  |
| 3:13.58 | it |  |  |  |
| 3:13.59 | disagrees |  |  |  |
| 3:13.60 | with |  |  |  |
| 3:13.61 | the Babylonian |  |  |  |
| 3:13.62 | which |  |  |  |
| 3:13.63 | all |  |  |  |
| 3:13.64 | the house of |  |  |  |
| 3:13.65 | Israel |  |  |  |
| 3:13.66 | rely |  |  |  |
| 3:13.67 | on it |  |  |  |
| 3:13.68 | both |  |  |  |
| 3:13.69 | for leniency |  |  |  |
| 3:13.70 | and |  |  |  |
| 3:13.71 | for stringency |  |  |  |
| 3:13.72 | as is made clear |  |  |  |
| 3:13.73 | all |  |  |  |
| 3:13.74 | this |  |  |  |
| 3:13.75 | in the words of |  |  |  |
| 3:13.76 | our rabbis |  |  |  |
| 3:13.77 | the decisors |  |  |  |
| 3:13.78 | in several |  |  |  |
| 3:13.79 | places |  |  |  |
| 3:13.80 | And see |  |  |  |
| 3:13.81 | below |  |  |  |
| 3:13.82 | section |  |  |  |
| 3:13.83 | 42 |  |  |  |
| 3:13.84 | what |  |  |  |
| 3:13.85 | I wrote |  |  |  |
| 3:13.86 | there |  |  |  |
| 3:13.87 | in the name of |  |  |  |
| 3:13.88 | Rabbeinu Asher |  |  | The Hebrew is Rabbeinu Asher, the Rosh. |
| 3:13.89 | And go and |  |  |  |
| 3:13.90 | learn |  |  |  |
| 3:13.91 | from what |  |  |  |
| 3:13.92 | wrote |  |  |  |
| 3:13.93 | Rabbi Tam |  |  | ר"ת is expanded as 'Rabbi Tam', the shorter reading. 'Rabbeinu Tam' is the other expansion, given in apparatus. |
| 3:13.94 | ben |  |  |  |
| 3:13.95 | Yichya |  |  |  |
| 3:13.96 | cited |  |  |  |
| 3:13.97 | in the book |  |  |  |
| 3:13.98 | Rivid |  |  | Rivid ha-Zahav: 'the golden chain'. The title is transliterated; rivid is a chain or necklace. |
| 3:13.99 | ha-Zahav |  |  |  |
| 3:13.100 | in the laws of |  |  |  |
| 3:13.101 | Passover |  |  |  |
| 3:13.102 | and these are his words: |  |  |  |
| 3:13.103 | In our time |  |  |  |
| 3:13.104 | - |  |  |  |
| 3:13.105 | has ceased |  |  |  |
| 3:13.106 | knowledge |  |  |  |
| 3:13.107 | and the hidden matters of |  |  |  |
| 3:13.108 | the Torah |  |  |  |
| 3:13.109 | have been hidden |  |  |  |
| 3:13.110 | from the eyes of |  |  |  |
| 3:13.111 | every |  |  |  |
| 3:13.112 | living thing |  |  |  |
| 3:13.113 | no |  |  |  |
| 3:13.114 | knows |  |  |  |
| 3:13.115 | man |  |  |  |
| 3:13.116 | the order of |  |  |  |
| 3:13.117 | the wisdom |  |  |  |
| 3:13.118 | this |  |  |  |
| 3:13.119 | and whoever |  |  |  |
| 3:13.120 | holds to |  |  |  |
| 3:13.121 | this |  |  |  |
| 3:13.122 | befalls |  |  |  |
| 3:13.123 | him |  |  |  |
| 3:13.124 | of the harm |  |  |  |
| 3:13.125 | and the danger |  |  |  |
| 3:13.126 | a share |  |  |  |
| 3:13.127 | great |  |  |  |
| 3:13.128 | that is to say |  |  | ר"ל: 'that is to say' (רוצה לומר). |
| 3:13.129 | the harm |  |  |  |
| 3:13.130 | and the danger of |  |  |  |
| 3:13.131 | heresy |  |  |  |
| 3:13.132 | in that there is no |  |  |  |
| 3:13.133 | for him |  |  |  |
| 3:13.134 | share |  |  |  |
| 3:13.135 | in the God of |  |  |  |
| 3:13.136 | Israel |  |  |  |
| 3:13.137 | And concerning |  |  |  |
| 3:13.138 | such |  |  |  |
| 3:13.139 | things |  |  |  |
| 3:13.140 | warned us |  |  |  |
| 3:13.141 | our Rabbis of blessed memory |  |  | רז"ל: 'our Rabbis of blessed memory'. |
| 3:13.142 | what is too wondrous |  |  |  |
| 3:13.143 | for you |  |  |  |
| 3:13.144 | do not |  |  |  |
| 3:13.145 | seek |  |  |  |
| 3:13.146 | and what is hidden |  |  |  |
| 3:13.147 | from you |  |  |  |
| 3:13.148 | do not |  |  |  |
| 3:13.149 | inquire into |  |  |  |
| 3:13.150 | And it is enough |  |  |  |
| 3:13.151 | for us |  |  |  |
| 3:13.152 | what |  |  |  |
| 3:13.153 | that taught us |  |  |  |
| 3:13.154 | the sages of |  |  |  |
| 3:13.155 | tradition |  |  | קבלה with האמתית: 'the true tradition' (owner's default, Q5), with the word Kabbalah in the Hebrew. |
| 3:13.156 | true |  |  |  |
| 3:13.157 | agreed upon |  |  |  |
| 3:13.158 | from every |  |  |  |
| 3:13.159 | side |  |  |  |
| 3:13.160 | in the explanation of |  |  |  |
| 3:13.161 | the Torah |  |  |  |
| 3:13.162 | and its commandments |  |  |  |
| 3:13.163 | according to |  |  | ע"פ: 'according to' (על פי). |
| 3:13.164 | the tradition |  |  |  |
| 3:13.165 | man |  |  |  |
| 3:13.166 | from the mouth of |  |  |  |
| 3:13.167 | man |  |  |  |
| 3:13.168 | until |  |  |  |
| 3:13.169 | the days of |  |  |  |
| 3:13.170 | Moses |  |  |  |
| 3:13.171 | and it is |  |  |  |
| 3:13.172 | called the |  |  |  |
| 3:13.173 | Torah |  |  |  |
| 3:13.174 | Oral |  |  | שבע"פ: 'of the mouth'; the Oral Torah. |
| 3:13.175 | which |  |  | The Hebrew is אשר בגללה, literally 'which on account of it'. |
| 3:13.176 | on account of it |  |  |  |
| 3:13.177 | went up |  |  |  |
| 3:13.178 | Moses |  |  |  |
| 3:13.179 | on high |  |  |  |
| 3:13.180 | and stood |  |  |  |
| 3:13.181 | there |  |  |  |
| 3:13.182 | forty |  |  |  |
| 3:13.183 | days |  |  |  |
| 3:13.184 | to learn |  |  |  |
| 3:13.185 | its general principles |  |  |  |
| 3:13.186 | and its particulars |  |  |  |
| 3:13.187 | by the rules |  |  | The Hebrew word is מדות, literally 'measures'. Here it is the rules transmitted to Moses. |
| 3:13.188 | transmitted |  |  |  |
| 3:13.189 | to him |  |  |  |
| 3:13.190 | from the mouth of |  |  |  |
| 3:13.191 | the Power |  |  | הגבורה is 'the Power', a name of God (decisions). |
| 3:13.192 | And with |  |  |  |
| 3:13.193 | these |  |  |  |
| 3:13.194 | we |  |  |  |
| 3:13.195 | are bound |  |  |  |
| 3:13.196 | to occupy ourselves |  |  |  |
| 3:13.197 | with them |  |  |  |
| 3:13.198 | for |  |  |  |
| 3:13.199 | they |  |  |  |
| 3:13.200 | are our life |  |  |  |
| 3:13.201 | and the length of |  |  |  |
| 3:13.202 | our days |  |  |  |
| 3:13.203 | and not one |  |  | וא' is ואחד, 'and one'; it is read with מהם and the negative in 205. |
| 3:13.204 | of them |  |  |  |
| 3:13.205 | not |  |  |  |
| 3:13.206 | we shall err |  |  |  |
| 3:13.207 | His one |  |  | אחת דתו is a set phrase, literally 'his one law'; the English keeps it. |
| 3:13.208 | law |  |  |  |
| 3:13.209 | is to put to death |  |  |  |
| 3:13.210 | whoever |  |  |  |
| 3:13.211 | defies |  |  |  |
| 3:13.212 | - |  |  |  |
| 3:13.213 | their word |  |  | Literally 'their mouth'; here it means what they say. |
| 3:13.214 | and whoever |  |  |  |
| 3:13.215 | transgresses |  |  |  |
| 3:13.216 | against |  |  |  |
| 3:13.217 | the words of |  |  |  |
| 3:13.218 | the Sages |  |  |  |
| 3:13.219 | is liable to |  |  |  |
| 3:13.220 | death |  |  |  |
| 3:13.221 | But |  |  |  |
| 3:13.222 | the remaining |  |  |  |
| 3:13.223 | allusions |  |  |  |
| 3:13.224 | which |  |  |  |
| 3:13.225 | in them |  |  |  |
| 3:13.226 | not |  |  |  |
| 3:13.227 | exhausted |  |  |  |
| 3:13.228 | in their lifetimes |  |  |  |
| 3:13.229 | the sages of |  |  |  |
| 3:13.230 | Kabbalah |  |  |  |
| 3:13.231 | although |  |  |  |
| 3:13.232 | they are |  |  |  |
| 3:13.233 | hidden matters of |  |  |  |
| 3:13.234 | the Torah |  |  |  |
| 3:13.235 | and matters of |  |  |  |
| 3:13.236 | the intellect |  |  |  |
| 3:13.237 | we were not |  |  |  |
| 3:13.238 | commanded |  |  |  |
| 3:13.239 | concerning them |  |  |  |
| 3:13.240 | For |  |  |  |
| 3:13.241 | these |  |  |  |
| 3:13.242 | remained |  |  |  |
| 3:13.243 | to be occupied |  |  |  |
| 3:13.244 | with them |  |  |  |
| 3:13.245 | the select |  |  |  |
| 3:13.246 | few |  |  |  |
| 3:13.247 | in the times |  |  |  |
| 3:13.248 | that |  |  |  |
| 3:13.249 | were |  |  |  |
| 3:13.250 | before us |  |  |  |
| 3:13.251 | etc. |  |  |  |
| 3:13.252 | And not |  |  |  |
| 3:13.253 | everyone |  |  |  |
| 3:13.254 | who wishes |  |  |  |
| 3:13.255 | to take |  |  |  |
| 3:13.256 | - |  |  |  |
| 3:13.257 | the Name |  |  |  |
| 3:13.258 | takes |  |  |  |
| 3:13.259 | For |  |  |  |
| 3:13.260 | matters |  |  |  |
| 3:13.261 | that |  |  |  |
| 3:13.262 | like its furnace |  |  | כבשונו is literally 'like its furnace'. The paragraph does not explain the figure, and the English keeps it. |
| 3:13.263 | forever |  |  |  |
| 3:13.264 | they are |  |  |  |
| 3:13.265 | And in our time |  |  |  |
| 3:13.266 | - |  |  |  |
| 3:13.267 | has ended |  |  |  |
| 3:13.268 | the merit |  |  |  |
| 3:13.269 | of the fathers |  |  |  |
| 3:13.270 | and no one |  |  |  |
| 3:13.271 | holds |  |  |  |
| 3:13.272 | justly |  |  |  |
| 3:13.273 | this |  |  |  |
| 3:13.274 | wisdom |  |  |  |
| 3:13.275 | On the contrary |  |  |  |
| 3:13.276 | they |  |  |  |
| 3:13.277 | destroy |  |  |  |
| 3:13.278 | the corners of |  |  |  |
| 3:13.279 | the Torah |  |  |  |
| 3:13.280 | and its pillars |  |  |  |
| 3:13.281 | they shudder |  |  |  |
| 3:13.282 | at the sound of |  |  |  |
| 3:13.283 | the trumpeters |  |  | מחצצים: 'trumpeters' (blowers of a trumpet). The sense of this clause is uncertain; the English keeps the words as read. |
| 3:13.284 | who pass over |  |  | העברים: 'those who pass over', literally 'the passers'. |
| 3:13.285 | the blind |  |  |  |
| 3:13.286 | do not |  |  |  |
| 3:13.287 | know |  |  |  |
| 3:13.288 | the way |  |  |  |
| 3:13.289 | where dwells |  |  |  |
| 3:13.290 | light |  |  |  |
| 3:13.291 | they grope |  |  |  |
| 3:13.292 | like the blind |  |  |  |
| 3:13.293 | and they weary themselves |  |  |  |
| 3:13.294 | seeking |  |  |  |
| 3:13.295 | the door |  |  |  |
| 3:13.296 | And where |  |  |  |
| 3:13.297 | they thought |  |  |  |
| 3:13.298 | to gain |  |  |  |
| 3:13.299 | suddenly |  |  |  |
| 3:13.300 | they will be broken |  |  |  |
| 3:13.301 | and there is |  |  |  |
| 3:13.302 | no healing |  |  |  |
| 3:13.303 | And many |  |  | The Hebrew is עמי הארץ, 'the unlearned' (decisions). |
| 3:13.304 | of the |  |  |  |
| 3:13.305 | unlearned |  |  |  |
| 3:13.306 | boast |  |  |  |
| 3:13.307 | saying |  |  |  |
| 3:13.308 | The secret of |  |  |  |
| 3:13.309 | the LORD |  |  |  |
| 3:13.310 | is ours |  |  |  |
| 3:13.311 | a heritage |  |  |  |
| 3:13.312 | for us |  |  |  |
| 3:13.313 | the waters |  |  |  |
| 3:13.314 | the upper |  |  |  |
| 3:13.315 | not |  |  |  |
| 3:13.316 | yours |  |  |  |
| 3:13.317 | you masters of |  |  |  |
| 3:13.318 | the Talmud |  |  |  |
| 3:13.319 | who walk |  |  |  |
| 3:13.320 | in darkness |  |  |  |
| 3:13.321 | And their heart |  |  |  |
| 3:13.322 | is not |  |  |  |
| 3:13.323 | with their people |  |  |  |
| 3:13.324 | is plastered over |  |  | טח: 'plastered over' (Isa 44:18, see apparatus). |
| 3:13.325 | the sight |  |  |  |
| 3:13.326 | their eyes |  |  |  |
| 3:13.327 | and substance |  |  |  |
| 3:13.328 | is not |  |  |  |
| 3:13.329 | in them |  |  |  |
| 3:13.330 | wise |  |  |  |
| 3:13.331 | to do harm |  |  |  |
| 3:13.332 | to themselves |  |  |  |
| 3:13.333 | For |  |  |  |
| 3:13.334 | they hold fast |  |  |  |
| 3:13.335 | to what |  |  |  |
| 3:13.336 | is too short |  |  |  |
| 3:13.337 | the hand |  |  |  |
| 3:13.338 | of their intellect |  |  | Literally 'the hand of their intellect is short'. |
| 3:13.339 | And they want |  |  |  |
| 3:13.340 | to climb |  |  |  |
| 3:13.341 | the rungs |  |  |  |
| 3:13.342 | there is no |  |  |  |
| 3:13.343 | doubt |  |  |  |
| 3:13.344 | in their power |  |  |  |
| 3:13.345 | to reach |  |  | The clause reads as affirmative. The irony lies in the next clause (the potsherd, Job 2:8), and the English keeps the words without adding any. |
| 3:13.346 | to |  |  |  |
| 3:13.347 | the goal |  |  |  |
| 3:13.348 | And they bring up |  |  |  |
| 3:13.349 | a potsherd |  |  |  |
| 3:13.350 | in their hands |  |  |  |
| 3:13.351 | to scrape |  |  |  |
| 3:13.352 | with it |  |  |  |
| 3:13.353 | And they become |  |  |  |
| 3:13.354 | to them |  |  |  |
| 3:13.355 | a snare |  |  |  |
| 3:13.356 | etc. |  |  |  |
| 3:13.357 | For |  |  |  |
| 3:13.358 | there comes out |  |  |  |
| 3:13.359 | from them |  |  |  |
| 3:13.360 | a stumbling-block |  |  |  |
| 3:13.361 | to breach |  |  |  |
| 3:13.362 | the fence of |  |  |  |
| 3:13.363 | the Torah |  |  |  |
| 3:13.364 | and to tear down |  |  |  |
| 3:13.365 | its walls |  |  |  |
| 3:13.366 | One boasts |  |  |  |
| 3:13.367 | of what |  |  |  |
| 3:13.368 | that he does not |  |  |  |
| 3:13.369 | understand |  |  |  |
| 3:13.370 | And he errs |  |  |  |
| 3:13.371 | in the beginnings of |  |  |  |
| 3:13.372 | the Torah |  |  |  |
| 3:13.373 | and its roots |  |  |  |
| 3:13.374 | and comes |  |  |  |
| 3:13.375 | to |  |  |  |
| 3:13.376 | heresy |  |  |  |
| 3:13.377 | And he is distanced |  |  |  |
| 3:13.378 | from the place |  |  |  |
| 3:13.379 | of his desire |  |  |  |
| 3:13.380 | which is |  |  |  |
| 3:13.381 | the nearness of |  |  |  |
| 3:13.382 | God |  |  | אלהים: 'God' (decisions). |
| 3:13.383 | it would have been better |  |  |  |
| 3:13.384 | for him |  |  |  |
| 3:13.385 | that he had not |  |  |  |
| 3:13.386 | been created |  |  |  |
| 3:13.387 | All |  |  |  |
| 3:13.388 | the more so |  |  |  |
| 3:13.389 | to change |  |  |  |
| 3:13.390 | a midrash |  |  |  |
| 3:13.391 | and to annul |  |  |  |
| 3:13.392 | a halakhah |  |  |  |
| 3:13.393 | Thus far his words. |  |  |  |

Phrases:
- 3:13.1 | 13) | 13) list marker
- 3:13.2 3:13.3 3:13.4 3:13.5 3:13.6 3:13.7 3:13.8 3:13.9 3:13.10 3:13.11 3:13.12 3:13.13 3:13.14 3:13.15 3:13.16 3:13.17 3:13.18 3:13.19 3:13.20 3:13.21 | We learned from the words of the Gaon that there are many books that are ascribed to the names of the great sages of Israel (falsely), for the sake of some benefit that accrues to them, | learned-we from-words-of the-Gaon that many books there-are that-ascribe them to the-name-of the-great sages-of Israel falsely for-the-sake-of some benefit that-accrues to-them
- 3:13.22 3:13.23 3:13.24 3:13.25 3:13.26 3:13.27 3:13.28 3:13.29 3:13.30 3:13.31 3:13.32 3:13.33 3:13.34 3:13.35 3:13.36 | and we learned also that whatever is not stated explicitly in the Mishnah and in the Gemara, and that we see is the opposite of the plain sense of the Talmud, is not a tradition, | and-learned-we also that-whatever is-not stated-explicitly in-the-Mishnah and-in-the-Gemara and-we-see that-it the-opposite-of the-plain-sense-of the-Talmud is-not tradition it
- 3:13.37 3:13.38 3:13.39 | and we do not rely on it. | and-not we-rely on-it
- 3:13.40 3:13.41 3:13.42 3:13.43 3:13.44 3:13.45 3:13.46 3:13.47 3:13.48 3:13.49 3:13.50 | And for this reason no decisors rule from the midrashim where they disagree with the Talmud, | and-for this-reason no decisors from the-midrashim where they disagree with the-Talmud
- 3:13.51 3:13.52 3:13.53 3:13.54 3:13.55 3:13.56 3:13.57 3:13.58 3:13.59 3:13.60 3:13.61 | and even the Jerusalem Talmud we do not heed where it disagrees with the Babylonian, | and-even the-Talmud Jerusalem not-we heed it where it disagrees with the-Babylonian
- 3:13.62 3:13.63 3:13.64 3:13.65 3:13.66 3:13.67 3:13.68 3:13.69 3:13.70 3:13.71 3:13.72 3:13.73 3:13.74 3:13.75 3:13.76 3:13.77 3:13.78 3:13.79 | on which all the house of Israel rely, both for leniency and for stringency, as all this is made clear in the words of our rabbis the decisors in several places. | which all the-house-of Israel rely on-it both for-leniency and for-stringency as-is-made-clear all this in-the-words-of our-rabbis the-decisors in-several places
- 3:13.80 3:13.81 3:13.82 3:13.83 3:13.84 3:13.85 3:13.86 3:13.87 3:13.88 | And see below, §42, what I wrote there in the name of Rabbeinu Asher, | and-see below section 42 what I-wrote there in-the-name-of Rabbeinu Asher
- 3:13.89 3:13.90 3:13.91 3:13.92 3:13.93 3:13.94 3:13.95 3:13.96 3:13.97 3:13.98 3:13.99 3:13.100 3:13.101 3:13.102 | And go and learn from what Rabbi Tam ben Yichya wrote, cited in the book Rivid ha-Zahav in the laws of Passover, and these are his words: | and-go-and learn from-what wrote Rabbi-Tam ben Yichya cited in-the-book Rivid ha-Zahav in-the-laws-of Passover and-these are-his-words
- 3:13.103 3:13.104 3:13.105 3:13.106 3:13.107 3:13.108 3:13.109 3:13.110 3:13.111 3:13.112 3:13.113 3:13.114 3:13.115 | "In our time knowledge has ceased, and the hidden matters of the Torah have been hidden from the eyes of every living thing; no man | in-our-time this has-ceased knowledge and-hidden-matters-of the-Torah have-been-hidden from-the-eyes-of every living-thing no knows man
- 3:13.116 3:13.117 3:13.118 | knows the order of this wisdom; | the-order-of the-wisdom this
- 3:13.119 3:13.120 3:13.121 3:13.122 3:13.123 3:13.124 3:13.125 3:13.126 3:13.127 | and whoever holds to this, there befalls him a great share of the harm and the danger | and-whoever holds-to this befalls him of-the-harm and-the-danger a-share great
- 3:13.128 3:13.129 3:13.130 3:13.131 3:13.132 3:13.133 3:13.134 3:13.135 3:13.136 | (that is to say, the harm and the danger of heresy, in that there is no share for him in the God of Israel). | that-is-to-say the-harm and-the-danger-of heresy in-that-there-is-no for-him share in-the-God-of Israel
- 3:13.137 3:13.138 3:13.139 3:13.140 3:13.141 3:13.142 3:13.143 3:13.144 3:13.145 3:13.146 3:13.147 3:13.148 3:13.149 | And concerning such things our Rabbis of blessed memory warned us, 'Do not seek what is too wondrous for you, and do not inquire into what is hidden from you.' | and-concerning such things warned-us our-Rabbis-of-blessed-memory what-is-too-wondrous for-you do-not seek and-what-is-hidden from-you do-not inquire
- 3:13.150 3:13.151 3:13.152 3:13.153 3:13.154 3:13.155 3:13.156 3:13.157 3:13.158 3:13.159 3:13.160 3:13.161 3:13.162 | And it is enough for us what the sages of the true tradition, agreed upon from every side, have taught us in the explanation of the Torah and its commandments. | and-enough for-us what that-taught-us the-sages-of tradition true agreed-upon from-every side in-the-explanation-of the-Torah and-its-commandments
- 3:13.163 3:13.164 3:13.165 3:13.166 3:13.167 3:13.168 3:13.169 3:13.170 | According to the tradition, from man to man, until the days of Moses; | according-to the-tradition man from-the-mouth-of man until the-days-of Moses
- 3:13.171 3:13.172 3:13.173 3:13.174 | and it is called the Oral Torah, | and-it is-called the Torah of-the-mouth
- 3:13.175 3:13.176 3:13.177 3:13.178 3:13.179 3:13.180 3:13.181 3:13.182 3:13.183 3:13.184 3:13.185 3:13.186 | on account of which Moses went up on high and stood there forty days to learn its general principles and its particulars | which on-account-of-it Moses went-up on-high and-stood there forty days to-learn its-general-principles and-its-particulars
- 3:13.187 3:13.188 3:13.189 3:13.190 3:13.191 | by the rules transmitted to him from the mouth of the Power. | by-the-measures transmitted to-him from-the-mouth-of the-Power
- 3:13.192 3:13.193 3:13.194 3:13.195 3:13.196 3:13.197 3:13.198 3:13.199 3:13.200 3:13.201 3:13.202 3:13.203 3:13.204 3:13.205 3:13.206 | And with these matters we are bound to occupy ourselves, for they are our life and the length of our days, and of them not one shall we err. | and-with these we are-bound to-occupy-ourselves with-them for they our-life and-the-length-of our-days and-one of-them not we-shall-err
- 3:13.207 3:13.208 3:13.209 3:13.210 3:13.211 3:13.212 3:13.213 3:13.214 3:13.215 3:13.216 3:13.217 3:13.218 3:13.219 3:13.220 | His one law is to put to death whoever defies their word, and whoever transgresses the words of the Sages is liable to death. | his-one law to-put-to-death who defies the-mouth-of-them and-whoever transgresses against the-words-of the-Sages is-liable death
- 3:13.221 3:13.222 3:13.223 3:13.224 3:13.225 3:13.226 3:13.227 3:13.228 3:13.229 3:13.230 | But the remaining allusions in them, which the sages of Kabbalah did not exhaust in their lifetimes, | but the-remaining allusions which in-them not exhausted their-days the-sages-of Kabbalah
- 3:13.231 3:13.232 3:13.233 3:13.234 3:13.235 3:13.236 3:13.237 3:13.238 3:13.239 | although they are hidden matters of the Torah and matters of the intellect; we were not commanded concerning them. | although they-are hidden-matters-of the-Torah and-matters-of the-intellect we were-not commanded concerning-them
- 3:13.240 3:13.241 3:13.242 3:13.243 3:13.244 3:13.245 3:13.246 3:13.247 3:13.248 3:13.249 3:13.250 3:13.251 | For these remained for the select few to be occupied with them, in the times that were before us, etc. | for these remained to-be-occupied with-them the-select few in-the-times that were before-us etc
- 3:13.252 3:13.253 3:13.254 3:13.255 3:13.256 3:13.257 3:13.258 | And not everyone who wishes to take up the Name takes it. | and-not everyone who-wishes to-take the-Name takes
- 3:13.259 3:13.260 3:13.261 3:13.262 3:13.263 3:13.264 | For these are matters that are like its furnace; they are so forever. | for matters that like-its-furnace forever they are
- 3:13.265 3:13.266 3:13.267 3:13.268 3:13.269 3:13.270 3:13.271 3:13.272 3:13.273 3:13.274 | And in our time the merit of the fathers has ended, and no one holds this wisdom justly: | and-in-our-time this has-ended the-merit of-the-fathers and-no-one holds justly this wisdom
- 3:13.275 3:13.276 3:13.277 3:13.278 3:13.279 3:13.280 | on the contrary, they destroy the corners of the Torah and its pillars; | on-the-contrary they destroy the-corners-of the-Torah and-its-pillars
- 3:13.281 3:13.282 3:13.283 3:13.284 3:13.285 3:13.286 3:13.287 3:13.288 3:13.289 3:13.290 | they shudder at the sound of the trumpeters; the blind, who pass over, do not know the way where light dwells; | they-shudder at-the-sound-of the-trumpeters the-passers-over the-blind not know the-way where dwells light
- 3:13.291 3:13.292 3:13.293 3:13.294 3:13.295 | they grope like the blind, and weary themselves seeking the door. | they-grope like-the-blind and-they-weary seeking the-door
- 3:13.296 3:13.297 3:13.298 3:13.299 3:13.300 3:13.301 3:13.302 | And where they thought to gain, suddenly they shall be broken, and there is no healing. | and-where they-thought to-gain suddenly they-shall-be-broken and-there-is no healing
- 3:13.303 3:13.304 3:13.305 3:13.306 3:13.307 3:13.308 3:13.309 3:13.310 | And many of the unlearned boast, saying: 'The secret of the LORD is ours; | and-many of-the-unlearned boast saying the-secret-of the-LORD for-us
- 3:13.311 3:13.312 3:13.313 3:13.314 3:13.315 3:13.316 3:13.317 3:13.318 3:13.319 3:13.320 | the upper waters are our heritage, not yours, you masters of the Talmud who walk in darkness.' | a-heritage for-us the-waters the-upper not for-you the-masters-of the-Talmud who walk in-darkness
- 3:13.321 3:13.322 3:13.323 3:13.324 3:13.325 3:13.326 | Their heart is not with their people; their eyes are plastered over, | and-their-heart not with-their-people plastered-over the-sight of-their-eyes
- 3:13.327 3:13.328 3:13.329 3:13.330 3:13.331 3:13.332 | and there is no substance in them; they are wise to do harm to themselves. | and-substance not in-them wise to-do-harm to-themselves
- 3:13.333 3:13.334 3:13.335 3:13.336 3:13.337 3:13.338 | For they hold fast to what the hand of their intellect is too short to reach. | for they hold-fast to-what is-short the-hand of-their-intellect
- 3:13.339 3:13.340 3:13.341 3:13.342 3:13.343 3:13.344 3:13.345 3:13.346 3:13.347 | And they want to climb the rungs; there is no doubt that it is in their power to reach the goal. | and-they-want to-climb up-the-rungs there-is no doubt in-their-power to-reach to the-goal
- 3:13.348 3:13.349 3:13.350 3:13.351 3:13.352 | And they bring up a potsherd in their hands to scrape with it. | and-they-bring-up a-potsherd in-their-hands to-scrape with-it
- 3:13.353 3:13.354 3:13.355 3:13.356 | And they become a snare to them, etc. | and-they-become to-them a-snare etc
- 3:13.357 3:13.358 3:13.359 3:13.360 3:13.361 3:13.362 3:13.363 3:13.364 3:13.365 | For there comes out from them a stumbling-block, to breach the fence of the Torah and to tear down its walls. | for comes-out from-them a-stumbling-block to-breach the-fence-of the-Torah and-to-tear-down its-walls
- 3:13.366 3:13.367 3:13.368 3:13.369 3:13.370 3:13.371 3:13.372 3:13.373 3:13.374 3:13.375 3:13.376 | One boasts of what he does not understand, and errs in the beginnings of the Torah and its roots, and comes to heresy. | one-boasts of-what that-he-not understands and-errs in-the-beginnings-of the-Torah and-its-roots and-comes to heresy
- 3:13.377 3:13.378 3:13.379 3:13.380 3:13.381 3:13.382 3:13.383 3:13.384 3:13.385 3:13.386 | And he is distanced from the place of his desire, which is the nearness of God; it would have been better for him had he not been created, | and-he-is-distanced from-the-place-of his-desire which is the-nearness-of God it-would-have-been-better for-him that-he-not been-created
- 3:13.387 3:13.388 3:13.389 3:13.390 3:13.391 3:13.392 | all the more so to change a midrash and to annul a halakhah." | all the-more-so to-change a-midrash and-to-annul a-halakhah
- 3:13.393 | Thus far his words. | thus far his words

Grammar notes:
- The paragraph is the author's own argument in its first half. From וז"ל (3:13.102) to עכ"ל (3:13.393) the words are Rabbi Tam's, as cited in Rivid ha-Zahav, and the quotation is marked in double quotation marks.
- ולא סמכינן עליה (3:13.37-39): סמכינן is the Talmudic Aramaic first-person plural of 'to rely', and the English gives it as the plain 'we do not rely on it'.
- The subject of יתפלצון (3:13.281) is not marked. The English takes the destroyers of the Torah's corners as its subject; a reader may take the blind as the subject instead.
- The speech of the unlearned (3:13.307-320) has no introductory formula beyond באמרם. It is marked as direct speech inside the quotation, in single quotation marks.
- אין ספק בידם להגיע (3:13.342-345) reads as an affirmative. The irony is carried by the next clause, the potsherd of 3:13.349-352.

Translator's note: Paragraph 13 of the book's third chapter. The opening 'we learned' frame is the author's; the central quotation is Rabbi Tam's, and it closes with the author's 'Thus far his words'.

Commentary rulings:
- apparatus: Base text: Wikisource, as translated for the unit under Q2 (the 1931 scan is not yet collated for this paragraph).
- apparatus: ר"ת (3:13.93): 'Rabbeinu Tam' is the other expansion; 'Rabbi Tam' is kept as the shorter.
- apparatus: Isa 44:18: טח מראות עיניהם (their eyes plastered over).
- apparatus: Job 38:19: אי זה הדרך ישכן אור; the author writes דרך ישכן אור.
- apparatus: Isa 59:10: ימששו כעורים (they grope like the blind).
- apparatus: Ps 25:14: סוד ה' ליראיו; the author writes סוד ה' לנו.
- apparatus: Prov 6:15: פתע יבוא ... ישבר ואין מרפא; the author writes פתע ישברו ואין מרפא.
- apparatus: Isa 59:1: הנה לא קצרה יד ה'; echoed in קצרה יד שכלם (3:13.336-338).
- apparatus: Job 2:8: ויקח לו חרש להתגרד בו; echoed in חרס בידם להתגרד בו (3:13.349-351).
- apparatus: Jer 4:22: חכמים המה להרע; echoed in חכמים להרע לנפשם (3:13.330-332).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

