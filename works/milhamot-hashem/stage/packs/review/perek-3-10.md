# Review pack: milhamot-hashem perek-3-10 (Milchamot ha-Shem 3:16–3:16)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-10.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-10`.

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
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
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


## Milchamot ha-Shem 3:16

Hebrew (Wikisource): טז) ואשר טענת עלי ואמרת מי נתן לך רשות לדרוש ולחקור אחר דבריהם, תמה אני מאד על דבריך אלה? וכי נתנה תורה לקרות בה כעוף המצפצף או כבהמה גועה ואינו יודע מה מדבר? והלא העושה כן נקרא רשע ערום כדגרסינן בסוטה (דף כב.) עולא אמר זה שקרא ושנה ולא שמש ת"ח, ר' אליעזר אמר ה"ז עם הארץ, ר' שמואל בר נחמני אמר הרי זה בור, ר' ינאי אמר ה"ז כותי, ר' אחא בר יעקב אמר ה"ז אמגוש, אמר רב נחמן בר יצחק מסתברא כרב אחא בר יעקב דאמרי אינשי רטין מגושא ולא ידע מה רטין, תני תנא ולא ידע מה תני ע"כ. ולפי דבריך עסקי הגוף ומכתבי מסחר שהם מעסקי חיי הגופות עדיפי מעסק התורה שהם חיי הנפש, אשר בהם נזכה לחיי העולם הבא ונקביל בהם פני אלהינו מלכנו? שהרי אם יביאו לך שני מכתבים מעסקי מסחר שהם סותרים זה את זה, אתה נותן לב בהם להבין כדי שלא תפסיד ואינך מקבלם באמונת אומן! ובתורה שהיא חיי הנפש תקבל את הסותר אמונתך שאמונת אומן אם טוב ואם רע בלי הבנה? הלא עיקר עסק התורה לדעת ולהבין מצותיה וחקותיה! ודוד המלך ע"ה אמר מתחנן לפני הקב"ה הבינני ואלמדה מצותיך. ואמר עוד טוב טעם ודעת למדני כי במצותיך האמנתי וקל וחומר הוא, ומה בשאר מצות של תורה צריך לב להבין שלא יטעה להתיר את האסור, או אפילו לאסור את המותר וכמאמר רז"ל הוי זהיר בתלמוד ששגגת תלמוד עולה זדון? במצות ידיעת השי"ת ויחודו על פי האמת והצדק שצונו השי"ת בתורה כמו שכתוב וידעת היום והשבות אל לבבך כי ה' הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד. ודוד המלך ע"ה צוה את בנו לאמר דע את אלהי אביך ועבדהו. והנביא ירמיהו מכריז ואמור אל יתהלל חכם בחכמתו, ואל יתהלל הגבור בגבורתו, אל יתהלל עשיר בעשרו, כי אם בזאת יתהלל המתהלל השכל וידוע אותי כי אני ה' עושה חסדו וכו'. על אחת כמה וכמה שאנחנו צריכים לידע את מי נעבוד.

