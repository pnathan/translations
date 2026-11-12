# Translation pack: milhamot-hashem perek-8-09 (Milchamot ha-Shem 8:11–8:11)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-09`.

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
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **תשבורת** → "mathematics"; **ותוכנים** → "astronomy".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
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


## Milchamot ha-Shem 8:11

Hebrew (Wikisource): ומ"מ אפשר שלא ירדו לעמקה לדעת את מרורות הפתנים אשר בקרבה, רק עשו כאשר עשינו אנחנו שנים רבות שפירשנו ממנה מפני שקרא עליה ערער. ולא דברנו אודותיה כמצווה עלינו להודיע את חקי האלהים ואת תורותיו ולמי ראוי לעבוד! ולכן חרה אף השי"ת בנו, ויענישינו על מריינו, שהיינו שומעים דברים המביאים לידי טעות, ועל שלא מחינו בכבוד קונינו. הסית השי"ת בנו את ארך הלשון אשר כחזקתו גבה לבו עד להשחית. וישיאו לבו ויצרו הרע להתגולל להוציא דבה רעה על לומדי תורה לשמה על אמתתה במשנה ובתלמוד. שהם מינים גמורים ושבבית שני תיקן רבן גמליאל על ידי שמואל הקטן ברכת המינים, ושכל התלמידים הלומדים משנה והרמב"ם יצאו לתרבות רעה על שאין לומדים זהר ותיקונים! ומיום ליום יביע אומר מלשינות ושקר לפני שרי הממשלה, שהמה אלה מכתבי עמל. ורצונם למרוד בממשלה הרוממה ולחסות תחת ממשלה אנגליז, ולהמריד את יושבי הארץ לחסות בממשלה אחרת נכריה ועל ידי רכילותו זאת בקש הואלי (המושל) לנתשם מעל אדמתם אל ארץ אחרת, ויתן אותנו בבית הסהר יותר מחדש ימים כשלשים איש, ויפל פחד ורעדה בלב העם מהדר גאונו ועריצותו! ושוב המציא במזימותיו רכילות אחרת במכתבים אל האימאם יחיא אלמתוכל יר"ה. שקצת מן היהודים רוצים להתנשא על המושלמנים וללבוש בגדי כבוד, ולרכוב על סוסים. ומבקשים בתי ספר ללמוד שפה תורכית ושאר לשונות כדי להתגדל על האיסלאם. ואלמלא אני מלמדם ומדריכם בדרכי כניעת היהדות ואיך הם מחוייבים לקבל עליהם עול האיסלאם. כי עתה פרצו ועלו על האיסלאם. ועל ידי זה זכה ליתן לו רשיון מן האימאם יר"ה לרדות בעם ה' אלה ולעשות כחפצו, אך לימודי שפת תוגרמה וידיעת גלילות הארץ וגבולותיה וחשבון ותשבורת וכיוצא לא אסר אותן האימאם יר"ה, כי הוא איש משכיל ונבון, ובמכתבו כתב שאין למנוע הלימודים האלה מנערי בני ישראל. ובעל הלשון הרחיב עוד את פיו לפני אחד השופטים של האימאם בפנינו ג"כ להלשין עלינו. כי אנחנו עוזבי דת הישראלית, ולומדים בספרי פילסופי יון, ומכחישים אמונת מציאות השדים. ובאמת לא נמצא אצלינו שום ספר מספרי יון, ולא זולתם מספרי הפילסופים החיצונים כלל! ורק זדון לבו השיאו להלשין עלינו שקר, והרבה שנאת חנם במחנה העברים על לא דבר! ויוסף עוד לחטוא לשלח מכתבים אל כל כפרי ועיירות התימן להוציא ולפרסם את הרבה הרעה הזאת עלינו בשם מינים! וערום יערים לעשות נוסחי דבות רעות לאמר, הן כבר עשינו משפט ברשעים האלה ונביא אותם אל בית הכלא, ועליכם המצוה למהר במכתבי קובלנא אל האימאם יר"ה, כי אלה מקלקלי דתינו ומפסידי אמונתינו, מהרו, חושו, חושו, להגדיל המדורה עליהם! ולא היה כותבם בכתב ידו, רק מוסרם לילדים מקרוביו כדי להעתיקם בכתב ידם, ושולחם לכל ערי התימן שלא בכתב ידו, כדי שיאמר לא פעלתי און.

