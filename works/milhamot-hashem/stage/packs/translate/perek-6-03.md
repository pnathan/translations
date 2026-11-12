# Translation pack: milhamot-hashem perek-6-03 (Milchamot ha-Shem 6:4–6:7)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-03`.

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
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
- **אין לנו להחמיר** → "we may not be stringent".
- **הגיע חתור** (infinitive after an adjective of motion): the grammar as
  written; the reviewer's reading may go in `readings`.
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


## Milchamot ha-Shem 6:4

Hebrew (Wikisource): ומדברי הר' הסמ"ג הרוחנו שעד ימיו היו כל ישראל תקועים באמונת יחוד השי"ת שהוא אחד ומיוחד מכל שאר האחדים כד' הרס"ג. והיא דעת תורתנו הקדושה שבכתב ושבעל פה שקבלו רז"ל גמרא ובמדרשים הידועים ומפורסמים לרז"ל בלי שום ספק.

| id | word | punctuation after |
|---|---|---|
| 6:4.1 | ומדברי |  |
| 6:4.2 | הר | ' |
| 6:4.3 | הסמ"ג |  |
| 6:4.4 | הרוחנו |  |
| 6:4.5 | שעד |  |
| 6:4.6 | ימיו |  |
| 6:4.7 | היו |  |
| 6:4.8 | כל |  |
| 6:4.9 | ישראל |  |
| 6:4.10 | תקועים |  |
| 6:4.11 | באמונת |  |
| 6:4.12 | יחוד |  |
| 6:4.13 | השי"ת |  |
| 6:4.14 | שהוא |  |
| 6:4.15 | אחד |  |
| 6:4.16 | ומיוחד |  |
| 6:4.17 | מכל |  |
| 6:4.18 | שאר |  |
| 6:4.19 | האחדים |  |
| 6:4.20 | כד | ' |
| 6:4.21 | הרס"ג | . |
| 6:4.22 | והיא |  |
| 6:4.23 | דעת |  |
| 6:4.24 | תורתנו |  |
| 6:4.25 | הקדושה |  |
| 6:4.26 | שבכתב |  |
| 6:4.27 | ושבעל |  |
| 6:4.28 | פה |  |
| 6:4.29 | שקבלו |  |
| 6:4.30 | רז"ל |  |
| 6:4.31 | גמרא |  |
| 6:4.32 | ובמדרשים |  |
| 6:4.33 | הידועים |  |
| 6:4.34 | ומפורסמים |  |
| 6:4.35 | לרז"ל |  |
| 6:4.36 | בלי |  |
| 6:4.37 | שום |  |
| 6:4.38 | ספק | . |


## Milchamot ha-Shem 6:5

Hebrew (Wikisource): ושלא כדעת הקבלה החדשה אשר לא נודע מבטן מי יצאה, ומיד בן נכר הגיע ליד ישראל כמו שכתב הר' חיים יוסף דוד אזולאי בספרו שם הגדולים ויובא לקמן והיא כמעט מינות גמורה, כי רק בלשונם יכזבו לו יתברך שהוא אחד, ובלבם יחשבו כמה עילות משתלשלים זה אחר זה וכל עילה נקרא בשם הקב"ה.

| id | word | punctuation after |
|---|---|---|
| 6:5.1 | ושלא |  |
| 6:5.2 | כדעת |  |
| 6:5.3 | הקבלה |  |
| 6:5.4 | החדשה |  |
| 6:5.5 | אשר |  |
| 6:5.6 | לא |  |
| 6:5.7 | נודע |  |
| 6:5.8 | מבטן |  |
| 6:5.9 | מי |  |
| 6:5.10 | יצאה | , |
| 6:5.11 | ומיד |  |
| 6:5.12 | בן |  |
| 6:5.13 | נכר |  |
| 6:5.14 | הגיע |  |
| 6:5.15 | ליד |  |
| 6:5.16 | ישראל |  |
| 6:5.17 | כמו |  |
| 6:5.18 | שכתב |  |
| 6:5.19 | הר | ' |
| 6:5.20 | חיים |  |
| 6:5.21 | יוסף |  |
| 6:5.22 | דוד |  |
| 6:5.23 | אזולאי |  |
| 6:5.24 | בספרו |  |
| 6:5.25 | שם |  |
| 6:5.26 | הגדולים |  |
| 6:5.27 | ויובא |  |
| 6:5.28 | לקמן |  |
| 6:5.29 | והיא |  |
| 6:5.30 | כמעט |  |
| 6:5.31 | מינות |  |
| 6:5.32 | גמורה | , |
| 6:5.33 | כי |  |
| 6:5.34 | רק |  |
| 6:5.35 | בלשונם |  |
| 6:5.36 | יכזבו |  |
| 6:5.37 | לו |  |
| 6:5.38 | יתברך |  |
| 6:5.39 | שהוא |  |
| 6:5.40 | אחד | , |
| 6:5.41 | ובלבם |  |
| 6:5.42 | יחשבו |  |
| 6:5.43 | כמה |  |
| 6:5.44 | עילות |  |
| 6:5.45 | משתלשלים |  |
| 6:5.46 | זה |  |
| 6:5.47 | אחר |  |
| 6:5.48 | זה |  |
| 6:5.49 | וכל |  |
| 6:5.50 | עילה |  |
| 6:5.51 | נקרא |  |
| 6:5.52 | בשם |  |
| 6:5.53 | הקב"ה | . |


