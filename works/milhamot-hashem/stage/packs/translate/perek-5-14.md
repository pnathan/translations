# Translation pack: milhamot-hashem perek-5-14 (Milchamot ha-Shem 5:21–5:22)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-14`.

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
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
- **יחוד / ייחוד / ייחדו**: the doctrine → "unity", fixed, not "monotheism" or
  "oneness" (אחדות → "oneness"). The Kabbalistic act → "unification" (יחודים
  → "unifications"). The verb, God as object → "to declare one" ("ליחדו" → "to
  declare Him one"); Kabbalistic entities as object → "to unify"; of
  dedicating a thing to one → "dedicate ... solely to" (note: "unify"). Where
  he plays on the senses: "unity" for the noun, "unify" for the verb, the play
  in a note.
- **שיתוף / שתוף / לשתף / משותף**: associating another with God → fixed,
  "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever associates
  the name of Heaven with something else"). The Maimonidean term for a word said of God and creatures →
  "equivocation / equivocal" ("בשתוף השם בלבד" → "by equivocation of the term
  only"), never "association", Guide I in the note; "shared" only for a thing
  held in common. צירוף: "conjunction" of God; "permutation" in letter
  mysticism.
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
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
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


## Milchamot ha-Shem 5:21

Hebrew (Wikisource): מט) ומעתה אשובה ארעה בשושנים דברי רז"ל בגמרא ובמדרשים המפורסמים שממקום קדוש יצאו ואין בהם שום ספק, אשמור דבריהם היקרים מפז ויהלום באחדותו יתברך אחדות גמור בלי חיבור וצירוף ושיתוף שום נברא או נאצל וידיעת דרכיו ית' על פי התורה והמצוה אשר בעבודה בחר באבות שנאמר (בראשית י"ח י"ט): "כי ידעתיו למען אשר יצוה את בניו ואת ביתו אחריו ושמרו דרך השם" וכו' וכן עשה אברהם אבינו וציוה לבניו יצחק וישמעאל אמונת היחוד שהוא יתברך אחד אמת, ואין אחד אמת זולתו, ושהוא קדמון לכל נברא, ראשון ואין ראשית לראשיתו!

| id | word | punctuation after |
|---|---|---|
| 5:21.1 | מט | ) |
| 5:21.2 | ומעתה |  |
| 5:21.3 | אשובה |  |
| 5:21.4 | ארעה |  |
| 5:21.5 | בשושנים |  |
| 5:21.6 | דברי |  |
| 5:21.7 | רז"ל |  |
| 5:21.8 | בגמרא |  |
| 5:21.9 | ובמדרשים |  |
| 5:21.10 | המפורסמים |  |
| 5:21.11 | שממקום |  |
| 5:21.12 | קדוש |  |
| 5:21.13 | יצאו |  |
| 5:21.14 | ואין |  |
| 5:21.15 | בהם |  |
| 5:21.16 | שום |  |
| 5:21.17 | ספק | , |
| 5:21.18 | אשמור |  |
| 5:21.19 | דבריהם |  |
| 5:21.20 | היקרים |  |
| 5:21.21 | מפז |  |
| 5:21.22 | ויהלום |  |
| 5:21.23 | באחדותו |  |
| 5:21.24 | יתברך |  |
| 5:21.25 | אחדות |  |
| 5:21.26 | גמור |  |
| 5:21.27 | בלי |  |
| 5:21.28 | חיבור |  |
| 5:21.29 | וצירוף |  |
| 5:21.30 | ושיתוף |  |
| 5:21.31 | שום |  |
| 5:21.32 | נברא |  |
| 5:21.33 | או |  |
| 5:21.34 | נאצל |  |
| 5:21.35 | וידיעת |  |
| 5:21.36 | דרכיו |  |
| 5:21.37 | ית | ' |
| 5:21.38 | על |  |
| 5:21.39 | פי |  |
| 5:21.40 | התורה |  |
| 5:21.41 | והמצוה |  |
| 5:21.42 | אשר |  |
| 5:21.43 | בעבודה |  |
| 5:21.44 | בחר |  |
| 5:21.45 | באבות |  |
| 5:21.46 | שנאמר |  |
| 5:21.47 | בראשית |  |
| 5:21.48 | י"ח |  |
| 5:21.49 | י"ט | ): |
| 5:21.50 | כי |  |
| 5:21.51 | ידעתיו |  |
| 5:21.52 | למען |  |
| 5:21.53 | אשר |  |
| 5:21.54 | יצוה |  |
| 5:21.55 | את |  |
| 5:21.56 | בניו |  |
| 5:21.57 | ואת |  |
| 5:21.58 | ביתו |  |
| 5:21.59 | אחריו |  |
| 5:21.60 | ושמרו |  |
| 5:21.61 | דרך |  |
| 5:21.62 | השם | " |
| 5:21.63 | וכו | ' |
| 5:21.64 | וכן |  |
| 5:21.65 | עשה |  |
| 5:21.66 | אברהם |  |
| 5:21.67 | אבינו |  |
| 5:21.68 | וציוה |  |
| 5:21.69 | לבניו |  |
| 5:21.70 | יצחק |  |
| 5:21.71 | וישמעאל |  |
| 5:21.72 | אמונת |  |
| 5:21.73 | היחוד |  |
| 5:21.74 | שהוא |  |
| 5:21.75 | יתברך |  |
| 5:21.76 | אחד |  |
| 5:21.77 | אמת | , |
| 5:21.78 | ואין |  |
| 5:21.79 | אחד |  |
| 5:21.80 | אמת |  |
| 5:21.81 | זולתו | , |
| 5:21.82 | ושהוא |  |
| 5:21.83 | קדמון |  |
| 5:21.84 | לכל |  |
| 5:21.85 | נברא | , |
| 5:21.86 | ראשון |  |
| 5:21.87 | ואין |  |
| 5:21.88 | ראשית |  |
| 5:21.89 | לראשיתו | ! |


