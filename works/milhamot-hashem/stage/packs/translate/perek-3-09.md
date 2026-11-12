# Translation pack: milhamot-hashem perek-3-09 (Milchamot ha-Shem 3:14–3:15)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-3-09.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-3-09`.

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
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
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


## Milchamot ha-Shem 3:14

Hebrew (Wikisource): יד) ראה והבן ידידנו קורא נעלם איך כמה"ר תם הנז"ל צוח ככרוכיא על העוסקים בקבלה החדשה, כי פרצו גדריה של תורה, וערו ערו יסודותיה ושרשיה ביחוד השי"ת כמו שיתבאר להלן, ותחת אשר חשבו להתקרב לשי"ת תעו מן הדרך ובאו והאמין באלהים רבים אחרים ולעבדם! וכדברי ר' תם ן' יחיא כ' גם כן מהר"י אלבו בס' העיקרים פרק כ"ח משני. ועוד האריך בספר רביד הזהב והביא ג"כ דברי הר' הגדול מהרש"ל בתשובותיו ובחבור ים של שלמה על המתיהרים בחידושים שלומדים בקבלה, וכתב דמאחר שנחתם התלמוד אין לנו להחמיר נגד הגמרא והוי כמו מינות כמבואר בחק יעקב סי' תפ"ט ס"ק י"א באורך. וכל זה הוא לפי ההנחה והיחוס שייחסו הזהר לרשב"י התנא, אך אחרי הבקורת שבקרו וחקרו חכמי לב אחריו מבטן מי יצאו הדברים הנזכרים בו ואימתי יצא החבור הנזכר ומיד מי יצא. ונתבאר באר היטב כי הוא חיבור אחד האחרונים ולא יצאו אותן הרעיונים מרשב"י התנא ע"ה כלל, כי נמצא בו כמה דברים בעיקרי אמונת תורתינו הקדושה היפך דברי רשב"י בתלמוד ובמדרשי רז"ל המפורסמים בישראל, ודאי שאין לסמוך עליו בשום דבר הנוגע למנהגי ישראל שנהגו בהם קדמונינו. ולא יפה עשו חכמי ימי הבינים ששינו עלינו כמה דברים בסידורי התפלות על פי הזהר והמקובלים, כי שינויים ההם יש מהם מביאים לידי מינות להאמין ברשויות הרבה ולעבוד אלהים אחרים כמו שיתבאר להלן.

