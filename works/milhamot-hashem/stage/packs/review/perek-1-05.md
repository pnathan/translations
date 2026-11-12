# Review pack: milhamot-hashem perek-1-05 (Milchamot ha-Shem 1:7–1:12)

Brief: `briefs/review.md`. Write: `works/milhamot-hashem/stage/review/perek-1-05.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-05`.

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


### Translation under review (1:7)

English: All the more so, a fortiori, regarding our God, blessed be He, and blessed be His name, that the Torah insisted that one should not attach to Him or associate with Him anything of all created things, whether corporeal or spiritual, as it is said: "He who sacrifices to the gods shall be destroyed, save to the LORD alone." And Rabbi Shimon bar Yochai the Tanna, peace be upon him, said: whoever associates the name of Heaven with something else is uprooted from the world. and we should not say that He is one as the others are one, since the Torah has made clear plainly in many places that He is one, unique, apart from all the other ones. Thus, you who mention the LORD should not think it right to refrain from making known His oneness, may He be blessed, to your children and your disciples, as you are commanded.

Literal: all the-more-so and-light-and-heavy in-our-God blessed He and-blessed His-name that-the-Torah insisted that-not to-attach and-to-associate with-Him any thing of-all the-created-beings whether corporeal or spiritual as it-is-said he-who-sacrifices to-gods shall-be-devoted-to-destruction except to-the-LORD alone and-he-said Rabbi-Shimon-bar-Yochai the-Tanna peace-be-upon-him everyone-who associates the-name the-name-of-Heaven with-something else is-uprooted from the-world and-not we-say as-one one since the-Torah has-made-clear plainly in-many places that-He one unique from-all the-other the-ones so thus you-who-mention (object) the-LORD let-it-not seem to-you from-making-known His-oneness may-He-be-blessed to-your-children and-to-your-disciples as-is-commanded upon-you

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:7.1 | all |  |  | The two Talmudic forms of the a fortiori argument: כל שכן ('all the more so') and וקל וחומר ('light and heavy'). |
| 1:7.2 | the more so |  |  |  |
| 1:7.3 | and light |  |  |  |
| 1:7.4 | and heavy |  |  |  |
| 1:7.5 | in our God |  |  |  |
| 1:7.6 | blessed |  |  |  |
| 1:7.7 | He |  |  |  |
| 1:7.8 | and blessed |  |  |  |
| 1:7.9 | His name |  |  |  |
| 1:7.10 | that the Torah |  |  |  |
| 1:7.11 | insisted |  |  |  |
| 1:7.12 | that not |  |  |  |
| 1:7.13 | to attach |  |  | The verb is paired with לשתף below; the English renders it as attaching one thing to another. |
| 1:7.14 | and to associate |  |  |  |
| 1:7.15 | with Him |  |  |  |
| 1:7.16 | any |  |  |  |
| 1:7.17 | thing |  |  |  |
| 1:7.18 | of all |  |  |  |
| 1:7.19 | the created beings |  |  |  |
| 1:7.20 | whether |  |  |  |
| 1:7.21 | corporeal |  |  |  |
| 1:7.22 | or |  |  |  |
| 1:7.23 | spiritual |  |  |  |
| 1:7.24 | as |  |  |  |
| 1:7.25 | it is said |  |  |  |
| 1:7.26 | he who sacrifices |  |  |  |
| 1:7.27 | to gods |  |  |  |
| 1:7.28 | shall be devoted to destruction |  |  |  |
| 1:7.29 | except |  |  |  |
| 1:7.30 | to the LORD |  |  |  |
| 1:7.31 | alone |  |  |  |
| 1:7.32 | and he said |  |  |  |
| 1:7.33 | Rabbi Shimon bar Yochai |  |  |  |
| 1:7.34 | the Tanna |  |  |  |
| 1:7.35 | peace be upon him |  |  |  |
| 1:7.36 | everyone who |  |  |  |
| 1:7.37 | associates |  |  |  |
| 1:7.38 | the name |  |  |  |
| 1:7.39 | of Heaven |  |  |  |
| 1:7.40 | with something |  |  |  |
| 1:7.41 | else |  |  |  |
| 1:7.42 | is uprooted |  |  |  |
| 1:7.43 | from |  |  |  |
| 1:7.44 | the world |  |  |  |
| 1:7.45 | and not |  |  |  |
| 1:7.46 | we say |  |  |  |
| 1:7.47 | as one |  |  | כלא חד is an Aramaic-style idiom, literally 'as one'; rendered by its sense in the context. |
| 1:7.48 | one |  |  |  |
| 1:7.49 | since |  |  |  |
| 1:7.50 | the Torah |  |  |  |
| 1:7.51 | has made clear |  |  |  |
| 1:7.52 | plainly |  |  |  |
| 1:7.53 | in many |  |  |  |
| 1:7.54 | places |  |  |  |
| 1:7.55 | that He |  |  |  |
| 1:7.56 | one |  |  |  |
| 1:7.57 | unique |  |  |  |
| 1:7.58 | from all |  |  |  |
| 1:7.59 | other |  |  |  |
| 1:7.60 | the ones |  |  |  |
| 1:7.61 | so |  |  |  |
| 1:7.62 | thus |  |  |  |
| 1:7.63 | you who mention |  |  |  |
| 1:7.64 | (object marker) |  |  |  |
| 1:7.65 | the LORD |  |  | The Tetragrammaton as written, ה׳: 'the LORD'. |
| 1:7.66 | let it not |  |  |  |
| 1:7.67 | seem |  |  |  |
| 1:7.68 | to you |  |  |  |
| 1:7.69 | from making known |  |  | להודיע: to make known. The clause אל דמי לכם מלהודיע is 'let it not seem right to you to refrain from making known'. |
| 1:7.70 | His oneness |  |  |  |
| 1:7.71 | may He be blessed |  |  |  |
| 1:7.72 | to your children |  |  |  |
| 1:7.73 | and to your disciples |  |  |  |
| 1:7.74 | as is commanded |  |  |  |
| 1:7.75 | upon you |  |  |  |

