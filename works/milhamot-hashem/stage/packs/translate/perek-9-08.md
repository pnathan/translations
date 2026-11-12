# Translation pack: milhamot-hashem perek-9-08 (Milchamot ha-Shem 9:10–9:10)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-9-08.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-9-08`.

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
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 9:10

Hebrew (Wikisource): קכז) לבי לחוקקי ישראל המתנבים בעם להבין ולהשכיל בתורת ה' בחכמה ותבונה. ובדעת ישרה על פי קבלת רז"ל בעלי המשנה והתלמוד והפוסקים הדולים ומשקים לכל צמא ממקור מים חיים, משה קבל תורה מסיני ומסרה ליהושע וכו' והם מלאו אחר ה' ותורתו הקדושה שלא לפנות אל האלילים, איך בכל מצותיה של תורה האריכו הרחיבו לבאר חילוקי דיניהם בכמה פרטים, וחומרות יתירות ממה שנזכר בתלמוד! ובדיני יחוד השי"ת שהוא העיקר הגדול, העלימו עיניהם ועזבוהו ולא דברו מאומה אפי' בפרט אחד מכל דיני יחודו ית'! ולמה זה ועל מה זה הניחוהו כעיר פרוצה אין חומה, והוא נוהג בכל זמן ובכל מקום בין בארץ בין בחוצא לארץ, בין בפני הבית בין שלא בפני הבית, כאלו נשתכחה מישראל תורת יחוד השי"ת היוצאת מדברי התלמוד בבלי וירושלמי ומדרשי רז"ל, אחרי מות ר' בחיי בעל חובת הלבבות ורבי יהודה הלוי והרס"ג והרמב"ם והסמ"ג ובעל העיקרים וזולתם! עד שבעבור זה נפלו קצת אחרונים בפח הנסיון אשר נסה ה' אלהינו אותנו על ידי המלך העכו"ם בשלחו ספר הזהר לחכמי טוליטולא, והמה שמחו בו שמחה גדולה, כי חפצו להוציא ישן תורת המשנה והתלמוד, מפני חדש תורת הזהר המבזה את המשנה והתלמוד כמבואר לעיל, ולא נתנו אל לבם למנוע מלעבוד אלהים חדשים, מקרבו באו, אשר הזהירנו ה' אלהינו לבלתי עבוד זולתו, אפי' נקראהו בשמו כמעשה דורו של אנוש. וחכמי טוליטולא לא עשו כן אדרבא שלחו מתנות רבות למלך העכו"ם בעבורו והגדילו והחשיבו דברי הפילסוף מחבר הזהר המתנבא בשקר בשם רשב"י שנגלה לו עתיק, יותר ממ"ח נביאים ושבע נביאות שנתנבאו להם לישראל, שמ"ח נביאים ושבע נביאות לא פיחתו ולא הותירו על מה שכתוב בתורה אפי' אות אחת חוץ ממקרא מגלה, ומחבר הזהר הוסיף אמונת אלהים אחרים חדשים אשר לא שערום אבותיהם. ואמר שמי שלא יעסוק בתורתו החדשה לית ליה רשו למיעל קמי מריה. ובספר אצהאר אלחק לאחד מחכמי הערביים בטענתו נגד אמונת השלוש שהיא אמונת הנצרים כתב בדף רפ"א ורפ"ב וז"ל. ועלמא אליהוד מן ודן מוסי' אלי' זמאננא הדא לא יעתרפון אתתלית (תלת קשרי דמהימנותא) וצאהר אן דאת אללה וצפאתה אלכמאליה קדימא גיה מתגיירה אזלא' ואבדא'. פלו כאן אלתתלית חקא', לכאן אלואגב עלי' מוסי' עליה אסלאם ואנביא בני אסראיל אן יביינוה חק אלביאן, פאלעגב כל אלעגב אן תכון אלשריעה אלמוסויה אלתי כאנת ואנבת אלאטאעה בגמיע אלאנביא אלי' עהד עיסי' כאלה עם ביאן הדה אלעקידה אלתי הי מדאר אלנגאה (היא מרכז ההצלה והוא מאמר הזהר דלית ליה רשו למיעל קמי מריה) עלי' זעם אהל אלתתלית ולא ימכן נגאת אחדכדונהא נבייא כאן או גיר נבי, ולא יביין מוסי' ע"ס. ולא נבי מן אלאנביא אלאסראלייה הדה אלעקידה בביאן ואצח בחית תפהם מנה הדה אלעקידה צראחה, ולא יבקש פיהא שך ויביין מוסי' אלאחכאם אלתי הי ענד מקרס אהל אלתתלית צעיפה נאקצה גדא באלתשריח אלאתם. ויכררהא מרה בעד אולי וכרה בעד אכרא, ויאכד ען מחאפצהתא תאכידא בלינא, ויוגב אלקתל למן לם יעתקדהא עכ"ל. (העתק ועיין בסוף הספר) וטענתו זאת בעצמה אמיתית צודקת על הקבלה החדשה, ומדבריו מוכח שלא פשטה עדיין אמונת הקבלה החדשה בישראל.

