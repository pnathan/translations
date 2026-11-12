# Translation pack: milhamot-hashem perek-8-06 (Milchamot ha-Shem 8:8–8:8)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-06`.

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
- **ובפתרון הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
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


## Milchamot ha-Shem 8:8

Hebrew (Wikisource): קד) הוא שהנביא ירמיה קורא בגרון בלי חשך כשופר מרים קולו בשם ה' לאמר מה מצאו אבותיכם בי עול כי רחקו מעלי כי עברו איי כתיים וראו וקדר שלחו והתבוננו מאד הן היתה כזאת, החמיר גוי אלהיו והמה לא אלהים. ועמי המיר כבודו בלא יועיל. כדגרסינן בפרק קמא דתענית כי שתים רעות עשה עמי אותי עזבו מקור מיים חיים לחצוב להם בורות, בורות נשברים. כי עברו איי כתיים וכו' תנא כתים עובדים לאש וקדרים עובדים למים. ואע"פ שיודעין שהמים מכבין את האש לא הימירו אלהיהם. ועמי המיר כבודו בלא יועיל. הן עתה הימירו המקובלים החדשים את האל הגדול הגבור והנורא אל רחום וחנון ארך אפים וכו' כאשר הודיע דרכיו למשה עבדו ולבני ישראל עלילותיו, באלוה קצר אפים אשר לפי דבריהם אמר עליו הנביא חבקוק מבקש מעתיק לאמר ה' פעלך שהוא זעיר אנפין בקרב שנים חייהו כדאיתא באידרא והובא לעיל סי' מ"ה. ועליו אמר שלמה המלך ע"ה קצר אפים יעשה אולת! אוי לאותה בושה, ואוי לאותה כלמה, נשכבה בבשתינו, ותכסינו כלמתינו. כי נתערבו דעות זרות של עובדי אלילים באמונת תורתינו הקדושה והטהורה. אוי! והדעת הלזו של מקובלים החדשים היא בעצמה אמונת הפילסופים הקדמונים עובדי ע"ז שהזכירו רז"ל במסכת מנחות וכבר הזכרתיה לעיל סי' נ"ט שהגוים מקרגתני כלפי מזרח ומצור כלפי מערב אין מכירין את ישראל ולא את אבינו שבשמים וקרו ליה אלהא דאלהיא! שהאלהים העליון נתן הכח והממשלה לאותן אלוהות שתחתיו כמו שפירש שם מהרש"א ז"ל. והמקובלים אמרו שניתן הכח והממשלה, לעתיק יומין ואריך אנפין ואבא ואימא, והם ג"כ נתנו הכח הנזכר לקצר אפים וזוגתו (זעיר ונוקביה) להשגיח בעולם התחתון. וקלי הדעת אשר קננו בלבם הדעות הכוזבות האלה והאמינו בהם התחילו לבקש להם סמוכות רעועות מן התורה ולהוציא הכתובים ממשמען האמתי לומר כי בראשית שם הוא. ושתי רשויות הן. אחד ברא את השני. בראשית ברא את האלהים שהוא השמים והארץ. זעיר ונוקביה שהם אלהים. ושהאלוה הזה משתבח בעצמו ואומר ה' אמר אלי בני אתה. שהזעיר אנפין משבח את עצמו שהשלטון מלפניו על כל הברואים (זהר בלק דף קצ"א ע"כ). ונתפשטו דרשות של הבל ורעות רוח כאלה על ידי מלך עובדי ע"ז בין האנשים החכמים בעיניהם ההולכים אחר דעתם שלא כדעת חז"ל שהרחיקו דעות אלו וכתבו לתלמי המלך אלהים ברא את בראשית, ועוד אמרו בפרק אין דורשין. שאל רבי ישמעאל את רבי עקיבה כשהיו מהלכין בדרך וכו' כדלעיל סי' כ"ה. כל זה הוצרכו רז"ל להרחיק הדעות הזרות האלה שלא יתערבו בישראל לומר שמים וארץ אלוהות הם.

