# Translation pack: milhamot-hashem perek-2-11 (Milchamot ha-Shem 2:23–2:27)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-2-11.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-2-11`.

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
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **אחד ... ואחד** of two kinds of case → "alike".
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


## Milchamot ha-Shem 2:23

Hebrew (Wikisource): חמישית העם היושבים באזיא אם יבוא אליהם מכתב פקודה מהמלך בכפם יתפשוהו, בפיהם ישקוהו, ועל ראשם יניחוהו, אף כי אנו נכבד את דברי אלהים חיים מלך על כל הארץ.

| id | word | punctuation after |
|---|---|---|
| 2:23.1 | חמישית |  |
| 2:23.2 | העם |  |
| 2:23.3 | היושבים |  |
| 2:23.4 | באזיא |  |
| 2:23.5 | אם |  |
| 2:23.6 | יבוא |  |
| 2:23.7 | אליהם |  |
| 2:23.8 | מכתב |  |
| 2:23.9 | פקודה |  |
| 2:23.10 | מהמלך |  |
| 2:23.11 | בכפם |  |
| 2:23.12 | יתפשוהו | , |
| 2:23.13 | בפיהם |  |
| 2:23.14 | ישקוהו | , |
| 2:23.15 | ועל |  |
| 2:23.16 | ראשם |  |
| 2:23.17 | יניחוהו | , |
| 2:23.18 | אף |  |
| 2:23.19 | כי |  |
| 2:23.20 | אנו |  |
| 2:23.21 | נכבד |  |
| 2:23.22 | את |  |
| 2:23.23 | דברי |  |
| 2:23.24 | אלהים |  |
| 2:23.25 | חיים |  |
| 2:23.26 | מלך |  |
| 2:23.27 | על |  |
| 2:23.28 | כל |  |
| 2:23.29 | הארץ | . |


## Milchamot ha-Shem 2:24

Hebrew (Wikisource): ששית כשהביאו שבעים ושנים זקנים את התורה לתלמי מלך מצרים, קם מכסאו והשתחוה ז' פעמים לקראתה, ואם עובד אלילים אשר לא הביא צוארו בעול תורתינו כבדה כל כך, אף כי אנחנו הבאים בבריתה ואשר ניתנה לנו למורשה.

| id | word | punctuation after |
|---|---|---|
| 2:24.1 | ששית |  |
| 2:24.2 | כשהביאו |  |
| 2:24.3 | שבעים |  |
| 2:24.4 | ושנים |  |
| 2:24.5 | זקנים |  |
| 2:24.6 | את |  |
| 2:24.7 | התורה |  |
| 2:24.8 | לתלמי |  |
| 2:24.9 | מלך |  |
| 2:24.10 | מצרים | , |
| 2:24.11 | קם |  |
| 2:24.12 | מכסאו |  |
| 2:24.13 | והשתחוה |  |
| 2:24.14 | ז | ' |
| 2:24.15 | פעמים |  |
| 2:24.16 | לקראתה | , |
| 2:24.17 | ואם |  |
| 2:24.18 | עובד |  |
| 2:24.19 | אלילים |  |
| 2:24.20 | אשר |  |
| 2:24.21 | לא |  |
| 2:24.22 | הביא |  |
| 2:24.23 | צוארו |  |
| 2:24.24 | בעול |  |
| 2:24.25 | תורתינו |  |
| 2:24.26 | כבדה |  |
| 2:24.27 | כל |  |
| 2:24.28 | כך | , |
| 2:24.29 | אף |  |
| 2:24.30 | כי |  |
| 2:24.31 | אנחנו |  |
| 2:24.32 | הבאים |  |
| 2:24.33 | בבריתה |  |
| 2:24.34 | ואשר |  |
| 2:24.35 | ניתנה |  |
| 2:24.36 | לנו |  |
| 2:24.37 | למורשה | . |


## Milchamot ha-Shem 2:25

