# Translation pack: milhamot-hashem perek-8-14 (Milchamot ha-Shem 8:17–8:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-14`.

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
- **קצר אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
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


## Milchamot ha-Shem 8:17

Hebrew (Wikisource): קיא) ואלה הדעות הן הנה הדעות הרעות שהיו לאבותינו בימי השופטים ובימי ירבעם בן נבט ושאר מלכי ישראל וקצת מלכי יהודה בזמן בית ראשון אשר עליהם נתנבאו כל הנביאים, וקראו בגרון הוי החוקקים חקקי און, ומכתבי עמל כתבו, והיה להם דבר ה' צו לצו לצו קו לקו קו זעיר שם זעיר וכמאמר רז"ל אין צו אלא ע"ז שנאמר כי הואיל הלך אחרי צו וכו' וכן קו לקו הם הפרצופים בעלי שלשה קוים המוזכרים בדברי המקובלים כמו שכתב מהרח"ו בעץ חיים שער עגולים ויושר, שספירות היושר הם בציור שלשה קוים, כמראה האדם בעל ראש, וזרועות ושוקים וגוף ורגלים וכו', וכן זעיר שם זעיר שם הוא זעיר אנפין, (קצר אפים) ונאמר ויעזבו את ה' אלהיהם ויעבדו את הבעלים ואת העשתרות, הבעלים הם האלוהות הזכרים, אבא וזעיר אנפין, ונקראו בעלים על שם הזווגים והבעילות שמזדווגין עם הנקיבות, כמו שאמרו רז"ל בירושלמי שבת פרק אמר ר' עקיבא ובמסכת אבות ע"ז פרק כל הצלמים וז"ל אמר ר' הונא רבי חמא בשם רב, הבעל ראש גוייה היה לו ונאפון הוא, וישימו להם בעל ברית לאלהים, עכ"ל הירושלמי הרי מבואר בהדיא מדברי הירוש' שנק' הצלם הזכר מעל על שם שיש לו ראש גויה והוא בועל ומנאף, וידוע הוא כי האלילים הקדמונים עשו להם צלמים מכוונים נגד הצורות והפרצופים שחשבו שיש בשמים, כדי להמשיך השפע מהצורות העליונות להם, כי הם עשוים בצלמם ודמותם, וגם כן נקראו בעלים על שם שעובדיהם שמו אותם אדונים להם כמו שכתב הרד"ק וכלי יקר. ואת העשתרות הם נקבות האלוהות שהם אימא ומלכית, כי עשתרות צאנך הם נקבות הצאן כמו שכתבו המפרשים, וכתן נאמר ואת העשתרות בלי סמיכות מלת הצאן, והם נקיבות האלוהות הנזכרים. וגם בין הכשרים היו האמונות האלו פושטים ולכן אמר נבוכדנאצר הרשע בעת שראה מלאך ה' בכבשן האש שירד להציל לחנניה מיאל ועזריה מן האור, וריויה די רביעאה דמי לבר אלהין, לפי שחשב שהזעיר אנפין הנקרא בן לאבא ואימא ירד להצילם מן האש, כדגרסינן בירושלמי שבת פרק במה אשה, וריוה די רביעאה דמי לבא אלהין אמר רבי ראובן באותה שעה ירד מלאך וסטרו לההוא רשיעא על פיו ואמר ליה תקין מילך, ובר אית ליה? חזר ואמר בריך אלההון די שדרך מישך ועביד נגו, די שלח מלאכיה ושיזיב לעבדוהי, בריה אין כתיב כאן אלא מלאכיה, ע"כ. ומזה יתברר לך שהדעות האלו של מחבר הזהר והמקובלים החדשים, זרים המה לדעת תורתינו הקדושה, ומהעמים אשר השאיר ה' לנסות בם את ישראל באו לנו, ונטו ישראל אחריהם בימי שפוט השופטים, וגם מן הכשדים שגלו ביניהם בעת חרבן בית ראשון.

