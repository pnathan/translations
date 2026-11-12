# Review pack: milhamot-hashem perek-3-04 (Milchamot ha-Shem 3:7–3:8)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-3-04.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-04`.

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
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
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


## Milchamot ha-Shem 3:7

Hebrew (Wikisource): ז) וז"ל המקובל בעל היכל הברכה, מצות עשה על הב"ד לדרוש ולחקור את העדים בשבע חקירות וכו' עוד יש בדיקות שהוא שאלות מה היה לבוש, עפר הארץ שהיה עומד עליה שחור או חול, או חמר, ואם כחשו זה לזה בטל העדות. משרשי המצוה הזאת הכלה הכלולה עולימתא שפירתא דלית בה עינין, הנהגת העולם הזה שיהא בחירה חפשית עיניו כיונים אונאה לבני נשא. ועינין דילה בהיכל זכות, עיני העדה, שדנין בידיעה על פי שנים עדים, תרין עטרין, והדעת הוא סוד עינים כנודע, ואור זה מתפשט בשבע בחינות וכ' ע"ש. ובמצות לא תלך רכיל כתב וז"ל, שלא להגיד לחברו, שכנגדו דבר עליו רע, שנאמר, לא תלך רכיל. ומדיני המצוה אף שמדבר אמת, ויש עון גדול בזה. וגם בכלל לאו זה שלא ידבר כגנות חברו אף אמת והוא כאילו כפר בעיקר וכו'. והאריך עוד בעונש המספר לשון הרע ובשכר הנשמר ממנו, יעו"ש.

| id | word | punctuation after |
|---|---|---|
| 3:7.1 | ז | ) |
| 3:7.2 | וז"ל |  |
| 3:7.3 | המקובל |  |
| 3:7.4 | בעל |  |
| 3:7.5 | היכל |  |
| 3:7.6 | הברכה | , |
| 3:7.7 | מצות |  |
| 3:7.8 | עשה |  |
| 3:7.9 | על |  |
| 3:7.10 | הב"ד |  |
| 3:7.11 | לדרוש |  |
| 3:7.12 | ולחקור |  |
| 3:7.13 | את |  |
| 3:7.14 | העדים |  |
| 3:7.15 | בשבע |  |
| 3:7.16 | חקירות |  |
| 3:7.17 | וכו | ' |
| 3:7.18 | עוד |  |
| 3:7.19 | יש |  |
| 3:7.20 | בדיקות |  |
| 3:7.21 | שהוא |  |
| 3:7.22 | שאלות |  |
| 3:7.23 | מה |  |
| 3:7.24 | היה |  |
| 3:7.25 | לבוש | , |
| 3:7.26 | עפר |  |
| 3:7.27 | הארץ |  |
| 3:7.28 | שהיה |  |
| 3:7.29 | עומד |  |
| 3:7.30 | עליה |  |
| 3:7.31 | שחור |  |
| 3:7.32 | או |  |
| 3:7.33 | חול | , |
| 3:7.34 | או |  |
| 3:7.35 | חמר | , |
| 3:7.36 | ואם |  |
| 3:7.37 | כחשו |  |
| 3:7.38 | זה |  |
| 3:7.39 | לזה |  |
| 3:7.40 | בטל |  |
| 3:7.41 | העדות | . |
| 3:7.42 | משרשי |  |
| 3:7.43 | המצוה |  |
| 3:7.44 | הזאת |  |
| 3:7.45 | הכלה |  |
| 3:7.46 | הכלולה |  |
| 3:7.47 | עולימתא |  |
| 3:7.48 | שפירתא |  |
| 3:7.49 | דלית |  |
| 3:7.50 | בה |  |
| 3:7.51 | עינין | , |
| 3:7.52 | הנהגת |  |
| 3:7.53 | העולם |  |
| 3:7.54 | הזה |  |
| 3:7.55 | שיהא |  |
| 3:7.56 | בחירה |  |
| 3:7.57 | חפשית |  |
| 3:7.58 | עיניו |  |
| 3:7.59 | כיונים |  |
| 3:7.60 | אונאה |  |
| 3:7.61 | לבני |  |
| 3:7.62 | נשא | . |
| 3:7.63 | ועינין |  |
| 3:7.64 | דילה |  |
| 3:7.65 | בהיכל |  |
| 3:7.66 | זכות | , |
| 3:7.67 | עיני |  |
| 3:7.68 | העדה | , |
| 3:7.69 | שדנין |  |
| 3:7.70 | בידיעה |  |
| 3:7.71 | על |  |
| 3:7.72 | פי |  |
| 3:7.73 | שנים |  |
| 3:7.74 | עדים | , |
| 3:7.75 | תרין |  |
| 3:7.76 | עטרין | , |
| 3:7.77 | והדעת |  |
| 3:7.78 | הוא |  |
| 3:7.79 | סוד |  |
| 3:7.80 | עינים |  |
| 3:7.81 | כנודע | , |
| 3:7.82 | ואור |  |
| 3:7.83 | זה |  |
| 3:7.84 | מתפשט |  |
| 3:7.85 | בשבע |  |
| 3:7.86 | בחינות |  |
| 3:7.87 | וכ | ' |
| 3:7.88 | ע"ש | . |
| 3:7.89 | ובמצות |  |
| 3:7.90 | לא |  |
| 3:7.91 | תלך |  |
| 3:7.92 | רכיל |  |
| 3:7.93 | כתב |  |
| 3:7.94 | וז"ל | , |
| 3:7.95 | שלא |  |
| 3:7.96 | להגיד |  |
| 3:7.97 | לחברו | , |
| 3:7.98 | שכנגדו |  |
| 3:7.99 | דבר |  |
| 3:7.100 | עליו |  |
| 3:7.101 | רע | , |
| 3:7.102 | שנאמר | , |
| 3:7.103 | לא |  |
| 3:7.104 | תלך |  |
| 3:7.105 | רכיל | . |
| 3:7.106 | ומדיני |  |
| 3:7.107 | המצוה |  |
| 3:7.108 | אף |  |
| 3:7.109 | שמדבר |  |
| 3:7.110 | אמת | , |
| 3:7.111 | ויש |  |
| 3:7.112 | עון |  |
| 3:7.113 | גדול |  |
| 3:7.114 | בזה | . |
| 3:7.115 | וגם |  |
| 3:7.116 | בכלל |  |
| 3:7.117 | לאו |  |
| 3:7.118 | זה |  |
| 3:7.119 | שלא |  |
| 3:7.120 | ידבר |  |
| 3:7.121 | כגנות |  |
| 3:7.122 | חברו |  |
| 3:7.123 | אף |  |
| 3:7.124 | אמת |  |
| 3:7.125 | והוא |  |
| 3:7.126 | כאילו |  |
| 3:7.127 | כפר |  |
| 3:7.128 | בעיקר |  |
| 3:7.129 | וכו | '. |
| 3:7.130 | והאריך |  |
| 3:7.131 | עוד |  |
| 3:7.132 | בעונש |  |
| 3:7.133 | המספר |  |
| 3:7.134 | לשון |  |
| 3:7.135 | הרע |  |
| 3:7.136 | ובשכר |  |
| 3:7.137 | הנשמר |  |
| 3:7.138 | ממנו | , |
| 3:7.139 | יעו"ש | . |


