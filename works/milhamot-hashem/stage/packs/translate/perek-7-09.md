# Translation pack: milhamot-hashem perek-7-09 (Milchamot ha-Shem 7:13–7:13)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-09`.

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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
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


## Milchamot ha-Shem 7:13

Hebrew (Wikisource): פו) והעיקר השלישי מס' שבילי אמונה להרחיק ממנו ית' מדת הגוף, כי זה אלוהינו האחד אשר זכרנו אינו גוף. ולא כח בגוף, ולא ישיגוהו משיגי הגוף ומקריו, כגון התנועה והעמידה, לא בעצם ולא במקרה, ועל כן הרחיקו רז"ל ממנו ית' שם החיבור והפירוד, ואמרו אין למעלה לא ישיבה ולא עמידה ולא עורף ולא עיפוי, עורף הוא פירוד, עיפוי הוא חיבור, נגזר מן ועפו בכתף פלשתים, רוצה לומר יהדפום בהדבקם בהם, ויאמר הנביא ואל מי תדמיוני ואשוה, ואלו היה גוף היה דומה לשאר הגופים וכל מה שנמצא בכתוב אשר ידמוהו בדמות הגופות כגון התנועה והעמידה כל זה דרך העברה כמו שאמרו דברה תורה כלשון בני אדם, וזה העיקר השלישי הוא האמור בכתוב כי לא ראיתם כל תמונה, רוצה לומר שלא השגתם לו ית' שום תמונה. כמו שאמרנו שאינו גוף ולא כח בגוף העכ"ל. וז"ל הרמב"ם בפ"א מהל' יסודי התורה המצוי הזה הוא אלהי העולם אדון כל הארץ, והוא המנהיג את הגלגל בכח שאין לו קץ ותכלית, בכח שאין לו הפסק וכו' וידיעת דבר זה מצות עשה שנ' אנכי ה' אלהיך. וכל המעלה על דעתו שיש שם אלוה אחר חוץ מזה, עובר בלא תעשה שנ' לא יהיה לך אלהים אחרים על פני. וכופר בעיקר, שזהו העיקר הגדול שהכל תלוי בו. אלוה זה אחד הוא, ואינו שנים, ולא יותר על שנים, אלא אחד שאין כיחודו אחד מן האחדים הנמצאים בעולם, לא אחד כמין שהוא כולל אחדים רבים ולא כאחד כגוף שהוא נחלק למחלקות ולקצוות, אלא יחוד שאין יחוד אחר בעולם כמותו. אלו היו אלוהות הרבה, היו גופים וגויות, מפני שאין הנמנין השוים במציאותן נפרדין זה מזה אלא במאורעין שיארעו לגופות ולגוית. ואלו היה היוצר גוף וגויה, היה לו קץ ותכלית, שאי אפשר להיות גוף שאין לו קץ, וכל שיש לגופו קץ ותכלית, יש לכחו קץ וסוף, ואלהינו ברוך שמו הואיל וכחו אין לו קץ ואינו פוסק שהרי הגלגל סובב תמיד, אין כחו כח בגוף. והואיל ואינו גוף, לא יארעו לו מאורעות הגופים כדי שיהיה נחלק ונפרד מאחר. לפיכך אי אפשר שיהיה אלא אחד, וידיעת דבר זה מצות עשה שנ' שמע ישראל ה' אלהינו ה' אחד. הרי מפורש בתורה ובנביאים שאין הקב"ה גוף וגויה שנ' כי ה' אלהיכם הוא אלהים שבשמים ממעל ועל הארץ מתחת. והגוף לא יהיה בשני מקומות ונאמר כי לא ראיתם כל תמונה ונאמר ואל מי תדמיוני ואשוה, ואלו היה גוף היה דומה לשאר גופים. אם כן מהו זה שכתוב בתורה ותחת רגליו, כתובים באצבע אלהים, יד ה', עיני ה', אזני ה', וכיצוא בדברים האלו? הכל לפי דעתן שלבני אדם שאין מכירין אלא הגופות, ודברה תורה כלשון בני אדם. והכל כנוים הם שנ' אם שנותי ברק חרבי וכי חרב יש לו ובחרב הוא הורג. אלא משל, והכל משל, ראיה לדבר זה שנביא אחד אומר שראה הקב"ה לבושיה כתלג חיור. ואחר ראהו חמוץ בגדים. משה רבינו עצמו ראהו על הים כגבור עושה מלחמה, ובסיני כשליח צבור עטוף. לומר שאין לו דמות וצורה. אלא הכל במראה הנבואה ובמחזה. אבל אמתת הדבר אין דעתו של אדם מבין ולא יכולה להשיגו ולחקרו, וזהו שאמר הכתוב החקר אלוה תמצא, אם עד תכלית שדי תמצא וכו'. וכיון שנתברר שאינו גוף וגוייה יתברר שלא יארעו מאורעות הגופים, לא חיבור ולא פירוד. לא מקום ולא מידה, לא עלייה ולא ירידה, ולא ימין ולא שמאל, לא פנים ולא אחור, ולא ישיבה ולא עמידה, ואינו מצוי בזמן עד שיהיה לו ראשית ואחרית ומנין שנים. ואינו משתנה, שאין לו דבר שיגרום לו שינוי, ואין לו לא מות ולא חיים כחיי הגוף, ולא סכלות ולא חכמת כחכמת האיש החכם, לא שינה ולא יקיצה, ולא כעס ולא שחוק, ולא שמחה ולא עצבות, ולא שתיקה ולא דיבור כדיבור בני אדם, וכך אמרו חכמים אין למעלה לא ישיבה ולא עמידה לא עורף ולא עיפוי, עד כאן ל' הרמב"ם ז"ל.