Phrases:
- 1:7.1 1:7.2 1:7.3 1:7.4 | All the more so, a fortiori, | all the-more-so and-light-and-heavy
- 1:7.5 1:7.6 1:7.7 | regarding our God, blessed be He, | in-our-God blessed He
- 1:7.8 1:7.9 | and blessed be His name, | and-blessed His-name
- 1:7.10 1:7.11 1:7.12 1:7.13 1:7.14 1:7.15 | that the Torah insisted that one should not attach to Him or associate with Him | that-the-Torah insisted that-not to-attach and-to-associate with-Him
- 1:7.16 1:7.17 1:7.18 1:7.19 | anything of all created things, | any thing of-all the-created-beings
- 1:7.20 1:7.21 1:7.22 1:7.23 | whether corporeal or spiritual, | whether corporeal or spiritual
- 1:7.24 1:7.25 1:7.26 1:7.27 1:7.28 1:7.29 1:7.30 1:7.31 | as it is said: "He who sacrifices to the gods shall be destroyed, save to the LORD alone." | as it-is-said he-who-sacrifices to-gods shall-be-devoted-to-destruction except to-the-LORD alone
- 1:7.32 1:7.33 1:7.34 1:7.35 | And Rabbi Shimon bar Yochai the Tanna, peace be upon him, said: | and-he-said Rabbi-Shimon-bar-Yochai the-Tanna peace-be-upon-him
- 1:7.36 1:7.37 1:7.38 1:7.39 1:7.40 1:7.41 1:7.42 1:7.43 1:7.44 | whoever associates the name of Heaven with something else is uprooted from the world. | everyone-who associates the-name the-name-of-Heaven with-something else is-uprooted from the-world
- 1:7.45 1:7.46 1:7.47 1:7.48 | and we should not say that He is one as the others are one, | and-not we-say as-one one
- 1:7.49 1:7.50 1:7.51 1:7.52 1:7.53 1:7.54 1:7.55 1:7.56 1:7.57 1:7.58 1:7.59 1:7.60 | since the Torah has made clear plainly in many places that He is one, unique, apart from all the other ones. | since the-Torah has-made-clear plainly in-many places that-He one unique from-all the-other the-ones
- 1:7.61 1:7.62 1:7.63 1:7.64 1:7.65 1:7.66 1:7.67 1:7.68 | Thus, you who mention the LORD should not think it right | so thus you-who-mention (object) the-LORD let-it-not seem to-you
- 1:7.69 1:7.70 1:7.71 | to refrain from making known His oneness, may He be blessed, | from-making-known His-oneness may-He-be-blessed
- 1:7.72 1:7.73 | to your children and your disciples, | to-your-children and-to-your-disciples
- 1:7.74 1:7.75 | as you are commanded. | as-is-commanded upon-you

