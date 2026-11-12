# Translation pack: milhamot-hashem perek-2-09 (Milchamot ha-Shem 2:17–2:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-09`.

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
- **קצר אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
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
- **ריבה / לרבות** → "include", never "multiply".
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **ובפתרון הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **הגשת מנחה** (argued from Gen 43 / 1 Sam 10) → "gift"; "meal offering" only
  in a sacrificial context.
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
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


## Milchamot ha-Shem 2:17

Hebrew (Wikisource): וגדולה מזו שעושים גם כן לזבוח לפני מלך או שר צאן ובקר והיא מעבודות שבפנים שהעושה לאחד ממיני ע"ז אף שאין עבודתה בכך חייב כדתנן בסנהדרין פ' ד' מיתות העובד ע"ז אחד העובד, ואחד המזבח, ואחד המקטר ואחד המנסך, ואחד המשתחוה וכו' ואמרינן בגמ' עלה מאי אחד העובד? אמר ר' ירמיה הכי קאמר אחד העובד כדרכה, ואחד המזבח ואחד המקטר ואחד המנסך ואחד המשתחוה, ואפי' שלא כדרכה. וכן פסקו כל הפוסקים דארבע עבודות חייב עליהן אפי' שלא כדרכה. ודבר זה נהוג עד היום אצל אחינו בני הכפרים והעירות שאם היה לאחד מהם אויב שהוא מבקש רעתו הוא מביא כבש ושוחטו לפני השר שלו להושיעו מאויבו. וקרא נמי כתיב: הקריבהו נא לפחתך וכו'. קדמונינו היה עליהם ג"כ החק הזה בבוא המלך אל המדינה לזבוח לפניו שור ונוטלים את דמי השור מקופה של צבור. ולא שמעו שום מפקפק בזה משום ע"ז, לפי שאין מכוין לעבדו בזה ולקבלו עליו באלוה, רק אות כבוד הוא עושה לו. ויעקב אבינו ע"ה אמר לבניו: קחו מזמרת הארץ בכליכם והורידו לאיש מנחה, ונאמ': ויביאו לו את המנחה אשר בידם הביתה וישתחו לו אפים וכו'. ובשאול בחיר ה' נאמר: ובני בליעל אמרו מה יושיענו זה ולא הביאו לו מנחה. ובנדרים שנינו: נדר בקרבן ואמר לא נדרתי אלא בקרבנות מלכים. אלמא שהשתחויה וזיבוח והגשת מנחה וכל ארבע עבודות פנים מותר לעשותן לאדם שאינו נעבד כהמן, שאם הוא נעבד כהמן אסור, שכן נמנע מרדכי להשתחות להמן מפני שעשה עצמו ע"ז.

