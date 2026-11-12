# Translation pack: milhamot-hashem perek-8-11 (Milchamot ha-Shem 8:13–8:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-11`.

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
- **יחוד / ייחוד / ייחדו**: the doctrine → "unity", fixed, not "monotheism" or
  "oneness" (אחדות → "oneness"). The Kabbalistic act → "unification" (יחודים
  → "unifications"). The verb, God as object → "to declare one" ("ליחדו" → "to
  declare Him one"); Kabbalistic entities as object → "to unify"; of
  dedicating a thing to one → "dedicate ... solely to" (note: "unify"). Where
  he plays on the senses: "unity" for the noun, "unify" for the verb, the play
  in a note.
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 8:13

Hebrew (Wikisource): קח) והנה כבר ביארנו כמה פעמים שדעת המקובלים החדשים שהזעיר אנפין הוא שנגלה על אבותינו בהר סיני ואמר אנכי ה' אלהיך, ואם כן איך הביאו רז"ל בשמות רבה פסוק אני ראשון ואני אחרון וכו' שאין לי אב, ולא בן ולא אח, על אנכי ה' אלהיך שאמרו המקובלים החדשים שהזעיר אנפין אמרו, ומה לו לאלהים העליון להשתבח בעצמו במקום זה תוך דברי הזעיר אנפין שיש לו אב ואם ובן ואחות? אלא ודאי שרז"ל בב"ר ובשמות רבה ודברים רבה, באו להעיר אותנו שלא נטעה באלו הדעות, ולבאר לנו, כי זה שאמר אנכי ה' אלהיך הוא שאמר אני ראשון שאין לי אב, ואני אחרון שאין לי בן, ומבלעדי אין אלהים שאין לי אח, הוא ה' אלהינו הסבה הראשונה הוא ולא אחר, והוא בכבודו ובעצמו מן השמים השמיענו את קולו, והזהרנו בדבור השני לא יהיה לך אלהים אחרים על פני וכו'! והוא שאמר ראו עתה כי אני אני הוא ואין אלהים עמדי! ודלא כדעת הזהר המסית דס"ל שהאלוה שאמר אני ראשון וכו' הוא האין סוף, ושאמר ראו עתה וכו' הוא אדם קדמון, ושאמר יהי אור וכו' הוא אבא שאמר לאימא, ושאמר נעשה אדם וכו' היא אימא שאמרה לאבא וכו' כדלעיל, ומכל מקום לא יחסו כל עבודותינו ותפלותינו וכל בקשותינו רק לקצר אפים (זעיר אנפין) ובצירוף פרצופים שלמעלה ממנו כמו שנת' לעיל ובקראינו אליו כאילו קראינו לאלהים עליון. כי האלהים העליון אין לו שם. ואם נקראיהו בשם יהוה לא יענה. כי לא נתיחסו שמות רק לפרצופים הנאצלים ונבראים ממנו. ושם יהוה ביחוד לפי דבריהם הוא האלוה קצר אפים דוקא. ועל האמונות הזרות האלה הרחוקות מדרכי תורתינו הקדושה וקבלת רז"ל קרא ירמיה הנביא ע"ה בדבר ה' לאמר (ירמיהו ה׳:י״ב) כחשו בה' ויאמרו לוא הוא, כחשו בה' הנאמר בו אלהים בצאתך לפני עמך בצעדך בישימון סלה ארץ ועשה אף שמים נטפו מפני אלהים זה סיני מפני אלהים אלהי ישראל, ונאמר ה' מסיני בא וזרח משעיר וכו'. ויאמר לא הוא, אלא הזעיר אנפין הוא שירד על הר סיני ובא להתחבר עם אשתו שכבר קדמה לו והוא ירד אחריה להתחבר ולהזדווג לה, לוא כתיב מלא וי"ו ואלף שחשבו מציאותו יתברך כאלו אינו, וגם קראהו אין, כלומר לא יש, ואמרו שהקורא אליו לא יענהו, כי הוא אין, לפי שידוע שאותיות השלילה והמניעה וההעדר אשר ישתמשו בהם הם לא ואין כמו שכתב הרמב"ם בספר המצות שורש ח', והערביים משתמשים להורות על השלילה וההעדר הנקרא בערבי נפי במלות לם, וליס, ולא, ומא, והמקובלים החדשים קראו לשי"ת אין שהוא לשון שלילה והעדר, ולכן אמרו שהוא יתברך אין לו שום שם ונקודה שהגבילהו כי אין הוא, וצירפו לפעמים גם מלת סוף, ואמרו אין סוף, והיא מליצה הדומה בעניינה לתרגום אין חקר לתבונתו, שתרגם יונתן בן עוזיאל לית סוף לסוכלתנותיה. והשמיטו מלת תבונתו שאמר הנביא ע"ה לבאר לנו שהיש המצוי הזה הוא ה' אלהינו אין חקר וסוף לתבונתו שהמציא כל נמצא בחכמה נפלאה. כמאמר דוד המלך ע"ה מה רבו מעשיך ה' כלם בחכמה עשית. שאינו ניכר אלא מתוך מעשיו. כמאמר הנביא שאו מרום עיניכם וראו מי ברא אלה וכו' וכן אמר דוד כי אראה שמיך מעשה אצבעותיך! וחשבו שכיון שאי אפשר להשיג מהותו והיאך הוא. אם כן אין הוא ח"ו! ולכן יחסו כל השמות הנזכרים בתורה ובנביאים ובלשון רז"ל הקב"ה או רבונו ש"ע הכל על הפרצופים שהמציאו בדמיונם שהם משתלשלים ונמשכים מהאין וקראון גופא דמלכא, ואלו הפרצופים הם עיקר האלהות, כי הם מושגים ומצויירים בשכל, אלא שהם אורות ואינם בעלי בשר ועצמות כמו שכ' הרמב"ם בפתיחת ספר המורה! ונתאוו לאלוהות הרבה כאבותינו אנשי דוד המדבר שאמרו אלה אלהיך ישראל כמאמר רשב"י ע"ה כדגרסינן בפרק ד' מיתות. אמר רב יהודה אלמלא וי"ו שבהעלוך נחייבו שונאיהם של ישראל כלייה. אמר לו רשב"י והלא כל המשתף שם שמים ודבר אחר נעקר מן העולם שנ' בלתו לה' לבדו אלא מה ת"ל אשר העלוך מלמד שאיוו לאלהות הרבה, ועיין לעיל סי' נ"ח.

