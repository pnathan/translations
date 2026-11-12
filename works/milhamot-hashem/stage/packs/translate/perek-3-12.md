# Translation pack: milhamot-hashem perek-3-12 (Milchamot ha-Shem 3:19–3:19)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-12.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-12`.

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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
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


## Milchamot ha-Shem 3:19

Hebrew (Wikisource): יט) וכח זה לחקור ולדרוש בתורה על פי קבלת רבותינו בעלי המשנה והתלמוד מסור לנו מרבותינו כמו שאמר התנא הפוך בה והפוך בה. דכלא בה ובה תחזי. וסיב ובלי בה.! כי אם ילמוד האדם ולא יבין לא יתגלה לו הסוד ולא ישיג דבר. כדגרסינן בברכות פ"ק לעולם ירוץ אדם לדבר מצוה ואפי'בשבת שנאמר אחרי ה' ילכו כאריה ישאג וכו' אמר ר' זירא אגרא דפרקא רהטא. (פי' רש"י אגרא דפרקא עיקר קיבול שכר הבריות הרצים לשמוע דרשה מפי חכם הוא שכר המרוצה שהרי רובם אינם מבינים להעמיד גרסא ולמי השמועה לאחר זמן שיקבלו שכר לימוד) אגרא דכלה דוחקא (שבת שלפני הרגל שהכל נאספים לשמוע הלכוה הרגל) אגרא דשמעתא סברא - פי' שהוא יגע וטורח ומחשב להבין טעמו של דבר) ע"כ. וכסנהדרין גרסי' כי בתחבולות תעשה לך מלחמה אמר ר' אחא בר' חנינא אמר ר' אסי אמר ר' יוחנן במי אתה מוצא מלחמתה של תורה. במי שיש בידו חבילות של משנה וכו' ופי' רש"י מלחמתה של תורה הוראותיה ולעמוד על בוריה ועל עיקרה, לא כאדם המפולפל ומחודד ובעל סברא ולא למד משניות וברייתות הרבה. כי מהיכן יתגלה לו הסוג? אלא בבעל משניות הרבה שאם יצטרך לו טעם בכאן ילמדנו מתוך משנה אחרת. או אם יקשה לו דבר יבין מתוך משניות שבידו. הא מני ר' פלוני היא ששמענוהו במקום אחר אומר כן עכ"ל הגמ' ורש"י ז"ל. וגרסי נמי בפ"ק דחגיגה ואף הוא פתח ודרש דברי חכמים כדרבונות וכמסמרות נטועים בעלי אסופות כולם ניתנו מרועה אחד. למה נמשלו דברי תורה מכוונים לב לומדיהן מדרכי מיתה לדרכי חיים. אי מה דרבן זה וכו' בעלי אסופות אלו תלמידי חכמים שיושבין אסופות אסופות ועוסקין בתורה. הללו מטמאין והללו מטהרין. הללו אוסרין, והללו מתירין, הללו פוסלין והללו מכשירין שמא יאמר היאך אני למד תורה מעתה תלמוד לומר כולם ניתנו מרועה אחד, כולן אל אחד נתנן. ופרנס אחד אמרן, מפי אדון כל המעשים ברוך הוא שנ' וידבר אלהים את כל הדברים האלה לאמר, אף אתה עשה אזנך כאפרכסת וקנה לך לך מבין לשמוע דברי המטהרין ודברי המטמאין, דברי האוסרין ודברי המתירין, דברי הפוסלין ודברי המכשירין ע"כ ופי' רש"י ז"ל אל אחד אמרן, אין לך אחד בבני המחלוקת מביא ראיה מתורת אל אחר אלא מתורתו של הקב"ה, ופרנס אחד אמרן אין לך מביא ראיה מדברי נביא הבא לחלוק על משה רבי' ע"ה, עשה אזנך כאפרכסת, מאחר שכולן לבן לשמים עשה אזן שומעת ולמוד ודע דברי כולן, וכשתדע להבחין איזה יבשר קבע הלכה כמותו, עכ"ל, וכן כ' המהרש"א שזה משל לעניינו, שיכניס לאזנו לשמוע הרבה דברי כולן בין המטמאין בין המטהרין אבל יקנה לו לב מבין להבחין מדעות החלוקות איזה יכשר ויקבע לו הלכה כמותו וכפי' רש"י ז"ל עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 3:19.1 | יט | ) |
| 3:19.2 | וכח |  |
| 3:19.3 | זה |  |
| 3:19.4 | לחקור |  |
| 3:19.5 | ולדרוש |  |
| 3:19.6 | בתורה |  |
| 3:19.7 | על |  |
| 3:19.8 | פי |  |
| 3:19.9 | קבלת |  |
| 3:19.10 | רבותינו |  |
| 3:19.11 | בעלי |  |
| 3:19.12 | המשנה |  |
| 3:19.13 | והתלמוד |  |
| 3:19.14 | מסור |  |
| 3:19.15 | לנו |  |
| 3:19.16 | מרבותינו |  |
| 3:19.17 | כמו |  |
| 3:19.18 | שאמר |  |
| 3:19.19 | התנא |  |
| 3:19.20 | הפוך |  |
| 3:19.21 | בה |  |
| 3:19.22 | והפוך |  |
| 3:19.23 | בה | . |
| 3:19.24 | דכלא |  |
| 3:19.25 | בה |  |
| 3:19.26 | ובה |  |
| 3:19.27 | תחזי | . |
| 3:19.28 | וסיב |  |
| 3:19.29 | ובלי |  |
| 3:19.30 | בה | .! |
| 3:19.31 | כי |  |
| 3:19.32 | אם |  |
| 3:19.33 | ילמוד |  |
| 3:19.34 | האדם |  |
| 3:19.35 | ולא |  |
| 3:19.36 | יבין |  |
| 3:19.37 | לא |  |
| 3:19.38 | יתגלה |  |
| 3:19.39 | לו |  |
| 3:19.40 | הסוד |  |
| 3:19.41 | ולא |  |
| 3:19.42 | ישיג |  |
| 3:19.43 | דבר | . |
| 3:19.44 | כדגרסינן |  |
| 3:19.45 | בברכות |  |
| 3:19.46 | פ"ק |  |
| 3:19.47 | לעולם |  |
| 3:19.48 | ירוץ |  |
| 3:19.49 | אדם |  |
| 3:19.50 | לדבר |  |
| 3:19.51 | מצוה |  |
| 3:19.52 | ואפי'בשבת |  |
| 3:19.53 | שנאמר |  |
| 3:19.54 | אחרי |  |
| 3:19.55 | ה | ' |
| 3:19.56 | ילכו |  |
| 3:19.57 | כאריה |  |
| 3:19.58 | ישאג |  |
| 3:19.59 | וכו | ' |
| 3:19.60 | אמר |  |
| 3:19.61 | ר | ' |
| 3:19.62 | זירא |  |
| 3:19.63 | אגרא |  |
| 3:19.64 | דפרקא |  |
| 3:19.65 | רהטא | . |
| 3:19.66 | פי | ' |
| 3:19.67 | רש"י |  |
| 3:19.68 | אגרא |  |
| 3:19.69 | דפרקא |  |
| 3:19.70 | עיקר |  |
| 3:19.71 | קיבול |  |
| 3:19.72 | שכר |  |
| 3:19.73 | הבריות |  |
| 3:19.74 | הרצים |  |
| 3:19.75 | לשמוע |  |
| 3:19.76 | דרשה |  |
| 3:19.77 | מפי |  |
| 3:19.78 | חכם |  |
| 3:19.79 | הוא |  |
| 3:19.80 | שכר |  |
| 3:19.81 | המרוצה |  |
| 3:19.82 | שהרי |  |
| 3:19.83 | רובם |  |
| 3:19.84 | אינם |  |
| 3:19.85 | מבינים |  |
| 3:19.86 | להעמיד |  |
| 3:19.87 | גרסא |  |
| 3:19.88 | ולמי |  |
| 3:19.89 | השמועה |  |
| 3:19.90 | לאחר |  |
| 3:19.91 | זמן |  |
| 3:19.92 | שיקבלו |  |
| 3:19.93 | שכר |  |
| 3:19.94 | לימוד | ) |
| 3:19.95 | אגרא |  |
| 3:19.96 | דכלה |  |
| 3:19.97 | דוחקא |  |
| 3:19.98 | שבת |  |
| 3:19.99 | שלפני |  |
| 3:19.100 | הרגל |  |
| 3:19.101 | שהכל |  |
| 3:19.102 | נאספים |  |
| 3:19.103 | לשמוע |  |
| 3:19.104 | הלכוה |  |
| 3:19.105 | הרגל | ) |
| 3:19.106 | אגרא |  |
| 3:19.107 | דשמעתא |  |
| 3:19.108 | סברא | - |
| 3:19.109 | פי | ' |
| 3:19.110 | שהוא |  |
| 3:19.111 | יגע |  |
| 3:19.112 | וטורח |  |
| 3:19.113 | ומחשב |  |
| 3:19.114 | להבין |  |
| 3:19.115 | טעמו |  |
| 3:19.116 | של |  |
| 3:19.117 | דבר | ) |
| 3:19.118 | ע"כ | . |
| 3:19.119 | וכסנהדרין |  |
| 3:19.120 | גרסי | ' |
| 3:19.121 | כי |  |
| 3:19.122 | בתחבולות |  |
| 3:19.123 | תעשה |  |
| 3:19.124 | לך |  |
| 3:19.125 | מלחמה |  |
| 3:19.126 | אמר |  |
| 3:19.127 | ר | ' |
| 3:19.128 | אחא |  |
| 3:19.129 | בר | ' |
| 3:19.130 | חנינא |  |
| 3:19.131 | אמר |  |
| 3:19.132 | ר | ' |
| 3:19.133 | אסי |  |
| 3:19.134 | אמר |  |
| 3:19.135 | ר | ' |
| 3:19.136 | יוחנן |  |
| 3:19.137 | במי |  |
| 3:19.138 | אתה |  |
| 3:19.139 | מוצא |  |
| 3:19.140 | מלחמתה |  |
| 3:19.141 | של |  |
| 3:19.142 | תורה | . |
| 3:19.143 | במי |  |
| 3:19.144 | שיש |  |
| 3:19.145 | בידו |  |
| 3:19.146 | חבילות |  |
| 3:19.147 | של |  |
| 3:19.148 | משנה |  |
| 3:19.149 | וכו | ' |
| 3:19.150 | ופי | ' |
| 3:19.151 | רש"י |  |
| 3:19.152 | מלחמתה |  |
| 3:19.153 | של |  |
| 3:19.154 | תורה |  |
| 3:19.155 | הוראותיה |  |
| 3:19.156 | ולעמוד |  |
| 3:19.157 | על |  |
| 3:19.158 | בוריה |  |
| 3:19.159 | ועל |  |
| 3:19.160 | עיקרה | , |
| 3:19.161 | לא |  |
| 3:19.162 | כאדם |  |
| 3:19.163 | המפולפל |  |
| 3:19.164 | ומחודד |  |
| 3:19.165 | ובעל |  |
| 3:19.166 | סברא |  |
| 3:19.167 | ולא |  |
| 3:19.168 | למד |  |
| 3:19.169 | משניות |  |
| 3:19.170 | וברייתות |  |
| 3:19.171 | הרבה | . |
| 3:19.172 | כי |  |
| 3:19.173 | מהיכן |  |
| 3:19.174 | יתגלה |  |
| 3:19.175 | לו |  |
| 3:19.176 | הסוג | ? |
| 3:19.177 | אלא |  |
| 3:19.178 | בבעל |  |
| 3:19.179 | משניות |  |
| 3:19.180 | הרבה |  |
| 3:19.181 | שאם |  |
| 3:19.182 | יצטרך |  |
| 3:19.183 | לו |  |
| 3:19.184 | טעם |  |
| 3:19.185 | בכאן |  |
| 3:19.186 | ילמדנו |  |
| 3:19.187 | מתוך |  |
| 3:19.188 | משנה |  |
| 3:19.189 | אחרת | . |
| 3:19.190 | או |  |
| 3:19.191 | אם |  |
| 3:19.192 | יקשה |  |
| 3:19.193 | לו |  |
| 3:19.194 | דבר |  |
| 3:19.195 | יבין |  |
| 3:19.196 | מתוך |  |
| 3:19.197 | משניות |  |
| 3:19.198 | שבידו | . |
| 3:19.199 | הא |  |
| 3:19.200 | מני |  |
| 3:19.201 | ר | ' |
| 3:19.202 | פלוני |  |
| 3:19.203 | היא |  |
| 3:19.204 | ששמענוהו |  |
| 3:19.205 | במקום |  |
| 3:19.206 | אחר |  |
| 3:19.207 | אומר |  |
| 3:19.208 | כן |  |
| 3:19.209 | עכ"ל |  |
| 3:19.210 | הגמ | ' |
| 3:19.211 | ורש"י |  |
| 3:19.212 | ז"ל | . |
| 3:19.213 | וגרסי |  |
| 3:19.214 | נמי |  |
| 3:19.215 | בפ"ק |  |
| 3:19.216 | דחגיגה |  |
| 3:19.217 | ואף |  |
| 3:19.218 | הוא |  |
| 3:19.219 | פתח |  |
| 3:19.220 | ודרש |  |
| 3:19.221 | דברי |  |
| 3:19.222 | חכמים |  |
| 3:19.223 | כדרבונות |  |
| 3:19.224 | וכמסמרות |  |
| 3:19.225 | נטועים |  |
| 3:19.226 | בעלי |  |
| 3:19.227 | אסופות |  |
| 3:19.228 | כולם |  |
| 3:19.229 | ניתנו |  |
| 3:19.230 | מרועה |  |
| 3:19.231 | אחד | . |
| 3:19.232 | למה |  |
| 3:19.233 | נמשלו |  |
| 3:19.234 | דברי |  |
| 3:19.235 | תורה |  |
| 3:19.236 | מכוונים |  |
| 3:19.237 | לב |  |
| 3:19.238 | לומדיהן |  |
| 3:19.239 | מדרכי |  |
| 3:19.240 | מיתה |  |
| 3:19.241 | לדרכי |  |
| 3:19.242 | חיים | . |
| 3:19.243 | אי |  |
| 3:19.244 | מה |  |
| 3:19.245 | דרבן |  |
| 3:19.246 | זה |  |
| 3:19.247 | וכו | ' |
| 3:19.248 | בעלי |  |
| 3:19.249 | אסופות |  |
| 3:19.250 | אלו |  |
| 3:19.251 | תלמידי |  |
| 3:19.252 | חכמים |  |
| 3:19.253 | שיושבין |  |
| 3:19.254 | אסופות |  |
| 3:19.255 | אסופות |  |
| 3:19.256 | ועוסקין |  |
| 3:19.257 | בתורה | . |
| 3:19.258 | הללו |  |
| 3:19.259 | מטמאין |  |
| 3:19.260 | והללו |  |
| 3:19.261 | מטהרין | . |
| 3:19.262 | הללו |  |
| 3:19.263 | אוסרין | , |
| 3:19.264 | והללו |  |
| 3:19.265 | מתירין | , |
| 3:19.266 | הללו |  |
| 3:19.267 | פוסלין |  |
| 3:19.268 | והללו |  |
| 3:19.269 | מכשירין |  |
| 3:19.270 | שמא |  |
| 3:19.271 | יאמר |  |
| 3:19.272 | היאך |  |
| 3:19.273 | אני |  |
| 3:19.274 | למד |  |
| 3:19.275 | תורה |  |
| 3:19.276 | מעתה |  |
| 3:19.277 | תלמוד |  |
| 3:19.278 | לומר |  |
| 3:19.279 | כולם |  |
| 3:19.280 | ניתנו |  |
| 3:19.281 | מרועה |  |
| 3:19.282 | אחד | , |
| 3:19.283 | כולן |  |
| 3:19.284 | אל |  |
| 3:19.285 | אחד |  |
| 3:19.286 | נתנן | . |
| 3:19.287 | ופרנס |  |
| 3:19.288 | אחד |  |
| 3:19.289 | אמרן | , |
| 3:19.290 | מפי |  |
| 3:19.291 | אדון |  |
| 3:19.292 | כל |  |
| 3:19.293 | המעשים |  |
| 3:19.294 | ברוך |  |
| 3:19.295 | הוא |  |
| 3:19.296 | שנ | ' |
| 3:19.297 | וידבר |  |
| 3:19.298 | אלהים |  |
| 3:19.299 | את |  |
| 3:19.300 | כל |  |
| 3:19.301 | הדברים |  |
| 3:19.302 | האלה |  |
| 3:19.303 | לאמר | , |
| 3:19.304 | אף |  |
| 3:19.305 | אתה |  |
| 3:19.306 | עשה |  |
| 3:19.307 | אזנך |  |
| 3:19.308 | כאפרכסת |  |
| 3:19.309 | וקנה |  |
| 3:19.310 | לך |  |
| 3:19.311 | לך |  |
| 3:19.312 | מבין |  |
| 3:19.313 | לשמוע |  |
| 3:19.314 | דברי |  |
| 3:19.315 | המטהרין |  |
| 3:19.316 | ודברי |  |
| 3:19.317 | המטמאין | , |
| 3:19.318 | דברי |  |
| 3:19.319 | האוסרין |  |
| 3:19.320 | ודברי |  |
| 3:19.321 | המתירין | , |
| 3:19.322 | דברי |  |
| 3:19.323 | הפוסלין |  |
| 3:19.324 | ודברי |  |
| 3:19.325 | המכשירין |  |
| 3:19.326 | ע"כ |  |
| 3:19.327 | ופי | ' |
| 3:19.328 | רש"י |  |
| 3:19.329 | ז"ל |  |
| 3:19.330 | אל |  |
| 3:19.331 | אחד |  |
| 3:19.332 | אמרן | , |
| 3:19.333 | אין |  |
| 3:19.334 | לך |  |
| 3:19.335 | אחד |  |
| 3:19.336 | בבני |  |
| 3:19.337 | המחלוקת |  |
| 3:19.338 | מביא |  |
| 3:19.339 | ראיה |  |
| 3:19.340 | מתורת |  |
| 3:19.341 | אל |  |
| 3:19.342 | אחר |  |
| 3:19.343 | אלא |  |
| 3:19.344 | מתורתו |  |
| 3:19.345 | של |  |
| 3:19.346 | הקב"ה | , |
| 3:19.347 | ופרנס |  |
| 3:19.348 | אחד |  |
| 3:19.349 | אמרן |  |
| 3:19.350 | אין |  |
| 3:19.351 | לך |  |
| 3:19.352 | מביא |  |
| 3:19.353 | ראיה |  |
| 3:19.354 | מדברי |  |
| 3:19.355 | נביא |  |
| 3:19.356 | הבא |  |
| 3:19.357 | לחלוק |  |
| 3:19.358 | על |  |
| 3:19.359 | משה |  |
| 3:19.360 | רבי | ' |
| 3:19.361 | ע"ה | , |
| 3:19.362 | עשה |  |
| 3:19.363 | אזנך |  |
| 3:19.364 | כאפרכסת | , |
| 3:19.365 | מאחר |  |
| 3:19.366 | שכולן |  |
| 3:19.367 | לבן |  |
| 3:19.368 | לשמים |  |
| 3:19.369 | עשה |  |
| 3:19.370 | אזן |  |
| 3:19.371 | שומעת |  |
| 3:19.372 | ולמוד |  |
| 3:19.373 | ודע |  |
| 3:19.374 | דברי |  |
| 3:19.375 | כולן | , |
| 3:19.376 | וכשתדע |  |
| 3:19.377 | להבחין |  |
| 3:19.378 | איזה |  |
| 3:19.379 | יבשר |  |
| 3:19.380 | קבע |  |
| 3:19.381 | הלכה |  |
| 3:19.382 | כמותו | , |
| 3:19.383 | עכ"ל | , |
| 3:19.384 | וכן |  |
| 3:19.385 | כ | ' |
| 3:19.386 | המהרש"א |  |
| 3:19.387 | שזה |  |
| 3:19.388 | משל |  |
| 3:19.389 | לעניינו | , |
| 3:19.390 | שיכניס |  |
| 3:19.391 | לאזנו |  |
| 3:19.392 | לשמוע |  |
| 3:19.393 | הרבה |  |
| 3:19.394 | דברי |  |
| 3:19.395 | כולן |  |
| 3:19.396 | בין |  |
| 3:19.397 | המטמאין |  |
| 3:19.398 | בין |  |
| 3:19.399 | המטהרין |  |
| 3:19.400 | אבל |  |
| 3:19.401 | יקנה |  |
| 3:19.402 | לו |  |
| 3:19.403 | לב |  |
| 3:19.404 | מבין |  |
| 3:19.405 | להבחין |  |
| 3:19.406 | מדעות |  |
| 3:19.407 | החלוקות |  |
| 3:19.408 | איזה |  |
| 3:19.409 | יכשר |  |
| 3:19.410 | ויקבע |  |
| 3:19.411 | לו |  |
| 3:19.412 | הלכה |  |
| 3:19.413 | כמותו |  |
| 3:19.414 | וכפי | ' |
| 3:19.415 | רש"י |  |
| 3:19.416 | ז"ל |  |
| 3:19.417 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