### Translation under review (3:7)

English: §7. And these are his words: "The Kabbalist, Master of the Palace of Blessing: a positive commandment on the court to inquire into and to examine the witnesses with seven inquiries, etc. There are also examinations, which are questions: what clothing he wore; the dust of the ground on which he stood, whether black, or sand, or clay; and if they contradict one another, the testimony is void. From the roots of this commandment: the included bride, the beautiful young maiden who has no eyes, the governance of this world, that there be free choice, whose eyes are like doves, a deception to the sons of man. Her eyes in the Palace of Merit: the eyes of the congregation, who judge with knowledge on the testimony of two witnesses, two crowns; and the Da'at is the secret of the eyes, as is known; and this light spreads in seven aspects." And he wrote, see there. And concerning the commandment "You shall not go about as a talebearer," he wrote: " that one should not tell his fellow that his adversary has spoken evil about him, as it is said, 'You shall not go about as a talebearer.' And the laws of the commandment: even one who speaks the truth, there is a great sin in this. And also within the scope of this prohibition: that one should not speak in disparagement of his fellow, even if true, and this is as if he denied the fundamental principle, etc." And he expanded further on the punishment of the one who tells slander and on the reward of the one who guards from it; see there.

Literal: section-7 and-these-are-his-words the-Kabbalist master-of the-palace of-blessing a-commandment positive incumbent-on the-court-(beit-din) to-inquire and-to-examine - the-witnesses with-seven inquiries etc. also there-are examinations which-are questions what was clothing the-dust of-the-ground that-was standing upon-it black or sand or clay and-if they-contradicted one the-other is-void the-testimony from-the-roots-of the-commandment this the-bride the-included the-young-maiden beautiful who-has-not in-her eyes the-governance-of the-world this that-there-be free-choice free its-eyes like-doves a-deception for-the-sons-of man and-her-eyes of-hers in-the-palace of-merit the-eyes-of the-congregation who-judge with-knowledge on the-testimony-of two witnesses two crowns and-the-Da'at is the-secret of-the-eyes as-is-known and-the-light this spreads in-seven aspects and-he-wrote see-there and-concerning not go-about as-a-talebearer he-wrote these-are-his-words that-one-should-not tell his-fellow his-adversary spoke about-him evil as-it-is-said not go-about as-a-talebearer and-the-laws-of the-commandment even one-who-speaks the-truth and-there-is a-sin great in-this and-also within-the-scope-of the-prohibition this that-one-should-not speak in-disparagement-of his-fellow even true and-this as-if he-denied the-fundamental-principle etc. and-he-expanded further on-the-punishment-of the-one-who-tells slander evil and-on-the-reward-of the-one-who-guards-himself from-it see-there

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:7.1 | section 7 |  |  | The printed paragraph number ז) is 7; rendered as the section number, as in the decisions. |
| 3:7.2 | and these are his words |  |  | Opening formula of the first quotation; it stands outside the opening quotation mark. |
| 3:7.3 | the Kabbalist |  |  |  |
| 3:7.4 | master of |  |  |  |
| 3:7.5 | the palace |  |  |  |
| 3:7.6 | of blessing |  |  |  |
| 3:7.7 | a commandment |  |  |  |
| 3:7.8 | positive |  |  |  |
| 3:7.9 | incumbent on |  |  |  |
| 3:7.10 | the court (beit din) |  |  | The abbreviation הב"ד is expanded: בית דין, the court. |
| 3:7.11 | to inquire |  |  |  |
| 3:7.12 | and to examine |  |  |  |
| 3:7.13 | - |  |  |  |
| 3:7.14 | the witnesses |  |  |  |
| 3:7.15 | with seven |  |  |  |
| 3:7.16 | inquiries |  |  |  |
| 3:7.17 | etc. |  |  |  |
| 3:7.18 | also |  |  |  |
| 3:7.19 | there are |  |  |  |
| 3:7.20 | examinations |  |  |  |
| 3:7.21 | which are |  |  |  |
| 3:7.22 | questions |  |  |  |
| 3:7.23 | what |  |  |  |
| 3:7.24 | was |  |  |  |
| 3:7.25 | clothing |  |  |  |
| 3:7.26 | the dust |  |  |  |
| 3:7.27 | of the ground |  |  |  |
| 3:7.28 | that was |  |  |  |
| 3:7.29 | standing |  |  |  |
| 3:7.30 | upon it |  |  |  |
| 3:7.31 | black |  |  |  |
| 3:7.32 | or |  |  |  |
| 3:7.33 | sand |  |  |  |
| 3:7.34 | or |  |  |  |
| 3:7.35 | clay |  |  |  |
| 3:7.36 | and if |  |  |  |
| 3:7.37 | they contradicted |  |  |  |
| 3:7.38 | one |  |  |  |
| 3:7.39 | the other |  |  |  |
| 3:7.40 | is void |  |  | בטל: is annulled, is void. |
| 3:7.41 | the testimony |  |  |  |
| 3:7.42 | from the roots of |  |  |  |
| 3:7.43 | the commandment |  |  |  |
| 3:7.44 | this |  |  |  |
| 3:7.45 | the bride |  |  |  |
| 3:7.46 | the included |  |  |  |
| 3:7.47 | the young maiden |  |  | Aramaic words; rendered by sense. |
| 3:7.48 | beautiful |  |  | Aramaic: 'beautiful'. |
| 3:7.49 | who has not |  |  | Aramaic: 'who has not' (דלית, 'there is not in her'). |
| 3:7.50 | in her |  |  |  |
| 3:7.51 | eyes |  |  |  |
| 3:7.52 | the governance of |  |  |  |
| 3:7.53 | the world |  |  |  |
| 3:7.54 | this |  |  |  |
| 3:7.55 | that there be |  |  |  |
| 3:7.56 | free choice |  |  |  |
| 3:7.57 | free |  |  |  |
| 3:7.58 | its eyes |  |  |  |
| 3:7.59 | like doves |  |  |  |
| 3:7.60 | a deception |  |  |  |
| 3:7.61 | for the sons of |  |  |  |
| 3:7.62 | man |  |  |  |
| 3:7.63 | and her eyes |  |  |  |
| 3:7.64 | of hers |  |  |  |
| 3:7.65 | in the palace |  |  |  |
| 3:7.66 | of merit |  |  |  |
| 3:7.67 | the eyes of |  |  |  |
| 3:7.68 | the congregation |  |  |  |
| 3:7.69 | who judge |  |  |  |
| 3:7.70 | with knowledge |  |  |  |
| 3:7.71 | on |  |  |  |
| 3:7.72 | the testimony of |  |  |  |
| 3:7.73 | two |  |  |  |
| 3:7.74 | witnesses |  |  |  |
| 3:7.75 | two |  |  |  |
| 3:7.76 | crowns |  |  |  |
| 3:7.77 | and the Da'at |  |  | Da'at: a Kabbalistic term, transliterated. |
| 3:7.78 | is |  |  |  |
| 3:7.79 | the secret |  |  |  |
| 3:7.80 | of the eyes |  |  |  |
| 3:7.81 | as is known |  |  |  |
| 3:7.82 | and the light |  |  |  |
| 3:7.83 | this |  |  |  |
| 3:7.84 | spreads |  |  |  |
| 3:7.85 | in seven |  |  |  |
| 3:7.86 | aspects |  |  |  |
| 3:7.87 | and he wrote |  |  |  |
| 3:7.88 | see there |  |  | עיין שם: 'see there' (the ע"ש of the author). |
| 3:7.89 | and concerning |  |  |  |
| 3:7.90 | not |  |  |  |
| 3:7.91 | go about |  |  |  |
| 3:7.92 | as a talebearer |  |  |  |
| 3:7.93 | he wrote |  |  | כתב: 'he wrote'; opens the second quotation. |
| 3:7.94 | these are his words |  |  |  |
| 3:7.95 | that one should not |  |  |  |
| 3:7.96 | tell |  |  |  |
| 3:7.97 | his fellow |  |  |  |
| 3:7.98 | his adversary |  |  | שכנגדו: 'his adversary', the one who speaks against him. |
| 3:7.99 | spoke |  |  |  |
| 3:7.100 | about him |  |  |  |
| 3:7.101 | evil |  |  |  |
| 3:7.102 | as it is said |  |  | שנאמר: formula introducing Lev 19:16. |
| 3:7.103 | not |  |  |  |
| 3:7.104 | go about |  |  |  |
| 3:7.105 | as a talebearer |  |  |  |
| 3:7.106 | and the laws of |  |  |  |
| 3:7.107 | the commandment |  |  |  |
| 3:7.108 | even |  |  |  |
| 3:7.109 | one who speaks |  |  |  |
| 3:7.110 | the truth |  |  |  |
| 3:7.111 | and there is |  |  |  |
| 3:7.112 | a sin |  |  |  |
| 3:7.113 | great |  |  |  |
| 3:7.114 | in this |  |  |  |
| 3:7.115 | and also |  |  |  |
| 3:7.116 | within the scope of |  |  |  |
| 3:7.117 | the prohibition |  |  | לאו: a prohibition (a negative commandment). |
| 3:7.118 | this |  |  |  |
| 3:7.119 | that one should not |  |  |  |
| 3:7.120 | speak |  |  |  |
| 3:7.121 | in disparagement of |  |  |  |
| 3:7.122 | his fellow |  |  |  |
| 3:7.123 | even |  |  |  |
| 3:7.124 | true |  |  |  |
| 3:7.125 | and this |  |  |  |
| 3:7.126 | as if |  |  |  |
| 3:7.127 | he denied |  |  |  |
| 3:7.128 | the fundamental principle |  |  |  |
| 3:7.129 | etc. |  |  | וכו' marks a truncated quotation and is kept as 'etc.'. |
| 3:7.130 | and he expanded |  |  |  |
| 3:7.131 | further |  |  |  |
| 3:7.132 | on the punishment of |  |  |  |
| 3:7.133 | the one who tells |  |  |  |
| 3:7.134 | slander |  |  |  |
| 3:7.135 | evil |  |  | לשון הרע is a fixed term: slander. |
| 3:7.136 | and on the reward of |  |  |  |
| 3:7.137 | the one who guards himself |  |  |  |
| 3:7.138 | from it |  |  |  |
| 3:7.139 | see there |  |  | יעו"ש = יעיין שם: 'see there'. |

Phrases:
- 3:7.1 | §7. | section-7
- 3:7.2 | And these are his words: | and-these-are-his-words
- 3:7.3 3:7.4 3:7.5 3:7.6 3:7.7 3:7.8 3:7.9 3:7.10 3:7.11 3:7.12 3:7.13 3:7.14 3:7.15 3:7.16 3:7.17 | "The Kabbalist, Master of the Palace of Blessing: a positive commandment on the court to inquire into and to examine the witnesses with seven inquiries, etc. | the-Kabbalist master-of the-palace of-blessing a-commandment positive incumbent-on the-court-(beit-din) to-inquire and-to-examine - the-witnesses with-seven inquiries etc.
- 3:7.18 3:7.19 3:7.20 3:7.21 3:7.22 3:7.23 3:7.24 3:7.25 | There are also examinations, which are questions: what clothing he wore; | also there-are examinations which-are questions what was clothing
- 3:7.26 3:7.27 3:7.28 3:7.29 3:7.30 3:7.31 3:7.32 3:7.33 3:7.34 3:7.35 | the dust of the ground on which he stood, whether black, or sand, or clay; | the-dust of-the-ground that-was standing upon-it black or sand or clay
- 3:7.36 3:7.37 3:7.38 3:7.39 3:7.40 3:7.41 | and if they contradict one another, the testimony is void. | and-if they-contradicted one the-other is-void the-testimony
- 3:7.42 3:7.43 3:7.44 3:7.45 3:7.46 3:7.47 3:7.48 3:7.49 3:7.50 3:7.51 | From the roots of this commandment: the included bride, the beautiful young maiden who has no eyes, | from-the-roots-of the-commandment this the-bride the-included the-young-maiden beautiful who-has-not in-her eyes
- 3:7.52 3:7.53 3:7.54 3:7.55 3:7.56 3:7.57 3:7.58 3:7.59 3:7.60 3:7.61 3:7.62 | the governance of this world, that there be free choice, whose eyes are like doves, a deception to the sons of man. | the-governance-of the-world this that-there-be free-choice free its-eyes like-doves a-deception for-the-sons-of man
- 3:7.63 3:7.64 3:7.65 3:7.66 3:7.67 3:7.68 | Her eyes in the Palace of Merit: the eyes of the congregation, | and-her-eyes of-hers in-the-palace of-merit the-eyes-of the-congregation
- 3:7.69 3:7.70 3:7.71 3:7.72 3:7.73 3:7.74 | who judge with knowledge on the testimony of two witnesses, | who-judge with-knowledge on the-testimony-of two witnesses
- 3:7.75 3:7.76 | two crowns; | two crowns
- 3:7.77 3:7.78 3:7.79 3:7.80 3:7.81 | and the Da'at is the secret of the eyes, as is known; | and-the-Da'at is the-secret of-the-eyes as-is-known
- 3:7.82 3:7.83 3:7.84 3:7.85 3:7.86 | and this light spreads in seven aspects." | and-the-light this spreads in-seven aspects
- 3:7.87 3:7.88 | And he wrote, see there. | and-he-wrote see-there
- 3:7.89 3:7.90 3:7.91 3:7.92 3:7.93 3:7.94 | And concerning the commandment "You shall not go about as a talebearer," he wrote: " | and-concerning not go-about as-a-talebearer he-wrote these-are-his-words
- 3:7.95 3:7.96 3:7.97 3:7.98 3:7.99 3:7.100 3:7.101 | that one should not tell his fellow that his adversary has spoken evil about him, | that-one-should-not tell his-fellow his-adversary spoke about-him evil
- 3:7.102 3:7.103 3:7.104 3:7.105 | as it is said, 'You shall not go about as a talebearer.' | as-it-is-said not go-about as-a-talebearer
- 3:7.106 3:7.107 3:7.108 3:7.109 3:7.110 | And the laws of the commandment: even one who speaks the truth, | and-the-laws-of the-commandment even one-who-speaks the-truth
- 3:7.111 3:7.112 3:7.113 3:7.114 | there is a great sin in this. | and-there-is a-sin great in-this
- 3:7.115 3:7.116 3:7.117 3:7.118 | And also within the scope of this prohibition: | and-also within-the-scope-of the-prohibition this
- 3:7.119 3:7.120 3:7.121 3:7.122 3:7.123 3:7.124 | that one should not speak in disparagement of his fellow, even if true, | that-one-should-not speak in-disparagement-of his-fellow even true
- 3:7.125 3:7.126 3:7.127 3:7.128 3:7.129 | and this is as if he denied the fundamental principle, etc." | and-this as-if he-denied the-fundamental-principle etc.
- 3:7.130 3:7.131 3:7.132 3:7.133 3:7.134 3:7.135 3:7.136 3:7.137 3:7.138 | And he expanded further on the punishment of the one who tells slander and on the reward of the one who guards from it; | and-he-expanded further on-the-punishment-of the-one-who-tells slander evil and-on-the-reward-of the-one-who-guards-himself from-it
- 3:7.139 | see there. | see-there

Grammar notes:
- The paragraph is two quotations. The first runs from וז"ל (3:7.2) to the author's own formula וכ' ע"ש; the words עולימתא שפירתא דלית בה עינין are Aramaic and are rendered by sense. The source is not named in the text.
- מצות עשה על הב"ד לדרוש ולחקור: 'a positive commandment on the court to inquire and examine'. על with a noun reads 'incumbent upon'; the infinitives give the content of the commandment.
- ואם כחשו זה לזה בטל העדות: a conditional with a perfect verb; the result clause is בטל, 'is void'. The two witnesses are the subject of the conflict (כחשו).
- The second quotation begins at כתב וז"ל (3:7.93-3:7.94) and ends at וכ' (וכו'); the closing words והאריך עוד are the author's own report and take no quotation marks.

Translator's note: The first quotation closes after 'seven aspects'; the author's ע"ש stands outside it. The second quotation has a Bible verse inside it, in single marks.

Commentary rulings:
- apparatus: Wikisource transcription. The 1931 scan is not yet collated; Wikisource additions are translated as base text (decisions, Q2).
- apparatus: Lev 19:16: 'לא תלך רכיל' (3:7.90-3:7.92 and 3:7.102-3:7.105). The quotation stops before בעמיך, so only the words on the page are translated.

## Milchamot ha-Shem 3:8

Hebrew (Wikisource): ח) ובאזהרת לא תעמוד על דם רעך, כ' אפי' עמידה כל דהו. לא תראה בצערו של חברך שאחד מלבין פניו ואתה שותק, עמדו ולא ענו עוד, אלא תהיה מוחה בכל יכלתך במי שמבזה חברך, או מדבר עליו רכילות ולשון הרע, אתה לא תהיה שומע ושותק, אלא תאמר לו רשע עוכר ישראל, למה תדבר סרה על דמות יוצרנו ועל שכינת עוזינו שהיא היו"ד שבעמיך, ומסתלקת בעון זה. ולכך נסמכה זה לזה, לא תלך רכיל ותאמר כיון שאיני מדבר אשמע ואשתוק ונהנה במה שחברו רשע מדבר, גם זה השומע רשע גמור. ועובר על לא תעמוד על דם רעך. ואין לו חלק באלהי ישראל כששומע רשע מדבר סרה על צדיק, או לוחש מלחש בשקרים על חברו רק תכף תדחה אותו ותאמר לו רשע לך מעלי, איני שומע לך, כי אם הוא שומע ושותק, וכל שכן אם נהנה, סופו נידון באש לא נופח, ולא יזכה לעמוד במחיצת הצדיקים, כשומע, כעונה, כמדבר, כולם יספו מן העולם בזה ובבא עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 3:8.1 | ח | ) |
| 3:8.2 | ובאזהרת |  |
| 3:8.3 | לא |  |
| 3:8.4 | תעמוד |  |
| 3:8.5 | על |  |
| 3:8.6 | דם |  |
| 3:8.7 | רעך | , |
| 3:8.8 | כ | ' |
| 3:8.9 | אפי | ' |
| 3:8.10 | עמידה |  |
| 3:8.11 | כל |  |
| 3:8.12 | דהו | . |
| 3:8.13 | לא |  |
| 3:8.14 | תראה |  |
| 3:8.15 | בצערו |  |
| 3:8.16 | של |  |
| 3:8.17 | חברך |  |
| 3:8.18 | שאחד |  |
| 3:8.19 | מלבין |  |
| 3:8.20 | פניו |  |
| 3:8.21 | ואתה |  |
| 3:8.22 | שותק | , |
| 3:8.23 | עמדו |  |
| 3:8.24 | ולא |  |
| 3:8.25 | ענו |  |
| 3:8.26 | עוד | , |
| 3:8.27 | אלא |  |
| 3:8.28 | תהיה |  |
| 3:8.29 | מוחה |  |
| 3:8.30 | בכל |  |
| 3:8.31 | יכלתך |  |
| 3:8.32 | במי |  |
| 3:8.33 | שמבזה |  |
| 3:8.34 | חברך | , |
| 3:8.35 | או |  |
| 3:8.36 | מדבר |  |
| 3:8.37 | עליו |  |
| 3:8.38 | רכילות |  |
| 3:8.39 | ולשון |  |
| 3:8.40 | הרע | , |
| 3:8.41 | אתה |  |
| 3:8.42 | לא |  |
| 3:8.43 | תהיה |  |
| 3:8.44 | שומע |  |
| 3:8.45 | ושותק | , |
| 3:8.46 | אלא |  |
| 3:8.47 | תאמר |  |
| 3:8.48 | לו |  |
| 3:8.49 | רשע |  |
| 3:8.50 | עוכר |  |
| 3:8.51 | ישראל | , |
| 3:8.52 | למה |  |
| 3:8.53 | תדבר |  |
| 3:8.54 | סרה |  |
| 3:8.55 | על |  |
| 3:8.56 | דמות |  |
| 3:8.57 | יוצרנו |  |
| 3:8.58 | ועל |  |
| 3:8.59 | שכינת |  |
| 3:8.60 | עוזינו |  |
| 3:8.61 | שהיא |  |
| 3:8.62 | היו"ד |  |
| 3:8.63 | שבעמיך | , |
| 3:8.64 | ומסתלקת |  |
| 3:8.65 | בעון |  |
| 3:8.66 | זה | . |
| 3:8.67 | ולכך |  |
| 3:8.68 | נסמכה |  |
| 3:8.69 | זה |  |
| 3:8.70 | לזה | , |
| 3:8.71 | לא |  |
| 3:8.72 | תלך |  |
| 3:8.73 | רכיל |  |
| 3:8.74 | ותאמר |  |
| 3:8.75 | כיון |  |
| 3:8.76 | שאיני |  |
| 3:8.77 | מדבר |  |
| 3:8.78 | אשמע |  |
| 3:8.79 | ואשתוק |  |
| 3:8.80 | ונהנה |  |
| 3:8.81 | במה |  |
| 3:8.82 | שחברו |  |
| 3:8.83 | רשע |  |
| 3:8.84 | מדבר | , |
| 3:8.85 | גם |  |
| 3:8.86 | זה |  |
| 3:8.87 | השומע |  |
| 3:8.88 | רשע |  |
| 3:8.89 | גמור | . |
| 3:8.90 | ועובר |  |
| 3:8.91 | על |  |
| 3:8.92 | לא |  |
| 3:8.93 | תעמוד |  |
| 3:8.94 | על |  |
| 3:8.95 | דם |  |
| 3:8.96 | רעך | . |
| 3:8.97 | ואין |  |
| 3:8.98 | לו |  |
| 3:8.99 | חלק |  |
| 3:8.100 | באלהי |  |
| 3:8.101 | ישראל |  |
| 3:8.102 | כששומע |  |
| 3:8.103 | רשע |  |
| 3:8.104 | מדבר |  |
| 3:8.105 | סרה |  |
| 3:8.106 | על |  |
| 3:8.107 | צדיק | , |
| 3:8.108 | או |  |
| 3:8.109 | לוחש |  |
| 3:8.110 | מלחש |  |
| 3:8.111 | בשקרים |  |
| 3:8.112 | על |  |
| 3:8.113 | חברו |  |
| 3:8.114 | רק |  |
| 3:8.115 | תכף |  |
| 3:8.116 | תדחה |  |
| 3:8.117 | אותו |  |
| 3:8.118 | ותאמר |  |
| 3:8.119 | לו |  |
| 3:8.120 | רשע |  |
| 3:8.121 | לך |  |
| 3:8.122 | מעלי | , |
| 3:8.123 | איני |  |
| 3:8.124 | שומע |  |
| 3:8.125 | לך | , |
| 3:8.126 | כי |  |
| 3:8.127 | אם |  |
| 3:8.128 | הוא |  |
| 3:8.129 | שומע |  |
| 3:8.130 | ושותק | , |
| 3:8.131 | וכל |  |
| 3:8.132 | שכן |  |
| 3:8.133 | אם |  |
| 3:8.134 | נהנה | , |
| 3:8.135 | סופו |  |
| 3:8.136 | נידון |  |
| 3:8.137 | באש |  |
| 3:8.138 | לא |  |
| 3:8.139 | נופח | , |
| 3:8.140 | ולא |  |
| 3:8.141 | יזכה |  |
| 3:8.142 | לעמוד |  |
| 3:8.143 | במחיצת |  |
| 3:8.144 | הצדיקים | , |
| 3:8.145 | כשומע | , |
| 3:8.146 | כעונה | , |
| 3:8.147 | כמדבר | , |
| 3:8.148 | כולם |  |
| 3:8.149 | יספו |  |
| 3:8.150 | מן |  |
| 3:8.151 | העולם |  |
| 3:8.152 | בזה |  |
| 3:8.153 | ובבא |  |
| 3:8.154 | עכ"ל | . |