## Milchamot ha-Shem 5:22

Hebrew (Wikisource): וזה לשון מדרש הגדול: תרח אבי אברהם היה עובד ע"ז והיה עושה צלמים וכו' וכשמלך נמרוד הרשע באותו זמן ועשה עצמו אלוה וכו' והיו כל האומות באים ומשתחוים לו וכורעין לפניו עד שזרח אורו של עולם אברהם אבינו, וכשנולד אברהם אבינו אמרו לו אצטגניניו קוסמיו וחרטומיו: היום נולד נער שהוא עתיד לירש את העולם ועתיד לאבד את המלכיות ולהוציאך ממלכותך, אם רצונך שלח אצל אביו ופדהו ממנו בבית מלא כסף וזהב ועקרהו מן העולם, והיה תרח יושב ביניהם, אמר להם דומין דבריכם לפרד שאומרים לו טול כור שעורים ונחתוך את ראשך. מה יעשה בשעורים אחר שיחתכו את ראשו? כך זה כשייהרג בנו מה יעשה בדמיו? אמרו לו מתוך דבריך ניכר שנולד לך בן זה, אמר להם לא אכחד מכם בן נולד לי וכבר מת, אמרו לו לא על מת אנו מספרים אלא על חי ואנו רואים מזלו ברקיע, עכשיו לך והבא אותו, מיד הלך תרח ולקח את אברהם ואת מנקתו והטמינם במערה והיו שם שלוש שנים. אחר כך יצא מן המערה ובא אצל אביו והיה לבו משוטט ביצירת העולם ובמאורות להשתחוות להן ולעבדן ולידע איזה מהן הוא אלוה, ראה הלבנה שאורה מבהיק בלילה מסוף העולם ועד סופו ופמליא שלה מרובה, אמר זה אלוה. עברה כל הלילה. בשחרית כיון שזרחה חמה וחשך מאור הלבנה ותשש כחה, אמר זה אלוה. לערב חשכה ויצאה הלבנה והכוכבים והמזלות, אמר בודאי יש אדון לאלו וכו' כיון שנעשה לו נס ויצא בשלום, באותה שעה כיון אברהם אבינו דעתו לשמים ואמר, יתברך שמו של הקב"ה שחמה ולבנה כוכבים ומזלות אינם משתחוים אלא לו. ושמא תאמר רב היה לאברהם אלא זימן לו הקב"ה שתי כליותיו כשני מעינות והיו מוציאין לו חכמה שנאמר הן אמת חפצת בטוחות ובסתום חכמה תודיעני" וכו'.

