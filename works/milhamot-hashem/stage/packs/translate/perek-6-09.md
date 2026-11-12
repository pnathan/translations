# Translation pack: milhamot-hashem perek-6-09 (Milchamot ha-Shem 6:14–6:14)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-09`.

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
- **רשב"י** → "Rabbi Shimon bar Yochai", always written out (רשב"י התנא → "...
  the Tanna"). Keep exactly the qualifier the author gives the writer who
  "named himself" after the Tanna ("אשר כינה את עצמו בשם רשב״י התנא"); no
  "pseudo-" or "so-called"; do not settle which he means.
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
- **עשתונות איש ירושלים** → "the thoughts of the man of Jerusalem" (note); no
  identification.
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
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
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


## Milchamot ha-Shem 6:14

Hebrew (Wikisource): ובסוף פרק קמא דחגיגה י' ע"א שנינו היתר נדרים פורחים באויר ואין להם על מה שיסמכו, הלכות שבת חגיגות והמעילות הרי הם כהררים התלויים בשערה מקרא מועט והלכות מרובות, הדינין והעבודות, הטהורות והטומאות ועריות יש להן על מה שיסמכו הן הן גופי תורה. ופי' בגמרא דהכי קאמר הן והן גופי תורה בין אלו שיש להן על מה שיסמכו בין אלו שהם כהררים התלוים בשערה אלו ואלו גופי תורה, כלומר עיקרה ועצמה של תורה. גם בסוף פרק שלישי דאבות שנינו גם כן ר' אלעזר חסמא אומר קינין ופתחי נדה הן הן גופי הלכות, ופי' ר' עובדיה הן הן גופי הלכות עיקר תורה שבעל פה, שמקבלים עליה שכר ע"כ. ובמסכת מגילה דף כ"ח ע"ב אמרו ודאשתמש בתגא חלף, תני ריש לקיש זה המשתמש במי ששונה הלכות כתרה של תורה. ע"כ מכל מה שהבאתי יתבאר לך שההלכות שבמשנה ובגמרא הן הן עיקרה ועצמה של תורה שבעל פה שקבל משה בסיני שאדם מקבל שכר עליהם וזוכה בהן לחיי העולם הבא, והן כתרה של תורה שנאמר בה בי מלכים ימלכו בי שרים ישורו. ובמדרש רבה פרשת קרח דרשו רז"ל ששים המה מלכות אלו ששים מסכתות ושמנים פלגשים אלו שמנים בתי מדרשות שהיו בירושלים כנגד פתחיה, ועלמות אין מספר זו משנה החיצונה, ופירשו המתנות כהונה, והרב עץ חיים, דקרי לבתי מדרשות פילגשים על שם שנושאים ונותנים וחולקים זה עם זה כמו פילגשם, וכי משנה החיצונה מה שלמדו היחידים חוץ לאלו בתי המדרשות. וביד יוסף פי' שכוונת המדרש על הברייתות וע"כ. וכן מצאתי סיוע לפירושו בשיר השירים רבה וז"ל ששים המה מלכות אלו ששים מסכתות של הלכות, ושמונים פילגשים אלו שמונים פרשיות שבתורת כהנים. ועלמות אין מספר אין קץ לתוספתות. אחת היא הן חולקין אלו עם אלו וכולהון דורשים מטעם אחד, מהלכה אחת, מגזירה שוה, מקל וחומר ע"כ. ראה כמה מחבבים רז"ל את המשנה והתלמוד וקראו להן מלכות. ולא שפחות כאשר עשה הפילסוף הרמאי האדוק בדעות חיצוניות הוא מחבר הזהר כמו שיתבאר לקמן. ובבא מציעא פרק אלו מציאות (בבא מציעא דף ל"ג) גרסינן. ת"ר העוסקים במקרא מדה ואינה מדה במשנה מדה ונוטל עליה שכר, בתלמוד, אין לך מדה גדולה מזו, ולעולם הוי רץ למשנה יותר מן התלמוד, ופריך היא גופה קשיא, אמרת בתלמוד אין לך מדה גדולה מזו, והדר אמרת ולעולם הוי רץ למשנה יותר מן התלמוד? אמר ר' יוחנן בימי ר' נשנית משנה זו, שבקו כולי עלמא מתניתין ואזלו בתר תלמודא, הדר דרש להו ולעולם הוי רץ למשנה יותר מן התלמוד עיין שם ובפי' רש"י. ובירושלמי מבואר בהדיא דרשב"י הוא דתני לה, כדגרסינן התם בפרק כל כתבי הדא אמרא שהמשנה קודמת למקרא, ודא מסייעא לרשב"י דתני רשב"י העוסק במקרא מדה שאינה מדה, במשנה מדה שנוטלין ממנה שכר, העוסק בתלמוד אין לך מדה גדולה מזו, לעולם הוי רץ אחר המשנה יותר מן התלמוד אמר ר' יוסי בר' בון הדא דאת אמר עד שלא שקע בו רבי רוב משניות, אבל מששיקע בו רבי רוב משניות לעולם הוי רץ אחר התלמוד יותר מן המשנה ע"כ. וכן בפ"ק דשבת סוף פסקא דלא ישב אדם לפני הספר קאמר התם דהקורא את שמע בעונתה כמשנה היא, ושלא בעונתה כאדם שקורא בתורה ששכרו פחות מהעוסק במשנה, ומסיים בה רשב"י כדעתיה דרשב"י אמר העוסק במקרא מדה שאינה מדה, ורבנין עבדין מקרא כמשנה עכ"ל, ובילקוט ריש פר' צו ר' חנינה אמר אין הגליות מתכנסות אלא בזכות משניות שנ' גם כי יתנו בגוים עתה אקבצם ע"כ. ראה תראה ידידי הקורא דברי רז"ל שהבאתי מקצת מהם כמה הם מחבבים ומחשיבים את המשנה שהיא עיקר תורה שבעל פה שנאמרה למשה בסיני, ובפרט שים עיניך לדברי תנא דידן רשב"י הנאמרים באמת בבבלי ובירושלמי כמה הוא מחבב ומחשיב את המשנה יותר מן המקרא, ואת התלמוד שהוא פי' טעמי המשנה יותר מכולם שאין לך מדה גדולה יותר ממנה, אלא דבירושלמי מוקי להא דאמר לעולם הוי רץ למשנה יותר מן התלמוד קודם שסידר רבי את המשנה שאז היו הרבה משניות שהם שגורים בפי התלמידים כל אחת בשם אומרה, והרבה בריתות היו להם, אבל לאחר שסידר רבי את המשנה והשמיט בה כמה משניות ובריתות, וגם סתם במשנה דברי יחידים שראה ר' דבריהם ושנאן סתם כדי לקבוע הלכה כמותם, לפיכך חזרו לומר הוי רץ לתלמוד יותר מן המשנה. ומדברי כולם למדנו כמה היה חביב להם לימוד המשנה והתלמוד, ושרשב"י הוא דתני לה, ומזה הוכחה גמורה ברורה שס' הזהר אינו לרשב"י התנא ע"ה, כי אם לפילסוף רמאי גדול שלמד כל חכמת ישראל בשני התלמודים ומדרשי רז"ל. והוא מאמין בדעות הכשדים והכלדיים פילסופים הראשונים שהיו חולקים על אברהם אבינו באמונתו באל אחד, ומאמינים בשתי סבות קודמות, וכל אחת פועלת היפך האחרת, והסבות הפועלות הטוב קורא להם בשם סטרא דקדושה, והם אדם קדמון וכתר עליון עם שאר הפרצופים וצורות המשתלשלים מהם, והסבות הפועלות הרע הוא קורא להם סטרא דקליפה, והם אדם בליעל וכתר עליון דקליפה עם שאר פרצופים וצורות דעשר ספירות דקליפה המשתלשלים מהם הנקראים בשם אל אחר, כי הוא מאמין ומאמת מציאות אלהים אחרים בעולם כדעת הכשדים והמצרים הקדמונים בפילוסופיא הקדומה שכבר נמר ריחה וסר טעמה. ובשפתי חלקות וגיזומים הוא מושך לב רבים מהרבנים הפתאים להאמין באלהים אחרים רבים ולזלזל בלמוד המשנה והתלמוד עד שלא ידעו מה לעשות והיאך יפרשו מן האיסורי' שנאמרו למשה בסיני כגון ח"י טריפות הנזכרים במשנה וחילוקי דיניהם, לפי שלא נתפרש בתורה בהדיא אלא דרוסה כמו שכ' הרמב"ם ז"ל בהל' מאכלות אסורות. ויכוצא בזה בשאר איסורים של תורה שלא נתפרשו בתורה שבכתב ונמסרו למשה רבי' ע"ה בעל פה. והמה מפורשים במשנה ובתלמוד כגון דיני תערובת איסור ובשר בחלב. ואיזה מהן אסור מדאוריתא ואיזה אסור מדרבנן. והפילסוף הנז' כיוון לפסול ולהשכיח מישראל כל הדינים המבוארים במשנה ובתלמוד. ובהדיא אמר שבסוף גלותא יתפרנסון מחבורא דיליה.

| id | word | punctuation after |
|---|---|---|
| 6:14.1 | ובסוף |  |
| 6:14.2 | פרק |  |
| 6:14.3 | קמא |  |
| 6:14.4 | דחגיגה |  |
| 6:14.5 | י | ' |
| 6:14.6 | ע"א |  |
| 6:14.7 | שנינו |  |
| 6:14.8 | היתר |  |
| 6:14.9 | נדרים |  |
| 6:14.10 | פורחים |  |
| 6:14.11 | באויר |  |
| 6:14.12 | ואין |  |
| 6:14.13 | להם |  |
| 6:14.14 | על |  |
| 6:14.15 | מה |  |
| 6:14.16 | שיסמכו | , |
| 6:14.17 | הלכות |  |
| 6:14.18 | שבת |  |
| 6:14.19 | חגיגות |  |
| 6:14.20 | והמעילות |  |
| 6:14.21 | הרי |  |
| 6:14.22 | הם |  |
| 6:14.23 | כהררים |  |
| 6:14.24 | התלויים |  |
| 6:14.25 | בשערה |  |
| 6:14.26 | מקרא |  |
| 6:14.27 | מועט |  |
| 6:14.28 | והלכות |  |
| 6:14.29 | מרובות | , |
| 6:14.30 | הדינין |  |
| 6:14.31 | והעבודות | , |
| 6:14.32 | הטהורות |  |
| 6:14.33 | והטומאות |  |
| 6:14.34 | ועריות |  |
| 6:14.35 | יש |  |
| 6:14.36 | להן |  |
| 6:14.37 | על |  |
| 6:14.38 | מה |  |
| 6:14.39 | שיסמכו |  |
| 6:14.40 | הן |  |
| 6:14.41 | הן |  |
| 6:14.42 | גופי |  |
| 6:14.43 | תורה | . |
| 6:14.44 | ופי | ' |
| 6:14.45 | בגמרא |  |
| 6:14.46 | דהכי |  |
| 6:14.47 | קאמר |  |
| 6:14.48 | הן |  |
| 6:14.49 | והן |  |
| 6:14.50 | גופי |  |
| 6:14.51 | תורה |  |
| 6:14.52 | בין |  |
| 6:14.53 | אלו |  |
| 6:14.54 | שיש |  |
| 6:14.55 | להן |  |
| 6:14.56 | על |  |
| 6:14.57 | מה |  |
| 6:14.58 | שיסמכו |  |
| 6:14.59 | בין |  |
| 6:14.60 | אלו |  |
| 6:14.61 | שהם |  |
| 6:14.62 | כהררים |  |
| 6:14.63 | התלוים |  |
| 6:14.64 | בשערה |  |
| 6:14.65 | אלו |  |
| 6:14.66 | ואלו |  |
| 6:14.67 | גופי |  |
| 6:14.68 | תורה | , |
| 6:14.69 | כלומר |  |
| 6:14.70 | עיקרה |  |
| 6:14.71 | ועצמה |  |
| 6:14.72 | של |  |
| 6:14.73 | תורה | . |
| 6:14.74 | גם |  |
| 6:14.75 | בסוף |  |
| 6:14.76 | פרק |  |
| 6:14.77 | שלישי |  |
| 6:14.78 | דאבות |  |
| 6:14.79 | שנינו |  |
| 6:14.80 | גם |  |
| 6:14.81 | כן |  |
| 6:14.82 | ר | ' |
| 6:14.83 | אלעזר |  |
| 6:14.84 | חסמא |  |
| 6:14.85 | אומר |  |
| 6:14.86 | קינין |  |
| 6:14.87 | ופתחי |  |
| 6:14.88 | נדה |  |
| 6:14.89 | הן |  |
| 6:14.90 | הן |  |
| 6:14.91 | גופי |  |
| 6:14.92 | הלכות | , |
| 6:14.93 | ופי | ' |
| 6:14.94 | ר | ' |
| 6:14.95 | עובדיה |  |
| 6:14.96 | הן |  |
| 6:14.97 | הן |  |
| 6:14.98 | גופי |  |
| 6:14.99 | הלכות |  |
| 6:14.100 | עיקר |  |
| 6:14.101 | תורה |  |
| 6:14.102 | שבעל |  |
| 6:14.103 | פה | , |
| 6:14.104 | שמקבלים |  |
| 6:14.105 | עליה |  |
| 6:14.106 | שכר |  |
| 6:14.107 | ע"כ | . |
| 6:14.108 | ובמסכת |  |
| 6:14.109 | מגילה |  |
| 6:14.110 | דף |  |
| 6:14.111 | כ"ח |  |
| 6:14.112 | ע"ב |  |
| 6:14.113 | אמרו |  |
| 6:14.114 | ודאשתמש |  |
| 6:14.115 | בתגא |  |
| 6:14.116 | חלף | , |
| 6:14.117 | תני |  |
| 6:14.118 | ריש |  |
| 6:14.119 | לקיש |  |
| 6:14.120 | זה |  |
| 6:14.121 | המשתמש |  |
| 6:14.122 | במי |  |
| 6:14.123 | ששונה |  |
| 6:14.124 | הלכות |  |
| 6:14.125 | כתרה |  |
| 6:14.126 | של |  |
| 6:14.127 | תורה | . |
| 6:14.128 | ע"כ |  |
| 6:14.129 | מכל |  |
| 6:14.130 | מה |  |
| 6:14.131 | שהבאתי |  |
| 6:14.132 | יתבאר |  |
| 6:14.133 | לך |  |
| 6:14.134 | שההלכות |  |
| 6:14.135 | שבמשנה |  |
| 6:14.136 | ובגמרא |  |
| 6:14.137 | הן |  |
| 6:14.138 | הן |  |
| 6:14.139 | עיקרה |  |
| 6:14.140 | ועצמה |  |
| 6:14.141 | של |  |
| 6:14.142 | תורה |  |
| 6:14.143 | שבעל |  |
| 6:14.144 | פה |  |
| 6:14.145 | שקבל |  |
| 6:14.146 | משה |  |
| 6:14.147 | בסיני |  |
| 6:14.148 | שאדם |  |
| 6:14.149 | מקבל |  |
| 6:14.150 | שכר |  |
| 6:14.151 | עליהם |  |
| 6:14.152 | וזוכה |  |
| 6:14.153 | בהן |  |
| 6:14.154 | לחיי |  |
| 6:14.155 | העולם |  |
| 6:14.156 | הבא | , |
| 6:14.157 | והן |  |
| 6:14.158 | כתרה |  |
| 6:14.159 | של |  |
| 6:14.160 | תורה |  |
| 6:14.161 | שנאמר |  |
| 6:14.162 | בה |  |
| 6:14.163 | בי |  |
| 6:14.164 | מלכים |  |
| 6:14.165 | ימלכו |  |
| 6:14.166 | בי |  |
| 6:14.167 | שרים |  |
| 6:14.168 | ישורו | . |
| 6:14.169 | ובמדרש |  |
| 6:14.170 | רבה |  |
| 6:14.171 | פרשת |  |
| 6:14.172 | קרח |  |
| 6:14.173 | דרשו |  |
| 6:14.174 | רז"ל |  |
| 6:14.175 | ששים |  |
| 6:14.176 | המה |  |
| 6:14.177 | מלכות |  |
| 6:14.178 | אלו |  |
| 6:14.179 | ששים |  |
| 6:14.180 | מסכתות |  |
| 6:14.181 | ושמנים |  |
| 6:14.182 | פלגשים |  |
| 6:14.183 | אלו |  |
| 6:14.184 | שמנים |  |
| 6:14.185 | בתי |  |
| 6:14.186 | מדרשות |  |
| 6:14.187 | שהיו |  |
| 6:14.188 | בירושלים |  |
| 6:14.189 | כנגד |  |
| 6:14.190 | פתחיה | , |
| 6:14.191 | ועלמות |  |
| 6:14.192 | אין |  |
| 6:14.193 | מספר |  |
| 6:14.194 | זו |  |
| 6:14.195 | משנה |  |
| 6:14.196 | החיצונה | , |
| 6:14.197 | ופירשו |  |
| 6:14.198 | המתנות |  |
| 6:14.199 | כהונה | , |
| 6:14.200 | והרב |  |
| 6:14.201 | עץ |  |
| 6:14.202 | חיים | , |
| 6:14.203 | דקרי |  |
| 6:14.204 | לבתי |  |
| 6:14.205 | מדרשות |  |
| 6:14.206 | פילגשים |  |
| 6:14.207 | על |  |
| 6:14.208 | שם |  |
| 6:14.209 | שנושאים |  |
| 6:14.210 | ונותנים |  |
| 6:14.211 | וחולקים |  |
| 6:14.212 | זה |  |
| 6:14.213 | עם |  |
| 6:14.214 | זה |  |
| 6:14.215 | כמו |  |
| 6:14.216 | פילגשם | , |
| 6:14.217 | וכי |  |
| 6:14.218 | משנה |  |
| 6:14.219 | החיצונה |  |
| 6:14.220 | מה |  |
| 6:14.221 | שלמדו |  |
| 6:14.222 | היחידים |  |
| 6:14.223 | חוץ |  |
| 6:14.224 | לאלו |  |
| 6:14.225 | בתי |  |
| 6:14.226 | המדרשות | . |
| 6:14.227 | וביד |  |
| 6:14.228 | יוסף |  |
| 6:14.229 | פי | ' |
| 6:14.230 | שכוונת |  |
| 6:14.231 | המדרש |  |
| 6:14.232 | על |  |
| 6:14.233 | הברייתות |  |
| 6:14.234 | וע"כ | . |
| 6:14.235 | וכן |  |
| 6:14.236 | מצאתי |  |
| 6:14.237 | סיוע |  |
| 6:14.238 | לפירושו |  |
| 6:14.239 | בשיר |  |
| 6:14.240 | השירים |  |
| 6:14.241 | רבה |  |
| 6:14.242 | וז"ל |  |
| 6:14.243 | ששים |  |
| 6:14.244 | המה |  |
| 6:14.245 | מלכות |  |
| 6:14.246 | אלו |  |
| 6:14.247 | ששים |  |
| 6:14.248 | מסכתות |  |
| 6:14.249 | של |  |
| 6:14.250 | הלכות | , |
| 6:14.251 | ושמונים |  |
| 6:14.252 | פילגשים |  |
| 6:14.253 | אלו |  |
| 6:14.254 | שמונים |  |
| 6:14.255 | פרשיות |  |
| 6:14.256 | שבתורת |  |
| 6:14.257 | כהנים | . |
| 6:14.258 | ועלמות |  |
| 6:14.259 | אין |  |
| 6:14.260 | מספר |  |
| 6:14.261 | אין |  |
| 6:14.262 | קץ |  |
| 6:14.263 | לתוספתות | . |
| 6:14.264 | אחת |  |
| 6:14.265 | היא |  |
| 6:14.266 | הן |  |
| 6:14.267 | חולקין |  |
| 6:14.268 | אלו |  |
| 6:14.269 | עם |  |
| 6:14.270 | אלו |  |
| 6:14.271 | וכולהון |  |
| 6:14.272 | דורשים |  |
| 6:14.273 | מטעם |  |
| 6:14.274 | אחד | , |
| 6:14.275 | מהלכה |  |
| 6:14.276 | אחת | , |
| 6:14.277 | מגזירה |  |
| 6:14.278 | שוה | , |
| 6:14.279 | מקל |  |
| 6:14.280 | וחומר |  |
| 6:14.281 | ע"כ | . |
| 6:14.282 | ראה |  |
| 6:14.283 | כמה |  |
| 6:14.284 | מחבבים |  |
| 6:14.285 | רז"ל |  |
| 6:14.286 | את |  |
| 6:14.287 | המשנה |  |
| 6:14.288 | והתלמוד |  |
| 6:14.289 | וקראו |  |
| 6:14.290 | להן |  |
| 6:14.291 | מלכות | . |
| 6:14.292 | ולא |  |
| 6:14.293 | שפחות |  |
| 6:14.294 | כאשר |  |
| 6:14.295 | עשה |  |
| 6:14.296 | הפילסוף |  |
| 6:14.297 | הרמאי |  |
| 6:14.298 | האדוק |  |
| 6:14.299 | בדעות |  |
| 6:14.300 | חיצוניות |  |
| 6:14.301 | הוא |  |
| 6:14.302 | מחבר |  |
| 6:14.303 | הזהר |  |
| 6:14.304 | כמו |  |
| 6:14.305 | שיתבאר |  |
| 6:14.306 | לקמן | . |
| 6:14.307 | ובבא |  |
| 6:14.308 | מציעא |  |
| 6:14.309 | פרק |  |
| 6:14.310 | אלו |  |
| 6:14.311 | מציאות |  |
| 6:14.312 | בבא |  |
| 6:14.313 | מציעא |  |
| 6:14.314 | דף |  |
| 6:14.315 | ל"ג | ) |
| 6:14.316 | גרסינן | . |
| 6:14.317 | ת"ר |  |
| 6:14.318 | העוסקים |  |
| 6:14.319 | במקרא |  |
| 6:14.320 | מדה |  |
| 6:14.321 | ואינה |  |
| 6:14.322 | מדה |  |
| 6:14.323 | במשנה |  |
| 6:14.324 | מדה |  |
| 6:14.325 | ונוטל |  |
| 6:14.326 | עליה |  |
| 6:14.327 | שכר | , |
| 6:14.328 | בתלמוד | , |
| 6:14.329 | אין |  |
| 6:14.330 | לך |  |
| 6:14.331 | מדה |  |
| 6:14.332 | גדולה |  |
| 6:14.333 | מזו | , |
| 6:14.334 | ולעולם |  |
| 6:14.335 | הוי |  |
| 6:14.336 | רץ |  |
| 6:14.337 | למשנה |  |
| 6:14.338 | יותר |  |
| 6:14.339 | מן |  |
| 6:14.340 | התלמוד | , |
| 6:14.341 | ופריך |  |
| 6:14.342 | היא |  |
| 6:14.343 | גופה |  |
| 6:14.344 | קשיא | , |
| 6:14.345 | אמרת |  |
| 6:14.346 | בתלמוד |  |
| 6:14.347 | אין |  |
| 6:14.348 | לך |  |
| 6:14.349 | מדה |  |
| 6:14.350 | גדולה |  |
| 6:14.351 | מזו | , |
| 6:14.352 | והדר |  |
| 6:14.353 | אמרת |  |
| 6:14.354 | ולעולם |  |
| 6:14.355 | הוי |  |
| 6:14.356 | רץ |  |
| 6:14.357 | למשנה |  |
| 6:14.358 | יותר |  |
| 6:14.359 | מן |  |
| 6:14.360 | התלמוד | ? |
| 6:14.361 | אמר |  |
| 6:14.362 | ר | ' |
| 6:14.363 | יוחנן |  |
| 6:14.364 | בימי |  |
| 6:14.365 | ר | ' |
| 6:14.366 | נשנית |  |
| 6:14.367 | משנה |  |
| 6:14.368 | זו | , |
| 6:14.369 | שבקו |  |
| 6:14.370 | כולי |  |
| 6:14.371 | עלמא |  |
| 6:14.372 | מתניתין |  |
| 6:14.373 | ואזלו |  |
| 6:14.374 | בתר |  |
| 6:14.375 | תלמודא | , |
| 6:14.376 | הדר |  |
| 6:14.377 | דרש |  |
| 6:14.378 | להו |  |
| 6:14.379 | ולעולם |  |
| 6:14.380 | הוי |  |
| 6:14.381 | רץ |  |
| 6:14.382 | למשנה |  |
| 6:14.383 | יותר |  |
| 6:14.384 | מן |  |
| 6:14.385 | התלמוד |  |
| 6:14.386 | עיין |  |
| 6:14.387 | שם |  |
| 6:14.388 | ובפי | ' |
| 6:14.389 | רש"י | . |
| 6:14.390 | ובירושלמי |  |
| 6:14.391 | מבואר |  |
| 6:14.392 | בהדיא |  |
| 6:14.393 | דרשב"י |  |
| 6:14.394 | הוא |  |
| 6:14.395 | דתני |  |
| 6:14.396 | לה | , |
| 6:14.397 | כדגרסינן |  |
| 6:14.398 | התם |  |
| 6:14.399 | בפרק |  |
| 6:14.400 | כל |  |
| 6:14.401 | כתבי |  |
| 6:14.402 | הדא |  |
| 6:14.403 | אמרא |  |
| 6:14.404 | שהמשנה |  |
| 6:14.405 | קודמת |  |
| 6:14.406 | למקרא | , |
| 6:14.407 | ודא |  |
| 6:14.408 | מסייעא |  |
| 6:14.409 | לרשב"י |  |
| 6:14.410 | דתני |  |
| 6:14.411 | רשב"י |  |
| 6:14.412 | העוסק |  |
| 6:14.413 | במקרא |  |
| 6:14.414 | מדה |  |
| 6:14.415 | שאינה |  |
| 6:14.416 | מדה | , |
| 6:14.417 | במשנה |  |
| 6:14.418 | מדה |  |
| 6:14.419 | שנוטלין |  |
| 6:14.420 | ממנה |  |
| 6:14.421 | שכר | , |
| 6:14.422 | העוסק |  |
| 6:14.423 | בתלמוד |  |
| 6:14.424 | אין |  |
| 6:14.425 | לך |  |
| 6:14.426 | מדה |  |
| 6:14.427 | גדולה |  |
| 6:14.428 | מזו | , |
| 6:14.429 | לעולם |  |
| 6:14.430 | הוי |  |
| 6:14.431 | רץ |  |
| 6:14.432 | אחר |  |
| 6:14.433 | המשנה |  |
| 6:14.434 | יותר |  |
| 6:14.435 | מן |  |
| 6:14.436 | התלמוד |  |
| 6:14.437 | אמר |  |
| 6:14.438 | ר | ' |
| 6:14.439 | יוסי |  |
| 6:14.440 | בר | ' |
| 6:14.441 | בון |  |
| 6:14.442 | הדא |  |
| 6:14.443 | דאת |  |
| 6:14.444 | אמר |  |
| 6:14.445 | עד |  |
| 6:14.446 | שלא |  |
| 6:14.447 | שקע |  |
| 6:14.448 | בו |  |
| 6:14.449 | רבי |  |
| 6:14.450 | רוב |  |
| 6:14.451 | משניות | , |
| 6:14.452 | אבל |  |
| 6:14.453 | מששיקע |  |
| 6:14.454 | בו |  |
| 6:14.455 | רבי |  |
| 6:14.456 | רוב |  |
| 6:14.457 | משניות |  |
| 6:14.458 | לעולם |  |
| 6:14.459 | הוי |  |
| 6:14.460 | רץ |  |
| 6:14.461 | אחר |  |
| 6:14.462 | התלמוד |  |
| 6:14.463 | יותר |  |
| 6:14.464 | מן |  |
| 6:14.465 | המשנה |  |
| 6:14.466 | ע"כ | . |
| 6:14.467 | וכן |  |
| 6:14.468 | בפ"ק |  |
| 6:14.469 | דשבת |  |
| 6:14.470 | סוף |  |
| 6:14.471 | פסקא |  |
| 6:14.472 | דלא |  |
| 6:14.473 | ישב |  |
| 6:14.474 | אדם |  |
| 6:14.475 | לפני |  |
| 6:14.476 | הספר |  |
| 6:14.477 | קאמר |  |
| 6:14.478 | התם |  |
| 6:14.479 | דהקורא |  |
| 6:14.480 | את |  |
| 6:14.481 | שמע |  |
| 6:14.482 | בעונתה |  |
| 6:14.483 | כמשנה |  |
| 6:14.484 | היא | , |
| 6:14.485 | ושלא |  |
| 6:14.486 | בעונתה |  |
| 6:14.487 | כאדם |  |
| 6:14.488 | שקורא |  |
| 6:14.489 | בתורה |  |
| 6:14.490 | ששכרו |  |
| 6:14.491 | פחות |  |
| 6:14.492 | מהעוסק |  |
| 6:14.493 | במשנה | , |
| 6:14.494 | ומסיים |  |
| 6:14.495 | בה |  |
| 6:14.496 | רשב"י |  |
| 6:14.497 | כדעתיה |  |
| 6:14.498 | דרשב"י |  |
| 6:14.499 | אמר |  |
| 6:14.500 | העוסק |  |
| 6:14.501 | במקרא |  |
| 6:14.502 | מדה |  |
| 6:14.503 | שאינה |  |
| 6:14.504 | מדה | , |
| 6:14.505 | ורבנין |  |
| 6:14.506 | עבדין |  |
| 6:14.507 | מקרא |  |
| 6:14.508 | כמשנה |  |
| 6:14.509 | עכ"ל | , |
| 6:14.510 | ובילקוט |  |
| 6:14.511 | ריש |  |
| 6:14.512 | פר | ' |
| 6:14.513 | צו |  |
| 6:14.514 | ר | ' |
| 6:14.515 | חנינה |  |
| 6:14.516 | אמר |  |
| 6:14.517 | אין |  |
| 6:14.518 | הגליות |  |
| 6:14.519 | מתכנסות |  |
| 6:14.520 | אלא |  |
| 6:14.521 | בזכות |  |
| 6:14.522 | משניות |  |
| 6:14.523 | שנ | ' |
| 6:14.524 | גם |  |
| 6:14.525 | כי |  |
| 6:14.526 | יתנו |  |
| 6:14.527 | בגוים |  |
| 6:14.528 | עתה |  |
| 6:14.529 | אקבצם |  |
| 6:14.530 | ע"כ | . |
| 6:14.531 | ראה |  |
| 6:14.532 | תראה |  |
| 6:14.533 | ידידי |  |
| 6:14.534 | הקורא |  |
| 6:14.535 | דברי |  |
| 6:14.536 | רז"ל |  |
| 6:14.537 | שהבאתי |  |
| 6:14.538 | מקצת |  |
| 6:14.539 | מהם |  |
| 6:14.540 | כמה |  |
| 6:14.541 | הם |  |
| 6:14.542 | מחבבים |  |
| 6:14.543 | ומחשיבים |  |
| 6:14.544 | את |  |
| 6:14.545 | המשנה |  |
| 6:14.546 | שהיא |  |
| 6:14.547 | עיקר |  |
| 6:14.548 | תורה |  |
| 6:14.549 | שבעל |  |
| 6:14.550 | פה |  |
| 6:14.551 | שנאמרה |  |
| 6:14.552 | למשה |  |
| 6:14.553 | בסיני | , |
| 6:14.554 | ובפרט |  |
| 6:14.555 | שים |  |
| 6:14.556 | עיניך |  |
| 6:14.557 | לדברי |  |
| 6:14.558 | תנא |  |
| 6:14.559 | דידן |  |
| 6:14.560 | רשב"י |  |
| 6:14.561 | הנאמרים |  |
| 6:14.562 | באמת |  |
| 6:14.563 | בבבלי |  |
| 6:14.564 | ובירושלמי |  |
| 6:14.565 | כמה |  |
| 6:14.566 | הוא |  |
| 6:14.567 | מחבב |  |
| 6:14.568 | ומחשיב |  |
| 6:14.569 | את |  |
| 6:14.570 | המשנה |  |
| 6:14.571 | יותר |  |
| 6:14.572 | מן |  |
| 6:14.573 | המקרא | , |
| 6:14.574 | ואת |  |
| 6:14.575 | התלמוד |  |
| 6:14.576 | שהוא |  |
| 6:14.577 | פי | ' |
| 6:14.578 | טעמי |  |
| 6:14.579 | המשנה |  |
| 6:14.580 | יותר |  |
| 6:14.581 | מכולם |  |
| 6:14.582 | שאין |  |
| 6:14.583 | לך |  |
| 6:14.584 | מדה |  |
| 6:14.585 | גדולה |  |
| 6:14.586 | יותר |  |
| 6:14.587 | ממנה | , |
| 6:14.588 | אלא |  |
| 6:14.589 | דבירושלמי |  |
| 6:14.590 | מוקי |  |
| 6:14.591 | להא |  |
| 6:14.592 | דאמר |  |
| 6:14.593 | לעולם |  |
| 6:14.594 | הוי |  |
| 6:14.595 | רץ |  |
| 6:14.596 | למשנה |  |
| 6:14.597 | יותר |  |
| 6:14.598 | מן |  |
| 6:14.599 | התלמוד |  |
| 6:14.600 | קודם |  |
| 6:14.601 | שסידר |  |
| 6:14.602 | רבי |  |
| 6:14.603 | את |  |
| 6:14.604 | המשנה |  |
| 6:14.605 | שאז |  |
| 6:14.606 | היו |  |
| 6:14.607 | הרבה |  |
| 6:14.608 | משניות |  |
| 6:14.609 | שהם |  |
| 6:14.610 | שגורים |  |
| 6:14.611 | בפי |  |
| 6:14.612 | התלמידים |  |
| 6:14.613 | כל |  |
| 6:14.614 | אחת |  |
| 6:14.615 | בשם |  |
| 6:14.616 | אומרה | , |
| 6:14.617 | והרבה |  |
| 6:14.618 | בריתות |  |
| 6:14.619 | היו |  |
| 6:14.620 | להם | , |
| 6:14.621 | אבל |  |
| 6:14.622 | לאחר |  |
| 6:14.623 | שסידר |  |
| 6:14.624 | רבי |  |
| 6:14.625 | את |  |
| 6:14.626 | המשנה |  |
| 6:14.627 | והשמיט |  |
| 6:14.628 | בה |  |
| 6:14.629 | כמה |  |
| 6:14.630 | משניות |  |
| 6:14.631 | ובריתות | , |
| 6:14.632 | וגם |  |
| 6:14.633 | סתם |  |
| 6:14.634 | במשנה |  |
| 6:14.635 | דברי |  |
| 6:14.636 | יחידים |  |
| 6:14.637 | שראה |  |
| 6:14.638 | ר | ' |
| 6:14.639 | דבריהם |  |
| 6:14.640 | ושנאן |  |
| 6:14.641 | סתם |  |
| 6:14.642 | כדי |  |
| 6:14.643 | לקבוע |  |
| 6:14.644 | הלכה |  |
| 6:14.645 | כמותם | , |
| 6:14.646 | לפיכך |  |
| 6:14.647 | חזרו |  |
| 6:14.648 | לומר |  |
| 6:14.649 | הוי |  |
| 6:14.650 | רץ |  |
| 6:14.651 | לתלמוד |  |
| 6:14.652 | יותר |  |
| 6:14.653 | מן |  |
| 6:14.654 | המשנה | . |
| 6:14.655 | ומדברי |  |
| 6:14.656 | כולם |  |
| 6:14.657 | למדנו |  |
| 6:14.658 | כמה |  |
| 6:14.659 | היה |  |
| 6:14.660 | חביב |  |
| 6:14.661 | להם |  |
| 6:14.662 | לימוד |  |
| 6:14.663 | המשנה |  |
| 6:14.664 | והתלמוד | , |
| 6:14.665 | ושרשב"י |  |
| 6:14.666 | הוא |  |
| 6:14.667 | דתני |  |
| 6:14.668 | לה | , |
| 6:14.669 | ומזה |  |
| 6:14.670 | הוכחה |  |
| 6:14.671 | גמורה |  |
| 6:14.672 | ברורה |  |
| 6:14.673 | שס | ' |
| 6:14.674 | הזהר |  |
| 6:14.675 | אינו |  |
| 6:14.676 | לרשב"י |  |
| 6:14.677 | התנא |  |
| 6:14.678 | ע"ה | , |
| 6:14.679 | כי |  |
| 6:14.680 | אם |  |
| 6:14.681 | לפילסוף |  |
| 6:14.682 | רמאי |  |
| 6:14.683 | גדול |  |
| 6:14.684 | שלמד |  |
| 6:14.685 | כל |  |
| 6:14.686 | חכמת |  |
| 6:14.687 | ישראל |  |
| 6:14.688 | בשני |  |
| 6:14.689 | התלמודים |  |
| 6:14.690 | ומדרשי |  |
| 6:14.691 | רז"ל | . |
| 6:14.692 | והוא |  |
| 6:14.693 | מאמין |  |
| 6:14.694 | בדעות |  |
| 6:14.695 | הכשדים |  |
| 6:14.696 | והכלדיים |  |
| 6:14.697 | פילסופים |  |
| 6:14.698 | הראשונים |  |
| 6:14.699 | שהיו |  |
| 6:14.700 | חולקים |  |
| 6:14.701 | על |  |
| 6:14.702 | אברהם |  |
| 6:14.703 | אבינו |  |
| 6:14.704 | באמונתו |  |
| 6:14.705 | באל |  |
| 6:14.706 | אחד | , |
| 6:14.707 | ומאמינים |  |
| 6:14.708 | בשתי |  |
| 6:14.709 | סבות |  |
| 6:14.710 | קודמות | , |
| 6:14.711 | וכל |  |
| 6:14.712 | אחת |  |
| 6:14.713 | פועלת |  |
| 6:14.714 | היפך |  |
| 6:14.715 | האחרת | , |
| 6:14.716 | והסבות |  |
| 6:14.717 | הפועלות |  |
| 6:14.718 | הטוב |  |
| 6:14.719 | קורא |  |
| 6:14.720 | להם |  |
| 6:14.721 | בשם |  |
| 6:14.722 | סטרא |  |
| 6:14.723 | דקדושה | , |
| 6:14.724 | והם |  |
| 6:14.725 | אדם |  |
| 6:14.726 | קדמון |  |
| 6:14.727 | וכתר |  |
| 6:14.728 | עליון |  |
| 6:14.729 | עם |  |
| 6:14.730 | שאר |  |
| 6:14.731 | הפרצופים |  |
| 6:14.732 | וצורות |  |
| 6:14.733 | המשתלשלים |  |
| 6:14.734 | מהם | , |
| 6:14.735 | והסבות |  |
| 6:14.736 | הפועלות |  |
| 6:14.737 | הרע |  |
| 6:14.738 | הוא |  |
| 6:14.739 | קורא |  |
| 6:14.740 | להם |  |
| 6:14.741 | סטרא |  |
| 6:14.742 | דקליפה | , |
| 6:14.743 | והם |  |
| 6:14.744 | אדם |  |
| 6:14.745 | בליעל |  |
| 6:14.746 | וכתר |  |
| 6:14.747 | עליון |  |
| 6:14.748 | דקליפה |  |
| 6:14.749 | עם |  |
| 6:14.750 | שאר |  |
| 6:14.751 | פרצופים |  |
| 6:14.752 | וצורות |  |
| 6:14.753 | דעשר |  |
| 6:14.754 | ספירות |  |
| 6:14.755 | דקליפה |  |
| 6:14.756 | המשתלשלים |  |
| 6:14.757 | מהם |  |
| 6:14.758 | הנקראים |  |
| 6:14.759 | בשם |  |
| 6:14.760 | אל |  |
| 6:14.761 | אחר | , |
| 6:14.762 | כי |  |
| 6:14.763 | הוא |  |
| 6:14.764 | מאמין |  |
| 6:14.765 | ומאמת |  |
| 6:14.766 | מציאות |  |
| 6:14.767 | אלהים |  |
| 6:14.768 | אחרים |  |
| 6:14.769 | בעולם |  |
| 6:14.770 | כדעת |  |
| 6:14.771 | הכשדים |  |
| 6:14.772 | והמצרים |  |
| 6:14.773 | הקדמונים |  |
| 6:14.774 | בפילוסופיא |  |
| 6:14.775 | הקדומה |  |
| 6:14.776 | שכבר |  |
| 6:14.777 | נמר |  |
| 6:14.778 | ריחה |  |
| 6:14.779 | וסר |  |
| 6:14.780 | טעמה | . |
| 6:14.781 | ובשפתי |  |
| 6:14.782 | חלקות |  |
| 6:14.783 | וגיזומים |  |
| 6:14.784 | הוא |  |
| 6:14.785 | מושך |  |
| 6:14.786 | לב |  |
| 6:14.787 | רבים |  |
| 6:14.788 | מהרבנים |  |
| 6:14.789 | הפתאים |  |
| 6:14.790 | להאמין |  |
| 6:14.791 | באלהים |  |
| 6:14.792 | אחרים |  |
| 6:14.793 | רבים |  |
| 6:14.794 | ולזלזל |  |
| 6:14.795 | בלמוד |  |
| 6:14.796 | המשנה |  |
| 6:14.797 | והתלמוד |  |
| 6:14.798 | עד |  |
| 6:14.799 | שלא |  |
| 6:14.800 | ידעו |  |
| 6:14.801 | מה |  |
| 6:14.802 | לעשות |  |
| 6:14.803 | והיאך |  |
| 6:14.804 | יפרשו |  |
| 6:14.805 | מן |  |
| 6:14.806 | האיסורי | ' |
| 6:14.807 | שנאמרו |  |
| 6:14.808 | למשה |  |
| 6:14.809 | בסיני |  |
| 6:14.810 | כגון |  |
| 6:14.811 | ח"י |  |
| 6:14.812 | טריפות |  |
| 6:14.813 | הנזכרים |  |
| 6:14.814 | במשנה |  |
| 6:14.815 | וחילוקי |  |
| 6:14.816 | דיניהם | , |
| 6:14.817 | לפי |  |
| 6:14.818 | שלא |  |
| 6:14.819 | נתפרש |  |
| 6:14.820 | בתורה |  |
| 6:14.821 | בהדיא |  |
| 6:14.822 | אלא |  |
| 6:14.823 | דרוסה |  |
| 6:14.824 | כמו |  |
| 6:14.825 | שכ | ' |
| 6:14.826 | הרמב"ם |  |
| 6:14.827 | ז"ל |  |
| 6:14.828 | בהל | ' |
| 6:14.829 | מאכלות |  |
| 6:14.830 | אסורות | . |
| 6:14.831 | ויכוצא |  |
| 6:14.832 | בזה |  |
| 6:14.833 | בשאר |  |
| 6:14.834 | איסורים |  |
| 6:14.835 | של |  |
| 6:14.836 | תורה |  |
| 6:14.837 | שלא |  |
| 6:14.838 | נתפרשו |  |
| 6:14.839 | בתורה |  |
| 6:14.840 | שבכתב |  |
| 6:14.841 | ונמסרו |  |
| 6:14.842 | למשה |  |
| 6:14.843 | רבי | ' |
| 6:14.844 | ע"ה |  |
| 6:14.845 | בעל |  |
| 6:14.846 | פה | . |
| 6:14.847 | והמה |  |
| 6:14.848 | מפורשים |  |
| 6:14.849 | במשנה |  |
| 6:14.850 | ובתלמוד |  |
| 6:14.851 | כגון |  |
| 6:14.852 | דיני |  |
| 6:14.853 | תערובת |  |
| 6:14.854 | איסור |  |
| 6:14.855 | ובשר |  |
| 6:14.856 | בחלב | . |
| 6:14.857 | ואיזה |  |
| 6:14.858 | מהן |  |
| 6:14.859 | אסור |  |
| 6:14.860 | מדאוריתא |  |
| 6:14.861 | ואיזה |  |
| 6:14.862 | אסור |  |
| 6:14.863 | מדרבנן | . |
| 6:14.864 | והפילסוף |  |
| 6:14.865 | הנז | ' |
| 6:14.866 | כיוון |  |
| 6:14.867 | לפסול |  |
| 6:14.868 | ולהשכיח |  |
| 6:14.869 | מישראל |  |
| 6:14.870 | כל |  |
| 6:14.871 | הדינים |  |
| 6:14.872 | המבוארים |  |
| 6:14.873 | במשנה |  |
| 6:14.874 | ובתלמוד | . |
| 6:14.875 | ובהדיא |  |
| 6:14.876 | אמר |  |
| 6:14.877 | שבסוף |  |
| 6:14.878 | גלותא |  |
| 6:14.879 | יתפרנסון |  |
| 6:14.880 | מחבורא |  |
| 6:14.881 | דיליה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