| id | word | punctuation after |
|---|---|---|
| 3:14.1 | יד | ) |
| 3:14.2 | ראה |  |
| 3:14.3 | והבן |  |
| 3:14.4 | ידידנו |  |
| 3:14.5 | קורא |  |
| 3:14.6 | נעלם |  |
| 3:14.7 | איך |  |
| 3:14.8 | כמה"ר |  |
| 3:14.9 | תם |  |
| 3:14.10 | הנז"ל |  |
| 3:14.11 | צוח |  |
| 3:14.12 | ככרוכיא |  |
| 3:14.13 | על |  |
| 3:14.14 | העוסקים |  |
| 3:14.15 | בקבלה |  |
| 3:14.16 | החדשה | , |
| 3:14.17 | כי |  |
| 3:14.18 | פרצו |  |
| 3:14.19 | גדריה |  |
| 3:14.20 | של |  |
| 3:14.21 | תורה | , |
| 3:14.22 | וערו |  |
| 3:14.23 | ערו |  |
| 3:14.24 | יסודותיה |  |
| 3:14.25 | ושרשיה |  |
| 3:14.26 | ביחוד |  |
| 3:14.27 | השי"ת |  |
| 3:14.28 | כמו |  |
| 3:14.29 | שיתבאר |  |
| 3:14.30 | להלן | , |
| 3:14.31 | ותחת |  |
| 3:14.32 | אשר |  |
| 3:14.33 | חשבו |  |
| 3:14.34 | להתקרב |  |
| 3:14.35 | לשי"ת |  |
| 3:14.36 | תעו |  |
| 3:14.37 | מן |  |
| 3:14.38 | הדרך |  |
| 3:14.39 | ובאו |  |
| 3:14.40 | והאמין |  |
| 3:14.41 | באלהים |  |
| 3:14.42 | רבים |  |
| 3:14.43 | אחרים |  |
| 3:14.44 | ולעבדם | ! |
| 3:14.45 | וכדברי |  |
| 3:14.46 | ר | ' |
| 3:14.47 | תם |  |
| 3:14.48 | ן | ' |
| 3:14.49 | יחיא |  |
| 3:14.50 | כ | ' |
| 3:14.51 | גם |  |
| 3:14.52 | כן |  |
| 3:14.53 | מהר"י |  |
| 3:14.54 | אלבו |  |
| 3:14.55 | בס | ' |
| 3:14.56 | העיקרים |  |
| 3:14.57 | פרק |  |
| 3:14.58 | כ"ח |  |
| 3:14.59 | משני | . |
| 3:14.60 | ועוד |  |
| 3:14.61 | האריך |  |
| 3:14.62 | בספר |  |
| 3:14.63 | רביד |  |
| 3:14.64 | הזהב |  |
| 3:14.65 | והביא |  |
| 3:14.66 | ג"כ |  |
| 3:14.67 | דברי |  |
| 3:14.68 | הר | ' |
| 3:14.69 | הגדול |  |
| 3:14.70 | מהרש"ל |  |
| 3:14.71 | בתשובותיו |  |
| 3:14.72 | ובחבור |  |
| 3:14.73 | ים |  |
| 3:14.74 | של |  |
| 3:14.75 | שלמה |  |
| 3:14.76 | על |  |
| 3:14.77 | המתיהרים |  |
| 3:14.78 | בחידושים |  |
| 3:14.79 | שלומדים |  |
| 3:14.80 | בקבלה | , |
| 3:14.81 | וכתב |  |
| 3:14.82 | דמאחר |  |
| 3:14.83 | שנחתם |  |
| 3:14.84 | התלמוד |  |
| 3:14.85 | אין |  |
| 3:14.86 | לנו |  |
| 3:14.87 | להחמיר |  |
| 3:14.88 | נגד |  |
| 3:14.89 | הגמרא |  |
| 3:14.90 | והוי |  |
| 3:14.91 | כמו |  |
| 3:14.92 | מינות |  |
| 3:14.93 | כמבואר |  |
| 3:14.94 | בחק |  |
| 3:14.95 | יעקב |  |
| 3:14.96 | סי | ' |
| 3:14.97 | תפ"ט |  |
| 3:14.98 | ס"ק |  |
| 3:14.99 | י"א |  |
| 3:14.100 | באורך | . |
| 3:14.101 | וכל |  |
| 3:14.102 | זה |  |
| 3:14.103 | הוא |  |
| 3:14.104 | לפי |  |
| 3:14.105 | ההנחה |  |
| 3:14.106 | והיחוס |  |
| 3:14.107 | שייחסו |  |
| 3:14.108 | הזהר |  |
| 3:14.109 | לרשב"י |  |
| 3:14.110 | התנא | , |
| 3:14.111 | אך |  |
| 3:14.112 | אחרי |  |
| 3:14.113 | הבקורת |  |
| 3:14.114 | שבקרו |  |
| 3:14.115 | וחקרו |  |
| 3:14.116 | חכמי |  |
| 3:14.117 | לב |  |
| 3:14.118 | אחריו |  |
| 3:14.119 | מבטן |  |
| 3:14.120 | מי |  |
| 3:14.121 | יצאו |  |
| 3:14.122 | הדברים |  |
| 3:14.123 | הנזכרים |  |
| 3:14.124 | בו |  |
| 3:14.125 | ואימתי |  |
| 3:14.126 | יצא |  |
| 3:14.127 | החבור |  |
| 3:14.128 | הנזכר |  |
| 3:14.129 | ומיד |  |
| 3:14.130 | מי |  |
| 3:14.131 | יצא | . |
| 3:14.132 | ונתבאר |  |
| 3:14.133 | באר |  |
| 3:14.134 | היטב |  |
| 3:14.135 | כי |  |
| 3:14.136 | הוא |  |
| 3:14.137 | חיבור |  |
| 3:14.138 | אחד |  |
| 3:14.139 | האחרונים |  |
| 3:14.140 | ולא |  |
| 3:14.141 | יצאו |  |
| 3:14.142 | אותן |  |
| 3:14.143 | הרעיונים |  |
| 3:14.144 | מרשב"י |  |
| 3:14.145 | התנא |  |
| 3:14.146 | ע"ה |  |
| 3:14.147 | כלל | , |
| 3:14.148 | כי |  |
| 3:14.149 | נמצא |  |
| 3:14.150 | בו |  |
| 3:14.151 | כמה |  |
| 3:14.152 | דברים |  |
| 3:14.153 | בעיקרי |  |
| 3:14.154 | אמונת |  |
| 3:14.155 | תורתינו |  |
| 3:14.156 | הקדושה |  |
| 3:14.157 | היפך |  |
| 3:14.158 | דברי |  |
| 3:14.159 | רשב"י |  |
| 3:14.160 | בתלמוד |  |
| 3:14.161 | ובמדרשי |  |
| 3:14.162 | רז"ל |  |
| 3:14.163 | המפורסמים |  |
| 3:14.164 | בישראל | , |
| 3:14.165 | ודאי |  |
| 3:14.166 | שאין |  |
| 3:14.167 | לסמוך |  |
| 3:14.168 | עליו |  |
| 3:14.169 | בשום |  |
| 3:14.170 | דבר |  |
| 3:14.171 | הנוגע |  |
| 3:14.172 | למנהגי |  |
| 3:14.173 | ישראל |  |
| 3:14.174 | שנהגו |  |
| 3:14.175 | בהם |  |
| 3:14.176 | קדמונינו | . |
| 3:14.177 | ולא |  |
| 3:14.178 | יפה |  |
| 3:14.179 | עשו |  |
| 3:14.180 | חכמי |  |
| 3:14.181 | ימי |  |
| 3:14.182 | הבינים |  |
| 3:14.183 | ששינו |  |
| 3:14.184 | עלינו |  |
| 3:14.185 | כמה |  |
| 3:14.186 | דברים |  |
| 3:14.187 | בסידורי |  |
| 3:14.188 | התפלות |  |
| 3:14.189 | על |  |
| 3:14.190 | פי |  |
| 3:14.191 | הזהר |  |
| 3:14.192 | והמקובלים | , |
| 3:14.193 | כי |  |
| 3:14.194 | שינויים |  |
| 3:14.195 | ההם |  |
| 3:14.196 | יש |  |
| 3:14.197 | מהם |  |
| 3:14.198 | מביאים |  |
| 3:14.199 | לידי |  |
| 3:14.200 | מינות |  |
| 3:14.201 | להאמין |  |
| 3:14.202 | ברשויות |  |
| 3:14.203 | הרבה |  |
| 3:14.204 | ולעבוד |  |
| 3:14.205 | אלהים |  |
| 3:14.206 | אחרים |  |
| 3:14.207 | כמו |  |
| 3:14.208 | שיתבאר |  |
| 3:14.209 | להלן | . |


