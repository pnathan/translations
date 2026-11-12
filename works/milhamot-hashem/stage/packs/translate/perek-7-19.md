# Translation pack: milhamot-hashem perek-7-19 (Milchamot ha-Shem 7:24–7:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-7-19.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-7-19`.

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
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
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
- **שיתוף / שתוף / לשתף / משותף**: associating another with God → fixed,
  "associate / association" ("כל המשתף שם שמים ודבר אחר" → "whoever associates
  the name of Heaven with something else"). The Maimonidean term for a word said of God and creatures →
  "equivocation / equivocal" ("בשתוף השם בלבד" → "by equivocation of the term
  only"), never "association", Guide I in the note; "shared" only for a thing
  held in common. צירוף: "conjunction" of God; "permutation" in letter
  mysticism.
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
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
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


## Milchamot ha-Shem 7:24

Hebrew (Wikisource): צו) גם הרס"ג בפירושו לס' איוב מרחיק מאד סברת האומר שבני האלהים מלאכים וביזה וזלזל את האומר שהם מלאכים כמו שעשה רשב"י בבראשית רבה וז"ל בלשון ערבי. ומא אבעד אלקיאס אן יכון הדא אלשטן מלכא, לאן אהל אלתוחיד מגמעין עלי' אן אלבארי נל ועז כלק מלאיכתה מקרבין. מן עלמה אנחם לא יעצון כמא אנה אצטפא אנביאה אלמרסלין מן עלמה אנהם לא יעצונה. ומעאנדה אלאוליא אלצאלחין מן כבאיר אלמעצי! פמן זעם אם הדא אלמעאנד לאיוב כאן מלכא פקד נעל אלמלאיכה יעצון. וכאלף מא נחן מגמעין עליה, וקד תעלם ארשדך אללה אן לאלנפס תלאת קוי, אלפכר, ואלגצב ואלשהות, פאמא קווה אלפכר פחי אלתר ימייז כתא אל אנסאן אלצואב מן אלכטא ואלחק מן אלבאטל, ואמא קווה אלגצב פהי אלתי כהא ילחק אלחיואן אלאקדאם ואלגראות ואלאנתקאם ואסתיפא אלטואיל, ואמא קוה אלשהוה פבהא ישתאק אלחיואן אלי אלגדא ואלי אלתלקיח וסאיר אלדעאוי. ותעלם מע דאלך אן צפה עדל אלנאטק אן תכון אלקוותן אעני קווה אלשהוה ואלגצב מדענה לאלפכר (פאן) גלב אליהא אחדהמא כאן דאלך גורא וכטאא, פמן גאוז אן תפעל אלמלאיכה אלחסד אלדי הו אחד אגואע אלגצב לזמה תגויז וקוע אלשהואת מן אלמלאיכה והו אלאכל ואלשרב ואלנכאח וסאיר אלאמור אלרדלה אלוסלה אלתי קד צאן אללה מלאכיתה ענהא, סימא וקד סבק פי עלמה אנהם לא יעצונה. פלא יגוז אן יכון באלפעל כלאף מא סבק פי עלמה וכו' וקד בלגני ען בעץ אלכוארג אן כאלפה בה אלי' אן זעם אן הדא אלשטן מלך. ואן אלחסד ואלגירה ואלבני יגוז וקועהם באלפעל מן אלמלאיכה כמה תוהם פי קצח' איוב וחסד אלשטן לה' תם ראא איצא אן באטלאקה עליהם דאלך ילזמה אטלאק אל שיאן ואלנבאח, פנאוז המא איצא. ואדעא אן קצה ויראו בני האלהים את בנות האדם כי טבת הנה. הם אמלאיכה אלנאכחון, אלחרם באסקאטהם מן אלסמא ותציירהם כאלנאס, ונצב להם שבחה הו קולה אני אמרתי אלהים אתם ובני עליון כלכם אבן כאדם תמותון וכאחד השרים תפלו, פלם יבק פאחשה ולא קביחה חתי וצף בהא אלמלאיכה, כל דאלך גאריא עלי' אצלה הדא, פרכב הדה אלשנאעאת אלעצימה לקלה' בצרה בתפסיר אלכלאם באלגה, ואנא מביןן לטאה פי גמיע מא וצעה בעון אללה, אעלם אסעדך אללה אן אלשבהה אלדי דעת הדא אלמבדע אלהרי וגירה אל' הדא אלקול אן אלשטן מלך הוא קול אללה צעלי' לאלשטן הנה כל אשר לו בידיך רק אליו אל תשלח ידיך, ואיצא קולה רק את נפשו שמור, ואיצא קולה ויך את איוב בשחין רע, ותוהם אן הדא מנסוב אלי' אלשטן. ואנא אשרח מא בה תזול הדה אל תלאת אלשבה במא פיה מקנע. פאקול אן אלמראד ואלראי פי לגה' בני אסראיל יסמא יד. כמא קאל דוד לאלמסתגיתה אליה היד יואב אתך בכל זאת, מענאה ראי ומראד וכו' וענד זואל הדה אלשבהה אקול אן כטא הדא אלמעבר פי תלריג ויראו בני האלהים את בנות האדם. אפצע מן מא קאלה פי קצה איוב ודאלך אנא לם יקף בתה עלי' מא יסתעמלון בני אסראיל פי אלכאץ ואלעאם, ודאלך אנה מן סבילהם אן יקולון פעל אלכואץ כדא ופעל אלעואם כדא ויכון אולאיך אלכואץ איצא נאנא, ואנמא פרדהם לתכציציהם, פעלי, הדא קולה ויראו בני האלהים נאסא איצא אנמא פרדהם מן אלעואם לתכציצהם ושרפהם. והדה עבארה מתגהה לתנפי אלמנכר ען אלמקרבין. ויזידהא איצא ויקויהא קולה בעד דאלך וירא ה' כי רבה רעת האדם, ולם ידם אלמלאיכה מעהם. וקולה פי עקב דאלך אמחה את האדם אשר בראתי ולם יעאקב אלמלאיכה וכו' עכ"ל. העתק תמצא בסוף הספר הרי בהדיא שהגאון מרחיק מאד דעת האומר שבני האלהים הם מלאכים ודוחה אותה בשתי ידים וקורא לה דעת חיצונית. ודבריו מכוונים לדברי הרשב"י שבבראשית רבה, שלא כדעת הזהר שאמר שהם מלאכים, ושונה ומשלם בה, וזו הוכחה גדולה שאינו מרשב"י שהרי רשב"י היה מקלל את האומר שהם מלאכים. ובענין השטן שפי' הגאון שהוא אדם, אע"פ שהראב"ע השיגו יש מקום סתירה להשגתו ולהעמיד דברי הגאון הרס"ג, ואע"פ שבפ"ק דבבא בתרא משמע כהראב"ע. הנה ידוע מה שכ' הנגיד בעל מבוא התלמוד בפי' הגדה, ואין לך ללמוד ממנה אלא מה שיעלה על הדעת, ויש לך לדעת שכל מה שקיימו חז"ל הלכה בענין מצוה, שהוא מפי משה רבינו עליו השלום מפי הגבורה, אין לך להוסיף עליו ולא לגרוע ממנו אבל מה שפירשו בפסוקים כל אחד כפי שיזדמן לו ומה שראה בדעתו ולפי מה שיעלה על הדעת מן הפירושים האלה לומדים אותם, והשאר אין סומכים עליהם עכ"ל. ועיין רש"י שבת פ' במה מדליקים בשאלת מהו לכבויי בוצינא מקמי באישא בשבתא. והיוצא מזה דכל למסבר קראי בעלמא ואין נוגע ממנו לדינא בעשיית המצות ובעיקרי תורתינו הקדושה, אין בומכים אלא על כל מאי דמסתבר טפי. אבל כל מאי דפנקא מינא לדינא כנדון דידן כמצות יחוד השם יתברך אין לזוז מקבלת חז"ל שאין לשתף עם השם ית' שום כח נברא בין שיהיה גשמי או רוחני כמאמר רשב"י ע"ח כל המשתף שם שמים ודבר אחר נעקר מן העולם.

| id | word | punctuation after |
|---|---|---|
| 7:24.1 | צו | ) |
| 7:24.2 | גם |  |
| 7:24.3 | הרס"ג |  |
| 7:24.4 | בפירושו |  |
| 7:24.5 | לס | ' |
| 7:24.6 | איוב |  |
| 7:24.7 | מרחיק |  |
| 7:24.8 | מאד |  |
| 7:24.9 | סברת |  |
| 7:24.10 | האומר |  |
| 7:24.11 | שבני |  |
| 7:24.12 | האלהים |  |
| 7:24.13 | מלאכים |  |
| 7:24.14 | וביזה |  |
| 7:24.15 | וזלזל |  |
| 7:24.16 | את |  |
| 7:24.17 | האומר |  |
| 7:24.18 | שהם |  |
| 7:24.19 | מלאכים |  |
| 7:24.20 | כמו |  |
| 7:24.21 | שעשה |  |
| 7:24.22 | רשב"י |  |
| 7:24.23 | בבראשית |  |
| 7:24.24 | רבה |  |
| 7:24.25 | וז"ל |  |
| 7:24.26 | בלשון |  |
| 7:24.27 | ערבי | . |
| 7:24.28 | ומא |  |
| 7:24.29 | אבעד |  |
| 7:24.30 | אלקיאס |  |
| 7:24.31 | אן |  |
| 7:24.32 | יכון |  |
| 7:24.33 | הדא |  |
| 7:24.34 | אלשטן |  |
| 7:24.35 | מלכא | , |
| 7:24.36 | לאן |  |
| 7:24.37 | אהל |  |
| 7:24.38 | אלתוחיד |  |
| 7:24.39 | מגמעין |  |
| 7:24.40 | עלי | ' |
| 7:24.41 | אן |  |
| 7:24.42 | אלבארי |  |
| 7:24.43 | נל |  |
| 7:24.44 | ועז |  |
| 7:24.45 | כלק |  |
| 7:24.46 | מלאיכתה |  |
| 7:24.47 | מקרבין | . |
| 7:24.48 | מן |  |
| 7:24.49 | עלמה |  |
| 7:24.50 | אנחם |  |
| 7:24.51 | לא |  |
| 7:24.52 | יעצון |  |
| 7:24.53 | כמא |  |
| 7:24.54 | אנה |  |
| 7:24.55 | אצטפא |  |
| 7:24.56 | אנביאה |  |
| 7:24.57 | אלמרסלין |  |
| 7:24.58 | מן |  |
| 7:24.59 | עלמה |  |
| 7:24.60 | אנהם |  |
| 7:24.61 | לא |  |
| 7:24.62 | יעצונה | . |
| 7:24.63 | ומעאנדה |  |
| 7:24.64 | אלאוליא |  |
| 7:24.65 | אלצאלחין |  |
| 7:24.66 | מן |  |
| 7:24.67 | כבאיר |  |
| 7:24.68 | אלמעצי | ! |
| 7:24.69 | פמן |  |
| 7:24.70 | זעם |  |
| 7:24.71 | אם |  |
| 7:24.72 | הדא |  |
| 7:24.73 | אלמעאנד |  |
| 7:24.74 | לאיוב |  |
| 7:24.75 | כאן |  |
| 7:24.76 | מלכא |  |
| 7:24.77 | פקד |  |
| 7:24.78 | נעל |  |
| 7:24.79 | אלמלאיכה |  |
| 7:24.80 | יעצון | . |
| 7:24.81 | וכאלף |  |
| 7:24.82 | מא |  |
| 7:24.83 | נחן |  |
| 7:24.84 | מגמעין |  |
| 7:24.85 | עליה | , |
| 7:24.86 | וקד |  |
| 7:24.87 | תעלם |  |
| 7:24.88 | ארשדך |  |
| 7:24.89 | אללה |  |
| 7:24.90 | אן |  |
| 7:24.91 | לאלנפס |  |
| 7:24.92 | תלאת |  |
| 7:24.93 | קוי | , |
| 7:24.94 | אלפכר | , |
| 7:24.95 | ואלגצב |  |
| 7:24.96 | ואלשהות | , |
| 7:24.97 | פאמא |  |
| 7:24.98 | קווה |  |
| 7:24.99 | אלפכר |  |
| 7:24.100 | פחי |  |
| 7:24.101 | אלתר |  |
| 7:24.102 | ימייז |  |
| 7:24.103 | כתא |  |
| 7:24.104 | אל |  |
| 7:24.105 | אנסאן |  |
| 7:24.106 | אלצואב |  |
| 7:24.107 | מן |  |
| 7:24.108 | אלכטא |  |
| 7:24.109 | ואלחק |  |
| 7:24.110 | מן |  |
| 7:24.111 | אלבאטל | , |
| 7:24.112 | ואמא |  |
| 7:24.113 | קווה |  |
| 7:24.114 | אלגצב |  |
| 7:24.115 | פהי |  |
| 7:24.116 | אלתי |  |
| 7:24.117 | כהא |  |
| 7:24.118 | ילחק |  |
| 7:24.119 | אלחיואן |  |
| 7:24.120 | אלאקדאם |  |
| 7:24.121 | ואלגראות |  |
| 7:24.122 | ואלאנתקאם |  |
| 7:24.123 | ואסתיפא |  |
| 7:24.124 | אלטואיל | , |
| 7:24.125 | ואמא |  |
| 7:24.126 | קוה |  |
| 7:24.127 | אלשהוה |  |
| 7:24.128 | פבהא |  |
| 7:24.129 | ישתאק |  |
| 7:24.130 | אלחיואן |  |
| 7:24.131 | אלי |  |
| 7:24.132 | אלגדא |  |
| 7:24.133 | ואלי |  |
| 7:24.134 | אלתלקיח |  |
| 7:24.135 | וסאיר |  |
| 7:24.136 | אלדעאוי | . |
| 7:24.137 | ותעלם |  |
| 7:24.138 | מע |  |
| 7:24.139 | דאלך |  |
| 7:24.140 | אן |  |
| 7:24.141 | צפה |  |
| 7:24.142 | עדל |  |
| 7:24.143 | אלנאטק |  |
| 7:24.144 | אן |  |
| 7:24.145 | תכון |  |
| 7:24.146 | אלקוותן |  |
| 7:24.147 | אעני |  |
| 7:24.148 | קווה |  |
| 7:24.149 | אלשהוה |  |
| 7:24.150 | ואלגצב |  |
| 7:24.151 | מדענה |  |
| 7:24.152 | לאלפכר |  |
| 7:24.153 | פאן | ) |
| 7:24.154 | גלב |  |
| 7:24.155 | אליהא |  |
| 7:24.156 | אחדהמא |  |
| 7:24.157 | כאן |  |
| 7:24.158 | דאלך |  |
| 7:24.159 | גורא |  |
| 7:24.160 | וכטאא | , |
| 7:24.161 | פמן |  |
| 7:24.162 | גאוז |  |
| 7:24.163 | אן |  |
| 7:24.164 | תפעל |  |
| 7:24.165 | אלמלאיכה |  |
| 7:24.166 | אלחסד |  |
| 7:24.167 | אלדי |  |
| 7:24.168 | הו |  |
| 7:24.169 | אחד |  |
| 7:24.170 | אגואע |  |
| 7:24.171 | אלגצב |  |
| 7:24.172 | לזמה |  |
| 7:24.173 | תגויז |  |
| 7:24.174 | וקוע |  |
| 7:24.175 | אלשהואת |  |
| 7:24.176 | מן |  |
| 7:24.177 | אלמלאיכה |  |
| 7:24.178 | והו |  |
| 7:24.179 | אלאכל |  |
| 7:24.180 | ואלשרב |  |
| 7:24.181 | ואלנכאח |  |
| 7:24.182 | וסאיר |  |
| 7:24.183 | אלאמור |  |
| 7:24.184 | אלרדלה |  |
| 7:24.185 | אלוסלה |  |
| 7:24.186 | אלתי |  |
| 7:24.187 | קד |  |
| 7:24.188 | צאן |  |
| 7:24.189 | אללה |  |
| 7:24.190 | מלאכיתה |  |
| 7:24.191 | ענהא | , |
| 7:24.192 | סימא |  |
| 7:24.193 | וקד |  |
| 7:24.194 | סבק |  |
| 7:24.195 | פי |  |
| 7:24.196 | עלמה |  |
| 7:24.197 | אנהם |  |
| 7:24.198 | לא |  |
| 7:24.199 | יעצונה | . |
| 7:24.200 | פלא |  |
| 7:24.201 | יגוז |  |
| 7:24.202 | אן |  |
| 7:24.203 | יכון |  |
| 7:24.204 | באלפעל |  |
| 7:24.205 | כלאף |  |
| 7:24.206 | מא |  |
| 7:24.207 | סבק |  |
| 7:24.208 | פי |  |
| 7:24.209 | עלמה |  |
| 7:24.210 | וכו | ' |
| 7:24.211 | וקד |  |
| 7:24.212 | בלגני |  |
| 7:24.213 | ען |  |
| 7:24.214 | בעץ |  |
| 7:24.215 | אלכוארג |  |
| 7:24.216 | אן |  |
| 7:24.217 | כאלפה |  |
| 7:24.218 | בה |  |
| 7:24.219 | אלי | ' |
| 7:24.220 | אן |  |
| 7:24.221 | זעם |  |
| 7:24.222 | אן |  |
| 7:24.223 | הדא |  |
| 7:24.224 | אלשטן |  |
| 7:24.225 | מלך | . |
| 7:24.226 | ואן |  |
| 7:24.227 | אלחסד |  |
| 7:24.228 | ואלגירה |  |
| 7:24.229 | ואלבני |  |
| 7:24.230 | יגוז |  |
| 7:24.231 | וקועהם |  |
| 7:24.232 | באלפעל |  |
| 7:24.233 | מן |  |
| 7:24.234 | אלמלאיכה |  |
| 7:24.235 | כמה |  |
| 7:24.236 | תוהם |  |
| 7:24.237 | פי |  |
| 7:24.238 | קצח | ' |
| 7:24.239 | איוב |  |
| 7:24.240 | וחסד |  |
| 7:24.241 | אלשטן |  |
| 7:24.242 | לה | ' |
| 7:24.243 | תם |  |
| 7:24.244 | ראא |  |
| 7:24.245 | איצא |  |
| 7:24.246 | אן |  |
| 7:24.247 | באטלאקה |  |
| 7:24.248 | עליהם |  |
| 7:24.249 | דאלך |  |
| 7:24.250 | ילזמה |  |
| 7:24.251 | אטלאק |  |
| 7:24.252 | אל |  |
| 7:24.253 | שיאן |  |
| 7:24.254 | ואלנבאח | , |
| 7:24.255 | פנאוז |  |
| 7:24.256 | המא |  |
| 7:24.257 | איצא | . |
| 7:24.258 | ואדעא |  |
| 7:24.259 | אן |  |
| 7:24.260 | קצה |  |
| 7:24.261 | ויראו |  |
| 7:24.262 | בני |  |
| 7:24.263 | האלהים |  |
| 7:24.264 | את |  |
| 7:24.265 | בנות |  |
| 7:24.266 | האדם |  |
| 7:24.267 | כי |  |
| 7:24.268 | טבת |  |
| 7:24.269 | הנה | . |
| 7:24.270 | הם |  |
| 7:24.271 | אמלאיכה |  |
| 7:24.272 | אלנאכחון | , |
| 7:24.273 | אלחרם |  |
| 7:24.274 | באסקאטהם |  |
| 7:24.275 | מן |  |
| 7:24.276 | אלסמא |  |
| 7:24.277 | ותציירהם |  |
| 7:24.278 | כאלנאס | , |
| 7:24.279 | ונצב |  |
| 7:24.280 | להם |  |
| 7:24.281 | שבחה |  |
| 7:24.282 | הו |  |
| 7:24.283 | קולה |  |
| 7:24.284 | אני |  |
| 7:24.285 | אמרתי |  |
| 7:24.286 | אלהים |  |
| 7:24.287 | אתם |  |
| 7:24.288 | ובני |  |
| 7:24.289 | עליון |  |
| 7:24.290 | כלכם |  |
| 7:24.291 | אבן |  |
| 7:24.292 | כאדם |  |
| 7:24.293 | תמותון |  |
| 7:24.294 | וכאחד |  |
| 7:24.295 | השרים |  |
| 7:24.296 | תפלו | , |
| 7:24.297 | פלם |  |
| 7:24.298 | יבק |  |
| 7:24.299 | פאחשה |  |
| 7:24.300 | ולא |  |
| 7:24.301 | קביחה |  |
| 7:24.302 | חתי |  |
| 7:24.303 | וצף |  |
| 7:24.304 | בהא |  |
| 7:24.305 | אלמלאיכה | , |
| 7:24.306 | כל |  |
| 7:24.307 | דאלך |  |
| 7:24.308 | גאריא |  |
| 7:24.309 | עלי | ' |
| 7:24.310 | אצלה |  |
| 7:24.311 | הדא | , |
| 7:24.312 | פרכב |  |
| 7:24.313 | הדה |  |
| 7:24.314 | אלשנאעאת |  |
| 7:24.315 | אלעצימה |  |
| 7:24.316 | לקלה | ' |
| 7:24.317 | בצרה |  |
| 7:24.318 | בתפסיר |  |
| 7:24.319 | אלכלאם |  |
| 7:24.320 | באלגה | , |
| 7:24.321 | ואנא |  |
| 7:24.322 | מביןן |  |
| 7:24.323 | לטאה |  |
| 7:24.324 | פי |  |
| 7:24.325 | גמיע |  |
| 7:24.326 | מא |  |
| 7:24.327 | וצעה |  |
| 7:24.328 | בעון |  |
| 7:24.329 | אללה | , |
| 7:24.330 | אעלם |  |
| 7:24.331 | אסעדך |  |
| 7:24.332 | אללה |  |
| 7:24.333 | אן |  |
| 7:24.334 | אלשבהה |  |
| 7:24.335 | אלדי |  |
| 7:24.336 | דעת |  |
| 7:24.337 | הדא |  |
| 7:24.338 | אלמבדע |  |
| 7:24.339 | אלהרי |  |
| 7:24.340 | וגירה |  |
| 7:24.341 | אל | ' |
| 7:24.342 | הדא |  |
| 7:24.343 | אלקול |  |
| 7:24.344 | אן |  |
| 7:24.345 | אלשטן |  |
| 7:24.346 | מלך |  |
| 7:24.347 | הוא |  |
| 7:24.348 | קול |  |
| 7:24.349 | אללה |  |
| 7:24.350 | צעלי | ' |
| 7:24.351 | לאלשטן |  |
| 7:24.352 | הנה |  |
| 7:24.353 | כל |  |
| 7:24.354 | אשר |  |
| 7:24.355 | לו |  |
| 7:24.356 | בידיך |  |
| 7:24.357 | רק |  |
| 7:24.358 | אליו |  |
| 7:24.359 | אל |  |
| 7:24.360 | תשלח |  |
| 7:24.361 | ידיך | , |
| 7:24.362 | ואיצא |  |
| 7:24.363 | קולה |  |
| 7:24.364 | רק |  |
| 7:24.365 | את |  |
| 7:24.366 | נפשו |  |
| 7:24.367 | שמור | , |
| 7:24.368 | ואיצא |  |
| 7:24.369 | קולה |  |
| 7:24.370 | ויך |  |
| 7:24.371 | את |  |
| 7:24.372 | איוב |  |
| 7:24.373 | בשחין |  |
| 7:24.374 | רע | , |
| 7:24.375 | ותוהם |  |
| 7:24.376 | אן |  |
| 7:24.377 | הדא |  |
| 7:24.378 | מנסוב |  |
| 7:24.379 | אלי | ' |
| 7:24.380 | אלשטן | . |
| 7:24.381 | ואנא |  |
| 7:24.382 | אשרח |  |
| 7:24.383 | מא |  |
| 7:24.384 | בה |  |
| 7:24.385 | תזול |  |
| 7:24.386 | הדה |  |
| 7:24.387 | אל |  |
| 7:24.388 | תלאת |  |
| 7:24.389 | אלשבה |  |
| 7:24.390 | במא |  |
| 7:24.391 | פיה |  |
| 7:24.392 | מקנע | . |
| 7:24.393 | פאקול |  |
| 7:24.394 | אן |  |
| 7:24.395 | אלמראד |  |
| 7:24.396 | ואלראי |  |
| 7:24.397 | פי |  |
| 7:24.398 | לגה | ' |
| 7:24.399 | בני |  |
| 7:24.400 | אסראיל |  |
| 7:24.401 | יסמא |  |
| 7:24.402 | יד | . |
| 7:24.403 | כמא |  |
| 7:24.404 | קאל |  |
| 7:24.405 | דוד |  |
| 7:24.406 | לאלמסתגיתה |  |
| 7:24.407 | אליה |  |
| 7:24.408 | היד |  |
| 7:24.409 | יואב |  |
| 7:24.410 | אתך |  |
| 7:24.411 | בכל |  |
| 7:24.412 | זאת | , |
| 7:24.413 | מענאה |  |
| 7:24.414 | ראי |  |
| 7:24.415 | ומראד |  |
| 7:24.416 | וכו | ' |
| 7:24.417 | וענד |  |
| 7:24.418 | זואל |  |
| 7:24.419 | הדה |  |
| 7:24.420 | אלשבהה |  |
| 7:24.421 | אקול |  |
| 7:24.422 | אן |  |
| 7:24.423 | כטא |  |
| 7:24.424 | הדא |  |
| 7:24.425 | אלמעבר |  |
| 7:24.426 | פי |  |
| 7:24.427 | תלריג |  |
| 7:24.428 | ויראו |  |
| 7:24.429 | בני |  |
| 7:24.430 | האלהים |  |
| 7:24.431 | את |  |
| 7:24.432 | בנות |  |
| 7:24.433 | האדם | . |
| 7:24.434 | אפצע |  |
| 7:24.435 | מן |  |
| 7:24.436 | מא |  |
| 7:24.437 | קאלה |  |
| 7:24.438 | פי |  |
| 7:24.439 | קצה |  |
| 7:24.440 | איוב |  |
| 7:24.441 | ודאלך |  |
| 7:24.442 | אנא |  |
| 7:24.443 | לם |  |
| 7:24.444 | יקף |  |
| 7:24.445 | בתה |  |
| 7:24.446 | עלי | ' |
| 7:24.447 | מא |  |
| 7:24.448 | יסתעמלון |  |
| 7:24.449 | בני |  |
| 7:24.450 | אסראיל |  |
| 7:24.451 | פי |  |
| 7:24.452 | אלכאץ |  |
| 7:24.453 | ואלעאם | , |
| 7:24.454 | ודאלך |  |
| 7:24.455 | אנה |  |
| 7:24.456 | מן |  |
| 7:24.457 | סבילהם |  |
| 7:24.458 | אן |  |
| 7:24.459 | יקולון |  |
| 7:24.460 | פעל |  |
| 7:24.461 | אלכואץ |  |
| 7:24.462 | כדא |  |
| 7:24.463 | ופעל |  |
| 7:24.464 | אלעואם |  |
| 7:24.465 | כדא |  |
| 7:24.466 | ויכון |  |
| 7:24.467 | אולאיך |  |
| 7:24.468 | אלכואץ |  |
| 7:24.469 | איצא |  |
| 7:24.470 | נאנא | , |
| 7:24.471 | ואנמא |  |
| 7:24.472 | פרדהם |  |
| 7:24.473 | לתכציציהם | , |
| 7:24.474 | פעלי | , |
| 7:24.475 | הדא |  |
| 7:24.476 | קולה |  |
| 7:24.477 | ויראו |  |
| 7:24.478 | בני |  |
| 7:24.479 | האלהים |  |
| 7:24.480 | נאסא |  |
| 7:24.481 | איצא |  |
| 7:24.482 | אנמא |  |
| 7:24.483 | פרדהם |  |
| 7:24.484 | מן |  |
| 7:24.485 | אלעואם |  |
| 7:24.486 | לתכציצהם |  |
| 7:24.487 | ושרפהם | . |
| 7:24.488 | והדה |  |
| 7:24.489 | עבארה |  |
| 7:24.490 | מתגהה |  |
| 7:24.491 | לתנפי |  |
| 7:24.492 | אלמנכר |  |
| 7:24.493 | ען |  |
| 7:24.494 | אלמקרבין | . |
| 7:24.495 | ויזידהא |  |
| 7:24.496 | איצא |  |
| 7:24.497 | ויקויהא |  |
| 7:24.498 | קולה |  |
| 7:24.499 | בעד |  |
| 7:24.500 | דאלך |  |
| 7:24.501 | וירא |  |
| 7:24.502 | ה | ' |
| 7:24.503 | כי |  |
| 7:24.504 | רבה |  |
| 7:24.505 | רעת |  |
| 7:24.506 | האדם | , |
| 7:24.507 | ולם |  |
| 7:24.508 | ידם |  |
| 7:24.509 | אלמלאיכה |  |
| 7:24.510 | מעהם | . |
| 7:24.511 | וקולה |  |
| 7:24.512 | פי |  |
| 7:24.513 | עקב |  |
| 7:24.514 | דאלך |  |
| 7:24.515 | אמחה |  |
| 7:24.516 | את |  |
| 7:24.517 | האדם |  |
| 7:24.518 | אשר |  |
| 7:24.519 | בראתי |  |
| 7:24.520 | ולם |  |
| 7:24.521 | יעאקב |  |
| 7:24.522 | אלמלאיכה |  |
| 7:24.523 | וכו | ' |
| 7:24.524 | עכ"ל | . |
| 7:24.525 | העתק |  |
| 7:24.526 | תמצא |  |
| 7:24.527 | בסוף |  |
| 7:24.528 | הספר |  |
| 7:24.529 | הרי |  |
| 7:24.530 | בהדיא |  |
| 7:24.531 | שהגאון |  |
| 7:24.532 | מרחיק |  |
| 7:24.533 | מאד |  |
| 7:24.534 | דעת |  |
| 7:24.535 | האומר |  |
| 7:24.536 | שבני |  |
| 7:24.537 | האלהים |  |
| 7:24.538 | הם |  |
| 7:24.539 | מלאכים |  |
| 7:24.540 | ודוחה |  |
| 7:24.541 | אותה |  |
| 7:24.542 | בשתי |  |
| 7:24.543 | ידים |  |
| 7:24.544 | וקורא |  |
| 7:24.545 | לה |  |
| 7:24.546 | דעת |  |
| 7:24.547 | חיצונית | . |
| 7:24.548 | ודבריו |  |
| 7:24.549 | מכוונים |  |
| 7:24.550 | לדברי |  |
| 7:24.551 | הרשב"י |  |
| 7:24.552 | שבבראשית |  |
| 7:24.553 | רבה | , |
| 7:24.554 | שלא |  |
| 7:24.555 | כדעת |  |
| 7:24.556 | הזהר |  |
| 7:24.557 | שאמר |  |
| 7:24.558 | שהם |  |
| 7:24.559 | מלאכים | , |
| 7:24.560 | ושונה |  |
| 7:24.561 | ומשלם |  |
| 7:24.562 | בה | , |
| 7:24.563 | וזו |  |
| 7:24.564 | הוכחה |  |
| 7:24.565 | גדולה |  |
| 7:24.566 | שאינו |  |
| 7:24.567 | מרשב"י |  |
| 7:24.568 | שהרי |  |
| 7:24.569 | רשב"י |  |
| 7:24.570 | היה |  |
| 7:24.571 | מקלל |  |
| 7:24.572 | את |  |
| 7:24.573 | האומר |  |
| 7:24.574 | שהם |  |
| 7:24.575 | מלאכים | . |
| 7:24.576 | ובענין |  |
| 7:24.577 | השטן |  |
| 7:24.578 | שפי | ' |
| 7:24.579 | הגאון |  |
| 7:24.580 | שהוא |  |
| 7:24.581 | אדם | , |
| 7:24.582 | אע"פ |  |
| 7:24.583 | שהראב"ע |  |
| 7:24.584 | השיגו |  |
| 7:24.585 | יש |  |
| 7:24.586 | מקום |  |
| 7:24.587 | סתירה |  |
| 7:24.588 | להשגתו |  |
| 7:24.589 | ולהעמיד |  |
| 7:24.590 | דברי |  |
| 7:24.591 | הגאון |  |
| 7:24.592 | הרס"ג | , |
| 7:24.593 | ואע"פ |  |
| 7:24.594 | שבפ"ק |  |
| 7:24.595 | דבבא |  |
| 7:24.596 | בתרא |  |
| 7:24.597 | משמע |  |
| 7:24.598 | כהראב"ע | . |
| 7:24.599 | הנה |  |
| 7:24.600 | ידוע |  |
| 7:24.601 | מה |  |
| 7:24.602 | שכ | ' |
| 7:24.603 | הנגיד |  |
| 7:24.604 | בעל |  |
| 7:24.605 | מבוא |  |
| 7:24.606 | התלמוד |  |
| 7:24.607 | בפי | ' |
| 7:24.608 | הגדה | , |
| 7:24.609 | ואין |  |
| 7:24.610 | לך |  |
| 7:24.611 | ללמוד |  |
| 7:24.612 | ממנה |  |
| 7:24.613 | אלא |  |
| 7:24.614 | מה |  |
| 7:24.615 | שיעלה |  |
| 7:24.616 | על |  |
| 7:24.617 | הדעת | , |
| 7:24.618 | ויש |  |
| 7:24.619 | לך |  |
| 7:24.620 | לדעת |  |
| 7:24.621 | שכל |  |
| 7:24.622 | מה |  |
| 7:24.623 | שקיימו |  |
| 7:24.624 | חז"ל |  |
| 7:24.625 | הלכה |  |
| 7:24.626 | בענין |  |
| 7:24.627 | מצוה | , |
| 7:24.628 | שהוא |  |
| 7:24.629 | מפי |  |
| 7:24.630 | משה |  |
| 7:24.631 | רבינו |  |
| 7:24.632 | עליו |  |
| 7:24.633 | השלום |  |
| 7:24.634 | מפי |  |
| 7:24.635 | הגבורה | , |
| 7:24.636 | אין |  |
| 7:24.637 | לך |  |
| 7:24.638 | להוסיף |  |
| 7:24.639 | עליו |  |
| 7:24.640 | ולא |  |
| 7:24.641 | לגרוע |  |
| 7:24.642 | ממנו |  |
| 7:24.643 | אבל |  |
| 7:24.644 | מה |  |
| 7:24.645 | שפירשו |  |
| 7:24.646 | בפסוקים |  |
| 7:24.647 | כל |  |
| 7:24.648 | אחד |  |
| 7:24.649 | כפי |  |
| 7:24.650 | שיזדמן |  |
| 7:24.651 | לו |  |
| 7:24.652 | ומה |  |
| 7:24.653 | שראה |  |
| 7:24.654 | בדעתו |  |
| 7:24.655 | ולפי |  |
| 7:24.656 | מה |  |
| 7:24.657 | שיעלה |  |
| 7:24.658 | על |  |
| 7:24.659 | הדעת |  |
| 7:24.660 | מן |  |
| 7:24.661 | הפירושים |  |
| 7:24.662 | האלה |  |
| 7:24.663 | לומדים |  |
| 7:24.664 | אותם | , |
| 7:24.665 | והשאר |  |
| 7:24.666 | אין |  |
| 7:24.667 | סומכים |  |
| 7:24.668 | עליהם |  |
| 7:24.669 | עכ"ל | . |
| 7:24.670 | ועיין |  |
| 7:24.671 | רש"י |  |
| 7:24.672 | שבת |  |
| 7:24.673 | פ | ' |
| 7:24.674 | במה |  |
| 7:24.675 | מדליקים |  |
| 7:24.676 | בשאלת |  |
| 7:24.677 | מהו |  |
| 7:24.678 | לכבויי |  |
| 7:24.679 | בוצינא |  |
| 7:24.680 | מקמי |  |
| 7:24.681 | באישא |  |
| 7:24.682 | בשבתא | . |
| 7:24.683 | והיוצא |  |
| 7:24.684 | מזה |  |
| 7:24.685 | דכל |  |
| 7:24.686 | למסבר |  |
| 7:24.687 | קראי |  |
| 7:24.688 | בעלמא |  |
| 7:24.689 | ואין |  |
| 7:24.690 | נוגע |  |
| 7:24.691 | ממנו |  |
| 7:24.692 | לדינא |  |
| 7:24.693 | בעשיית |  |
| 7:24.694 | המצות |  |
| 7:24.695 | ובעיקרי |  |
| 7:24.696 | תורתינו |  |
| 7:24.697 | הקדושה | , |
| 7:24.698 | אין |  |
| 7:24.699 | בומכים |  |
| 7:24.700 | אלא |  |
| 7:24.701 | על |  |
| 7:24.702 | כל |  |
| 7:24.703 | מאי |  |
| 7:24.704 | דמסתבר |  |
| 7:24.705 | טפי | . |
| 7:24.706 | אבל |  |
| 7:24.707 | כל |  |
| 7:24.708 | מאי |  |
| 7:24.709 | דפנקא |  |
| 7:24.710 | מינא |  |
| 7:24.711 | לדינא |  |
| 7:24.712 | כנדון |  |
| 7:24.713 | דידן |  |
| 7:24.714 | כמצות |  |
| 7:24.715 | יחוד |  |
| 7:24.716 | השם |  |
| 7:24.717 | יתברך |  |
| 7:24.718 | אין |  |
| 7:24.719 | לזוז |  |
| 7:24.720 | מקבלת |  |
| 7:24.721 | חז"ל |  |
| 7:24.722 | שאין |  |
| 7:24.723 | לשתף |  |
| 7:24.724 | עם |  |
| 7:24.725 | השם |  |
| 7:24.726 | ית | ' |
| 7:24.727 | שום |  |
| 7:24.728 | כח |  |
| 7:24.729 | נברא |  |
| 7:24.730 | בין |  |
| 7:24.731 | שיהיה |  |
| 7:24.732 | גשמי |  |
| 7:24.733 | או |  |
| 7:24.734 | רוחני |  |
| 7:24.735 | כמאמר |  |
| 7:24.736 | רשב"י |  |
| 7:24.737 | ע"ח |  |
| 7:24.738 | כל |  |
| 7:24.739 | המשתף |  |
| 7:24.740 | שם |  |
| 7:24.741 | שמים |  |
| 7:24.742 | ודבר |  |
| 7:24.743 | אחר |  |
| 7:24.744 | נעקר |  |
| 7:24.745 | מן |  |
| 7:24.746 | העולם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