### Translation under review (3:8)

English: §8. And in the warning, "You shall not stand idly by the blood of your fellow," he wrote: "even standing by, to any degree. You shall not see your fellow's distress, that one puts him to shame, and you are silent; he stood and did not answer again. But you shall protest with all your power against him who disgraces your fellow, or speaks about him with talebearing and evil talk, You shall not be a listener and silent, but you shall say to him: 'Wicked one, troubler of Israel, why do you speak defiance against the likeness of our Maker, and against the Shekhinah, our strength, which is the yod in "among your people," and departs through this sin.' Therefore this was juxtaposed to that, 'You shall not go about as a talebearer,' and one might say: since I do not speak, I will listen and be silent, and take pleasure in what his fellow, the wicked one, speaks. This listener too is an utterly wicked person. and transgresses 'You shall not stand idly by the blood of your fellow.' He has no portion in the God of Israel when he hears a wicked one speak defiance against a righteous person, or whispers a whisper of lies about his fellow; rather, you shall immediately push him away and say to him: 'Wicked one, go away from me, I do not listen to you.' But if he hears and is silent, and all the more so if he takes pleasure, his end is to be judged in unblown fire, and he will not merit to stand in the company of the righteous, as listener, as answerer, as speaker. All of them shall perish from the world in this one, and in the world to come." Thus far his words.

