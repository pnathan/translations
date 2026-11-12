# Translation pack: milhamot-hashem perek-5-08 (Milchamot ha-Shem 5:14–5:14)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-08`.

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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **ריבה / לרבות** → "include", never "multiply".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **ואין צריך לומר** → "needless to say".
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


## Milchamot ha-Shem 5:14

Hebrew (Wikisource): מו) ועל יסוד רעוע זה בנה מחבר הזהר מעשה שבדה מלבו ואמר: כשחלה רשב"י וקרבו ימיו למות, שבחר לדון דינו לפני המלך העליון שהוא עתיקא קדישא, ולא לפני זעיר אנפין הנקרא בית דין כמבואר בזהר חדש, ומובא בסוף הזהר בתוספות המובאים שם וז"ל: כשחלה רשב"י עאלו קמיה רבי פנחס ור' חייא ור' אבא. אמרו ליה מאן קיימא דעלמא שכיב. אמר לון לאו בי דינא דלעילא (זעיר אנפין) מעיינין בדינאי דהא אנא חזי דלית אנא מתיהיב למלכא ולדיינא דלעילא. דאנא לאו כשאר אינשי. אלא האי דינא דילי קב"ה דייניה ולאו בדיניה. והיינו דאמר דוד שפטני אלהים וריבה ריבי. וגם שלמה אמר לעשות משפט עבדו. הוא בלחודוי ולא אחרא. דהא תנן כשאדם שכיב, בי דינא דלעילא מסתכלין בדינוי, אית מנהון נוטין לכף זכות דאחזיין זכותיה דבר נש, ואית מנהון נוטין לכף חובה דאחזיין חוביה דבר נש. ולא אפיק אנש מדיניה כמה דהוה בעי. אבל מאן דדאין קמי מלכא עלאה דשליט על כלא הוא טב. ולא יכיל איניש למהוי בההוא דינא בר טב. מאי טעמא דהא תנן כל מכילוי דמלכא עלאה נטו לזכותא תדיר והוא כוליה צד דמהימנותא. ובידוי לשבקא חובין וחטאין הה"ד כי עמך הסליחה למען תיורא, עמך הסליחה ולא עם אחר, ובגין כך בעינא קמיה דהוא דאין דינאי. ואנא איעול תליסר בני לעלמא דאתי וכו'. הרי בהדיא שסובר שרשב"י בחר לדון קמי מלכא עלאה שהוא עתיק הנקרא אצלם רעוא דרעוין שהוא ותרן וסלחן וידון לו לאדם כחפצו מפני שאינו מקפיד על מעשה התחתונים בין טוב ובין רע. כיעקב כן עשו אצלו. ואין אדם נידון ויוצא מלפניו אלא זכאי. מה שאין כן זעיר אנפין הנקרא ב"ד אינו דן את הנידון לפניו כחפצו של נידון רק כפי ראות הדיין מאן דזכי לדינא לדינא, ומאן דזכי לרחמי רחמי, ואין לך אמונת הרבה רשויות בשמים יותר מזה.

