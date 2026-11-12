# Translation pack: milhamot-hashem perek-2-14 (Milchamot ha-Shem 2:30–2:30)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-14`.

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
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **המתפלספים** → "would-be philosophers".
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


## Milchamot ha-Shem 2:30

Hebrew (Wikisource): ומה שכתבת שז"א בן לאבא לאמא והוא אב לעלולים שאחריו, ויש לו אחות נוקבא דיליה, הכל משל וכינוי בספירות. ומה שכתבת שאין לעבוד שום פרצוף, זה ידוע בפי כל המקובלים, ואין צריכים לאזהרתך. ומה שכתבת בשם הרמב"ם חמשה הן הנקראים מינים וכו' מי הוא האומר שאין שם אלוה, או שהם שנים, או שהוא בעל גוף, או שאינו לבדו האחד או שעובד זולתו כדי להיות מליץ, או האומר שמצא סמנים או האומר שנברא העולם יש מיש ח"ו. אלא שכבר עלתה טינא בלבך, ואתה מעיין הספרים הקדושים ולוקח לך הדברים כפשטן וסומך על הבנתך. ושכחת מה שאמר חכם החכמים, ואל בינתך אל תשען. ומפני שכבר הודעת לי שאתה מבין תתיר לי ספקותי אשר שלחתי אליך ואשר אני עתיד לשלוח. ומה שכתבת והבנת רבים וגדולים חכמי לב הפורשים ממנה ראשונים ואחרונים, ח"ו שיש שום חכם פורש ממנה. ואם יש פורש ממנה הוא עצמו מודה בפה מלא שהוא מחסרון ידיעה בה. ומה שכתבת כמו שכ' הריב"ש וחו"י, שקר אתה דובר עליהם. והם שני עדים כשרים להעיד עליך שאתה דובר שקרים. וז"ל הריב"ש בקיצור כי ר' פרץ הכהן לא היה מדבר ולא מחשיב באותן הספירות. גם שמע מפיו שר' שמשון מקינון היה אומר אני מתפלל לדעת זה התינוק כלומר להוציא מלב המקובלים שהם מתפללים פעם לספירה אחת ופעם לספירה אחרת. גם בתפלת י"ח יש להם בכל אחת ואחת כוונה לספירה ידועה. וכל זה הוא דבר זר מאוד בעיני מי שאינו מקובל כמותם. וחושבים שזה אמונת שניות. וששמע אחד מן המתפלספים מספר בגנות המקובלים והיה אומר הנצרים מאמיני השלוש והמקובלים מאמיני העשיריות. ושאל להחכם הישיש דון יוסף ן' שושן והוא היה חכם בתלמוד. וראה פילוסופיא והיה מקובל וחסיד וגדול ומדקדק במצות. ואמר לו איך אתם המקובלים מכוונים בברכה אחת לספירה ידועה ובברכה אחת לספירה אחרת. ועוד הכי יש אלוהות לספירות? שיתפלל אדם להם? וענו לו חלילה שתהיה התפלה כי אם לשי"ת עילת העילות אלא שמכוין המחשבה להמשיך השפע לאותה ספירה המתיחסת לאותו דבר שהוא מבקש עליו כמו שתאמר שבברכת על הצדיקים יכוין לספירה הנקראת חסד שהיא מדת רחמים וכו'. והביא לו משל על זה. ענה ואמר והנה טוב מאד. ואמר איני תוקע עצמי באותה חכמה אחר שלא קיבלתיה מפי חכם מקובל. ואם ראיתי באורים על סודות הרמב"ן. הם מגלים טפח ומכסים טפחים. ואמר שחכם אחד פירש הקבלה שלא כד' הרמב"ן. ולכן אמר שדברי החכם ההוא יהיו לו לבדו. וחתם דבריו ולזה אני אומר שאין לסמוך בדברים כאלו אלא מפי חכם מקובל.

