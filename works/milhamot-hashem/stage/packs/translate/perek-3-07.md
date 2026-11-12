# Translation pack: milhamot-hashem perek-3-07 (Milchamot ha-Shem 3:12–3:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-07.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-07`.

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
- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
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
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז / עכו"ם / עבודה אחרת**: עבודה זרה → "idolatry" (note once:
  "strange worship"); the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of
  idols". עכו"ם → "idolaters" (note: "worshippers of stars and
  constellations"), never "gentiles" or "Christians", even of a Christian king
  or the Trinity. עבודה אחרת → "other worship" (note first in a unit). אליל →
  "idol"; האלילה → "the she-idol"; אל זר → "a strange god"; אל אחר → "another
  god"; אלהים אחרים → "other gods".
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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
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


## Milchamot ha-Shem 3:12

Hebrew (Wikisource): יב) הנה כי כן בודאי שכל דבריך הפלגות וגיזומים להבהיל את השוטים והמשוגעים אוהבי הנפלאות והנמנעות ולא בדעת תדבר. וכל דבריך לא בהשכל. ובמרמה אתה ממאן לדעת את ה' כמאמר הנביא ירמיה ע"ה בשם ה' שבתך בתוך מרמה במרמה מאנו דעת אותי נאם ה'. וכבר אמר אחד מן הרמאים שבכם ידענו שהאמת אתכם. אך מה לכם לגלות דבר זה ולפרסם לקצת התלמידים. טוב לכם להניחם בטעותם ממש כדברי הרשע שבארבעה בנים האומר מה העבודה הזאת לכם וכו'. והראיה שאתה טח תפל שהרי פרסמת במכתבך זה שהרב בעל מצרף האמונה הוא ח"ו מין וכופר. וקראת את ספרו משרף האמונה. והוא בתומו הולך לא סר מדברי הזהר ומקדש מלך, וכס"א, והרש"ב, ויושר לבב, וספר הברית ור' האר"י ומהרח"ו בס' עץ חיים ונחלת יוסף וסולתם ימין ושמאל. וכמו שיתבאר לקמן, ותחלה אביא לך כמה ראיות. שאין הקבלה הזאת הלכה למשה מסיני שהרי אמרו רז"ל בעירובין דכל מתניתא דלא מיתניא בי רבי אושעיא לאו מתניתא היא ולא תתיבון מינה, דמשבשתא היא. הרי דכל מתניתא שלא נכנסה לבית מדרשו של רבי חייא או של ר' אושעיא לא מיבעיא דלא הויא הלכה למשה מסיני. אלא דאפילו לאותובי מינה על אמורא לא מותבינן. ובירושלמי עירובין פרק קמא בפסקא דלחיים. וז"ל אמר ליה אחא כן אמר ר' שמעון בן לקיש כל משנה שלא נכנסה לחבורה אין סומכין עליה ע"כ, ועיין לקמן סי' מ"ב. ובקונדריס קדום מחבורי קדמונינו מצאתי כתוב משם ר' סעדיה גאון וז"ל כתאב שיעור קומה לם יקע עליה נקל אמנתא אר ליס הוא פי אל משנה ולא פי אל תלמוד, ולא דליל נסתדל בה הל קאלה רבי ישמעאל עלי' אלחקיקה או גירה קאלה פנסבה אליה כמא תנסב כתיר מן אלכתב אלמפתעלה אלי' בעץ אלעלמא אלמשהורין, עכ"ל. (תרגום מאת הרב המחבר: "ספר שיעור קומה אין בידינו קבלה עליו מחכמי אמונתנו, לפי שאינו לא במשנה ולא בתלמוד, וגם אין ראיה ללמוד ממנה אם אמרו ר' ישמעאל באמת או אחרים זולתו אמרו וייחסו לר' ישמעאל כמו שמייחסים הרבה מן הספרים העשויים לקצת מן החכמים המפורסמים").

