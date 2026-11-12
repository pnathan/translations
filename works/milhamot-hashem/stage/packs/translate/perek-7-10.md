# Translation pack: milhamot-hashem perek-7-10 (Milchamot ha-Shem 7:14–7:14)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-10`.

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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אלפין** → "alefs".
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


## Milchamot ha-Shem 7:14

Hebrew (Wikisource): פז) ורב סעדיה גאון בפי' דניאל כ' וז"ל ודע כי דינא יתיב וספרין פתיחו אשר ביארתי למעלה יום דין ופקודה, הוא יום העתיד לדרוש כל מעשה בני אדם החיים והמתים, ואשר כתב למעלה ועתיק יומין יתיב, כל האומר כי ראה ה' צבאות עין בעין. הרי הוא פושע, שמדמה דמות לשם ית' ואין לו דמות, דכ' ואל מי תדמיון אל, וכל האומר כי ראה מלאך הוא מן השונים, והוא מן האומרים קטון וגדול, למען כי אין שופט העולם כי אם הקב"ה לבדו, ככתוב לפני ה' כי בא לשפוט את הארץ, ואין מלאך אשר יוכל לדבר, כי אלף אלפין ישמשוניה. ועת הדע כי דניאל ראה כל אלה בחלום ולא במראה עינים וכל החלום יש לו פתרון. כי הם משל, כי גם כל נבואת הנביאים אשר ראו ותחת רגליו כמעשה לבנת נספיר. וממראה מתניו ולמעלה, יושב על כסא רם ונשא, לא ראו חק חקוקו, כי אין דמות לבורא ית' כלל, ולא דמות מלאך, כי בעת שהיה מדבר עמם מראה להם חזיון חדש. וכל דמות שראו יש לו פתרון. מקל שקד, כי שוקד אני רבבות גדולות, מעשיו בלא מספר, וכאשר ראה מיכה ברוח נבואה, וכל צבא השמים עומדים מימינו ומשמאלו, לא נוכל לדמות ליוצר הכל מימינו ומשמאלו, ולא יעמוד לפניו כל יצור, ואולם דברי כל אלה במשל כמעשה בני אדם עכ"ל. והנה כבר ראו עיניך ומלאו אזניך ממה שקדם כי ה' אלהינו לא שנה ולא ישתנה. כמאמר הנשביא אני ה' לא שניתי. וכבר האריכו רבותינו הקדמונים הרס"ג ובעל חובת הלבבות ור' יהודה הליו והרמב"ם בפי' המשנה ובס' היד ובמורה, ובעל שבילי אמונה שאין ליחס לשי"ת שום שינוי והתפעלות. וגם לא מינה אחר לשפוט את הארץ. ודברים אלו מוסמכים מכל רבותינו כמבואר בדבריהם. ורבותינו בעלי הקבלה החדשה שגו בכל אלה. כי האמינו לדברי נביא השקר המתנבא בשם עתיק יומין לעבוד אל אחר קצא אפים (זעיר אנפין) הוא מחבר הזהר אשר בילדי נכרים ישפיק. ותלה אותם ברשב"י התנא בשקר. ולכן שגו ברואה שלא די שיחסו לשי"ת רבוי וחלוק צורות ופרצופים שונים וקוים. אלא שיחסו לו שינוים רבים בעצמותו ית' ואמרו שמתחלה היה מקום וחלל גדול פנוי וריק להיות שם כל ברואיו, ונעשה ככדור עגול וחלול באמצעותו, ושוב נעשה בו שינוי אחר, כי המשיך מאורו הגדול המקיף, ספירות דעיגולי אדם קדמאה סתימא, ואחר זה עיגולי אדם קדמון. ועיגולי עתיק יומין, ועיגולי אריך אנפין, ועיגולי אבא ואימא. גם עיגולי זעיר ונוקביה. ושוב נעשו עוד בו שינויים רבים להמשיך ג"כ מעצמותו ספירות דיושר בתנית אדם. והם עיקר האלוהות לפי דעתם. תחלה פרצוף אדם קדמון. ואחריו פרצוף עתיק יומין. ואחריו אריך אנפין. ושניהם כולם זכר ונקבה. רק שזה ימין ושמאל. וזה פנים ואחור, ואחריהם פרצופי אבא ואימא זכר ונקבה נפרדים זה מזה. ואחריהם פרצופי זעיר ונוקביה ג"כ נפרדים. והם עיקר האלוהות אשר לפי דבריהם אלהים נקרא בכל תפלותינו. וכל ברכותינו אליהם יכונו. כי הם אלוהינו לפי דעתם. הרי נתבאר לך כמה שינוים וחליפות נעשו באלהינו על פי הקבלה החדשה הסותרת אמונת תוה"ק.