| id | word | punctuation after |
|---|---|---|
| 2:17.1 | וגדולה |  |
| 2:17.2 | מזו |  |
| 2:17.3 | שעושים |  |
| 2:17.4 | גם |  |
| 2:17.5 | כן |  |
| 2:17.6 | לזבוח |  |
| 2:17.7 | לפני |  |
| 2:17.8 | מלך |  |
| 2:17.9 | או |  |
| 2:17.10 | שר |  |
| 2:17.11 | צאן |  |
| 2:17.12 | ובקר |  |
| 2:17.13 | והיא |  |
| 2:17.14 | מעבודות |  |
| 2:17.15 | שבפנים |  |
| 2:17.16 | שהעושה |  |
| 2:17.17 | לאחד |  |
| 2:17.18 | ממיני |  |
| 2:17.19 | ע"ז |  |
| 2:17.20 | אף |  |
| 2:17.21 | שאין |  |
| 2:17.22 | עבודתה |  |
| 2:17.23 | בכך |  |
| 2:17.24 | חייב |  |
| 2:17.25 | כדתנן |  |
| 2:17.26 | בסנהדרין |  |
| 2:17.27 | פ | ' |
| 2:17.28 | ד | ' |
| 2:17.29 | מיתות |  |
| 2:17.30 | העובד |  |
| 2:17.31 | ע"ז |  |
| 2:17.32 | אחד |  |
| 2:17.33 | העובד | , |
| 2:17.34 | ואחד |  |
| 2:17.35 | המזבח | , |
| 2:17.36 | ואחד |  |
| 2:17.37 | המקטר |  |
| 2:17.38 | ואחד |  |
| 2:17.39 | המנסך | , |
| 2:17.40 | ואחד |  |
| 2:17.41 | המשתחוה |  |
| 2:17.42 | וכו | ' |
| 2:17.43 | ואמרינן |  |
| 2:17.44 | בגמ | ' |
| 2:17.45 | עלה |  |
| 2:17.46 | מאי |  |
| 2:17.47 | אחד |  |
| 2:17.48 | העובד | ? |
| 2:17.49 | אמר |  |
| 2:17.50 | ר | ' |
| 2:17.51 | ירמיה |  |
| 2:17.52 | הכי |  |
| 2:17.53 | קאמר |  |
| 2:17.54 | אחד |  |
| 2:17.55 | העובד |  |
| 2:17.56 | כדרכה | , |
| 2:17.57 | ואחד |  |
| 2:17.58 | המזבח |  |
| 2:17.59 | ואחד |  |
| 2:17.60 | המקטר |  |
| 2:17.61 | ואחד |  |
| 2:17.62 | המנסך |  |
| 2:17.63 | ואחד |  |
| 2:17.64 | המשתחוה | , |
| 2:17.65 | ואפי | ' |
| 2:17.66 | שלא |  |
| 2:17.67 | כדרכה | . |
| 2:17.68 | וכן |  |
| 2:17.69 | פסקו |  |
| 2:17.70 | כל |  |
| 2:17.71 | הפוסקים |  |
| 2:17.72 | דארבע |  |
| 2:17.73 | עבודות |  |
| 2:17.74 | חייב |  |
| 2:17.75 | עליהן |  |
| 2:17.76 | אפי | ' |
| 2:17.77 | שלא |  |
| 2:17.78 | כדרכה | . |
| 2:17.79 | ודבר |  |
| 2:17.80 | זה |  |
| 2:17.81 | נהוג |  |
| 2:17.82 | עד |  |
| 2:17.83 | היום |  |
| 2:17.84 | אצל |  |
| 2:17.85 | אחינו |  |
| 2:17.86 | בני |  |
| 2:17.87 | הכפרים |  |
| 2:17.88 | והעירות |  |
| 2:17.89 | שאם |  |
| 2:17.90 | היה |  |
| 2:17.91 | לאחד |  |
| 2:17.92 | מהם |  |
| 2:17.93 | אויב |  |
| 2:17.94 | שהוא |  |
| 2:17.95 | מבקש |  |
| 2:17.96 | רעתו |  |
| 2:17.97 | הוא |  |
| 2:17.98 | מביא |  |
| 2:17.99 | כבש |  |
| 2:17.100 | ושוחטו |  |
| 2:17.101 | לפני |  |
| 2:17.102 | השר |  |
| 2:17.103 | שלו |  |
| 2:17.104 | להושיעו |  |
| 2:17.105 | מאויבו | . |
| 2:17.106 | וקרא |  |
| 2:17.107 | נמי |  |
| 2:17.108 | כתיב | : |
| 2:17.109 | הקריבהו |  |
| 2:17.110 | נא |  |
| 2:17.111 | לפחתך |  |
| 2:17.112 | וכו | '. |
| 2:17.113 | קדמונינו |  |
| 2:17.114 | היה |  |
| 2:17.115 | עליהם |  |
| 2:17.116 | ג"כ |  |
| 2:17.117 | החק |  |
| 2:17.118 | הזה |  |
| 2:17.119 | בבוא |  |
| 2:17.120 | המלך |  |
| 2:17.121 | אל |  |
| 2:17.122 | המדינה |  |
| 2:17.123 | לזבוח |  |
| 2:17.124 | לפניו |  |
| 2:17.125 | שור |  |
| 2:17.126 | ונוטלים |  |
| 2:17.127 | את |  |
| 2:17.128 | דמי |  |
| 2:17.129 | השור |  |
| 2:17.130 | מקופה |  |
| 2:17.131 | של |  |
| 2:17.132 | צבור | . |
| 2:17.133 | ולא |  |
| 2:17.134 | שמעו |  |
| 2:17.135 | שום |  |
| 2:17.136 | מפקפק |  |
| 2:17.137 | בזה |  |
| 2:17.138 | משום |  |
| 2:17.139 | ע"ז | , |
| 2:17.140 | לפי |  |
| 2:17.141 | שאין |  |
| 2:17.142 | מכוין |  |
| 2:17.143 | לעבדו |  |
| 2:17.144 | בזה |  |
| 2:17.145 | ולקבלו |  |
| 2:17.146 | עליו |  |
| 2:17.147 | באלוה | , |
| 2:17.148 | רק |  |
| 2:17.149 | אות |  |
| 2:17.150 | כבוד |  |
| 2:17.151 | הוא |  |
| 2:17.152 | עושה |  |
| 2:17.153 | לו | . |
| 2:17.154 | ויעקב |  |
| 2:17.155 | אבינו |  |
| 2:17.156 | ע"ה |  |
| 2:17.157 | אמר |  |
| 2:17.158 | לבניו | : |
| 2:17.159 | קחו |  |
| 2:17.160 | מזמרת |  |
| 2:17.161 | הארץ |  |
| 2:17.162 | בכליכם |  |
| 2:17.163 | והורידו |  |
| 2:17.164 | לאיש |  |
| 2:17.165 | מנחה | , |
| 2:17.166 | ונאמ | ': |
| 2:17.167 | ויביאו |  |
| 2:17.168 | לו |  |
| 2:17.169 | את |  |
| 2:17.170 | המנחה |  |
| 2:17.171 | אשר |  |
| 2:17.172 | בידם |  |
| 2:17.173 | הביתה |  |
| 2:17.174 | וישתחו |  |
| 2:17.175 | לו |  |
| 2:17.176 | אפים |  |
| 2:17.177 | וכו | '. |
| 2:17.178 | ובשאול |  |
| 2:17.179 | בחיר |  |
| 2:17.180 | ה | ' |
| 2:17.181 | נאמר | : |
| 2:17.182 | ובני |  |
| 2:17.183 | בליעל |  |
| 2:17.184 | אמרו |  |
| 2:17.185 | מה |  |
| 2:17.186 | יושיענו |  |
| 2:17.187 | זה |  |
| 2:17.188 | ולא |  |
| 2:17.189 | הביאו |  |
| 2:17.190 | לו |  |
| 2:17.191 | מנחה | . |
| 2:17.192 | ובנדרים |  |
| 2:17.193 | שנינו | : |
| 2:17.194 | נדר |  |
| 2:17.195 | בקרבן |  |
| 2:17.196 | ואמר |  |
| 2:17.197 | לא |  |
| 2:17.198 | נדרתי |  |
| 2:17.199 | אלא |  |
| 2:17.200 | בקרבנות |  |
| 2:17.201 | מלכים | . |
| 2:17.202 | אלמא |  |
| 2:17.203 | שהשתחויה |  |
| 2:17.204 | וזיבוח |  |
| 2:17.205 | והגשת |  |
| 2:17.206 | מנחה |  |
| 2:17.207 | וכל |  |
| 2:17.208 | ארבע |  |
| 2:17.209 | עבודות |  |
| 2:17.210 | פנים |  |
| 2:17.211 | מותר |  |
| 2:17.212 | לעשותן |  |
| 2:17.213 | לאדם |  |
| 2:17.214 | שאינו |  |
| 2:17.215 | נעבד |  |
| 2:17.216 | כהמן | , |
| 2:17.217 | שאם |  |
| 2:17.218 | הוא |  |
| 2:17.219 | נעבד |  |
| 2:17.220 | כהמן |  |
| 2:17.221 | אסור | , |
| 2:17.222 | שכן |  |
| 2:17.223 | נמנע |  |
| 2:17.224 | מרדכי |  |
| 2:17.225 | להשתחות |  |
| 2:17.226 | להמן |  |
| 2:17.227 | מפני |  |
| 2:17.228 | שעשה |  |
| 2:17.229 | עצמו |  |
| 2:17.230 | ע"ז | . |


