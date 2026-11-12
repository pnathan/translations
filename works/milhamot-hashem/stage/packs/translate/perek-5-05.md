# Translation pack: milhamot-hashem perek-5-05 (Milchamot ha-Shem 5:8–5:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-05`.

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
- **קצר אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **שיתוף / שתוף / לשתף / משותף**: associating another with God → fixed,
  "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever associates
  the name of Heaven with something else"). The Maimonidean term for a word said of God and creatures →
  "equivocation / equivocal" ("בשתוף השם בלבד" → "by equivocation of the term
  only"), never "association", Guide I in the note; "shared" only for a thing
  held in common. צירוף: "conjunction" of God; "permutation" in letter
  mysticism.
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **המתפלספים** → "would-be philosophers".
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


## Milchamot ha-Shem 5:8

Hebrew (Wikisource): ובתקוני זהר חדש תקונא קדמאה דף י"ו ע"ג וז"ל הא הכא רזא דתלת שמהן יה"ו משולשין ומתמן בני, היי ומזוני. חיי וההבמה תחיה את בעליה. מזוני פותח את ידיך ומשביע לכל חי רצון. ואיהי אימא עילאה. בני בן י"ה ואיהי יד רמה בינת דאתוון דילה יה"ו, ומן יוד נפיק ידוד ומן ה"א נפיק אלהינו. ומן וא"ו ידוד דאינן י"ד אתוון ורמיזין בקריאת שמע דתמן ידוד אלהינו ידוד ובלה אחד עכ"ל. מי חכם ויבין את זאת ויודיעני מה הפרש וחילוק יש בין אמונת הכתות שהזכיר והזהיר ממנה הר"ר חטר ז"ל ובין דברי הזהר ותיקוני הזהר הללו? ולמה יקראו לעכו"ם מאמיני השלוש. ולא יקראו לאלה מאמיני החמוש. או מאמינ יהי"ב דפרטות כמו שכתב הריב"ש בשם אחד מן המתפלספים?

| id | word | punctuation after |
|---|---|---|
| 5:8.1 | ובתקוני |  |
| 5:8.2 | זהר |  |
| 5:8.3 | חדש |  |
| 5:8.4 | תקונא |  |
| 5:8.5 | קדמאה |  |
| 5:8.6 | דף |  |
| 5:8.7 | י"ו |  |
| 5:8.8 | ע"ג |  |
| 5:8.9 | וז"ל |  |
| 5:8.10 | הא |  |
| 5:8.11 | הכא |  |
| 5:8.12 | רזא |  |
| 5:8.13 | דתלת |  |
| 5:8.14 | שמהן |  |
| 5:8.15 | יה"ו |  |
| 5:8.16 | משולשין |  |
| 5:8.17 | ומתמן |  |
| 5:8.18 | בני | , |
| 5:8.19 | היי |  |
| 5:8.20 | ומזוני | . |
| 5:8.21 | חיי |  |
| 5:8.22 | וההבמה |  |
| 5:8.23 | תחיה |  |
| 5:8.24 | את |  |
| 5:8.25 | בעליה | . |
| 5:8.26 | מזוני |  |
| 5:8.27 | פותח |  |
| 5:8.28 | את |  |
| 5:8.29 | ידיך |  |
| 5:8.30 | ומשביע |  |
| 5:8.31 | לכל |  |
| 5:8.32 | חי |  |
| 5:8.33 | רצון | . |
| 5:8.34 | ואיהי |  |
| 5:8.35 | אימא |  |
| 5:8.36 | עילאה | . |
| 5:8.37 | בני |  |
| 5:8.38 | בן |  |
| 5:8.39 | י"ה |  |
| 5:8.40 | ואיהי |  |
| 5:8.41 | יד |  |
| 5:8.42 | רמה |  |
| 5:8.43 | בינת |  |
| 5:8.44 | דאתוון |  |
| 5:8.45 | דילה |  |
| 5:8.46 | יה"ו | , |
| 5:8.47 | ומן |  |
| 5:8.48 | יוד |  |
| 5:8.49 | נפיק |  |
| 5:8.50 | ידוד |  |
| 5:8.51 | ומן |  |
| 5:8.52 | ה"א |  |
| 5:8.53 | נפיק |  |
| 5:8.54 | אלהינו | . |
| 5:8.55 | ומן |  |
| 5:8.56 | וא"ו |  |
| 5:8.57 | ידוד |  |
| 5:8.58 | דאינן |  |
| 5:8.59 | י"ד |  |
| 5:8.60 | אתוון |  |
| 5:8.61 | ורמיזין |  |
| 5:8.62 | בקריאת |  |
| 5:8.63 | שמע |  |
| 5:8.64 | דתמן |  |
| 5:8.65 | ידוד |  |
| 5:8.66 | אלהינו |  |
| 5:8.67 | ידוד |  |
| 5:8.68 | ובלה |  |
| 5:8.69 | אחד |  |
| 5:8.70 | עכ"ל | . |
| 5:8.71 | מי |  |
| 5:8.72 | חכם |  |
| 5:8.73 | ויבין |  |
| 5:8.74 | את |  |
| 5:8.75 | זאת |  |
| 5:8.76 | ויודיעני |  |
| 5:8.77 | מה |  |
| 5:8.78 | הפרש |  |
| 5:8.79 | וחילוק |  |
| 5:8.80 | יש |  |
| 5:8.81 | בין |  |
| 5:8.82 | אמונת |  |
| 5:8.83 | הכתות |  |
| 5:8.84 | שהזכיר |  |
| 5:8.85 | והזהיר |  |
| 5:8.86 | ממנה |  |
| 5:8.87 | הר"ר |  |
| 5:8.88 | חטר |  |
| 5:8.89 | ז"ל |  |
| 5:8.90 | ובין |  |
| 5:8.91 | דברי |  |
| 5:8.92 | הזהר |  |
| 5:8.93 | ותיקוני |  |
| 5:8.94 | הזהר |  |
| 5:8.95 | הללו | ? |
| 5:8.96 | ולמה |  |
| 5:8.97 | יקראו |  |
| 5:8.98 | לעכו"ם |  |
| 5:8.99 | מאמיני |  |
| 5:8.100 | השלוש | . |
| 5:8.101 | ולא |  |
| 5:8.102 | יקראו |  |
| 5:8.103 | לאלה |  |
| 5:8.104 | מאמיני |  |
| 5:8.105 | החמוש | . |
| 5:8.106 | או |  |
| 5:8.107 | מאמינ |  |
| 5:8.108 | יהי"ב |  |
| 5:8.109 | דפרטות |  |
| 5:8.110 | כמו |  |
| 5:8.111 | שכתב |  |
| 5:8.112 | הריב"ש |  |
| 5:8.113 | בשם |  |
| 5:8.114 | אחד |  |
| 5:8.115 | מן |  |
| 5:8.116 | המתפלספים | ? |


## Milchamot ha-Shem 5:9

Hebrew (Wikisource): מג) תנן במסכת אבות הכל צפוי והרשות נתונה וכו' ופי' הרמב"ם ז"ל כל מה שבעולם ידוע אצלו יתברך והוא משיג אותו (רוצה לומר בין עבר בין עתיד), ולא תאמר כיון שהקב"ה יודע מה שיעשה האדם. וא"כ מוכרח במעשיו שיהיה צדיק או רשע כי הרשות נתונה בידו לעשות טוב או רע, ואין שום דבר שיכריחהו כלל.