Hebrew (Wikisource): שביעית יסודי ועיקרי דתינו הם מציאות בורא אחד יחיד קדמון אשר איננו גשם ולא כח בגשם, נותן התורה לעמו ישראל על ידי המחוקק הנאמן משה עבד ה' אדון כל הנביאים. הוא יתברך המציץ ומשגיח בעולמות שברא. ועיניו משוטטות בכל, וחוקר ובוחן לבות בני אדם ושופט בצדק על עלילות מצעדי גבר. ממנו לבד נבקש עזרה בצר לנו, והוא יפרה נצר מגזע ישי אשר יקבץ נדחי ישראל מארבע כנפות הארץ, ומכל קצוצי פאה אשר זרו אותם שמה, עצמות מתיהם היבשות תחינה ונבלתם יקומון. הן כל אלה עקרי ושרשי דתינו. ומה שמץ ע"ז נמצא בם? אף לדעת אותם אשר דעותיהם ואמונותיהם נבדלו מדעותינו ומאמונתנו עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 2:25.1 | שביעית |  |
| 2:25.2 | יסודי |  |
| 2:25.3 | ועיקרי |  |
| 2:25.4 | דתינו |  |
| 2:25.5 | הם |  |
| 2:25.6 | מציאות |  |
| 2:25.7 | בורא |  |
| 2:25.8 | אחד |  |
| 2:25.9 | יחיד |  |
| 2:25.10 | קדמון |  |
| 2:25.11 | אשר |  |
| 2:25.12 | איננו |  |
| 2:25.13 | גשם |  |
| 2:25.14 | ולא |  |
| 2:25.15 | כח |  |
| 2:25.16 | בגשם | , |
| 2:25.17 | נותן |  |
| 2:25.18 | התורה |  |
| 2:25.19 | לעמו |  |
| 2:25.20 | ישראל |  |
| 2:25.21 | על |  |
| 2:25.22 | ידי |  |
| 2:25.23 | המחוקק |  |
| 2:25.24 | הנאמן |  |
| 2:25.25 | משה |  |
| 2:25.26 | עבד |  |
| 2:25.27 | ה | ' |
| 2:25.28 | אדון |  |
| 2:25.29 | כל |  |
| 2:25.30 | הנביאים | . |
| 2:25.31 | הוא |  |
| 2:25.32 | יתברך |  |
| 2:25.33 | המציץ |  |
| 2:25.34 | ומשגיח |  |
| 2:25.35 | בעולמות |  |
| 2:25.36 | שברא | . |
| 2:25.37 | ועיניו |  |
| 2:25.38 | משוטטות |  |
| 2:25.39 | בכל | , |
| 2:25.40 | וחוקר |  |
| 2:25.41 | ובוחן |  |
| 2:25.42 | לבות |  |
| 2:25.43 | בני |  |
| 2:25.44 | אדם |  |
| 2:25.45 | ושופט |  |
| 2:25.46 | בצדק |  |
| 2:25.47 | על |  |
| 2:25.48 | עלילות |  |
| 2:25.49 | מצעדי |  |
| 2:25.50 | גבר | . |
| 2:25.51 | ממנו |  |
| 2:25.52 | לבד |  |
| 2:25.53 | נבקש |  |
| 2:25.54 | עזרה |  |
| 2:25.55 | בצר |  |
| 2:25.56 | לנו | , |
| 2:25.57 | והוא |  |
| 2:25.58 | יפרה |  |
| 2:25.59 | נצר |  |
| 2:25.60 | מגזע |  |
| 2:25.61 | ישי |  |
| 2:25.62 | אשר |  |
| 2:25.63 | יקבץ |  |
| 2:25.64 | נדחי |  |
| 2:25.65 | ישראל |  |
| 2:25.66 | מארבע |  |
| 2:25.67 | כנפות |  |
| 2:25.68 | הארץ | , |
| 2:25.69 | ומכל |  |
| 2:25.70 | קצוצי |  |
| 2:25.71 | פאה |  |
| 2:25.72 | אשר |  |
| 2:25.73 | זרו |  |
| 2:25.74 | אותם |  |
| 2:25.75 | שמה | , |
| 2:25.76 | עצמות |  |
| 2:25.77 | מתיהם |  |
| 2:25.78 | היבשות |  |
| 2:25.79 | תחינה |  |
| 2:25.80 | ונבלתם |  |
| 2:25.81 | יקומון | . |
| 2:25.82 | הן |  |
| 2:25.83 | כל |  |
| 2:25.84 | אלה |  |
| 2:25.85 | עקרי |  |
| 2:25.86 | ושרשי |  |
| 2:25.87 | דתינו | . |
| 2:25.88 | ומה |  |
| 2:25.89 | שמץ |  |
| 2:25.90 | ע"ז |  |
| 2:25.91 | נמצא |  |
| 2:25.92 | בם | ? |
| 2:25.93 | אף |  |
| 2:25.94 | לדעת |  |
| 2:25.95 | אותם |  |
| 2:25.96 | אשר |  |
| 2:25.97 | דעותיהם |  |
| 2:25.98 | ואמונותיהם |  |
| 2:25.99 | נבדלו |  |
| 2:25.100 | מדעותינו |  |
| 2:25.101 | ומאמונתנו |  |
| 2:25.102 | עכ"ל | . |