## Milchamot ha-Shem 3:15

Hebrew (Wikisource): טו) גם הביא עוד הר' רביד הזהב דברי מרן בב"י דחייש לדברי הזהר במקום שלא נזכר בגמרא היפך מדבריו, וכתב עליו הרמ"א ז"ל שאע"פ שלא נזכר בגמרא היפך דבריו, אין לזה מדברי הפוסקים, ואין לחוש לדברי הזהר ע"ש. עוד הביא דברי הרדב"ז ז"ל שכתב שכל דבר שנזכר בגמרא או באחד מן הפוסקים, או בעלי ההלכות אפי' שיהיה היפך ממה שכתוב בספרי הקבלה אני מורה בו בין להקל בין להחמיר, ולא אחוש למה שכתוב באותן הספרים. ולעצמי אם הוא חומרא אני נוהג אותו ואם קולא לא אחוש לה. ועוד כתב בשם הרא"ם ושאר פוסקים שאין לחוש לדברי המקובלים נגד שום פוסק ע"ש. אמור מעתה שמדברי הפוסקים הנזכרים בתשובת הרב רבי הזהב והם מרן ז"ל ומור"ם והרדב"ז וכנה"ג ומהר"י הלוי כ' מהר"ש הלוי ור' אליהו הלוי, והר"מ אלקבולי ור"ת בן יחיא והר' חק יעקב ומהרש"ל ומהר"ם אלשקר שאומרים שאין לחוש לדברי הזהר והמקובלים שהוא נגד המשנה והתלמוד ואפילו נגד פוסק יחידי. הנה מזה הוכיח לעצמך שכתבת שהם הלכה למשה מסיני, שדבריך בדותא וסכלות גמורה.