Grammar notes:
- כל שכן and וקל וחומר are the two Talmudic forms of the a fortiori argument, and the English names the argument once, at the start.
- The שלא clause depends on הקפידה: the Torah insisted that one not attach to God or associate with Him anything created.
- Exod 22:19 is quoted with שנאמר, an explicit quotation, so it takes double quotation marks; the verse itself is in the apparatus.
- ולא אמרינן כלא חד כיון ש… is the reasoning: since the Torah has made clear that He is one, we may not speak of Him as one of the ones.
- הנה כן המזכירים את ה׳ is an address in the second person to the reader; אל דמי לכם is 'it should not seem to you'.

Translator's note: Rendered by sense: כלא חד is the one clause whose sense is taken from its context; the literal is 'as one'.

Commentary rulings:
- apparatus: Exod 22:19, quoted with שנאמר: זובח לאלהים יחרם בלתי לה׳ לבדו. The author writes לה׳ where the verse has ליהוה; the English gives 'the LORD'.

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


### Translation under review (1:8)

English: Because the beginning of this composition was made at first by the question of one student: whom shall we worship? Therefore I placed at its opening the question of the student, and its reply from one of the learners who persisted the first, second and third time.

Literal: and-since that-the-beginning-of this-composition this was-made at-first by-way-of the-hand-of the-question-of a-student one (object) whom we-shall-worship on-so I-placed its-beginning the-question-of the-student and-its-reply from-one of-the-learners who-persisted the-time first and-second and-third

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:8.1 | and since |  |  |  |
| 1:8.2 | that the beginning of |  |  |  |
| 1:8.3 | the composition |  |  |  |
| 1:8.4 | this |  |  |  |
| 1:8.5 | was made |  |  |  |
| 1:8.6 | at first |  |  |  |
| 1:8.7 | by way of |  |  |  |
| 1:8.8 | the hand of |  |  |  |
| 1:8.9 | the question of |  |  |  |
| 1:8.10 | a student |  |  |  |
| 1:8.11 | one |  |  |  |
| 1:8.12 | (object marker) |  |  |  |
| 1:8.13 | whom |  |  | The student's question in direct speech: the Hebrew has no question mark, and the English adds one. |
| 1:8.14 | we shall worship |  |  | את מי נעבוד is the question itself, 'whom shall we worship?' |
| 1:8.15 | therefore |  |  | על כן, two words: 'therefore'. |
| 1:8.16 | so |  |  |  |
| 1:8.17 | I placed |  |  |  |
| 1:8.18 | its beginning |  |  |  |
| 1:8.19 | the question of |  |  |  |
| 1:8.20 | the student |  |  |  |
| 1:8.21 | and its reply |  |  |  |
| 1:8.22 | from one |  |  |  |
| 1:8.23 | of the learners |  |  |  |
| 1:8.24 | who persisted |  |  |  |
| 1:8.25 | the time |  |  |  |
| 1:8.26 | first |  |  |  |
| 1:8.27 | and second |  |  |  |
| 1:8.28 | and third |  |  |  |

Phrases:
- 1:8.1 1:8.2 1:8.3 1:8.4 1:8.5 1:8.6 | Because the beginning of this composition was made at first | and-since that-the-beginning-of this-composition this was-made at-first
- 1:8.7 1:8.8 1:8.9 1:8.10 1:8.11 | by the question of one student: | by-way-of the-hand-of the-question-of a-student one
- 1:8.12 1:8.13 1:8.14 | whom shall we worship? | (object) whom we-shall-worship
- 1:8.15 1:8.16 | Therefore | on-so
- 1:8.17 1:8.18 | I placed at its opening | I-placed its-beginning
- 1:8.19 1:8.20 | the question of the student, | the-question-of the-student
- 1:8.21 1:8.22 1:8.23 1:8.24 | and its reply from one of the learners who persisted | and-its-reply from-one of-the-learners who-persisted
- 1:8.25 1:8.26 1:8.27 1:8.28 | the first, second and third time. | the-time first and-second and-third