| id | word | punctuation after |
|---|---|---|
| 3:16.1 | טז | ) |
| 3:16.2 | ואשר |  |
| 3:16.3 | טענת |  |
| 3:16.4 | עלי |  |
| 3:16.5 | ואמרת |  |
| 3:16.6 | מי |  |
| 3:16.7 | נתן |  |
| 3:16.8 | לך |  |
| 3:16.9 | רשות |  |
| 3:16.10 | לדרוש |  |
| 3:16.11 | ולחקור |  |
| 3:16.12 | אחר |  |
| 3:16.13 | דבריהם | , |
| 3:16.14 | תמה |  |
| 3:16.15 | אני |  |
| 3:16.16 | מאד |  |
| 3:16.17 | על |  |
| 3:16.18 | דבריך |  |
| 3:16.19 | אלה | ? |
| 3:16.20 | וכי |  |
| 3:16.21 | נתנה |  |
| 3:16.22 | תורה |  |
| 3:16.23 | לקרות |  |
| 3:16.24 | בה |  |
| 3:16.25 | כעוף |  |
| 3:16.26 | המצפצף |  |
| 3:16.27 | או |  |
| 3:16.28 | כבהמה |  |
| 3:16.29 | גועה |  |
| 3:16.30 | ואינו |  |
| 3:16.31 | יודע |  |
| 3:16.32 | מה |  |
| 3:16.33 | מדבר | ? |
| 3:16.34 | והלא |  |
| 3:16.35 | העושה |  |
| 3:16.36 | כן |  |
| 3:16.37 | נקרא |  |
| 3:16.38 | רשע |  |
| 3:16.39 | ערום |  |
| 3:16.40 | כדגרסינן |  |
| 3:16.41 | בסוטה |  |
| 3:16.42 | דף |  |
| 3:16.43 | כב | .) |
| 3:16.44 | עולא |  |
| 3:16.45 | אמר |  |
| 3:16.46 | זה |  |
| 3:16.47 | שקרא |  |
| 3:16.48 | ושנה |  |
| 3:16.49 | ולא |  |
| 3:16.50 | שמש |  |
| 3:16.51 | ת"ח | , |
| 3:16.52 | ר | ' |
| 3:16.53 | אליעזר |  |
| 3:16.54 | אמר |  |
| 3:16.55 | ה"ז |  |
| 3:16.56 | עם |  |
| 3:16.57 | הארץ | , |
| 3:16.58 | ר | ' |
| 3:16.59 | שמואל |  |
| 3:16.60 | בר |  |
| 3:16.61 | נחמני |  |
| 3:16.62 | אמר |  |
| 3:16.63 | הרי |  |
| 3:16.64 | זה |  |
| 3:16.65 | בור | , |
| 3:16.66 | ר | ' |
| 3:16.67 | ינאי |  |
| 3:16.68 | אמר |  |
| 3:16.69 | ה"ז |  |
| 3:16.70 | כותי | , |
| 3:16.71 | ר | ' |
| 3:16.72 | אחא |  |
| 3:16.73 | בר |  |
| 3:16.74 | יעקב |  |
| 3:16.75 | אמר |  |
| 3:16.76 | ה"ז |  |
| 3:16.77 | אמגוש | , |
| 3:16.78 | אמר |  |
| 3:16.79 | רב |  |
| 3:16.80 | נחמן |  |
| 3:16.81 | בר |  |
| 3:16.82 | יצחק |  |
| 3:16.83 | מסתברא |  |
| 3:16.84 | כרב |  |
| 3:16.85 | אחא |  |
| 3:16.86 | בר |  |
| 3:16.87 | יעקב |  |
| 3:16.88 | דאמרי |  |
| 3:16.89 | אינשי |  |
| 3:16.90 | רטין |  |
| 3:16.91 | מגושא |  |
| 3:16.92 | ולא |  |
| 3:16.93 | ידע |  |
| 3:16.94 | מה |  |
| 3:16.95 | רטין | , |
| 3:16.96 | תני |  |
| 3:16.97 | תנא |  |
| 3:16.98 | ולא |  |
| 3:16.99 | ידע |  |
| 3:16.100 | מה |  |
| 3:16.101 | תני |  |
| 3:16.102 | ע"כ | . |
| 3:16.103 | ולפי |  |
| 3:16.104 | דבריך |  |
| 3:16.105 | עסקי |  |
| 3:16.106 | הגוף |  |
| 3:16.107 | ומכתבי |  |
| 3:16.108 | מסחר |  |
| 3:16.109 | שהם |  |
| 3:16.110 | מעסקי |  |
| 3:16.111 | חיי |  |
| 3:16.112 | הגופות |  |
| 3:16.113 | עדיפי |  |
| 3:16.114 | מעסק |  |
| 3:16.115 | התורה |  |
| 3:16.116 | שהם |  |
| 3:16.117 | חיי |  |
| 3:16.118 | הנפש | , |
| 3:16.119 | אשר |  |
| 3:16.120 | בהם |  |
| 3:16.121 | נזכה |  |
| 3:16.122 | לחיי |  |
| 3:16.123 | העולם |  |
| 3:16.124 | הבא |  |
| 3:16.125 | ונקביל |  |
| 3:16.126 | בהם |  |
| 3:16.127 | פני |  |
| 3:16.128 | אלהינו |  |
| 3:16.129 | מלכנו | ? |
| 3:16.130 | שהרי |  |
| 3:16.131 | אם |  |
| 3:16.132 | יביאו |  |
| 3:16.133 | לך |  |
| 3:16.134 | שני |  |
| 3:16.135 | מכתבים |  |
| 3:16.136 | מעסקי |  |
| 3:16.137 | מסחר |  |
| 3:16.138 | שהם |  |
| 3:16.139 | סותרים |  |
| 3:16.140 | זה |  |
| 3:16.141 | את |  |
| 3:16.142 | זה | , |
| 3:16.143 | אתה |  |
| 3:16.144 | נותן |  |
| 3:16.145 | לב |  |
| 3:16.146 | בהם |  |
| 3:16.147 | להבין |  |
| 3:16.148 | כדי |  |
| 3:16.149 | שלא |  |
| 3:16.150 | תפסיד |  |
| 3:16.151 | ואינך |  |
| 3:16.152 | מקבלם |  |
| 3:16.153 | באמונת |  |
| 3:16.154 | אומן | ! |
| 3:16.155 | ובתורה |  |
| 3:16.156 | שהיא |  |
| 3:16.157 | חיי |  |
| 3:16.158 | הנפש |  |
| 3:16.159 | תקבל |  |
| 3:16.160 | את |  |
| 3:16.161 | הסותר |  |
| 3:16.162 | אמונתך |  |
| 3:16.163 | שאמונת |  |
| 3:16.164 | אומן |  |
| 3:16.165 | אם |  |
| 3:16.166 | טוב |  |
| 3:16.167 | ואם |  |
| 3:16.168 | רע |  |
| 3:16.169 | בלי |  |
| 3:16.170 | הבנה | ? |
| 3:16.171 | הלא |  |
| 3:16.172 | עיקר |  |
| 3:16.173 | עסק |  |
| 3:16.174 | התורה |  |
| 3:16.175 | לדעת |  |
| 3:16.176 | ולהבין |  |
| 3:16.177 | מצותיה |  |
| 3:16.178 | וחקותיה | ! |
| 3:16.179 | ודוד |  |
| 3:16.180 | המלך |  |
| 3:16.181 | ע"ה |  |
| 3:16.182 | אמר |  |
| 3:16.183 | מתחנן |  |
| 3:16.184 | לפני |  |
| 3:16.185 | הקב"ה |  |
| 3:16.186 | הבינני |  |
| 3:16.187 | ואלמדה |  |
| 3:16.188 | מצותיך | . |
| 3:16.189 | ואמר |  |
| 3:16.190 | עוד |  |
| 3:16.191 | טוב |  |
| 3:16.192 | טעם |  |
| 3:16.193 | ודעת |  |
| 3:16.194 | למדני |  |
| 3:16.195 | כי |  |
| 3:16.196 | במצותיך |  |
| 3:16.197 | האמנתי |  |
| 3:16.198 | וקל |  |
| 3:16.199 | וחומר |  |
| 3:16.200 | הוא | , |
| 3:16.201 | ומה |  |
| 3:16.202 | בשאר |  |
| 3:16.203 | מצות |  |
| 3:16.204 | של |  |
| 3:16.205 | תורה |  |
| 3:16.206 | צריך |  |
| 3:16.207 | לב |  |
| 3:16.208 | להבין |  |
| 3:16.209 | שלא |  |
| 3:16.210 | יטעה |  |
| 3:16.211 | להתיר |  |
| 3:16.212 | את |  |
| 3:16.213 | האסור | , |
| 3:16.214 | או |  |
| 3:16.215 | אפילו |  |
| 3:16.216 | לאסור |  |
| 3:16.217 | את |  |
| 3:16.218 | המותר |  |
| 3:16.219 | וכמאמר |  |
| 3:16.220 | רז"ל |  |
| 3:16.221 | הוי |  |
| 3:16.222 | זהיר |  |
| 3:16.223 | בתלמוד |  |
| 3:16.224 | ששגגת |  |
| 3:16.225 | תלמוד |  |
| 3:16.226 | עולה |  |
| 3:16.227 | זדון | ? |
| 3:16.228 | במצות |  |
| 3:16.229 | ידיעת |  |
| 3:16.230 | השי"ת |  |
| 3:16.231 | ויחודו |  |
| 3:16.232 | על |  |
| 3:16.233 | פי |  |
| 3:16.234 | האמת |  |
| 3:16.235 | והצדק |  |
| 3:16.236 | שצונו |  |
| 3:16.237 | השי"ת |  |
| 3:16.238 | בתורה |  |
| 3:16.239 | כמו |  |
| 3:16.240 | שכתוב |  |
| 3:16.241 | וידעת |  |
| 3:16.242 | היום |  |
| 3:16.243 | והשבות |  |
| 3:16.244 | אל |  |
| 3:16.245 | לבבך |  |
| 3:16.246 | כי |  |
| 3:16.247 | ה | ' |
| 3:16.248 | הוא |  |
| 3:16.249 | האלהים |  |
| 3:16.250 | בשמים |  |
| 3:16.251 | ממעל |  |
| 3:16.252 | ועל |  |
| 3:16.253 | הארץ |  |
| 3:16.254 | מתחת |  |
| 3:16.255 | אין |  |
| 3:16.256 | עוד | . |
| 3:16.257 | ודוד |  |
| 3:16.258 | המלך |  |
| 3:16.259 | ע"ה |  |
| 3:16.260 | צוה |  |
| 3:16.261 | את |  |
| 3:16.262 | בנו |  |
| 3:16.263 | לאמר |  |
| 3:16.264 | דע |  |
| 3:16.265 | את |  |
| 3:16.266 | אלהי |  |
| 3:16.267 | אביך |  |
| 3:16.268 | ועבדהו | . |
| 3:16.269 | והנביא |  |
| 3:16.270 | ירמיהו |  |
| 3:16.271 | מכריז |  |
| 3:16.272 | ואמור |  |
| 3:16.273 | אל |  |
| 3:16.274 | יתהלל |  |
| 3:16.275 | חכם |  |
| 3:16.276 | בחכמתו | , |
| 3:16.277 | ואל |  |
| 3:16.278 | יתהלל |  |
| 3:16.279 | הגבור |  |
| 3:16.280 | בגבורתו | , |
| 3:16.281 | אל |  |
| 3:16.282 | יתהלל |  |
| 3:16.283 | עשיר |  |
| 3:16.284 | בעשרו | , |
| 3:16.285 | כי |  |
| 3:16.286 | אם |  |
| 3:16.287 | בזאת |  |
| 3:16.288 | יתהלל |  |
| 3:16.289 | המתהלל |  |
| 3:16.290 | השכל |  |
| 3:16.291 | וידוע |  |
| 3:16.292 | אותי |  |
| 3:16.293 | כי |  |
| 3:16.294 | אני |  |
| 3:16.295 | ה | ' |
| 3:16.296 | עושה |  |
| 3:16.297 | חסדו |  |
| 3:16.298 | וכו | '. |
| 3:16.299 | על |  |
| 3:16.300 | אחת |  |
| 3:16.301 | כמה |  |
| 3:16.302 | וכמה |  |
| 3:16.303 | שאנחנו |  |
| 3:16.304 | צריכים |  |
| 3:16.305 | לידע |  |
| 3:16.306 | את |  |
| 3:16.307 | מי |  |
| 3:16.308 | נעבוד | . |


