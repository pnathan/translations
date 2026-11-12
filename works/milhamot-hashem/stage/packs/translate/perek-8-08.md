# Translation pack: milhamot-hashem perek-8-08 (Milchamot ha-Shem 8:10–8:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-08`.

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
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
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


## Milchamot ha-Shem 8:10

Hebrew (Wikisource): קו) ושרידי משכילי עמינו אשר הרגישו בזרות הקבלה החדשה כרבי' תם ן' יחיא, ורבי פרץ, ורבי שמשון מקינון, והרישב"ש והרש"ל הרחיקו את עצמם מללמדה כמו שאמר הריב"ש איני תוקע את עצמי וכו', כי הרגיש בזרותה מדרך תורתינו הקדושה והטהורה מכל סיג וחלאה, אך לא דבר כנגדה רק כלאחר יד באמרו וכי יש אלהות לספירות? והמשל שאמר לו דון יוסף ן' שושן לא ערב לו, שידע באמת שהוא כיחש לו, אך לא חפץ להעציבו לומר שקר דברת. ולפי' אמר לו והנה טוב מאד אם הוא כדבריך. כי ידע כי רק דחיה בעלמא קא מדחי ליה, כי גם אמרם כי לנשמתו אתה עובד דחייה בעלמא היא. והמה באמת עובדים להם ומקבלים אלהותם עליהם. ומיחדים את הכלים עם נשמותיהם כמו שכתב שושן סודות והרש"ב בפרשת בהר דף ק"ט ומשנת חסידים סוף מסכת העשייה ועז לאלהים שהגוף והנשמה חד אינון. וכ"כ היכל הברכה פרשת ויקרא. ואלו היתה מניעת הריב"ש מטעם שלא קיבלה מחכם מקובל היה לו ללמדה ולקבלה מפי דון יוסף ן' שושן! אלא ודאי דברים בגו! שראה בחכמתו שזו היא יציאה מדת תורתינו הקדושה לעבוד אלהים אחרים, מעשה ידי מחשבות אדם ורעיונותיו ודמיונותיו אשר עליהם באה האזהרה בתורה ולא תתורו אחרי לבבכם ואחרי עיניכם ובאה הקבלה אחרי לבבכם זו מינות, ואחרי עיניכם זו זנות, כי לבא ועינא תרי סרסורי דחטאה נינהו, ומן הידוע שההולך אחר דעתו להמציא לו רעיונות בתורה. שלא על דרך רז"ל שקבלו תורה שבעל פה שהיא המשנה והתלמוד כמפורסם ודאי יבוא לידי מינות (כדגרסינן במדרש הגדול פרשת שמיני יין ושכר אל תשת. רבי חייא פתח פקודי ה' ישרים משמחי לב, וכתיב יתהלך במישרים. משמחי לב, וכתיב ויין ישמח לבב אנוש. מה היין הזה משמח את הלב. כך דברי תורה העוסק בהן ומחשב בהן יותר מדאי סוף שהוא יוצא למינות ע"כ וכבר הבאתיו בהקדמה). ועל זה רמז הריב"ש במלות קצרות וכי יש אלהות לספירות? ולכך אמר שהמפרש את הקבלה וכו' יהיו לו לבדו, ואיני תוקע את עצמי וכו' וכן עשה רבינו תם ן' יחיא ומהרש"ל המובאים דבריהם בספר רביד הזהב ור"י אלבו בספר העיקרים.

| id | word | punctuation after |
|---|---|---|
| 8:10.1 | קו | ) |
| 8:10.2 | ושרידי |  |
| 8:10.3 | משכילי |  |
| 8:10.4 | עמינו |  |
| 8:10.5 | אשר |  |
| 8:10.6 | הרגישו |  |
| 8:10.7 | בזרות |  |
| 8:10.8 | הקבלה |  |
| 8:10.9 | החדשה |  |
| 8:10.10 | כרבי | ' |
| 8:10.11 | תם |  |
| 8:10.12 | ן | ' |
| 8:10.13 | יחיא | , |
| 8:10.14 | ורבי |  |
| 8:10.15 | פרץ | , |
| 8:10.16 | ורבי |  |
| 8:10.17 | שמשון |  |
| 8:10.18 | מקינון | , |
| 8:10.19 | והרישב"ש |  |
| 8:10.20 | והרש"ל |  |
| 8:10.21 | הרחיקו |  |
| 8:10.22 | את |  |
| 8:10.23 | עצמם |  |
| 8:10.24 | מללמדה |  |
| 8:10.25 | כמו |  |
| 8:10.26 | שאמר |  |
| 8:10.27 | הריב"ש |  |
| 8:10.28 | איני |  |
| 8:10.29 | תוקע |  |
| 8:10.30 | את |  |
| 8:10.31 | עצמי |  |
| 8:10.32 | וכו | ', |
| 8:10.33 | כי |  |
| 8:10.34 | הרגיש |  |
| 8:10.35 | בזרותה |  |
| 8:10.36 | מדרך |  |
| 8:10.37 | תורתינו |  |
| 8:10.38 | הקדושה |  |
| 8:10.39 | והטהורה |  |
| 8:10.40 | מכל |  |
| 8:10.41 | סיג |  |
| 8:10.42 | וחלאה | , |
| 8:10.43 | אך |  |
| 8:10.44 | לא |  |
| 8:10.45 | דבר |  |
| 8:10.46 | כנגדה |  |
| 8:10.47 | רק |  |
| 8:10.48 | כלאחר |  |
| 8:10.49 | יד |  |
| 8:10.50 | באמרו |  |
| 8:10.51 | וכי |  |
| 8:10.52 | יש |  |
| 8:10.53 | אלהות |  |
| 8:10.54 | לספירות | ? |
| 8:10.55 | והמשל |  |
| 8:10.56 | שאמר |  |
| 8:10.57 | לו |  |
| 8:10.58 | דון |  |
| 8:10.59 | יוסף |  |
| 8:10.60 | ן | ' |
| 8:10.61 | שושן |  |
| 8:10.62 | לא |  |
| 8:10.63 | ערב |  |
| 8:10.64 | לו | , |
| 8:10.65 | שידע |  |
| 8:10.66 | באמת |  |
| 8:10.67 | שהוא |  |
| 8:10.68 | כיחש |  |
| 8:10.69 | לו | , |
| 8:10.70 | אך |  |
| 8:10.71 | לא |  |
| 8:10.72 | חפץ |  |
| 8:10.73 | להעציבו |  |
| 8:10.74 | לומר |  |
| 8:10.75 | שקר |  |
| 8:10.76 | דברת | . |
| 8:10.77 | ולפי | ' |
| 8:10.78 | אמר |  |
| 8:10.79 | לו |  |
| 8:10.80 | והנה |  |
| 8:10.81 | טוב |  |
| 8:10.82 | מאד |  |
| 8:10.83 | אם |  |
| 8:10.84 | הוא |  |
| 8:10.85 | כדבריך | . |
| 8:10.86 | כי |  |
| 8:10.87 | ידע |  |
| 8:10.88 | כי |  |
| 8:10.89 | רק |  |
| 8:10.90 | דחיה |  |
| 8:10.91 | בעלמא |  |
| 8:10.92 | קא |  |
| 8:10.93 | מדחי |  |
| 8:10.94 | ליה | , |
| 8:10.95 | כי |  |
| 8:10.96 | גם |  |
| 8:10.97 | אמרם |  |
| 8:10.98 | כי |  |
| 8:10.99 | לנשמתו |  |
| 8:10.100 | אתה |  |
| 8:10.101 | עובד |  |
| 8:10.102 | דחייה |  |
| 8:10.103 | בעלמא |  |
| 8:10.104 | היא | . |
| 8:10.105 | והמה |  |
| 8:10.106 | באמת |  |
| 8:10.107 | עובדים |  |
| 8:10.108 | להם |  |
| 8:10.109 | ומקבלים |  |
| 8:10.110 | אלהותם |  |
| 8:10.111 | עליהם | . |
| 8:10.112 | ומיחדים |  |
| 8:10.113 | את |  |
| 8:10.114 | הכלים |  |
| 8:10.115 | עם |  |
| 8:10.116 | נשמותיהם |  |
| 8:10.117 | כמו |  |
| 8:10.118 | שכתב |  |
| 8:10.119 | שושן |  |
| 8:10.120 | סודות |  |
| 8:10.121 | והרש"ב |  |
| 8:10.122 | בפרשת |  |
| 8:10.123 | בהר |  |
| 8:10.124 | דף |  |
| 8:10.125 | ק"ט |  |
| 8:10.126 | ומשנת |  |
| 8:10.127 | חסידים |  |
| 8:10.128 | סוף |  |
| 8:10.129 | מסכת |  |
| 8:10.130 | העשייה |  |
| 8:10.131 | ועז |  |
| 8:10.132 | לאלהים |  |
| 8:10.133 | שהגוף |  |
| 8:10.134 | והנשמה |  |
| 8:10.135 | חד |  |
| 8:10.136 | אינון | . |
| 8:10.137 | וכ"כ |  |
| 8:10.138 | היכל |  |
| 8:10.139 | הברכה |  |
| 8:10.140 | פרשת |  |
| 8:10.141 | ויקרא | . |
| 8:10.142 | ואלו |  |
| 8:10.143 | היתה |  |
| 8:10.144 | מניעת |  |
| 8:10.145 | הריב"ש |  |
| 8:10.146 | מטעם |  |
| 8:10.147 | שלא |  |
| 8:10.148 | קיבלה |  |
| 8:10.149 | מחכם |  |
| 8:10.150 | מקובל |  |
| 8:10.151 | היה |  |
| 8:10.152 | לו |  |
| 8:10.153 | ללמדה |  |
| 8:10.154 | ולקבלה |  |
| 8:10.155 | מפי |  |
| 8:10.156 | דון |  |
| 8:10.157 | יוסף |  |
| 8:10.158 | ן | ' |
| 8:10.159 | שושן | ! |
| 8:10.160 | אלא |  |
| 8:10.161 | ודאי |  |
| 8:10.162 | דברים |  |
| 8:10.163 | בגו | ! |
| 8:10.164 | שראה |  |
| 8:10.165 | בחכמתו |  |
| 8:10.166 | שזו |  |
| 8:10.167 | היא |  |
| 8:10.168 | יציאה |  |
| 8:10.169 | מדת |  |
| 8:10.170 | תורתינו |  |
| 8:10.171 | הקדושה |  |
| 8:10.172 | לעבוד |  |
| 8:10.173 | אלהים |  |
| 8:10.174 | אחרים | , |
| 8:10.175 | מעשה |  |
| 8:10.176 | ידי |  |
| 8:10.177 | מחשבות |  |
| 8:10.178 | אדם |  |
| 8:10.179 | ורעיונותיו |  |
| 8:10.180 | ודמיונותיו |  |
| 8:10.181 | אשר |  |
| 8:10.182 | עליהם |  |
| 8:10.183 | באה |  |
| 8:10.184 | האזהרה |  |
| 8:10.185 | בתורה |  |
| 8:10.186 | ולא |  |
| 8:10.187 | תתורו |  |
| 8:10.188 | אחרי |  |
| 8:10.189 | לבבכם |  |
| 8:10.190 | ואחרי |  |
| 8:10.191 | עיניכם |  |
| 8:10.192 | ובאה |  |
| 8:10.193 | הקבלה |  |
| 8:10.194 | אחרי |  |
| 8:10.195 | לבבכם |  |
| 8:10.196 | זו |  |
| 8:10.197 | מינות | , |
| 8:10.198 | ואחרי |  |
| 8:10.199 | עיניכם |  |
| 8:10.200 | זו |  |
| 8:10.201 | זנות | , |
| 8:10.202 | כי |  |
| 8:10.203 | לבא |  |
| 8:10.204 | ועינא |  |
| 8:10.205 | תרי |  |
| 8:10.206 | סרסורי |  |
| 8:10.207 | דחטאה |  |
| 8:10.208 | נינהו | , |
| 8:10.209 | ומן |  |
| 8:10.210 | הידוע |  |
| 8:10.211 | שההולך |  |
| 8:10.212 | אחר |  |
| 8:10.213 | דעתו |  |
| 8:10.214 | להמציא |  |
| 8:10.215 | לו |  |
| 8:10.216 | רעיונות |  |
| 8:10.217 | בתורה | . |
| 8:10.218 | שלא |  |
| 8:10.219 | על |  |
| 8:10.220 | דרך |  |
| 8:10.221 | רז"ל |  |
| 8:10.222 | שקבלו |  |
| 8:10.223 | תורה |  |
| 8:10.224 | שבעל |  |
| 8:10.225 | פה |  |
| 8:10.226 | שהיא |  |
| 8:10.227 | המשנה |  |
| 8:10.228 | והתלמוד |  |
| 8:10.229 | כמפורסם |  |
| 8:10.230 | ודאי |  |
| 8:10.231 | יבוא |  |
| 8:10.232 | לידי |  |
| 8:10.233 | מינות |  |
| 8:10.234 | כדגרסינן |  |
| 8:10.235 | במדרש |  |
| 8:10.236 | הגדול |  |
| 8:10.237 | פרשת |  |
| 8:10.238 | שמיני |  |
| 8:10.239 | יין |  |
| 8:10.240 | ושכר |  |
| 8:10.241 | אל |  |
| 8:10.242 | תשת | . |
| 8:10.243 | רבי |  |
| 8:10.244 | חייא |  |
| 8:10.245 | פתח |  |
| 8:10.246 | פקודי |  |
| 8:10.247 | ה | ' |
| 8:10.248 | ישרים |  |
| 8:10.249 | משמחי |  |
| 8:10.250 | לב | , |
| 8:10.251 | וכתיב |  |
| 8:10.252 | יתהלך |  |
| 8:10.253 | במישרים | . |
| 8:10.254 | משמחי |  |
| 8:10.255 | לב | , |
| 8:10.256 | וכתיב |  |
| 8:10.257 | ויין |  |
| 8:10.258 | ישמח |  |
| 8:10.259 | לבב |  |
| 8:10.260 | אנוש | . |
| 8:10.261 | מה |  |
| 8:10.262 | היין |  |
| 8:10.263 | הזה |  |
| 8:10.264 | משמח |  |
| 8:10.265 | את |  |
| 8:10.266 | הלב | . |
| 8:10.267 | כך |  |
| 8:10.268 | דברי |  |
| 8:10.269 | תורה |  |
| 8:10.270 | העוסק |  |
| 8:10.271 | בהן |  |
| 8:10.272 | ומחשב |  |
| 8:10.273 | בהן |  |
| 8:10.274 | יותר |  |
| 8:10.275 | מדאי |  |
| 8:10.276 | סוף |  |
| 8:10.277 | שהוא |  |
| 8:10.278 | יוצא |  |
| 8:10.279 | למינות |  |
| 8:10.280 | ע"כ |  |
| 8:10.281 | וכבר |  |
| 8:10.282 | הבאתיו |  |
| 8:10.283 | בהקדמה | ). |
| 8:10.284 | ועל |  |
| 8:10.285 | זה |  |
| 8:10.286 | רמז |  |
| 8:10.287 | הריב"ש |  |
| 8:10.288 | במלות |  |
| 8:10.289 | קצרות |  |
| 8:10.290 | וכי |  |
| 8:10.291 | יש |  |
| 8:10.292 | אלהות |  |
| 8:10.293 | לספירות | ? |
| 8:10.294 | ולכך |  |
| 8:10.295 | אמר |  |
| 8:10.296 | שהמפרש |  |
| 8:10.297 | את |  |
| 8:10.298 | הקבלה |  |
| 8:10.299 | וכו | ' |
| 8:10.300 | יהיו |  |
| 8:10.301 | לו |  |
| 8:10.302 | לבדו | , |
| 8:10.303 | ואיני |  |
| 8:10.304 | תוקע |  |
| 8:10.305 | את |  |
| 8:10.306 | עצמי |  |
| 8:10.307 | וכו | ' |
| 8:10.308 | וכן |  |
| 8:10.309 | עשה |  |
| 8:10.310 | רבינו |  |
| 8:10.311 | תם |  |
| 8:10.312 | ן | ' |
| 8:10.313 | יחיא |  |
| 8:10.314 | ומהרש"ל |  |
| 8:10.315 | המובאים |  |
| 8:10.316 | דבריהם |  |
| 8:10.317 | בספר |  |
| 8:10.318 | רביד |  |
| 8:10.319 | הזהב |  |
| 8:10.320 | ור"י |  |
| 8:10.321 | אלבו |  |
| 8:10.322 | בספר |  |
| 8:10.323 | העיקרים | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

