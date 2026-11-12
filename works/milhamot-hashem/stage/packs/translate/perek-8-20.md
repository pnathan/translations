# Translation pack: milhamot-hashem perek-8-20 (Milchamot ha-Shem 8:24–8:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-20.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-20`.

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
- **הוי"ה / הוי'ה / הויה** (the Name as a name) → "YHVH", never "being";
  "מילוי הויה ע"ב ס"ג" → "the fillings of YHVH, 72, 63" (note: AB, SaG). אדנות as a name → "Adonai" (note: "lordship").
  עשר הויות → "ten havayot" (note).
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
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
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


## Milchamot ha-Shem 8:24

Hebrew (Wikisource): קיז) וכלל הדבר כי הפילסוף המסית מחבר הזהר שם רעיונותיו ולבו בשאלת מה למעלה ומה למטה וכו' והפקיר איסור זה הנזכר במשנה, ולא הניח אלא מה למעלה מכיפת עיגול הא"ס אשר לדעתו נצתמצם למעלה סביב, ואע"פ שתמצא אותו אומר דאית אתר דלא קימא לשאלה ביה, אין זה אלא לגנוב את העין כאלו הוא מאמין באיסור זה, והראיה שהרי דבר במהות האלוה וצמצומו שנעשה ככדור עגול חלול, וכיון שדבר באין סוף שהוא למעלה מן הכל, איך יהיה אסור לשאול ולדבר במה שלמטה ממנו תוך חללו? וכבר גילה דעתו בהדיא, שכל אסור ופסול וטמא הנזכרים במשנה הם מסטרא דיצר הרע, ולכך נהג בזה הפקר ודבר במהות יוצר הכל ואיך ציורו ואופן חנייתו, ועשה לו גופים ופרצופים קבועים שונים זה מזה בכמותם ואיכותם ומקומם, זה למעלה וזה למטה, וזה ארוך וזה קצר כמו שיראה המעיין בו ובעץ חיים ופריו לכהרח"ו, ושבש ובלבל אמונת ישראל הטהורה המקובלת בידם, והשוה אותה לאמונת העכו"ם. ורבים מעמנו נמשכו אחריו לתומם בהאמנים שמהתנא הקדוש רשב"י יצאו הדברים הכבושים אשר היו קבורים, ולא שמו אל לבם איך אפשר לקבל תורה חדשה מיד עכו"ם, ואיך נקברה בארץ בארגז, ומי קברה? ולמה נתגלית בימי רשב"י וחזרה ונקברה ונשכח זכרה מכל היהודים? אשר על כן חלילה לכל משכיל בתורת ה' שבכתב ושבעל פה להאמין שדברים כאלה יצאו מפי רשב"י ע"ה ולא משום תנא או אמורא, כי אדרבא מצינו ראינו שהיו מרחיקים אלו הדעות הרעות ומתנגדים להם, וכמו שכבר הבאתי קצת מאמרים שמהם נלמד התנגדותם לאלו הדעות החדשות הנזכרים בקבלה החדשה, דממה נפשך הואיל ויש לפרצופים דין אלהות כמו שכתב הרש"ב בגליון פרשת בהר דף ק"ט והרב משנת חסידים סוף מסכת העשייה, ובית קדש הקדשים דף נ"ו ע"ב! נמצאנו מאמינים ועובדים לאלוהות רבים נבראים שכל אחד מהם הוא סבה למה שאחריו כמו שכתב לעיל בשם זוהר בראשית דף כ"ב שהוא קורא לכל אחד מהם בשם הוי'ה ואדנות ואלהים וקב"ה ושאר שמות המיוחדים, עד נדר שכשישמע האיש הישראלי את חבירו מזכיר את השם הנכבד הויה או אלהים או הקב"ה, צריך השומע לשאול איזה הוא הקב"ה שאתה מזכיר, אם עתיק, או אריך, או אבא, או אימא, או זעיר, או נוקבא דיליה, כמו שעשה הוא עצמו מחבר הזהר בפסוק ויאמר אלהים נעשה אדם, ואמר מאן אלהים, כדלעיל סי' כ'! כי האלהים העליון לא תוכל לומר (לפי דבריהם) כי הוא אין לו שם ולא נקודה, ולא תוכל לכנותו בשם מן השמות הקדושים כמו שכתבו המקובלים החדשים, ונמצאנו מאמינים באלהים רבים נבראים ח"ו.

