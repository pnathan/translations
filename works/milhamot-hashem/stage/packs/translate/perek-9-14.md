# Translation pack: milhamot-hashem perek-9-14 (Milchamot ha-Shem 9:17–9:17)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-14.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-14`.

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
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
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
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 9:17

Hebrew (Wikisource): קלה) רב לכם נשיאי ישראל מכל דברי מינות האמורים בזהר בכמה מקומות, ובפרט בפרשת בהר דף ק"כ שהאלהים העליון שאם ראוי לעבדו (לפי דבריה) ברא והמציא את האלהים התחתונים הרבים, והוא ממלא אותם בחללם וחוצה להם, והוא המאחד ומקשר את נקיבותיהם עם הזכרים לקרא להם בשם קב"ה ושכינתיה, ואותם נעבוד, ונירא את נקיבותיהם, ודוקא לפרצוף אחרון שבהם הנקרא קצר אפים (ז"א) עם מלכת השמים המזדווגת עמו, משום דלא שייך עבודה ותפלה רק לפרצופים שכבר נתגשמו ונתעבו קצת כמבואר במה שקדם מדברי כסא אליהו יושר לבב וספר הברית ועץ חיים להרח"ו שער טנת"ה סוף פרק ב' ובשער הנקודים פ"א שכל מה שהאו יורד למטה הוא מתעבה יותר וניכר ונרגש ומתגלה יותר, אבל באין סוף הוא הסבה הראשונה ואפי' בפרצופים ראשונים (שהמציאו המקובלים החדשים לאלהינו) כאדם קדמאה ואדם קדמון ועתיק ואריך שעדיין לא נתעבו ונתגשמו כל כך לא שייכא עבודה וקריאה, כי אין המה, ולא יאות לקרוא להם בשם אלהים כמפורש יוצא מספר עץ חיים להרח"ו שער הנקודים פ"ז ובהג"ה שם. שזהו האלהים העליון עשה את האדם ישר (והמה בקשו חשבונות רבים) כי האדם קדמון היה בו בחינת יושר וזהו עשה את האדם ישר (הגהה א"מ ןקרא המאציל אלהים בדרך שאלה. כי הרי אפי' באריך אנפין אין שם אלהים) ע"כ. הרי לך בהדיא דפרצופים העליונים לא שייכא בהו אלהות וקריאה כלל, כי לא אלהים המה ורק להעלולים התחתונים נקרא ונעבוד, והמה אלהינו (לפי דעתם) כמו שכתב הרב יושר לבב וספר הברית וכס"א ונח"י ומצרף האמונה! ולפי דבריהם שאומרים דלית רשותא לבר נש למיעל קמי מריה אלא בידיעת אמונה זו, איך אמרו רז"ל רצה הקב"ה לזכות את ישראל לפיכך הרבה להם תורה ומצות שנ' ה' חפץ למען צדקו יגדיל תורה ויאדיר! ולמה זה לא רצה הקב"ה לזכות אותנו ולהודענו על ידי נביאנו הרבים שעמדו להם לישראל ממשה רבינו ע"ה עד שהנביא האחרון מלאכי שאמר זכרו תורת משה עבדי וכו'. את אמונה זו שהעבודה היא להעלולים התחתונים. ושהעובד לאלהים עליון לא יענהו ואדרבא ענוש יענש ומחה ה' את שמו כמו שכתב ספר הברית?