| id | word | punctuation after |
|---|---|---|
| 5:14.1 | מו | ) |
| 5:14.2 | ועל |  |
| 5:14.3 | יסוד |  |
| 5:14.4 | רעוע |  |
| 5:14.5 | זה |  |
| 5:14.6 | בנה |  |
| 5:14.7 | מחבר |  |
| 5:14.8 | הזהר |  |
| 5:14.9 | מעשה |  |
| 5:14.10 | שבדה |  |
| 5:14.11 | מלבו |  |
| 5:14.12 | ואמר | : |
| 5:14.13 | כשחלה |  |
| 5:14.14 | רשב"י |  |
| 5:14.15 | וקרבו |  |
| 5:14.16 | ימיו |  |
| 5:14.17 | למות | , |
| 5:14.18 | שבחר |  |
| 5:14.19 | לדון |  |
| 5:14.20 | דינו |  |
| 5:14.21 | לפני |  |
| 5:14.22 | המלך |  |
| 5:14.23 | העליון |  |
| 5:14.24 | שהוא |  |
| 5:14.25 | עתיקא |  |
| 5:14.26 | קדישא | , |
| 5:14.27 | ולא |  |
| 5:14.28 | לפני |  |
| 5:14.29 | זעיר |  |
| 5:14.30 | אנפין |  |
| 5:14.31 | הנקרא |  |
| 5:14.32 | בית |  |
| 5:14.33 | דין |  |
| 5:14.34 | כמבואר |  |
| 5:14.35 | בזהר |  |
| 5:14.36 | חדש | , |
| 5:14.37 | ומובא |  |
| 5:14.38 | בסוף |  |
| 5:14.39 | הזהר |  |
| 5:14.40 | בתוספות |  |
| 5:14.41 | המובאים |  |
| 5:14.42 | שם |  |
| 5:14.43 | וז"ל | : |
| 5:14.44 | כשחלה |  |
| 5:14.45 | רשב"י |  |
| 5:14.46 | עאלו |  |
| 5:14.47 | קמיה |  |
| 5:14.48 | רבי |  |
| 5:14.49 | פנחס |  |
| 5:14.50 | ור | ' |
| 5:14.51 | חייא |  |
| 5:14.52 | ור | ' |
| 5:14.53 | אבא | . |
| 5:14.54 | אמרו |  |
| 5:14.55 | ליה |  |
| 5:14.56 | מאן |  |
| 5:14.57 | קיימא |  |
| 5:14.58 | דעלמא |  |
| 5:14.59 | שכיב | . |
| 5:14.60 | אמר |  |
| 5:14.61 | לון |  |
| 5:14.62 | לאו |  |
| 5:14.63 | בי |  |
| 5:14.64 | דינא |  |
| 5:14.65 | דלעילא |  |
| 5:14.66 | זעיר |  |
| 5:14.67 | אנפין | ) |
| 5:14.68 | מעיינין |  |
| 5:14.69 | בדינאי |  |
| 5:14.70 | דהא |  |
| 5:14.71 | אנא |  |
| 5:14.72 | חזי |  |
| 5:14.73 | דלית |  |
| 5:14.74 | אנא |  |
| 5:14.75 | מתיהיב |  |
| 5:14.76 | למלכא |  |
| 5:14.77 | ולדיינא |  |
| 5:14.78 | דלעילא | . |
| 5:14.79 | דאנא |  |
| 5:14.80 | לאו |  |
| 5:14.81 | כשאר |  |
| 5:14.82 | אינשי | . |
| 5:14.83 | אלא |  |
| 5:14.84 | האי |  |
| 5:14.85 | דינא |  |
| 5:14.86 | דילי |  |
| 5:14.87 | קב"ה |  |
| 5:14.88 | דייניה |  |
| 5:14.89 | ולאו |  |
| 5:14.90 | בדיניה | . |
| 5:14.91 | והיינו |  |
| 5:14.92 | דאמר |  |
| 5:14.93 | דוד |  |
| 5:14.94 | שפטני |  |
| 5:14.95 | אלהים |  |
| 5:14.96 | וריבה |  |
| 5:14.97 | ריבי | . |
| 5:14.98 | וגם |  |
| 5:14.99 | שלמה |  |
| 5:14.100 | אמר |  |
| 5:14.101 | לעשות |  |
| 5:14.102 | משפט |  |
| 5:14.103 | עבדו | . |
| 5:14.104 | הוא |  |
| 5:14.105 | בלחודוי |  |
| 5:14.106 | ולא |  |
| 5:14.107 | אחרא | . |
| 5:14.108 | דהא |  |
| 5:14.109 | תנן |  |
| 5:14.110 | כשאדם |  |
| 5:14.111 | שכיב | , |
| 5:14.112 | בי |  |
| 5:14.113 | דינא |  |
| 5:14.114 | דלעילא |  |
| 5:14.115 | מסתכלין |  |
| 5:14.116 | בדינוי | , |
| 5:14.117 | אית |  |
| 5:14.118 | מנהון |  |
| 5:14.119 | נוטין |  |
| 5:14.120 | לכף |  |
| 5:14.121 | זכות |  |
| 5:14.122 | דאחזיין |  |
| 5:14.123 | זכותיה |  |
| 5:14.124 | דבר |  |
| 5:14.125 | נש | , |
| 5:14.126 | ואית |  |
| 5:14.127 | מנהון |  |
| 5:14.128 | נוטין |  |
| 5:14.129 | לכף |  |
| 5:14.130 | חובה |  |
| 5:14.131 | דאחזיין |  |
| 5:14.132 | חוביה |  |
| 5:14.133 | דבר |  |
| 5:14.134 | נש | . |
| 5:14.135 | ולא |  |
| 5:14.136 | אפיק |  |
| 5:14.137 | אנש |  |
| 5:14.138 | מדיניה |  |
| 5:14.139 | כמה |  |
| 5:14.140 | דהוה |  |
| 5:14.141 | בעי | . |
| 5:14.142 | אבל |  |
| 5:14.143 | מאן |  |
| 5:14.144 | דדאין |  |
| 5:14.145 | קמי |  |
| 5:14.146 | מלכא |  |
| 5:14.147 | עלאה |  |
| 5:14.148 | דשליט |  |
| 5:14.149 | על |  |
| 5:14.150 | כלא |  |
| 5:14.151 | הוא |  |
| 5:14.152 | טב | . |
| 5:14.153 | ולא |  |
| 5:14.154 | יכיל |  |
| 5:14.155 | איניש |  |
| 5:14.156 | למהוי |  |
| 5:14.157 | בההוא |  |
| 5:14.158 | דינא |  |
| 5:14.159 | בר |  |
| 5:14.160 | טב | . |
| 5:14.161 | מאי |  |
| 5:14.162 | טעמא |  |
| 5:14.163 | דהא |  |
| 5:14.164 | תנן |  |
| 5:14.165 | כל |  |
| 5:14.166 | מכילוי |  |
| 5:14.167 | דמלכא |  |
| 5:14.168 | עלאה |  |
| 5:14.169 | נטו |  |
| 5:14.170 | לזכותא |  |
| 5:14.171 | תדיר |  |
| 5:14.172 | והוא |  |
| 5:14.173 | כוליה |  |
| 5:14.174 | צד |  |
| 5:14.175 | דמהימנותא | . |
| 5:14.176 | ובידוי |  |
| 5:14.177 | לשבקא |  |
| 5:14.178 | חובין |  |
| 5:14.179 | וחטאין |  |
| 5:14.180 | הה"ד |  |
| 5:14.181 | כי |  |
| 5:14.182 | עמך |  |
| 5:14.183 | הסליחה |  |
| 5:14.184 | למען |  |
| 5:14.185 | תיורא | , |
| 5:14.186 | עמך |  |
| 5:14.187 | הסליחה |  |
| 5:14.188 | ולא |  |
| 5:14.189 | עם |  |
| 5:14.190 | אחר | , |
| 5:14.191 | ובגין |  |
| 5:14.192 | כך |  |
| 5:14.193 | בעינא |  |
| 5:14.194 | קמיה |  |
| 5:14.195 | דהוא |  |
| 5:14.196 | דאין |  |
| 5:14.197 | דינאי | . |
| 5:14.198 | ואנא |  |
| 5:14.199 | איעול |  |
| 5:14.200 | תליסר |  |
| 5:14.201 | בני |  |
| 5:14.202 | לעלמא |  |
| 5:14.203 | דאתי |  |
| 5:14.204 | וכו | '. |
| 5:14.205 | הרי |  |
| 5:14.206 | בהדיא |  |
| 5:14.207 | שסובר |  |
| 5:14.208 | שרשב"י |  |
| 5:14.209 | בחר |  |
| 5:14.210 | לדון |  |
| 5:14.211 | קמי |  |
| 5:14.212 | מלכא |  |
| 5:14.213 | עלאה |  |
| 5:14.214 | שהוא |  |
| 5:14.215 | עתיק |  |
| 5:14.216 | הנקרא |  |
| 5:14.217 | אצלם |  |
| 5:14.218 | רעוא |  |
| 5:14.219 | דרעוין |  |
| 5:14.220 | שהוא |  |
| 5:14.221 | ותרן |  |
| 5:14.222 | וסלחן |  |
| 5:14.223 | וידון |  |
| 5:14.224 | לו |  |
| 5:14.225 | לאדם |  |
| 5:14.226 | כחפצו |  |
| 5:14.227 | מפני |  |
| 5:14.228 | שאינו |  |
| 5:14.229 | מקפיד |  |
| 5:14.230 | על |  |
| 5:14.231 | מעשה |  |
| 5:14.232 | התחתונים |  |
| 5:14.233 | בין |  |
| 5:14.234 | טוב |  |
| 5:14.235 | ובין |  |
| 5:14.236 | רע | . |
| 5:14.237 | כיעקב |  |
| 5:14.238 | כן |  |
| 5:14.239 | עשו |  |
| 5:14.240 | אצלו | . |
| 5:14.241 | ואין |  |
| 5:14.242 | אדם |  |
| 5:14.243 | נידון |  |
| 5:14.244 | ויוצא |  |
| 5:14.245 | מלפניו |  |
| 5:14.246 | אלא |  |
| 5:14.247 | זכאי | . |
| 5:14.248 | מה |  |
| 5:14.249 | שאין |  |
| 5:14.250 | כן |  |
| 5:14.251 | זעיר |  |
| 5:14.252 | אנפין |  |
| 5:14.253 | הנקרא |  |
| 5:14.254 | ב"ד |  |
| 5:14.255 | אינו |  |
| 5:14.256 | דן |  |
| 5:14.257 | את |  |
| 5:14.258 | הנידון |  |
| 5:14.259 | לפניו |  |
| 5:14.260 | כחפצו |  |
| 5:14.261 | של |  |
| 5:14.262 | נידון |  |
| 5:14.263 | רק |  |
| 5:14.264 | כפי |  |
| 5:14.265 | ראות |  |
| 5:14.266 | הדיין |  |
| 5:14.267 | מאן |  |
| 5:14.268 | דזכי |  |
| 5:14.269 | לדינא |  |
| 5:14.270 | לדינא | , |
| 5:14.271 | ומאן |  |
| 5:14.272 | דזכי |  |
| 5:14.273 | לרחמי |  |
| 5:14.274 | רחמי | , |
| 5:14.275 | ואין |  |
| 5:14.276 | לך |  |
| 5:14.277 | אמונת |  |
| 5:14.278 | הרבה |  |
| 5:14.279 | רשויות |  |
| 5:14.280 | בשמים |  |
| 5:14.281 | יותר |  |
| 5:14.282 | מזה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

