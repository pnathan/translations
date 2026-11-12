# Translation pack: milhamot-hashem perek-3-03 (Milchamot ha-Shem 3:5–3:6)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-03.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-03`.

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
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **עבודה / לעבוד / לעבד**: fixed, "worship / to worship", for God and idols,
  in prose and verses ("ואותו תעבוד" → "and Him you shall worship"); the one
  exception: Exod 12:26 "מה העבודה הזאת לכם" → "this service" (the Passover
  rite, note). "serve" only for a human master. An ambiguous לעבד / לעבד כ-:
  by verse or context, the other in `readings`.
- **עבודה זרה / ע"ז / עכו"ם / עבודה אחרת**: עבודה זרה → "idolatry" (note once:
  "strange worship"); the tractate "Avodah Zarah"; עובדי ע"ז → "worshippers of
  idols". עכו"ם → "idolaters" (note: "worshippers of stars and
  constellations"), never "gentiles" or "Christians", even of a Christian king
  or the Trinity. עבודה אחרת → "other worship" (note first in a unit). אליל →
  "idol"; האלילה → "the she-idol"; אל זר → "a strange god"; אל אחר → "another
  god"; אלהים אחרים → "other gods".
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
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
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


## Milchamot ha-Shem 3:5

Hebrew (Wikisource): ובפעם השלישית היתה תשובתכם כשאלת אחד מארבעה בנים מה העבודה הזאת לכם, לכם ולא לו, כאילו לא נצטוינו בתורה וידעת היום והשבות אל לבבך כי ה' הוא האלהים וכו' וכמה מקראות באו בתורה ובנביאים ובכתובים על זה וכל שכן בדברי רז"ל. גם חזרת להיות מכת מספרי לשון הרע וממקבליו באמרך שכבר יש עלינו כמה עדים שאנחנו מבזים את החכמים ולא נוכל להכחיש עוד וכו'. והן ידוע ומפורסם שאין מקבלין עדות אלא בפני בעל דין. ומן התורה אחד דיני ממונות ואחד דיני נפשות בדרישה וחקירה כמו שכ' הרמב"ם ז"ל בפרק ג' מהלכות עדות אבל אמרו חכמים כדי שלא תנעול דלת בפני לוין אין עדי ממון צריכים דרישה וחקירה וכו'. במה דברים אמורים בהודאות והלואות, מתנות ומכירות, אבל בדיני קנסות צריכים דרישה וחקירה, ואין צריך לומר במלקיות ובגלות, ובדין מרומה אף בדיני ממונות צריך דרישה וחקירה.

| id | word | punctuation after |
|---|---|---|
| 3:5.1 | ובפעם |  |
| 3:5.2 | השלישית |  |
| 3:5.3 | היתה |  |
| 3:5.4 | תשובתכם |  |
| 3:5.5 | כשאלת |  |
| 3:5.6 | אחד |  |
| 3:5.7 | מארבעה |  |
| 3:5.8 | בנים |  |
| 3:5.9 | מה |  |
| 3:5.10 | העבודה |  |
| 3:5.11 | הזאת |  |
| 3:5.12 | לכם | , |
| 3:5.13 | לכם |  |
| 3:5.14 | ולא |  |
| 3:5.15 | לו | , |
| 3:5.16 | כאילו |  |
| 3:5.17 | לא |  |
| 3:5.18 | נצטוינו |  |
| 3:5.19 | בתורה |  |
| 3:5.20 | וידעת |  |
| 3:5.21 | היום |  |
| 3:5.22 | והשבות |  |
| 3:5.23 | אל |  |
| 3:5.24 | לבבך |  |
| 3:5.25 | כי |  |
| 3:5.26 | ה | ' |
| 3:5.27 | הוא |  |
| 3:5.28 | האלהים |  |
| 3:5.29 | וכו | ' |
| 3:5.30 | וכמה |  |
| 3:5.31 | מקראות |  |
| 3:5.32 | באו |  |
| 3:5.33 | בתורה |  |
| 3:5.34 | ובנביאים |  |
| 3:5.35 | ובכתובים |  |
| 3:5.36 | על |  |
| 3:5.37 | זה |  |
| 3:5.38 | וכל |  |
| 3:5.39 | שכן |  |
| 3:5.40 | בדברי |  |
| 3:5.41 | רז"ל | . |
| 3:5.42 | גם |  |
| 3:5.43 | חזרת |  |
| 3:5.44 | להיות |  |
| 3:5.45 | מכת |  |
| 3:5.46 | מספרי |  |
| 3:5.47 | לשון |  |
| 3:5.48 | הרע |  |
| 3:5.49 | וממקבליו |  |
| 3:5.50 | באמרך |  |
| 3:5.51 | שכבר |  |
| 3:5.52 | יש |  |
| 3:5.53 | עלינו |  |
| 3:5.54 | כמה |  |
| 3:5.55 | עדים |  |
| 3:5.56 | שאנחנו |  |
| 3:5.57 | מבזים |  |
| 3:5.58 | את |  |
| 3:5.59 | החכמים |  |
| 3:5.60 | ולא |  |
| 3:5.61 | נוכל |  |
| 3:5.62 | להכחיש |  |
| 3:5.63 | עוד |  |
| 3:5.64 | וכו | '. |
| 3:5.65 | והן |  |
| 3:5.66 | ידוע |  |
| 3:5.67 | ומפורסם |  |
| 3:5.68 | שאין |  |
| 3:5.69 | מקבלין |  |
| 3:5.70 | עדות |  |
| 3:5.71 | אלא |  |
| 3:5.72 | בפני |  |
| 3:5.73 | בעל |  |
| 3:5.74 | דין | . |
| 3:5.75 | ומן |  |
| 3:5.76 | התורה |  |
| 3:5.77 | אחד |  |
| 3:5.78 | דיני |  |
| 3:5.79 | ממונות |  |
| 3:5.80 | ואחד |  |
| 3:5.81 | דיני |  |
| 3:5.82 | נפשות |  |
| 3:5.83 | בדרישה |  |
| 3:5.84 | וחקירה |  |
| 3:5.85 | כמו |  |
| 3:5.86 | שכ | ' |
| 3:5.87 | הרמב"ם |  |
| 3:5.88 | ז"ל |  |
| 3:5.89 | בפרק |  |
| 3:5.90 | ג | ' |
| 3:5.91 | מהלכות |  |
| 3:5.92 | עדות |  |
| 3:5.93 | אבל |  |
| 3:5.94 | אמרו |  |
| 3:5.95 | חכמים |  |
| 3:5.96 | כדי |  |
| 3:5.97 | שלא |  |
| 3:5.98 | תנעול |  |
| 3:5.99 | דלת |  |
| 3:5.100 | בפני |  |
| 3:5.101 | לוין |  |
| 3:5.102 | אין |  |
| 3:5.103 | עדי |  |
| 3:5.104 | ממון |  |
| 3:5.105 | צריכים |  |
| 3:5.106 | דרישה |  |
| 3:5.107 | וחקירה |  |
| 3:5.108 | וכו | '. |
| 3:5.109 | במה |  |
| 3:5.110 | דברים |  |
| 3:5.111 | אמורים |  |
| 3:5.112 | בהודאות |  |
| 3:5.113 | והלואות | , |
| 3:5.114 | מתנות |  |
| 3:5.115 | ומכירות | , |
| 3:5.116 | אבל |  |
| 3:5.117 | בדיני |  |
| 3:5.118 | קנסות |  |
| 3:5.119 | צריכים |  |
| 3:5.120 | דרישה |  |
| 3:5.121 | וחקירה | , |
| 3:5.122 | ואין |  |
| 3:5.123 | צריך |  |
| 3:5.124 | לומר |  |
| 3:5.125 | במלקיות |  |
| 3:5.126 | ובגלות | , |
| 3:5.127 | ובדין |  |
| 3:5.128 | מרומה |  |
| 3:5.129 | אף |  |
| 3:5.130 | בדיני |  |
| 3:5.131 | ממונות |  |
| 3:5.132 | צריך |  |
| 3:5.133 | דרישה |  |
| 3:5.134 | וחקירה | . |


## Milchamot ha-Shem 3:6

Hebrew (Wikisource): ו) ובנדון דידן כולהו איתינהו ביה, ואיך מלאך לבך לקבל לשון הרע מאנשים סכלים ההולכים יחפים כמנודים לשמים, ולחשוב אותה לעדות גמורה, להחזיק את עמיתך בתורה ובמצות דאוריתא ודרבנן לאפיקורסים ומינים שהם מכלל אותם שמורידין ולא מעלין? וידוע הוא ומפורסם, כי זה הדור, דור שוא ושקר, והיא בנשיאי ישראל, וכל אחד מתכבד בקלון חברו ורבו אשר העמידו בקרו אורה זו תורה. ובמשימי לילות כימים ועבדי לתורה שלא תשתכח מישראל כפי יכולתם שפכתם בוז, ותשימו את לומדי התורה ומחזיקים בה נבזים ושפלים ככלי אין חפץ בו. וכל עושי רע טוב בעיניכם, ומכובד יותר מהם, ועל ידי דבות שקר שמתם לומדי תורה לשמצה בעיני כל רואיהם עד שבאו למדרגה פחותה מאד, שעמי הארץ משתבחים בעצמם לומר אשרינו מה טוב חלקנו וערבה שנתינו עד עלות השמש וכבוד ננחל יותר מאלה העומדים בבית ה' בלילות, לקיים מה שנאמר ואם בחקותי תמאסו, אם במלמדי חקותי תמאסו, כמו שכ הרמב"ם ז"ל בהל' תלמוד תורה.