| id | word | punctuation after |
|---|---|---|
| 5:9.1 | מג | ) |
| 5:9.2 | תנן |  |
| 5:9.3 | במסכת |  |
| 5:9.4 | אבות |  |
| 5:9.5 | הכל |  |
| 5:9.6 | צפוי |  |
| 5:9.7 | והרשות |  |
| 5:9.8 | נתונה |  |
| 5:9.9 | וכו | ' |
| 5:9.10 | ופי | ' |
| 5:9.11 | הרמב"ם |  |
| 5:9.12 | ז"ל |  |
| 5:9.13 | כל |  |
| 5:9.14 | מה |  |
| 5:9.15 | שבעולם |  |
| 5:9.16 | ידוע |  |
| 5:9.17 | אצלו |  |
| 5:9.18 | יתברך |  |
| 5:9.19 | והוא |  |
| 5:9.20 | משיג |  |
| 5:9.21 | אותו |  |
| 5:9.22 | רוצה |  |
| 5:9.23 | לומר |  |
| 5:9.24 | בין |  |
| 5:9.25 | עבר |  |
| 5:9.26 | בין |  |
| 5:9.27 | עתיד | ), |
| 5:9.28 | ולא |  |
| 5:9.29 | תאמר |  |
| 5:9.30 | כיון |  |
| 5:9.31 | שהקב"ה |  |
| 5:9.32 | יודע |  |
| 5:9.33 | מה |  |
| 5:9.34 | שיעשה |  |
| 5:9.35 | האדם | . |
| 5:9.36 | וא"כ |  |
| 5:9.37 | מוכרח |  |
| 5:9.38 | במעשיו |  |
| 5:9.39 | שיהיה |  |
| 5:9.40 | צדיק |  |
| 5:9.41 | או |  |
| 5:9.42 | רשע |  |
| 5:9.43 | כי |  |
| 5:9.44 | הרשות |  |
| 5:9.45 | נתונה |  |
| 5:9.46 | בידו |  |
| 5:9.47 | לעשות |  |
| 5:9.48 | טוב |  |
| 5:9.49 | או |  |
| 5:9.50 | רע | , |
| 5:9.51 | ואין |  |
| 5:9.52 | שום |  |
| 5:9.53 | דבר |  |
| 5:9.54 | שיכריחהו |  |
| 5:9.55 | כלל | . |