| id | word | punctuation after |
|---|---|---|
| 2:30.1 | ומה |  |
| 2:30.2 | שכתבת |  |
| 2:30.3 | שז"א |  |
| 2:30.4 | בן |  |
| 2:30.5 | לאבא |  |
| 2:30.6 | לאמא |  |
| 2:30.7 | והוא |  |
| 2:30.8 | אב |  |
| 2:30.9 | לעלולים |  |
| 2:30.10 | שאחריו | , |
| 2:30.11 | ויש |  |
| 2:30.12 | לו |  |
| 2:30.13 | אחות |  |
| 2:30.14 | נוקבא |  |
| 2:30.15 | דיליה | , |
| 2:30.16 | הכל |  |
| 2:30.17 | משל |  |
| 2:30.18 | וכינוי |  |
| 2:30.19 | בספירות | . |
| 2:30.20 | ומה |  |
| 2:30.21 | שכתבת |  |
| 2:30.22 | שאין |  |
| 2:30.23 | לעבוד |  |
| 2:30.24 | שום |  |
| 2:30.25 | פרצוף | , |
| 2:30.26 | זה |  |
| 2:30.27 | ידוע |  |
| 2:30.28 | בפי |  |
| 2:30.29 | כל |  |
| 2:30.30 | המקובלים | , |
| 2:30.31 | ואין |  |
| 2:30.32 | צריכים |  |
| 2:30.33 | לאזהרתך | . |
| 2:30.34 | ומה |  |
| 2:30.35 | שכתבת |  |
| 2:30.36 | בשם |  |
| 2:30.37 | הרמב"ם |  |
| 2:30.38 | חמשה |  |
| 2:30.39 | הן |  |
| 2:30.40 | הנקראים |  |
| 2:30.41 | מינים |  |
| 2:30.42 | וכו | ' |
| 2:30.43 | מי |  |
| 2:30.44 | הוא |  |
| 2:30.45 | האומר |  |
| 2:30.46 | שאין |  |
| 2:30.47 | שם |  |
| 2:30.48 | אלוה | , |
| 2:30.49 | או |  |
| 2:30.50 | שהם |  |
| 2:30.51 | שנים | , |
| 2:30.52 | או |  |
| 2:30.53 | שהוא |  |
| 2:30.54 | בעל |  |
| 2:30.55 | גוף | , |
| 2:30.56 | או |  |
| 2:30.57 | שאינו |  |
| 2:30.58 | לבדו |  |
| 2:30.59 | האחד |  |
| 2:30.60 | או |  |
| 2:30.61 | שעובד |  |
| 2:30.62 | זולתו |  |
| 2:30.63 | כדי |  |
| 2:30.64 | להיות |  |
| 2:30.65 | מליץ | , |
| 2:30.66 | או |  |
| 2:30.67 | האומר |  |
| 2:30.68 | שמצא |  |
| 2:30.69 | סמנים |  |
| 2:30.70 | או |  |
| 2:30.71 | האומר |  |
| 2:30.72 | שנברא |  |
| 2:30.73 | העולם |  |
| 2:30.74 | יש |  |
| 2:30.75 | מיש |  |
| 2:30.76 | ח"ו | . |
| 2:30.77 | אלא |  |
| 2:30.78 | שכבר |  |
| 2:30.79 | עלתה |  |
| 2:30.80 | טינא |  |
| 2:30.81 | בלבך | , |
| 2:30.82 | ואתה |  |
| 2:30.83 | מעיין |  |
| 2:30.84 | הספרים |  |
| 2:30.85 | הקדושים |  |
| 2:30.86 | ולוקח |  |
| 2:30.87 | לך |  |
| 2:30.88 | הדברים |  |
| 2:30.89 | כפשטן |  |
| 2:30.90 | וסומך |  |
| 2:30.91 | על |  |
| 2:30.92 | הבנתך | . |
| 2:30.93 | ושכחת |  |
| 2:30.94 | מה |  |
| 2:30.95 | שאמר |  |
| 2:30.96 | חכם |  |
| 2:30.97 | החכמים | , |
| 2:30.98 | ואל |  |
| 2:30.99 | בינתך |  |
| 2:30.100 | אל |  |
| 2:30.101 | תשען | . |
| 2:30.102 | ומפני |  |
| 2:30.103 | שכבר |  |
| 2:30.104 | הודעת |  |
| 2:30.105 | לי |  |
| 2:30.106 | שאתה |  |
| 2:30.107 | מבין |  |
| 2:30.108 | תתיר |  |
| 2:30.109 | לי |  |
| 2:30.110 | ספקותי |  |
| 2:30.111 | אשר |  |
| 2:30.112 | שלחתי |  |
| 2:30.113 | אליך |  |
| 2:30.114 | ואשר |  |
| 2:30.115 | אני |  |
| 2:30.116 | עתיד |  |
| 2:30.117 | לשלוח | . |
| 2:30.118 | ומה |  |
| 2:30.119 | שכתבת |  |
| 2:30.120 | והבנת |  |
| 2:30.121 | רבים |  |
| 2:30.122 | וגדולים |  |
| 2:30.123 | חכמי |  |
| 2:30.124 | לב |  |
| 2:30.125 | הפורשים |  |
| 2:30.126 | ממנה |  |
| 2:30.127 | ראשונים |  |
| 2:30.128 | ואחרונים | , |
| 2:30.129 | ח"ו |  |
| 2:30.130 | שיש |  |
| 2:30.131 | שום |  |
| 2:30.132 | חכם |  |
| 2:30.133 | פורש |  |
| 2:30.134 | ממנה | . |
| 2:30.135 | ואם |  |
| 2:30.136 | יש |  |
| 2:30.137 | פורש |  |
| 2:30.138 | ממנה |  |
| 2:30.139 | הוא |  |
| 2:30.140 | עצמו |  |
| 2:30.141 | מודה |  |
| 2:30.142 | בפה |  |
| 2:30.143 | מלא |  |
| 2:30.144 | שהוא |  |
| 2:30.145 | מחסרון |  |
| 2:30.146 | ידיעה |  |
| 2:30.147 | בה | . |
| 2:30.148 | ומה |  |
| 2:30.149 | שכתבת |  |
| 2:30.150 | כמו |  |
| 2:30.151 | שכ | ' |
| 2:30.152 | הריב"ש |  |
| 2:30.153 | וחו"י | , |
| 2:30.154 | שקר |  |
| 2:30.155 | אתה |  |
| 2:30.156 | דובר |  |
| 2:30.157 | עליהם | . |
| 2:30.158 | והם |  |
| 2:30.159 | שני |  |
| 2:30.160 | עדים |  |
| 2:30.161 | כשרים |  |
| 2:30.162 | להעיד |  |
| 2:30.163 | עליך |  |
| 2:30.164 | שאתה |  |
| 2:30.165 | דובר |  |
| 2:30.166 | שקרים | . |
| 2:30.167 | וז"ל |  |
| 2:30.168 | הריב"ש |  |
| 2:30.169 | בקיצור |  |
| 2:30.170 | כי |  |
| 2:30.171 | ר | ' |
| 2:30.172 | פרץ |  |
| 2:30.173 | הכהן |  |
| 2:30.174 | לא |  |
| 2:30.175 | היה |  |
| 2:30.176 | מדבר |  |
| 2:30.177 | ולא |  |
| 2:30.178 | מחשיב |  |
| 2:30.179 | באותן |  |
| 2:30.180 | הספירות | . |
| 2:30.181 | גם |  |
| 2:30.182 | שמע |  |
| 2:30.183 | מפיו |  |
| 2:30.184 | שר | ' |
| 2:30.185 | שמשון |  |
| 2:30.186 | מקינון |  |
| 2:30.187 | היה |  |
| 2:30.188 | אומר |  |
| 2:30.189 | אני |  |
| 2:30.190 | מתפלל |  |
| 2:30.191 | לדעת |  |
| 2:30.192 | זה |  |
| 2:30.193 | התינוק |  |
| 2:30.194 | כלומר |  |
| 2:30.195 | להוציא |  |
| 2:30.196 | מלב |  |
| 2:30.197 | המקובלים |  |
| 2:30.198 | שהם |  |
| 2:30.199 | מתפללים |  |
| 2:30.200 | פעם |  |
| 2:30.201 | לספירה |  |
| 2:30.202 | אחת |  |
| 2:30.203 | ופעם |  |
| 2:30.204 | לספירה |  |
| 2:30.205 | אחרת | . |
| 2:30.206 | גם |  |
| 2:30.207 | בתפלת |  |
| 2:30.208 | י"ח |  |
| 2:30.209 | יש |  |
| 2:30.210 | להם |  |
| 2:30.211 | בכל |  |
| 2:30.212 | אחת |  |
| 2:30.213 | ואחת |  |
| 2:30.214 | כוונה |  |
| 2:30.215 | לספירה |  |
| 2:30.216 | ידועה | . |
| 2:30.217 | וכל |  |
| 2:30.218 | זה |  |
| 2:30.219 | הוא |  |
| 2:30.220 | דבר |  |
| 2:30.221 | זר |  |
| 2:30.222 | מאוד |  |
| 2:30.223 | בעיני |  |
| 2:30.224 | מי |  |
| 2:30.225 | שאינו |  |
| 2:30.226 | מקובל |  |
| 2:30.227 | כמותם | . |
| 2:30.228 | וחושבים |  |
| 2:30.229 | שזה |  |
| 2:30.230 | אמונת |  |
| 2:30.231 | שניות | . |
| 2:30.232 | וששמע |  |
| 2:30.233 | אחד |  |
| 2:30.234 | מן |  |
| 2:30.235 | המתפלספים |  |
| 2:30.236 | מספר |  |
| 2:30.237 | בגנות |  |
| 2:30.238 | המקובלים |  |
| 2:30.239 | והיה |  |
| 2:30.240 | אומר |  |
| 2:30.241 | הנצרים |  |
| 2:30.242 | מאמיני |  |
| 2:30.243 | השלוש |  |
| 2:30.244 | והמקובלים |  |
| 2:30.245 | מאמיני |  |
| 2:30.246 | העשיריות | . |
| 2:30.247 | ושאל |  |
| 2:30.248 | להחכם |  |
| 2:30.249 | הישיש |  |
| 2:30.250 | דון |  |
| 2:30.251 | יוסף |  |
| 2:30.252 | ן | ' |
| 2:30.253 | שושן |  |
| 2:30.254 | והוא |  |
| 2:30.255 | היה |  |
| 2:30.256 | חכם |  |
| 2:30.257 | בתלמוד | . |
| 2:30.258 | וראה |  |
| 2:30.259 | פילוסופיא |  |
| 2:30.260 | והיה |  |
| 2:30.261 | מקובל |  |
| 2:30.262 | וחסיד |  |
| 2:30.263 | וגדול |  |
| 2:30.264 | ומדקדק |  |
| 2:30.265 | במצות | . |
| 2:30.266 | ואמר |  |
| 2:30.267 | לו |  |
| 2:30.268 | איך |  |
| 2:30.269 | אתם |  |
| 2:30.270 | המקובלים |  |
| 2:30.271 | מכוונים |  |
| 2:30.272 | בברכה |  |
| 2:30.273 | אחת |  |
| 2:30.274 | לספירה |  |
| 2:30.275 | ידועה |  |
| 2:30.276 | ובברכה |  |
| 2:30.277 | אחת |  |
| 2:30.278 | לספירה |  |
| 2:30.279 | אחרת | . |
| 2:30.280 | ועוד |  |
| 2:30.281 | הכי |  |
| 2:30.282 | יש |  |
| 2:30.283 | אלוהות |  |
| 2:30.284 | לספירות | ? |
| 2:30.285 | שיתפלל |  |
| 2:30.286 | אדם |  |
| 2:30.287 | להם | ? |
| 2:30.288 | וענו |  |
| 2:30.289 | לו |  |
| 2:30.290 | חלילה |  |
| 2:30.291 | שתהיה |  |
| 2:30.292 | התפלה |  |
| 2:30.293 | כי |  |
| 2:30.294 | אם |  |
| 2:30.295 | לשי"ת |  |
| 2:30.296 | עילת |  |
| 2:30.297 | העילות |  |
| 2:30.298 | אלא |  |
| 2:30.299 | שמכוין |  |
| 2:30.300 | המחשבה |  |
| 2:30.301 | להמשיך |  |
| 2:30.302 | השפע |  |
| 2:30.303 | לאותה |  |
| 2:30.304 | ספירה |  |
| 2:30.305 | המתיחסת |  |
| 2:30.306 | לאותו |  |
| 2:30.307 | דבר |  |
| 2:30.308 | שהוא |  |
| 2:30.309 | מבקש |  |
| 2:30.310 | עליו |  |
| 2:30.311 | כמו |  |
| 2:30.312 | שתאמר |  |
| 2:30.313 | שבברכת |  |
| 2:30.314 | על |  |
| 2:30.315 | הצדיקים |  |
| 2:30.316 | יכוין |  |
| 2:30.317 | לספירה |  |
| 2:30.318 | הנקראת |  |
| 2:30.319 | חסד |  |
| 2:30.320 | שהיא |  |
| 2:30.321 | מדת |  |
| 2:30.322 | רחמים |  |
| 2:30.323 | וכו | '. |
| 2:30.324 | והביא |  |
| 2:30.325 | לו |  |
| 2:30.326 | משל |  |
| 2:30.327 | על |  |
| 2:30.328 | זה | . |
| 2:30.329 | ענה |  |
| 2:30.330 | ואמר |  |
| 2:30.331 | והנה |  |
| 2:30.332 | טוב |  |
| 2:30.333 | מאד | . |
| 2:30.334 | ואמר |  |
| 2:30.335 | איני |  |
| 2:30.336 | תוקע |  |
| 2:30.337 | עצמי |  |
| 2:30.338 | באותה |  |
| 2:30.339 | חכמה |  |
| 2:30.340 | אחר |  |
| 2:30.341 | שלא |  |
| 2:30.342 | קיבלתיה |  |
| 2:30.343 | מפי |  |
| 2:30.344 | חכם |  |
| 2:30.345 | מקובל | . |
| 2:30.346 | ואם |  |
| 2:30.347 | ראיתי |  |
| 2:30.348 | באורים |  |
| 2:30.349 | על |  |
| 2:30.350 | סודות |  |
| 2:30.351 | הרמב"ן | . |
| 2:30.352 | הם |  |
| 2:30.353 | מגלים |  |
| 2:30.354 | טפח |  |
| 2:30.355 | ומכסים |  |
| 2:30.356 | טפחים | . |
| 2:30.357 | ואמר |  |
| 2:30.358 | שחכם |  |
| 2:30.359 | אחד |  |
| 2:30.360 | פירש |  |
| 2:30.361 | הקבלה |  |
| 2:30.362 | שלא |  |
| 2:30.363 | כד | ' |
| 2:30.364 | הרמב"ן | . |
| 2:30.365 | ולכן |  |
| 2:30.366 | אמר |  |
| 2:30.367 | שדברי |  |
| 2:30.368 | החכם |  |
| 2:30.369 | ההוא |  |
| 2:30.370 | יהיו |  |
| 2:30.371 | לו |  |
| 2:30.372 | לבדו | . |
| 2:30.373 | וחתם |  |
| 2:30.374 | דבריו |  |
| 2:30.375 | ולזה |  |
| 2:30.376 | אני |  |
| 2:30.377 | אומר |  |
| 2:30.378 | שאין |  |
| 2:30.379 | לסמוך |  |
| 2:30.380 | בדברים |  |
| 2:30.381 | כאלו |  |
| 2:30.382 | אלא |  |
| 2:30.383 | מפי |  |
| 2:30.384 | חכם |  |
| 2:30.385 | מקובל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