### Translation under review (3:16)

English: §16. You who objected to me and said, "Who gave you permission to inquire into and investigate their words?" I am greatly astonished at these words of yours? Was the Torah given to be read in like a chirping bird, or like a bellowing beast that does not know what it is saying? Is not one who does so called a wicked and crafty person, as we learn in Sotah (folio 22): "Ulla said: this is one who read and studied and did not serve the Torah scholars. Rabbi Eliezer said: this one is an unlearned person. Rabbi Shmuel bar Nachmani said: this one is an ignoramus. Rabbi Yannai said: this one is a Kuthean. Rabbi Acha bar Yaakov said: this one is a Magian. Rav Nachman bar Yitzchak said: it stands to reason like Rav Acha bar Yaakov, as people say, 'a Magian's rtin and does not know what rtin is'; a tanna taught, and did not know what he taught." Thus far. And according to your words, the affairs of the body and commercial letters, which are among the affairs of bodily life, are preferable to the affairs of the Torah, which are the life of the soul, through which we merit the life of the world to come, and through them greet the face of our God, our King? For if they bring you two letters of commerce that contradict each other, you pay attention to them, to understand, so that you do not lose, and do not accept them on a craftsman's faith! And in the Torah, which is the life of the soul, would you accept what contradicts your faith, for a craftsman's faith, whether good or bad, without understanding? Is not the main business of the Torah to know and to understand its commandments and statutes! And King David, peace be upon him, said, pleading before the Holy One, blessed be He, "Give me understanding, that I may learn Your commandments." And he said further: "Teach me good judgment and knowledge, for I have trusted in Your commandments." And this is an a fortiori argument: if in the other commandments of the Torah one needs a heart to understand, lest he err by permitting the forbidden, or even by forbidding the permitted, and as our Sages of blessed memory said: Be careful in study, for the error of study counts as wilful sin? how much more so in the commandment of knowing the Name, may He be blessed, and His unity, in accordance with truth and justice, which the Name, may He be blessed, commanded us in the Torah, as it is written: "Know this day, and take it to your heart: that the LORD is God in heaven above, and on the earth below; there is no other." And King David, peace be upon him, charged his son, saying: "Know the God of your father and serve Him." And the prophet Jeremiah proclaims, saying: "Let not the wise man glory in his wisdom, let not the strong man glory in his strength, let not the rich man glory in his riches; but in this let him who glories glory, that he has understanding and knows Me, that I am the LORD who performs His kindness, etc." All the more so that we need to know whom we serve.