Grammar notes:
- ולפי שתחלת החבור הזה נעשה… על ידי שאלת תלמיד אחד: the clause gives the reason for the book's form, a dialogue opened by a student's question.
- The student's question את מי נעבוד is a direct question inside the author's sentence; it is kept as a question.
- על כן הנחתי התחלתו: 'therefore I placed its beginning', in the author's first person.
- פעם ראשונה ושנייה ושלישית is an adverbial string: 'the first, second and third time'.

Translator's note: The student's question is kept as a question; the English places the question mark, which the Hebrew does not write.

Commentary rulings:

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


### Translation under review (1:9)

English: The humble Yichya, son of Sali' al-Kafich, may his Rock keep him and give him life.

Literal: the-humble Yichya son-of Sali' al-Kafich may-his-Rock-keep-him-and-give-him-life

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:9.1 | the humble |  |  |  |
| 1:9.2 | Yichya |  |  |  |
| 1:9.3 | son of |  |  |  |
| 1:9.4 | Sali' |  |  | The patronymic is abbreviated in the text and is kept as written. |
| 1:9.5 | al-Kafich |  |  |  |
| 1:9.6 | may his Rock keep him and give him life |  |  | An abbreviated blessing: the letters stand for words of blessing for the author. |

Phrases:
- 1:9.1 1:9.2 | The humble Yichya, | the-humble Yichya
- 1:9.3 1:9.4 1:9.5 | son of Sali' al-Kafich, | son-of Sali' al-Kafich
- 1:9.6 | may his Rock keep him and give him life. | may-his-Rock-keep-him-and-give-him-life

Grammar notes:
- הצעיר is an adjective in apposition to יחיה: the author's self-description, 'the humble'.
- בן סלי׳ אלקאפח is the patronymic and the surname; the abbreviated form is kept as written.
- יצ״ו is an abbreviated blessing in the optative: 'may his Rock keep him and give him life'.

Translator's note: The signature as in the decisions for this work.

Commentary rulings:

## Milchamot ha-Shem 1:10

Hebrew (Wikisource): הקדמת המו"ל.

| id | word | punctuation after |
|---|---|---|
| 1:10.1 | הקדמת |  |
| 1:10.2 | המו"ל | . |


### Translation under review (1:10)

English: The Publisher's Introduction.

Literal: the-introduction of-the-publisher

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:10.1 | the introduction |  |  |  |
| 1:10.2 | the publisher |  |  | המו״ל: the publisher. The voice of the publisher runs through 1:10 to 1:13. |

Phrases:
- 1:10.1 1:10.2 | The Publisher's Introduction. | the-introduction of-the-publisher

Grammar notes:
- הקדמת המו״ל is a construct chain: 'the introduction of the publisher'.
- The heading is in the publisher's voice, not the author's (see the decisions).

Commentary rulings:

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


### Translation under review (1:11)

English: This book, The Wars of the LORD, makes known the way of the truths of His unity, may He be blessed, according to the Torah and its interpreters, our Rabbis of blessed memory, in the Mishnah, in the Bavli and the Yerushalmi, and the Midrashim of our Rabbis of blessed memory, and in accordance with the view of the Geonim, Rabbi Saadia Gaon, and our Rabbi Bahya in Chovot ha-Levavot, and Rabbi Yehuda Halevi in his poems, and in the Kuzari, the Rambam, the SeMaG and the author of the Ikkarim, of blessed memory, and as is accepted by every Jew. In it is a severe warning not to worship any created thing, in accordance with the tradition of all our prophets and all our sages, the Tannaim, the Amoraim, the Geonim and the early decisors, who received it from the mouth of Moses, Aaron and the elders, not to deviate from them and not to add to them.