| id | word | punctuation after |
|---|---|---|
| 5:22.1 | וזה |  |
| 5:22.2 | לשון |  |
| 5:22.3 | מדרש |  |
| 5:22.4 | הגדול | : |
| 5:22.5 | תרח |  |
| 5:22.6 | אבי |  |
| 5:22.7 | אברהם |  |
| 5:22.8 | היה |  |
| 5:22.9 | עובד |  |
| 5:22.10 | ע"ז |  |
| 5:22.11 | והיה |  |
| 5:22.12 | עושה |  |
| 5:22.13 | צלמים |  |
| 5:22.14 | וכו | ' |
| 5:22.15 | וכשמלך |  |
| 5:22.16 | נמרוד |  |
| 5:22.17 | הרשע |  |
| 5:22.18 | באותו |  |
| 5:22.19 | זמן |  |
| 5:22.20 | ועשה |  |
| 5:22.21 | עצמו |  |
| 5:22.22 | אלוה |  |
| 5:22.23 | וכו | ' |
| 5:22.24 | והיו |  |
| 5:22.25 | כל |  |
| 5:22.26 | האומות |  |
| 5:22.27 | באים |  |
| 5:22.28 | ומשתחוים |  |
| 5:22.29 | לו |  |
| 5:22.30 | וכורעין |  |
| 5:22.31 | לפניו |  |
| 5:22.32 | עד |  |
| 5:22.33 | שזרח |  |
| 5:22.34 | אורו |  |
| 5:22.35 | של |  |
| 5:22.36 | עולם |  |
| 5:22.37 | אברהם |  |
| 5:22.38 | אבינו | , |
| 5:22.39 | וכשנולד |  |
| 5:22.40 | אברהם |  |
| 5:22.41 | אבינו |  |
| 5:22.42 | אמרו |  |
| 5:22.43 | לו |  |
| 5:22.44 | אצטגניניו |  |
| 5:22.45 | קוסמיו |  |
| 5:22.46 | וחרטומיו | : |
| 5:22.47 | היום |  |
| 5:22.48 | נולד |  |
| 5:22.49 | נער |  |
| 5:22.50 | שהוא |  |
| 5:22.51 | עתיד |  |
| 5:22.52 | לירש |  |
| 5:22.53 | את |  |
| 5:22.54 | העולם |  |
| 5:22.55 | ועתיד |  |
| 5:22.56 | לאבד |  |
| 5:22.57 | את |  |
| 5:22.58 | המלכיות |  |
| 5:22.59 | ולהוציאך |  |
| 5:22.60 | ממלכותך | , |
| 5:22.61 | אם |  |
| 5:22.62 | רצונך |  |
| 5:22.63 | שלח |  |
| 5:22.64 | אצל |  |
| 5:22.65 | אביו |  |
| 5:22.66 | ופדהו |  |
| 5:22.67 | ממנו |  |
| 5:22.68 | בבית |  |
| 5:22.69 | מלא |  |
| 5:22.70 | כסף |  |
| 5:22.71 | וזהב |  |
| 5:22.72 | ועקרהו |  |
| 5:22.73 | מן |  |
| 5:22.74 | העולם | , |
| 5:22.75 | והיה |  |
| 5:22.76 | תרח |  |
| 5:22.77 | יושב |  |
| 5:22.78 | ביניהם | , |
| 5:22.79 | אמר |  |
| 5:22.80 | להם |  |
| 5:22.81 | דומין |  |
| 5:22.82 | דבריכם |  |
| 5:22.83 | לפרד |  |
| 5:22.84 | שאומרים |  |
| 5:22.85 | לו |  |
| 5:22.86 | טול |  |
| 5:22.87 | כור |  |
| 5:22.88 | שעורים |  |
| 5:22.89 | ונחתוך |  |
| 5:22.90 | את |  |
| 5:22.91 | ראשך | . |
| 5:22.92 | מה |  |
| 5:22.93 | יעשה |  |
| 5:22.94 | בשעורים |  |
| 5:22.95 | אחר |  |
| 5:22.96 | שיחתכו |  |
| 5:22.97 | את |  |
| 5:22.98 | ראשו | ? |
| 5:22.99 | כך |  |
| 5:22.100 | זה |  |
| 5:22.101 | כשייהרג |  |
| 5:22.102 | בנו |  |
| 5:22.103 | מה |  |
| 5:22.104 | יעשה |  |
| 5:22.105 | בדמיו | ? |
| 5:22.106 | אמרו |  |
| 5:22.107 | לו |  |
| 5:22.108 | מתוך |  |
| 5:22.109 | דבריך |  |
| 5:22.110 | ניכר |  |
| 5:22.111 | שנולד |  |
| 5:22.112 | לך |  |
| 5:22.113 | בן |  |
| 5:22.114 | זה | , |
| 5:22.115 | אמר |  |
| 5:22.116 | להם |  |
| 5:22.117 | לא |  |
| 5:22.118 | אכחד |  |
| 5:22.119 | מכם |  |
| 5:22.120 | בן |  |
| 5:22.121 | נולד |  |
| 5:22.122 | לי |  |
| 5:22.123 | וכבר |  |
| 5:22.124 | מת | , |
| 5:22.125 | אמרו |  |
| 5:22.126 | לו |  |
| 5:22.127 | לא |  |
| 5:22.128 | על |  |
| 5:22.129 | מת |  |
| 5:22.130 | אנו |  |
| 5:22.131 | מספרים |  |
| 5:22.132 | אלא |  |
| 5:22.133 | על |  |
| 5:22.134 | חי |  |
| 5:22.135 | ואנו |  |
| 5:22.136 | רואים |  |
| 5:22.137 | מזלו |  |
| 5:22.138 | ברקיע | , |
| 5:22.139 | עכשיו |  |
| 5:22.140 | לך |  |
| 5:22.141 | והבא |  |
| 5:22.142 | אותו | , |
| 5:22.143 | מיד |  |
| 5:22.144 | הלך |  |
| 5:22.145 | תרח |  |
| 5:22.146 | ולקח |  |
| 5:22.147 | את |  |
| 5:22.148 | אברהם |  |
| 5:22.149 | ואת |  |
| 5:22.150 | מנקתו |  |
| 5:22.151 | והטמינם |  |
| 5:22.152 | במערה |  |
| 5:22.153 | והיו |  |
| 5:22.154 | שם |  |
| 5:22.155 | שלוש |  |
| 5:22.156 | שנים | . |
| 5:22.157 | אחר |  |
| 5:22.158 | כך |  |
| 5:22.159 | יצא |  |
| 5:22.160 | מן |  |
| 5:22.161 | המערה |  |
| 5:22.162 | ובא |  |
| 5:22.163 | אצל |  |
| 5:22.164 | אביו |  |
| 5:22.165 | והיה |  |
| 5:22.166 | לבו |  |
| 5:22.167 | משוטט |  |
| 5:22.168 | ביצירת |  |
| 5:22.169 | העולם |  |
| 5:22.170 | ובמאורות |  |
| 5:22.171 | להשתחוות |  |
| 5:22.172 | להן |  |
| 5:22.173 | ולעבדן |  |
| 5:22.174 | ולידע |  |
| 5:22.175 | איזה |  |
| 5:22.176 | מהן |  |
| 5:22.177 | הוא |  |
| 5:22.178 | אלוה | , |
| 5:22.179 | ראה |  |
| 5:22.180 | הלבנה |  |
| 5:22.181 | שאורה |  |
| 5:22.182 | מבהיק |  |
| 5:22.183 | בלילה |  |
| 5:22.184 | מסוף |  |
| 5:22.185 | העולם |  |
| 5:22.186 | ועד |  |
| 5:22.187 | סופו |  |
| 5:22.188 | ופמליא |  |
| 5:22.189 | שלה |  |
| 5:22.190 | מרובה | , |
| 5:22.191 | אמר |  |
| 5:22.192 | זה |  |
| 5:22.193 | אלוה | . |
| 5:22.194 | עברה |  |
| 5:22.195 | כל |  |
| 5:22.196 | הלילה | . |
| 5:22.197 | בשחרית |  |
| 5:22.198 | כיון |  |
| 5:22.199 | שזרחה |  |
| 5:22.200 | חמה |  |
| 5:22.201 | וחשך |  |
| 5:22.202 | מאור |  |
| 5:22.203 | הלבנה |  |
| 5:22.204 | ותשש |  |
| 5:22.205 | כחה | , |
| 5:22.206 | אמר |  |
| 5:22.207 | זה |  |
| 5:22.208 | אלוה | . |
| 5:22.209 | לערב |  |
| 5:22.210 | חשכה |  |
| 5:22.211 | ויצאה |  |
| 5:22.212 | הלבנה |  |
| 5:22.213 | והכוכבים |  |
| 5:22.214 | והמזלות | , |
| 5:22.215 | אמר |  |
| 5:22.216 | בודאי |  |
| 5:22.217 | יש |  |
| 5:22.218 | אדון |  |
| 5:22.219 | לאלו |  |
| 5:22.220 | וכו | ' |
| 5:22.221 | כיון |  |
| 5:22.222 | שנעשה |  |
| 5:22.223 | לו |  |
| 5:22.224 | נס |  |
| 5:22.225 | ויצא |  |
| 5:22.226 | בשלום | , |
| 5:22.227 | באותה |  |
| 5:22.228 | שעה |  |
| 5:22.229 | כיון |  |
| 5:22.230 | אברהם |  |
| 5:22.231 | אבינו |  |
| 5:22.232 | דעתו |  |
| 5:22.233 | לשמים |  |
| 5:22.234 | ואמר | , |
| 5:22.235 | יתברך |  |
| 5:22.236 | שמו |  |
| 5:22.237 | של |  |
| 5:22.238 | הקב"ה |  |
| 5:22.239 | שחמה |  |
| 5:22.240 | ולבנה |  |
| 5:22.241 | כוכבים |  |
| 5:22.242 | ומזלות |  |
| 5:22.243 | אינם |  |
| 5:22.244 | משתחוים |  |
| 5:22.245 | אלא |  |
| 5:22.246 | לו | . |
| 5:22.247 | ושמא |  |
| 5:22.248 | תאמר |  |
| 5:22.249 | רב |  |
| 5:22.250 | היה |  |
| 5:22.251 | לאברהם |  |
| 5:22.252 | אלא |  |
| 5:22.253 | זימן |  |
| 5:22.254 | לו |  |
| 5:22.255 | הקב"ה |  |
| 5:22.256 | שתי |  |
| 5:22.257 | כליותיו |  |
| 5:22.258 | כשני |  |
| 5:22.259 | מעינות |  |
| 5:22.260 | והיו |  |
| 5:22.261 | מוציאין |  |
| 5:22.262 | לו |  |
| 5:22.263 | חכמה |  |
| 5:22.264 | שנאמר |  |
| 5:22.265 | הן |  |
| 5:22.266 | אמת |  |
| 5:22.267 | חפצת |  |
| 5:22.268 | בטוחות |  |
| 5:22.269 | ובסתום |  |
| 5:22.270 | חכמה |  |
| 5:22.271 | תודיעני | " |
| 5:22.272 | וכו | '. |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