| id | word | punctuation after |
|---|---|---|
| 8:13.1 | קח | ) |
| 8:13.2 | והנה |  |
| 8:13.3 | כבר |  |
| 8:13.4 | ביארנו |  |
| 8:13.5 | כמה |  |
| 8:13.6 | פעמים |  |
| 8:13.7 | שדעת |  |
| 8:13.8 | המקובלים |  |
| 8:13.9 | החדשים |  |
| 8:13.10 | שהזעיר |  |
| 8:13.11 | אנפין |  |
| 8:13.12 | הוא |  |
| 8:13.13 | שנגלה |  |
| 8:13.14 | על |  |
| 8:13.15 | אבותינו |  |
| 8:13.16 | בהר |  |
| 8:13.17 | סיני |  |
| 8:13.18 | ואמר |  |
| 8:13.19 | אנכי |  |
| 8:13.20 | ה | ' |
| 8:13.21 | אלהיך | , |
| 8:13.22 | ואם |  |
| 8:13.23 | כן |  |
| 8:13.24 | איך |  |
| 8:13.25 | הביאו |  |
| 8:13.26 | רז"ל |  |
| 8:13.27 | בשמות |  |
| 8:13.28 | רבה |  |
| 8:13.29 | פסוק |  |
| 8:13.30 | אני |  |
| 8:13.31 | ראשון |  |
| 8:13.32 | ואני |  |
| 8:13.33 | אחרון |  |
| 8:13.34 | וכו | ' |
| 8:13.35 | שאין |  |
| 8:13.36 | לי |  |
| 8:13.37 | אב | , |
| 8:13.38 | ולא |  |
| 8:13.39 | בן |  |
| 8:13.40 | ולא |  |
| 8:13.41 | אח | , |
| 8:13.42 | על |  |
| 8:13.43 | אנכי |  |
| 8:13.44 | ה | ' |
| 8:13.45 | אלהיך |  |
| 8:13.46 | שאמרו |  |
| 8:13.47 | המקובלים |  |
| 8:13.48 | החדשים |  |
| 8:13.49 | שהזעיר |  |
| 8:13.50 | אנפין |  |
| 8:13.51 | אמרו | , |
| 8:13.52 | ומה |  |
| 8:13.53 | לו |  |
| 8:13.54 | לאלהים |  |
| 8:13.55 | העליון |  |
| 8:13.56 | להשתבח |  |
| 8:13.57 | בעצמו |  |
| 8:13.58 | במקום |  |
| 8:13.59 | זה |  |
| 8:13.60 | תוך |  |
| 8:13.61 | דברי |  |
| 8:13.62 | הזעיר |  |
| 8:13.63 | אנפין |  |
| 8:13.64 | שיש |  |
| 8:13.65 | לו |  |
| 8:13.66 | אב |  |
| 8:13.67 | ואם |  |
| 8:13.68 | ובן |  |
| 8:13.69 | ואחות | ? |
| 8:13.70 | אלא |  |
| 8:13.71 | ודאי |  |
| 8:13.72 | שרז"ל |  |
| 8:13.73 | בב"ר |  |
| 8:13.74 | ובשמות |  |
| 8:13.75 | רבה |  |
| 8:13.76 | ודברים |  |
| 8:13.77 | רבה | , |
| 8:13.78 | באו |  |
| 8:13.79 | להעיר |  |
| 8:13.80 | אותנו |  |
| 8:13.81 | שלא |  |
| 8:13.82 | נטעה |  |
| 8:13.83 | באלו |  |
| 8:13.84 | הדעות | , |
| 8:13.85 | ולבאר |  |
| 8:13.86 | לנו | , |
| 8:13.87 | כי |  |
| 8:13.88 | זה |  |
| 8:13.89 | שאמר |  |
| 8:13.90 | אנכי |  |
| 8:13.91 | ה | ' |
| 8:13.92 | אלהיך |  |
| 8:13.93 | הוא |  |
| 8:13.94 | שאמר |  |
| 8:13.95 | אני |  |
| 8:13.96 | ראשון |  |
| 8:13.97 | שאין |  |
| 8:13.98 | לי |  |
| 8:13.99 | אב | , |
| 8:13.100 | ואני |  |
| 8:13.101 | אחרון |  |
| 8:13.102 | שאין |  |
| 8:13.103 | לי |  |
| 8:13.104 | בן | , |
| 8:13.105 | ומבלעדי |  |
| 8:13.106 | אין |  |
| 8:13.107 | אלהים |  |
| 8:13.108 | שאין |  |
| 8:13.109 | לי |  |
| 8:13.110 | אח | , |
| 8:13.111 | הוא |  |
| 8:13.112 | ה | ' |
| 8:13.113 | אלהינו |  |
| 8:13.114 | הסבה |  |
| 8:13.115 | הראשונה |  |
| 8:13.116 | הוא |  |
| 8:13.117 | ולא |  |
| 8:13.118 | אחר | , |
| 8:13.119 | והוא |  |
| 8:13.120 | בכבודו |  |
| 8:13.121 | ובעצמו |  |
| 8:13.122 | מן |  |
| 8:13.123 | השמים |  |
| 8:13.124 | השמיענו |  |
| 8:13.125 | את |  |
| 8:13.126 | קולו | , |
| 8:13.127 | והזהרנו |  |
| 8:13.128 | בדבור |  |
| 8:13.129 | השני |  |
| 8:13.130 | לא |  |
| 8:13.131 | יהיה |  |
| 8:13.132 | לך |  |
| 8:13.133 | אלהים |  |
| 8:13.134 | אחרים |  |
| 8:13.135 | על |  |
| 8:13.136 | פני |  |
| 8:13.137 | וכו | '! |
| 8:13.138 | והוא |  |
| 8:13.139 | שאמר |  |
| 8:13.140 | ראו |  |
| 8:13.141 | עתה |  |
| 8:13.142 | כי |  |
| 8:13.143 | אני |  |
| 8:13.144 | אני |  |
| 8:13.145 | הוא |  |
| 8:13.146 | ואין |  |
| 8:13.147 | אלהים |  |
| 8:13.148 | עמדי | ! |
| 8:13.149 | ודלא |  |
| 8:13.150 | כדעת |  |
| 8:13.151 | הזהר |  |
| 8:13.152 | המסית |  |
| 8:13.153 | דס"ל |  |
| 8:13.154 | שהאלוה |  |
| 8:13.155 | שאמר |  |
| 8:13.156 | אני |  |
| 8:13.157 | ראשון |  |
| 8:13.158 | וכו | ' |
| 8:13.159 | הוא |  |
| 8:13.160 | האין |  |
| 8:13.161 | סוף | , |
| 8:13.162 | ושאמר |  |
| 8:13.163 | ראו |  |
| 8:13.164 | עתה |  |
| 8:13.165 | וכו | ' |
| 8:13.166 | הוא |  |
| 8:13.167 | אדם |  |
| 8:13.168 | קדמון | , |
| 8:13.169 | ושאמר |  |
| 8:13.170 | יהי |  |
| 8:13.171 | אור |  |
| 8:13.172 | וכו | ' |
| 8:13.173 | הוא |  |
| 8:13.174 | אבא |  |
| 8:13.175 | שאמר |  |
| 8:13.176 | לאימא | , |
| 8:13.177 | ושאמר |  |
| 8:13.178 | נעשה |  |
| 8:13.179 | אדם |  |
| 8:13.180 | וכו | ' |
| 8:13.181 | היא |  |
| 8:13.182 | אימא |  |
| 8:13.183 | שאמרה |  |
| 8:13.184 | לאבא |  |
| 8:13.185 | וכו | ' |
| 8:13.186 | כדלעיל | , |
| 8:13.187 | ומכל |  |
| 8:13.188 | מקום |  |
| 8:13.189 | לא |  |
| 8:13.190 | יחסו |  |
| 8:13.191 | כל |  |
| 8:13.192 | עבודותינו |  |
| 8:13.193 | ותפלותינו |  |
| 8:13.194 | וכל |  |
| 8:13.195 | בקשותינו |  |
| 8:13.196 | רק |  |
| 8:13.197 | לקצר |  |
| 8:13.198 | אפים |  |
| 8:13.199 | זעיר |  |
| 8:13.200 | אנפין | ) |
| 8:13.201 | ובצירוף |  |
| 8:13.202 | פרצופים |  |
| 8:13.203 | שלמעלה |  |
| 8:13.204 | ממנו |  |
| 8:13.205 | כמו |  |
| 8:13.206 | שנת | ' |
| 8:13.207 | לעיל |  |
| 8:13.208 | ובקראינו |  |
| 8:13.209 | אליו |  |
| 8:13.210 | כאילו |  |
| 8:13.211 | קראינו |  |
| 8:13.212 | לאלהים |  |
| 8:13.213 | עליון | . |
| 8:13.214 | כי |  |
| 8:13.215 | האלהים |  |
| 8:13.216 | העליון |  |
| 8:13.217 | אין |  |
| 8:13.218 | לו |  |
| 8:13.219 | שם | . |
| 8:13.220 | ואם |  |
| 8:13.221 | נקראיהו |  |
| 8:13.222 | בשם |  |
| 8:13.223 | יהוה |  |
| 8:13.224 | לא |  |
| 8:13.225 | יענה | . |
| 8:13.226 | כי |  |
| 8:13.227 | לא |  |
| 8:13.228 | נתיחסו |  |
| 8:13.229 | שמות |  |
| 8:13.230 | רק |  |
| 8:13.231 | לפרצופים |  |
| 8:13.232 | הנאצלים |  |
| 8:13.233 | ונבראים |  |
| 8:13.234 | ממנו | . |
| 8:13.235 | ושם |  |
| 8:13.236 | יהוה |  |
| 8:13.237 | ביחוד |  |
| 8:13.238 | לפי |  |
| 8:13.239 | דבריהם |  |
| 8:13.240 | הוא |  |
| 8:13.241 | האלוה |  |
| 8:13.242 | קצר |  |
| 8:13.243 | אפים |  |
| 8:13.244 | דוקא | . |
| 8:13.245 | ועל |  |
| 8:13.246 | האמונות |  |
| 8:13.247 | הזרות |  |
| 8:13.248 | האלה |  |
| 8:13.249 | הרחוקות |  |
| 8:13.250 | מדרכי |  |
| 8:13.251 | תורתינו |  |
| 8:13.252 | הקדושה |  |
| 8:13.253 | וקבלת |  |
| 8:13.254 | רז"ל |  |
| 8:13.255 | קרא |  |
| 8:13.256 | ירמיה |  |
| 8:13.257 | הנביא |  |
| 8:13.258 | ע"ה |  |
| 8:13.259 | בדבר |  |
| 8:13.260 | ה | ' |
| 8:13.261 | לאמר |  |
| 8:13.262 | ירמיהו |  |
| 8:13.263 | ה׳:י״ב | ) |
| 8:13.264 | כחשו |  |
| 8:13.265 | בה | ' |
| 8:13.266 | ויאמרו |  |
| 8:13.267 | לוא |  |
| 8:13.268 | הוא | , |
| 8:13.269 | כחשו |  |
| 8:13.270 | בה | ' |
| 8:13.271 | הנאמר |  |
| 8:13.272 | בו |  |
| 8:13.273 | אלהים |  |
| 8:13.274 | בצאתך |  |
| 8:13.275 | לפני |  |
| 8:13.276 | עמך |  |
| 8:13.277 | בצעדך |  |
| 8:13.278 | בישימון |  |
| 8:13.279 | סלה |  |
| 8:13.280 | ארץ |  |
| 8:13.281 | ועשה |  |
| 8:13.282 | אף |  |
| 8:13.283 | שמים |  |
| 8:13.284 | נטפו |  |
| 8:13.285 | מפני |  |
| 8:13.286 | אלהים |  |
| 8:13.287 | זה |  |
| 8:13.288 | סיני |  |
| 8:13.289 | מפני |  |
| 8:13.290 | אלהים |  |
| 8:13.291 | אלהי |  |
| 8:13.292 | ישראל | , |
| 8:13.293 | ונאמר |  |
| 8:13.294 | ה | ' |
| 8:13.295 | מסיני |  |
| 8:13.296 | בא |  |
| 8:13.297 | וזרח |  |
| 8:13.298 | משעיר |  |
| 8:13.299 | וכו | '. |
| 8:13.300 | ויאמר |  |
| 8:13.301 | לא |  |
| 8:13.302 | הוא | , |
| 8:13.303 | אלא |  |
| 8:13.304 | הזעיר |  |
| 8:13.305 | אנפין |  |
| 8:13.306 | הוא |  |
| 8:13.307 | שירד |  |
| 8:13.308 | על |  |
| 8:13.309 | הר |  |
| 8:13.310 | סיני |  |
| 8:13.311 | ובא |  |
| 8:13.312 | להתחבר |  |
| 8:13.313 | עם |  |
| 8:13.314 | אשתו |  |
| 8:13.315 | שכבר |  |
| 8:13.316 | קדמה |  |
| 8:13.317 | לו |  |
| 8:13.318 | והוא |  |
| 8:13.319 | ירד |  |
| 8:13.320 | אחריה |  |
| 8:13.321 | להתחבר |  |
| 8:13.322 | ולהזדווג |  |
| 8:13.323 | לה | , |
| 8:13.324 | לוא |  |
| 8:13.325 | כתיב |  |
| 8:13.326 | מלא |  |
| 8:13.327 | וי"ו |  |
| 8:13.328 | ואלף |  |
| 8:13.329 | שחשבו |  |
| 8:13.330 | מציאותו |  |
| 8:13.331 | יתברך |  |
| 8:13.332 | כאלו |  |
| 8:13.333 | אינו | , |
| 8:13.334 | וגם |  |
| 8:13.335 | קראהו |  |
| 8:13.336 | אין | , |
| 8:13.337 | כלומר |  |
| 8:13.338 | לא |  |
| 8:13.339 | יש | , |
| 8:13.340 | ואמרו |  |
| 8:13.341 | שהקורא |  |
| 8:13.342 | אליו |  |
| 8:13.343 | לא |  |
| 8:13.344 | יענהו | , |
| 8:13.345 | כי |  |
| 8:13.346 | הוא |  |
| 8:13.347 | אין | , |
| 8:13.348 | לפי |  |
| 8:13.349 | שידוע |  |
| 8:13.350 | שאותיות |  |
| 8:13.351 | השלילה |  |
| 8:13.352 | והמניעה |  |
| 8:13.353 | וההעדר |  |
| 8:13.354 | אשר |  |
| 8:13.355 | ישתמשו |  |
| 8:13.356 | בהם |  |
| 8:13.357 | הם |  |
| 8:13.358 | לא |  |
| 8:13.359 | ואין |  |
| 8:13.360 | כמו |  |
| 8:13.361 | שכתב |  |
| 8:13.362 | הרמב"ם |  |
| 8:13.363 | בספר |  |
| 8:13.364 | המצות |  |
| 8:13.365 | שורש |  |
| 8:13.366 | ח | ', |
| 8:13.367 | והערביים |  |
| 8:13.368 | משתמשים |  |
| 8:13.369 | להורות |  |
| 8:13.370 | על |  |
| 8:13.371 | השלילה |  |
| 8:13.372 | וההעדר |  |
| 8:13.373 | הנקרא |  |
| 8:13.374 | בערבי |  |
| 8:13.375 | נפי |  |
| 8:13.376 | במלות |  |
| 8:13.377 | לם | , |
| 8:13.378 | וליס | , |
| 8:13.379 | ולא | , |
| 8:13.380 | ומא | , |
| 8:13.381 | והמקובלים |  |
| 8:13.382 | החדשים |  |
| 8:13.383 | קראו |  |
| 8:13.384 | לשי"ת |  |
| 8:13.385 | אין |  |
| 8:13.386 | שהוא |  |
| 8:13.387 | לשון |  |
| 8:13.388 | שלילה |  |
| 8:13.389 | והעדר | , |
| 8:13.390 | ולכן |  |
| 8:13.391 | אמרו |  |
| 8:13.392 | שהוא |  |
| 8:13.393 | יתברך |  |
| 8:13.394 | אין |  |
| 8:13.395 | לו |  |
| 8:13.396 | שום |  |
| 8:13.397 | שם |  |
| 8:13.398 | ונקודה |  |
| 8:13.399 | שהגבילהו |  |
| 8:13.400 | כי |  |
| 8:13.401 | אין |  |
| 8:13.402 | הוא | , |
| 8:13.403 | וצירפו |  |
| 8:13.404 | לפעמים |  |
| 8:13.405 | גם |  |
| 8:13.406 | מלת |  |
| 8:13.407 | סוף | , |
| 8:13.408 | ואמרו |  |
| 8:13.409 | אין |  |
| 8:13.410 | סוף | , |
| 8:13.411 | והיא |  |
| 8:13.412 | מליצה |  |
| 8:13.413 | הדומה |  |
| 8:13.414 | בעניינה |  |
| 8:13.415 | לתרגום |  |
| 8:13.416 | אין |  |
| 8:13.417 | חקר |  |
| 8:13.418 | לתבונתו | , |
| 8:13.419 | שתרגם |  |
| 8:13.420 | יונתן |  |
| 8:13.421 | בן |  |
| 8:13.422 | עוזיאל |  |
| 8:13.423 | לית |  |
| 8:13.424 | סוף |  |
| 8:13.425 | לסוכלתנותיה | . |
| 8:13.426 | והשמיטו |  |
| 8:13.427 | מלת |  |
| 8:13.428 | תבונתו |  |
| 8:13.429 | שאמר |  |
| 8:13.430 | הנביא |  |
| 8:13.431 | ע"ה |  |
| 8:13.432 | לבאר |  |
| 8:13.433 | לנו |  |
| 8:13.434 | שהיש |  |
| 8:13.435 | המצוי |  |
| 8:13.436 | הזה |  |
| 8:13.437 | הוא |  |
| 8:13.438 | ה | ' |
| 8:13.439 | אלהינו |  |
| 8:13.440 | אין |  |
| 8:13.441 | חקר |  |
| 8:13.442 | וסוף |  |
| 8:13.443 | לתבונתו |  |
| 8:13.444 | שהמציא |  |
| 8:13.445 | כל |  |
| 8:13.446 | נמצא |  |
| 8:13.447 | בחכמה |  |
| 8:13.448 | נפלאה | . |
| 8:13.449 | כמאמר |  |
| 8:13.450 | דוד |  |
| 8:13.451 | המלך |  |
| 8:13.452 | ע"ה |  |
| 8:13.453 | מה |  |
| 8:13.454 | רבו |  |
| 8:13.455 | מעשיך |  |
| 8:13.456 | ה | ' |
| 8:13.457 | כלם |  |
| 8:13.458 | בחכמה |  |
| 8:13.459 | עשית | . |
| 8:13.460 | שאינו |  |
| 8:13.461 | ניכר |  |
| 8:13.462 | אלא |  |
| 8:13.463 | מתוך |  |
| 8:13.464 | מעשיו | . |
| 8:13.465 | כמאמר |  |
| 8:13.466 | הנביא |  |
| 8:13.467 | שאו |  |
| 8:13.468 | מרום |  |
| 8:13.469 | עיניכם |  |
| 8:13.470 | וראו |  |
| 8:13.471 | מי |  |
| 8:13.472 | ברא |  |
| 8:13.473 | אלה |  |
| 8:13.474 | וכו | ' |
| 8:13.475 | וכן |  |
| 8:13.476 | אמר |  |
| 8:13.477 | דוד |  |
| 8:13.478 | כי |  |
| 8:13.479 | אראה |  |
| 8:13.480 | שמיך |  |
| 8:13.481 | מעשה |  |
| 8:13.482 | אצבעותיך | ! |
| 8:13.483 | וחשבו |  |
| 8:13.484 | שכיון |  |
| 8:13.485 | שאי |  |
| 8:13.486 | אפשר |  |
| 8:13.487 | להשיג |  |
| 8:13.488 | מהותו |  |
| 8:13.489 | והיאך |  |
| 8:13.490 | הוא | . |
| 8:13.491 | אם |  |
| 8:13.492 | כן |  |
| 8:13.493 | אין |  |
| 8:13.494 | הוא |  |
| 8:13.495 | ח"ו | ! |
| 8:13.496 | ולכן |  |
| 8:13.497 | יחסו |  |
| 8:13.498 | כל |  |
| 8:13.499 | השמות |  |
| 8:13.500 | הנזכרים |  |
| 8:13.501 | בתורה |  |
| 8:13.502 | ובנביאים |  |
| 8:13.503 | ובלשון |  |
| 8:13.504 | רז"ל |  |
| 8:13.505 | הקב"ה |  |
| 8:13.506 | או |  |
| 8:13.507 | רבונו |  |
| 8:13.508 | ש"ע |  |
| 8:13.509 | הכל |  |
| 8:13.510 | על |  |
| 8:13.511 | הפרצופים |  |
| 8:13.512 | שהמציאו |  |
| 8:13.513 | בדמיונם |  |
| 8:13.514 | שהם |  |
| 8:13.515 | משתלשלים |  |
| 8:13.516 | ונמשכים |  |
| 8:13.517 | מהאין |  |
| 8:13.518 | וקראון |  |
| 8:13.519 | גופא |  |
| 8:13.520 | דמלכא | , |
| 8:13.521 | ואלו |  |
| 8:13.522 | הפרצופים |  |
| 8:13.523 | הם |  |
| 8:13.524 | עיקר |  |
| 8:13.525 | האלהות | , |
| 8:13.526 | כי |  |
| 8:13.527 | הם |  |
| 8:13.528 | מושגים |  |
| 8:13.529 | ומצויירים |  |
| 8:13.530 | בשכל | , |
| 8:13.531 | אלא |  |
| 8:13.532 | שהם |  |
| 8:13.533 | אורות |  |
| 8:13.534 | ואינם |  |
| 8:13.535 | בעלי |  |
| 8:13.536 | בשר |  |
| 8:13.537 | ועצמות |  |
| 8:13.538 | כמו |  |
| 8:13.539 | שכ | ' |
| 8:13.540 | הרמב"ם |  |
| 8:13.541 | בפתיחת |  |
| 8:13.542 | ספר |  |
| 8:13.543 | המורה | ! |
| 8:13.544 | ונתאוו |  |
| 8:13.545 | לאלוהות |  |
| 8:13.546 | הרבה |  |
| 8:13.547 | כאבותינו |  |
| 8:13.548 | אנשי |  |
| 8:13.549 | דוד |  |
| 8:13.550 | המדבר |  |
| 8:13.551 | שאמרו |  |
| 8:13.552 | אלה |  |
| 8:13.553 | אלהיך |  |
| 8:13.554 | ישראל |  |
| 8:13.555 | כמאמר |  |
| 8:13.556 | רשב"י |  |
| 8:13.557 | ע"ה |  |
| 8:13.558 | כדגרסינן |  |
| 8:13.559 | בפרק |  |
| 8:13.560 | ד | ' |
| 8:13.561 | מיתות | . |
| 8:13.562 | אמר |  |
| 8:13.563 | רב |  |
| 8:13.564 | יהודה |  |
| 8:13.565 | אלמלא |  |
| 8:13.566 | וי"ו |  |
| 8:13.567 | שבהעלוך |  |
| 8:13.568 | נחייבו |  |
| 8:13.569 | שונאיהם |  |
| 8:13.570 | של |  |
| 8:13.571 | ישראל |  |
| 8:13.572 | כלייה | . |
| 8:13.573 | אמר |  |
| 8:13.574 | לו |  |
| 8:13.575 | רשב"י |  |
| 8:13.576 | והלא |  |
| 8:13.577 | כל |  |
| 8:13.578 | המשתף |  |
| 8:13.579 | שם |  |
| 8:13.580 | שמים |  |
| 8:13.581 | ודבר |  |
| 8:13.582 | אחר |  |
| 8:13.583 | נעקר |  |
| 8:13.584 | מן |  |
| 8:13.585 | העולם |  |
| 8:13.586 | שנ | ' |
| 8:13.587 | בלתו |  |
| 8:13.588 | לה | ' |
| 8:13.589 | לבדו |  |
| 8:13.590 | אלא |  |
| 8:13.591 | מה |  |
| 8:13.592 | ת"ל |  |
| 8:13.593 | אשר |  |
| 8:13.594 | העלוך |  |
| 8:13.595 | מלמד |  |
| 8:13.596 | שאיוו |  |
| 8:13.597 | לאלהות |  |
| 8:13.598 | הרבה | , |
| 8:13.599 | ועיין |  |
| 8:13.600 | לעיל |  |
| 8:13.601 | סי | ' |
| 8:13.602 | נ"ח | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