Literal: the-book the-wars-of the-LORD that-makes-known the-way-of the-truths-of His-unity may-He-be-blessed according-to the-Torah and-its-interpreters our-Rabbis-of-blessed-memory in-the-Mishnah and-in-the-Talmud Bavli and-Yerushalmi and-the-Midrashim-of our-Rabbis-of-blessed-memory and-in-accordance-with the-Geonim Rabbi-Saadia-Gaon and-our-Rabbi Bahya in-Chovot-ha-Levavot and-Rabbi Yehuda Halevi in-his-poems and-in-the-Kuzari the-Rambam and-the-SeMaG and-the-author-of the-Ikkarim of-blessed-memory and-as-is-accepted by every man-of Israel and-in-it a-warning severe that-not to-worship any created-thing in-accordance-with the-tradition-of all our-prophets and-all our-sages the-Tannaim and-the-Amoraim and-the-Geonim and-the-decisors the-early from-the-mouth-of Moses and-Aaron and-the-elders not to-deviate from-them and-not to-add to-them

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:11.1 | the book |  |  | The title: The Wars of the LORD (Num 21:14), as the title page gives it. |
| 1:11.2 | the wars of |  |  |  |
| 1:11.3 | the LORD |  |  |  |
| 1:11.4 | that makes known |  |  |  |
| 1:11.5 | the way of |  |  |  |
| 1:11.6 | the truths of |  |  |  |
| 1:11.7 | His unity |  |  |  |
| 1:11.8 | may He be blessed |  |  |  |
| 1:11.9 | according to |  |  |  |
| 1:11.10 | the Torah |  |  |  |
| 1:11.11 | and its interpreters |  |  |  |
| 1:11.12 | our Rabbis of blessed memory |  |  |  |
| 1:11.13 | in the Mishnah |  |  |  |
| 1:11.14 | and in the Talmud |  |  |  |
| 1:11.15 | Bavli |  |  | Bavli: the Babylonian Talmud. |
| 1:11.16 | and the Yerushalmi |  |  | Yerushalmi: the Jerusalem Talmud. |
| 1:11.17 | and the Midrashim of |  |  |  |
| 1:11.18 | our Rabbis of blessed memory |  |  |  |
| 1:11.19 | and in accordance with |  |  |  |
| 1:11.20 | the Geonim |  |  |  |
| 1:11.21 | Rabbi Saadia Gaon |  |  |  |
| 1:11.22 | and our Rabbi |  |  |  |
| 1:11.23 | Bahya |  |  |  |
| 1:11.24 | in Chovot |  |  |  |
| 1:11.25 | ha-Levavot |  |  | Chovot ha-Levavot, 'Duties of the Hearts', by Bahya. |
| 1:11.26 | and Rabbi |  |  |  |
| 1:11.27 | Yehuda |  |  |  |
| 1:11.28 | Halevi |  |  | Yehuda Halevi, poet and author of the Kuzari. |
| 1:11.29 | in his poems |  |  |  |
| 1:11.30 | and in the Kuzari |  |  | The Kuzari, Halevi's dialogue set at the court of the Khazar king. |
| 1:11.31 | the Rambam |  |  |  |
| 1:11.32 | and the SeMaG |  |  | SeMaG: an abbreviation for the medieval code Sefer Mitzvot Gadol. |
| 1:11.33 | and the author of |  |  |  |
| 1:11.34 | the Ikkarim |  |  | Ikkarim: 'Principles', the book of the principles of faith. |
| 1:11.35 | of blessed memory |  |  |  |
| 1:11.36 | and as is accepted |  |  |  |
| 1:11.37 | by |  |  |  |
| 1:11.38 | every |  |  |  |
| 1:11.39 | man |  |  |  |
| 1:11.40 | Israel |  |  |  |
| 1:11.41 | and in it |  |  |  |
| 1:11.42 | a warning |  |  |  |
| 1:11.43 | severe |  |  |  |
| 1:11.44 | that not |  |  |  |
| 1:11.45 | to worship |  |  |  |
| 1:11.46 | any |  |  |  |
| 1:11.47 | created thing |  |  |  |
| 1:11.48 | in accordance with |  |  |  |
| 1:11.49 | the tradition of |  |  | The Hebrew is the word for Kabbalah. The author claims it here, for the chain of tradition from the prophets and sages, not for the Kabbalists. |
| 1:11.50 | all |  |  |  |
| 1:11.51 | our prophets |  |  |  |
| 1:11.52 | and all |  |  |  |
| 1:11.53 | our sages |  |  |  |
| 1:11.54 | the Tannaim |  |  |  |
| 1:11.55 | the Amoraim |  |  |  |
| 1:11.56 | and the Geonim |  |  |  |
| 1:11.57 | and the decisors |  |  |  |
| 1:11.58 | the early |  |  |  |
| 1:11.59 | from the mouth of |  |  | From the mouth of: the chain Moses, Aaron and the elders. |
| 1:11.60 | Moses |  |  |  |
| 1:11.61 | and Aaron |  |  |  |
| 1:11.62 | and the elders |  |  |  |
| 1:11.63 | not |  |  |  |
| 1:11.64 | to deviate |  |  |  |
| 1:11.65 | from them |  |  |  |
| 1:11.66 | and not |  |  |  |
| 1:11.67 | to add |  |  |  |
| 1:11.68 | to them |  |  |  |