## Milchamot ha-Shem 2:18

Hebrew (Wikisource): ובס' תשועת ישראל לר' מנשה בן ישראל כ' וז"ל: נדרשתי להודיע מנהגי הכבוד והכניעות אשר יכבדו בם היהודים את ספר התורה, כי צרינו אשר אינם יודעים מה זאת מכעיסים אותנו לאמר כי עובדי אלילים אנחנו, וכי נשתחוה לארון עץ וליריעות עור אשר בתוכה, לכן אערוך תשובתי בסדר הניאות כיד ה' הטובה עלי והוא כל איש ישראל יראה את עצמו מחוייב לקום ולעמוד בהוצאת ספר תורה מארון הקודש עדי תונח על התיבה להראותה לכל העם ולקרות בה. ומצינו כי כן עשו בימי נחמיה ככתוב וכפתחו עמדו כל העם. וזאת יעשו מפני היראה והכבוד אשר יכבדו דברי אלהים חיים הכתובים בה, ומאותו טעם יכופו ראשם בנשאם אותה מהארון עד התיבה, ולא משום דבר אחר מן הטעמים אשר נזכיר.

| id | word | punctuation after |
|---|---|---|
| 2:18.1 | ובס | ' |
| 2:18.2 | תשועת |  |
| 2:18.3 | ישראל |  |
| 2:18.4 | לר | ' |
| 2:18.5 | מנשה |  |
| 2:18.6 | בן |  |
| 2:18.7 | ישראל |  |
| 2:18.8 | כ | ' |
| 2:18.9 | וז"ל | : |
| 2:18.10 | נדרשתי |  |
| 2:18.11 | להודיע |  |
| 2:18.12 | מנהגי |  |
| 2:18.13 | הכבוד |  |
| 2:18.14 | והכניעות |  |
| 2:18.15 | אשר |  |
| 2:18.16 | יכבדו |  |
| 2:18.17 | בם |  |
| 2:18.18 | היהודים |  |
| 2:18.19 | את |  |
| 2:18.20 | ספר |  |
| 2:18.21 | התורה | , |
| 2:18.22 | כי |  |
| 2:18.23 | צרינו |  |
| 2:18.24 | אשר |  |
| 2:18.25 | אינם |  |
| 2:18.26 | יודעים |  |
| 2:18.27 | מה |  |
| 2:18.28 | זאת |  |
| 2:18.29 | מכעיסים |  |
| 2:18.30 | אותנו |  |
| 2:18.31 | לאמר |  |
| 2:18.32 | כי |  |
| 2:18.33 | עובדי |  |
| 2:18.34 | אלילים |  |
| 2:18.35 | אנחנו | , |
| 2:18.36 | וכי |  |
| 2:18.37 | נשתחוה |  |
| 2:18.38 | לארון |  |
| 2:18.39 | עץ |  |
| 2:18.40 | וליריעות |  |
| 2:18.41 | עור |  |
| 2:18.42 | אשר |  |
| 2:18.43 | בתוכה | , |
| 2:18.44 | לכן |  |
| 2:18.45 | אערוך |  |
| 2:18.46 | תשובתי |  |
| 2:18.47 | בסדר |  |
| 2:18.48 | הניאות |  |
| 2:18.49 | כיד |  |
| 2:18.50 | ה | ' |
| 2:18.51 | הטובה |  |
| 2:18.52 | עלי |  |
| 2:18.53 | והוא |  |
| 2:18.54 | כל |  |
| 2:18.55 | איש |  |
| 2:18.56 | ישראל |  |
| 2:18.57 | יראה |  |
| 2:18.58 | את |  |
| 2:18.59 | עצמו |  |
| 2:18.60 | מחוייב |  |
| 2:18.61 | לקום |  |
| 2:18.62 | ולעמוד |  |
| 2:18.63 | בהוצאת |  |
| 2:18.64 | ספר |  |
| 2:18.65 | תורה |  |
| 2:18.66 | מארון |  |
| 2:18.67 | הקודש |  |
| 2:18.68 | עדי |  |
| 2:18.69 | תונח |  |
| 2:18.70 | על |  |
| 2:18.71 | התיבה |  |
| 2:18.72 | להראותה |  |
| 2:18.73 | לכל |  |
| 2:18.74 | העם |  |
| 2:18.75 | ולקרות |  |
| 2:18.76 | בה | . |
| 2:18.77 | ומצינו |  |
| 2:18.78 | כי |  |
| 2:18.79 | כן |  |
| 2:18.80 | עשו |  |
| 2:18.81 | בימי |  |
| 2:18.82 | נחמיה |  |
| 2:18.83 | ככתוב |  |
| 2:18.84 | וכפתחו |  |
| 2:18.85 | עמדו |  |
| 2:18.86 | כל |  |
| 2:18.87 | העם | . |
| 2:18.88 | וזאת |  |
| 2:18.89 | יעשו |  |
| 2:18.90 | מפני |  |
| 2:18.91 | היראה |  |
| 2:18.92 | והכבוד |  |
| 2:18.93 | אשר |  |
| 2:18.94 | יכבדו |  |
| 2:18.95 | דברי |  |
| 2:18.96 | אלהים |  |
| 2:18.97 | חיים |  |
| 2:18.98 | הכתובים |  |
| 2:18.99 | בה | , |
| 2:18.100 | ומאותו |  |
| 2:18.101 | טעם |  |
| 2:18.102 | יכופו |  |
| 2:18.103 | ראשם |  |
| 2:18.104 | בנשאם |  |
| 2:18.105 | אותה |  |
| 2:18.106 | מהארון |  |
| 2:18.107 | עד |  |
| 2:18.108 | התיבה | , |
| 2:18.109 | ולא |  |
| 2:18.110 | משום |  |
| 2:18.111 | דבר |  |
| 2:18.112 | אחר |  |
| 2:18.113 | מן |  |
| 2:18.114 | הטעמים |  |
| 2:18.115 | אשר |  |
| 2:18.116 | נזכיר | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