| id | word | punctuation after |
|---|---|---|
| 7:13.1 | פו | ) |
| 7:13.2 | והעיקר |  |
| 7:13.3 | השלישי |  |
| 7:13.4 | מס | ' |
| 7:13.5 | שבילי |  |
| 7:13.6 | אמונה |  |
| 7:13.7 | להרחיק |  |
| 7:13.8 | ממנו |  |
| 7:13.9 | ית | ' |
| 7:13.10 | מדת |  |
| 7:13.11 | הגוף | , |
| 7:13.12 | כי |  |
| 7:13.13 | זה |  |
| 7:13.14 | אלוהינו |  |
| 7:13.15 | האחד |  |
| 7:13.16 | אשר |  |
| 7:13.17 | זכרנו |  |
| 7:13.18 | אינו |  |
| 7:13.19 | גוף | . |
| 7:13.20 | ולא |  |
| 7:13.21 | כח |  |
| 7:13.22 | בגוף | , |
| 7:13.23 | ולא |  |
| 7:13.24 | ישיגוהו |  |
| 7:13.25 | משיגי |  |
| 7:13.26 | הגוף |  |
| 7:13.27 | ומקריו | , |
| 7:13.28 | כגון |  |
| 7:13.29 | התנועה |  |
| 7:13.30 | והעמידה | , |
| 7:13.31 | לא |  |
| 7:13.32 | בעצם |  |
| 7:13.33 | ולא |  |
| 7:13.34 | במקרה | , |
| 7:13.35 | ועל |  |
| 7:13.36 | כן |  |
| 7:13.37 | הרחיקו |  |
| 7:13.38 | רז"ל |  |
| 7:13.39 | ממנו |  |
| 7:13.40 | ית | ' |
| 7:13.41 | שם |  |
| 7:13.42 | החיבור |  |
| 7:13.43 | והפירוד | , |
| 7:13.44 | ואמרו |  |
| 7:13.45 | אין |  |
| 7:13.46 | למעלה |  |
| 7:13.47 | לא |  |
| 7:13.48 | ישיבה |  |
| 7:13.49 | ולא |  |
| 7:13.50 | עמידה |  |
| 7:13.51 | ולא |  |
| 7:13.52 | עורף |  |
| 7:13.53 | ולא |  |
| 7:13.54 | עיפוי | , |
| 7:13.55 | עורף |  |
| 7:13.56 | הוא |  |
| 7:13.57 | פירוד | , |
| 7:13.58 | עיפוי |  |
| 7:13.59 | הוא |  |
| 7:13.60 | חיבור | , |
| 7:13.61 | נגזר |  |
| 7:13.62 | מן |  |
| 7:13.63 | ועפו |  |
| 7:13.64 | בכתף |  |
| 7:13.65 | פלשתים | , |
| 7:13.66 | רוצה |  |
| 7:13.67 | לומר |  |
| 7:13.68 | יהדפום |  |
| 7:13.69 | בהדבקם |  |
| 7:13.70 | בהם | , |
| 7:13.71 | ויאמר |  |
| 7:13.72 | הנביא |  |
| 7:13.73 | ואל |  |
| 7:13.74 | מי |  |
| 7:13.75 | תדמיוני |  |
| 7:13.76 | ואשוה | , |
| 7:13.77 | ואלו |  |
| 7:13.78 | היה |  |
| 7:13.79 | גוף |  |
| 7:13.80 | היה |  |
| 7:13.81 | דומה |  |
| 7:13.82 | לשאר |  |
| 7:13.83 | הגופים |  |
| 7:13.84 | וכל |  |
| 7:13.85 | מה |  |
| 7:13.86 | שנמצא |  |
| 7:13.87 | בכתוב |  |
| 7:13.88 | אשר |  |
| 7:13.89 | ידמוהו |  |
| 7:13.90 | בדמות |  |
| 7:13.91 | הגופות |  |
| 7:13.92 | כגון |  |
| 7:13.93 | התנועה |  |
| 7:13.94 | והעמידה |  |
| 7:13.95 | כל |  |
| 7:13.96 | זה |  |
| 7:13.97 | דרך |  |
| 7:13.98 | העברה |  |
| 7:13.99 | כמו |  |
| 7:13.100 | שאמרו |  |
| 7:13.101 | דברה |  |
| 7:13.102 | תורה |  |
| 7:13.103 | כלשון |  |
| 7:13.104 | בני |  |
| 7:13.105 | אדם | , |
| 7:13.106 | וזה |  |
| 7:13.107 | העיקר |  |
| 7:13.108 | השלישי |  |
| 7:13.109 | הוא |  |
| 7:13.110 | האמור |  |
| 7:13.111 | בכתוב |  |
| 7:13.112 | כי |  |
| 7:13.113 | לא |  |
| 7:13.114 | ראיתם |  |
| 7:13.115 | כל |  |
| 7:13.116 | תמונה | , |
| 7:13.117 | רוצה |  |
| 7:13.118 | לומר |  |
| 7:13.119 | שלא |  |
| 7:13.120 | השגתם |  |
| 7:13.121 | לו |  |
| 7:13.122 | ית | ' |
| 7:13.123 | שום |  |
| 7:13.124 | תמונה | . |
| 7:13.125 | כמו |  |
| 7:13.126 | שאמרנו |  |
| 7:13.127 | שאינו |  |
| 7:13.128 | גוף |  |
| 7:13.129 | ולא |  |
| 7:13.130 | כח |  |
| 7:13.131 | בגוף |  |
| 7:13.132 | העכ"ל | . |
| 7:13.133 | וז"ל |  |
| 7:13.134 | הרמב"ם |  |
| 7:13.135 | בפ"א |  |
| 7:13.136 | מהל | ' |
| 7:13.137 | יסודי |  |
| 7:13.138 | התורה |  |
| 7:13.139 | המצוי |  |
| 7:13.140 | הזה |  |
| 7:13.141 | הוא |  |
| 7:13.142 | אלהי |  |
| 7:13.143 | העולם |  |
| 7:13.144 | אדון |  |
| 7:13.145 | כל |  |
| 7:13.146 | הארץ | , |
| 7:13.147 | והוא |  |
| 7:13.148 | המנהיג |  |
| 7:13.149 | את |  |
| 7:13.150 | הגלגל |  |
| 7:13.151 | בכח |  |
| 7:13.152 | שאין |  |
| 7:13.153 | לו |  |
| 7:13.154 | קץ |  |
| 7:13.155 | ותכלית | , |
| 7:13.156 | בכח |  |
| 7:13.157 | שאין |  |
| 7:13.158 | לו |  |
| 7:13.159 | הפסק |  |
| 7:13.160 | וכו | ' |
| 7:13.161 | וידיעת |  |
| 7:13.162 | דבר |  |
| 7:13.163 | זה |  |
| 7:13.164 | מצות |  |
| 7:13.165 | עשה |  |
| 7:13.166 | שנ | ' |
| 7:13.167 | אנכי |  |
| 7:13.168 | ה | ' |
| 7:13.169 | אלהיך | . |
| 7:13.170 | וכל |  |
| 7:13.171 | המעלה |  |
| 7:13.172 | על |  |
| 7:13.173 | דעתו |  |
| 7:13.174 | שיש |  |
| 7:13.175 | שם |  |
| 7:13.176 | אלוה |  |
| 7:13.177 | אחר |  |
| 7:13.178 | חוץ |  |
| 7:13.179 | מזה | , |
| 7:13.180 | עובר |  |
| 7:13.181 | בלא |  |
| 7:13.182 | תעשה |  |
| 7:13.183 | שנ | ' |
| 7:13.184 | לא |  |
| 7:13.185 | יהיה |  |
| 7:13.186 | לך |  |
| 7:13.187 | אלהים |  |
| 7:13.188 | אחרים |  |
| 7:13.189 | על |  |
| 7:13.190 | פני | . |
| 7:13.191 | וכופר |  |
| 7:13.192 | בעיקר | , |
| 7:13.193 | שזהו |  |
| 7:13.194 | העיקר |  |
| 7:13.195 | הגדול |  |
| 7:13.196 | שהכל |  |
| 7:13.197 | תלוי |  |
| 7:13.198 | בו | . |
| 7:13.199 | אלוה |  |
| 7:13.200 | זה |  |
| 7:13.201 | אחד |  |
| 7:13.202 | הוא | , |
| 7:13.203 | ואינו |  |
| 7:13.204 | שנים | , |
| 7:13.205 | ולא |  |
| 7:13.206 | יותר |  |
| 7:13.207 | על |  |
| 7:13.208 | שנים | , |
| 7:13.209 | אלא |  |
| 7:13.210 | אחד |  |
| 7:13.211 | שאין |  |
| 7:13.212 | כיחודו |  |
| 7:13.213 | אחד |  |
| 7:13.214 | מן |  |
| 7:13.215 | האחדים |  |
| 7:13.216 | הנמצאים |  |
| 7:13.217 | בעולם | , |
| 7:13.218 | לא |  |
| 7:13.219 | אחד |  |
| 7:13.220 | כמין |  |
| 7:13.221 | שהוא |  |
| 7:13.222 | כולל |  |
| 7:13.223 | אחדים |  |
| 7:13.224 | רבים |  |
| 7:13.225 | ולא |  |
| 7:13.226 | כאחד |  |
| 7:13.227 | כגוף |  |
| 7:13.228 | שהוא |  |
| 7:13.229 | נחלק |  |
| 7:13.230 | למחלקות |  |
| 7:13.231 | ולקצוות | , |
| 7:13.232 | אלא |  |
| 7:13.233 | יחוד |  |
| 7:13.234 | שאין |  |
| 7:13.235 | יחוד |  |
| 7:13.236 | אחר |  |
| 7:13.237 | בעולם |  |
| 7:13.238 | כמותו | . |
| 7:13.239 | אלו |  |
| 7:13.240 | היו |  |
| 7:13.241 | אלוהות |  |
| 7:13.242 | הרבה | , |
| 7:13.243 | היו |  |
| 7:13.244 | גופים |  |
| 7:13.245 | וגויות | , |
| 7:13.246 | מפני |  |
| 7:13.247 | שאין |  |
| 7:13.248 | הנמנין |  |
| 7:13.249 | השוים |  |
| 7:13.250 | במציאותן |  |
| 7:13.251 | נפרדין |  |
| 7:13.252 | זה |  |
| 7:13.253 | מזה |  |
| 7:13.254 | אלא |  |
| 7:13.255 | במאורעין |  |
| 7:13.256 | שיארעו |  |
| 7:13.257 | לגופות |  |
| 7:13.258 | ולגוית | . |
| 7:13.259 | ואלו |  |
| 7:13.260 | היה |  |
| 7:13.261 | היוצר |  |
| 7:13.262 | גוף |  |
| 7:13.263 | וגויה | , |
| 7:13.264 | היה |  |
| 7:13.265 | לו |  |
| 7:13.266 | קץ |  |
| 7:13.267 | ותכלית | , |
| 7:13.268 | שאי |  |
| 7:13.269 | אפשר |  |
| 7:13.270 | להיות |  |
| 7:13.271 | גוף |  |
| 7:13.272 | שאין |  |
| 7:13.273 | לו |  |
| 7:13.274 | קץ | , |
| 7:13.275 | וכל |  |
| 7:13.276 | שיש |  |
| 7:13.277 | לגופו |  |
| 7:13.278 | קץ |  |
| 7:13.279 | ותכלית | , |
| 7:13.280 | יש |  |
| 7:13.281 | לכחו |  |
| 7:13.282 | קץ |  |
| 7:13.283 | וסוף | , |
| 7:13.284 | ואלהינו |  |
| 7:13.285 | ברוך |  |
| 7:13.286 | שמו |  |
| 7:13.287 | הואיל |  |
| 7:13.288 | וכחו |  |
| 7:13.289 | אין |  |
| 7:13.290 | לו |  |
| 7:13.291 | קץ |  |
| 7:13.292 | ואינו |  |
| 7:13.293 | פוסק |  |
| 7:13.294 | שהרי |  |
| 7:13.295 | הגלגל |  |
| 7:13.296 | סובב |  |
| 7:13.297 | תמיד | , |
| 7:13.298 | אין |  |
| 7:13.299 | כחו |  |
| 7:13.300 | כח |  |
| 7:13.301 | בגוף | . |
| 7:13.302 | והואיל |  |
| 7:13.303 | ואינו |  |
| 7:13.304 | גוף | , |
| 7:13.305 | לא |  |
| 7:13.306 | יארעו |  |
| 7:13.307 | לו |  |
| 7:13.308 | מאורעות |  |
| 7:13.309 | הגופים |  |
| 7:13.310 | כדי |  |
| 7:13.311 | שיהיה |  |
| 7:13.312 | נחלק |  |
| 7:13.313 | ונפרד |  |
| 7:13.314 | מאחר | . |
| 7:13.315 | לפיכך |  |
| 7:13.316 | אי |  |
| 7:13.317 | אפשר |  |
| 7:13.318 | שיהיה |  |
| 7:13.319 | אלא |  |
| 7:13.320 | אחד | , |
| 7:13.321 | וידיעת |  |
| 7:13.322 | דבר |  |
| 7:13.323 | זה |  |
| 7:13.324 | מצות |  |
| 7:13.325 | עשה |  |
| 7:13.326 | שנ | ' |
| 7:13.327 | שמע |  |
| 7:13.328 | ישראל |  |
| 7:13.329 | ה | ' |
| 7:13.330 | אלהינו |  |
| 7:13.331 | ה | ' |
| 7:13.332 | אחד | . |
| 7:13.333 | הרי |  |
| 7:13.334 | מפורש |  |
| 7:13.335 | בתורה |  |
| 7:13.336 | ובנביאים |  |
| 7:13.337 | שאין |  |
| 7:13.338 | הקב"ה |  |
| 7:13.339 | גוף |  |
| 7:13.340 | וגויה |  |
| 7:13.341 | שנ | ' |
| 7:13.342 | כי |  |
| 7:13.343 | ה | ' |
| 7:13.344 | אלהיכם |  |
| 7:13.345 | הוא |  |
| 7:13.346 | אלהים |  |
| 7:13.347 | שבשמים |  |
| 7:13.348 | ממעל |  |
| 7:13.349 | ועל |  |
| 7:13.350 | הארץ |  |
| 7:13.351 | מתחת | . |
| 7:13.352 | והגוף |  |
| 7:13.353 | לא |  |
| 7:13.354 | יהיה |  |
| 7:13.355 | בשני |  |
| 7:13.356 | מקומות |  |
| 7:13.357 | ונאמר |  |
| 7:13.358 | כי |  |
| 7:13.359 | לא |  |
| 7:13.360 | ראיתם |  |
| 7:13.361 | כל |  |
| 7:13.362 | תמונה |  |
| 7:13.363 | ונאמר |  |
| 7:13.364 | ואל |  |
| 7:13.365 | מי |  |
| 7:13.366 | תדמיוני |  |
| 7:13.367 | ואשוה | , |
| 7:13.368 | ואלו |  |
| 7:13.369 | היה |  |
| 7:13.370 | גוף |  |
| 7:13.371 | היה |  |
| 7:13.372 | דומה |  |
| 7:13.373 | לשאר |  |
| 7:13.374 | גופים | . |
| 7:13.375 | אם |  |
| 7:13.376 | כן |  |
| 7:13.377 | מהו |  |
| 7:13.378 | זה |  |
| 7:13.379 | שכתוב |  |
| 7:13.380 | בתורה |  |
| 7:13.381 | ותחת |  |
| 7:13.382 | רגליו | , |
| 7:13.383 | כתובים |  |
| 7:13.384 | באצבע |  |
| 7:13.385 | אלהים | , |
| 7:13.386 | יד |  |
| 7:13.387 | ה | ', |
| 7:13.388 | עיני |  |
| 7:13.389 | ה | ', |
| 7:13.390 | אזני |  |
| 7:13.391 | ה | ', |
| 7:13.392 | וכיצוא |  |
| 7:13.393 | בדברים |  |
| 7:13.394 | האלו | ? |
| 7:13.395 | הכל |  |
| 7:13.396 | לפי |  |
| 7:13.397 | דעתן |  |
| 7:13.398 | שלבני |  |
| 7:13.399 | אדם |  |
| 7:13.400 | שאין |  |
| 7:13.401 | מכירין |  |
| 7:13.402 | אלא |  |
| 7:13.403 | הגופות | , |
| 7:13.404 | ודברה |  |
| 7:13.405 | תורה |  |
| 7:13.406 | כלשון |  |
| 7:13.407 | בני |  |
| 7:13.408 | אדם | . |
| 7:13.409 | והכל |  |
| 7:13.410 | כנוים |  |
| 7:13.411 | הם |  |
| 7:13.412 | שנ | ' |
| 7:13.413 | אם |  |
| 7:13.414 | שנותי |  |
| 7:13.415 | ברק |  |
| 7:13.416 | חרבי |  |
| 7:13.417 | וכי |  |
| 7:13.418 | חרב |  |
| 7:13.419 | יש |  |
| 7:13.420 | לו |  |
| 7:13.421 | ובחרב |  |
| 7:13.422 | הוא |  |
| 7:13.423 | הורג | . |
| 7:13.424 | אלא |  |
| 7:13.425 | משל | , |
| 7:13.426 | והכל |  |
| 7:13.427 | משל | , |
| 7:13.428 | ראיה |  |
| 7:13.429 | לדבר |  |
| 7:13.430 | זה |  |
| 7:13.431 | שנביא |  |
| 7:13.432 | אחד |  |
| 7:13.433 | אומר |  |
| 7:13.434 | שראה |  |
| 7:13.435 | הקב"ה |  |
| 7:13.436 | לבושיה |  |
| 7:13.437 | כתלג |  |
| 7:13.438 | חיור | . |
| 7:13.439 | ואחר |  |
| 7:13.440 | ראהו |  |
| 7:13.441 | חמוץ |  |
| 7:13.442 | בגדים | . |
| 7:13.443 | משה |  |
| 7:13.444 | רבינו |  |
| 7:13.445 | עצמו |  |
| 7:13.446 | ראהו |  |
| 7:13.447 | על |  |
| 7:13.448 | הים |  |
| 7:13.449 | כגבור |  |
| 7:13.450 | עושה |  |
| 7:13.451 | מלחמה | , |
| 7:13.452 | ובסיני |  |
| 7:13.453 | כשליח |  |
| 7:13.454 | צבור |  |
| 7:13.455 | עטוף | . |
| 7:13.456 | לומר |  |
| 7:13.457 | שאין |  |
| 7:13.458 | לו |  |
| 7:13.459 | דמות |  |
| 7:13.460 | וצורה | . |
| 7:13.461 | אלא |  |
| 7:13.462 | הכל |  |
| 7:13.463 | במראה |  |
| 7:13.464 | הנבואה |  |
| 7:13.465 | ובמחזה | . |
| 7:13.466 | אבל |  |
| 7:13.467 | אמתת |  |
| 7:13.468 | הדבר |  |
| 7:13.469 | אין |  |
| 7:13.470 | דעתו |  |
| 7:13.471 | של |  |
| 7:13.472 | אדם |  |
| 7:13.473 | מבין |  |
| 7:13.474 | ולא |  |
| 7:13.475 | יכולה |  |
| 7:13.476 | להשיגו |  |
| 7:13.477 | ולחקרו | , |
| 7:13.478 | וזהו |  |
| 7:13.479 | שאמר |  |
| 7:13.480 | הכתוב |  |
| 7:13.481 | החקר |  |
| 7:13.482 | אלוה |  |
| 7:13.483 | תמצא | , |
| 7:13.484 | אם |  |
| 7:13.485 | עד |  |
| 7:13.486 | תכלית |  |
| 7:13.487 | שדי |  |
| 7:13.488 | תמצא |  |
| 7:13.489 | וכו | '. |
| 7:13.490 | וכיון |  |
| 7:13.491 | שנתברר |  |
| 7:13.492 | שאינו |  |
| 7:13.493 | גוף |  |
| 7:13.494 | וגוייה |  |
| 7:13.495 | יתברר |  |
| 7:13.496 | שלא |  |
| 7:13.497 | יארעו |  |
| 7:13.498 | מאורעות |  |
| 7:13.499 | הגופים | , |
| 7:13.500 | לא |  |
| 7:13.501 | חיבור |  |
| 7:13.502 | ולא |  |
| 7:13.503 | פירוד | . |
| 7:13.504 | לא |  |
| 7:13.505 | מקום |  |
| 7:13.506 | ולא |  |
| 7:13.507 | מידה | , |
| 7:13.508 | לא |  |
| 7:13.509 | עלייה |  |
| 7:13.510 | ולא |  |
| 7:13.511 | ירידה | , |
| 7:13.512 | ולא |  |
| 7:13.513 | ימין |  |
| 7:13.514 | ולא |  |
| 7:13.515 | שמאל | , |
| 7:13.516 | לא |  |
| 7:13.517 | פנים |  |
| 7:13.518 | ולא |  |
| 7:13.519 | אחור | , |
| 7:13.520 | ולא |  |
| 7:13.521 | ישיבה |  |
| 7:13.522 | ולא |  |
| 7:13.523 | עמידה | , |
| 7:13.524 | ואינו |  |
| 7:13.525 | מצוי |  |
| 7:13.526 | בזמן |  |
| 7:13.527 | עד |  |
| 7:13.528 | שיהיה |  |
| 7:13.529 | לו |  |
| 7:13.530 | ראשית |  |
| 7:13.531 | ואחרית |  |
| 7:13.532 | ומנין |  |
| 7:13.533 | שנים | . |
| 7:13.534 | ואינו |  |
| 7:13.535 | משתנה | , |
| 7:13.536 | שאין |  |
| 7:13.537 | לו |  |
| 7:13.538 | דבר |  |
| 7:13.539 | שיגרום |  |
| 7:13.540 | לו |  |
| 7:13.541 | שינוי | , |
| 7:13.542 | ואין |  |
| 7:13.543 | לו |  |
| 7:13.544 | לא |  |
| 7:13.545 | מות |  |
| 7:13.546 | ולא |  |
| 7:13.547 | חיים |  |
| 7:13.548 | כחיי |  |
| 7:13.549 | הגוף | , |
| 7:13.550 | ולא |  |
| 7:13.551 | סכלות |  |
| 7:13.552 | ולא |  |
| 7:13.553 | חכמת |  |
| 7:13.554 | כחכמת |  |
| 7:13.555 | האיש |  |
| 7:13.556 | החכם | , |
| 7:13.557 | לא |  |
| 7:13.558 | שינה |  |
| 7:13.559 | ולא |  |
| 7:13.560 | יקיצה | , |
| 7:13.561 | ולא |  |
| 7:13.562 | כעס |  |
| 7:13.563 | ולא |  |
| 7:13.564 | שחוק | , |
| 7:13.565 | ולא |  |
| 7:13.566 | שמחה |  |
| 7:13.567 | ולא |  |
| 7:13.568 | עצבות | , |
| 7:13.569 | ולא |  |
| 7:13.570 | שתיקה |  |
| 7:13.571 | ולא |  |
| 7:13.572 | דיבור |  |
| 7:13.573 | כדיבור |  |
| 7:13.574 | בני |  |
| 7:13.575 | אדם | , |
| 7:13.576 | וכך |  |
| 7:13.577 | אמרו |  |
| 7:13.578 | חכמים |  |
| 7:13.579 | אין |  |
| 7:13.580 | למעלה |  |
| 7:13.581 | לא |  |
| 7:13.582 | ישיבה |  |
| 7:13.583 | ולא |  |
| 7:13.584 | עמידה |  |
| 7:13.585 | לא |  |
| 7:13.586 | עורף |  |
| 7:13.587 | ולא |  |
| 7:13.588 | עיפוי | , |
| 7:13.589 | עד |  |
| 7:13.590 | כאן |  |
| 7:13.591 | ל | ' |
| 7:13.592 | הרמב"ם |  |
| 7:13.593 | ז"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