Phrases:
- 1:11.1 1:11.2 1:11.3 | This book, The Wars of the LORD, | the-book the-wars-of the-LORD
- 1:11.4 1:11.5 1:11.6 1:11.7 1:11.8 | makes known the way of the truths of His unity, may He be blessed, | that-makes-known the-way-of the-truths-of His-unity may-He-be-blessed
- 1:11.9 1:11.10 1:11.11 1:11.12 | according to the Torah and its interpreters, our Rabbis of blessed memory, | according-to the-Torah and-its-interpreters our-Rabbis-of-blessed-memory
- 1:11.13 1:11.14 1:11.15 1:11.16 | in the Mishnah, in the Bavli and the Yerushalmi, | in-the-Mishnah and-in-the-Talmud Bavli and-Yerushalmi
- 1:11.17 1:11.18 | and the Midrashim of our Rabbis of blessed memory, | and-the-Midrashim-of our-Rabbis-of-blessed-memory
- 1:11.19 1:11.20 1:11.21 | and in accordance with the view of the Geonim, Rabbi Saadia Gaon, | and-in-accordance-with the-Geonim Rabbi-Saadia-Gaon
- 1:11.22 1:11.23 1:11.24 1:11.25 | and our Rabbi Bahya in Chovot ha-Levavot, | and-our-Rabbi Bahya in-Chovot-ha-Levavot
- 1:11.26 1:11.27 1:11.28 1:11.29 | and Rabbi Yehuda Halevi in his poems, | and-Rabbi Yehuda Halevi in-his-poems
- 1:11.30 1:11.31 | and in the Kuzari, the Rambam, | and-in-the-Kuzari the-Rambam
- 1:11.32 1:11.33 1:11.34 1:11.35 | the SeMaG and the author of the Ikkarim, of blessed memory, | and-the-SeMaG and-the-author-of the-Ikkarim of-blessed-memory
- 1:11.36 1:11.37 1:11.38 1:11.39 1:11.40 | and as is accepted by every Jew. | and-as-is-accepted by every man-of Israel
- 1:11.41 1:11.42 1:11.43 1:11.44 1:11.45 1:11.46 1:11.47 | In it is a severe warning not to worship any created thing, | and-in-it a-warning severe that-not to-worship any created-thing
- 1:11.48 1:11.49 1:11.50 1:11.51 | in accordance with the tradition of all our prophets | in-accordance-with the-tradition-of all our-prophets
- 1:11.52 1:11.53 1:11.54 1:11.55 1:11.56 1:11.57 1:11.58 | and all our sages, the Tannaim, the Amoraim, the Geonim and the early decisors, | and-all our-sages the-Tannaim and-the-Amoraim and-the-Geonim and-the-decisors the-early
- 1:11.59 1:11.60 1:11.61 1:11.62 | who received it from the mouth of Moses, Aaron and the elders, | from-the-mouth-of Moses and-Aaron and-the-elders
- 1:11.63 1:11.64 1:11.65 1:11.66 1:11.67 1:11.68 | not to deviate from them and not to add to them. | not to-deviate from-them and-not to-add to-them

Grammar notes:
- The sentence has the subject ספר with the participle המודיע and the object דרך אמתות יחודו; the rest is a list of sources.
- The list is joined by ו and by ב (in): 'in the Mishnah', 'in the Bavli', 'in the Kuzari'.
- ובו אזהרה חמורה שלא לעבוד: 'and in it a severe warning not to worship'; the שלא clause gives the content of the warning.
- כפי קבלת … מפי משה ואהרן וזקנים לא לזוז מהם ולא להוסיף עליהם: the chain of tradition ends in two infinitives, 'not to deviate' and 'not to add'.

Translator's note: Title per the decisions (Q1). The Geonim and the Kuzari are given as context.