## Milchamot ha-Shem 6:6

Hebrew (Wikisource): סג) והר"מ בעל הגדרים בשורש אמונה כתב וז"ל, אמונה הוא הענין המצוייר בנפש כשיאמינו בו כך כמו שצויר, יהיה אמת מה שצייר או בלתי אמת, כי היא אמונה אצלו רצויה אצל המאמין, והאמונה האמתית היא ההאמנה במה שצייר שהוא חוץ לשכל כפי מה שהו מצויר בשכל, כי ההאמנה אינה ענין הנאמר בפה, כי אין אמונה אלא אחר ציור. ואם יהיה עם זאת האמונה שאי אפשר חילוף זאת האמונה בשום פנים. אז תהיה האמונה ההיא אמתית וקיימת, כי תסכים למציאות והאמות כאשר יהיה בצודק, והוא לא יאמין במקבילו, היה אמות אמרתי ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:6.1 | סג | ) |
| 6:6.2 | והר"מ |  |
| 6:6.3 | בעל |  |
| 6:6.4 | הגדרים |  |
| 6:6.5 | בשורש |  |
| 6:6.6 | אמונה |  |
| 6:6.7 | כתב |  |
| 6:6.8 | וז"ל | , |
| 6:6.9 | אמונה |  |
| 6:6.10 | הוא |  |
| 6:6.11 | הענין |  |
| 6:6.12 | המצוייר |  |
| 6:6.13 | בנפש |  |
| 6:6.14 | כשיאמינו |  |
| 6:6.15 | בו |  |
| 6:6.16 | כך |  |
| 6:6.17 | כמו |  |
| 6:6.18 | שצויר | , |
| 6:6.19 | יהיה |  |
| 6:6.20 | אמת |  |
| 6:6.21 | מה |  |
| 6:6.22 | שצייר |  |
| 6:6.23 | או |  |
| 6:6.24 | בלתי |  |
| 6:6.25 | אמת | , |
| 6:6.26 | כי |  |
| 6:6.27 | היא |  |
| 6:6.28 | אמונה |  |
| 6:6.29 | אצלו |  |
| 6:6.30 | רצויה |  |
| 6:6.31 | אצל |  |
| 6:6.32 | המאמין | , |
| 6:6.33 | והאמונה |  |
| 6:6.34 | האמתית |  |
| 6:6.35 | היא |  |
| 6:6.36 | ההאמנה |  |
| 6:6.37 | במה |  |
| 6:6.38 | שצייר |  |
| 6:6.39 | שהוא |  |
| 6:6.40 | חוץ |  |
| 6:6.41 | לשכל |  |
| 6:6.42 | כפי |  |
| 6:6.43 | מה |  |
| 6:6.44 | שהו |  |
| 6:6.45 | מצויר |  |
| 6:6.46 | בשכל | , |
| 6:6.47 | כי |  |
| 6:6.48 | ההאמנה |  |
| 6:6.49 | אינה |  |
| 6:6.50 | ענין |  |
| 6:6.51 | הנאמר |  |
| 6:6.52 | בפה | , |
| 6:6.53 | כי |  |
| 6:6.54 | אין |  |
| 6:6.55 | אמונה |  |
| 6:6.56 | אלא |  |
| 6:6.57 | אחר |  |
| 6:6.58 | ציור | . |
| 6:6.59 | ואם |  |
| 6:6.60 | יהיה |  |
| 6:6.61 | עם |  |
| 6:6.62 | זאת |  |
| 6:6.63 | האמונה |  |
| 6:6.64 | שאי |  |
| 6:6.65 | אפשר |  |
| 6:6.66 | חילוף |  |
| 6:6.67 | זאת |  |
| 6:6.68 | האמונה |  |
| 6:6.69 | בשום |  |
| 6:6.70 | פנים | . |
| 6:6.71 | אז |  |
| 6:6.72 | תהיה |  |
| 6:6.73 | האמונה |  |
| 6:6.74 | ההיא |  |
| 6:6.75 | אמתית |  |
| 6:6.76 | וקיימת | , |
| 6:6.77 | כי |  |
| 6:6.78 | תסכים |  |
| 6:6.79 | למציאות |  |
| 6:6.80 | והאמות |  |
| 6:6.81 | כאשר |  |
| 6:6.82 | יהיה |  |
| 6:6.83 | בצודק | , |
| 6:6.84 | והוא |  |
| 6:6.85 | לא |  |
| 6:6.86 | יאמין |  |
| 6:6.87 | במקבילו | , |
| 6:6.88 | היה |  |
| 6:6.89 | אמות |  |
| 6:6.90 | אמרתי |  |
| 6:6.91 | ע"כ | . |