## Milchamot ha-Shem 2:26

Hebrew (Wikisource): ולפי דרכינו דרך רז"ל בבריתא הנז"ל אין צורך לכל זה שאין כוונתנו שאנו משתחוים מול הס"ת אלא למי שבחר בתורה, ובהדיא אנו אומרים בואו נשתחוה ונכרעה נברכה לפני ה' עושנו. והשתחוית אברהם אבינו למלאכים וכן יהושע, איבעית אימא כיוון שנדמו להם כאנשים קרינא בהו מאי דאמור רבנן בבריתא לא תשתחוה להם, להם אי אתה משתחוה אבל אתה משתחוה לאדם שכמותך כיון שנגלו להם בדמות אנשים ואיבעית אימא שכוונת אברהם ויהושע להשתחות למי ששלחם וכמו שכ' בעל העיקרים.

| id | word | punctuation after |
|---|---|---|
| 2:26.1 | ולפי |  |
| 2:26.2 | דרכינו |  |
| 2:26.3 | דרך |  |
| 2:26.4 | רז"ל |  |
| 2:26.5 | בבריתא |  |
| 2:26.6 | הנז"ל |  |
| 2:26.7 | אין |  |
| 2:26.8 | צורך |  |
| 2:26.9 | לכל |  |
| 2:26.10 | זה |  |
| 2:26.11 | שאין |  |
| 2:26.12 | כוונתנו |  |
| 2:26.13 | שאנו |  |
| 2:26.14 | משתחוים |  |
| 2:26.15 | מול |  |
| 2:26.16 | הס"ת |  |
| 2:26.17 | אלא |  |
| 2:26.18 | למי |  |
| 2:26.19 | שבחר |  |
| 2:26.20 | בתורה | , |
| 2:26.21 | ובהדיא |  |
| 2:26.22 | אנו |  |
| 2:26.23 | אומרים |  |
| 2:26.24 | בואו |  |
| 2:26.25 | נשתחוה |  |
| 2:26.26 | ונכרעה |  |
| 2:26.27 | נברכה |  |
| 2:26.28 | לפני |  |
| 2:26.29 | ה | ' |
| 2:26.30 | עושנו | . |
| 2:26.31 | והשתחוית |  |
| 2:26.32 | אברהם |  |
| 2:26.33 | אבינו |  |
| 2:26.34 | למלאכים |  |
| 2:26.35 | וכן |  |
| 2:26.36 | יהושע | , |
| 2:26.37 | איבעית |  |
| 2:26.38 | אימא |  |
| 2:26.39 | כיוון |  |
| 2:26.40 | שנדמו |  |
| 2:26.41 | להם |  |
| 2:26.42 | כאנשים |  |
| 2:26.43 | קרינא |  |
| 2:26.44 | בהו |  |
| 2:26.45 | מאי |  |
| 2:26.46 | דאמור |  |
| 2:26.47 | רבנן |  |
| 2:26.48 | בבריתא |  |
| 2:26.49 | לא |  |
| 2:26.50 | תשתחוה |  |
| 2:26.51 | להם | , |
| 2:26.52 | להם |  |
| 2:26.53 | אי |  |
| 2:26.54 | אתה |  |
| 2:26.55 | משתחוה |  |
| 2:26.56 | אבל |  |
| 2:26.57 | אתה |  |
| 2:26.58 | משתחוה |  |
| 2:26.59 | לאדם |  |
| 2:26.60 | שכמותך |  |
| 2:26.61 | כיון |  |
| 2:26.62 | שנגלו |  |
| 2:26.63 | להם |  |
| 2:26.64 | בדמות |  |
| 2:26.65 | אנשים |  |
| 2:26.66 | ואיבעית |  |
| 2:26.67 | אימא |  |
| 2:26.68 | שכוונת |  |
| 2:26.69 | אברהם |  |
| 2:26.70 | ויהושע |  |
| 2:26.71 | להשתחות |  |
| 2:26.72 | למי |  |
| 2:26.73 | ששלחם |  |
| 2:26.74 | וכמו |  |
| 2:26.75 | שכ | ' |
| 2:26.76 | בעל |  |
| 2:26.77 | העיקרים | . |


## Milchamot ha-Shem 2:27