Literal: section-8 and-in-the-warning not stand-idly by the-blood-of your-fellow he-wrote even standing-by any degree not you-see his-distress of your-fellow that-one puts-to-shame his-face and-you are-silent he-stood and-not answer again but you-shall-be protesting with-all your-power against who-disgraces your-fellow or speaks about-him talebearing and-slander evil you not be a-listener and-silent but you-shall-say to-him wicked-one troubler of-Israel why do-you-speak defiance against the-likeness-of our-Maker and-against the-Shekhinah our-strength which-is the-yod in-'among-your-people' and-departs through this-sin therefore was-juxtaposed this to-that not go-about as-a-talebearer and-one-might-say since I-am-not speaking I-will-hear and-be-silent and-take-pleasure in-what his-fellow wicked-one speaks also this listener wicked utterly and-transgresses against not stand by blood-of your-fellow and-he-has-no for-him portion in-the-God-of Israel when-he-hears a-wicked-one speaks defiance against a-righteous-person or whispers a-whisper of-lies about his-fellow but immediately you-push him and-say to-him wicked-one go away-from-me I-do-not listen to-you but if he hears and-is-silent and-all the-more-so if he-takes-pleasure his-end is-judged in-fire not blown and-not he-will-merit to-stand in-the-enclosure of-the-righteous as-listener as-answerer as-speaker all-of-them shall-perish from the-world in-this and-in-the-world-to-come Thus-far-his-words

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 3:8.1 | section 8 |  |  | The printed paragraph number ח) is 8; rendered as the section number. |
| 3:8.2 | and in the warning |  |  |  |
| 3:8.3 | not |  |  |  |
| 3:8.4 | stand idly |  |  |  |
| 3:8.5 | by |  |  |  |
| 3:8.6 | the blood of |  |  |  |
| 3:8.7 | your fellow |  |  |  |
| 3:8.8 | he wrote |  |  | כ' = כתב, 'he wrote'; it opens the quotation that runs to עכ"ל. |
| 3:8.9 | even |  |  |  |
| 3:8.10 | standing by |  |  |  |
| 3:8.11 | any |  |  |  |
| 3:8.12 | degree |  |  |  |
| 3:8.13 | not |  |  |  |
| 3:8.14 | you see |  |  |  |
| 3:8.15 | his distress |  |  |  |
| 3:8.16 | of |  |  |  |
| 3:8.17 | your fellow |  |  |  |
| 3:8.18 | that one |  |  |  |
| 3:8.19 | puts to shame |  |  | Idiom: להלבין פנים, 'to whiten the face', means to shame. |
| 3:8.20 | his face |  |  |  |
| 3:8.21 | and you |  |  |  |
| 3:8.22 | are silent |  |  |  |
| 3:8.23 | he stood |  |  | עמדו is read as 'he stood'; it may also be a noun, 'his standing'. |
| 3:8.24 | and not |  |  |  |
| 3:8.25 | answer |  |  |  |
| 3:8.26 | again |  |  |  |
| 3:8.27 | but |  |  |  |
| 3:8.28 | you shall be |  |  |  |
| 3:8.29 | protesting |  |  |  |
| 3:8.30 | with all |  |  |  |
| 3:8.31 | your power |  |  |  |
| 3:8.32 | against |  |  |  |
| 3:8.33 | who disgraces |  |  |  |
| 3:8.34 | your fellow |  |  |  |
| 3:8.35 | or |  |  |  |
| 3:8.36 | speaks |  |  |  |
| 3:8.37 | about him |  |  |  |
| 3:8.38 | talebearing |  |  |  |
| 3:8.39 | and slander |  |  |  |
| 3:8.40 | evil |  |  |  |
| 3:8.41 | you |  |  |  |
| 3:8.42 | not |  |  |  |
| 3:8.43 | be |  |  |  |
| 3:8.44 | a listener |  |  |  |
| 3:8.45 | and silent |  |  |  |
| 3:8.46 | but |  |  |  |
| 3:8.47 | you shall say |  |  |  |
| 3:8.48 | to him |  |  |  |
| 3:8.49 | wicked one |  |  |  |
| 3:8.50 | troubler |  |  |  |
| 3:8.51 | of Israel |  |  |  |
| 3:8.52 | why |  |  |  |
| 3:8.53 | do you speak |  |  |  |
| 3:8.54 | defiance |  |  |  |
| 3:8.55 | against |  |  |  |
| 3:8.56 | the likeness of |  |  |  |
| 3:8.57 | our Maker |  |  |  |
| 3:8.58 | and against |  |  |  |
| 3:8.59 | the Shekhinah |  |  |  |
| 3:8.60 | our strength |  |  |  |
| 3:8.61 | which is |  |  |  |
| 3:8.62 | the yod |  |  | היו"ד: the letter yod, which the author takes as the Shekhinah. |
| 3:8.63 | in 'among your people' |  |  | The yod is the one in בעמיך of Lev 19:16, which the author cites. |
| 3:8.64 | and departs |  |  |  |
| 3:8.65 | through |  |  |  |
| 3:8.66 | this sin |  |  |  |
| 3:8.67 | therefore |  |  |  |
| 3:8.68 | was juxtaposed |  |  |  |
| 3:8.69 | this |  |  |  |
| 3:8.70 | to that |  |  |  |
| 3:8.71 | not |  |  |  |
| 3:8.72 | go about |  |  |  |
| 3:8.73 | as a talebearer |  |  |  |
| 3:8.74 | and one might say |  |  | ותאמר opens the hypothetical objection, read as a generic 'one might say'. |
| 3:8.75 | since |  |  |  |
| 3:8.76 | I am not |  |  |  |
| 3:8.77 | speaking |  |  |  |
| 3:8.78 | I will hear |  |  |  |
| 3:8.79 | and be silent |  |  |  |
| 3:8.80 | and take pleasure |  |  |  |
| 3:8.81 | in what |  |  |  |
| 3:8.82 | his fellow |  |  |  |
| 3:8.83 | wicked one |  |  |  |
| 3:8.84 | speaks |  |  |  |
| 3:8.85 | also |  |  |  |
| 3:8.86 | this |  |  |  |
| 3:8.87 | listener |  |  |  |
| 3:8.88 | wicked |  |  |  |
| 3:8.89 | utterly |  |  |  |
| 3:8.90 | and transgresses |  |  |  |
| 3:8.91 | against |  |  |  |
| 3:8.92 | not |  |  |  |
| 3:8.93 | stand |  |  |  |
| 3:8.94 | by |  |  |  |
| 3:8.95 | blood of |  |  |  |
| 3:8.96 | your fellow |  |  |  |
| 3:8.97 | and he has no |  |  |  |
| 3:8.98 | for him |  |  |  |
| 3:8.99 | portion |  |  |  |
| 3:8.100 | in the God of |  |  | באלהי ישראל: 'in the God of Israel'. |
| 3:8.101 | Israel |  |  |  |
| 3:8.102 | when he hears |  |  |  |
| 3:8.103 | a wicked one |  |  |  |
| 3:8.104 | speaks |  |  |  |
| 3:8.105 | defiance |  |  |  |
| 3:8.106 | against |  |  |  |
| 3:8.107 | a righteous person |  |  |  |
| 3:8.108 | or |  |  |  |
| 3:8.109 | whispers |  |  |  |
| 3:8.110 | a whisper |  |  |  |
| 3:8.111 | of lies |  |  |  |
| 3:8.112 | about |  |  |  |
| 3:8.113 | his fellow |  |  |  |
| 3:8.114 | but |  |  |  |
| 3:8.115 | immediately |  |  |  |
| 3:8.116 | you push |  |  |  |
| 3:8.117 | him |  |  |  |
| 3:8.118 | and say |  |  |  |
| 3:8.119 | to him |  |  |  |
| 3:8.120 | wicked one |  |  |  |
| 3:8.121 | go |  |  |  |
| 3:8.122 | away from me |  |  |  |
| 3:8.123 | I do not |  |  |  |
| 3:8.124 | listen |  |  |  |
| 3:8.125 | to you |  |  |  |
| 3:8.126 | but |  |  |  |
| 3:8.127 | if |  |  |  |
| 3:8.128 | he |  |  |  |
| 3:8.129 | hears |  |  |  |
| 3:8.130 | and is silent |  |  |  |
| 3:8.131 | and all |  |  | וכל שכן: 'and all the more so', a kal va-chomer. |
| 3:8.132 | the more so |  |  |  |
| 3:8.133 | if |  |  |  |
| 3:8.134 | he takes pleasure |  |  |  |
| 3:8.135 | his end |  |  |  |
| 3:8.136 | is judged |  |  |  |
| 3:8.137 | in fire |  |  |  |
| 3:8.138 | not |  |  |  |
| 3:8.139 | blown |  |  | לא נופח: fire that is not fanned. |
| 3:8.140 | and not |  |  |  |
| 3:8.141 | he will merit |  |  |  |
| 3:8.142 | to stand |  |  |  |
| 3:8.143 | in the enclosure |  |  | מחיצת: the enclosure, the company of the righteous. |
| 3:8.144 | of the righteous |  |  |  |
| 3:8.145 | as listener |  |  |  |
| 3:8.146 | as answerer |  |  |  |
| 3:8.147 | as speaker |  |  |  |
| 3:8.148 | all of them |  |  |  |
| 3:8.149 | shall perish |  |  |  |
| 3:8.150 | from |  |  |  |
| 3:8.151 | the world |  |  |  |
| 3:8.152 | in this |  |  |  |
| 3:8.153 | and in the world to come |  |  | ובבא: 'and in the coming (world)', the world to come. |
| 3:8.154 | Thus far his words |  |  | עכ"ל: closing formula, 'thus far his words', outside the closing quotation mark. |