Literal: 16 - you who objected to me and said: who gave you permission to inquire and to investigate after their words? Astonished I greatly at your words, these? Is it indeed that the Torah was given to be read in it, like a bird that chirps, or like a beast that bellows and does not know what it is saying? Is it not that one who does so is called wicked, crafty, as we learn in Sotah folio 22: Ulla said: this is one who read and studied and not served the Torah scholars. Rabbi Eliezer said: this one is people of the land. Rabbi Shmuel bar Nachmani said: behold this one an ignoramus. Rabbi Yannai said: this one is a Kuthean. Rabbi Acha bar Yaakov said: this one is a Magian. Said Rav Nachman bar Yitzchak: it stands to reason like Rav Acha bar Yaakov, as they say, people: rtin of a Magian, and not knows what rtin; taught a tanna and not knows what he taught. Thus far. And according to your words the affairs of the body, and letters of commerce which are of the affairs of life of the bodies, are preferable to the affair of the Torah which are life of the soul, in which we merit to the life of the world, the coming, and we greet with them the face of our God, our King? For if they bring to you two letters of commerce which contradict one the other, you give heart to them, to understand so that not you lose, and not you accept them with faith of a craftsman! And in the Torah which is life of the soul, will you accept the contradicting your faith, for faith of a craftsman whether good or bad without understanding? Is it not the main business of the Torah to know and to understand its commandments and its statutes! And David the king, peace on him, said, pleading before the Holy One, blessed be He: make me understand and that I learn Your commandments. And said further: good judgment and knowledge teach me, for in Your commandments I trusted. Light and heavy this is: if in the rest of commandments of the Torah needs a heart to understand, lest he err by permitting - the forbidden, or even by forbidding - the permitted, and as the saying of our Sages of blessed memory: be careful in study, that the error of study counts as wilful sin? in the commandment of knowing the Name, may He be blessed, and His unity, according to the mouth of truth and justice which commanded us the Name in the Torah, as it is written: and know this day and take to your heart that the LORD He is God in heaven above and on the earth below there is no other. And David the king, peace on him, charged his son saying: know - the God of your father and serve Him. And the prophet Jeremiah proclaims saying: let not the wise glory in his wisdom, and let not the strong glory in his strength, let not the rich glory in his riches; but only in this let glory the one who glories, has understanding and knows Me, that I the LORD who performs His kindness, etc. For one how much and how much, that we need to know - whom we serve.

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:16.1 | §16. |  |  | The author's section label (טז) is the book's continuous numbering: §16. |
| 3:16.2 | and who |  |  |  |
| 3:16.3 | objected |  |  |  |
| 3:16.4 | to me |  |  |  |
| 3:16.5 | and said |  |  |  |
| 3:16.6 | who |  |  |  |
| 3:16.7 | gave |  |  |  |
| 3:16.8 | to you |  |  |  |
| 3:16.9 | permission |  |  |  |
| 3:16.10 | to inquire |  |  |  |
| 3:16.11 | and to investigate |  |  |  |
| 3:16.12 | into |  |  |  |
| 3:16.13 | their words |  |  |  |
| 3:16.14 | astonished |  |  |  |
| 3:16.15 | I |  |  |  |
| 3:16.16 | greatly |  |  |  |
| 3:16.17 | at |  |  |  |
| 3:16.18 | your words |  |  |  |
| 3:16.19 | these |  |  |  |
| 3:16.20 | is it indeed that |  |  |  |
| 3:16.21 | was given |  |  |  |
| 3:16.22 | the Torah |  |  |  |
| 3:16.23 | to be read |  |  |  |
| 3:16.24 | in it |  |  |  |
| 3:16.25 | like a bird |  |  |  |
| 3:16.26 | that chirps |  |  |  |
| 3:16.27 | or |  |  |  |
| 3:16.28 | like a beast |  |  |  |
| 3:16.29 | that bellows |  |  |  |
| 3:16.30 | and does not |  |  |  |
| 3:16.31 | know |  |  |  |
| 3:16.32 | what |  |  |  |
| 3:16.33 | it is saying |  |  |  |
| 3:16.34 | is it not that |  |  |  |
| 3:16.35 | one who does |  |  |  |
| 3:16.36 | so |  |  |  |
| 3:16.37 | is called |  |  |  |
| 3:16.38 | wicked |  |  |  |
| 3:16.39 | crafty |  |  |  |
| 3:16.40 | as we learn |  |  |  |
| 3:16.41 | in Sotah |  |  |  |
| 3:16.42 | folio |  |  |  |
| 3:16.43 | 22 |  |  | The author writes 'כב.' with no a or b; the side is not supplied. |
| 3:16.44 | Ulla |  |  |  |
| 3:16.45 | said |  |  |  |
| 3:16.46 | this |  |  |  |
| 3:16.47 | one who read |  |  |  |
| 3:16.48 | and studied |  |  |  |
| 3:16.49 | and not |  |  |  |
| 3:16.50 | served |  |  |  |
| 3:16.51 | the Torah scholars |  |  | ת״ח = talmidei ḥakhamim, Torah scholars. |
| 3:16.52 | Rabbi |  |  | ר׳ = Rabbi. |
| 3:16.53 | Eliezer |  |  |  |
| 3:16.54 | said |  |  |  |
| 3:16.55 | this one is |  |  | ה״ז = הרי זה, 'this one is'. |
| 3:16.56 | people of |  |  | עם הארץ: 'the unlearned', as in Avot (the Hebrew term is עם הארץ). |
| 3:16.57 | the land |  |  |  |
| 3:16.58 | Rabbi |  |  |  |
| 3:16.59 | Shmuel |  |  |  |
| 3:16.60 | bar |  |  |  |
| 3:16.61 | Nachmani |  |  |  |
| 3:16.62 | said |  |  |  |
| 3:16.63 | this one is |  |  | הרי זה: the full form of ה״ז, 'this one is'. |
| 3:16.64 | this one |  |  |  |
| 3:16.65 | an ignoramus |  |  | בור: an uneducated, ignorant man, the Talmud's term. |
| 3:16.66 | Rabbi |  |  |  |
| 3:16.67 | Yannai |  |  |  |
| 3:16.68 | said |  |  |  |
| 3:16.69 | this one is |  |  |  |
| 3:16.70 | a Kuthean |  |  | כותי: the Talmud's term for a Samaritan. |
| 3:16.71 | Rabbi |  |  |  |
| 3:16.72 | Acha |  |  |  |
| 3:16.73 | bar |  |  |  |
| 3:16.74 | Yaakov |  |  |  |
| 3:16.75 | said |  |  |  |
| 3:16.76 | this one is |  |  |  |
| 3:16.77 | a Magian |  |  | Singular of אמגושים: a Magian. |
| 3:16.78 | said |  |  |  |
| 3:16.79 | Rav |  |  |  |
| 3:16.80 | Nachman |  |  |  |
| 3:16.81 | bar |  |  |  |
| 3:16.82 | Yitzchak |  |  |  |
| 3:16.83 | it stands to reason |  |  |  |
| 3:16.84 | like Rav |  |  |  |
| 3:16.85 | Acha |  |  |  |
| 3:16.86 | bar |  |  |  |
| 3:16.87 | Yaakov |  |  | The author writes רב here and ר׳ earlier; see apparatus. |
| 3:16.88 | as they say |  |  |  |
| 3:16.89 | people |  |  |  |
| 3:16.90 | rtin |  |  | An unexplained Aramaic word; the Talmud gives no sense, so it stays a transliteration. |
| 3:16.91 | of a Magian |  |  |  |
| 3:16.92 | and not |  |  |  |
| 3:16.93 | knows |  |  |  |
| 3:16.94 | what |  |  |  |
| 3:16.95 | rtin |  |  |  |
| 3:16.96 | taught |  |  |  |
| 3:16.97 | a tanna |  |  | תנא: a tanna, a teacher of the Mishnaic tradition. |
| 3:16.98 | and not |  |  |  |
| 3:16.99 | knows |  |  |  |
| 3:16.100 | what |  |  |  |
| 3:16.101 | he taught |  |  |  |
| 3:16.102 | thus far |  |  | ע״כ closes the quotation: 'Thus far.' |
| 3:16.103 | and according to |  |  |  |
| 3:16.104 | your words |  |  |  |
| 3:16.105 | the affairs of |  |  |  |
| 3:16.106 | the body |  |  |  |
| 3:16.107 | and letters of |  |  |  |
| 3:16.108 | commerce |  |  |  |
| 3:16.109 | which are |  |  |  |
| 3:16.110 | of the affairs of |  |  |  |
| 3:16.111 | life of |  |  |  |
| 3:16.112 | the bodies |  |  |  |
| 3:16.113 | are preferable |  |  |  |
| 3:16.114 | to the affair of |  |  |  |
| 3:16.115 | the Torah |  |  |  |
| 3:16.116 | which are |  |  |  |
| 3:16.117 | the life of |  |  |  |
| 3:16.118 | the soul |  |  |  |
| 3:16.119 | in which |  |  |  |
| 3:16.120 | in them |  |  |  |
| 3:16.121 | we merit |  |  |  |
| 3:16.122 | the life of |  |  |  |
| 3:16.123 | the world |  |  |  |
| 3:16.124 | to come |  |  |  |
| 3:16.125 | and we greet |  |  |  |
| 3:16.126 | with them |  |  |  |
| 3:16.127 | the face of |  |  |  |
| 3:16.128 | our God |  |  |  |
| 3:16.129 | our King |  |  |  |
| 3:16.130 | for indeed |  |  |  |
| 3:16.131 | if |  |  |  |
| 3:16.132 | they bring |  |  |  |
| 3:16.133 | to you |  |  |  |
| 3:16.134 | two |  |  |  |
| 3:16.135 | letters |  |  |  |
| 3:16.136 | of the affairs of |  |  |  |
| 3:16.137 | commerce |  |  |  |
| 3:16.138 | which |  |  |  |
| 3:16.139 | contradict |  |  |  |
| 3:16.140 | one |  |  |  |
| 3:16.141 | - |  |  | את זה: 'each other' (literally 'this with this'). |
| 3:16.142 | the other |  |  |  |
| 3:16.143 | you |  |  |  |
| 3:16.144 | give |  |  |  |
| 3:16.145 | heart |  |  |  |
| 3:16.146 | to them |  |  |  |
| 3:16.147 | to understand |  |  |  |
| 3:16.148 | so as |  |  |  |
| 3:16.149 | lest |  |  |  |
| 3:16.150 | you lose |  |  |  |
| 3:16.151 | and not you |  |  |  |
| 3:16.152 | accept them |  |  |  |
| 3:16.153 | with the faith of |  |  |  |
| 3:16.154 | a craftsman |  |  | אמונת אומן: trust in a craftsman's word without understanding of his work. |
| 3:16.155 | and in the Torah |  |  |  |
| 3:16.156 | which is |  |  |  |
| 3:16.157 | the life of |  |  |  |
| 3:16.158 | the soul |  |  |  |
| 3:16.159 | would you accept |  |  |  |
| 3:16.160 | - |  |  |  |
| 3:16.161 | what contradicts |  |  |  |
| 3:16.162 | your faith |  |  |  |
| 3:16.163 | for the faith of |  |  |  |
| 3:16.164 | a craftsman |  |  |  |
| 3:16.165 | whether |  |  |  |
| 3:16.166 | good |  |  |  |
| 3:16.167 | or |  |  |  |
| 3:16.168 | bad |  |  |  |
| 3:16.169 | without |  |  |  |
| 3:16.170 | understanding |  |  |  |
| 3:16.171 | is it not |  |  |  |
| 3:16.172 | the main |  |  |  |
| 3:16.173 | business of |  |  |  |
| 3:16.174 | the Torah |  |  |  |
| 3:16.175 | to know |  |  |  |
| 3:16.176 | and to understand |  |  |  |
| 3:16.177 | its commandments |  |  |  |
| 3:16.178 | and its statutes |  |  |  |
| 3:16.179 | and David |  |  |  |
| 3:16.180 | the king |  |  |  |
| 3:16.181 | peace be upon him |  |  | ע״ה: peace be upon him. |
| 3:16.182 | said |  |  |  |
| 3:16.183 | pleading |  |  |  |
| 3:16.184 | before |  |  |  |
| 3:16.185 | the Holy One, blessed be He |  |  | הקב״ה: the Holy One, blessed be He. |
| 3:16.186 | make me understand |  |  |  |
| 3:16.187 | that I may learn |  |  |  |
| 3:16.188 | Your commandments |  |  |  |
| 3:16.189 | and said |  |  |  |
| 3:16.190 | further |  |  |  |
| 3:16.191 | good |  |  |  |
| 3:16.192 | judgment |  |  |  |
| 3:16.193 | and knowledge |  |  |  |
| 3:16.194 | teach me |  |  |  |
| 3:16.195 | for |  |  |  |
| 3:16.196 | in Your commandments |  |  |  |
| 3:16.197 | I have trusted |  |  |  |
| 3:16.198 | light |  |  | קל וחומר: 'light and heavy', the a fortiori argument. |
| 3:16.199 | and heavy |  |  |  |
| 3:16.200 | this is |  |  |  |
| 3:16.201 | if |  |  | The first clause of the kal va-chomer; 'how much more so' is the implied second clause (see the phrase note). |
| 3:16.202 | in the rest of |  |  |  |
| 3:16.203 | commandments |  |  |  |
| 3:16.204 | of |  |  |  |
| 3:16.205 | the Torah |  |  |  |
| 3:16.206 | needs |  |  |  |
| 3:16.207 | a heart |  |  |  |
| 3:16.208 | to understand |  |  |  |
| 3:16.209 | lest |  |  |  |
| 3:16.210 | he err |  |  |  |
| 3:16.211 | by permitting |  |  |  |
| 3:16.212 | - |  |  | את is the object marker, with no English counterpart. |
| 3:16.213 | the forbidden |  |  |  |
| 3:16.214 | or |  |  |  |
| 3:16.215 | even |  |  |  |
| 3:16.216 | by forbidding |  |  |  |
| 3:16.217 | - |  |  | את is the object marker, with no English counterpart. |
| 3:16.218 | the permitted |  |  |  |
| 3:16.219 | and as the saying of |  |  |  |
| 3:16.220 | our Sages of blessed memory |  |  | רז״ל: our Sages of blessed memory. |
| 3:16.221 | Be |  |  | A saying of the Sages with no source given; no quotation marks, as the source is not certain. |
| 3:16.222 | careful |  |  |  |
| 3:16.223 | in study |  |  |  |
| 3:16.224 | that the error of |  |  |  |
| 3:16.225 | study |  |  |  |
| 3:16.226 | counts as |  |  |  |
| 3:16.227 | wilful sin |  |  |  |
| 3:16.228 | in the commandment of |  |  |  |
| 3:16.229 | knowing |  |  |  |
| 3:16.230 | the Name, may He be blessed |  |  | השי״ת = השם יתברך: the Name, may He be blessed. |
| 3:16.231 | and His unity |  |  |  |
| 3:16.232 | according to |  |  |  |
| 3:16.233 | the mouth of |  |  |  |
| 3:16.234 | truth |  |  |  |
| 3:16.235 | and justice |  |  |  |
| 3:16.236 | which commanded us |  |  |  |
| 3:16.237 | the Name, may He be blessed |  |  | השי״ת: the Name, may He be blessed. |
| 3:16.238 | in the Torah |  |  |  |
| 3:16.239 | as |  |  |  |
| 3:16.240 | it is written |  |  |  |
| 3:16.241 | and know |  |  |  |
| 3:16.242 | this day |  |  |  |
| 3:16.243 | and take to heart |  |  |  |
| 3:16.244 | to |  |  |  |
| 3:16.245 | your heart |  |  |  |
| 3:16.246 | that |  |  |  |
| 3:16.247 | the LORD |  |  | ה׳ as the Tetragrammaton: the LORD. |
| 3:16.248 | He |  |  |  |
| 3:16.249 | is God |  |  |  |
| 3:16.250 | in heaven |  |  |  |
| 3:16.251 | above |  |  |  |
| 3:16.252 | and on |  |  |  |
| 3:16.253 | the earth |  |  |  |
| 3:16.254 | below |  |  |  |
| 3:16.255 | there is no |  |  |  |
| 3:16.256 | other |  |  |  |
| 3:16.257 | and David |  |  |  |
| 3:16.258 | the king |  |  |  |
| 3:16.259 | peace be upon him |  |  | ע״ה: peace be upon him. |
| 3:16.260 | charged |  |  |  |
| 3:16.261 | - |  |  |  |
| 3:16.262 | his son |  |  |  |
| 3:16.263 | saying |  |  |  |
| 3:16.264 | know |  |  |  |
| 3:16.265 | - |  |  |  |
| 3:16.266 | the God of |  |  |  |
| 3:16.267 | your father |  |  |  |
| 3:16.268 | and serve Him |  |  |  |
| 3:16.269 | and the prophet |  |  |  |
| 3:16.270 | Jeremiah |  |  |  |
| 3:16.271 | proclaims |  |  |  |
| 3:16.272 | saying |  |  |  |
| 3:16.273 | let not |  |  |  |
| 3:16.274 | glory |  |  |  |
| 3:16.275 | the wise man |  |  |  |
| 3:16.276 | in his wisdom |  |  |  |
| 3:16.277 | and let not |  |  |  |
| 3:16.278 | glory |  |  |  |
| 3:16.279 | the strong man |  |  |  |
| 3:16.280 | in his strength |  |  |  |
| 3:16.281 | let not |  |  |  |
| 3:16.282 | glory |  |  |  |
| 3:16.283 | the rich man |  |  |  |
| 3:16.284 | in his riches |  |  |  |
| 3:16.285 | but |  |  |  |
| 3:16.286 | only |  |  |  |
| 3:16.287 | in this |  |  |  |
| 3:16.288 | let glory |  |  |  |
| 3:16.289 | he who glories |  |  |  |
| 3:16.290 | has understanding |  |  |  |
| 3:16.291 | and knows |  |  |  |
| 3:16.292 | Me |  |  |  |
| 3:16.293 | that |  |  |  |
| 3:16.294 | I |  |  |  |
| 3:16.295 | the LORD |  |  |  |
| 3:16.296 | who performs |  |  |  |
| 3:16.297 | His kindness |  |  |  |
| 3:16.298 | etc. |  |  | The author cuts the verse with וכו׳; 'etc.' is kept. |
| 3:16.299 | for |  |  | על אחת כמה וכמה: 'all the more so' (literally 'for one, how much and how much'). |
| 3:16.300 | one |  |  |  |
| 3:16.301 | how much |  |  |  |
| 3:16.302 | and how much |  |  |  |
| 3:16.303 | that we |  |  |  |
| 3:16.304 | need |  |  |  |
| 3:16.305 | to know |  |  |  |
| 3:16.306 | - |  |  |  |
| 3:16.307 | whom |  |  |  |
| 3:16.308 | we serve |  |  |  |

