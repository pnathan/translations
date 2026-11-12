# Translation pack: milhamot-hashem perek-7-13 (Milchamot ha-Shem 7:18–7:18)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-13.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-13`.

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
- **עשתונות איש ירושלים** → "the thoughts of the man of Jerusalem" (note); no
  identification.
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

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
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
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


## Milchamot ha-Shem 7:18

Hebrew (Wikisource): ורז"ל אמרו בסנהדרין כשעלה משה למרום מצאו להקב"ה שיושב וכותב ארך אפים. אמר לפניו רבונו של עולם לצדיקים? אמר לו אף לרשעים. אמר לפניו רשעים יאבדו אמר ליה השתא חזית דמבעי לך. בשעה שחטאו ישראל אמרו לו לא כך אמרת ועתה יגדל נא כח ה' כאשר דברת לאמר ה' ארך אפים וכו'. וכן כל הנביאים מתארים לשי"ת בתאר ארך אפים. דוד המלך ע"ה גם במזמור פ"ו אמר ואתה ה' אל רחום וחנון ארך אפים ורב חסד ואמת, ובמזמור קמ"ה אמר חנון ורחום ה' ארך אפים וגדל חסד. הנביא נחום אמר ה' ארך אפים וגדל כח ונקה לא ינקה. יונה אמר בתפלתו כי ידעתי כי אתה אל חנון ורחום ארך אפים ורב חסד וניחם על הרעה. יואל אמר ושובו של ה' אלהיכם כי חנון ורחום הוא ארך אפים ורב חסד וניחם על הרעה, אנשי כנסת הגדולה אמרו בתפלתם ואתה אלוה סליחות חנון ורחום ארך אפים ורב חסד ולא עזבתם (נחמיה ט׳:י״ז) ושלמה המלך ע"ה אמר טוב ארך רוח מגבה רוח (קהלת ז׳:ח׳), ופירש רש"י על ארך רוח המאריך רוגזו ואינו ממהר לריב. ובמשלי אמר קצר אפים יעשה אולת, ואיש מזמות ישנא, ופירש רש"י קצר אפים הממהר לנקום כעסו, ואיש מזמות. מחשבות עצה רעה ישנא, ובפי' קב ונקי פי' קצר אפים הוא היפך ארך אפים, האדם שהוא נמהר לשפוך חמתו ולנקום נקמתו יעשה כל מעשיו באולת כסילות, ועיין עוד במפרשים ובס' הנז' משבח ארך אפים ואמר ארך אפים רב תבונות, וקצר רוח מרים אולת, ופירשו המפרשים מי שהוא ארך אפים שאינו ממהר לכעוס והוא מושל ברוחו הוא רב תבונות, ומי שהוא קצר אפים וקצר רוח שלא ימשול ברוחו הוא מרים האולת מן הארץ, ומגביהה ומניחה על ראשו, ומגלה אותה לאשר לא ידעה, או יאמר שהוא מעורר האולת. וי"מ מרים אולת שהוא גורם לכסילים להרים בו יד. כלומר למרוד בו ע"כ. ושם סי' ט"ו איש חמה יגרה מדון, וארך אפים ישקיט ריב, ופי' מי שהוא בעל חמה דהיינו קצר אפים הוא יבעיר המדון והקטטה ויגדילנו אע"פ שהריב קטן יעלנו ויגדילנו בחמתו, והאדם המאריך אפו ישקיט ריב ויניחנו. ובסי' י"ו טוב ארך אפים מגבור ומושל ברוחו מלוכד עיר וכ' המפרשים שטוב ארך אפים המושל ביצרו וכובש תאותו מלהנקם, מן הגבור המתנקם באויבו. כי הכובש את יצרו, אויבו שהוא היצר מבפנים, והוא מתגבר עליו וכובשו, וזה הגבור אויבו מבחוץ והוא מתגבר על האויב החיצון, והמושל ברוחו הוא מושל על האויב הפנימי. המורם מכל האמור, כי תואר קצר אפים הוא תואר אולת וכסילות, וכבר נודע הדבר אשר דברו רז"ל בעל חובת הלבבות ור"י הלוי והרמב"ם במורה פרק כ"ו ומ"ו ומ"ז לראשון ורס"ג בהאמונות והדעות מאמר האחדות ושבילי אמונה וזולתם, שלא יתואר השי"ת אלא בתארים נעלים, המורים על השלימות, להורות שהור ית' שלם בכל מיני שלימות. ולא יתואר בשום תואר חסרון או העדר, כל שכן בתאר פחיתות וסכילות וסכלות.