| id | word | punctuation after |
|---|---|---|
| 3:6.1 | ו | ) |
| 3:6.2 | ובנדון |  |
| 3:6.3 | דידן |  |
| 3:6.4 | כולהו |  |
| 3:6.5 | איתינהו |  |
| 3:6.6 | ביה | , |
| 3:6.7 | ואיך |  |
| 3:6.8 | מלאך |  |
| 3:6.9 | לבך |  |
| 3:6.10 | לקבל |  |
| 3:6.11 | לשון |  |
| 3:6.12 | הרע |  |
| 3:6.13 | מאנשים |  |
| 3:6.14 | סכלים |  |
| 3:6.15 | ההולכים |  |
| 3:6.16 | יחפים |  |
| 3:6.17 | כמנודים |  |
| 3:6.18 | לשמים | , |
| 3:6.19 | ולחשוב |  |
| 3:6.20 | אותה |  |
| 3:6.21 | לעדות |  |
| 3:6.22 | גמורה | , |
| 3:6.23 | להחזיק |  |
| 3:6.24 | את |  |
| 3:6.25 | עמיתך |  |
| 3:6.26 | בתורה |  |
| 3:6.27 | ובמצות |  |
| 3:6.28 | דאוריתא |  |
| 3:6.29 | ודרבנן |  |
| 3:6.30 | לאפיקורסים |  |
| 3:6.31 | ומינים |  |
| 3:6.32 | שהם |  |
| 3:6.33 | מכלל |  |
| 3:6.34 | אותם |  |
| 3:6.35 | שמורידין |  |
| 3:6.36 | ולא |  |
| 3:6.37 | מעלין | ? |
| 3:6.38 | וידוע |  |
| 3:6.39 | הוא |  |
| 3:6.40 | ומפורסם | , |
| 3:6.41 | כי |  |
| 3:6.42 | זה |  |
| 3:6.43 | הדור | , |
| 3:6.44 | דור |  |
| 3:6.45 | שוא |  |
| 3:6.46 | ושקר | , |
| 3:6.47 | והיא |  |
| 3:6.48 | בנשיאי |  |
| 3:6.49 | ישראל | , |
| 3:6.50 | וכל |  |
| 3:6.51 | אחד |  |
| 3:6.52 | מתכבד |  |
| 3:6.53 | בקלון |  |
| 3:6.54 | חברו |  |
| 3:6.55 | ורבו |  |
| 3:6.56 | אשר |  |
| 3:6.57 | העמידו |  |
| 3:6.58 | בקרו |  |
| 3:6.59 | אורה |  |
| 3:6.60 | זו |  |
| 3:6.61 | תורה | . |
| 3:6.62 | ובמשימי |  |
| 3:6.63 | לילות |  |
| 3:6.64 | כימים |  |
| 3:6.65 | ועבדי |  |
| 3:6.66 | לתורה |  |
| 3:6.67 | שלא |  |
| 3:6.68 | תשתכח |  |
| 3:6.69 | מישראל |  |
| 3:6.70 | כפי |  |
| 3:6.71 | יכולתם |  |
| 3:6.72 | שפכתם |  |
| 3:6.73 | בוז | , |
| 3:6.74 | ותשימו |  |
| 3:6.75 | את |  |
| 3:6.76 | לומדי |  |
| 3:6.77 | התורה |  |
| 3:6.78 | ומחזיקים |  |
| 3:6.79 | בה |  |
| 3:6.80 | נבזים |  |
| 3:6.81 | ושפלים |  |
| 3:6.82 | ככלי |  |
| 3:6.83 | אין |  |
| 3:6.84 | חפץ |  |
| 3:6.85 | בו | . |
| 3:6.86 | וכל |  |
| 3:6.87 | עושי |  |
| 3:6.88 | רע |  |
| 3:6.89 | טוב |  |
| 3:6.90 | בעיניכם | , |
| 3:6.91 | ומכובד |  |
| 3:6.92 | יותר |  |
| 3:6.93 | מהם | , |
| 3:6.94 | ועל |  |
| 3:6.95 | ידי |  |
| 3:6.96 | דבות |  |
| 3:6.97 | שקר |  |
| 3:6.98 | שמתם |  |
| 3:6.99 | לומדי |  |
| 3:6.100 | תורה |  |
| 3:6.101 | לשמצה |  |
| 3:6.102 | בעיני |  |
| 3:6.103 | כל |  |
| 3:6.104 | רואיהם |  |
| 3:6.105 | עד |  |
| 3:6.106 | שבאו |  |
| 3:6.107 | למדרגה |  |
| 3:6.108 | פחותה |  |
| 3:6.109 | מאד | , |
| 3:6.110 | שעמי |  |
| 3:6.111 | הארץ |  |
| 3:6.112 | משתבחים |  |
| 3:6.113 | בעצמם |  |
| 3:6.114 | לומר |  |
| 3:6.115 | אשרינו |  |
| 3:6.116 | מה |  |
| 3:6.117 | טוב |  |
| 3:6.118 | חלקנו |  |
| 3:6.119 | וערבה |  |
| 3:6.120 | שנתינו |  |
| 3:6.121 | עד |  |
| 3:6.122 | עלות |  |
| 3:6.123 | השמש |  |
| 3:6.124 | וכבוד |  |
| 3:6.125 | ננחל |  |
| 3:6.126 | יותר |  |
| 3:6.127 | מאלה |  |
| 3:6.128 | העומדים |  |
| 3:6.129 | בבית |  |
| 3:6.130 | ה | ' |
| 3:6.131 | בלילות | , |
| 3:6.132 | לקיים |  |
| 3:6.133 | מה |  |
| 3:6.134 | שנאמר |  |
| 3:6.135 | ואם |  |
| 3:6.136 | בחקותי |  |
| 3:6.137 | תמאסו | , |
| 3:6.138 | אם |  |
| 3:6.139 | במלמדי |  |
| 3:6.140 | חקותי |  |
| 3:6.141 | תמאסו | , |
| 3:6.142 | כמו |  |
| 3:6.143 | שכ |  |
| 3:6.144 | הרמב"ם |  |
| 3:6.145 | ז"ל |  |
| 3:6.146 | בהל | ' |
| 3:6.147 | תלמוד |  |
| 3:6.148 | תורה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