Commentary rulings:
- apparatus: Num 21:14, the title: ספר מלחמת ה׳ (The Wars of the LORD), as the title page gives it.

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


### Translation under review (1:12)

English: Know, esteemed reader, that you will not grasp the intention of this book, unless you read and study all of it, for its author made nights as days, and labored and scrutinized many books, and drew from seas and rivers proofs strong as a cast mirror, and the matters did not come from his own opinion. Therefore I ask you, esteemed reader, that you not come with the fist of Esau, in mockery, but with the voice of Jacob, and to consider the matters, and whoever is lacking in strength, let him be silent from speech and be still, and may El Shaddai give you mercy, and the merit of the author will protect us, and upon all His people Israel.

Literal: and-know to-you reader esteemed that not you-will-grasp the-intention-of the-book this unless after you-read and-you-study (object) all-of-it for its-author made nights as-days and-labored and-examined in-many books and-drew from-seas and-rivers proofs like-a-mirror molten strong and-not from-his-mind came-out the-matters and-therefore I-ask-you reader esteemed that-not you-come with-his-fist of Esau in-mockery but with-his-voice of Jacob and-to-consider the-matters and-whoever who-is lacking of-strength let-him-be-silent from-speech and-be-still and-El Shaddai may-he-grant to-you mercy and-the-merit-of the-author shall-shield us and-upon-all His-people Israel

| token | gloss | lex | sense | note |
|---|---|---|---|---|
| 1:12.1 | and know |  |  | Imperative 'know', with 'to you'; the author addresses the reader. |
| 1:12.2 | to you |  |  |  |
| 1:12.3 | reader |  |  |  |
| 1:12.4 | esteemed |  |  |  |
| 1:12.5 | that |  |  |  |
| 1:12.6 | not |  |  |  |
| 1:12.7 | you will grasp |  |  | Second-person future with לא: 'you will not grasp'. |
| 1:12.8 | the intention of |  |  |  |
| 1:12.9 | the book |  |  |  |
| 1:12.10 | this |  |  |  |
| 1:12.11 | unless |  |  |  |
| 1:12.12 | after |  |  |  |
| 1:12.13 | you read |  |  |  |
| 1:12.14 | and you study |  |  |  |
| 1:12.15 | (object marker) |  |  |  |
| 1:12.16 | all of it |  |  |  |
| 1:12.17 | for |  |  |  |
| 1:12.18 | its author |  |  |  |
| 1:12.19 | made |  |  |  |
| 1:12.20 | nights |  |  |  |
| 1:12.21 | as days |  |  |  |
| 1:12.22 | and labored |  |  |  |
| 1:12.23 | and examined |  |  |  |
| 1:12.24 | many |  |  |  |
| 1:12.25 | books |  |  |  |
| 1:12.26 | and drew |  |  |  |
| 1:12.27 | from seas |  |  |  |
| 1:12.28 | and rivers |  |  |  |
| 1:12.29 | proofs |  |  |  |
| 1:12.30 | like a mirror |  |  |  |
| 1:12.31 | molten |  |  | מוצק: molten or cast metal; the mirror was of cast metal. |
| 1:12.32 | strong |  |  |  |
| 1:12.33 | and not |  |  |  |
| 1:12.34 | from his mind |  |  |  |
| 1:12.35 | came out |  |  |  |
| 1:12.36 | the matters |  |  |  |
| 1:12.37 | and therefore |  |  |  |
| 1:12.38 | I ask you |  |  |  |
| 1:12.39 | reader |  |  |  |
| 1:12.40 | esteemed |  |  |  |
| 1:12.41 | that not |  |  |  |
| 1:12.42 | you come |  |  |  |
| 1:12.43 | with his fist |  |  | A fist, the image for the hands of Esau (Gen 27:22; see the apparatus). |
| 1:12.44 | of |  |  |  |
| 1:12.45 | Esau |  |  |  |
| 1:12.46 | in mockery |  |  |  |
| 1:12.47 | but |  |  |  |
| 1:12.48 | with his voice |  |  |  |
| 1:12.49 | of |  |  |  |
| 1:12.50 | Jacob |  |  |  |
| 1:12.51 | and to consider |  |  |  |
| 1:12.52 | the matters |  |  |  |
| 1:12.53 | and whoever |  |  |  |
| 1:12.54 | who is |  |  |  |
| 1:12.55 | lacking |  |  |  |
| 1:12.56 | of strength |  |  |  |
| 1:12.57 | let him be silent |  |  |  |
| 1:12.58 | from speech |  |  |  |
| 1:12.59 | and be still |  |  |  |
| 1:12.60 | and El |  |  | El Shaddai, a divine name. |
| 1:12.61 | Shaddai |  |  |  |
| 1:12.62 | may he grant |  |  |  |
| 1:12.63 | to you |  |  |  |
| 1:12.64 | mercy |  |  |  |
| 1:12.65 | and the merit of |  |  |  |
| 1:12.66 | the author |  |  |  |
| 1:12.67 | shall shield |  |  |  |
| 1:12.68 | us |  |  |  |
| 1:12.69 | and upon all |  |  | ע״כ is read as ועל כל, 'and upon all', here. Elsewhere it closes a quotation or means 'thus far'. |
| 1:12.70 | His people |  |  |  |
| 1:12.71 | Israel |  |  |  |

