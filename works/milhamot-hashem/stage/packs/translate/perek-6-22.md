# Translation pack: milhamot-hashem perek-6-22 (Milchamot ha-Shem 6:33–6:33)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-22.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-22`.

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
- **האלהי** before a name → "the divine" (note); never dropped.
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

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
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


## Milchamot ha-Shem 6:33

Hebrew (Wikisource): ובגמרא מכות דף י"כ גרסינן אמר ר"ל שלש טעויות עתיד לטעות שרה של אדם ושנ' מי זה בא מאדום חמוץ בגדים מבצרה טועה שאינה קולטת אלא בצר והוא גולה לבצרה טועה שאינה קולטת אלא אדם והוא מלאך, טועה שאינה קולטת אלא שוגג והוא מזיד, ע"כ ופי' רש"י ז"ל שרה של אדום סמאל שיברה לבצרה, שבתחלה יפרע הקב"ה ממנו כשיגיע קץ אדום ליחרב כדכתיב יפקוד ה' על צבא המרום במרום, ואחר כך ועל מלכי האדמה באדמה, חמוץ בגדים מדמו של סמאל, ואע"פ שאין המלאכים בשר ודם, כתב בו הכתוב בעין הריגת אדם לשבר את האוזן מה שהיא יכולה לשמוע עכ"ל הרי בהדיא שסמא"ל מלאך ולא אלוה, ויש לשאול לפי שיטת המקובלים החדשים, כיון שיש לכל אומה שר הממונה עליה, ושר של ישראל מיכאל, ושל מצרים דהב, ושל אדום סמאל וכו'. למה נתנו המקובלים החדשים כל זו הגדולה לשמאל להקרא בשם אל אחר, ולהיות לו עשר ספירות דקליפה, ואדם בליעל ואריך ואבא ואמא וזעיר ונבה דקליפה. ואמרו שגם זה לעומת זה עשה האלהים, היכלות הטומאה נגד היכלות הקדושה, עם כל הפרצופים שלה, ולמה לא עשו כן לשאר כל השרים של כל האומות? ולמה יגרע מיכאל השר של ישראל מכל אלו המעלות? ולמה הניחוהו בלא אשה, ושמאל השיאוהו בלילית עם ת"פ מחנותיה? כלל הדברים לא מציונ בתורה ולא בנביאים ולא בכתובים ולא בדברי רבותינו ז"ל מציאות אל אחד בעולם, לא קדושים ולא טמאים זולתי ה' אלהינו ב"ה. והכל משרתיו עושי רצונו. ורק לפי מחשבת עובדיהן דברה תורה, שהם בוחרים להם כל אחד מכל הנבראים שהמציא הקב"ה בעולם ומשתחוים לו ועובדים ומקריבים לו קרבנות, וחושבין אותו לאמצעי או לאלוה. כמבואר בדברי הרמב"ם ז"ל בפ"א מלה' ע"ז. יש שבחרו בשמש ויש שבחרו בירח או באחד הכוכבים או המזלות או לאש או למים לעבדם ולהשתחוות להם. ועל כולם הזהירנו השי"ת ואמר לא יהיה לך אלהים אחרים, לא תשתחוה לאל אחר, לא שיש מציאות אל אחר בעולם מבלעדי ה' אלהינו. ולפי שבאמת אינן אלוהות כמאמר רז"ל, לכן תרגם רבי' סעדיה אלהים אחרים מעבוד אלר, ולא תרגם אלאה אלר. וכן תרגם זבח לאלהים יחרם ומן דבח לאלמעודעת פאלי'קתל לפי שבאמת אינן אלוהות.