| id | word | punctuation after |
|---|---|---|
| 3:12.1 | יב | ) |
| 3:12.2 | הנה |  |
| 3:12.3 | כי |  |
| 3:12.4 | כן |  |
| 3:12.5 | בודאי |  |
| 3:12.6 | שכל |  |
| 3:12.7 | דבריך |  |
| 3:12.8 | הפלגות |  |
| 3:12.9 | וגיזומים |  |
| 3:12.10 | להבהיל |  |
| 3:12.11 | את |  |
| 3:12.12 | השוטים |  |
| 3:12.13 | והמשוגעים |  |
| 3:12.14 | אוהבי |  |
| 3:12.15 | הנפלאות |  |
| 3:12.16 | והנמנעות |  |
| 3:12.17 | ולא |  |
| 3:12.18 | בדעת |  |
| 3:12.19 | תדבר | . |
| 3:12.20 | וכל |  |
| 3:12.21 | דבריך |  |
| 3:12.22 | לא |  |
| 3:12.23 | בהשכל | . |
| 3:12.24 | ובמרמה |  |
| 3:12.25 | אתה |  |
| 3:12.26 | ממאן |  |
| 3:12.27 | לדעת |  |
| 3:12.28 | את |  |
| 3:12.29 | ה | ' |
| 3:12.30 | כמאמר |  |
| 3:12.31 | הנביא |  |
| 3:12.32 | ירמיה |  |
| 3:12.33 | ע"ה |  |
| 3:12.34 | בשם |  |
| 3:12.35 | ה | ' |
| 3:12.36 | שבתך |  |
| 3:12.37 | בתוך |  |
| 3:12.38 | מרמה |  |
| 3:12.39 | במרמה |  |
| 3:12.40 | מאנו |  |
| 3:12.41 | דעת |  |
| 3:12.42 | אותי |  |
| 3:12.43 | נאם |  |
| 3:12.44 | ה | '. |
| 3:12.45 | וכבר |  |
| 3:12.46 | אמר |  |
| 3:12.47 | אחד |  |
| 3:12.48 | מן |  |
| 3:12.49 | הרמאים |  |
| 3:12.50 | שבכם |  |
| 3:12.51 | ידענו |  |
| 3:12.52 | שהאמת |  |
| 3:12.53 | אתכם | . |
| 3:12.54 | אך |  |
| 3:12.55 | מה |  |
| 3:12.56 | לכם |  |
| 3:12.57 | לגלות |  |
| 3:12.58 | דבר |  |
| 3:12.59 | זה |  |
| 3:12.60 | ולפרסם |  |
| 3:12.61 | לקצת |  |
| 3:12.62 | התלמידים | . |
| 3:12.63 | טוב |  |
| 3:12.64 | לכם |  |
| 3:12.65 | להניחם |  |
| 3:12.66 | בטעותם |  |
| 3:12.67 | ממש |  |
| 3:12.68 | כדברי |  |
| 3:12.69 | הרשע |  |
| 3:12.70 | שבארבעה |  |
| 3:12.71 | בנים |  |
| 3:12.72 | האומר |  |
| 3:12.73 | מה |  |
| 3:12.74 | העבודה |  |
| 3:12.75 | הזאת |  |
| 3:12.76 | לכם |  |
| 3:12.77 | וכו | '. |
| 3:12.78 | והראיה |  |
| 3:12.79 | שאתה |  |
| 3:12.80 | טח |  |
| 3:12.81 | תפל |  |
| 3:12.82 | שהרי |  |
| 3:12.83 | פרסמת |  |
| 3:12.84 | במכתבך |  |
| 3:12.85 | זה |  |
| 3:12.86 | שהרב |  |
| 3:12.87 | בעל |  |
| 3:12.88 | מצרף |  |
| 3:12.89 | האמונה |  |
| 3:12.90 | הוא |  |
| 3:12.91 | ח"ו |  |
| 3:12.92 | מין |  |
| 3:12.93 | וכופר | . |
| 3:12.94 | וקראת |  |
| 3:12.95 | את |  |
| 3:12.96 | ספרו |  |
| 3:12.97 | משרף |  |
| 3:12.98 | האמונה | . |
| 3:12.99 | והוא |  |
| 3:12.100 | בתומו |  |
| 3:12.101 | הולך |  |
| 3:12.102 | לא |  |
| 3:12.103 | סר |  |
| 3:12.104 | מדברי |  |
| 3:12.105 | הזהר |  |
| 3:12.106 | ומקדש |  |
| 3:12.107 | מלך | , |
| 3:12.108 | וכס"א | , |
| 3:12.109 | והרש"ב | , |
| 3:12.110 | ויושר |  |
| 3:12.111 | לבב | , |
| 3:12.112 | וספר |  |
| 3:12.113 | הברית |  |
| 3:12.114 | ור | ' |
| 3:12.115 | האר"י |  |
| 3:12.116 | ומהרח"ו |  |
| 3:12.117 | בס | ' |
| 3:12.118 | עץ |  |
| 3:12.119 | חיים |  |
| 3:12.120 | ונחלת |  |
| 3:12.121 | יוסף |  |
| 3:12.122 | וסולתם |  |
| 3:12.123 | ימין |  |
| 3:12.124 | ושמאל | . |
| 3:12.125 | וכמו |  |
| 3:12.126 | שיתבאר |  |
| 3:12.127 | לקמן | , |
| 3:12.128 | ותחלה |  |
| 3:12.129 | אביא |  |
| 3:12.130 | לך |  |
| 3:12.131 | כמה |  |
| 3:12.132 | ראיות | . |
| 3:12.133 | שאין |  |
| 3:12.134 | הקבלה |  |
| 3:12.135 | הזאת |  |
| 3:12.136 | הלכה |  |
| 3:12.137 | למשה |  |
| 3:12.138 | מסיני |  |
| 3:12.139 | שהרי |  |
| 3:12.140 | אמרו |  |
| 3:12.141 | רז"ל |  |
| 3:12.142 | בעירובין |  |
| 3:12.143 | דכל |  |
| 3:12.144 | מתניתא |  |
| 3:12.145 | דלא |  |
| 3:12.146 | מיתניא |  |
| 3:12.147 | בי |  |
| 3:12.148 | רבי |  |
| 3:12.149 | אושעיא |  |
| 3:12.150 | לאו |  |
| 3:12.151 | מתניתא |  |
| 3:12.152 | היא |  |
| 3:12.153 | ולא |  |
| 3:12.154 | תתיבון |  |
| 3:12.155 | מינה | , |
| 3:12.156 | דמשבשתא |  |
| 3:12.157 | היא | . |
| 3:12.158 | הרי |  |
| 3:12.159 | דכל |  |
| 3:12.160 | מתניתא |  |
| 3:12.161 | שלא |  |
| 3:12.162 | נכנסה |  |
| 3:12.163 | לבית |  |
| 3:12.164 | מדרשו |  |
| 3:12.165 | של |  |
| 3:12.166 | רבי |  |
| 3:12.167 | חייא |  |
| 3:12.168 | או |  |
| 3:12.169 | של |  |
| 3:12.170 | ר | ' |
| 3:12.171 | אושעיא |  |
| 3:12.172 | לא |  |
| 3:12.173 | מיבעיא |  |
| 3:12.174 | דלא |  |
| 3:12.175 | הויא |  |
| 3:12.176 | הלכה |  |
| 3:12.177 | למשה |  |
| 3:12.178 | מסיני | . |
| 3:12.179 | אלא |  |
| 3:12.180 | דאפילו |  |
| 3:12.181 | לאותובי |  |
| 3:12.182 | מינה |  |
| 3:12.183 | על |  |
| 3:12.184 | אמורא |  |
| 3:12.185 | לא |  |
| 3:12.186 | מותבינן | . |
| 3:12.187 | ובירושלמי |  |
| 3:12.188 | עירובין |  |
| 3:12.189 | פרק |  |
| 3:12.190 | קמא |  |
| 3:12.191 | בפסקא |  |
| 3:12.192 | דלחיים | . |
| 3:12.193 | וז"ל |  |
| 3:12.194 | אמר |  |
| 3:12.195 | ליה |  |
| 3:12.196 | אחא |  |
| 3:12.197 | כן |  |
| 3:12.198 | אמר |  |
| 3:12.199 | ר | ' |
| 3:12.200 | שמעון |  |
| 3:12.201 | בן |  |
| 3:12.202 | לקיש |  |
| 3:12.203 | כל |  |
| 3:12.204 | משנה |  |
| 3:12.205 | שלא |  |
| 3:12.206 | נכנסה |  |
| 3:12.207 | לחבורה |  |
| 3:12.208 | אין |  |
| 3:12.209 | סומכין |  |
| 3:12.210 | עליה |  |
| 3:12.211 | ע"כ | , |
| 3:12.212 | ועיין |  |
| 3:12.213 | לקמן |  |
| 3:12.214 | סי | ' |
| 3:12.215 | מ"ב | . |
| 3:12.216 | ובקונדריס |  |
| 3:12.217 | קדום |  |
| 3:12.218 | מחבורי |  |
| 3:12.219 | קדמונינו |  |
| 3:12.220 | מצאתי |  |
| 3:12.221 | כתוב |  |
| 3:12.222 | משם |  |
| 3:12.223 | ר | ' |
| 3:12.224 | סעדיה |  |
| 3:12.225 | גאון |  |
| 3:12.226 | וז"ל |  |
| 3:12.227 | כתאב |  |
| 3:12.228 | שיעור |  |
| 3:12.229 | קומה |  |
| 3:12.230 | לם |  |
| 3:12.231 | יקע |  |
| 3:12.232 | עליה |  |
| 3:12.233 | נקל |  |
| 3:12.234 | אמנתא |  |
| 3:12.235 | אר |  |
| 3:12.236 | ליס |  |
| 3:12.237 | הוא |  |
| 3:12.238 | פי |  |
| 3:12.239 | אל |  |
| 3:12.240 | משנה |  |
| 3:12.241 | ולא |  |
| 3:12.242 | פי |  |
| 3:12.243 | אל |  |
| 3:12.244 | תלמוד | , |
| 3:12.245 | ולא |  |
| 3:12.246 | דליל |  |
| 3:12.247 | נסתדל |  |
| 3:12.248 | בה |  |
| 3:12.249 | הל |  |
| 3:12.250 | קאלה |  |
| 3:12.251 | רבי |  |
| 3:12.252 | ישמעאל |  |
| 3:12.253 | עלי | ' |
| 3:12.254 | אלחקיקה |  |
| 3:12.255 | או |  |
| 3:12.256 | גירה |  |
| 3:12.257 | קאלה |  |
| 3:12.258 | פנסבה |  |
| 3:12.259 | אליה |  |
| 3:12.260 | כמא |  |
| 3:12.261 | תנסב |  |
| 3:12.262 | כתיר |  |
| 3:12.263 | מן |  |
| 3:12.264 | אלכתב |  |
| 3:12.265 | אלמפתעלה |  |
| 3:12.266 | אלי | ' |
| 3:12.267 | בעץ |  |
| 3:12.268 | אלעלמא |  |
| 3:12.269 | אלמשהורין | , |
| 3:12.270 | עכ"ל | . |
| 3:12.271 | תרגום |  |
| 3:12.272 | מאת |  |
| 3:12.273 | הרב |  |
| 3:12.274 | המחבר | : |
| 3:12.275 | ספר |  |
| 3:12.276 | שיעור |  |
| 3:12.277 | קומה |  |
| 3:12.278 | אין |  |
| 3:12.279 | בידינו |  |
| 3:12.280 | קבלה |  |
| 3:12.281 | עליו |  |
| 3:12.282 | מחכמי |  |
| 3:12.283 | אמונתנו | , |
| 3:12.284 | לפי |  |
| 3:12.285 | שאינו |  |
| 3:12.286 | לא |  |
| 3:12.287 | במשנה |  |
| 3:12.288 | ולא |  |
| 3:12.289 | בתלמוד | , |
| 3:12.290 | וגם |  |
| 3:12.291 | אין |  |
| 3:12.292 | ראיה |  |
| 3:12.293 | ללמוד |  |
| 3:12.294 | ממנה |  |
| 3:12.295 | אם |  |
| 3:12.296 | אמרו |  |
| 3:12.297 | ר | ' |
| 3:12.298 | ישמעאל |  |
| 3:12.299 | באמת |  |
| 3:12.300 | או |  |
| 3:12.301 | אחרים |  |
| 3:12.302 | זולתו |  |
| 3:12.303 | אמרו |  |
| 3:12.304 | וייחסו |  |
| 3:12.305 | לר | ' |
| 3:12.306 | ישמעאל |  |
| 3:12.307 | כמו |  |
| 3:12.308 | שמייחסים |  |
| 3:12.309 | הרבה |  |
| 3:12.310 | מן |  |
| 3:12.311 | הספרים |  |
| 3:12.312 | העשויים |  |
| 3:12.313 | לקצת |  |
| 3:12.314 | מן |  |
| 3:12.315 | החכמים |  |
| 3:12.316 | המפורסמים | "). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