Phrases:
- 3:16.1 3:16.2 3:16.3 3:16.4 3:16.5 | §16. You who objected to me and said, | 16 - you-who objected to-me and-said
- 3:16.6 3:16.7 3:16.8 3:16.9 3:16.10 3:16.11 3:16.12 3:16.13 | "Who gave you permission to inquire into and investigate their words?" | who gave to-you permission to-inquire and-to-investigate after their-words
- 3:16.14 3:16.15 3:16.16 3:16.17 3:16.18 3:16.19 | I am greatly astonished at these words of yours? | astonished I greatly at your-words these
- 3:16.20 3:16.21 3:16.22 3:16.23 3:16.24 | Was the Torah given to be read in | is-it-indeed given the-Torah to-be-read in-it
- 3:16.25 3:16.26 | like a chirping bird, | like-a-bird that-chirps
- 3:16.27 3:16.28 3:16.29 | or like a bellowing beast | or like-a-beast that-bellows
- 3:16.30 3:16.31 3:16.32 3:16.33 | that does not know what it is saying? | and-not knows what it-says
- 3:16.34 3:16.35 3:16.36 3:16.37 3:16.38 3:16.39 | Is not one who does so called a wicked and crafty person, | is-it-not one who-does so is-called wicked crafty
- 3:16.40 3:16.41 | as we learn in Sotah | as-we-learn in-Sotah
- 3:16.42 3:16.43 | (folio 22): | folio 22
- 3:16.44 3:16.45 | "Ulla said: | Ulla said
- 3:16.46 3:16.47 3:16.48 3:16.49 3:16.50 3:16.51 | this is one who read and studied and did not serve the Torah scholars. | this (is) one who-read and-studied and-not served the-Torah-scholars
- 3:16.52 3:16.53 3:16.54 3:16.55 3:16.56 3:16.57 | Rabbi Eliezer said: this one is an unlearned person. | Rabbi Eliezer said: this-one-is people-of the-land
- 3:16.58 3:16.59 3:16.60 3:16.61 3:16.62 3:16.63 3:16.64 3:16.65 | Rabbi Shmuel bar Nachmani said: this one is an ignoramus. | Rabbi Shmuel bar Nachmani said: behold this-one an-ignoramus
- 3:16.66 3:16.67 3:16.68 3:16.69 3:16.70 | Rabbi Yannai said: this one is a Kuthean. | Rabbi Yannai said: this-one-is a-Kuthean
- 3:16.71 3:16.72 3:16.73 3:16.74 3:16.75 3:16.76 3:16.77 | Rabbi Acha bar Yaakov said: this one is a Magian. | Rabbi Acha bar Yaakov said: this-one-is a-Magian
- 3:16.78 3:16.79 3:16.80 3:16.81 3:16.82 | Rav Nachman bar Yitzchak said: | said Rav Nachman bar Yitzchak
- 3:16.83 3:16.84 3:16.85 3:16.86 3:16.87 | it stands to reason like Rav Acha bar Yaakov, | it-stands-to-reason like-Rav Acha bar Yaakov
- 3:16.88 3:16.89 3:16.90 3:16.91 | as people say, 'a Magian's rtin | as-they-say people: rtin of-a-Magian
- 3:16.92 3:16.93 3:16.94 3:16.95 | and does not know what rtin is'; | and-not knows what rtin
- 3:16.96 3:16.97 3:16.98 3:16.99 3:16.100 3:16.101 | a tanna taught, and did not know what he taught." | taught a-tanna and-not knows what he-taught
- 3:16.102 | Thus far. | thus far
- 3:16.103 3:16.104 3:16.105 3:16.106 | And according to your words, the affairs of the body | and-according-to your-words the-affairs-of the-body
- 3:16.107 3:16.108 3:16.109 3:16.110 3:16.111 3:16.112 | and commercial letters, which are among the affairs of bodily life, | and letters-of commerce which-are of-the-affairs-of life-of the-bodies
- 3:16.113 3:16.114 3:16.115 3:16.116 3:16.117 3:16.118 | are preferable to the affairs of the Torah, which are the life of the soul, | are-preferable to-the-affair-of the-Torah which-are life-of the-soul
- 3:16.119 3:16.120 3:16.121 3:16.122 3:16.123 3:16.124 | through which we merit the life of the world to come, | in-which we-merit to-life-of the-world the-coming
- 3:16.125 3:16.126 3:16.127 3:16.128 3:16.129 | and through them greet the face of our God, our King? | and-we-greet with-them the-face-of our-God our-King
- 3:16.130 3:16.131 3:16.132 3:16.133 | For if they bring you | for if they-bring to-you
- 3:16.134 3:16.135 3:16.136 3:16.137 3:16.138 3:16.139 | two letters of commerce that contradict | two letters of-commerce which contradict
- 3:16.140 3:16.141 3:16.142 | each other, | one the-other
- 3:16.143 3:16.144 3:16.145 3:16.146 | you pay attention to them, | you give heart to-them
- 3:16.147 3:16.148 3:16.149 3:16.150 | to understand, so that you do not lose, | to-understand so-that not you-lose
- 3:16.151 3:16.152 3:16.153 3:16.154 | and do not accept them on a craftsman's faith! | and-not you accept-them with-faith-of a-craftsman
- 3:16.155 3:16.156 3:16.157 3:16.158 | And in the Torah, which is the life of the soul, | and-in-the-Torah which-is life-of the-soul
- 3:16.159 3:16.160 3:16.161 3:16.162 | would you accept what contradicts your faith, | will-you-accept the-contradicting your-faith
- 3:16.163 3:16.164 3:16.165 3:16.166 3:16.167 3:16.168 3:16.169 3:16.170 | for a craftsman's faith, whether good or bad, without understanding? | for faith-of a-craftsman whether good or-bad without understanding
- 3:16.171 3:16.172 3:16.173 3:16.174 3:16.175 3:16.176 3:16.177 3:16.178 | Is not the main business of the Torah to know and to understand its commandments and statutes! | is-it-not the-main business-of the-Torah to-know and-to-understand its-commandments and-its-statutes
- 3:16.179 3:16.180 3:16.181 3:16.182 3:16.183 3:16.184 3:16.185 | And King David, peace be upon him, said, pleading before the Holy One, blessed be He, | and-David the-king peace-on-him said pleading before the-Holy-One-blessed-be-He
- 3:16.186 3:16.187 3:16.188 | "Give me understanding, that I may learn Your commandments." | make-me-understand and-that-I-learn Your-commandments
- 3:16.189 3:16.190 | And he said further: | and-said further
- 3:16.191 3:16.192 3:16.193 3:16.194 3:16.195 3:16.196 3:16.197 | "Teach me good judgment and knowledge, for I have trusted in Your commandments." | good judgment and-knowledge teach-me for in-Your-commandments I-trusted
- 3:16.198 3:16.199 3:16.200 | And this is an a fortiori argument: | light and-heavy this-is
- 3:16.201 3:16.202 3:16.203 3:16.204 3:16.205 3:16.206 3:16.207 3:16.208 | if in the other commandments of the Torah one needs a heart to understand, | if in-the-rest-of commandments-of the-Torah needs a-heart to-understand
- 3:16.209 3:16.210 3:16.211 | lest he err by permitting | lest he-err by-permitting
- 3:16.212 3:16.213 | the forbidden, | - the-forbidden
- 3:16.214 3:16.215 3:16.216 | or even by forbidding | or even by-forbidding
- 3:16.217 3:16.218 | the permitted, | - the-permitted
- 3:16.219 3:16.220 | and as our Sages of blessed memory said: | and-as the-saying-of our-Sages of-blessed-memory
- 3:16.221 3:16.222 3:16.223 3:16.224 3:16.225 3:16.226 3:16.227 | Be careful in study, for the error of study counts as wilful sin? | be careful in-study that-the-error-of study counts-as wilful-sin
- 3:16.228 3:16.229 3:16.230 3:16.231 | how much more so in the commandment of knowing the Name, may He be blessed, and His unity, | in-the-commandment-of knowing the-Name-may-He-be-blessed and-His-unity
- 3:16.232 3:16.233 3:16.234 3:16.235 3:16.236 3:16.237 3:16.238 | in accordance with truth and justice, which the Name, may He be blessed, commanded us in the Torah, | according-to the-mouth-of truth and-justice which-commanded-us the-Name in-the-Torah
- 3:16.239 3:16.240 | as it is written: | as it-is-written
- 3:16.241 3:16.242 3:16.243 3:16.244 3:16.245 | "Know this day, and take it to your heart: | and-know this-day and-take-to to-your-heart
- 3:16.246 3:16.247 3:16.248 3:16.249 3:16.250 3:16.251 | that the LORD is God in heaven above, | that the-LORD He is-God in-heaven above
- 3:16.252 3:16.253 3:16.254 3:16.255 3:16.256 | and on the earth below; there is no other." | and-on the-earth below there-is-no other
- 3:16.257 3:16.258 3:16.259 3:16.260 3:16.261 3:16.262 3:16.263 | And King David, peace be upon him, charged his son, saying: | and-David the-king peace-on-him charged his-son saying
- 3:16.264 3:16.265 3:16.266 3:16.267 3:16.268 | "Know the God of your father and serve Him." | know - the-God-of your-father and-serve-Him
- 3:16.269 3:16.270 3:16.271 3:16.272 | And the prophet Jeremiah proclaims, saying: | and-the-prophet Jeremiah proclaims saying
- 3:16.273 3:16.274 3:16.275 3:16.276 | "Let not the wise man glory in his wisdom, | let-not the-wise glory in-his-wisdom
- 3:16.277 3:16.278 3:16.279 3:16.280 | let not the strong man glory in his strength, | and-let-not the-strong glory in-his-strength
- 3:16.281 3:16.282 3:16.283 3:16.284 | let not the rich man glory in his riches; | let-not the-rich glory in-his-riches
- 3:16.285 3:16.286 3:16.287 3:16.288 3:16.289 | but in this let him who glories glory, | but only in-this let-glory the-one-who-glories
- 3:16.290 3:16.291 3:16.292 3:16.293 3:16.294 | that he has understanding and knows Me, that I | has-understanding and-knows Me that I
- 3:16.295 3:16.296 3:16.297 3:16.298 | am the LORD who performs His kindness, etc." | the-LORD who-performs His-kindness etc.
- 3:16.299 3:16.300 3:16.301 3:16.302 | All the more so | for one how-much and-how-much
- 3:16.303 3:16.304 3:16.305 3:16.306 3:16.307 3:16.308 | that we need to know whom we serve. | that we need to-know - whom we-serve