| id | word | punctuation after |
|---|---|---|
| 8:8.1 | קד | ) |
| 8:8.2 | הוא |  |
| 8:8.3 | שהנביא |  |
| 8:8.4 | ירמיה |  |
| 8:8.5 | קורא |  |
| 8:8.6 | בגרון |  |
| 8:8.7 | בלי |  |
| 8:8.8 | חשך |  |
| 8:8.9 | כשופר |  |
| 8:8.10 | מרים |  |
| 8:8.11 | קולו |  |
| 8:8.12 | בשם |  |
| 8:8.13 | ה | ' |
| 8:8.14 | לאמר |  |
| 8:8.15 | מה |  |
| 8:8.16 | מצאו |  |
| 8:8.17 | אבותיכם |  |
| 8:8.18 | בי |  |
| 8:8.19 | עול |  |
| 8:8.20 | כי |  |
| 8:8.21 | רחקו |  |
| 8:8.22 | מעלי |  |
| 8:8.23 | כי |  |
| 8:8.24 | עברו |  |
| 8:8.25 | איי |  |
| 8:8.26 | כתיים |  |
| 8:8.27 | וראו |  |
| 8:8.28 | וקדר |  |
| 8:8.29 | שלחו |  |
| 8:8.30 | והתבוננו |  |
| 8:8.31 | מאד |  |
| 8:8.32 | הן |  |
| 8:8.33 | היתה |  |
| 8:8.34 | כזאת | , |
| 8:8.35 | החמיר |  |
| 8:8.36 | גוי |  |
| 8:8.37 | אלהיו |  |
| 8:8.38 | והמה |  |
| 8:8.39 | לא |  |
| 8:8.40 | אלהים | . |
| 8:8.41 | ועמי |  |
| 8:8.42 | המיר |  |
| 8:8.43 | כבודו |  |
| 8:8.44 | בלא |  |
| 8:8.45 | יועיל | . |
| 8:8.46 | כדגרסינן |  |
| 8:8.47 | בפרק |  |
| 8:8.48 | קמא |  |
| 8:8.49 | דתענית |  |
| 8:8.50 | כי |  |
| 8:8.51 | שתים |  |
| 8:8.52 | רעות |  |
| 8:8.53 | עשה |  |
| 8:8.54 | עמי |  |
| 8:8.55 | אותי |  |
| 8:8.56 | עזבו |  |
| 8:8.57 | מקור |  |
| 8:8.58 | מיים |  |
| 8:8.59 | חיים |  |
| 8:8.60 | לחצוב |  |
| 8:8.61 | להם |  |
| 8:8.62 | בורות | , |
| 8:8.63 | בורות |  |
| 8:8.64 | נשברים | . |
| 8:8.65 | כי |  |
| 8:8.66 | עברו |  |
| 8:8.67 | איי |  |
| 8:8.68 | כתיים |  |
| 8:8.69 | וכו | ' |
| 8:8.70 | תנא |  |
| 8:8.71 | כתים |  |
| 8:8.72 | עובדים |  |
| 8:8.73 | לאש |  |
| 8:8.74 | וקדרים |  |
| 8:8.75 | עובדים |  |
| 8:8.76 | למים | . |
| 8:8.77 | ואע"פ |  |
| 8:8.78 | שיודעין |  |
| 8:8.79 | שהמים |  |
| 8:8.80 | מכבין |  |
| 8:8.81 | את |  |
| 8:8.82 | האש |  |
| 8:8.83 | לא |  |
| 8:8.84 | הימירו |  |
| 8:8.85 | אלהיהם | . |
| 8:8.86 | ועמי |  |
| 8:8.87 | המיר |  |
| 8:8.88 | כבודו |  |
| 8:8.89 | בלא |  |
| 8:8.90 | יועיל | . |
| 8:8.91 | הן |  |
| 8:8.92 | עתה |  |
| 8:8.93 | הימירו |  |
| 8:8.94 | המקובלים |  |
| 8:8.95 | החדשים |  |
| 8:8.96 | את |  |
| 8:8.97 | האל |  |
| 8:8.98 | הגדול |  |
| 8:8.99 | הגבור |  |
| 8:8.100 | והנורא |  |
| 8:8.101 | אל |  |
| 8:8.102 | רחום |  |
| 8:8.103 | וחנון |  |
| 8:8.104 | ארך |  |
| 8:8.105 | אפים |  |
| 8:8.106 | וכו | ' |
| 8:8.107 | כאשר |  |
| 8:8.108 | הודיע |  |
| 8:8.109 | דרכיו |  |
| 8:8.110 | למשה |  |
| 8:8.111 | עבדו |  |
| 8:8.112 | ולבני |  |
| 8:8.113 | ישראל |  |
| 8:8.114 | עלילותיו | , |
| 8:8.115 | באלוה |  |
| 8:8.116 | קצר |  |
| 8:8.117 | אפים |  |
| 8:8.118 | אשר |  |
| 8:8.119 | לפי |  |
| 8:8.120 | דבריהם |  |
| 8:8.121 | אמר |  |
| 8:8.122 | עליו |  |
| 8:8.123 | הנביא |  |
| 8:8.124 | חבקוק |  |
| 8:8.125 | מבקש |  |
| 8:8.126 | מעתיק |  |
| 8:8.127 | לאמר |  |
| 8:8.128 | ה | ' |
| 8:8.129 | פעלך |  |
| 8:8.130 | שהוא |  |
| 8:8.131 | זעיר |  |
| 8:8.132 | אנפין |  |
| 8:8.133 | בקרב |  |
| 8:8.134 | שנים |  |
| 8:8.135 | חייהו |  |
| 8:8.136 | כדאיתא |  |
| 8:8.137 | באידרא |  |
| 8:8.138 | והובא |  |
| 8:8.139 | לעיל |  |
| 8:8.140 | סי | ' |
| 8:8.141 | מ"ה | . |
| 8:8.142 | ועליו |  |
| 8:8.143 | אמר |  |
| 8:8.144 | שלמה |  |
| 8:8.145 | המלך |  |
| 8:8.146 | ע"ה |  |
| 8:8.147 | קצר |  |
| 8:8.148 | אפים |  |
| 8:8.149 | יעשה |  |
| 8:8.150 | אולת | ! |
| 8:8.151 | אוי |  |
| 8:8.152 | לאותה |  |
| 8:8.153 | בושה | , |
| 8:8.154 | ואוי |  |
| 8:8.155 | לאותה |  |
| 8:8.156 | כלמה | , |
| 8:8.157 | נשכבה |  |
| 8:8.158 | בבשתינו | , |
| 8:8.159 | ותכסינו |  |
| 8:8.160 | כלמתינו | . |
| 8:8.161 | כי |  |
| 8:8.162 | נתערבו |  |
| 8:8.163 | דעות |  |
| 8:8.164 | זרות |  |
| 8:8.165 | של |  |
| 8:8.166 | עובדי |  |
| 8:8.167 | אלילים |  |
| 8:8.168 | באמונת |  |
| 8:8.169 | תורתינו |  |
| 8:8.170 | הקדושה |  |
| 8:8.171 | והטהורה | . |
| 8:8.172 | אוי | ! |
| 8:8.173 | והדעת |  |
| 8:8.174 | הלזו |  |
| 8:8.175 | של |  |
| 8:8.176 | מקובלים |  |
| 8:8.177 | החדשים |  |
| 8:8.178 | היא |  |
| 8:8.179 | בעצמה |  |
| 8:8.180 | אמונת |  |
| 8:8.181 | הפילסופים |  |
| 8:8.182 | הקדמונים |  |
| 8:8.183 | עובדי |  |
| 8:8.184 | ע"ז |  |
| 8:8.185 | שהזכירו |  |
| 8:8.186 | רז"ל |  |
| 8:8.187 | במסכת |  |
| 8:8.188 | מנחות |  |
| 8:8.189 | וכבר |  |
| 8:8.190 | הזכרתיה |  |
| 8:8.191 | לעיל |  |
| 8:8.192 | סי | ' |
| 8:8.193 | נ"ט |  |
| 8:8.194 | שהגוים |  |
| 8:8.195 | מקרגתני |  |
| 8:8.196 | כלפי |  |
| 8:8.197 | מזרח |  |
| 8:8.198 | ומצור |  |
| 8:8.199 | כלפי |  |
| 8:8.200 | מערב |  |
| 8:8.201 | אין |  |
| 8:8.202 | מכירין |  |
| 8:8.203 | את |  |
| 8:8.204 | ישראל |  |
| 8:8.205 | ולא |  |
| 8:8.206 | את |  |
| 8:8.207 | אבינו |  |
| 8:8.208 | שבשמים |  |
| 8:8.209 | וקרו |  |
| 8:8.210 | ליה |  |
| 8:8.211 | אלהא |  |
| 8:8.212 | דאלהיא | ! |
| 8:8.213 | שהאלהים |  |
| 8:8.214 | העליון |  |
| 8:8.215 | נתן |  |
| 8:8.216 | הכח |  |
| 8:8.217 | והממשלה |  |
| 8:8.218 | לאותן |  |
| 8:8.219 | אלוהות |  |
| 8:8.220 | שתחתיו |  |
| 8:8.221 | כמו |  |
| 8:8.222 | שפירש |  |
| 8:8.223 | שם |  |
| 8:8.224 | מהרש"א |  |
| 8:8.225 | ז"ל | . |
| 8:8.226 | והמקובלים |  |
| 8:8.227 | אמרו |  |
| 8:8.228 | שניתן |  |
| 8:8.229 | הכח |  |
| 8:8.230 | והממשלה | , |
| 8:8.231 | לעתיק |  |
| 8:8.232 | יומין |  |
| 8:8.233 | ואריך |  |
| 8:8.234 | אנפין |  |
| 8:8.235 | ואבא |  |
| 8:8.236 | ואימא | , |
| 8:8.237 | והם |  |
| 8:8.238 | ג"כ |  |
| 8:8.239 | נתנו |  |
| 8:8.240 | הכח |  |
| 8:8.241 | הנזכר |  |
| 8:8.242 | לקצר |  |
| 8:8.243 | אפים |  |
| 8:8.244 | וזוגתו |  |
| 8:8.245 | זעיר |  |
| 8:8.246 | ונוקביה | ) |
| 8:8.247 | להשגיח |  |
| 8:8.248 | בעולם |  |
| 8:8.249 | התחתון | . |
| 8:8.250 | וקלי |  |
| 8:8.251 | הדעת |  |
| 8:8.252 | אשר |  |
| 8:8.253 | קננו |  |
| 8:8.254 | בלבם |  |
| 8:8.255 | הדעות |  |
| 8:8.256 | הכוזבות |  |
| 8:8.257 | האלה |  |
| 8:8.258 | והאמינו |  |
| 8:8.259 | בהם |  |
| 8:8.260 | התחילו |  |
| 8:8.261 | לבקש |  |
| 8:8.262 | להם |  |
| 8:8.263 | סמוכות |  |
| 8:8.264 | רעועות |  |
| 8:8.265 | מן |  |
| 8:8.266 | התורה |  |
| 8:8.267 | ולהוציא |  |
| 8:8.268 | הכתובים |  |
| 8:8.269 | ממשמען |  |
| 8:8.270 | האמתי |  |
| 8:8.271 | לומר |  |
| 8:8.272 | כי |  |
| 8:8.273 | בראשית |  |
| 8:8.274 | שם |  |
| 8:8.275 | הוא | . |
| 8:8.276 | ושתי |  |
| 8:8.277 | רשויות |  |
| 8:8.278 | הן | . |
| 8:8.279 | אחד |  |
| 8:8.280 | ברא |  |
| 8:8.281 | את |  |
| 8:8.282 | השני | . |
| 8:8.283 | בראשית |  |
| 8:8.284 | ברא |  |
| 8:8.285 | את |  |
| 8:8.286 | האלהים |  |
| 8:8.287 | שהוא |  |
| 8:8.288 | השמים |  |
| 8:8.289 | והארץ | . |
| 8:8.290 | זעיר |  |
| 8:8.291 | ונוקביה |  |
| 8:8.292 | שהם |  |
| 8:8.293 | אלהים | . |
| 8:8.294 | ושהאלוה |  |
| 8:8.295 | הזה |  |
| 8:8.296 | משתבח |  |
| 8:8.297 | בעצמו |  |
| 8:8.298 | ואומר |  |
| 8:8.299 | ה | ' |
| 8:8.300 | אמר |  |
| 8:8.301 | אלי |  |
| 8:8.302 | בני |  |
| 8:8.303 | אתה | . |
| 8:8.304 | שהזעיר |  |
| 8:8.305 | אנפין |  |
| 8:8.306 | משבח |  |
| 8:8.307 | את |  |
| 8:8.308 | עצמו |  |
| 8:8.309 | שהשלטון |  |
| 8:8.310 | מלפניו |  |
| 8:8.311 | על |  |
| 8:8.312 | כל |  |
| 8:8.313 | הברואים |  |
| 8:8.314 | זהר |  |
| 8:8.315 | בלק |  |
| 8:8.316 | דף |  |
| 8:8.317 | קצ"א |  |
| 8:8.318 | ע"כ | ). |
| 8:8.319 | ונתפשטו |  |
| 8:8.320 | דרשות |  |
| 8:8.321 | של |  |
| 8:8.322 | הבל |  |
| 8:8.323 | ורעות |  |
| 8:8.324 | רוח |  |
| 8:8.325 | כאלה |  |
| 8:8.326 | על |  |
| 8:8.327 | ידי |  |
| 8:8.328 | מלך |  |
| 8:8.329 | עובדי |  |
| 8:8.330 | ע"ז |  |
| 8:8.331 | בין |  |
| 8:8.332 | האנשים |  |
| 8:8.333 | החכמים |  |
| 8:8.334 | בעיניהם |  |
| 8:8.335 | ההולכים |  |
| 8:8.336 | אחר |  |
| 8:8.337 | דעתם |  |
| 8:8.338 | שלא |  |
| 8:8.339 | כדעת |  |
| 8:8.340 | חז"ל |  |
| 8:8.341 | שהרחיקו |  |
| 8:8.342 | דעות |  |
| 8:8.343 | אלו |  |
| 8:8.344 | וכתבו |  |
| 8:8.345 | לתלמי |  |
| 8:8.346 | המלך |  |
| 8:8.347 | אלהים |  |
| 8:8.348 | ברא |  |
| 8:8.349 | את |  |
| 8:8.350 | בראשית | , |
| 8:8.351 | ועוד |  |
| 8:8.352 | אמרו |  |
| 8:8.353 | בפרק |  |
| 8:8.354 | אין |  |
| 8:8.355 | דורשין | . |
| 8:8.356 | שאל |  |
| 8:8.357 | רבי |  |
| 8:8.358 | ישמעאל |  |
| 8:8.359 | את |  |
| 8:8.360 | רבי |  |
| 8:8.361 | עקיבה |  |
| 8:8.362 | כשהיו |  |
| 8:8.363 | מהלכין |  |
| 8:8.364 | בדרך |  |
| 8:8.365 | וכו | ' |
| 8:8.366 | כדלעיל |  |
| 8:8.367 | סי | ' |
| 8:8.368 | כ"ה | . |
| 8:8.369 | כל |  |
| 8:8.370 | זה |  |
| 8:8.371 | הוצרכו |  |
| 8:8.372 | רז"ל |  |
| 8:8.373 | להרחיק |  |
| 8:8.374 | הדעות |  |
| 8:8.375 | הזרות |  |
| 8:8.376 | האלה |  |
| 8:8.377 | שלא |  |
| 8:8.378 | יתערבו |  |
| 8:8.379 | בישראל |  |
| 8:8.380 | לומר |  |
| 8:8.381 | שמים |  |
| 8:8.382 | וארץ |  |
| 8:8.383 | אלוהות |  |
| 8:8.384 | הם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