Phrases:
- 1:12.1 1:12.2 1:12.3 1:12.4 | Know, esteemed reader, | and-know to-you reader esteemed
- 1:12.5 1:12.6 1:12.7 1:12.8 1:12.9 1:12.10 | that you will not grasp the intention of this book, | that not you-will-grasp the-intention-of the-book this
- 1:12.11 1:12.12 1:12.13 1:12.14 1:12.15 1:12.16 | unless you read and study all of it, | unless after you-read and-you-study (object) all-of-it
- 1:12.17 1:12.18 1:12.19 1:12.20 1:12.21 | for its author made nights as days, | for its-author made nights as-days
- 1:12.22 1:12.23 1:12.24 1:12.25 | and labored and scrutinized many books, | and-labored and-examined in-many books
- 1:12.26 1:12.27 1:12.28 1:12.29 1:12.30 1:12.31 1:12.32 | and drew from seas and rivers proofs strong as a cast mirror, | and-drew from-seas and-rivers proofs like-a-mirror molten strong
- 1:12.33 1:12.34 1:12.35 1:12.36 | and the matters did not come from his own opinion. | and-not from-his-mind came-out the-matters
- 1:12.37 1:12.38 1:12.39 1:12.40 | Therefore I ask you, esteemed reader, | and-therefore I-ask-you reader esteemed
- 1:12.41 1:12.42 1:12.43 1:12.44 1:12.45 1:12.46 | that you not come with the fist of Esau, in mockery, | that-not you-come with-his-fist of Esau in-mockery
- 1:12.47 1:12.48 1:12.49 1:12.50 1:12.51 1:12.52 | but with the voice of Jacob, and to consider the matters, | but with-his-voice of Jacob and-to-consider the-matters
- 1:12.53 1:12.54 1:12.55 1:12.56 1:12.57 1:12.58 1:12.59 | and whoever is lacking in strength, let him be silent from speech and be still, | and-whoever who-is lacking of-strength let-him-be-silent from-speech and-be-still
- 1:12.60 1:12.61 1:12.62 1:12.63 1:12.64 | and may El Shaddai give you mercy, | and-El Shaddai may-he-grant to-you mercy
- 1:12.65 1:12.66 1:12.67 1:12.68 | and the merit of the author will protect us, | and-the-merit-of the-author shall-shield us
- 1:12.69 1:12.70 1:12.71 | and upon all His people Israel. | and-upon-all His-people Israel

Grammar notes:
- ודע לך: an imperative with 'to you', 'know'; the author addresses the reader in the second person throughout.
- לא תשיג כונת הספר הזה אלא אחרי שתקרא: 'you will not grasp… unless after you read…'.
- כי מחברו עשה לילות כימים: the second כי is causal, 'for its author made nights as days'.
- ולכן אבקשך שלא תבוא באגרונו של עשו בהתולים אלא בקולו של יעקב: the request contrasts the fist of Esau with the voice of Jacob.

Translator's note: The allusion to Gen 27:22 is kept without quotation marks; the author's 'fist' is kept where the verse has 'hands'.

Commentary rulings:
- apparatus: Gen 27:22, an allusion without quotation marks: 'the voice is the voice of Jacob, but the hands are the hands of Esau'. The author writes אגרון (fist) where the verse has ידים (hands).

## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