Grammar notes:
- ואשר טענת עלי ואמרת: the author addresses an objector in the second person ('you'), and the English keeps that voice to the end of the objection.
- The Sotah 22 quotation is Talmudic; its proverb 'דאמרי אינשי רטין מגושא' and the words מסתברא, תני תנא are Bavli Aramaic, translated by sense.
- ומה בשאר מצות של תורה צריך לב להבין: a kal va-chomer. The protasis is complete; the apodosis ('how much more so') is elided and is supplied in the English.
- Deut 4:39 and Ps 119:66 and 119:73 are quoted with the Hebrew imperatives and perfects rendered in the sense of the verse ('Know'; 'Teach me'; 'Give me understanding').
- אל יתהלל is a negative jussive, 'let not ... glory'; כי אם בזאת יתהלל is the same verb, 'but in this let ... glory', and the English keeps the parallel.

Translator's note: The paragraph is the author's reply to an objector (§16). The Sotah 22 passage is quoted at length; its Aramaic words are translated by sense. The kal va-chomer in the middle has an elided apodosis, supplied as 'how much more so'. Ulla's saying takes the main clause 'this is'. The closing ע״כ is 'Thus far', outside the quotation mark. The author's rtin is left untranslated, since no sense is given in the text.

Commentary rulings:
- apparatus: Wikisource additions (Q2): translated as base text, flagged here until the 1931 scan is collated.
- apparatus: Section number: the opening label טז) is §16.
- apparatus: Sotah 22: the author writes 'דף כב.' with no a or b; nothing is added.
- apparatus: The author writes רב in 'כרב אחא בר יעקב' and ר׳ in 'ר׳ אחא בר יעקב'; the same name appears to be meant.
- apparatus: Deut 4:39: the author quotes 'וידעת היום והשבות אל לבבך כי ה׳ הוא האלהים בשמים ממעל ועל הארץ מתחת אין עוד' in full.
- apparatus: 1 Chron 28:9: the author quotes 'דע את אלהי אביך ועבדהו' and stops there; the verse continues 'בלבב שלם', and nothing is added.
- apparatus: Ps 119:73: 'הבינני ואלמדה מצותיך'. Ps 119:66: 'טוב טעם ודעת למדני כי במצותיך האמנתי'.
- apparatus: Jer 9:22–23: the author writes 'עושה חסדו' where the verse reads 'עושה חסד משפט וצדקה בארץ'; the author's wording is followed and 'וכו׳' keeps the cut.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

