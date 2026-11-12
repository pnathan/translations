# Translation pack: milhamot-hashem perek-1-05 (Milchamot ha-Shem 1:7–1:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-05`.

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
- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
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
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
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


## Milchamot ha-Shem 1:7

Hebrew (Wikisource): כל שכן וקל וחומר באלהינו ברוך הוא וברוך שמו שתורה הקפידה שלא לטפול ולשתף עמו שום דבר מכל הנבראים בין גשמי בין רוחני כמו שנאמר זובח לאלהים יחרם בלתי לה' לבדו, ואמר רשב״י התנא ע״ה כל המשתף שם שמים ודבר אחר נעקר מן העולם, ולא אמרינן כלא חד כיון שהתורה ביארה בהדיא בכמה מקומות שהוא אחד מיוחד מכל שאר האחדים הנה כן המזכירים את ה' אל דמי לכם מלהודיע אחדותו ית' לבניכם ולתלמידיכם כמצווה עליכם.

| id | word | punctuation after |
|---|---|---|
| 1:7.1 | כל |  |
| 1:7.2 | שכן |  |
| 1:7.3 | וקל |  |
| 1:7.4 | וחומר |  |
| 1:7.5 | באלהינו |  |
| 1:7.6 | ברוך |  |
| 1:7.7 | הוא |  |
| 1:7.8 | וברוך |  |
| 1:7.9 | שמו |  |
| 1:7.10 | שתורה |  |
| 1:7.11 | הקפידה |  |
| 1:7.12 | שלא |  |
| 1:7.13 | לטפול |  |
| 1:7.14 | ולשתף |  |
| 1:7.15 | עמו |  |
| 1:7.16 | שום |  |
| 1:7.17 | דבר |  |
| 1:7.18 | מכל |  |
| 1:7.19 | הנבראים |  |
| 1:7.20 | בין |  |
| 1:7.21 | גשמי |  |
| 1:7.22 | בין |  |
| 1:7.23 | רוחני |  |
| 1:7.24 | כמו |  |
| 1:7.25 | שנאמר |  |
| 1:7.26 | זובח |  |
| 1:7.27 | לאלהים |  |
| 1:7.28 | יחרם |  |
| 1:7.29 | בלתי |  |
| 1:7.30 | לה | ' |
| 1:7.31 | לבדו | , |
| 1:7.32 | ואמר |  |
| 1:7.33 | רשב״י |  |
| 1:7.34 | התנא |  |
| 1:7.35 | ע״ה |  |
| 1:7.36 | כל |  |
| 1:7.37 | המשתף |  |
| 1:7.38 | שם |  |
| 1:7.39 | שמים |  |
| 1:7.40 | ודבר |  |
| 1:7.41 | אחר |  |
| 1:7.42 | נעקר |  |
| 1:7.43 | מן |  |
| 1:7.44 | העולם | , |
| 1:7.45 | ולא |  |
| 1:7.46 | אמרינן |  |
| 1:7.47 | כלא |  |
| 1:7.48 | חד |  |
| 1:7.49 | כיון |  |
| 1:7.50 | שהתורה |  |
| 1:7.51 | ביארה |  |
| 1:7.52 | בהדיא |  |
| 1:7.53 | בכמה |  |
| 1:7.54 | מקומות |  |
| 1:7.55 | שהוא |  |
| 1:7.56 | אחד |  |
| 1:7.57 | מיוחד |  |
| 1:7.58 | מכל |  |
| 1:7.59 | שאר |  |
| 1:7.60 | האחדים |  |
| 1:7.61 | הנה |  |
| 1:7.62 | כן |  |
| 1:7.63 | המזכירים |  |
| 1:7.64 | את |  |
| 1:7.65 | ה | ' |
| 1:7.66 | אל |  |
| 1:7.67 | דמי |  |
| 1:7.68 | לכם |  |
| 1:7.69 | מלהודיע |  |
| 1:7.70 | אחדותו |  |
| 1:7.71 | ית | ' |
| 1:7.72 | לבניכם |  |
| 1:7.73 | ולתלמידיכם |  |
| 1:7.74 | כמצווה |  |
| 1:7.75 | עליכם | . |


## Milchamot ha-Shem 1:8

Hebrew (Wikisource): ולפי שתחלת החבור הזה נעשה בראשונה על ידי שאלת תלמיד אחד את מי נעבוד על כן הנחתי התחלתו שאלת התלמיד ותשובתו מאחד הלומדים המתעקש פעם ראשונה ושנייה ושלישית.