| id | word | punctuation after |
|---|---|---|
| 9:10.1 | קכז | ) |
| 9:10.2 | לבי |  |
| 9:10.3 | לחוקקי |  |
| 9:10.4 | ישראל |  |
| 9:10.5 | המתנבים |  |
| 9:10.6 | בעם |  |
| 9:10.7 | להבין |  |
| 9:10.8 | ולהשכיל |  |
| 9:10.9 | בתורת |  |
| 9:10.10 | ה | ' |
| 9:10.11 | בחכמה |  |
| 9:10.12 | ותבונה | . |
| 9:10.13 | ובדעת |  |
| 9:10.14 | ישרה |  |
| 9:10.15 | על |  |
| 9:10.16 | פי |  |
| 9:10.17 | קבלת |  |
| 9:10.18 | רז"ל |  |
| 9:10.19 | בעלי |  |
| 9:10.20 | המשנה |  |
| 9:10.21 | והתלמוד |  |
| 9:10.22 | והפוסקים |  |
| 9:10.23 | הדולים |  |
| 9:10.24 | ומשקים |  |
| 9:10.25 | לכל |  |
| 9:10.26 | צמא |  |
| 9:10.27 | ממקור |  |
| 9:10.28 | מים |  |
| 9:10.29 | חיים | , |
| 9:10.30 | משה |  |
| 9:10.31 | קבל |  |
| 9:10.32 | תורה |  |
| 9:10.33 | מסיני |  |
| 9:10.34 | ומסרה |  |
| 9:10.35 | ליהושע |  |
| 9:10.36 | וכו | ' |
| 9:10.37 | והם |  |
| 9:10.38 | מלאו |  |
| 9:10.39 | אחר |  |
| 9:10.40 | ה | ' |
| 9:10.41 | ותורתו |  |
| 9:10.42 | הקדושה |  |
| 9:10.43 | שלא |  |
| 9:10.44 | לפנות |  |
| 9:10.45 | אל |  |
| 9:10.46 | האלילים | , |
| 9:10.47 | איך |  |
| 9:10.48 | בכל |  |
| 9:10.49 | מצותיה |  |
| 9:10.50 | של |  |
| 9:10.51 | תורה |  |
| 9:10.52 | האריכו |  |
| 9:10.53 | הרחיבו |  |
| 9:10.54 | לבאר |  |
| 9:10.55 | חילוקי |  |
| 9:10.56 | דיניהם |  |
| 9:10.57 | בכמה |  |
| 9:10.58 | פרטים | , |
| 9:10.59 | וחומרות |  |
| 9:10.60 | יתירות |  |
| 9:10.61 | ממה |  |
| 9:10.62 | שנזכר |  |
| 9:10.63 | בתלמוד | ! |
| 9:10.64 | ובדיני |  |
| 9:10.65 | יחוד |  |
| 9:10.66 | השי"ת |  |
| 9:10.67 | שהוא |  |
| 9:10.68 | העיקר |  |
| 9:10.69 | הגדול | , |
| 9:10.70 | העלימו |  |
| 9:10.71 | עיניהם |  |
| 9:10.72 | ועזבוהו |  |
| 9:10.73 | ולא |  |
| 9:10.74 | דברו |  |
| 9:10.75 | מאומה |  |
| 9:10.76 | אפי | ' |
| 9:10.77 | בפרט |  |
| 9:10.78 | אחד |  |
| 9:10.79 | מכל |  |
| 9:10.80 | דיני |  |
| 9:10.81 | יחודו |  |
| 9:10.82 | ית | '! |
| 9:10.83 | ולמה |  |
| 9:10.84 | זה |  |
| 9:10.85 | ועל |  |
| 9:10.86 | מה |  |
| 9:10.87 | זה |  |
| 9:10.88 | הניחוהו |  |
| 9:10.89 | כעיר |  |
| 9:10.90 | פרוצה |  |
| 9:10.91 | אין |  |
| 9:10.92 | חומה | , |
| 9:10.93 | והוא |  |
| 9:10.94 | נוהג |  |
| 9:10.95 | בכל |  |
| 9:10.96 | זמן |  |
| 9:10.97 | ובכל |  |
| 9:10.98 | מקום |  |
| 9:10.99 | בין |  |
| 9:10.100 | בארץ |  |
| 9:10.101 | בין |  |
| 9:10.102 | בחוצא |  |
| 9:10.103 | לארץ | , |
| 9:10.104 | בין |  |
| 9:10.105 | בפני |  |
| 9:10.106 | הבית |  |
| 9:10.107 | בין |  |
| 9:10.108 | שלא |  |
| 9:10.109 | בפני |  |
| 9:10.110 | הבית | , |
| 9:10.111 | כאלו |  |
| 9:10.112 | נשתכחה |  |
| 9:10.113 | מישראל |  |
| 9:10.114 | תורת |  |
| 9:10.115 | יחוד |  |
| 9:10.116 | השי"ת |  |
| 9:10.117 | היוצאת |  |
| 9:10.118 | מדברי |  |
| 9:10.119 | התלמוד |  |
| 9:10.120 | בבלי |  |
| 9:10.121 | וירושלמי |  |
| 9:10.122 | ומדרשי |  |
| 9:10.123 | רז"ל | , |
| 9:10.124 | אחרי |  |
| 9:10.125 | מות |  |
| 9:10.126 | ר | ' |
| 9:10.127 | בחיי |  |
| 9:10.128 | בעל |  |
| 9:10.129 | חובת |  |
| 9:10.130 | הלבבות |  |
| 9:10.131 | ורבי |  |
| 9:10.132 | יהודה |  |
| 9:10.133 | הלוי |  |
| 9:10.134 | והרס"ג |  |
| 9:10.135 | והרמב"ם |  |
| 9:10.136 | והסמ"ג |  |
| 9:10.137 | ובעל |  |
| 9:10.138 | העיקרים |  |
| 9:10.139 | וזולתם | ! |
| 9:10.140 | עד |  |
| 9:10.141 | שבעבור |  |
| 9:10.142 | זה |  |
| 9:10.143 | נפלו |  |
| 9:10.144 | קצת |  |
| 9:10.145 | אחרונים |  |
| 9:10.146 | בפח |  |
| 9:10.147 | הנסיון |  |
| 9:10.148 | אשר |  |
| 9:10.149 | נסה |  |
| 9:10.150 | ה | ' |
| 9:10.151 | אלהינו |  |
| 9:10.152 | אותנו |  |
| 9:10.153 | על |  |
| 9:10.154 | ידי |  |
| 9:10.155 | המלך |  |
| 9:10.156 | העכו"ם |  |
| 9:10.157 | בשלחו |  |
| 9:10.158 | ספר |  |
| 9:10.159 | הזהר |  |
| 9:10.160 | לחכמי |  |
| 9:10.161 | טוליטולא | , |
| 9:10.162 | והמה |  |
| 9:10.163 | שמחו |  |
| 9:10.164 | בו |  |
| 9:10.165 | שמחה |  |
| 9:10.166 | גדולה | , |
| 9:10.167 | כי |  |
| 9:10.168 | חפצו |  |
| 9:10.169 | להוציא |  |
| 9:10.170 | ישן |  |
| 9:10.171 | תורת |  |
| 9:10.172 | המשנה |  |
| 9:10.173 | והתלמוד | , |
| 9:10.174 | מפני |  |
| 9:10.175 | חדש |  |
| 9:10.176 | תורת |  |
| 9:10.177 | הזהר |  |
| 9:10.178 | המבזה |  |
| 9:10.179 | את |  |
| 9:10.180 | המשנה |  |
| 9:10.181 | והתלמוד |  |
| 9:10.182 | כמבואר |  |
| 9:10.183 | לעיל | , |
| 9:10.184 | ולא |  |
| 9:10.185 | נתנו |  |
| 9:10.186 | אל |  |
| 9:10.187 | לבם |  |
| 9:10.188 | למנוע |  |
| 9:10.189 | מלעבוד |  |
| 9:10.190 | אלהים |  |
| 9:10.191 | חדשים | , |
| 9:10.192 | מקרבו |  |
| 9:10.193 | באו | , |
| 9:10.194 | אשר |  |
| 9:10.195 | הזהירנו |  |
| 9:10.196 | ה | ' |
| 9:10.197 | אלהינו |  |
| 9:10.198 | לבלתי |  |
| 9:10.199 | עבוד |  |
| 9:10.200 | זולתו | , |
| 9:10.201 | אפי | ' |
| 9:10.202 | נקראהו |  |
| 9:10.203 | בשמו |  |
| 9:10.204 | כמעשה |  |
| 9:10.205 | דורו |  |
| 9:10.206 | של |  |
| 9:10.207 | אנוש | . |
| 9:10.208 | וחכמי |  |
| 9:10.209 | טוליטולא |  |
| 9:10.210 | לא |  |
| 9:10.211 | עשו |  |
| 9:10.212 | כן |  |
| 9:10.213 | אדרבא |  |
| 9:10.214 | שלחו |  |
| 9:10.215 | מתנות |  |
| 9:10.216 | רבות |  |
| 9:10.217 | למלך |  |
| 9:10.218 | העכו"ם |  |
| 9:10.219 | בעבורו |  |
| 9:10.220 | והגדילו |  |
| 9:10.221 | והחשיבו |  |
| 9:10.222 | דברי |  |
| 9:10.223 | הפילסוף |  |
| 9:10.224 | מחבר |  |
| 9:10.225 | הזהר |  |
| 9:10.226 | המתנבא |  |
| 9:10.227 | בשקר |  |
| 9:10.228 | בשם |  |
| 9:10.229 | רשב"י |  |
| 9:10.230 | שנגלה |  |
| 9:10.231 | לו |  |
| 9:10.232 | עתיק | , |
| 9:10.233 | יותר |  |
| 9:10.234 | ממ"ח |  |
| 9:10.235 | נביאים |  |
| 9:10.236 | ושבע |  |
| 9:10.237 | נביאות |  |
| 9:10.238 | שנתנבאו |  |
| 9:10.239 | להם |  |
| 9:10.240 | לישראל | , |
| 9:10.241 | שמ"ח |  |
| 9:10.242 | נביאים |  |
| 9:10.243 | ושבע |  |
| 9:10.244 | נביאות |  |
| 9:10.245 | לא |  |
| 9:10.246 | פיחתו |  |
| 9:10.247 | ולא |  |
| 9:10.248 | הותירו |  |
| 9:10.249 | על |  |
| 9:10.250 | מה |  |
| 9:10.251 | שכתוב |  |
| 9:10.252 | בתורה |  |
| 9:10.253 | אפי | ' |
| 9:10.254 | אות |  |
| 9:10.255 | אחת |  |
| 9:10.256 | חוץ |  |
| 9:10.257 | ממקרא |  |
| 9:10.258 | מגלה | , |
| 9:10.259 | ומחבר |  |
| 9:10.260 | הזהר |  |
| 9:10.261 | הוסיף |  |
| 9:10.262 | אמונת |  |
| 9:10.263 | אלהים |  |
| 9:10.264 | אחרים |  |
| 9:10.265 | חדשים |  |
| 9:10.266 | אשר |  |
| 9:10.267 | לא |  |
| 9:10.268 | שערום |  |
| 9:10.269 | אבותיהם | . |
| 9:10.270 | ואמר |  |
| 9:10.271 | שמי |  |
| 9:10.272 | שלא |  |
| 9:10.273 | יעסוק |  |
| 9:10.274 | בתורתו |  |
| 9:10.275 | החדשה |  |
| 9:10.276 | לית |  |
| 9:10.277 | ליה |  |
| 9:10.278 | רשו |  |
| 9:10.279 | למיעל |  |
| 9:10.280 | קמי |  |
| 9:10.281 | מריה | . |
| 9:10.282 | ובספר |  |
| 9:10.283 | אצהאר |  |
| 9:10.284 | אלחק |  |
| 9:10.285 | לאחד |  |
| 9:10.286 | מחכמי |  |
| 9:10.287 | הערביים |  |
| 9:10.288 | בטענתו |  |
| 9:10.289 | נגד |  |
| 9:10.290 | אמונת |  |
| 9:10.291 | השלוש |  |
| 9:10.292 | שהיא |  |
| 9:10.293 | אמונת |  |
| 9:10.294 | הנצרים |  |
| 9:10.295 | כתב |  |
| 9:10.296 | בדף |  |
| 9:10.297 | רפ"א |  |
| 9:10.298 | ורפ"ב |  |
| 9:10.299 | וז"ל | . |
| 9:10.300 | ועלמא |  |
| 9:10.301 | אליהוד |  |
| 9:10.302 | מן |  |
| 9:10.303 | ודן |  |
| 9:10.304 | מוסי | ' |
| 9:10.305 | אלי | ' |
| 9:10.306 | זמאננא |  |
| 9:10.307 | הדא |  |
| 9:10.308 | לא |  |
| 9:10.309 | יעתרפון |  |
| 9:10.310 | אתתלית |  |
| 9:10.311 | תלת |  |
| 9:10.312 | קשרי |  |
| 9:10.313 | דמהימנותא | ) |
| 9:10.314 | וצאהר |  |
| 9:10.315 | אן |  |
| 9:10.316 | דאת |  |
| 9:10.317 | אללה |  |
| 9:10.318 | וצפאתה |  |
| 9:10.319 | אלכמאליה |  |
| 9:10.320 | קדימא |  |
| 9:10.321 | גיה |  |
| 9:10.322 | מתגיירה |  |
| 9:10.323 | אזלא | ' |
| 9:10.324 | ואבדא | '. |
| 9:10.325 | פלו |  |
| 9:10.326 | כאן |  |
| 9:10.327 | אלתתלית |  |
| 9:10.328 | חקא | ', |
| 9:10.329 | לכאן |  |
| 9:10.330 | אלואגב |  |
| 9:10.331 | עלי | ' |
| 9:10.332 | מוסי | ' |
| 9:10.333 | עליה |  |
| 9:10.334 | אסלאם |  |
| 9:10.335 | ואנביא |  |
| 9:10.336 | בני |  |
| 9:10.337 | אסראיל |  |
| 9:10.338 | אן |  |
| 9:10.339 | יביינוה |  |
| 9:10.340 | חק |  |
| 9:10.341 | אלביאן | , |
| 9:10.342 | פאלעגב |  |
| 9:10.343 | כל |  |
| 9:10.344 | אלעגב |  |
| 9:10.345 | אן |  |
| 9:10.346 | תכון |  |
| 9:10.347 | אלשריעה |  |
| 9:10.348 | אלמוסויה |  |
| 9:10.349 | אלתי |  |
| 9:10.350 | כאנת |  |
| 9:10.351 | ואנבת |  |
| 9:10.352 | אלאטאעה |  |
| 9:10.353 | בגמיע |  |
| 9:10.354 | אלאנביא |  |
| 9:10.355 | אלי | ' |
| 9:10.356 | עהד |  |
| 9:10.357 | עיסי | ' |
| 9:10.358 | כאלה |  |
| 9:10.359 | עם |  |
| 9:10.360 | ביאן |  |
| 9:10.361 | הדה |  |
| 9:10.362 | אלעקידה |  |
| 9:10.363 | אלתי |  |
| 9:10.364 | הי |  |
| 9:10.365 | מדאר |  |
| 9:10.366 | אלנגאה |  |
| 9:10.367 | היא |  |
| 9:10.368 | מרכז |  |
| 9:10.369 | ההצלה |  |
| 9:10.370 | והוא |  |
| 9:10.371 | מאמר |  |
| 9:10.372 | הזהר |  |
| 9:10.373 | דלית |  |
| 9:10.374 | ליה |  |
| 9:10.375 | רשו |  |
| 9:10.376 | למיעל |  |
| 9:10.377 | קמי |  |
| 9:10.378 | מריה | ) |
| 9:10.379 | עלי | ' |
| 9:10.380 | זעם |  |
| 9:10.381 | אהל |  |
| 9:10.382 | אלתתלית |  |
| 9:10.383 | ולא |  |
| 9:10.384 | ימכן |  |
| 9:10.385 | נגאת |  |
| 9:10.386 | אחדכדונהא |  |
| 9:10.387 | נבייא |  |
| 9:10.388 | כאן |  |
| 9:10.389 | או |  |
| 9:10.390 | גיר |  |
| 9:10.391 | נבי | , |
| 9:10.392 | ולא |  |
| 9:10.393 | יביין |  |
| 9:10.394 | מוסי | ' |
| 9:10.395 | ע"ס | . |
| 9:10.396 | ולא |  |
| 9:10.397 | נבי |  |
| 9:10.398 | מן |  |
| 9:10.399 | אלאנביא |  |
| 9:10.400 | אלאסראלייה |  |
| 9:10.401 | הדה |  |
| 9:10.402 | אלעקידה |  |
| 9:10.403 | בביאן |  |
| 9:10.404 | ואצח |  |
| 9:10.405 | בחית |  |
| 9:10.406 | תפהם |  |
| 9:10.407 | מנה |  |
| 9:10.408 | הדה |  |
| 9:10.409 | אלעקידה |  |
| 9:10.410 | צראחה | , |
| 9:10.411 | ולא |  |
| 9:10.412 | יבקש |  |
| 9:10.413 | פיהא |  |
| 9:10.414 | שך |  |
| 9:10.415 | ויביין |  |
| 9:10.416 | מוסי | ' |
| 9:10.417 | אלאחכאם |  |
| 9:10.418 | אלתי |  |
| 9:10.419 | הי |  |
| 9:10.420 | ענד |  |
| 9:10.421 | מקרס |  |
| 9:10.422 | אהל |  |
| 9:10.423 | אלתתלית |  |
| 9:10.424 | צעיפה |  |
| 9:10.425 | נאקצה |  |
| 9:10.426 | גדא |  |
| 9:10.427 | באלתשריח |  |
| 9:10.428 | אלאתם | . |
| 9:10.429 | ויכררהא |  |
| 9:10.430 | מרה |  |
| 9:10.431 | בעד |  |
| 9:10.432 | אולי |  |
| 9:10.433 | וכרה |  |
| 9:10.434 | בעד |  |
| 9:10.435 | אכרא | , |
| 9:10.436 | ויאכד |  |
| 9:10.437 | ען |  |
| 9:10.438 | מחאפצהתא |  |
| 9:10.439 | תאכידא |  |
| 9:10.440 | בלינא | , |
| 9:10.441 | ויוגב |  |
| 9:10.442 | אלקתל |  |
| 9:10.443 | למן |  |
| 9:10.444 | לם |  |
| 9:10.445 | יעתקדהא |  |
| 9:10.446 | עכ"ל | . |
| 9:10.447 | העתק |  |
| 9:10.448 | ועיין |  |
| 9:10.449 | בסוף |  |
| 9:10.450 | הספר | ) |
| 9:10.451 | וטענתו |  |
| 9:10.452 | זאת |  |
| 9:10.453 | בעצמה |  |
| 9:10.454 | אמיתית |  |
| 9:10.455 | צודקת |  |
| 9:10.456 | על |  |
| 9:10.457 | הקבלה |  |
| 9:10.458 | החדשה | , |
| 9:10.459 | ומדבריו |  |
| 9:10.460 | מוכח |  |
| 9:10.461 | שלא |  |
| 9:10.462 | פשטה |  |
| 9:10.463 | עדיין |  |
| 9:10.464 | אמונת |  |
| 9:10.465 | הקבלה |  |
| 9:10.466 | החדשה |  |
| 9:10.467 | בישראל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