| id | word | punctuation after |
|---|---|---|
| 3:15.1 | טו | ) |
| 3:15.2 | גם |  |
| 3:15.3 | הביא |  |
| 3:15.4 | עוד |  |
| 3:15.5 | הר | ' |
| 3:15.6 | רביד |  |
| 3:15.7 | הזהב |  |
| 3:15.8 | דברי |  |
| 3:15.9 | מרן |  |
| 3:15.10 | בב"י |  |
| 3:15.11 | דחייש |  |
| 3:15.12 | לדברי |  |
| 3:15.13 | הזהר |  |
| 3:15.14 | במקום |  |
| 3:15.15 | שלא |  |
| 3:15.16 | נזכר |  |
| 3:15.17 | בגמרא |  |
| 3:15.18 | היפך |  |
| 3:15.19 | מדבריו | , |
| 3:15.20 | וכתב |  |
| 3:15.21 | עליו |  |
| 3:15.22 | הרמ"א |  |
| 3:15.23 | ז"ל |  |
| 3:15.24 | שאע"פ |  |
| 3:15.25 | שלא |  |
| 3:15.26 | נזכר |  |
| 3:15.27 | בגמרא |  |
| 3:15.28 | היפך |  |
| 3:15.29 | דבריו | , |
| 3:15.30 | אין |  |
| 3:15.31 | לזה |  |
| 3:15.32 | מדברי |  |
| 3:15.33 | הפוסקים | , |
| 3:15.34 | ואין |  |
| 3:15.35 | לחוש |  |
| 3:15.36 | לדברי |  |
| 3:15.37 | הזהר |  |
| 3:15.38 | ע"ש | . |
| 3:15.39 | עוד |  |
| 3:15.40 | הביא |  |
| 3:15.41 | דברי |  |
| 3:15.42 | הרדב"ז |  |
| 3:15.43 | ז"ל |  |
| 3:15.44 | שכתב |  |
| 3:15.45 | שכל |  |
| 3:15.46 | דבר |  |
| 3:15.47 | שנזכר |  |
| 3:15.48 | בגמרא |  |
| 3:15.49 | או |  |
| 3:15.50 | באחד |  |
| 3:15.51 | מן |  |
| 3:15.52 | הפוסקים | , |
| 3:15.53 | או |  |
| 3:15.54 | בעלי |  |
| 3:15.55 | ההלכות |  |
| 3:15.56 | אפי | ' |
| 3:15.57 | שיהיה |  |
| 3:15.58 | היפך |  |
| 3:15.59 | ממה |  |
| 3:15.60 | שכתוב |  |
| 3:15.61 | בספרי |  |
| 3:15.62 | הקבלה |  |
| 3:15.63 | אני |  |
| 3:15.64 | מורה |  |
| 3:15.65 | בו |  |
| 3:15.66 | בין |  |
| 3:15.67 | להקל |  |
| 3:15.68 | בין |  |
| 3:15.69 | להחמיר | , |
| 3:15.70 | ולא |  |
| 3:15.71 | אחוש |  |
| 3:15.72 | למה |  |
| 3:15.73 | שכתוב |  |
| 3:15.74 | באותן |  |
| 3:15.75 | הספרים | . |
| 3:15.76 | ולעצמי |  |
| 3:15.77 | אם |  |
| 3:15.78 | הוא |  |
| 3:15.79 | חומרא |  |
| 3:15.80 | אני |  |
| 3:15.81 | נוהג |  |
| 3:15.82 | אותו |  |
| 3:15.83 | ואם |  |
| 3:15.84 | קולא |  |
| 3:15.85 | לא |  |
| 3:15.86 | אחוש |  |
| 3:15.87 | לה | . |
| 3:15.88 | ועוד |  |
| 3:15.89 | כתב |  |
| 3:15.90 | בשם |  |
| 3:15.91 | הרא"ם |  |
| 3:15.92 | ושאר |  |
| 3:15.93 | פוסקים |  |
| 3:15.94 | שאין |  |
| 3:15.95 | לחוש |  |
| 3:15.96 | לדברי |  |
| 3:15.97 | המקובלים |  |
| 3:15.98 | נגד |  |
| 3:15.99 | שום |  |
| 3:15.100 | פוסק |  |
| 3:15.101 | ע"ש | . |
| 3:15.102 | אמור |  |
| 3:15.103 | מעתה |  |
| 3:15.104 | שמדברי |  |
| 3:15.105 | הפוסקים |  |
| 3:15.106 | הנזכרים |  |
| 3:15.107 | בתשובת |  |
| 3:15.108 | הרב |  |
| 3:15.109 | רבי |  |
| 3:15.110 | הזהב |  |
| 3:15.111 | והם |  |
| 3:15.112 | מרן |  |
| 3:15.113 | ז"ל |  |
| 3:15.114 | ומור"ם |  |
| 3:15.115 | והרדב"ז |  |
| 3:15.116 | וכנה"ג |  |
| 3:15.117 | ומהר"י |  |
| 3:15.118 | הלוי |  |
| 3:15.119 | כ | ' |
| 3:15.120 | מהר"ש |  |
| 3:15.121 | הלוי |  |
| 3:15.122 | ור | ' |
| 3:15.123 | אליהו |  |
| 3:15.124 | הלוי | , |
| 3:15.125 | והר"מ |  |
| 3:15.126 | אלקבולי |  |
| 3:15.127 | ור"ת |  |
| 3:15.128 | בן |  |
| 3:15.129 | יחיא |  |
| 3:15.130 | והר | ' |
| 3:15.131 | חק |  |
| 3:15.132 | יעקב |  |
| 3:15.133 | ומהרש"ל |  |
| 3:15.134 | ומהר"ם |  |
| 3:15.135 | אלשקר |  |
| 3:15.136 | שאומרים |  |
| 3:15.137 | שאין |  |
| 3:15.138 | לחוש |  |
| 3:15.139 | לדברי |  |
| 3:15.140 | הזהר |  |
| 3:15.141 | והמקובלים |  |
| 3:15.142 | שהוא |  |
| 3:15.143 | נגד |  |
| 3:15.144 | המשנה |  |
| 3:15.145 | והתלמוד |  |
| 3:15.146 | ואפילו |  |
| 3:15.147 | נגד |  |
| 3:15.148 | פוסק |  |
| 3:15.149 | יחידי | . |
| 3:15.150 | הנה |  |
| 3:15.151 | מזה |  |
| 3:15.152 | הוכיח |  |
| 3:15.153 | לעצמך |  |
| 3:15.154 | שכתבת |  |
| 3:15.155 | שהם |  |
| 3:15.156 | הלכה |  |
| 3:15.157 | למשה |  |
| 3:15.158 | מסיני | , |
| 3:15.159 | שדבריך |  |
| 3:15.160 | בדותא |  |
| 3:15.161 | וסכלות |  |
| 3:15.162 | גמורה | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