| id | word | punctuation after |
|---|---|---|
| 1:8.1 | ולפי |  |
| 1:8.2 | שתחלת |  |
| 1:8.3 | החבור |  |
| 1:8.4 | הזה |  |
| 1:8.5 | נעשה |  |
| 1:8.6 | בראשונה |  |
| 1:8.7 | על |  |
| 1:8.8 | ידי |  |
| 1:8.9 | שאלת |  |
| 1:8.10 | תלמיד |  |
| 1:8.11 | אחד |  |
| 1:8.12 | את |  |
| 1:8.13 | מי |  |
| 1:8.14 | נעבוד |  |
| 1:8.15 | על |  |
| 1:8.16 | כן |  |
| 1:8.17 | הנחתי |  |
| 1:8.18 | התחלתו |  |
| 1:8.19 | שאלת |  |
| 1:8.20 | התלמיד |  |
| 1:8.21 | ותשובתו |  |
| 1:8.22 | מאחד |  |
| 1:8.23 | הלומדים |  |
| 1:8.24 | המתעקש |  |
| 1:8.25 | פעם |  |
| 1:8.26 | ראשונה |  |
| 1:8.27 | ושנייה |  |
| 1:8.28 | ושלישית | . |


## Milchamot ha-Shem 1:9

Hebrew (Wikisource): הצעיר יחיה בן סלי' אלקאפח יצ״ו

| id | word | punctuation after |
|---|---|---|
| 1:9.1 | הצעיר |  |
| 1:9.2 | יחיה |  |
| 1:9.3 | בן |  |
| 1:9.4 | סלי | ' |
| 1:9.5 | אלקאפח |  |
| 1:9.6 | יצ״ו |  |


## Milchamot ha-Shem 1:10

Hebrew (Wikisource): הקדמת המו"ל.

| id | word | punctuation after |
|---|---|---|
| 1:10.1 | הקדמת |  |
| 1:10.2 | המו"ל | . |


## Milchamot ha-Shem 1:11

Hebrew (Wikisource): ספר מלחמת ה' המודיע דרך אמתות יחודו יתברך עפ"י התורה ומפרשיה רז"ל במשנה ובתלמוד בבלי וירושלמי ומדרשי רז"ל וכדעת הגאונים רס"ג ורבינו בחיי בחובת הלבבות ור' יהודה הלוי בשיריו ובכוזרי הרמב"ם וסמ"ג ובעל העיקרים ז"ל וכמקובל אצל כל איש ישראל ובו אזהרה חמורה שלא לעבוד שום נברא כפי קבלת כל נביאינו וכל חכמינו התנאים והאמוראים והגאונים והפוסקים הראשונים מפי משה ואהרן וזקנים לא לזוז מהם ולא להוסיף עליהם.

| id | word | punctuation after |
|---|---|---|
| 1:11.1 | ספר |  |
| 1:11.2 | מלחמת |  |
| 1:11.3 | ה | ' |
| 1:11.4 | המודיע |  |
| 1:11.5 | דרך |  |
| 1:11.6 | אמתות |  |
| 1:11.7 | יחודו |  |
| 1:11.8 | יתברך |  |
| 1:11.9 | עפ"י |  |
| 1:11.10 | התורה |  |
| 1:11.11 | ומפרשיה |  |
| 1:11.12 | רז"ל |  |
| 1:11.13 | במשנה |  |
| 1:11.14 | ובתלמוד |  |
| 1:11.15 | בבלי |  |
| 1:11.16 | וירושלמי |  |
| 1:11.17 | ומדרשי |  |
| 1:11.18 | רז"ל |  |
| 1:11.19 | וכדעת |  |
| 1:11.20 | הגאונים |  |
| 1:11.21 | רס"ג |  |
| 1:11.22 | ורבינו |  |
| 1:11.23 | בחיי |  |
| 1:11.24 | בחובת |  |
| 1:11.25 | הלבבות |  |
| 1:11.26 | ור | ' |
| 1:11.27 | יהודה |  |
| 1:11.28 | הלוי |  |
| 1:11.29 | בשיריו |  |
| 1:11.30 | ובכוזרי |  |
| 1:11.31 | הרמב"ם |  |
| 1:11.32 | וסמ"ג |  |
| 1:11.33 | ובעל |  |
| 1:11.34 | העיקרים |  |
| 1:11.35 | ז"ל |  |
| 1:11.36 | וכמקובל |  |
| 1:11.37 | אצל |  |
| 1:11.38 | כל |  |
| 1:11.39 | איש |  |
| 1:11.40 | ישראל |  |
| 1:11.41 | ובו |  |
| 1:11.42 | אזהרה |  |
| 1:11.43 | חמורה |  |
| 1:11.44 | שלא |  |
| 1:11.45 | לעבוד |  |
| 1:11.46 | שום |  |
| 1:11.47 | נברא |  |
| 1:11.48 | כפי |  |
| 1:11.49 | קבלת |  |
| 1:11.50 | כל |  |
| 1:11.51 | נביאינו |  |
| 1:11.52 | וכל |  |
| 1:11.53 | חכמינו |  |
| 1:11.54 | התנאים |  |
| 1:11.55 | והאמוראים |  |
| 1:11.56 | והגאונים |  |
| 1:11.57 | והפוסקים |  |
| 1:11.58 | הראשונים |  |
| 1:11.59 | מפי |  |
| 1:11.60 | משה |  |
| 1:11.61 | ואהרן |  |
| 1:11.62 | וזקנים |  |
| 1:11.63 | לא |  |
| 1:11.64 | לזוז |  |
| 1:11.65 | מהם |  |
| 1:11.66 | ולא |  |
| 1:11.67 | להוסיף |  |
| 1:11.68 | עליהם | . |