| id | word | punctuation after |
|---|---|---|
| 8:11.1 | ומ"מ |  |
| 8:11.2 | אפשר |  |
| 8:11.3 | שלא |  |
| 8:11.4 | ירדו |  |
| 8:11.5 | לעמקה |  |
| 8:11.6 | לדעת |  |
| 8:11.7 | את |  |
| 8:11.8 | מרורות |  |
| 8:11.9 | הפתנים |  |
| 8:11.10 | אשר |  |
| 8:11.11 | בקרבה | , |
| 8:11.12 | רק |  |
| 8:11.13 | עשו |  |
| 8:11.14 | כאשר |  |
| 8:11.15 | עשינו |  |
| 8:11.16 | אנחנו |  |
| 8:11.17 | שנים |  |
| 8:11.18 | רבות |  |
| 8:11.19 | שפירשנו |  |
| 8:11.20 | ממנה |  |
| 8:11.21 | מפני |  |
| 8:11.22 | שקרא |  |
| 8:11.23 | עליה |  |
| 8:11.24 | ערער | . |
| 8:11.25 | ולא |  |
| 8:11.26 | דברנו |  |
| 8:11.27 | אודותיה |  |
| 8:11.28 | כמצווה |  |
| 8:11.29 | עלינו |  |
| 8:11.30 | להודיע |  |
| 8:11.31 | את |  |
| 8:11.32 | חקי |  |
| 8:11.33 | האלהים |  |
| 8:11.34 | ואת |  |
| 8:11.35 | תורותיו |  |
| 8:11.36 | ולמי |  |
| 8:11.37 | ראוי |  |
| 8:11.38 | לעבוד | ! |
| 8:11.39 | ולכן |  |
| 8:11.40 | חרה |  |
| 8:11.41 | אף |  |
| 8:11.42 | השי"ת |  |
| 8:11.43 | בנו | , |
| 8:11.44 | ויענישינו |  |
| 8:11.45 | על |  |
| 8:11.46 | מריינו | , |
| 8:11.47 | שהיינו |  |
| 8:11.48 | שומעים |  |
| 8:11.49 | דברים |  |
| 8:11.50 | המביאים |  |
| 8:11.51 | לידי |  |
| 8:11.52 | טעות | , |
| 8:11.53 | ועל |  |
| 8:11.54 | שלא |  |
| 8:11.55 | מחינו |  |
| 8:11.56 | בכבוד |  |
| 8:11.57 | קונינו | . |
| 8:11.58 | הסית |  |
| 8:11.59 | השי"ת |  |
| 8:11.60 | בנו |  |
| 8:11.61 | את |  |
| 8:11.62 | ארך |  |
| 8:11.63 | הלשון |  |
| 8:11.64 | אשר |  |
| 8:11.65 | כחזקתו |  |
| 8:11.66 | גבה |  |
| 8:11.67 | לבו |  |
| 8:11.68 | עד |  |
| 8:11.69 | להשחית | . |
| 8:11.70 | וישיאו |  |
| 8:11.71 | לבו |  |
| 8:11.72 | ויצרו |  |
| 8:11.73 | הרע |  |
| 8:11.74 | להתגולל |  |
| 8:11.75 | להוציא |  |
| 8:11.76 | דבה |  |
| 8:11.77 | רעה |  |
| 8:11.78 | על |  |
| 8:11.79 | לומדי |  |
| 8:11.80 | תורה |  |
| 8:11.81 | לשמה |  |
| 8:11.82 | על |  |
| 8:11.83 | אמתתה |  |
| 8:11.84 | במשנה |  |
| 8:11.85 | ובתלמוד | . |
| 8:11.86 | שהם |  |
| 8:11.87 | מינים |  |
| 8:11.88 | גמורים |  |
| 8:11.89 | ושבבית |  |
| 8:11.90 | שני |  |
| 8:11.91 | תיקן |  |
| 8:11.92 | רבן |  |
| 8:11.93 | גמליאל |  |
| 8:11.94 | על |  |
| 8:11.95 | ידי |  |
| 8:11.96 | שמואל |  |
| 8:11.97 | הקטן |  |
| 8:11.98 | ברכת |  |
| 8:11.99 | המינים | , |
| 8:11.100 | ושכל |  |
| 8:11.101 | התלמידים |  |
| 8:11.102 | הלומדים |  |
| 8:11.103 | משנה |  |
| 8:11.104 | והרמב"ם |  |
| 8:11.105 | יצאו |  |
| 8:11.106 | לתרבות |  |
| 8:11.107 | רעה |  |
| 8:11.108 | על |  |
| 8:11.109 | שאין |  |
| 8:11.110 | לומדים |  |
| 8:11.111 | זהר |  |
| 8:11.112 | ותיקונים | ! |
| 8:11.113 | ומיום |  |
| 8:11.114 | ליום |  |
| 8:11.115 | יביע |  |
| 8:11.116 | אומר |  |
| 8:11.117 | מלשינות |  |
| 8:11.118 | ושקר |  |
| 8:11.119 | לפני |  |
| 8:11.120 | שרי |  |
| 8:11.121 | הממשלה | , |
| 8:11.122 | שהמה |  |
| 8:11.123 | אלה |  |
| 8:11.124 | מכתבי |  |
| 8:11.125 | עמל | . |
| 8:11.126 | ורצונם |  |
| 8:11.127 | למרוד |  |
| 8:11.128 | בממשלה |  |
| 8:11.129 | הרוממה |  |
| 8:11.130 | ולחסות |  |
| 8:11.131 | תחת |  |
| 8:11.132 | ממשלה |  |
| 8:11.133 | אנגליז | , |
| 8:11.134 | ולהמריד |  |
| 8:11.135 | את |  |
| 8:11.136 | יושבי |  |
| 8:11.137 | הארץ |  |
| 8:11.138 | לחסות |  |
| 8:11.139 | בממשלה |  |
| 8:11.140 | אחרת |  |
| 8:11.141 | נכריה |  |
| 8:11.142 | ועל |  |
| 8:11.143 | ידי |  |
| 8:11.144 | רכילותו |  |
| 8:11.145 | זאת |  |
| 8:11.146 | בקש |  |
| 8:11.147 | הואלי |  |
| 8:11.148 | המושל | ) |
| 8:11.149 | לנתשם |  |
| 8:11.150 | מעל |  |
| 8:11.151 | אדמתם |  |
| 8:11.152 | אל |  |
| 8:11.153 | ארץ |  |
| 8:11.154 | אחרת | , |
| 8:11.155 | ויתן |  |
| 8:11.156 | אותנו |  |
| 8:11.157 | בבית |  |
| 8:11.158 | הסהר |  |
| 8:11.159 | יותר |  |
| 8:11.160 | מחדש |  |
| 8:11.161 | ימים |  |
| 8:11.162 | כשלשים |  |
| 8:11.163 | איש | , |
| 8:11.164 | ויפל |  |
| 8:11.165 | פחד |  |
| 8:11.166 | ורעדה |  |
| 8:11.167 | בלב |  |
| 8:11.168 | העם |  |
| 8:11.169 | מהדר |  |
| 8:11.170 | גאונו |  |
| 8:11.171 | ועריצותו | ! |
| 8:11.172 | ושוב |  |
| 8:11.173 | המציא |  |
| 8:11.174 | במזימותיו |  |
| 8:11.175 | רכילות |  |
| 8:11.176 | אחרת |  |
| 8:11.177 | במכתבים |  |
| 8:11.178 | אל |  |
| 8:11.179 | האימאם |  |
| 8:11.180 | יחיא |  |
| 8:11.181 | אלמתוכל |  |
| 8:11.182 | יר"ה | . |
| 8:11.183 | שקצת |  |
| 8:11.184 | מן |  |
| 8:11.185 | היהודים |  |
| 8:11.186 | רוצים |  |
| 8:11.187 | להתנשא |  |
| 8:11.188 | על |  |
| 8:11.189 | המושלמנים |  |
| 8:11.190 | וללבוש |  |
| 8:11.191 | בגדי |  |
| 8:11.192 | כבוד | , |
| 8:11.193 | ולרכוב |  |
| 8:11.194 | על |  |
| 8:11.195 | סוסים | . |
| 8:11.196 | ומבקשים |  |
| 8:11.197 | בתי |  |
| 8:11.198 | ספר |  |
| 8:11.199 | ללמוד |  |
| 8:11.200 | שפה |  |
| 8:11.201 | תורכית |  |
| 8:11.202 | ושאר |  |
| 8:11.203 | לשונות |  |
| 8:11.204 | כדי |  |
| 8:11.205 | להתגדל |  |
| 8:11.206 | על |  |
| 8:11.207 | האיסלאם | . |
| 8:11.208 | ואלמלא |  |
| 8:11.209 | אני |  |
| 8:11.210 | מלמדם |  |
| 8:11.211 | ומדריכם |  |
| 8:11.212 | בדרכי |  |
| 8:11.213 | כניעת |  |
| 8:11.214 | היהדות |  |
| 8:11.215 | ואיך |  |
| 8:11.216 | הם |  |
| 8:11.217 | מחוייבים |  |
| 8:11.218 | לקבל |  |
| 8:11.219 | עליהם |  |
| 8:11.220 | עול |  |
| 8:11.221 | האיסלאם | . |
| 8:11.222 | כי |  |
| 8:11.223 | עתה |  |
| 8:11.224 | פרצו |  |
| 8:11.225 | ועלו |  |
| 8:11.226 | על |  |
| 8:11.227 | האיסלאם | . |
| 8:11.228 | ועל |  |
| 8:11.229 | ידי |  |
| 8:11.230 | זה |  |
| 8:11.231 | זכה |  |
| 8:11.232 | ליתן |  |
| 8:11.233 | לו |  |
| 8:11.234 | רשיון |  |
| 8:11.235 | מן |  |
| 8:11.236 | האימאם |  |
| 8:11.237 | יר"ה |  |
| 8:11.238 | לרדות |  |
| 8:11.239 | בעם |  |
| 8:11.240 | ה | ' |
| 8:11.241 | אלה |  |
| 8:11.242 | ולעשות |  |
| 8:11.243 | כחפצו | , |
| 8:11.244 | אך |  |
| 8:11.245 | לימודי |  |
| 8:11.246 | שפת |  |
| 8:11.247 | תוגרמה |  |
| 8:11.248 | וידיעת |  |
| 8:11.249 | גלילות |  |
| 8:11.250 | הארץ |  |
| 8:11.251 | וגבולותיה |  |
| 8:11.252 | וחשבון |  |
| 8:11.253 | ותשבורת |  |
| 8:11.254 | וכיוצא |  |
| 8:11.255 | לא |  |
| 8:11.256 | אסר |  |
| 8:11.257 | אותן |  |
| 8:11.258 | האימאם |  |
| 8:11.259 | יר"ה | , |
| 8:11.260 | כי |  |
| 8:11.261 | הוא |  |
| 8:11.262 | איש |  |
| 8:11.263 | משכיל |  |
| 8:11.264 | ונבון | , |
| 8:11.265 | ובמכתבו |  |
| 8:11.266 | כתב |  |
| 8:11.267 | שאין |  |
| 8:11.268 | למנוע |  |
| 8:11.269 | הלימודים |  |
| 8:11.270 | האלה |  |
| 8:11.271 | מנערי |  |
| 8:11.272 | בני |  |
| 8:11.273 | ישראל | . |
| 8:11.274 | ובעל |  |
| 8:11.275 | הלשון |  |
| 8:11.276 | הרחיב |  |
| 8:11.277 | עוד |  |
| 8:11.278 | את |  |
| 8:11.279 | פיו |  |
| 8:11.280 | לפני |  |
| 8:11.281 | אחד |  |
| 8:11.282 | השופטים |  |
| 8:11.283 | של |  |
| 8:11.284 | האימאם |  |
| 8:11.285 | בפנינו |  |
| 8:11.286 | ג"כ |  |
| 8:11.287 | להלשין |  |
| 8:11.288 | עלינו | . |
| 8:11.289 | כי |  |
| 8:11.290 | אנחנו |  |
| 8:11.291 | עוזבי |  |
| 8:11.292 | דת |  |
| 8:11.293 | הישראלית | , |
| 8:11.294 | ולומדים |  |
| 8:11.295 | בספרי |  |
| 8:11.296 | פילסופי |  |
| 8:11.297 | יון | , |
| 8:11.298 | ומכחישים |  |
| 8:11.299 | אמונת |  |
| 8:11.300 | מציאות |  |
| 8:11.301 | השדים | . |
| 8:11.302 | ובאמת |  |
| 8:11.303 | לא |  |
| 8:11.304 | נמצא |  |
| 8:11.305 | אצלינו |  |
| 8:11.306 | שום |  |
| 8:11.307 | ספר |  |
| 8:11.308 | מספרי |  |
| 8:11.309 | יון | , |
| 8:11.310 | ולא |  |
| 8:11.311 | זולתם |  |
| 8:11.312 | מספרי |  |
| 8:11.313 | הפילסופים |  |
| 8:11.314 | החיצונים |  |
| 8:11.315 | כלל | ! |
| 8:11.316 | ורק |  |
| 8:11.317 | זדון |  |
| 8:11.318 | לבו |  |
| 8:11.319 | השיאו |  |
| 8:11.320 | להלשין |  |
| 8:11.321 | עלינו |  |
| 8:11.322 | שקר | , |
| 8:11.323 | והרבה |  |
| 8:11.324 | שנאת |  |
| 8:11.325 | חנם |  |
| 8:11.326 | במחנה |  |
| 8:11.327 | העברים |  |
| 8:11.328 | על |  |
| 8:11.329 | לא |  |
| 8:11.330 | דבר | ! |
| 8:11.331 | ויוסף |  |
| 8:11.332 | עוד |  |
| 8:11.333 | לחטוא |  |
| 8:11.334 | לשלח |  |
| 8:11.335 | מכתבים |  |
| 8:11.336 | אל |  |
| 8:11.337 | כל |  |
| 8:11.338 | כפרי |  |
| 8:11.339 | ועיירות |  |
| 8:11.340 | התימן |  |
| 8:11.341 | להוציא |  |
| 8:11.342 | ולפרסם |  |
| 8:11.343 | את |  |
| 8:11.344 | הרבה |  |
| 8:11.345 | הרעה |  |
| 8:11.346 | הזאת |  |
| 8:11.347 | עלינו |  |
| 8:11.348 | בשם |  |
| 8:11.349 | מינים | ! |
| 8:11.350 | וערום |  |
| 8:11.351 | יערים |  |
| 8:11.352 | לעשות |  |
| 8:11.353 | נוסחי |  |
| 8:11.354 | דבות |  |
| 8:11.355 | רעות |  |
| 8:11.356 | לאמר | , |
| 8:11.357 | הן |  |
| 8:11.358 | כבר |  |
| 8:11.359 | עשינו |  |
| 8:11.360 | משפט |  |
| 8:11.361 | ברשעים |  |
| 8:11.362 | האלה |  |
| 8:11.363 | ונביא |  |
| 8:11.364 | אותם |  |
| 8:11.365 | אל |  |
| 8:11.366 | בית |  |
| 8:11.367 | הכלא | , |
| 8:11.368 | ועליכם |  |
| 8:11.369 | המצוה |  |
| 8:11.370 | למהר |  |
| 8:11.371 | במכתבי |  |
| 8:11.372 | קובלנא |  |
| 8:11.373 | אל |  |
| 8:11.374 | האימאם |  |
| 8:11.375 | יר"ה | , |
| 8:11.376 | כי |  |
| 8:11.377 | אלה |  |
| 8:11.378 | מקלקלי |  |
| 8:11.379 | דתינו |  |
| 8:11.380 | ומפסידי |  |
| 8:11.381 | אמונתינו | , |
| 8:11.382 | מהרו | , |
| 8:11.383 | חושו | , |
| 8:11.384 | חושו | , |
| 8:11.385 | להגדיל |  |
| 8:11.386 | המדורה |  |
| 8:11.387 | עליהם | ! |
| 8:11.388 | ולא |  |
| 8:11.389 | היה |  |
| 8:11.390 | כותבם |  |
| 8:11.391 | בכתב |  |
| 8:11.392 | ידו | , |
| 8:11.393 | רק |  |
| 8:11.394 | מוסרם |  |
| 8:11.395 | לילדים |  |
| 8:11.396 | מקרוביו |  |
| 8:11.397 | כדי |  |
| 8:11.398 | להעתיקם |  |
| 8:11.399 | בכתב |  |
| 8:11.400 | ידם | , |
| 8:11.401 | ושולחם |  |
| 8:11.402 | לכל |  |
| 8:11.403 | ערי |  |
| 8:11.404 | התימן |  |
| 8:11.405 | שלא |  |
| 8:11.406 | בכתב |  |
| 8:11.407 | ידו | , |
| 8:11.408 | כדי |  |
| 8:11.409 | שיאמר |  |
| 8:11.410 | לא |  |
| 8:11.411 | פעלתי |  |
| 8:11.412 | און | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