## Milchamot ha-Shem 5:10

Hebrew (Wikisource): ובסוף השמנה פרקים כתב שאין ידיעת השי"ת כידיעתנו ואעפ"י שאנו אומרים בו לשון ידיעה כמו שאנו אומרים ידיעה בנו אין זה אלא בשתוף השם בלבד. וכמו שאים בנו כח לידע מציאות אמתתו ית' כמו שנ' החקר אלוה תמצא אם עד תכלית שדי תמצא. כל אין בנו כח לחקור ולהשיג ולמצוא דעתו, כי הוא ודעתו אחד, לא כדעת האדם שהוא ודעתו שנים, הוא שהנביא אומר כי לא מחשבותי מחשבותיכם. וכ"כ בחבורו פ"ה מהלכות תשובה ע"ש ובמורה משלישי. וכן כתב רבינו סעדיה גאון בספרו הנכבד האמונות והדעות מאמר ד'. ועיין תוספת יו"ט שהאריך בענין זה במשנה הנז"ל. והענין הוא כי אלהינו ב"ה איננו נופל תחת הזמן. כי הזמן הוא מפעולותיו והוא יודע כל חלקיו. עבר והוה ועתיד אין דבר נסתר ממנו כמו שאמרו רז"ל אנשי כנה"ג בתפלת מוסף דראש השנה שהוא יתברך צופה ומביט עד סוף כל הדורות. ר"ל שהכל גלוי לפניו יתברך עבר והוה ועתיד. ולפיכך כינו חז"ל אותה בלשון צפיה והבטה, ולא בלשון ידיעה, ללמדנו שהכל סדור וצפוי לפנוי ית' ומביט בו, ולכן כתב הרמב"ם במלות קצרות שאין בנו כח לידע אמיתת ידיעתו, לפי שידיעת האדם אינה משגת אלא את ההוה ולא מה שעתיד להיות.