| id | word | punctuation after |
|---|---|---|
| 9:17.1 | קלה | ) |
| 9:17.2 | רב |  |
| 9:17.3 | לכם |  |
| 9:17.4 | נשיאי |  |
| 9:17.5 | ישראל |  |
| 9:17.6 | מכל |  |
| 9:17.7 | דברי |  |
| 9:17.8 | מינות |  |
| 9:17.9 | האמורים |  |
| 9:17.10 | בזהר |  |
| 9:17.11 | בכמה |  |
| 9:17.12 | מקומות | , |
| 9:17.13 | ובפרט |  |
| 9:17.14 | בפרשת |  |
| 9:17.15 | בהר |  |
| 9:17.16 | דף |  |
| 9:17.17 | ק"כ |  |
| 9:17.18 | שהאלהים |  |
| 9:17.19 | העליון |  |
| 9:17.20 | שאם |  |
| 9:17.21 | ראוי |  |
| 9:17.22 | לעבדו |  |
| 9:17.23 | לפי |  |
| 9:17.24 | דבריה | ) |
| 9:17.25 | ברא |  |
| 9:17.26 | והמציא |  |
| 9:17.27 | את |  |
| 9:17.28 | האלהים |  |
| 9:17.29 | התחתונים |  |
| 9:17.30 | הרבים | , |
| 9:17.31 | והוא |  |
| 9:17.32 | ממלא |  |
| 9:17.33 | אותם |  |
| 9:17.34 | בחללם |  |
| 9:17.35 | וחוצה |  |
| 9:17.36 | להם | , |
| 9:17.37 | והוא |  |
| 9:17.38 | המאחד |  |
| 9:17.39 | ומקשר |  |
| 9:17.40 | את |  |
| 9:17.41 | נקיבותיהם |  |
| 9:17.42 | עם |  |
| 9:17.43 | הזכרים |  |
| 9:17.44 | לקרא |  |
| 9:17.45 | להם |  |
| 9:17.46 | בשם |  |
| 9:17.47 | קב"ה |  |
| 9:17.48 | ושכינתיה | , |
| 9:17.49 | ואותם |  |
| 9:17.50 | נעבוד | , |
| 9:17.51 | ונירא |  |
| 9:17.52 | את |  |
| 9:17.53 | נקיבותיהם | , |
| 9:17.54 | ודוקא |  |
| 9:17.55 | לפרצוף |  |
| 9:17.56 | אחרון |  |
| 9:17.57 | שבהם |  |
| 9:17.58 | הנקרא |  |
| 9:17.59 | קצר |  |
| 9:17.60 | אפים |  |
| 9:17.61 | ז"א | ) |
| 9:17.62 | עם |  |
| 9:17.63 | מלכת |  |
| 9:17.64 | השמים |  |
| 9:17.65 | המזדווגת |  |
| 9:17.66 | עמו | , |
| 9:17.67 | משום |  |
| 9:17.68 | דלא |  |
| 9:17.69 | שייך |  |
| 9:17.70 | עבודה |  |
| 9:17.71 | ותפלה |  |
| 9:17.72 | רק |  |
| 9:17.73 | לפרצופים |  |
| 9:17.74 | שכבר |  |
| 9:17.75 | נתגשמו |  |
| 9:17.76 | ונתעבו |  |
| 9:17.77 | קצת |  |
| 9:17.78 | כמבואר |  |
| 9:17.79 | במה |  |
| 9:17.80 | שקדם |  |
| 9:17.81 | מדברי |  |
| 9:17.82 | כסא |  |
| 9:17.83 | אליהו |  |
| 9:17.84 | יושר |  |
| 9:17.85 | לבב |  |
| 9:17.86 | וספר |  |
| 9:17.87 | הברית |  |
| 9:17.88 | ועץ |  |
| 9:17.89 | חיים |  |
| 9:17.90 | להרח"ו |  |
| 9:17.91 | שער |  |
| 9:17.92 | טנת"ה |  |
| 9:17.93 | סוף |  |
| 9:17.94 | פרק |  |
| 9:17.95 | ב | ' |
| 9:17.96 | ובשער |  |
| 9:17.97 | הנקודים |  |
| 9:17.98 | פ"א |  |
| 9:17.99 | שכל |  |
| 9:17.100 | מה |  |
| 9:17.101 | שהאו |  |
| 9:17.102 | יורד |  |
| 9:17.103 | למטה |  |
| 9:17.104 | הוא |  |
| 9:17.105 | מתעבה |  |
| 9:17.106 | יותר |  |
| 9:17.107 | וניכר |  |
| 9:17.108 | ונרגש |  |
| 9:17.109 | ומתגלה |  |
| 9:17.110 | יותר | , |
| 9:17.111 | אבל |  |
| 9:17.112 | באין |  |
| 9:17.113 | סוף |  |
| 9:17.114 | הוא |  |
| 9:17.115 | הסבה |  |
| 9:17.116 | הראשונה |  |
| 9:17.117 | ואפי | ' |
| 9:17.118 | בפרצופים |  |
| 9:17.119 | ראשונים |  |
| 9:17.120 | שהמציאו |  |
| 9:17.121 | המקובלים |  |
| 9:17.122 | החדשים |  |
| 9:17.123 | לאלהינו | ) |
| 9:17.124 | כאדם |  |
| 9:17.125 | קדמאה |  |
| 9:17.126 | ואדם |  |
| 9:17.127 | קדמון |  |
| 9:17.128 | ועתיק |  |
| 9:17.129 | ואריך |  |
| 9:17.130 | שעדיין |  |
| 9:17.131 | לא |  |
| 9:17.132 | נתעבו |  |
| 9:17.133 | ונתגשמו |  |
| 9:17.134 | כל |  |
| 9:17.135 | כך |  |
| 9:17.136 | לא |  |
| 9:17.137 | שייכא |  |
| 9:17.138 | עבודה |  |
| 9:17.139 | וקריאה | , |
| 9:17.140 | כי |  |
| 9:17.141 | אין |  |
| 9:17.142 | המה | , |
| 9:17.143 | ולא |  |
| 9:17.144 | יאות |  |
| 9:17.145 | לקרוא |  |
| 9:17.146 | להם |  |
| 9:17.147 | בשם |  |
| 9:17.148 | אלהים |  |
| 9:17.149 | כמפורש |  |
| 9:17.150 | יוצא |  |
| 9:17.151 | מספר |  |
| 9:17.152 | עץ |  |
| 9:17.153 | חיים |  |
| 9:17.154 | להרח"ו |  |
| 9:17.155 | שער |  |
| 9:17.156 | הנקודים |  |
| 9:17.157 | פ"ז |  |
| 9:17.158 | ובהג"ה |  |
| 9:17.159 | שם | . |
| 9:17.160 | שזהו |  |
| 9:17.161 | האלהים |  |
| 9:17.162 | העליון |  |
| 9:17.163 | עשה |  |
| 9:17.164 | את |  |
| 9:17.165 | האדם |  |
| 9:17.166 | ישר |  |
| 9:17.167 | והמה |  |
| 9:17.168 | בקשו |  |
| 9:17.169 | חשבונות |  |
| 9:17.170 | רבים | ) |
| 9:17.171 | כי |  |
| 9:17.172 | האדם |  |
| 9:17.173 | קדמון |  |
| 9:17.174 | היה |  |
| 9:17.175 | בו |  |
| 9:17.176 | בחינת |  |
| 9:17.177 | יושר |  |
| 9:17.178 | וזהו |  |
| 9:17.179 | עשה |  |
| 9:17.180 | את |  |
| 9:17.181 | האדם |  |
| 9:17.182 | ישר |  |
| 9:17.183 | הגהה |  |
| 9:17.184 | א"מ |  |
| 9:17.185 | ןקרא |  |
| 9:17.186 | המאציל |  |
| 9:17.187 | אלהים |  |
| 9:17.188 | בדרך |  |
| 9:17.189 | שאלה | . |
| 9:17.190 | כי |  |
| 9:17.191 | הרי |  |
| 9:17.192 | אפי | ' |
| 9:17.193 | באריך |  |
| 9:17.194 | אנפין |  |
| 9:17.195 | אין |  |
| 9:17.196 | שם |  |
| 9:17.197 | אלהים | ) |
| 9:17.198 | ע"כ | . |
| 9:17.199 | הרי |  |
| 9:17.200 | לך |  |
| 9:17.201 | בהדיא |  |
| 9:17.202 | דפרצופים |  |
| 9:17.203 | העליונים |  |
| 9:17.204 | לא |  |
| 9:17.205 | שייכא |  |
| 9:17.206 | בהו |  |
| 9:17.207 | אלהות |  |
| 9:17.208 | וקריאה |  |
| 9:17.209 | כלל | , |
| 9:17.210 | כי |  |
| 9:17.211 | לא |  |
| 9:17.212 | אלהים |  |
| 9:17.213 | המה |  |
| 9:17.214 | ורק |  |
| 9:17.215 | להעלולים |  |
| 9:17.216 | התחתונים |  |
| 9:17.217 | נקרא |  |
| 9:17.218 | ונעבוד | , |
| 9:17.219 | והמה |  |
| 9:17.220 | אלהינו |  |
| 9:17.221 | לפי |  |
| 9:17.222 | דעתם | ) |
| 9:17.223 | כמו |  |
| 9:17.224 | שכתב |  |
| 9:17.225 | הרב |  |
| 9:17.226 | יושר |  |
| 9:17.227 | לבב |  |
| 9:17.228 | וספר |  |
| 9:17.229 | הברית |  |
| 9:17.230 | וכס"א |  |
| 9:17.231 | ונח"י |  |
| 9:17.232 | ומצרף |  |
| 9:17.233 | האמונה | ! |
| 9:17.234 | ולפי |  |
| 9:17.235 | דבריהם |  |
| 9:17.236 | שאומרים |  |
| 9:17.237 | דלית |  |
| 9:17.238 | רשותא |  |
| 9:17.239 | לבר |  |
| 9:17.240 | נש |  |
| 9:17.241 | למיעל |  |
| 9:17.242 | קמי |  |
| 9:17.243 | מריה |  |
| 9:17.244 | אלא |  |
| 9:17.245 | בידיעת |  |
| 9:17.246 | אמונה |  |
| 9:17.247 | זו | , |
| 9:17.248 | איך |  |
| 9:17.249 | אמרו |  |
| 9:17.250 | רז"ל |  |
| 9:17.251 | רצה |  |
| 9:17.252 | הקב"ה |  |
| 9:17.253 | לזכות |  |
| 9:17.254 | את |  |
| 9:17.255 | ישראל |  |
| 9:17.256 | לפיכך |  |
| 9:17.257 | הרבה |  |
| 9:17.258 | להם |  |
| 9:17.259 | תורה |  |
| 9:17.260 | ומצות |  |
| 9:17.261 | שנ | ' |
| 9:17.262 | ה | ' |
| 9:17.263 | חפץ |  |
| 9:17.264 | למען |  |
| 9:17.265 | צדקו |  |
| 9:17.266 | יגדיל |  |
| 9:17.267 | תורה |  |
| 9:17.268 | ויאדיר | ! |
| 9:17.269 | ולמה |  |
| 9:17.270 | זה |  |
| 9:17.271 | לא |  |
| 9:17.272 | רצה |  |
| 9:17.273 | הקב"ה |  |
| 9:17.274 | לזכות |  |
| 9:17.275 | אותנו |  |
| 9:17.276 | ולהודענו |  |
| 9:17.277 | על |  |
| 9:17.278 | ידי |  |
| 9:17.279 | נביאנו |  |
| 9:17.280 | הרבים |  |
| 9:17.281 | שעמדו |  |
| 9:17.282 | להם |  |
| 9:17.283 | לישראל |  |
| 9:17.284 | ממשה |  |
| 9:17.285 | רבינו |  |
| 9:17.286 | ע"ה |  |
| 9:17.287 | עד |  |
| 9:17.288 | שהנביא |  |
| 9:17.289 | האחרון |  |
| 9:17.290 | מלאכי |  |
| 9:17.291 | שאמר |  |
| 9:17.292 | זכרו |  |
| 9:17.293 | תורת |  |
| 9:17.294 | משה |  |
| 9:17.295 | עבדי |  |
| 9:17.296 | וכו | '. |
| 9:17.297 | את |  |
| 9:17.298 | אמונה |  |
| 9:17.299 | זו |  |
| 9:17.300 | שהעבודה |  |
| 9:17.301 | היא |  |
| 9:17.302 | להעלולים |  |
| 9:17.303 | התחתונים | . |
| 9:17.304 | ושהעובד |  |
| 9:17.305 | לאלהים |  |
| 9:17.306 | עליון |  |
| 9:17.307 | לא |  |
| 9:17.308 | יענהו |  |
| 9:17.309 | ואדרבא |  |
| 9:17.310 | ענוש |  |
| 9:17.311 | יענש |  |
| 9:17.312 | ומחה |  |
| 9:17.313 | ה | ' |
| 9:17.314 | את |  |
| 9:17.315 | שמו |  |
| 9:17.316 | כמו |  |
| 9:17.317 | שכתב |  |
| 9:17.318 | ספר |  |
| 9:17.319 | הברית | ? |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