| id | word | punctuation after |
|---|---|---|
| 7:18.1 | ורז"ל |  |
| 7:18.2 | אמרו |  |
| 7:18.3 | בסנהדרין |  |
| 7:18.4 | כשעלה |  |
| 7:18.5 | משה |  |
| 7:18.6 | למרום |  |
| 7:18.7 | מצאו |  |
| 7:18.8 | להקב"ה |  |
| 7:18.9 | שיושב |  |
| 7:18.10 | וכותב |  |
| 7:18.11 | ארך |  |
| 7:18.12 | אפים | . |
| 7:18.13 | אמר |  |
| 7:18.14 | לפניו |  |
| 7:18.15 | רבונו |  |
| 7:18.16 | של |  |
| 7:18.17 | עולם |  |
| 7:18.18 | לצדיקים | ? |
| 7:18.19 | אמר |  |
| 7:18.20 | לו |  |
| 7:18.21 | אף |  |
| 7:18.22 | לרשעים | . |
| 7:18.23 | אמר |  |
| 7:18.24 | לפניו |  |
| 7:18.25 | רשעים |  |
| 7:18.26 | יאבדו |  |
| 7:18.27 | אמר |  |
| 7:18.28 | ליה |  |
| 7:18.29 | השתא |  |
| 7:18.30 | חזית |  |
| 7:18.31 | דמבעי |  |
| 7:18.32 | לך | . |
| 7:18.33 | בשעה |  |
| 7:18.34 | שחטאו |  |
| 7:18.35 | ישראל |  |
| 7:18.36 | אמרו |  |
| 7:18.37 | לו |  |
| 7:18.38 | לא |  |
| 7:18.39 | כך |  |
| 7:18.40 | אמרת |  |
| 7:18.41 | ועתה |  |
| 7:18.42 | יגדל |  |
| 7:18.43 | נא |  |
| 7:18.44 | כח |  |
| 7:18.45 | ה | ' |
| 7:18.46 | כאשר |  |
| 7:18.47 | דברת |  |
| 7:18.48 | לאמר |  |
| 7:18.49 | ה | ' |
| 7:18.50 | ארך |  |
| 7:18.51 | אפים |  |
| 7:18.52 | וכו | '. |
| 7:18.53 | וכן |  |
| 7:18.54 | כל |  |
| 7:18.55 | הנביאים |  |
| 7:18.56 | מתארים |  |
| 7:18.57 | לשי"ת |  |
| 7:18.58 | בתאר |  |
| 7:18.59 | ארך |  |
| 7:18.60 | אפים | . |
| 7:18.61 | דוד |  |
| 7:18.62 | המלך |  |
| 7:18.63 | ע"ה |  |
| 7:18.64 | גם |  |
| 7:18.65 | במזמור |  |
| 7:18.66 | פ"ו |  |
| 7:18.67 | אמר |  |
| 7:18.68 | ואתה |  |
| 7:18.69 | ה | ' |
| 7:18.70 | אל |  |
| 7:18.71 | רחום |  |
| 7:18.72 | וחנון |  |
| 7:18.73 | ארך |  |
| 7:18.74 | אפים |  |
| 7:18.75 | ורב |  |
| 7:18.76 | חסד |  |
| 7:18.77 | ואמת | , |
| 7:18.78 | ובמזמור |  |
| 7:18.79 | קמ"ה |  |
| 7:18.80 | אמר |  |
| 7:18.81 | חנון |  |
| 7:18.82 | ורחום |  |
| 7:18.83 | ה | ' |
| 7:18.84 | ארך |  |
| 7:18.85 | אפים |  |
| 7:18.86 | וגדל |  |
| 7:18.87 | חסד | . |
| 7:18.88 | הנביא |  |
| 7:18.89 | נחום |  |
| 7:18.90 | אמר |  |
| 7:18.91 | ה | ' |
| 7:18.92 | ארך |  |
| 7:18.93 | אפים |  |
| 7:18.94 | וגדל |  |
| 7:18.95 | כח |  |
| 7:18.96 | ונקה |  |
| 7:18.97 | לא |  |
| 7:18.98 | ינקה | . |
| 7:18.99 | יונה |  |
| 7:18.100 | אמר |  |
| 7:18.101 | בתפלתו |  |
| 7:18.102 | כי |  |
| 7:18.103 | ידעתי |  |
| 7:18.104 | כי |  |
| 7:18.105 | אתה |  |
| 7:18.106 | אל |  |
| 7:18.107 | חנון |  |
| 7:18.108 | ורחום |  |
| 7:18.109 | ארך |  |
| 7:18.110 | אפים |  |
| 7:18.111 | ורב |  |
| 7:18.112 | חסד |  |
| 7:18.113 | וניחם |  |
| 7:18.114 | על |  |
| 7:18.115 | הרעה | . |
| 7:18.116 | יואל |  |
| 7:18.117 | אמר |  |
| 7:18.118 | ושובו |  |
| 7:18.119 | של |  |
| 7:18.120 | ה | ' |
| 7:18.121 | אלהיכם |  |
| 7:18.122 | כי |  |
| 7:18.123 | חנון |  |
| 7:18.124 | ורחום |  |
| 7:18.125 | הוא |  |
| 7:18.126 | ארך |  |
| 7:18.127 | אפים |  |
| 7:18.128 | ורב |  |
| 7:18.129 | חסד |  |
| 7:18.130 | וניחם |  |
| 7:18.131 | על |  |
| 7:18.132 | הרעה | , |
| 7:18.133 | אנשי |  |
| 7:18.134 | כנסת |  |
| 7:18.135 | הגדולה |  |
| 7:18.136 | אמרו |  |
| 7:18.137 | בתפלתם |  |
| 7:18.138 | ואתה |  |
| 7:18.139 | אלוה |  |
| 7:18.140 | סליחות |  |
| 7:18.141 | חנון |  |
| 7:18.142 | ורחום |  |
| 7:18.143 | ארך |  |
| 7:18.144 | אפים |  |
| 7:18.145 | ורב |  |
| 7:18.146 | חסד |  |
| 7:18.147 | ולא |  |
| 7:18.148 | עזבתם |  |
| 7:18.149 | נחמיה |  |
| 7:18.150 | ט׳:י״ז | ) |
| 7:18.151 | ושלמה |  |
| 7:18.152 | המלך |  |
| 7:18.153 | ע"ה |  |
| 7:18.154 | אמר |  |
| 7:18.155 | טוב |  |
| 7:18.156 | ארך |  |
| 7:18.157 | רוח |  |
| 7:18.158 | מגבה |  |
| 7:18.159 | רוח |  |
| 7:18.160 | קהלת |  |
| 7:18.161 | ז׳:ח׳ | ), |
| 7:18.162 | ופירש |  |
| 7:18.163 | רש"י |  |
| 7:18.164 | על |  |
| 7:18.165 | ארך |  |
| 7:18.166 | רוח |  |
| 7:18.167 | המאריך |  |
| 7:18.168 | רוגזו |  |
| 7:18.169 | ואינו |  |
| 7:18.170 | ממהר |  |
| 7:18.171 | לריב | . |
| 7:18.172 | ובמשלי |  |
| 7:18.173 | אמר |  |
| 7:18.174 | קצר |  |
| 7:18.175 | אפים |  |
| 7:18.176 | יעשה |  |
| 7:18.177 | אולת | , |
| 7:18.178 | ואיש |  |
| 7:18.179 | מזמות |  |
| 7:18.180 | ישנא | , |
| 7:18.181 | ופירש |  |
| 7:18.182 | רש"י |  |
| 7:18.183 | קצר |  |
| 7:18.184 | אפים |  |
| 7:18.185 | הממהר |  |
| 7:18.186 | לנקום |  |
| 7:18.187 | כעסו | , |
| 7:18.188 | ואיש |  |
| 7:18.189 | מזמות | . |
| 7:18.190 | מחשבות |  |
| 7:18.191 | עצה |  |
| 7:18.192 | רעה |  |
| 7:18.193 | ישנא | , |
| 7:18.194 | ובפי | ' |
| 7:18.195 | קב |  |
| 7:18.196 | ונקי |  |
| 7:18.197 | פי | ' |
| 7:18.198 | קצר |  |
| 7:18.199 | אפים |  |
| 7:18.200 | הוא |  |
| 7:18.201 | היפך |  |
| 7:18.202 | ארך |  |
| 7:18.203 | אפים | , |
| 7:18.204 | האדם |  |
| 7:18.205 | שהוא |  |
| 7:18.206 | נמהר |  |
| 7:18.207 | לשפוך |  |
| 7:18.208 | חמתו |  |
| 7:18.209 | ולנקום |  |
| 7:18.210 | נקמתו |  |
| 7:18.211 | יעשה |  |
| 7:18.212 | כל |  |
| 7:18.213 | מעשיו |  |
| 7:18.214 | באולת |  |
| 7:18.215 | כסילות | , |
| 7:18.216 | ועיין |  |
| 7:18.217 | עוד |  |
| 7:18.218 | במפרשים |  |
| 7:18.219 | ובס | ' |
| 7:18.220 | הנז | ' |
| 7:18.221 | משבח |  |
| 7:18.222 | ארך |  |
| 7:18.223 | אפים |  |
| 7:18.224 | ואמר |  |
| 7:18.225 | ארך |  |
| 7:18.226 | אפים |  |
| 7:18.227 | רב |  |
| 7:18.228 | תבונות | , |
| 7:18.229 | וקצר |  |
| 7:18.230 | רוח |  |
| 7:18.231 | מרים |  |
| 7:18.232 | אולת | , |
| 7:18.233 | ופירשו |  |
| 7:18.234 | המפרשים |  |
| 7:18.235 | מי |  |
| 7:18.236 | שהוא |  |
| 7:18.237 | ארך |  |
| 7:18.238 | אפים |  |
| 7:18.239 | שאינו |  |
| 7:18.240 | ממהר |  |
| 7:18.241 | לכעוס |  |
| 7:18.242 | והוא |  |
| 7:18.243 | מושל |  |
| 7:18.244 | ברוחו |  |
| 7:18.245 | הוא |  |
| 7:18.246 | רב |  |
| 7:18.247 | תבונות | , |
| 7:18.248 | ומי |  |
| 7:18.249 | שהוא |  |
| 7:18.250 | קצר |  |
| 7:18.251 | אפים |  |
| 7:18.252 | וקצר |  |
| 7:18.253 | רוח |  |
| 7:18.254 | שלא |  |
| 7:18.255 | ימשול |  |
| 7:18.256 | ברוחו |  |
| 7:18.257 | הוא |  |
| 7:18.258 | מרים |  |
| 7:18.259 | האולת |  |
| 7:18.260 | מן |  |
| 7:18.261 | הארץ | , |
| 7:18.262 | ומגביהה |  |
| 7:18.263 | ומניחה |  |
| 7:18.264 | על |  |
| 7:18.265 | ראשו | , |
| 7:18.266 | ומגלה |  |
| 7:18.267 | אותה |  |
| 7:18.268 | לאשר |  |
| 7:18.269 | לא |  |
| 7:18.270 | ידעה | , |
| 7:18.271 | או |  |
| 7:18.272 | יאמר |  |
| 7:18.273 | שהוא |  |
| 7:18.274 | מעורר |  |
| 7:18.275 | האולת | . |
| 7:18.276 | וי"מ |  |
| 7:18.277 | מרים |  |
| 7:18.278 | אולת |  |
| 7:18.279 | שהוא |  |
| 7:18.280 | גורם |  |
| 7:18.281 | לכסילים |  |
| 7:18.282 | להרים |  |
| 7:18.283 | בו |  |
| 7:18.284 | יד | . |
| 7:18.285 | כלומר |  |
| 7:18.286 | למרוד |  |
| 7:18.287 | בו |  |
| 7:18.288 | ע"כ | . |
| 7:18.289 | ושם |  |
| 7:18.290 | סי | ' |
| 7:18.291 | ט"ו |  |
| 7:18.292 | איש |  |
| 7:18.293 | חמה |  |
| 7:18.294 | יגרה |  |
| 7:18.295 | מדון | , |
| 7:18.296 | וארך |  |
| 7:18.297 | אפים |  |
| 7:18.298 | ישקיט |  |
| 7:18.299 | ריב | , |
| 7:18.300 | ופי | ' |
| 7:18.301 | מי |  |
| 7:18.302 | שהוא |  |
| 7:18.303 | בעל |  |
| 7:18.304 | חמה |  |
| 7:18.305 | דהיינו |  |
| 7:18.306 | קצר |  |
| 7:18.307 | אפים |  |
| 7:18.308 | הוא |  |
| 7:18.309 | יבעיר |  |
| 7:18.310 | המדון |  |
| 7:18.311 | והקטטה |  |
| 7:18.312 | ויגדילנו |  |
| 7:18.313 | אע"פ |  |
| 7:18.314 | שהריב |  |
| 7:18.315 | קטן |  |
| 7:18.316 | יעלנו |  |
| 7:18.317 | ויגדילנו |  |
| 7:18.318 | בחמתו | , |
| 7:18.319 | והאדם |  |
| 7:18.320 | המאריך |  |
| 7:18.321 | אפו |  |
| 7:18.322 | ישקיט |  |
| 7:18.323 | ריב |  |
| 7:18.324 | ויניחנו | . |
| 7:18.325 | ובסי | ' |
| 7:18.326 | י"ו |  |
| 7:18.327 | טוב |  |
| 7:18.328 | ארך |  |
| 7:18.329 | אפים |  |
| 7:18.330 | מגבור |  |
| 7:18.331 | ומושל |  |
| 7:18.332 | ברוחו |  |
| 7:18.333 | מלוכד |  |
| 7:18.334 | עיר |  |
| 7:18.335 | וכ | ' |
| 7:18.336 | המפרשים |  |
| 7:18.337 | שטוב |  |
| 7:18.338 | ארך |  |
| 7:18.339 | אפים |  |
| 7:18.340 | המושל |  |
| 7:18.341 | ביצרו |  |
| 7:18.342 | וכובש |  |
| 7:18.343 | תאותו |  |
| 7:18.344 | מלהנקם | , |
| 7:18.345 | מן |  |
| 7:18.346 | הגבור |  |
| 7:18.347 | המתנקם |  |
| 7:18.348 | באויבו | . |
| 7:18.349 | כי |  |
| 7:18.350 | הכובש |  |
| 7:18.351 | את |  |
| 7:18.352 | יצרו | , |
| 7:18.353 | אויבו |  |
| 7:18.354 | שהוא |  |
| 7:18.355 | היצר |  |
| 7:18.356 | מבפנים | , |
| 7:18.357 | והוא |  |
| 7:18.358 | מתגבר |  |
| 7:18.359 | עליו |  |
| 7:18.360 | וכובשו | , |
| 7:18.361 | וזה |  |
| 7:18.362 | הגבור |  |
| 7:18.363 | אויבו |  |
| 7:18.364 | מבחוץ |  |
| 7:18.365 | והוא |  |
| 7:18.366 | מתגבר |  |
| 7:18.367 | על |  |
| 7:18.368 | האויב |  |
| 7:18.369 | החיצון | , |
| 7:18.370 | והמושל |  |
| 7:18.371 | ברוחו |  |
| 7:18.372 | הוא |  |
| 7:18.373 | מושל |  |
| 7:18.374 | על |  |
| 7:18.375 | האויב |  |
| 7:18.376 | הפנימי | . |
| 7:18.377 | המורם |  |
| 7:18.378 | מכל |  |
| 7:18.379 | האמור | , |
| 7:18.380 | כי |  |
| 7:18.381 | תואר |  |
| 7:18.382 | קצר |  |
| 7:18.383 | אפים |  |
| 7:18.384 | הוא |  |
| 7:18.385 | תואר |  |
| 7:18.386 | אולת |  |
| 7:18.387 | וכסילות | , |
| 7:18.388 | וכבר |  |
| 7:18.389 | נודע |  |
| 7:18.390 | הדבר |  |
| 7:18.391 | אשר |  |
| 7:18.392 | דברו |  |
| 7:18.393 | רז"ל |  |
| 7:18.394 | בעל |  |
| 7:18.395 | חובת |  |
| 7:18.396 | הלבבות |  |
| 7:18.397 | ור"י |  |
| 7:18.398 | הלוי |  |
| 7:18.399 | והרמב"ם |  |
| 7:18.400 | במורה |  |
| 7:18.401 | פרק |  |
| 7:18.402 | כ"ו |  |
| 7:18.403 | ומ"ו |  |
| 7:18.404 | ומ"ז |  |
| 7:18.405 | לראשון |  |
| 7:18.406 | ורס"ג |  |
| 7:18.407 | בהאמונות |  |
| 7:18.408 | והדעות |  |
| 7:18.409 | מאמר |  |
| 7:18.410 | האחדות |  |
| 7:18.411 | ושבילי |  |
| 7:18.412 | אמונה |  |
| 7:18.413 | וזולתם | , |
| 7:18.414 | שלא |  |
| 7:18.415 | יתואר |  |
| 7:18.416 | השי"ת |  |
| 7:18.417 | אלא |  |
| 7:18.418 | בתארים |  |
| 7:18.419 | נעלים | , |
| 7:18.420 | המורים |  |
| 7:18.421 | על |  |
| 7:18.422 | השלימות | , |
| 7:18.423 | להורות |  |
| 7:18.424 | שהור |  |
| 7:18.425 | ית | ' |
| 7:18.426 | שלם |  |
| 7:18.427 | בכל |  |
| 7:18.428 | מיני |  |
| 7:18.429 | שלימות | . |
| 7:18.430 | ולא |  |
| 7:18.431 | יתואר |  |
| 7:18.432 | בשום |  |
| 7:18.433 | תואר |  |
| 7:18.434 | חסרון |  |
| 7:18.435 | או |  |
| 7:18.436 | העדר | , |
| 7:18.437 | כל |  |
| 7:18.438 | שכן |  |
| 7:18.439 | בתאר |  |
| 7:18.440 | פחיתות |  |
| 7:18.441 | וסכילות |  |
| 7:18.442 | וסכלות | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