| id | word | punctuation after |
|---|---|---|
| 5:10.1 | ובסוף |  |
| 5:10.2 | השמנה |  |
| 5:10.3 | פרקים |  |
| 5:10.4 | כתב |  |
| 5:10.5 | שאין |  |
| 5:10.6 | ידיעת |  |
| 5:10.7 | השי"ת |  |
| 5:10.8 | כידיעתנו |  |
| 5:10.9 | ואעפ"י |  |
| 5:10.10 | שאנו |  |
| 5:10.11 | אומרים |  |
| 5:10.12 | בו |  |
| 5:10.13 | לשון |  |
| 5:10.14 | ידיעה |  |
| 5:10.15 | כמו |  |
| 5:10.16 | שאנו |  |
| 5:10.17 | אומרים |  |
| 5:10.18 | ידיעה |  |
| 5:10.19 | בנו |  |
| 5:10.20 | אין |  |
| 5:10.21 | זה |  |
| 5:10.22 | אלא |  |
| 5:10.23 | בשתוף |  |
| 5:10.24 | השם |  |
| 5:10.25 | בלבד | . |
| 5:10.26 | וכמו |  |
| 5:10.27 | שאים |  |
| 5:10.28 | בנו |  |
| 5:10.29 | כח |  |
| 5:10.30 | לידע |  |
| 5:10.31 | מציאות |  |
| 5:10.32 | אמתתו |  |
| 5:10.33 | ית | ' |
| 5:10.34 | כמו |  |
| 5:10.35 | שנ | ' |
| 5:10.36 | החקר |  |
| 5:10.37 | אלוה |  |
| 5:10.38 | תמצא |  |
| 5:10.39 | אם |  |
| 5:10.40 | עד |  |
| 5:10.41 | תכלית |  |
| 5:10.42 | שדי |  |
| 5:10.43 | תמצא | . |
| 5:10.44 | כל |  |
| 5:10.45 | אין |  |
| 5:10.46 | בנו |  |
| 5:10.47 | כח |  |
| 5:10.48 | לחקור |  |
| 5:10.49 | ולהשיג |  |
| 5:10.50 | ולמצוא |  |
| 5:10.51 | דעתו | , |
| 5:10.52 | כי |  |
| 5:10.53 | הוא |  |
| 5:10.54 | ודעתו |  |
| 5:10.55 | אחד | , |
| 5:10.56 | לא |  |
| 5:10.57 | כדעת |  |
| 5:10.58 | האדם |  |
| 5:10.59 | שהוא |  |
| 5:10.60 | ודעתו |  |
| 5:10.61 | שנים | , |
| 5:10.62 | הוא |  |
| 5:10.63 | שהנביא |  |
| 5:10.64 | אומר |  |
| 5:10.65 | כי |  |
| 5:10.66 | לא |  |
| 5:10.67 | מחשבותי |  |
| 5:10.68 | מחשבותיכם | . |
| 5:10.69 | וכ"כ |  |
| 5:10.70 | בחבורו |  |
| 5:10.71 | פ"ה |  |
| 5:10.72 | מהלכות |  |
| 5:10.73 | תשובה |  |
| 5:10.74 | ע"ש |  |
| 5:10.75 | ובמורה |  |
| 5:10.76 | משלישי | . |
| 5:10.77 | וכן |  |
| 5:10.78 | כתב |  |
| 5:10.79 | רבינו |  |
| 5:10.80 | סעדיה |  |
| 5:10.81 | גאון |  |
| 5:10.82 | בספרו |  |
| 5:10.83 | הנכבד |  |
| 5:10.84 | האמונות |  |
| 5:10.85 | והדעות |  |
| 5:10.86 | מאמר |  |
| 5:10.87 | ד | '. |
| 5:10.88 | ועיין |  |
| 5:10.89 | תוספת |  |
| 5:10.90 | יו"ט |  |
| 5:10.91 | שהאריך |  |
| 5:10.92 | בענין |  |
| 5:10.93 | זה |  |
| 5:10.94 | במשנה |  |
| 5:10.95 | הנז"ל | . |
| 5:10.96 | והענין |  |
| 5:10.97 | הוא |  |
| 5:10.98 | כי |  |
| 5:10.99 | אלהינו |  |
| 5:10.100 | ב"ה |  |
| 5:10.101 | איננו |  |
| 5:10.102 | נופל |  |
| 5:10.103 | תחת |  |
| 5:10.104 | הזמן | . |
| 5:10.105 | כי |  |
| 5:10.106 | הזמן |  |
| 5:10.107 | הוא |  |
| 5:10.108 | מפעולותיו |  |
| 5:10.109 | והוא |  |
| 5:10.110 | יודע |  |
| 5:10.111 | כל |  |
| 5:10.112 | חלקיו | . |
| 5:10.113 | עבר |  |
| 5:10.114 | והוה |  |
| 5:10.115 | ועתיד |  |
| 5:10.116 | אין |  |
| 5:10.117 | דבר |  |
| 5:10.118 | נסתר |  |
| 5:10.119 | ממנו |  |
| 5:10.120 | כמו |  |
| 5:10.121 | שאמרו |  |
| 5:10.122 | רז"ל |  |
| 5:10.123 | אנשי |  |
| 5:10.124 | כנה"ג |  |
| 5:10.125 | בתפלת |  |
| 5:10.126 | מוסף |  |
| 5:10.127 | דראש |  |
| 5:10.128 | השנה |  |
| 5:10.129 | שהוא |  |
| 5:10.130 | יתברך |  |
| 5:10.131 | צופה |  |
| 5:10.132 | ומביט |  |
| 5:10.133 | עד |  |
| 5:10.134 | סוף |  |
| 5:10.135 | כל |  |
| 5:10.136 | הדורות | . |
| 5:10.137 | ר"ל |  |
| 5:10.138 | שהכל |  |
| 5:10.139 | גלוי |  |
| 5:10.140 | לפניו |  |
| 5:10.141 | יתברך |  |
| 5:10.142 | עבר |  |
| 5:10.143 | והוה |  |
| 5:10.144 | ועתיד | . |
| 5:10.145 | ולפיכך |  |
| 5:10.146 | כינו |  |
| 5:10.147 | חז"ל |  |
| 5:10.148 | אותה |  |
| 5:10.149 | בלשון |  |
| 5:10.150 | צפיה |  |
| 5:10.151 | והבטה | , |
| 5:10.152 | ולא |  |
| 5:10.153 | בלשון |  |
| 5:10.154 | ידיעה | , |
| 5:10.155 | ללמדנו |  |
| 5:10.156 | שהכל |  |
| 5:10.157 | סדור |  |
| 5:10.158 | וצפוי |  |
| 5:10.159 | לפנוי |  |
| 5:10.160 | ית | ' |
| 5:10.161 | ומביט |  |
| 5:10.162 | בו | , |
| 5:10.163 | ולכן |  |
| 5:10.164 | כתב |  |
| 5:10.165 | הרמב"ם |  |
| 5:10.166 | במלות |  |
| 5:10.167 | קצרות |  |
| 5:10.168 | שאין |  |
| 5:10.169 | בנו |  |
| 5:10.170 | כח |  |
| 5:10.171 | לידע |  |
| 5:10.172 | אמיתת |  |
| 5:10.173 | ידיעתו | , |
| 5:10.174 | לפי |  |
| 5:10.175 | שידיעת |  |
| 5:10.176 | האדם |  |
| 5:10.177 | אינה |  |
| 5:10.178 | משגת |  |
| 5:10.179 | אלא |  |
| 5:10.180 | את |  |
| 5:10.181 | ההוה |  |
| 5:10.182 | ולא |  |
| 5:10.183 | מה |  |
| 5:10.184 | שעתיד |  |
| 5:10.185 | להיות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