| id | word | punctuation after |
|---|---|---|
| 6:33.1 | ובגמרא |  |
| 6:33.2 | מכות |  |
| 6:33.3 | דף |  |
| 6:33.4 | י"כ |  |
| 6:33.5 | גרסינן |  |
| 6:33.6 | אמר |  |
| 6:33.7 | ר"ל |  |
| 6:33.8 | שלש |  |
| 6:33.9 | טעויות |  |
| 6:33.10 | עתיד |  |
| 6:33.11 | לטעות |  |
| 6:33.12 | שרה |  |
| 6:33.13 | של |  |
| 6:33.14 | אדם |  |
| 6:33.15 | ושנ | ' |
| 6:33.16 | מי |  |
| 6:33.17 | זה |  |
| 6:33.18 | בא |  |
| 6:33.19 | מאדום |  |
| 6:33.20 | חמוץ |  |
| 6:33.21 | בגדים |  |
| 6:33.22 | מבצרה |  |
| 6:33.23 | טועה |  |
| 6:33.24 | שאינה |  |
| 6:33.25 | קולטת |  |
| 6:33.26 | אלא |  |
| 6:33.27 | בצר |  |
| 6:33.28 | והוא |  |
| 6:33.29 | גולה |  |
| 6:33.30 | לבצרה |  |
| 6:33.31 | טועה |  |
| 6:33.32 | שאינה |  |
| 6:33.33 | קולטת |  |
| 6:33.34 | אלא |  |
| 6:33.35 | אדם |  |
| 6:33.36 | והוא |  |
| 6:33.37 | מלאך | , |
| 6:33.38 | טועה |  |
| 6:33.39 | שאינה |  |
| 6:33.40 | קולטת |  |
| 6:33.41 | אלא |  |
| 6:33.42 | שוגג |  |
| 6:33.43 | והוא |  |
| 6:33.44 | מזיד | , |
| 6:33.45 | ע"כ |  |
| 6:33.46 | ופי | ' |
| 6:33.47 | רש"י |  |
| 6:33.48 | ז"ל |  |
| 6:33.49 | שרה |  |
| 6:33.50 | של |  |
| 6:33.51 | אדום |  |
| 6:33.52 | סמאל |  |
| 6:33.53 | שיברה |  |
| 6:33.54 | לבצרה | , |
| 6:33.55 | שבתחלה |  |
| 6:33.56 | יפרע |  |
| 6:33.57 | הקב"ה |  |
| 6:33.58 | ממנו |  |
| 6:33.59 | כשיגיע |  |
| 6:33.60 | קץ |  |
| 6:33.61 | אדום |  |
| 6:33.62 | ליחרב |  |
| 6:33.63 | כדכתיב |  |
| 6:33.64 | יפקוד |  |
| 6:33.65 | ה | ' |
| 6:33.66 | על |  |
| 6:33.67 | צבא |  |
| 6:33.68 | המרום |  |
| 6:33.69 | במרום | , |
| 6:33.70 | ואחר |  |
| 6:33.71 | כך |  |
| 6:33.72 | ועל |  |
| 6:33.73 | מלכי |  |
| 6:33.74 | האדמה |  |
| 6:33.75 | באדמה | , |
| 6:33.76 | חמוץ |  |
| 6:33.77 | בגדים |  |
| 6:33.78 | מדמו |  |
| 6:33.79 | של |  |
| 6:33.80 | סמאל | , |
| 6:33.81 | ואע"פ |  |
| 6:33.82 | שאין |  |
| 6:33.83 | המלאכים |  |
| 6:33.84 | בשר |  |
| 6:33.85 | ודם | , |
| 6:33.86 | כתב |  |
| 6:33.87 | בו |  |
| 6:33.88 | הכתוב |  |
| 6:33.89 | בעין |  |
| 6:33.90 | הריגת |  |
| 6:33.91 | אדם |  |
| 6:33.92 | לשבר |  |
| 6:33.93 | את |  |
| 6:33.94 | האוזן |  |
| 6:33.95 | מה |  |
| 6:33.96 | שהיא |  |
| 6:33.97 | יכולה |  |
| 6:33.98 | לשמוע |  |
| 6:33.99 | עכ"ל |  |
| 6:33.100 | הרי |  |
| 6:33.101 | בהדיא |  |
| 6:33.102 | שסמא"ל |  |
| 6:33.103 | מלאך |  |
| 6:33.104 | ולא |  |
| 6:33.105 | אלוה | , |
| 6:33.106 | ויש |  |
| 6:33.107 | לשאול |  |
| 6:33.108 | לפי |  |
| 6:33.109 | שיטת |  |
| 6:33.110 | המקובלים |  |
| 6:33.111 | החדשים | , |
| 6:33.112 | כיון |  |
| 6:33.113 | שיש |  |
| 6:33.114 | לכל |  |
| 6:33.115 | אומה |  |
| 6:33.116 | שר |  |
| 6:33.117 | הממונה |  |
| 6:33.118 | עליה | , |
| 6:33.119 | ושר |  |
| 6:33.120 | של |  |
| 6:33.121 | ישראל |  |
| 6:33.122 | מיכאל | , |
| 6:33.123 | ושל |  |
| 6:33.124 | מצרים |  |
| 6:33.125 | דהב | , |
| 6:33.126 | ושל |  |
| 6:33.127 | אדום |  |
| 6:33.128 | סמאל |  |
| 6:33.129 | וכו | '. |
| 6:33.130 | למה |  |
| 6:33.131 | נתנו |  |
| 6:33.132 | המקובלים |  |
| 6:33.133 | החדשים |  |
| 6:33.134 | כל |  |
| 6:33.135 | זו |  |
| 6:33.136 | הגדולה |  |
| 6:33.137 | לשמאל |  |
| 6:33.138 | להקרא |  |
| 6:33.139 | בשם |  |
| 6:33.140 | אל |  |
| 6:33.141 | אחר | , |
| 6:33.142 | ולהיות |  |
| 6:33.143 | לו |  |
| 6:33.144 | עשר |  |
| 6:33.145 | ספירות |  |
| 6:33.146 | דקליפה | , |
| 6:33.147 | ואדם |  |
| 6:33.148 | בליעל |  |
| 6:33.149 | ואריך |  |
| 6:33.150 | ואבא |  |
| 6:33.151 | ואמא |  |
| 6:33.152 | וזעיר |  |
| 6:33.153 | ונבה |  |
| 6:33.154 | דקליפה | . |
| 6:33.155 | ואמרו |  |
| 6:33.156 | שגם |  |
| 6:33.157 | זה |  |
| 6:33.158 | לעומת |  |
| 6:33.159 | זה |  |
| 6:33.160 | עשה |  |
| 6:33.161 | האלהים | , |
| 6:33.162 | היכלות |  |
| 6:33.163 | הטומאה |  |
| 6:33.164 | נגד |  |
| 6:33.165 | היכלות |  |
| 6:33.166 | הקדושה | , |
| 6:33.167 | עם |  |
| 6:33.168 | כל |  |
| 6:33.169 | הפרצופים |  |
| 6:33.170 | שלה | , |
| 6:33.171 | ולמה |  |
| 6:33.172 | לא |  |
| 6:33.173 | עשו |  |
| 6:33.174 | כן |  |
| 6:33.175 | לשאר |  |
| 6:33.176 | כל |  |
| 6:33.177 | השרים |  |
| 6:33.178 | של |  |
| 6:33.179 | כל |  |
| 6:33.180 | האומות | ? |
| 6:33.181 | ולמה |  |
| 6:33.182 | יגרע |  |
| 6:33.183 | מיכאל |  |
| 6:33.184 | השר |  |
| 6:33.185 | של |  |
| 6:33.186 | ישראל |  |
| 6:33.187 | מכל |  |
| 6:33.188 | אלו |  |
| 6:33.189 | המעלות | ? |
| 6:33.190 | ולמה |  |
| 6:33.191 | הניחוהו |  |
| 6:33.192 | בלא |  |
| 6:33.193 | אשה | , |
| 6:33.194 | ושמאל |  |
| 6:33.195 | השיאוהו |  |
| 6:33.196 | בלילית |  |
| 6:33.197 | עם |  |
| 6:33.198 | ת"פ |  |
| 6:33.199 | מחנותיה | ? |
| 6:33.200 | כלל |  |
| 6:33.201 | הדברים |  |
| 6:33.202 | לא |  |
| 6:33.203 | מציונ |  |
| 6:33.204 | בתורה |  |
| 6:33.205 | ולא |  |
| 6:33.206 | בנביאים |  |
| 6:33.207 | ולא |  |
| 6:33.208 | בכתובים |  |
| 6:33.209 | ולא |  |
| 6:33.210 | בדברי |  |
| 6:33.211 | רבותינו |  |
| 6:33.212 | ז"ל |  |
| 6:33.213 | מציאות |  |
| 6:33.214 | אל |  |
| 6:33.215 | אחד |  |
| 6:33.216 | בעולם | , |
| 6:33.217 | לא |  |
| 6:33.218 | קדושים |  |
| 6:33.219 | ולא |  |
| 6:33.220 | טמאים |  |
| 6:33.221 | זולתי |  |
| 6:33.222 | ה | ' |
| 6:33.223 | אלהינו |  |
| 6:33.224 | ב"ה | . |
| 6:33.225 | והכל |  |
| 6:33.226 | משרתיו |  |
| 6:33.227 | עושי |  |
| 6:33.228 | רצונו | . |
| 6:33.229 | ורק |  |
| 6:33.230 | לפי |  |
| 6:33.231 | מחשבת |  |
| 6:33.232 | עובדיהן |  |
| 6:33.233 | דברה |  |
| 6:33.234 | תורה | , |
| 6:33.235 | שהם |  |
| 6:33.236 | בוחרים |  |
| 6:33.237 | להם |  |
| 6:33.238 | כל |  |
| 6:33.239 | אחד |  |
| 6:33.240 | מכל |  |
| 6:33.241 | הנבראים |  |
| 6:33.242 | שהמציא |  |
| 6:33.243 | הקב"ה |  |
| 6:33.244 | בעולם |  |
| 6:33.245 | ומשתחוים |  |
| 6:33.246 | לו |  |
| 6:33.247 | ועובדים |  |
| 6:33.248 | ומקריבים |  |
| 6:33.249 | לו |  |
| 6:33.250 | קרבנות | , |
| 6:33.251 | וחושבין |  |
| 6:33.252 | אותו |  |
| 6:33.253 | לאמצעי |  |
| 6:33.254 | או |  |
| 6:33.255 | לאלוה | . |
| 6:33.256 | כמבואר |  |
| 6:33.257 | בדברי |  |
| 6:33.258 | הרמב"ם |  |
| 6:33.259 | ז"ל |  |
| 6:33.260 | בפ"א |  |
| 6:33.261 | מלה | ' |
| 6:33.262 | ע"ז | . |
| 6:33.263 | יש |  |
| 6:33.264 | שבחרו |  |
| 6:33.265 | בשמש |  |
| 6:33.266 | ויש |  |
| 6:33.267 | שבחרו |  |
| 6:33.268 | בירח |  |
| 6:33.269 | או |  |
| 6:33.270 | באחד |  |
| 6:33.271 | הכוכבים |  |
| 6:33.272 | או |  |
| 6:33.273 | המזלות |  |
| 6:33.274 | או |  |
| 6:33.275 | לאש |  |
| 6:33.276 | או |  |
| 6:33.277 | למים |  |
| 6:33.278 | לעבדם |  |
| 6:33.279 | ולהשתחוות |  |
| 6:33.280 | להם | . |
| 6:33.281 | ועל |  |
| 6:33.282 | כולם |  |
| 6:33.283 | הזהירנו |  |
| 6:33.284 | השי"ת |  |
| 6:33.285 | ואמר |  |
| 6:33.286 | לא |  |
| 6:33.287 | יהיה |  |
| 6:33.288 | לך |  |
| 6:33.289 | אלהים |  |
| 6:33.290 | אחרים | , |
| 6:33.291 | לא |  |
| 6:33.292 | תשתחוה |  |
| 6:33.293 | לאל |  |
| 6:33.294 | אחר | , |
| 6:33.295 | לא |  |
| 6:33.296 | שיש |  |
| 6:33.297 | מציאות |  |
| 6:33.298 | אל |  |
| 6:33.299 | אחר |  |
| 6:33.300 | בעולם |  |
| 6:33.301 | מבלעדי |  |
| 6:33.302 | ה | ' |
| 6:33.303 | אלהינו | . |
| 6:33.304 | ולפי |  |
| 6:33.305 | שבאמת |  |
| 6:33.306 | אינן |  |
| 6:33.307 | אלוהות |  |
| 6:33.308 | כמאמר |  |
| 6:33.309 | רז"ל | , |
| 6:33.310 | לכן |  |
| 6:33.311 | תרגם |  |
| 6:33.312 | רבי | ' |
| 6:33.313 | סעדיה |  |
| 6:33.314 | אלהים |  |
| 6:33.315 | אחרים |  |
| 6:33.316 | מעבוד |  |
| 6:33.317 | אלר | , |
| 6:33.318 | ולא |  |
| 6:33.319 | תרגם |  |
| 6:33.320 | אלאה |  |
| 6:33.321 | אלר | . |
| 6:33.322 | וכן |  |
| 6:33.323 | תרגם |  |
| 6:33.324 | זבח |  |
| 6:33.325 | לאלהים |  |
| 6:33.326 | יחרם |  |
| 6:33.327 | ומן |  |
| 6:33.328 | דבח |  |
| 6:33.329 | לאלמעודעת |  |
| 6:33.330 | פאלי'קתל |  |
| 6:33.331 | לפי |  |
| 6:33.332 | שבאמת |  |
| 6:33.333 | אינן |  |
| 6:33.334 | אלוהות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