| id | word | punctuation after |
|---|---|---|
| 8:24.1 | קיז | ) |
| 8:24.2 | וכלל |  |
| 8:24.3 | הדבר |  |
| 8:24.4 | כי |  |
| 8:24.5 | הפילסוף |  |
| 8:24.6 | המסית |  |
| 8:24.7 | מחבר |  |
| 8:24.8 | הזהר |  |
| 8:24.9 | שם |  |
| 8:24.10 | רעיונותיו |  |
| 8:24.11 | ולבו |  |
| 8:24.12 | בשאלת |  |
| 8:24.13 | מה |  |
| 8:24.14 | למעלה |  |
| 8:24.15 | ומה |  |
| 8:24.16 | למטה |  |
| 8:24.17 | וכו | ' |
| 8:24.18 | והפקיר |  |
| 8:24.19 | איסור |  |
| 8:24.20 | זה |  |
| 8:24.21 | הנזכר |  |
| 8:24.22 | במשנה | , |
| 8:24.23 | ולא |  |
| 8:24.24 | הניח |  |
| 8:24.25 | אלא |  |
| 8:24.26 | מה |  |
| 8:24.27 | למעלה |  |
| 8:24.28 | מכיפת |  |
| 8:24.29 | עיגול |  |
| 8:24.30 | הא"ס |  |
| 8:24.31 | אשר |  |
| 8:24.32 | לדעתו |  |
| 8:24.33 | נצתמצם |  |
| 8:24.34 | למעלה |  |
| 8:24.35 | סביב | , |
| 8:24.36 | ואע"פ |  |
| 8:24.37 | שתמצא |  |
| 8:24.38 | אותו |  |
| 8:24.39 | אומר |  |
| 8:24.40 | דאית |  |
| 8:24.41 | אתר |  |
| 8:24.42 | דלא |  |
| 8:24.43 | קימא |  |
| 8:24.44 | לשאלה |  |
| 8:24.45 | ביה | , |
| 8:24.46 | אין |  |
| 8:24.47 | זה |  |
| 8:24.48 | אלא |  |
| 8:24.49 | לגנוב |  |
| 8:24.50 | את |  |
| 8:24.51 | העין |  |
| 8:24.52 | כאלו |  |
| 8:24.53 | הוא |  |
| 8:24.54 | מאמין |  |
| 8:24.55 | באיסור |  |
| 8:24.56 | זה | , |
| 8:24.57 | והראיה |  |
| 8:24.58 | שהרי |  |
| 8:24.59 | דבר |  |
| 8:24.60 | במהות |  |
| 8:24.61 | האלוה |  |
| 8:24.62 | וצמצומו |  |
| 8:24.63 | שנעשה |  |
| 8:24.64 | ככדור |  |
| 8:24.65 | עגול |  |
| 8:24.66 | חלול | , |
| 8:24.67 | וכיון |  |
| 8:24.68 | שדבר |  |
| 8:24.69 | באין |  |
| 8:24.70 | סוף |  |
| 8:24.71 | שהוא |  |
| 8:24.72 | למעלה |  |
| 8:24.73 | מן |  |
| 8:24.74 | הכל | , |
| 8:24.75 | איך |  |
| 8:24.76 | יהיה |  |
| 8:24.77 | אסור |  |
| 8:24.78 | לשאול |  |
| 8:24.79 | ולדבר |  |
| 8:24.80 | במה |  |
| 8:24.81 | שלמטה |  |
| 8:24.82 | ממנו |  |
| 8:24.83 | תוך |  |
| 8:24.84 | חללו | ? |
| 8:24.85 | וכבר |  |
| 8:24.86 | גילה |  |
| 8:24.87 | דעתו |  |
| 8:24.88 | בהדיא | , |
| 8:24.89 | שכל |  |
| 8:24.90 | אסור |  |
| 8:24.91 | ופסול |  |
| 8:24.92 | וטמא |  |
| 8:24.93 | הנזכרים |  |
| 8:24.94 | במשנה |  |
| 8:24.95 | הם |  |
| 8:24.96 | מסטרא |  |
| 8:24.97 | דיצר |  |
| 8:24.98 | הרע | , |
| 8:24.99 | ולכך |  |
| 8:24.100 | נהג |  |
| 8:24.101 | בזה |  |
| 8:24.102 | הפקר |  |
| 8:24.103 | ודבר |  |
| 8:24.104 | במהות |  |
| 8:24.105 | יוצר |  |
| 8:24.106 | הכל |  |
| 8:24.107 | ואיך |  |
| 8:24.108 | ציורו |  |
| 8:24.109 | ואופן |  |
| 8:24.110 | חנייתו | , |
| 8:24.111 | ועשה |  |
| 8:24.112 | לו |  |
| 8:24.113 | גופים |  |
| 8:24.114 | ופרצופים |  |
| 8:24.115 | קבועים |  |
| 8:24.116 | שונים |  |
| 8:24.117 | זה |  |
| 8:24.118 | מזה |  |
| 8:24.119 | בכמותם |  |
| 8:24.120 | ואיכותם |  |
| 8:24.121 | ומקומם | , |
| 8:24.122 | זה |  |
| 8:24.123 | למעלה |  |
| 8:24.124 | וזה |  |
| 8:24.125 | למטה | , |
| 8:24.126 | וזה |  |
| 8:24.127 | ארוך |  |
| 8:24.128 | וזה |  |
| 8:24.129 | קצר |  |
| 8:24.130 | כמו |  |
| 8:24.131 | שיראה |  |
| 8:24.132 | המעיין |  |
| 8:24.133 | בו |  |
| 8:24.134 | ובעץ |  |
| 8:24.135 | חיים |  |
| 8:24.136 | ופריו |  |
| 8:24.137 | לכהרח"ו | , |
| 8:24.138 | ושבש |  |
| 8:24.139 | ובלבל |  |
| 8:24.140 | אמונת |  |
| 8:24.141 | ישראל |  |
| 8:24.142 | הטהורה |  |
| 8:24.143 | המקובלת |  |
| 8:24.144 | בידם | , |
| 8:24.145 | והשוה |  |
| 8:24.146 | אותה |  |
| 8:24.147 | לאמונת |  |
| 8:24.148 | העכו"ם | . |
| 8:24.149 | ורבים |  |
| 8:24.150 | מעמנו |  |
| 8:24.151 | נמשכו |  |
| 8:24.152 | אחריו |  |
| 8:24.153 | לתומם |  |
| 8:24.154 | בהאמנים |  |
| 8:24.155 | שמהתנא |  |
| 8:24.156 | הקדוש |  |
| 8:24.157 | רשב"י |  |
| 8:24.158 | יצאו |  |
| 8:24.159 | הדברים |  |
| 8:24.160 | הכבושים |  |
| 8:24.161 | אשר |  |
| 8:24.162 | היו |  |
| 8:24.163 | קבורים | , |
| 8:24.164 | ולא |  |
| 8:24.165 | שמו |  |
| 8:24.166 | אל |  |
| 8:24.167 | לבם |  |
| 8:24.168 | איך |  |
| 8:24.169 | אפשר |  |
| 8:24.170 | לקבל |  |
| 8:24.171 | תורה |  |
| 8:24.172 | חדשה |  |
| 8:24.173 | מיד |  |
| 8:24.174 | עכו"ם | , |
| 8:24.175 | ואיך |  |
| 8:24.176 | נקברה |  |
| 8:24.177 | בארץ |  |
| 8:24.178 | בארגז | , |
| 8:24.179 | ומי |  |
| 8:24.180 | קברה | ? |
| 8:24.181 | ולמה |  |
| 8:24.182 | נתגלית |  |
| 8:24.183 | בימי |  |
| 8:24.184 | רשב"י |  |
| 8:24.185 | וחזרה |  |
| 8:24.186 | ונקברה |  |
| 8:24.187 | ונשכח |  |
| 8:24.188 | זכרה |  |
| 8:24.189 | מכל |  |
| 8:24.190 | היהודים | ? |
| 8:24.191 | אשר |  |
| 8:24.192 | על |  |
| 8:24.193 | כן |  |
| 8:24.194 | חלילה |  |
| 8:24.195 | לכל |  |
| 8:24.196 | משכיל |  |
| 8:24.197 | בתורת |  |
| 8:24.198 | ה | ' |
| 8:24.199 | שבכתב |  |
| 8:24.200 | ושבעל |  |
| 8:24.201 | פה |  |
| 8:24.202 | להאמין |  |
| 8:24.203 | שדברים |  |
| 8:24.204 | כאלה |  |
| 8:24.205 | יצאו |  |
| 8:24.206 | מפי |  |
| 8:24.207 | רשב"י |  |
| 8:24.208 | ע"ה |  |
| 8:24.209 | ולא |  |
| 8:24.210 | משום |  |
| 8:24.211 | תנא |  |
| 8:24.212 | או |  |
| 8:24.213 | אמורא | , |
| 8:24.214 | כי |  |
| 8:24.215 | אדרבא |  |
| 8:24.216 | מצינו |  |
| 8:24.217 | ראינו |  |
| 8:24.218 | שהיו |  |
| 8:24.219 | מרחיקים |  |
| 8:24.220 | אלו |  |
| 8:24.221 | הדעות |  |
| 8:24.222 | הרעות |  |
| 8:24.223 | ומתנגדים |  |
| 8:24.224 | להם | , |
| 8:24.225 | וכמו |  |
| 8:24.226 | שכבר |  |
| 8:24.227 | הבאתי |  |
| 8:24.228 | קצת |  |
| 8:24.229 | מאמרים |  |
| 8:24.230 | שמהם |  |
| 8:24.231 | נלמד |  |
| 8:24.232 | התנגדותם |  |
| 8:24.233 | לאלו |  |
| 8:24.234 | הדעות |  |
| 8:24.235 | החדשות |  |
| 8:24.236 | הנזכרים |  |
| 8:24.237 | בקבלה |  |
| 8:24.238 | החדשה | , |
| 8:24.239 | דממה |  |
| 8:24.240 | נפשך |  |
| 8:24.241 | הואיל |  |
| 8:24.242 | ויש |  |
| 8:24.243 | לפרצופים |  |
| 8:24.244 | דין |  |
| 8:24.245 | אלהות |  |
| 8:24.246 | כמו |  |
| 8:24.247 | שכתב |  |
| 8:24.248 | הרש"ב |  |
| 8:24.249 | בגליון |  |
| 8:24.250 | פרשת |  |
| 8:24.251 | בהר |  |
| 8:24.252 | דף |  |
| 8:24.253 | ק"ט |  |
| 8:24.254 | והרב |  |
| 8:24.255 | משנת |  |
| 8:24.256 | חסידים |  |
| 8:24.257 | סוף |  |
| 8:24.258 | מסכת |  |
| 8:24.259 | העשייה | , |
| 8:24.260 | ובית |  |
| 8:24.261 | קדש |  |
| 8:24.262 | הקדשים |  |
| 8:24.263 | דף |  |
| 8:24.264 | נ"ו |  |
| 8:24.265 | ע"ב | ! |
| 8:24.266 | נמצאנו |  |
| 8:24.267 | מאמינים |  |
| 8:24.268 | ועובדים |  |
| 8:24.269 | לאלוהות |  |
| 8:24.270 | רבים |  |
| 8:24.271 | נבראים |  |
| 8:24.272 | שכל |  |
| 8:24.273 | אחד |  |
| 8:24.274 | מהם |  |
| 8:24.275 | הוא |  |
| 8:24.276 | סבה |  |
| 8:24.277 | למה |  |
| 8:24.278 | שאחריו |  |
| 8:24.279 | כמו |  |
| 8:24.280 | שכתב |  |
| 8:24.281 | לעיל |  |
| 8:24.282 | בשם |  |
| 8:24.283 | זוהר |  |
| 8:24.284 | בראשית |  |
| 8:24.285 | דף |  |
| 8:24.286 | כ"ב |  |
| 8:24.287 | שהוא |  |
| 8:24.288 | קורא |  |
| 8:24.289 | לכל |  |
| 8:24.290 | אחד |  |
| 8:24.291 | מהם |  |
| 8:24.292 | בשם |  |
| 8:24.293 | הוי'ה |  |
| 8:24.294 | ואדנות |  |
| 8:24.295 | ואלהים |  |
| 8:24.296 | וקב"ה |  |
| 8:24.297 | ושאר |  |
| 8:24.298 | שמות |  |
| 8:24.299 | המיוחדים | , |
| 8:24.300 | עד |  |
| 8:24.301 | נדר |  |
| 8:24.302 | שכשישמע |  |
| 8:24.303 | האיש |  |
| 8:24.304 | הישראלי |  |
| 8:24.305 | את |  |
| 8:24.306 | חבירו |  |
| 8:24.307 | מזכיר |  |
| 8:24.308 | את |  |
| 8:24.309 | השם |  |
| 8:24.310 | הנכבד |  |
| 8:24.311 | הויה |  |
| 8:24.312 | או |  |
| 8:24.313 | אלהים |  |
| 8:24.314 | או |  |
| 8:24.315 | הקב"ה | , |
| 8:24.316 | צריך |  |
| 8:24.317 | השומע |  |
| 8:24.318 | לשאול |  |
| 8:24.319 | איזה |  |
| 8:24.320 | הוא |  |
| 8:24.321 | הקב"ה |  |
| 8:24.322 | שאתה |  |
| 8:24.323 | מזכיר | , |
| 8:24.324 | אם |  |
| 8:24.325 | עתיק | , |
| 8:24.326 | או |  |
| 8:24.327 | אריך | , |
| 8:24.328 | או |  |
| 8:24.329 | אבא | , |
| 8:24.330 | או |  |
| 8:24.331 | אימא | , |
| 8:24.332 | או |  |
| 8:24.333 | זעיר | , |
| 8:24.334 | או |  |
| 8:24.335 | נוקבא |  |
| 8:24.336 | דיליה | , |
| 8:24.337 | כמו |  |
| 8:24.338 | שעשה |  |
| 8:24.339 | הוא |  |
| 8:24.340 | עצמו |  |
| 8:24.341 | מחבר |  |
| 8:24.342 | הזהר |  |
| 8:24.343 | בפסוק |  |
| 8:24.344 | ויאמר |  |
| 8:24.345 | אלהים |  |
| 8:24.346 | נעשה |  |
| 8:24.347 | אדם | , |
| 8:24.348 | ואמר |  |
| 8:24.349 | מאן |  |
| 8:24.350 | אלהים | , |
| 8:24.351 | כדלעיל |  |
| 8:24.352 | סי | ' |
| 8:24.353 | כ | '! |
| 8:24.354 | כי |  |
| 8:24.355 | האלהים |  |
| 8:24.356 | העליון |  |
| 8:24.357 | לא |  |
| 8:24.358 | תוכל |  |
| 8:24.359 | לומר |  |
| 8:24.360 | לפי |  |
| 8:24.361 | דבריהם | ) |
| 8:24.362 | כי |  |
| 8:24.363 | הוא |  |
| 8:24.364 | אין |  |
| 8:24.365 | לו |  |
| 8:24.366 | שם |  |
| 8:24.367 | ולא |  |
| 8:24.368 | נקודה | , |
| 8:24.369 | ולא |  |
| 8:24.370 | תוכל |  |
| 8:24.371 | לכנותו |  |
| 8:24.372 | בשם |  |
| 8:24.373 | מן |  |
| 8:24.374 | השמות |  |
| 8:24.375 | הקדושים |  |
| 8:24.376 | כמו |  |
| 8:24.377 | שכתבו |  |
| 8:24.378 | המקובלים |  |
| 8:24.379 | החדשים | , |
| 8:24.380 | ונמצאנו |  |
| 8:24.381 | מאמינים |  |
| 8:24.382 | באלהים |  |
| 8:24.383 | רבים |  |
| 8:24.384 | נבראים |  |
| 8:24.385 | ח"ו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

