# Translation pack: milhamot-hashem perek-6-18 (Milchamot ha-Shem 6:28–6:28)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-18.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-18`.

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
- **שכינה**: fixed, "the Shekhinah", never "the Divine Presence"; שכינתיה → "His
  Shekhinah".
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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
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


## Milchamot ha-Shem 6:28

Hebrew (Wikisource): עג) ובירושלמי פ' הרואה גרסינן המינין שאלו את ר' שמלאי כמה אלוהות בראו את העולם, אמר להם ולי אתם שולאים, לכו ושאלו את אדם הראשון, שנ' כי שאל נא לימים ראשונים וגו' אשר ברא אלהים אדם על הארץ, אין כתיב כאן אלא למן היום אשר ברא אלהים אדם על הארץ, אמרו ליה והא כתיב בראשית ברא אלהים, אמר להו וכי כתיב בראו, אין כתיב אלא ברא, אמר ר' שמלאי כל מקום שפקרו המינין תשובתן בצדן, חזרו ושאלו אותו מה ההן דכתיב נעשה אדם בצלמינו כדמותינו, אמר להם ויבראו אלהים את האדם בצלמם אין כתיב כאן, אלא ויברא אלהים את האדם בצלמו, אמרו לו תלמידיו לאלה דחית בקנה, לנו מה אתה משיב, אמר להם לשעבר אדם נברא מן העפר, וחוה נבראת מן האדם, מאדם ואילך בצלמינו כדמותינו, אי אפשר לאיש בלא אשה, ואי אפשר לאשה בלא איש, אי אפשר לשניהם בלא שכינה, וחזרו ושאלו אותן, מה ההן דכתיב אל אלהים ה', אל אלהים ה' הוא יודע, אמר להם הם יודעים אין כתיב כאן. אלא הוא יודע כתיב, אמרו לו תלמידי הרבי לאלו דחית בקנה, לנו מה אתה משיב, אמר להן שלשתן שם אחד כאיניש דאמר בסיליוס קיסר אנושתום (אדום מושל ושליט). חזרו ושאלו אותו, מהו דכתיב אל אלהים ה' דבר ויקרא ארץ, אמר להן וכי דברו-ויקראו כתוב כאן, אין כתיב אלא דבר ויקראו ארץ, אמרו לו תלמידיו רבי לאלה דחית בקנה, לנו מה אתה משיב אמר להם שלשתן שם אחד הן כאיניש דאמר אומנון בניןן ארכיטקטנן (נ"א ארכי קטונדס ערוך אמן). חזרו ושאלו אותו מהו דכתיב כי אלהים קדושים הוא אמר להם קדושים המה אין כתוב כאן אלא הוא אל קנא הוא, אמרו לו תלמידיו רבי לאלו דחית בקנה. ולנו מה אתה משיב, אמר להם קדוש הוא בכל מיני קדושות, דאמר רבי יודן בשם רבי אחא הקב"ה דרכו בקדושה, דיבורו בקדושה, ישובו בקדושה, אלהים בקדש דרכך, הילוכו בקדושה, הליכות אל מלכי בקדש, מושבו בקדושה אלהים ישב על כסא קדשו, דיבורו בקדושה אלהים דיבר בקדשו, חשיפת דרועו בקדושה, חשף ה' את זרוע קדשו נורא ואדיר בקדושה, מי כמוך נאדר בקדש, חזרו ושאלו אותו מהו ההן דכתיב מי גוי גדול אשר לו אלהים קרובים אליו. אמר להם כה' אלהינו בכל קראינו אליהם אין כתיב כאן, אלא בכל קראינו אליו. אמרו לו תלמידיו רבי לאלו דחית בקנה. לנו מה אתה משיב, אמר להם קרוב הוא בכל מיני קריבות דאמר ר' פנחס בשפ ר' יהודה בר' סימון. ע"ז נראית קרובה ואינה אלא רחוקה. מה טעם ישאוהו על כתף יסבלוהו וגו' סוף דבר אלוהו עמו בבית והוא צועק עד שימות ולא ישמע ולא יושיע מצרתו. אבל הקב"ה נראה רחוק ואין קרוב ממנו, דאמר רבי לוי מהארץ עד לרקיע מהלך חמש מאות שנה. וכן לכל רקיע ורקיע, ואמר רבי חלבו ורבי ברכיה בשם רבי אבא סמוקא אף טלפי החיות מהלך חמש מאות חמש עשרה מנין ישרה, ראה כמה גבוה הוא מעולמו, ואדם נכנס לבית הכנסת ועומד אחורי העמוד ומתפלל בלחישה והקב"ה מאזין את תפלתו שנאמר וחנה היא מדברת אל לבה רק שפתיה נעות והאזין הקב"ה תפלתה. וכן כל בריותיו שנ' תפלה לעני כי יעטף כאדם המשיח באזן חבירו והוא שומע, וכי יש לך אלוה קרוב משה שהוא קרוב לבריותיו מפה לאזן ע"כ.

| id | word | punctuation after |
|---|---|---|
| 6:28.1 | עג | ) |
| 6:28.2 | ובירושלמי |  |
| 6:28.3 | פ | ' |
| 6:28.4 | הרואה |  |
| 6:28.5 | גרסינן |  |
| 6:28.6 | המינין |  |
| 6:28.7 | שאלו |  |
| 6:28.8 | את |  |
| 6:28.9 | ר | ' |
| 6:28.10 | שמלאי |  |
| 6:28.11 | כמה |  |
| 6:28.12 | אלוהות |  |
| 6:28.13 | בראו |  |
| 6:28.14 | את |  |
| 6:28.15 | העולם | , |
| 6:28.16 | אמר |  |
| 6:28.17 | להם |  |
| 6:28.18 | ולי |  |
| 6:28.19 | אתם |  |
| 6:28.20 | שולאים | , |
| 6:28.21 | לכו |  |
| 6:28.22 | ושאלו |  |
| 6:28.23 | את |  |
| 6:28.24 | אדם |  |
| 6:28.25 | הראשון | , |
| 6:28.26 | שנ | ' |
| 6:28.27 | כי |  |
| 6:28.28 | שאל |  |
| 6:28.29 | נא |  |
| 6:28.30 | לימים |  |
| 6:28.31 | ראשונים |  |
| 6:28.32 | וגו | ' |
| 6:28.33 | אשר |  |
| 6:28.34 | ברא |  |
| 6:28.35 | אלהים |  |
| 6:28.36 | אדם |  |
| 6:28.37 | על |  |
| 6:28.38 | הארץ | , |
| 6:28.39 | אין |  |
| 6:28.40 | כתיב |  |
| 6:28.41 | כאן |  |
| 6:28.42 | אלא |  |
| 6:28.43 | למן |  |
| 6:28.44 | היום |  |
| 6:28.45 | אשר |  |
| 6:28.46 | ברא |  |
| 6:28.47 | אלהים |  |
| 6:28.48 | אדם |  |
| 6:28.49 | על |  |
| 6:28.50 | הארץ | , |
| 6:28.51 | אמרו |  |
| 6:28.52 | ליה |  |
| 6:28.53 | והא |  |
| 6:28.54 | כתיב |  |
| 6:28.55 | בראשית |  |
| 6:28.56 | ברא |  |
| 6:28.57 | אלהים | , |
| 6:28.58 | אמר |  |
| 6:28.59 | להו |  |
| 6:28.60 | וכי |  |
| 6:28.61 | כתיב |  |
| 6:28.62 | בראו | , |
| 6:28.63 | אין |  |
| 6:28.64 | כתיב |  |
| 6:28.65 | אלא |  |
| 6:28.66 | ברא | , |
| 6:28.67 | אמר |  |
| 6:28.68 | ר | ' |
| 6:28.69 | שמלאי |  |
| 6:28.70 | כל |  |
| 6:28.71 | מקום |  |
| 6:28.72 | שפקרו |  |
| 6:28.73 | המינין |  |
| 6:28.74 | תשובתן |  |
| 6:28.75 | בצדן | , |
| 6:28.76 | חזרו |  |
| 6:28.77 | ושאלו |  |
| 6:28.78 | אותו |  |
| 6:28.79 | מה |  |
| 6:28.80 | ההן |  |
| 6:28.81 | דכתיב |  |
| 6:28.82 | נעשה |  |
| 6:28.83 | אדם |  |
| 6:28.84 | בצלמינו |  |
| 6:28.85 | כדמותינו | , |
| 6:28.86 | אמר |  |
| 6:28.87 | להם |  |
| 6:28.88 | ויבראו |  |
| 6:28.89 | אלהים |  |
| 6:28.90 | את |  |
| 6:28.91 | האדם |  |
| 6:28.92 | בצלמם |  |
| 6:28.93 | אין |  |
| 6:28.94 | כתיב |  |
| 6:28.95 | כאן | , |
| 6:28.96 | אלא |  |
| 6:28.97 | ויברא |  |
| 6:28.98 | אלהים |  |
| 6:28.99 | את |  |
| 6:28.100 | האדם |  |
| 6:28.101 | בצלמו | , |
| 6:28.102 | אמרו |  |
| 6:28.103 | לו |  |
| 6:28.104 | תלמידיו |  |
| 6:28.105 | לאלה |  |
| 6:28.106 | דחית |  |
| 6:28.107 | בקנה | , |
| 6:28.108 | לנו |  |
| 6:28.109 | מה |  |
| 6:28.110 | אתה |  |
| 6:28.111 | משיב | , |
| 6:28.112 | אמר |  |
| 6:28.113 | להם |  |
| 6:28.114 | לשעבר |  |
| 6:28.115 | אדם |  |
| 6:28.116 | נברא |  |
| 6:28.117 | מן |  |
| 6:28.118 | העפר | , |
| 6:28.119 | וחוה |  |
| 6:28.120 | נבראת |  |
| 6:28.121 | מן |  |
| 6:28.122 | האדם | , |
| 6:28.123 | מאדם |  |
| 6:28.124 | ואילך |  |
| 6:28.125 | בצלמינו |  |
| 6:28.126 | כדמותינו | , |
| 6:28.127 | אי |  |
| 6:28.128 | אפשר |  |
| 6:28.129 | לאיש |  |
| 6:28.130 | בלא |  |
| 6:28.131 | אשה | , |
| 6:28.132 | ואי |  |
| 6:28.133 | אפשר |  |
| 6:28.134 | לאשה |  |
| 6:28.135 | בלא |  |
| 6:28.136 | איש | , |
| 6:28.137 | אי |  |
| 6:28.138 | אפשר |  |
| 6:28.139 | לשניהם |  |
| 6:28.140 | בלא |  |
| 6:28.141 | שכינה | , |
| 6:28.142 | וחזרו |  |
| 6:28.143 | ושאלו |  |
| 6:28.144 | אותן | , |
| 6:28.145 | מה |  |
| 6:28.146 | ההן |  |
| 6:28.147 | דכתיב |  |
| 6:28.148 | אל |  |
| 6:28.149 | אלהים |  |
| 6:28.150 | ה | ', |
| 6:28.151 | אל |  |
| 6:28.152 | אלהים |  |
| 6:28.153 | ה | ' |
| 6:28.154 | הוא |  |
| 6:28.155 | יודע | , |
| 6:28.156 | אמר |  |
| 6:28.157 | להם |  |
| 6:28.158 | הם |  |
| 6:28.159 | יודעים |  |
| 6:28.160 | אין |  |
| 6:28.161 | כתיב |  |
| 6:28.162 | כאן | . |
| 6:28.163 | אלא |  |
| 6:28.164 | הוא |  |
| 6:28.165 | יודע |  |
| 6:28.166 | כתיב | , |
| 6:28.167 | אמרו |  |
| 6:28.168 | לו |  |
| 6:28.169 | תלמידי |  |
| 6:28.170 | הרבי |  |
| 6:28.171 | לאלו |  |
| 6:28.172 | דחית |  |
| 6:28.173 | בקנה | , |
| 6:28.174 | לנו |  |
| 6:28.175 | מה |  |
| 6:28.176 | אתה |  |
| 6:28.177 | משיב | , |
| 6:28.178 | אמר |  |
| 6:28.179 | להן |  |
| 6:28.180 | שלשתן |  |
| 6:28.181 | שם |  |
| 6:28.182 | אחד |  |
| 6:28.183 | כאיניש |  |
| 6:28.184 | דאמר |  |
| 6:28.185 | בסיליוס |  |
| 6:28.186 | קיסר |  |
| 6:28.187 | אנושתום |  |
| 6:28.188 | אדום |  |
| 6:28.189 | מושל |  |
| 6:28.190 | ושליט | ). |
| 6:28.191 | חזרו |  |
| 6:28.192 | ושאלו |  |
| 6:28.193 | אותו | , |
| 6:28.194 | מהו |  |
| 6:28.195 | דכתיב |  |
| 6:28.196 | אל |  |
| 6:28.197 | אלהים |  |
| 6:28.198 | ה | ' |
| 6:28.199 | דבר |  |
| 6:28.200 | ויקרא |  |
| 6:28.201 | ארץ | , |
| 6:28.202 | אמר |  |
| 6:28.203 | להן |  |
| 6:28.204 | וכי |  |
| 6:28.205 | דברו-ויקראו |  |
| 6:28.206 | כתוב |  |
| 6:28.207 | כאן | , |
| 6:28.208 | אין |  |
| 6:28.209 | כתיב |  |
| 6:28.210 | אלא |  |
| 6:28.211 | דבר |  |
| 6:28.212 | ויקראו |  |
| 6:28.213 | ארץ | , |
| 6:28.214 | אמרו |  |
| 6:28.215 | לו |  |
| 6:28.216 | תלמידיו |  |
| 6:28.217 | רבי |  |
| 6:28.218 | לאלה |  |
| 6:28.219 | דחית |  |
| 6:28.220 | בקנה | , |
| 6:28.221 | לנו |  |
| 6:28.222 | מה |  |
| 6:28.223 | אתה |  |
| 6:28.224 | משיב |  |
| 6:28.225 | אמר |  |
| 6:28.226 | להם |  |
| 6:28.227 | שלשתן |  |
| 6:28.228 | שם |  |
| 6:28.229 | אחד |  |
| 6:28.230 | הן |  |
| 6:28.231 | כאיניש |  |
| 6:28.232 | דאמר |  |
| 6:28.233 | אומנון |  |
| 6:28.234 | בניןן |  |
| 6:28.235 | ארכיטקטנן |  |
| 6:28.236 | נ"א |  |
| 6:28.237 | ארכי |  |
| 6:28.238 | קטונדס |  |
| 6:28.239 | ערוך |  |
| 6:28.240 | אמן | ). |
| 6:28.241 | חזרו |  |
| 6:28.242 | ושאלו |  |
| 6:28.243 | אותו |  |
| 6:28.244 | מהו |  |
| 6:28.245 | דכתיב |  |
| 6:28.246 | כי |  |
| 6:28.247 | אלהים |  |
| 6:28.248 | קדושים |  |
| 6:28.249 | הוא |  |
| 6:28.250 | אמר |  |
| 6:28.251 | להם |  |
| 6:28.252 | קדושים |  |
| 6:28.253 | המה |  |
| 6:28.254 | אין |  |
| 6:28.255 | כתוב |  |
| 6:28.256 | כאן |  |
| 6:28.257 | אלא |  |
| 6:28.258 | הוא |  |
| 6:28.259 | אל |  |
| 6:28.260 | קנא |  |
| 6:28.261 | הוא | , |
| 6:28.262 | אמרו |  |
| 6:28.263 | לו |  |
| 6:28.264 | תלמידיו |  |
| 6:28.265 | רבי |  |
| 6:28.266 | לאלו |  |
| 6:28.267 | דחית |  |
| 6:28.268 | בקנה | . |
| 6:28.269 | ולנו |  |
| 6:28.270 | מה |  |
| 6:28.271 | אתה |  |
| 6:28.272 | משיב | , |
| 6:28.273 | אמר |  |
| 6:28.274 | להם |  |
| 6:28.275 | קדוש |  |
| 6:28.276 | הוא |  |
| 6:28.277 | בכל |  |
| 6:28.278 | מיני |  |
| 6:28.279 | קדושות | , |
| 6:28.280 | דאמר |  |
| 6:28.281 | רבי |  |
| 6:28.282 | יודן |  |
| 6:28.283 | בשם |  |
| 6:28.284 | רבי |  |
| 6:28.285 | אחא |  |
| 6:28.286 | הקב"ה |  |
| 6:28.287 | דרכו |  |
| 6:28.288 | בקדושה | , |
| 6:28.289 | דיבורו |  |
| 6:28.290 | בקדושה | , |
| 6:28.291 | ישובו |  |
| 6:28.292 | בקדושה | , |
| 6:28.293 | אלהים |  |
| 6:28.294 | בקדש |  |
| 6:28.295 | דרכך | , |
| 6:28.296 | הילוכו |  |
| 6:28.297 | בקדושה | , |
| 6:28.298 | הליכות |  |
| 6:28.299 | אל |  |
| 6:28.300 | מלכי |  |
| 6:28.301 | בקדש | , |
| 6:28.302 | מושבו |  |
| 6:28.303 | בקדושה |  |
| 6:28.304 | אלהים |  |
| 6:28.305 | ישב |  |
| 6:28.306 | על |  |
| 6:28.307 | כסא |  |
| 6:28.308 | קדשו | , |
| 6:28.309 | דיבורו |  |
| 6:28.310 | בקדושה |  |
| 6:28.311 | אלהים |  |
| 6:28.312 | דיבר |  |
| 6:28.313 | בקדשו | , |
| 6:28.314 | חשיפת |  |
| 6:28.315 | דרועו |  |
| 6:28.316 | בקדושה | , |
| 6:28.317 | חשף |  |
| 6:28.318 | ה | ' |
| 6:28.319 | את |  |
| 6:28.320 | זרוע |  |
| 6:28.321 | קדשו |  |
| 6:28.322 | נורא |  |
| 6:28.323 | ואדיר |  |
| 6:28.324 | בקדושה | , |
| 6:28.325 | מי |  |
| 6:28.326 | כמוך |  |
| 6:28.327 | נאדר |  |
| 6:28.328 | בקדש | , |
| 6:28.329 | חזרו |  |
| 6:28.330 | ושאלו |  |
| 6:28.331 | אותו |  |
| 6:28.332 | מהו |  |
| 6:28.333 | ההן |  |
| 6:28.334 | דכתיב |  |
| 6:28.335 | מי |  |
| 6:28.336 | גוי |  |
| 6:28.337 | גדול |  |
| 6:28.338 | אשר |  |
| 6:28.339 | לו |  |
| 6:28.340 | אלהים |  |
| 6:28.341 | קרובים |  |
| 6:28.342 | אליו | . |
| 6:28.343 | אמר |  |
| 6:28.344 | להם |  |
| 6:28.345 | כה | ' |
| 6:28.346 | אלהינו |  |
| 6:28.347 | בכל |  |
| 6:28.348 | קראינו |  |
| 6:28.349 | אליהם |  |
| 6:28.350 | אין |  |
| 6:28.351 | כתיב |  |
| 6:28.352 | כאן | , |
| 6:28.353 | אלא |  |
| 6:28.354 | בכל |  |
| 6:28.355 | קראינו |  |
| 6:28.356 | אליו | . |
| 6:28.357 | אמרו |  |
| 6:28.358 | לו |  |
| 6:28.359 | תלמידיו |  |
| 6:28.360 | רבי |  |
| 6:28.361 | לאלו |  |
| 6:28.362 | דחית |  |
| 6:28.363 | בקנה | . |
| 6:28.364 | לנו |  |
| 6:28.365 | מה |  |
| 6:28.366 | אתה |  |
| 6:28.367 | משיב | , |
| 6:28.368 | אמר |  |
| 6:28.369 | להם |  |
| 6:28.370 | קרוב |  |
| 6:28.371 | הוא |  |
| 6:28.372 | בכל |  |
| 6:28.373 | מיני |  |
| 6:28.374 | קריבות |  |
| 6:28.375 | דאמר |  |
| 6:28.376 | ר | ' |
| 6:28.377 | פנחס |  |
| 6:28.378 | בשפ |  |
| 6:28.379 | ר | ' |
| 6:28.380 | יהודה |  |
| 6:28.381 | בר | ' |
| 6:28.382 | סימון | . |
| 6:28.383 | ע"ז |  |
| 6:28.384 | נראית |  |
| 6:28.385 | קרובה |  |
| 6:28.386 | ואינה |  |
| 6:28.387 | אלא |  |
| 6:28.388 | רחוקה | . |
| 6:28.389 | מה |  |
| 6:28.390 | טעם |  |
| 6:28.391 | ישאוהו |  |
| 6:28.392 | על |  |
| 6:28.393 | כתף |  |
| 6:28.394 | יסבלוהו |  |
| 6:28.395 | וגו | ' |
| 6:28.396 | סוף |  |
| 6:28.397 | דבר |  |
| 6:28.398 | אלוהו |  |
| 6:28.399 | עמו |  |
| 6:28.400 | בבית |  |
| 6:28.401 | והוא |  |
| 6:28.402 | צועק |  |
| 6:28.403 | עד |  |
| 6:28.404 | שימות |  |
| 6:28.405 | ולא |  |
| 6:28.406 | ישמע |  |
| 6:28.407 | ולא |  |
| 6:28.408 | יושיע |  |
| 6:28.409 | מצרתו | . |
| 6:28.410 | אבל |  |
| 6:28.411 | הקב"ה |  |
| 6:28.412 | נראה |  |
| 6:28.413 | רחוק |  |
| 6:28.414 | ואין |  |
| 6:28.415 | קרוב |  |
| 6:28.416 | ממנו | , |
| 6:28.417 | דאמר |  |
| 6:28.418 | רבי |  |
| 6:28.419 | לוי |  |
| 6:28.420 | מהארץ |  |
| 6:28.421 | עד |  |
| 6:28.422 | לרקיע |  |
| 6:28.423 | מהלך |  |
| 6:28.424 | חמש |  |
| 6:28.425 | מאות |  |
| 6:28.426 | שנה | . |
| 6:28.427 | וכן |  |
| 6:28.428 | לכל |  |
| 6:28.429 | רקיע |  |
| 6:28.430 | ורקיע | , |
| 6:28.431 | ואמר |  |
| 6:28.432 | רבי |  |
| 6:28.433 | חלבו |  |
| 6:28.434 | ורבי |  |
| 6:28.435 | ברכיה |  |
| 6:28.436 | בשם |  |
| 6:28.437 | רבי |  |
| 6:28.438 | אבא |  |
| 6:28.439 | סמוקא |  |
| 6:28.440 | אף |  |
| 6:28.441 | טלפי |  |
| 6:28.442 | החיות |  |
| 6:28.443 | מהלך |  |
| 6:28.444 | חמש |  |
| 6:28.445 | מאות |  |
| 6:28.446 | חמש |  |
| 6:28.447 | עשרה |  |
| 6:28.448 | מנין |  |
| 6:28.449 | ישרה | , |
| 6:28.450 | ראה |  |
| 6:28.451 | כמה |  |
| 6:28.452 | גבוה |  |
| 6:28.453 | הוא |  |
| 6:28.454 | מעולמו | , |
| 6:28.455 | ואדם |  |
| 6:28.456 | נכנס |  |
| 6:28.457 | לבית |  |
| 6:28.458 | הכנסת |  |
| 6:28.459 | ועומד |  |
| 6:28.460 | אחורי |  |
| 6:28.461 | העמוד |  |
| 6:28.462 | ומתפלל |  |
| 6:28.463 | בלחישה |  |
| 6:28.464 | והקב"ה |  |
| 6:28.465 | מאזין |  |
| 6:28.466 | את |  |
| 6:28.467 | תפלתו |  |
| 6:28.468 | שנאמר |  |
| 6:28.469 | וחנה |  |
| 6:28.470 | היא |  |
| 6:28.471 | מדברת |  |
| 6:28.472 | אל |  |
| 6:28.473 | לבה |  |
| 6:28.474 | רק |  |
| 6:28.475 | שפתיה |  |
| 6:28.476 | נעות |  |
| 6:28.477 | והאזין |  |
| 6:28.478 | הקב"ה |  |
| 6:28.479 | תפלתה | . |
| 6:28.480 | וכן |  |
| 6:28.481 | כל |  |
| 6:28.482 | בריותיו |  |
| 6:28.483 | שנ | ' |
| 6:28.484 | תפלה |  |
| 6:28.485 | לעני |  |
| 6:28.486 | כי |  |
| 6:28.487 | יעטף |  |
| 6:28.488 | כאדם |  |
| 6:28.489 | המשיח |  |
| 6:28.490 | באזן |  |
| 6:28.491 | חבירו |  |
| 6:28.492 | והוא |  |
| 6:28.493 | שומע | , |
| 6:28.494 | וכי |  |
| 6:28.495 | יש |  |
| 6:28.496 | לך |  |
| 6:28.497 | אלוה |  |
| 6:28.498 | קרוב |  |
| 6:28.499 | משה |  |
| 6:28.500 | שהוא |  |
| 6:28.501 | קרוב |  |
| 6:28.502 | לבריותיו |  |
| 6:28.503 | מפה |  |
| 6:28.504 | לאזן |  |
| 6:28.505 | ע"כ | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