| id | word | punctuation after |
|---|---|---|
| 8:17.1 | קיא | ) |
| 8:17.2 | ואלה |  |
| 8:17.3 | הדעות |  |
| 8:17.4 | הן |  |
| 8:17.5 | הנה |  |
| 8:17.6 | הדעות |  |
| 8:17.7 | הרעות |  |
| 8:17.8 | שהיו |  |
| 8:17.9 | לאבותינו |  |
| 8:17.10 | בימי |  |
| 8:17.11 | השופטים |  |
| 8:17.12 | ובימי |  |
| 8:17.13 | ירבעם |  |
| 8:17.14 | בן |  |
| 8:17.15 | נבט |  |
| 8:17.16 | ושאר |  |
| 8:17.17 | מלכי |  |
| 8:17.18 | ישראל |  |
| 8:17.19 | וקצת |  |
| 8:17.20 | מלכי |  |
| 8:17.21 | יהודה |  |
| 8:17.22 | בזמן |  |
| 8:17.23 | בית |  |
| 8:17.24 | ראשון |  |
| 8:17.25 | אשר |  |
| 8:17.26 | עליהם |  |
| 8:17.27 | נתנבאו |  |
| 8:17.28 | כל |  |
| 8:17.29 | הנביאים | , |
| 8:17.30 | וקראו |  |
| 8:17.31 | בגרון |  |
| 8:17.32 | הוי |  |
| 8:17.33 | החוקקים |  |
| 8:17.34 | חקקי |  |
| 8:17.35 | און | , |
| 8:17.36 | ומכתבי |  |
| 8:17.37 | עמל |  |
| 8:17.38 | כתבו | , |
| 8:17.39 | והיה |  |
| 8:17.40 | להם |  |
| 8:17.41 | דבר |  |
| 8:17.42 | ה | ' |
| 8:17.43 | צו |  |
| 8:17.44 | לצו |  |
| 8:17.45 | לצו |  |
| 8:17.46 | קו |  |
| 8:17.47 | לקו |  |
| 8:17.48 | קו |  |
| 8:17.49 | זעיר |  |
| 8:17.50 | שם |  |
| 8:17.51 | זעיר |  |
| 8:17.52 | וכמאמר |  |
| 8:17.53 | רז"ל |  |
| 8:17.54 | אין |  |
| 8:17.55 | צו |  |
| 8:17.56 | אלא |  |
| 8:17.57 | ע"ז |  |
| 8:17.58 | שנאמר |  |
| 8:17.59 | כי |  |
| 8:17.60 | הואיל |  |
| 8:17.61 | הלך |  |
| 8:17.62 | אחרי |  |
| 8:17.63 | צו |  |
| 8:17.64 | וכו | ' |
| 8:17.65 | וכן |  |
| 8:17.66 | קו |  |
| 8:17.67 | לקו |  |
| 8:17.68 | הם |  |
| 8:17.69 | הפרצופים |  |
| 8:17.70 | בעלי |  |
| 8:17.71 | שלשה |  |
| 8:17.72 | קוים |  |
| 8:17.73 | המוזכרים |  |
| 8:17.74 | בדברי |  |
| 8:17.75 | המקובלים |  |
| 8:17.76 | כמו |  |
| 8:17.77 | שכתב |  |
| 8:17.78 | מהרח"ו |  |
| 8:17.79 | בעץ |  |
| 8:17.80 | חיים |  |
| 8:17.81 | שער |  |
| 8:17.82 | עגולים |  |
| 8:17.83 | ויושר | , |
| 8:17.84 | שספירות |  |
| 8:17.85 | היושר |  |
| 8:17.86 | הם |  |
| 8:17.87 | בציור |  |
| 8:17.88 | שלשה |  |
| 8:17.89 | קוים | , |
| 8:17.90 | כמראה |  |
| 8:17.91 | האדם |  |
| 8:17.92 | בעל |  |
| 8:17.93 | ראש | , |
| 8:17.94 | וזרועות |  |
| 8:17.95 | ושוקים |  |
| 8:17.96 | וגוף |  |
| 8:17.97 | ורגלים |  |
| 8:17.98 | וכו | ', |
| 8:17.99 | וכן |  |
| 8:17.100 | זעיר |  |
| 8:17.101 | שם |  |
| 8:17.102 | זעיר |  |
| 8:17.103 | שם |  |
| 8:17.104 | הוא |  |
| 8:17.105 | זעיר |  |
| 8:17.106 | אנפין | , |
| 8:17.107 | קצר |  |
| 8:17.108 | אפים | ) |
| 8:17.109 | ונאמר |  |
| 8:17.110 | ויעזבו |  |
| 8:17.111 | את |  |
| 8:17.112 | ה | ' |
| 8:17.113 | אלהיהם |  |
| 8:17.114 | ויעבדו |  |
| 8:17.115 | את |  |
| 8:17.116 | הבעלים |  |
| 8:17.117 | ואת |  |
| 8:17.118 | העשתרות | , |
| 8:17.119 | הבעלים |  |
| 8:17.120 | הם |  |
| 8:17.121 | האלוהות |  |
| 8:17.122 | הזכרים | , |
| 8:17.123 | אבא |  |
| 8:17.124 | וזעיר |  |
| 8:17.125 | אנפין | , |
| 8:17.126 | ונקראו |  |
| 8:17.127 | בעלים |  |
| 8:17.128 | על |  |
| 8:17.129 | שם |  |
| 8:17.130 | הזווגים |  |
| 8:17.131 | והבעילות |  |
| 8:17.132 | שמזדווגין |  |
| 8:17.133 | עם |  |
| 8:17.134 | הנקיבות | , |
| 8:17.135 | כמו |  |
| 8:17.136 | שאמרו |  |
| 8:17.137 | רז"ל |  |
| 8:17.138 | בירושלמי |  |
| 8:17.139 | שבת |  |
| 8:17.140 | פרק |  |
| 8:17.141 | אמר |  |
| 8:17.142 | ר | ' |
| 8:17.143 | עקיבא |  |
| 8:17.144 | ובמסכת |  |
| 8:17.145 | אבות |  |
| 8:17.146 | ע"ז |  |
| 8:17.147 | פרק |  |
| 8:17.148 | כל |  |
| 8:17.149 | הצלמים |  |
| 8:17.150 | וז"ל |  |
| 8:17.151 | אמר |  |
| 8:17.152 | ר | ' |
| 8:17.153 | הונא |  |
| 8:17.154 | רבי |  |
| 8:17.155 | חמא |  |
| 8:17.156 | בשם |  |
| 8:17.157 | רב | , |
| 8:17.158 | הבעל |  |
| 8:17.159 | ראש |  |
| 8:17.160 | גוייה |  |
| 8:17.161 | היה |  |
| 8:17.162 | לו |  |
| 8:17.163 | ונאפון |  |
| 8:17.164 | הוא | , |
| 8:17.165 | וישימו |  |
| 8:17.166 | להם |  |
| 8:17.167 | בעל |  |
| 8:17.168 | ברית |  |
| 8:17.169 | לאלהים | , |
| 8:17.170 | עכ"ל |  |
| 8:17.171 | הירושלמי |  |
| 8:17.172 | הרי |  |
| 8:17.173 | מבואר |  |
| 8:17.174 | בהדיא |  |
| 8:17.175 | מדברי |  |
| 8:17.176 | הירוש | ' |
| 8:17.177 | שנק | ' |
| 8:17.178 | הצלם |  |
| 8:17.179 | הזכר |  |
| 8:17.180 | מעל |  |
| 8:17.181 | על |  |
| 8:17.182 | שם |  |
| 8:17.183 | שיש |  |
| 8:17.184 | לו |  |
| 8:17.185 | ראש |  |
| 8:17.186 | גויה |  |
| 8:17.187 | והוא |  |
| 8:17.188 | בועל |  |
| 8:17.189 | ומנאף | , |
| 8:17.190 | וידוע |  |
| 8:17.191 | הוא |  |
| 8:17.192 | כי |  |
| 8:17.193 | האלילים |  |
| 8:17.194 | הקדמונים |  |
| 8:17.195 | עשו |  |
| 8:17.196 | להם |  |
| 8:17.197 | צלמים |  |
| 8:17.198 | מכוונים |  |
| 8:17.199 | נגד |  |
| 8:17.200 | הצורות |  |
| 8:17.201 | והפרצופים |  |
| 8:17.202 | שחשבו |  |
| 8:17.203 | שיש |  |
| 8:17.204 | בשמים | , |
| 8:17.205 | כדי |  |
| 8:17.206 | להמשיך |  |
| 8:17.207 | השפע |  |
| 8:17.208 | מהצורות |  |
| 8:17.209 | העליונות |  |
| 8:17.210 | להם | , |
| 8:17.211 | כי |  |
| 8:17.212 | הם |  |
| 8:17.213 | עשוים |  |
| 8:17.214 | בצלמם |  |
| 8:17.215 | ודמותם | , |
| 8:17.216 | וגם |  |
| 8:17.217 | כן |  |
| 8:17.218 | נקראו |  |
| 8:17.219 | בעלים |  |
| 8:17.220 | על |  |
| 8:17.221 | שם |  |
| 8:17.222 | שעובדיהם |  |
| 8:17.223 | שמו |  |
| 8:17.224 | אותם |  |
| 8:17.225 | אדונים |  |
| 8:17.226 | להם |  |
| 8:17.227 | כמו |  |
| 8:17.228 | שכתב |  |
| 8:17.229 | הרד"ק |  |
| 8:17.230 | וכלי |  |
| 8:17.231 | יקר | . |
| 8:17.232 | ואת |  |
| 8:17.233 | העשתרות |  |
| 8:17.234 | הם |  |
| 8:17.235 | נקבות |  |
| 8:17.236 | האלוהות |  |
| 8:17.237 | שהם |  |
| 8:17.238 | אימא |  |
| 8:17.239 | ומלכית | , |
| 8:17.240 | כי |  |
| 8:17.241 | עשתרות |  |
| 8:17.242 | צאנך |  |
| 8:17.243 | הם |  |
| 8:17.244 | נקבות |  |
| 8:17.245 | הצאן |  |
| 8:17.246 | כמו |  |
| 8:17.247 | שכתבו |  |
| 8:17.248 | המפרשים | , |
| 8:17.249 | וכתן |  |
| 8:17.250 | נאמר |  |
| 8:17.251 | ואת |  |
| 8:17.252 | העשתרות |  |
| 8:17.253 | בלי |  |
| 8:17.254 | סמיכות |  |
| 8:17.255 | מלת |  |
| 8:17.256 | הצאן | , |
| 8:17.257 | והם |  |
| 8:17.258 | נקיבות |  |
| 8:17.259 | האלוהות |  |
| 8:17.260 | הנזכרים | . |
| 8:17.261 | וגם |  |
| 8:17.262 | בין |  |
| 8:17.263 | הכשרים |  |
| 8:17.264 | היו |  |
| 8:17.265 | האמונות |  |
| 8:17.266 | האלו |  |
| 8:17.267 | פושטים |  |
| 8:17.268 | ולכן |  |
| 8:17.269 | אמר |  |
| 8:17.270 | נבוכדנאצר |  |
| 8:17.271 | הרשע |  |
| 8:17.272 | בעת |  |
| 8:17.273 | שראה |  |
| 8:17.274 | מלאך |  |
| 8:17.275 | ה | ' |
| 8:17.276 | בכבשן |  |
| 8:17.277 | האש |  |
| 8:17.278 | שירד |  |
| 8:17.279 | להציל |  |
| 8:17.280 | לחנניה |  |
| 8:17.281 | מיאל |  |
| 8:17.282 | ועזריה |  |
| 8:17.283 | מן |  |
| 8:17.284 | האור | , |
| 8:17.285 | וריויה |  |
| 8:17.286 | די |  |
| 8:17.287 | רביעאה |  |
| 8:17.288 | דמי |  |
| 8:17.289 | לבר |  |
| 8:17.290 | אלהין | , |
| 8:17.291 | לפי |  |
| 8:17.292 | שחשב |  |
| 8:17.293 | שהזעיר |  |
| 8:17.294 | אנפין |  |
| 8:17.295 | הנקרא |  |
| 8:17.296 | בן |  |
| 8:17.297 | לאבא |  |
| 8:17.298 | ואימא |  |
| 8:17.299 | ירד |  |
| 8:17.300 | להצילם |  |
| 8:17.301 | מן |  |
| 8:17.302 | האש | , |
| 8:17.303 | כדגרסינן |  |
| 8:17.304 | בירושלמי |  |
| 8:17.305 | שבת |  |
| 8:17.306 | פרק |  |
| 8:17.307 | במה |  |
| 8:17.308 | אשה | , |
| 8:17.309 | וריוה |  |
| 8:17.310 | די |  |
| 8:17.311 | רביעאה |  |
| 8:17.312 | דמי |  |
| 8:17.313 | לבא |  |
| 8:17.314 | אלהין |  |
| 8:17.315 | אמר |  |
| 8:17.316 | רבי |  |
| 8:17.317 | ראובן |  |
| 8:17.318 | באותה |  |
| 8:17.319 | שעה |  |
| 8:17.320 | ירד |  |
| 8:17.321 | מלאך |  |
| 8:17.322 | וסטרו |  |
| 8:17.323 | לההוא |  |
| 8:17.324 | רשיעא |  |
| 8:17.325 | על |  |
| 8:17.326 | פיו |  |
| 8:17.327 | ואמר |  |
| 8:17.328 | ליה |  |
| 8:17.329 | תקין |  |
| 8:17.330 | מילך | , |
| 8:17.331 | ובר |  |
| 8:17.332 | אית |  |
| 8:17.333 | ליה | ? |
| 8:17.334 | חזר |  |
| 8:17.335 | ואמר |  |
| 8:17.336 | בריך |  |
| 8:17.337 | אלההון |  |
| 8:17.338 | די |  |
| 8:17.339 | שדרך |  |
| 8:17.340 | מישך |  |
| 8:17.341 | ועביד |  |
| 8:17.342 | נגו | , |
| 8:17.343 | די |  |
| 8:17.344 | שלח |  |
| 8:17.345 | מלאכיה |  |
| 8:17.346 | ושיזיב |  |
| 8:17.347 | לעבדוהי | , |
| 8:17.348 | בריה |  |
| 8:17.349 | אין |  |
| 8:17.350 | כתיב |  |
| 8:17.351 | כאן |  |
| 8:17.352 | אלא |  |
| 8:17.353 | מלאכיה | , |
| 8:17.354 | ע"כ | . |
| 8:17.355 | ומזה |  |
| 8:17.356 | יתברר |  |
| 8:17.357 | לך |  |
| 8:17.358 | שהדעות |  |
| 8:17.359 | האלו |  |
| 8:17.360 | של |  |
| 8:17.361 | מחבר |  |
| 8:17.362 | הזהר |  |
| 8:17.363 | והמקובלים |  |
| 8:17.364 | החדשים | , |
| 8:17.365 | זרים |  |
| 8:17.366 | המה |  |
| 8:17.367 | לדעת |  |
| 8:17.368 | תורתינו |  |
| 8:17.369 | הקדושה | , |
| 8:17.370 | ומהעמים |  |
| 8:17.371 | אשר |  |
| 8:17.372 | השאיר |  |
| 8:17.373 | ה | ' |
| 8:17.374 | לנסות |  |
| 8:17.375 | בם |  |
| 8:17.376 | את |  |
| 8:17.377 | ישראל |  |
| 8:17.378 | באו |  |
| 8:17.379 | לנו | , |
| 8:17.380 | ונטו |  |
| 8:17.381 | ישראל |  |
| 8:17.382 | אחריהם |  |
| 8:17.383 | בימי |  |
| 8:17.384 | שפוט |  |
| 8:17.385 | השופטים | , |
| 8:17.386 | וגם |  |
| 8:17.387 | מן |  |
| 8:17.388 | הכשדים |  |
| 8:17.389 | שגלו |  |
| 8:17.390 | ביניהם |  |
| 8:17.391 | בעת |  |
| 8:17.392 | חרבן |  |
| 8:17.393 | בית |  |
| 8:17.394 | ראשון | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