## Milchamot ha-Shem 1:12

Hebrew (Wikisource): ודע לך קורא נכבד כי לא תשיג כונת הספר הזה אלא אחרי שתקרא ותשנה את כלו כי מחברו עשה לילות כימים ויגע ודקדק בכמה ספרים ודלה מימים ונהרים ראיות כראי מוצק חזקים ולא מדעתו יצאו הדברים ולכן אבקשך קורא נכבד שלא תבוא באגרונו של עשו בהתולים אלא בקולו של יעקב ולהתבונן בדברים, ומי שהוא חסר אונים יאלם מדברים וידום ואל שדי יתן לכם רחמים וזכות המחבר תגן עלינו וע"כ עמו ישראל.

| id | word | punctuation after |
|---|---|---|
| 1:12.1 | ודע |  |
| 1:12.2 | לך |  |
| 1:12.3 | קורא |  |
| 1:12.4 | נכבד |  |
| 1:12.5 | כי |  |
| 1:12.6 | לא |  |
| 1:12.7 | תשיג |  |
| 1:12.8 | כונת |  |
| 1:12.9 | הספר |  |
| 1:12.10 | הזה |  |
| 1:12.11 | אלא |  |
| 1:12.12 | אחרי |  |
| 1:12.13 | שתקרא |  |
| 1:12.14 | ותשנה |  |
| 1:12.15 | את |  |
| 1:12.16 | כלו |  |
| 1:12.17 | כי |  |
| 1:12.18 | מחברו |  |
| 1:12.19 | עשה |  |
| 1:12.20 | לילות |  |
| 1:12.21 | כימים |  |
| 1:12.22 | ויגע |  |
| 1:12.23 | ודקדק |  |
| 1:12.24 | בכמה |  |
| 1:12.25 | ספרים |  |
| 1:12.26 | ודלה |  |
| 1:12.27 | מימים |  |
| 1:12.28 | ונהרים |  |
| 1:12.29 | ראיות |  |
| 1:12.30 | כראי |  |
| 1:12.31 | מוצק |  |
| 1:12.32 | חזקים |  |
| 1:12.33 | ולא |  |
| 1:12.34 | מדעתו |  |
| 1:12.35 | יצאו |  |
| 1:12.36 | הדברים |  |
| 1:12.37 | ולכן |  |
| 1:12.38 | אבקשך |  |
| 1:12.39 | קורא |  |
| 1:12.40 | נכבד |  |
| 1:12.41 | שלא |  |
| 1:12.42 | תבוא |  |
| 1:12.43 | באגרונו |  |
| 1:12.44 | של |  |
| 1:12.45 | עשו |  |
| 1:12.46 | בהתולים |  |
| 1:12.47 | אלא |  |
| 1:12.48 | בקולו |  |
| 1:12.49 | של |  |
| 1:12.50 | יעקב |  |
| 1:12.51 | ולהתבונן |  |
| 1:12.52 | בדברים | , |
| 1:12.53 | ומי |  |
| 1:12.54 | שהוא |  |
| 1:12.55 | חסר |  |
| 1:12.56 | אונים |  |
| 1:12.57 | יאלם |  |
| 1:12.58 | מדברים |  |
| 1:12.59 | וידום |  |
| 1:12.60 | ואל |  |
| 1:12.61 | שדי |  |
| 1:12.62 | יתן |  |
| 1:12.63 | לכם |  |
| 1:12.64 | רחמים |  |
| 1:12.65 | וזכות |  |
| 1:12.66 | המחבר |  |
| 1:12.67 | תגן |  |
| 1:12.68 | עלינו |  |
| 1:12.69 | וע"כ |  |
| 1:12.70 | עמו |  |
| 1:12.71 | ישראל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