## Milchamot ha-Shem 6:7

Hebrew (Wikisource): ונכתבו בצדו מהמפרש וז"ל אמונה הוא הענין המצויר בנפש, לאפוקי אמונת הסותרים שאין לה ציור בנפש, לכן לא תכנס בגדר אמונה. והאמונה האמתית וכו' זה הגדר לקוח מדברי הרב המורה, ושם נאמר שהאמונה האמתית לה שם בשלושה מסכימים בקשר הסכמה אחת, שהם הציור המחשבי, ואמונת הלב, והמציאות. כמו שמצייר במחשבתו צורת הארי ומאמין בלבו שהוא נמצא כן במציאות חוץ לנפש. ואחר שהנפש לא תתרצה באמונתה כי אם בהסכמת השכל, לכן עלינו זאת המצוה לחזק אמונת ודעות התורה במחקר השכלי והסכמת החכמה עם התורה, כי בהצמדם יחד תחזק הנפש באמונתה. והיתה האמת אזור מתנה, והאמונה אזור חלציה. ויחדיו יהיו תמים, בל יזוזו מלבו כל הימים. וכן אמרו רז"ל ושאינו יודע לשאול את פתח לו וכו' וכן הורו הר' בעל החובות וה' בעל בינה לעתים וסמכה אקרא דכ' דרך אמונה בחרתי משפטיך שויתי, פי' האמונה והדעות האמנתי אחר בחירת השכל באמתתן, משפטיך שויתי, אך בדברים הנוגעים במשפט וחוקים שויתי הבלתי מבוררים בשכל עם המבוררים וקצת חכמים בעיניהם יחשבו שהאמונה הנבחרת היא המתנגדת לשכל, ונמשכו בזה אחר דעת בעלי השלישים. ולא כן חלק יעקב, כי אמונתינו לא יסור אותם השכל, ככתוב אל תדמה ישראל אל עמים לא עלו מתוך ההפכה, וריב להם עם הדעת (אסף חלק א' מאמר א') וכן יאמר הכוזרי מאמר ו' סי' ל"א חלילה לאל מן השקר ושיבא בתורה מה שהשכל מרחיק אותו וישימהו שקר, ומי חכם יבין דעת, כי רובי מצות התורה לא ניתנו לדרוש מפי התורה שבכתב כ"א כפי ביאור חכמים ז"ל בתורה שבעל פה.