Phrases:
- 3:8.1 | §8. | section-8
- 3:8.2 3:8.3 3:8.4 3:8.5 3:8.6 3:8.7 | And in the warning, "You shall not stand idly by the blood of your fellow," | and-in-the-warning not stand-idly by the-blood-of your-fellow
- 3:8.8 3:8.9 3:8.10 3:8.11 3:8.12 | he wrote: "even standing by, to any degree. | he-wrote even standing-by any degree
- 3:8.13 3:8.14 3:8.15 3:8.16 3:8.17 3:8.18 3:8.19 3:8.20 | You shall not see your fellow's distress, that one puts him to shame, | not you-see his-distress of your-fellow that-one puts-to-shame his-face
- 3:8.21 3:8.22 | and you are silent; | and-you are-silent
- 3:8.23 3:8.24 3:8.25 3:8.26 | he stood and did not answer again. | he-stood and-not answer again
- 3:8.27 3:8.28 3:8.29 3:8.30 3:8.31 | But you shall protest with all your power | but you-shall-be protesting with-all your-power
- 3:8.32 3:8.33 3:8.34 | against him who disgraces your fellow, | against who-disgraces your-fellow
- 3:8.35 3:8.36 3:8.37 3:8.38 3:8.39 3:8.40 | or speaks about him with talebearing and evil talk, | or speaks about-him talebearing and-slander evil
- 3:8.41 3:8.42 3:8.43 3:8.44 3:8.45 | You shall not be a listener and silent, | you not be a-listener and-silent
- 3:8.46 3:8.47 3:8.48 3:8.49 3:8.50 3:8.51 | but you shall say to him: 'Wicked one, troubler of Israel, | but you-shall-say to-him wicked-one troubler of-Israel
- 3:8.52 3:8.53 3:8.54 3:8.55 3:8.56 3:8.57 | why do you speak defiance against the likeness of our Maker, | why do-you-speak defiance against the-likeness-of our-Maker
- 3:8.58 3:8.59 3:8.60 3:8.61 3:8.62 3:8.63 | and against the Shekhinah, our strength, which is the yod in "among your people," | and-against the-Shekhinah our-strength which-is the-yod in-'among-your-people'
- 3:8.64 3:8.65 3:8.66 | and departs through this sin.' | and-departs through this-sin
- 3:8.67 3:8.68 3:8.69 3:8.70 | Therefore this was juxtaposed to that, | therefore was-juxtaposed this to-that
- 3:8.71 3:8.72 3:8.73 | 'You shall not go about as a talebearer,' | not go-about as-a-talebearer
- 3:8.74 3:8.75 3:8.76 3:8.77 | and one might say: since I do not speak, | and-one-might-say since I-am-not speaking
- 3:8.78 3:8.79 3:8.80 3:8.81 3:8.82 3:8.83 3:8.84 | I will listen and be silent, and take pleasure in what his fellow, the wicked one, speaks. | I-will-hear and-be-silent and-take-pleasure in-what his-fellow wicked-one speaks
- 3:8.85 3:8.86 3:8.87 3:8.88 3:8.89 | This listener too is an utterly wicked person. | also this listener wicked utterly
- 3:8.90 3:8.91 3:8.92 3:8.93 3:8.94 3:8.95 3:8.96 | and transgresses 'You shall not stand idly by the blood of your fellow.' | and-transgresses against not stand by blood-of your-fellow
- 3:8.97 3:8.98 3:8.99 3:8.100 3:8.101 | He has no portion in the God of Israel | and-he-has-no for-him portion in-the-God-of Israel
- 3:8.102 3:8.103 3:8.104 3:8.105 3:8.106 3:8.107 | when he hears a wicked one speak defiance against a righteous person, | when-he-hears a-wicked-one speaks defiance against a-righteous-person
- 3:8.108 3:8.109 3:8.110 3:8.111 3:8.112 3:8.113 | or whispers a whisper of lies about his fellow; | or whispers a-whisper of-lies about his-fellow
- 3:8.114 3:8.115 3:8.116 3:8.117 | rather, you shall immediately push him away | but immediately you-push him
- 3:8.118 3:8.119 3:8.120 3:8.121 3:8.122 | and say to him: 'Wicked one, go away from me, | and-say to-him wicked-one go away-from-me
- 3:8.123 3:8.124 3:8.125 | I do not listen to you.' | I-do-not listen to-you
- 3:8.126 3:8.127 3:8.128 3:8.129 3:8.130 | But if he hears and is silent, | but if he hears and-is-silent
- 3:8.131 3:8.132 3:8.133 3:8.134 | and all the more so if he takes pleasure, | and-all the-more-so if he-takes-pleasure
- 3:8.135 3:8.136 3:8.137 3:8.138 3:8.139 | his end is to be judged in unblown fire, | his-end is-judged in-fire not blown
- 3:8.140 3:8.141 3:8.142 3:8.143 3:8.144 | and he will not merit to stand in the company of the righteous, | and-not he-will-merit to-stand in-the-enclosure of-the-righteous
- 3:8.145 3:8.146 3:8.147 | as listener, as answerer, as speaker. | as-listener as-answerer as-speaker
- 3:8.148 3:8.149 3:8.150 3:8.151 3:8.152 | All of them shall perish from the world in this one, | all-of-them shall-perish from the-world in-this
- 3:8.153 | and in the world to come." | and-in-the-world-to-come
- 3:8.154 | Thus far his words. | Thus-far-his-words

Grammar notes:
- The quotation opens with כ' (he wrote) after the verse לא תעמוד על דם רעך, and closes at עכ"ל. Within it, direct address is in single marks, and the verse לא תלך רכיל is a single-marked quotation inside the outer quotation.
- ולכך נסמכה זה לזה: 'for this reason this was placed next to that'. The verse לא תלך רכיל is the second member of the pair; ותאמר begins the objection that the author answers.
- וכל שכן ('all the more so') introduces the consequence. The three participles כשומע, כעונה, כמדבר ('as listener, as answerer, as speaker') are the ways of taking part, and כולם יספו ('all of them shall perish') refers back to them.

Translator's note: The opening verse is in double marks because it stands outside the quotation; the verse quoted inside the quotation (3:8.90-3:8.96) is in single marks.

Commentary rulings:
- apparatus: Lev 19:16: 'לא תעמוד על דם רעך' (3:8.2-3:8.7 and 3:8.90-3:8.96), quoted as the author gives it.
- apparatus: Lev 19:16, 'בעמיך', is the source of the yod of the Shekhinah (3:8.62-3:8.63); the author's reading, not a translation of the verse's sense.

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