Hebrew (Wikisource): ושאלתך מענין הנבואה אעפ"י שאינו שייך לעניינינו ולא שאלת רק להטעותנו ולהשיאנו לדבר אחר, הנה זה מבואר מדרגות הנבואה ברמב"ם הל' יסודי התורה, ובחלק שני מספר המורה, ובספר האמונות והדעות להרס"ג. ולפי שמצאתיך חכם בעיניך אחזור לבקש ממך להשיב לי תשובה ניצחת מדברי רז"ל בתלמוד, כי רק אתה לבדך (לפי דעתך) לך נגלו כל תעלומות, ובינה יתירה ורוח חדש ניתן בקרבך. הורינו אני וחבירי איך הדברים יעלו בקנה אחד מלאות וטובות. כי נפשינו חשקה לשמוע לימודיך, ולהשתעשע בתורתך התמימה, ולך תהיה צדקה, וגם כבוד גדול תנחל מאיתנו.

| id | word | punctuation after |
|---|---|---|
| 2:27.1 | ושאלתך |  |
| 2:27.2 | מענין |  |
| 2:27.3 | הנבואה |  |
| 2:27.4 | אעפ"י |  |
| 2:27.5 | שאינו |  |
| 2:27.6 | שייך |  |
| 2:27.7 | לעניינינו |  |
| 2:27.8 | ולא |  |
| 2:27.9 | שאלת |  |
| 2:27.10 | רק |  |
| 2:27.11 | להטעותנו |  |
| 2:27.12 | ולהשיאנו |  |
| 2:27.13 | לדבר |  |
| 2:27.14 | אחר | , |
| 2:27.15 | הנה |  |
| 2:27.16 | זה |  |
| 2:27.17 | מבואר |  |
| 2:27.18 | מדרגות |  |
| 2:27.19 | הנבואה |  |
| 2:27.20 | ברמב"ם |  |
| 2:27.21 | הל | ' |
| 2:27.22 | יסודי |  |
| 2:27.23 | התורה | , |
| 2:27.24 | ובחלק |  |
| 2:27.25 | שני |  |
| 2:27.26 | מספר |  |
| 2:27.27 | המורה | , |
| 2:27.28 | ובספר |  |
| 2:27.29 | האמונות |  |
| 2:27.30 | והדעות |  |
| 2:27.31 | להרס"ג | . |
| 2:27.32 | ולפי |  |
| 2:27.33 | שמצאתיך |  |
| 2:27.34 | חכם |  |
| 2:27.35 | בעיניך |  |
| 2:27.36 | אחזור |  |
| 2:27.37 | לבקש |  |
| 2:27.38 | ממך |  |
| 2:27.39 | להשיב |  |
| 2:27.40 | לי |  |
| 2:27.41 | תשובה |  |
| 2:27.42 | ניצחת |  |
| 2:27.43 | מדברי |  |
| 2:27.44 | רז"ל |  |
| 2:27.45 | בתלמוד | , |
| 2:27.46 | כי |  |
| 2:27.47 | רק |  |
| 2:27.48 | אתה |  |
| 2:27.49 | לבדך |  |
| 2:27.50 | לפי |  |
| 2:27.51 | דעתך | ) |
| 2:27.52 | לך |  |
| 2:27.53 | נגלו |  |
| 2:27.54 | כל |  |
| 2:27.55 | תעלומות | , |
| 2:27.56 | ובינה |  |
| 2:27.57 | יתירה |  |
| 2:27.58 | ורוח |  |
| 2:27.59 | חדש |  |
| 2:27.60 | ניתן |  |
| 2:27.61 | בקרבך | . |
| 2:27.62 | הורינו |  |
| 2:27.63 | אני |  |
| 2:27.64 | וחבירי |  |
| 2:27.65 | איך |  |
| 2:27.66 | הדברים |  |
| 2:27.67 | יעלו |  |
| 2:27.68 | בקנה |  |
| 2:27.69 | אחד |  |
| 2:27.70 | מלאות |  |
| 2:27.71 | וטובות | . |
| 2:27.72 | כי |  |
| 2:27.73 | נפשינו |  |
| 2:27.74 | חשקה |  |
| 2:27.75 | לשמוע |  |
| 2:27.76 | לימודיך | , |
| 2:27.77 | ולהשתעשע |  |
| 2:27.78 | בתורתך |  |
| 2:27.79 | התמימה | , |
| 2:27.80 | ולך |  |
| 2:27.81 | תהיה |  |
| 2:27.82 | צדקה | , |
| 2:27.83 | וגם |  |
| 2:27.84 | כבוד |  |
| 2:27.85 | גדול |  |
| 2:27.86 | תנחל |  |
| 2:27.87 | מאיתנו | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