| id | word | punctuation after |
|---|---|---|
| 6:7.1 | ונכתבו |  |
| 6:7.2 | בצדו |  |
| 6:7.3 | מהמפרש |  |
| 6:7.4 | וז"ל |  |
| 6:7.5 | אמונה |  |
| 6:7.6 | הוא |  |
| 6:7.7 | הענין |  |
| 6:7.8 | המצויר |  |
| 6:7.9 | בנפש | , |
| 6:7.10 | לאפוקי |  |
| 6:7.11 | אמונת |  |
| 6:7.12 | הסותרים |  |
| 6:7.13 | שאין |  |
| 6:7.14 | לה |  |
| 6:7.15 | ציור |  |
| 6:7.16 | בנפש | , |
| 6:7.17 | לכן |  |
| 6:7.18 | לא |  |
| 6:7.19 | תכנס |  |
| 6:7.20 | בגדר |  |
| 6:7.21 | אמונה | . |
| 6:7.22 | והאמונה |  |
| 6:7.23 | האמתית |  |
| 6:7.24 | וכו | ' |
| 6:7.25 | זה |  |
| 6:7.26 | הגדר |  |
| 6:7.27 | לקוח |  |
| 6:7.28 | מדברי |  |
| 6:7.29 | הרב |  |
| 6:7.30 | המורה | , |
| 6:7.31 | ושם |  |
| 6:7.32 | נאמר |  |
| 6:7.33 | שהאמונה |  |
| 6:7.34 | האמתית |  |
| 6:7.35 | לה |  |
| 6:7.36 | שם |  |
| 6:7.37 | בשלושה |  |
| 6:7.38 | מסכימים |  |
| 6:7.39 | בקשר |  |
| 6:7.40 | הסכמה |  |
| 6:7.41 | אחת | , |
| 6:7.42 | שהם |  |
| 6:7.43 | הציור |  |
| 6:7.44 | המחשבי | , |
| 6:7.45 | ואמונת |  |
| 6:7.46 | הלב | , |
| 6:7.47 | והמציאות | . |
| 6:7.48 | כמו |  |
| 6:7.49 | שמצייר |  |
| 6:7.50 | במחשבתו |  |
| 6:7.51 | צורת |  |
| 6:7.52 | הארי |  |
| 6:7.53 | ומאמין |  |
| 6:7.54 | בלבו |  |
| 6:7.55 | שהוא |  |
| 6:7.56 | נמצא |  |
| 6:7.57 | כן |  |
| 6:7.58 | במציאות |  |
| 6:7.59 | חוץ |  |
| 6:7.60 | לנפש | . |
| 6:7.61 | ואחר |  |
| 6:7.62 | שהנפש |  |
| 6:7.63 | לא |  |
| 6:7.64 | תתרצה |  |
| 6:7.65 | באמונתה |  |
| 6:7.66 | כי |  |
| 6:7.67 | אם |  |
| 6:7.68 | בהסכמת |  |
| 6:7.69 | השכל | , |
| 6:7.70 | לכן |  |
| 6:7.71 | עלינו |  |
| 6:7.72 | זאת |  |
| 6:7.73 | המצוה |  |
| 6:7.74 | לחזק |  |
| 6:7.75 | אמונת |  |
| 6:7.76 | ודעות |  |
| 6:7.77 | התורה |  |
| 6:7.78 | במחקר |  |
| 6:7.79 | השכלי |  |
| 6:7.80 | והסכמת |  |
| 6:7.81 | החכמה |  |
| 6:7.82 | עם |  |
| 6:7.83 | התורה | , |
| 6:7.84 | כי |  |
| 6:7.85 | בהצמדם |  |
| 6:7.86 | יחד |  |
| 6:7.87 | תחזק |  |
| 6:7.88 | הנפש |  |
| 6:7.89 | באמונתה | . |
| 6:7.90 | והיתה |  |
| 6:7.91 | האמת |  |
| 6:7.92 | אזור |  |
| 6:7.93 | מתנה | , |
| 6:7.94 | והאמונה |  |
| 6:7.95 | אזור |  |
| 6:7.96 | חלציה | . |
| 6:7.97 | ויחדיו |  |
| 6:7.98 | יהיו |  |
| 6:7.99 | תמים | , |
| 6:7.100 | בל |  |
| 6:7.101 | יזוזו |  |
| 6:7.102 | מלבו |  |
| 6:7.103 | כל |  |
| 6:7.104 | הימים | . |
| 6:7.105 | וכן |  |
| 6:7.106 | אמרו |  |
| 6:7.107 | רז"ל |  |
| 6:7.108 | ושאינו |  |
| 6:7.109 | יודע |  |
| 6:7.110 | לשאול |  |
| 6:7.111 | את |  |
| 6:7.112 | פתח |  |
| 6:7.113 | לו |  |
| 6:7.114 | וכו | ' |
| 6:7.115 | וכן |  |
| 6:7.116 | הורו |  |
| 6:7.117 | הר | ' |
| 6:7.118 | בעל |  |
| 6:7.119 | החובות |  |
| 6:7.120 | וה | ' |
| 6:7.121 | בעל |  |
| 6:7.122 | בינה |  |
| 6:7.123 | לעתים |  |
| 6:7.124 | וסמכה |  |
| 6:7.125 | אקרא |  |
| 6:7.126 | דכ | ' |
| 6:7.127 | דרך |  |
| 6:7.128 | אמונה |  |
| 6:7.129 | בחרתי |  |
| 6:7.130 | משפטיך |  |
| 6:7.131 | שויתי | , |
| 6:7.132 | פי | ' |
| 6:7.133 | האמונה |  |
| 6:7.134 | והדעות |  |
| 6:7.135 | האמנתי |  |
| 6:7.136 | אחר |  |
| 6:7.137 | בחירת |  |
| 6:7.138 | השכל |  |
| 6:7.139 | באמתתן | , |
| 6:7.140 | משפטיך |  |
| 6:7.141 | שויתי | , |
| 6:7.142 | אך |  |
| 6:7.143 | בדברים |  |
| 6:7.144 | הנוגעים |  |
| 6:7.145 | במשפט |  |
| 6:7.146 | וחוקים |  |
| 6:7.147 | שויתי |  |
| 6:7.148 | הבלתי |  |
| 6:7.149 | מבוררים |  |
| 6:7.150 | בשכל |  |
| 6:7.151 | עם |  |
| 6:7.152 | המבוררים |  |
| 6:7.153 | וקצת |  |
| 6:7.154 | חכמים |  |
| 6:7.155 | בעיניהם |  |
| 6:7.156 | יחשבו |  |
| 6:7.157 | שהאמונה |  |
| 6:7.158 | הנבחרת |  |
| 6:7.159 | היא |  |
| 6:7.160 | המתנגדת |  |
| 6:7.161 | לשכל | , |
| 6:7.162 | ונמשכו |  |
| 6:7.163 | בזה |  |
| 6:7.164 | אחר |  |
| 6:7.165 | דעת |  |
| 6:7.166 | בעלי |  |
| 6:7.167 | השלישים | . |
| 6:7.168 | ולא |  |
| 6:7.169 | כן |  |
| 6:7.170 | חלק |  |
| 6:7.171 | יעקב | , |
| 6:7.172 | כי |  |
| 6:7.173 | אמונתינו |  |
| 6:7.174 | לא |  |
| 6:7.175 | יסור |  |
| 6:7.176 | אותם |  |
| 6:7.177 | השכל | , |
| 6:7.178 | ככתוב |  |
| 6:7.179 | אל |  |
| 6:7.180 | תדמה |  |
| 6:7.181 | ישראל |  |
| 6:7.182 | אל |  |
| 6:7.183 | עמים |  |
| 6:7.184 | לא |  |
| 6:7.185 | עלו |  |
| 6:7.186 | מתוך |  |
| 6:7.187 | ההפכה | , |
| 6:7.188 | וריב |  |
| 6:7.189 | להם |  |
| 6:7.190 | עם |  |
| 6:7.191 | הדעת |  |
| 6:7.192 | אסף |  |
| 6:7.193 | חלק |  |
| 6:7.194 | א | ' |
| 6:7.195 | מאמר |  |
| 6:7.196 | א | ') |
| 6:7.197 | וכן |  |
| 6:7.198 | יאמר |  |
| 6:7.199 | הכוזרי |  |
| 6:7.200 | מאמר |  |
| 6:7.201 | ו | ' |
| 6:7.202 | סי | ' |
| 6:7.203 | ל"א |  |
| 6:7.204 | חלילה |  |
| 6:7.205 | לאל |  |
| 6:7.206 | מן |  |
| 6:7.207 | השקר |  |
| 6:7.208 | ושיבא |  |
| 6:7.209 | בתורה |  |
| 6:7.210 | מה |  |
| 6:7.211 | שהשכל |  |
| 6:7.212 | מרחיק |  |
| 6:7.213 | אותו |  |
| 6:7.214 | וישימהו |  |
| 6:7.215 | שקר | , |
| 6:7.216 | ומי |  |
| 6:7.217 | חכם |  |
| 6:7.218 | יבין |  |
| 6:7.219 | דעת | , |
| 6:7.220 | כי |  |
| 6:7.221 | רובי |  |
| 6:7.222 | מצות |  |
| 6:7.223 | התורה |  |
| 6:7.224 | לא |  |
| 6:7.225 | ניתנו |  |
| 6:7.226 | לדרוש |  |
| 6:7.227 | מפי |  |
| 6:7.228 | התורה |  |
| 6:7.229 | שבכתב |  |
| 6:7.230 | כ"א |  |
| 6:7.231 | כפי |  |
| 6:7.232 | ביאור |  |
| 6:7.233 | חכמים |  |
| 6:7.234 | ז"ל |  |
| 6:7.235 | בתורה |  |
| 6:7.236 | שבעל |  |
| 6:7.237 | פה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