| id | word | punctuation after |
|---|---|---|
| 7:14.1 | פז | ) |
| 7:14.2 | ורב |  |
| 7:14.3 | סעדיה |  |
| 7:14.4 | גאון |  |
| 7:14.5 | בפי | ' |
| 7:14.6 | דניאל |  |
| 7:14.7 | כ | ' |
| 7:14.8 | וז"ל |  |
| 7:14.9 | ודע |  |
| 7:14.10 | כי |  |
| 7:14.11 | דינא |  |
| 7:14.12 | יתיב |  |
| 7:14.13 | וספרין |  |
| 7:14.14 | פתיחו |  |
| 7:14.15 | אשר |  |
| 7:14.16 | ביארתי |  |
| 7:14.17 | למעלה |  |
| 7:14.18 | יום |  |
| 7:14.19 | דין |  |
| 7:14.20 | ופקודה | , |
| 7:14.21 | הוא |  |
| 7:14.22 | יום |  |
| 7:14.23 | העתיד |  |
| 7:14.24 | לדרוש |  |
| 7:14.25 | כל |  |
| 7:14.26 | מעשה |  |
| 7:14.27 | בני |  |
| 7:14.28 | אדם |  |
| 7:14.29 | החיים |  |
| 7:14.30 | והמתים | , |
| 7:14.31 | ואשר |  |
| 7:14.32 | כתב |  |
| 7:14.33 | למעלה |  |
| 7:14.34 | ועתיק |  |
| 7:14.35 | יומין |  |
| 7:14.36 | יתיב | , |
| 7:14.37 | כל |  |
| 7:14.38 | האומר |  |
| 7:14.39 | כי |  |
| 7:14.40 | ראה |  |
| 7:14.41 | ה | ' |
| 7:14.42 | צבאות |  |
| 7:14.43 | עין |  |
| 7:14.44 | בעין | . |
| 7:14.45 | הרי |  |
| 7:14.46 | הוא |  |
| 7:14.47 | פושע | , |
| 7:14.48 | שמדמה |  |
| 7:14.49 | דמות |  |
| 7:14.50 | לשם |  |
| 7:14.51 | ית | ' |
| 7:14.52 | ואין |  |
| 7:14.53 | לו |  |
| 7:14.54 | דמות | , |
| 7:14.55 | דכ | ' |
| 7:14.56 | ואל |  |
| 7:14.57 | מי |  |
| 7:14.58 | תדמיון |  |
| 7:14.59 | אל | , |
| 7:14.60 | וכל |  |
| 7:14.61 | האומר |  |
| 7:14.62 | כי |  |
| 7:14.63 | ראה |  |
| 7:14.64 | מלאך |  |
| 7:14.65 | הוא |  |
| 7:14.66 | מן |  |
| 7:14.67 | השונים | , |
| 7:14.68 | והוא |  |
| 7:14.69 | מן |  |
| 7:14.70 | האומרים |  |
| 7:14.71 | קטון |  |
| 7:14.72 | וגדול | , |
| 7:14.73 | למען |  |
| 7:14.74 | כי |  |
| 7:14.75 | אין |  |
| 7:14.76 | שופט |  |
| 7:14.77 | העולם |  |
| 7:14.78 | כי |  |
| 7:14.79 | אם |  |
| 7:14.80 | הקב"ה |  |
| 7:14.81 | לבדו | , |
| 7:14.82 | ככתוב |  |
| 7:14.83 | לפני |  |
| 7:14.84 | ה | ' |
| 7:14.85 | כי |  |
| 7:14.86 | בא |  |
| 7:14.87 | לשפוט |  |
| 7:14.88 | את |  |
| 7:14.89 | הארץ | , |
| 7:14.90 | ואין |  |
| 7:14.91 | מלאך |  |
| 7:14.92 | אשר |  |
| 7:14.93 | יוכל |  |
| 7:14.94 | לדבר | , |
| 7:14.95 | כי |  |
| 7:14.96 | אלף |  |
| 7:14.97 | אלפין |  |
| 7:14.98 | ישמשוניה | . |
| 7:14.99 | ועת |  |
| 7:14.100 | הדע |  |
| 7:14.101 | כי |  |
| 7:14.102 | דניאל |  |
| 7:14.103 | ראה |  |
| 7:14.104 | כל |  |
| 7:14.105 | אלה |  |
| 7:14.106 | בחלום |  |
| 7:14.107 | ולא |  |
| 7:14.108 | במראה |  |
| 7:14.109 | עינים |  |
| 7:14.110 | וכל |  |
| 7:14.111 | החלום |  |
| 7:14.112 | יש |  |
| 7:14.113 | לו |  |
| 7:14.114 | פתרון | . |
| 7:14.115 | כי |  |
| 7:14.116 | הם |  |
| 7:14.117 | משל | , |
| 7:14.118 | כי |  |
| 7:14.119 | גם |  |
| 7:14.120 | כל |  |
| 7:14.121 | נבואת |  |
| 7:14.122 | הנביאים |  |
| 7:14.123 | אשר |  |
| 7:14.124 | ראו |  |
| 7:14.125 | ותחת |  |
| 7:14.126 | רגליו |  |
| 7:14.127 | כמעשה |  |
| 7:14.128 | לבנת |  |
| 7:14.129 | נספיר | . |
| 7:14.130 | וממראה |  |
| 7:14.131 | מתניו |  |
| 7:14.132 | ולמעלה | , |
| 7:14.133 | יושב |  |
| 7:14.134 | על |  |
| 7:14.135 | כסא |  |
| 7:14.136 | רם |  |
| 7:14.137 | ונשא | , |
| 7:14.138 | לא |  |
| 7:14.139 | ראו |  |
| 7:14.140 | חק |  |
| 7:14.141 | חקוקו | , |
| 7:14.142 | כי |  |
| 7:14.143 | אין |  |
| 7:14.144 | דמות |  |
| 7:14.145 | לבורא |  |
| 7:14.146 | ית | ' |
| 7:14.147 | כלל | , |
| 7:14.148 | ולא |  |
| 7:14.149 | דמות |  |
| 7:14.150 | מלאך | , |
| 7:14.151 | כי |  |
| 7:14.152 | בעת |  |
| 7:14.153 | שהיה |  |
| 7:14.154 | מדבר |  |
| 7:14.155 | עמם |  |
| 7:14.156 | מראה |  |
| 7:14.157 | להם |  |
| 7:14.158 | חזיון |  |
| 7:14.159 | חדש | . |
| 7:14.160 | וכל |  |
| 7:14.161 | דמות |  |
| 7:14.162 | שראו |  |
| 7:14.163 | יש |  |
| 7:14.164 | לו |  |
| 7:14.165 | פתרון | . |
| 7:14.166 | מקל |  |
| 7:14.167 | שקד | , |
| 7:14.168 | כי |  |
| 7:14.169 | שוקד |  |
| 7:14.170 | אני |  |
| 7:14.171 | רבבות |  |
| 7:14.172 | גדולות | , |
| 7:14.173 | מעשיו |  |
| 7:14.174 | בלא |  |
| 7:14.175 | מספר | , |
| 7:14.176 | וכאשר |  |
| 7:14.177 | ראה |  |
| 7:14.178 | מיכה |  |
| 7:14.179 | ברוח |  |
| 7:14.180 | נבואה | , |
| 7:14.181 | וכל |  |
| 7:14.182 | צבא |  |
| 7:14.183 | השמים |  |
| 7:14.184 | עומדים |  |
| 7:14.185 | מימינו |  |
| 7:14.186 | ומשמאלו | , |
| 7:14.187 | לא |  |
| 7:14.188 | נוכל |  |
| 7:14.189 | לדמות |  |
| 7:14.190 | ליוצר |  |
| 7:14.191 | הכל |  |
| 7:14.192 | מימינו |  |
| 7:14.193 | ומשמאלו | , |
| 7:14.194 | ולא |  |
| 7:14.195 | יעמוד |  |
| 7:14.196 | לפניו |  |
| 7:14.197 | כל |  |
| 7:14.198 | יצור | , |
| 7:14.199 | ואולם |  |
| 7:14.200 | דברי |  |
| 7:14.201 | כל |  |
| 7:14.202 | אלה |  |
| 7:14.203 | במשל |  |
| 7:14.204 | כמעשה |  |
| 7:14.205 | בני |  |
| 7:14.206 | אדם |  |
| 7:14.207 | עכ"ל | . |
| 7:14.208 | והנה |  |
| 7:14.209 | כבר |  |
| 7:14.210 | ראו |  |
| 7:14.211 | עיניך |  |
| 7:14.212 | ומלאו |  |
| 7:14.213 | אזניך |  |
| 7:14.214 | ממה |  |
| 7:14.215 | שקדם |  |
| 7:14.216 | כי |  |
| 7:14.217 | ה | ' |
| 7:14.218 | אלהינו |  |
| 7:14.219 | לא |  |
| 7:14.220 | שנה |  |
| 7:14.221 | ולא |  |
| 7:14.222 | ישתנה | . |
| 7:14.223 | כמאמר |  |
| 7:14.224 | הנשביא |  |
| 7:14.225 | אני |  |
| 7:14.226 | ה | ' |
| 7:14.227 | לא |  |
| 7:14.228 | שניתי | . |
| 7:14.229 | וכבר |  |
| 7:14.230 | האריכו |  |
| 7:14.231 | רבותינו |  |
| 7:14.232 | הקדמונים |  |
| 7:14.233 | הרס"ג |  |
| 7:14.234 | ובעל |  |
| 7:14.235 | חובת |  |
| 7:14.236 | הלבבות |  |
| 7:14.237 | ור | ' |
| 7:14.238 | יהודה |  |
| 7:14.239 | הליו |  |
| 7:14.240 | והרמב"ם |  |
| 7:14.241 | בפי | ' |
| 7:14.242 | המשנה |  |
| 7:14.243 | ובס | ' |
| 7:14.244 | היד |  |
| 7:14.245 | ובמורה | , |
| 7:14.246 | ובעל |  |
| 7:14.247 | שבילי |  |
| 7:14.248 | אמונה |  |
| 7:14.249 | שאין |  |
| 7:14.250 | ליחס |  |
| 7:14.251 | לשי"ת |  |
| 7:14.252 | שום |  |
| 7:14.253 | שינוי |  |
| 7:14.254 | והתפעלות | . |
| 7:14.255 | וגם |  |
| 7:14.256 | לא |  |
| 7:14.257 | מינה |  |
| 7:14.258 | אחר |  |
| 7:14.259 | לשפוט |  |
| 7:14.260 | את |  |
| 7:14.261 | הארץ | . |
| 7:14.262 | ודברים |  |
| 7:14.263 | אלו |  |
| 7:14.264 | מוסמכים |  |
| 7:14.265 | מכל |  |
| 7:14.266 | רבותינו |  |
| 7:14.267 | כמבואר |  |
| 7:14.268 | בדבריהם | . |
| 7:14.269 | ורבותינו |  |
| 7:14.270 | בעלי |  |
| 7:14.271 | הקבלה |  |
| 7:14.272 | החדשה |  |
| 7:14.273 | שגו |  |
| 7:14.274 | בכל |  |
| 7:14.275 | אלה | . |
| 7:14.276 | כי |  |
| 7:14.277 | האמינו |  |
| 7:14.278 | לדברי |  |
| 7:14.279 | נביא |  |
| 7:14.280 | השקר |  |
| 7:14.281 | המתנבא |  |
| 7:14.282 | בשם |  |
| 7:14.283 | עתיק |  |
| 7:14.284 | יומין |  |
| 7:14.285 | לעבוד |  |
| 7:14.286 | אל |  |
| 7:14.287 | אחר |  |
| 7:14.288 | קצא |  |
| 7:14.289 | אפים |  |
| 7:14.290 | זעיר |  |
| 7:14.291 | אנפין | ) |
| 7:14.292 | הוא |  |
| 7:14.293 | מחבר |  |
| 7:14.294 | הזהר |  |
| 7:14.295 | אשר |  |
| 7:14.296 | בילדי |  |
| 7:14.297 | נכרים |  |
| 7:14.298 | ישפיק | . |
| 7:14.299 | ותלה |  |
| 7:14.300 | אותם |  |
| 7:14.301 | ברשב"י |  |
| 7:14.302 | התנא |  |
| 7:14.303 | בשקר | . |
| 7:14.304 | ולכן |  |
| 7:14.305 | שגו |  |
| 7:14.306 | ברואה |  |
| 7:14.307 | שלא |  |
| 7:14.308 | די |  |
| 7:14.309 | שיחסו |  |
| 7:14.310 | לשי"ת |  |
| 7:14.311 | רבוי |  |
| 7:14.312 | וחלוק |  |
| 7:14.313 | צורות |  |
| 7:14.314 | ופרצופים |  |
| 7:14.315 | שונים |  |
| 7:14.316 | וקוים | . |
| 7:14.317 | אלא |  |
| 7:14.318 | שיחסו |  |
| 7:14.319 | לו |  |
| 7:14.320 | שינוים |  |
| 7:14.321 | רבים |  |
| 7:14.322 | בעצמותו |  |
| 7:14.323 | ית | ' |
| 7:14.324 | ואמרו |  |
| 7:14.325 | שמתחלה |  |
| 7:14.326 | היה |  |
| 7:14.327 | מקום |  |
| 7:14.328 | וחלל |  |
| 7:14.329 | גדול |  |
| 7:14.330 | פנוי |  |
| 7:14.331 | וריק |  |
| 7:14.332 | להיות |  |
| 7:14.333 | שם |  |
| 7:14.334 | כל |  |
| 7:14.335 | ברואיו | , |
| 7:14.336 | ונעשה |  |
| 7:14.337 | ככדור |  |
| 7:14.338 | עגול |  |
| 7:14.339 | וחלול |  |
| 7:14.340 | באמצעותו | , |
| 7:14.341 | ושוב |  |
| 7:14.342 | נעשה |  |
| 7:14.343 | בו |  |
| 7:14.344 | שינוי |  |
| 7:14.345 | אחר | , |
| 7:14.346 | כי |  |
| 7:14.347 | המשיך |  |
| 7:14.348 | מאורו |  |
| 7:14.349 | הגדול |  |
| 7:14.350 | המקיף | , |
| 7:14.351 | ספירות |  |
| 7:14.352 | דעיגולי |  |
| 7:14.353 | אדם |  |
| 7:14.354 | קדמאה |  |
| 7:14.355 | סתימא | , |
| 7:14.356 | ואחר |  |
| 7:14.357 | זה |  |
| 7:14.358 | עיגולי |  |
| 7:14.359 | אדם |  |
| 7:14.360 | קדמון | . |
| 7:14.361 | ועיגולי |  |
| 7:14.362 | עתיק |  |
| 7:14.363 | יומין | , |
| 7:14.364 | ועיגולי |  |
| 7:14.365 | אריך |  |
| 7:14.366 | אנפין | , |
| 7:14.367 | ועיגולי |  |
| 7:14.368 | אבא |  |
| 7:14.369 | ואימא | . |
| 7:14.370 | גם |  |
| 7:14.371 | עיגולי |  |
| 7:14.372 | זעיר |  |
| 7:14.373 | ונוקביה | . |
| 7:14.374 | ושוב |  |
| 7:14.375 | נעשו |  |
| 7:14.376 | עוד |  |
| 7:14.377 | בו |  |
| 7:14.378 | שינויים |  |
| 7:14.379 | רבים |  |
| 7:14.380 | להמשיך |  |
| 7:14.381 | ג"כ |  |
| 7:14.382 | מעצמותו |  |
| 7:14.383 | ספירות |  |
| 7:14.384 | דיושר |  |
| 7:14.385 | בתנית |  |
| 7:14.386 | אדם | . |
| 7:14.387 | והם |  |
| 7:14.388 | עיקר |  |
| 7:14.389 | האלוהות |  |
| 7:14.390 | לפי |  |
| 7:14.391 | דעתם | . |
| 7:14.392 | תחלה |  |
| 7:14.393 | פרצוף |  |
| 7:14.394 | אדם |  |
| 7:14.395 | קדמון | . |
| 7:14.396 | ואחריו |  |
| 7:14.397 | פרצוף |  |
| 7:14.398 | עתיק |  |
| 7:14.399 | יומין | . |
| 7:14.400 | ואחריו |  |
| 7:14.401 | אריך |  |
| 7:14.402 | אנפין | . |
| 7:14.403 | ושניהם |  |
| 7:14.404 | כולם |  |
| 7:14.405 | זכר |  |
| 7:14.406 | ונקבה | . |
| 7:14.407 | רק |  |
| 7:14.408 | שזה |  |
| 7:14.409 | ימין |  |
| 7:14.410 | ושמאל | . |
| 7:14.411 | וזה |  |
| 7:14.412 | פנים |  |
| 7:14.413 | ואחור | , |
| 7:14.414 | ואחריהם |  |
| 7:14.415 | פרצופי |  |
| 7:14.416 | אבא |  |
| 7:14.417 | ואימא |  |
| 7:14.418 | זכר |  |
| 7:14.419 | ונקבה |  |
| 7:14.420 | נפרדים |  |
| 7:14.421 | זה |  |
| 7:14.422 | מזה | . |
| 7:14.423 | ואחריהם |  |
| 7:14.424 | פרצופי |  |
| 7:14.425 | זעיר |  |
| 7:14.426 | ונוקביה |  |
| 7:14.427 | ג"כ |  |
| 7:14.428 | נפרדים | . |
| 7:14.429 | והם |  |
| 7:14.430 | עיקר |  |
| 7:14.431 | האלוהות |  |
| 7:14.432 | אשר |  |
| 7:14.433 | לפי |  |
| 7:14.434 | דבריהם |  |
| 7:14.435 | אלהים |  |
| 7:14.436 | נקרא |  |
| 7:14.437 | בכל |  |
| 7:14.438 | תפלותינו | . |
| 7:14.439 | וכל |  |
| 7:14.440 | ברכותינו |  |
| 7:14.441 | אליהם |  |
| 7:14.442 | יכונו | . |
| 7:14.443 | כי |  |
| 7:14.444 | הם |  |
| 7:14.445 | אלוהינו |  |
| 7:14.446 | לפי |  |
| 7:14.447 | דעתם | . |
| 7:14.448 | הרי |  |
| 7:14.449 | נתבאר |  |
| 7:14.450 | לך |  |
| 7:14.451 | כמה |  |
| 7:14.452 | שינוים |  |
| 7:14.453 | וחליפות |  |
| 7:14.454 | נעשו |  |
| 7:14.455 | באלהינו |  |
| 7:14.456 | על |  |
| 7:14.457 | פי |  |
| 7:14.458 | הקבלה |  |
| 7:14.459 | החדשה |  |
| 7:14.460 | הסותרת |  |
| 7:14.461 | אמונת |  |
| 7:14.462 | תוה"ק | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

