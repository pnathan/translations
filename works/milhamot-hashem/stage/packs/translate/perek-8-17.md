# Translation pack: milhamot-hashem perek-8-17 (Milchamot ha-Shem 8:20–8:21)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-8-17.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-8-17`.

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
- **משכילים / השכלה** → "the educated" / "enlightenment" only where the
  Haskalah is meant; else "those of understanding".
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


## Milchamot ha-Shem 8:20

Hebrew (Wikisource): קיג) ולכן אמרו רז"ל על אנשי כנה"ג שהחזירו את העטרה לישנה באמרם האל הגדול הגבור והנורא, ושכל המאורעות שאירעו לישראל מטוב ועד רע, הן הן גבורותיו ונוראותיו של אלהים עליון, והוא אחד אשר אין יחיד כיחודו. ונמשכה ונתפשטה אמונת אחדותו ית' הטהורה בכל ישראל כל משך ימי האלף הרביעי והחמישי ומעט מזער מהאלף השישי, עד שעברו איזה שנים מאלף הששי, וינסה הקב"ה את ישראל במרמת המלך העכו"ם שהוציא את ספר הזהר כאלו נתחבר ע"י רשב"י, ונפתו רבים מחכמינו לדבריו, וההמינו לשקריו לעשות לנו תורה שבעל פה חדשה, ואלהים חדשים, ולדבר תועה על המשנה והתלמוד שהם מטוב ורע מעורבים, ולא מאת השם אלא ממטטרון כדלעיל. והרבה מחכמי ישראל לא עמדו בנסיון ויאמינו בכמה דעות מדומות ונתפשטו בישראל אצל העוסקים בו, וחברו באמונה הזאת ספרים, ונספחו אל המשנה והתלמוד, כאלו הם שוים באמונת האלהות, ולא ביחנו יפה להבין דבריו אשר ידבר במליצות סתומות ומושאלות ובלעני שפה. אמנם רבותינו התלמודיים הדבקים בעסק המשנה והתלמוד ובמדרשי רז"ל וגם המון העם אשר לא ידעו דבר מהקבלה החדשה, הנם עומדים באמונתם השלימה וטהורה מכל סיג הלאה, וכל שכן האנשים המעיינים בשום שכל בדברי הרס"ג, והכוזרי, והרמב"ם, והח"ל, וכיוצא בהם, כולם כאחד תקועים באמונת האחדות האמתית. ועטרת אלהינו ברוך הוא עומדת אצלם בישנה כבימי אנשי כנה"ג. אך דא עקא. כי המה מאמינים בספרי הקבלה החדשה הצבועים למראה עינים. כי קדושים המה ושיש בהם דעות נשגבות אשר לא תכיל דעתם להשיגם. ושהמה שלובי יד עם האמונה הטהורה. ואינם מרגישים במרורות הפתנים אשר בקרבה. העוקרים יסודות תוה"ק. ועמודיה יתפלצון. ורבים מחכמינו שנו בה. ויהיו לאבן מכשול לבאים אחריהם.

| id | word | punctuation after |
|---|---|---|
| 8:20.1 | קיג | ) |
| 8:20.2 | ולכן |  |
| 8:20.3 | אמרו |  |
| 8:20.4 | רז"ל |  |
| 8:20.5 | על |  |
| 8:20.6 | אנשי |  |
| 8:20.7 | כנה"ג |  |
| 8:20.8 | שהחזירו |  |
| 8:20.9 | את |  |
| 8:20.10 | העטרה |  |
| 8:20.11 | לישנה |  |
| 8:20.12 | באמרם |  |
| 8:20.13 | האל |  |
| 8:20.14 | הגדול |  |
| 8:20.15 | הגבור |  |
| 8:20.16 | והנורא | , |
| 8:20.17 | ושכל |  |
| 8:20.18 | המאורעות |  |
| 8:20.19 | שאירעו |  |
| 8:20.20 | לישראל |  |
| 8:20.21 | מטוב |  |
| 8:20.22 | ועד |  |
| 8:20.23 | רע | , |
| 8:20.24 | הן |  |
| 8:20.25 | הן |  |
| 8:20.26 | גבורותיו |  |
| 8:20.27 | ונוראותיו |  |
| 8:20.28 | של |  |
| 8:20.29 | אלהים |  |
| 8:20.30 | עליון | , |
| 8:20.31 | והוא |  |
| 8:20.32 | אחד |  |
| 8:20.33 | אשר |  |
| 8:20.34 | אין |  |
| 8:20.35 | יחיד |  |
| 8:20.36 | כיחודו | . |
| 8:20.37 | ונמשכה |  |
| 8:20.38 | ונתפשטה |  |
| 8:20.39 | אמונת |  |
| 8:20.40 | אחדותו |  |
| 8:20.41 | ית | ' |
| 8:20.42 | הטהורה |  |
| 8:20.43 | בכל |  |
| 8:20.44 | ישראל |  |
| 8:20.45 | כל |  |
| 8:20.46 | משך |  |
| 8:20.47 | ימי |  |
| 8:20.48 | האלף |  |
| 8:20.49 | הרביעי |  |
| 8:20.50 | והחמישי |  |
| 8:20.51 | ומעט |  |
| 8:20.52 | מזער |  |
| 8:20.53 | מהאלף |  |
| 8:20.54 | השישי | , |
| 8:20.55 | עד |  |
| 8:20.56 | שעברו |  |
| 8:20.57 | איזה |  |
| 8:20.58 | שנים |  |
| 8:20.59 | מאלף |  |
| 8:20.60 | הששי | , |
| 8:20.61 | וינסה |  |
| 8:20.62 | הקב"ה |  |
| 8:20.63 | את |  |
| 8:20.64 | ישראל |  |
| 8:20.65 | במרמת |  |
| 8:20.66 | המלך |  |
| 8:20.67 | העכו"ם |  |
| 8:20.68 | שהוציא |  |
| 8:20.69 | את |  |
| 8:20.70 | ספר |  |
| 8:20.71 | הזהר |  |
| 8:20.72 | כאלו |  |
| 8:20.73 | נתחבר |  |
| 8:20.74 | ע"י |  |
| 8:20.75 | רשב"י | , |
| 8:20.76 | ונפתו |  |
| 8:20.77 | רבים |  |
| 8:20.78 | מחכמינו |  |
| 8:20.79 | לדבריו | , |
| 8:20.80 | וההמינו |  |
| 8:20.81 | לשקריו |  |
| 8:20.82 | לעשות |  |
| 8:20.83 | לנו |  |
| 8:20.84 | תורה |  |
| 8:20.85 | שבעל |  |
| 8:20.86 | פה |  |
| 8:20.87 | חדשה | , |
| 8:20.88 | ואלהים |  |
| 8:20.89 | חדשים | , |
| 8:20.90 | ולדבר |  |
| 8:20.91 | תועה |  |
| 8:20.92 | על |  |
| 8:20.93 | המשנה |  |
| 8:20.94 | והתלמוד |  |
| 8:20.95 | שהם |  |
| 8:20.96 | מטוב |  |
| 8:20.97 | ורע |  |
| 8:20.98 | מעורבים | , |
| 8:20.99 | ולא |  |
| 8:20.100 | מאת |  |
| 8:20.101 | השם |  |
| 8:20.102 | אלא |  |
| 8:20.103 | ממטטרון |  |
| 8:20.104 | כדלעיל | . |
| 8:20.105 | והרבה |  |
| 8:20.106 | מחכמי |  |
| 8:20.107 | ישראל |  |
| 8:20.108 | לא |  |
| 8:20.109 | עמדו |  |
| 8:20.110 | בנסיון |  |
| 8:20.111 | ויאמינו |  |
| 8:20.112 | בכמה |  |
| 8:20.113 | דעות |  |
| 8:20.114 | מדומות |  |
| 8:20.115 | ונתפשטו |  |
| 8:20.116 | בישראל |  |
| 8:20.117 | אצל |  |
| 8:20.118 | העוסקים |  |
| 8:20.119 | בו | , |
| 8:20.120 | וחברו |  |
| 8:20.121 | באמונה |  |
| 8:20.122 | הזאת |  |
| 8:20.123 | ספרים | , |
| 8:20.124 | ונספחו |  |
| 8:20.125 | אל |  |
| 8:20.126 | המשנה |  |
| 8:20.127 | והתלמוד | , |
| 8:20.128 | כאלו |  |
| 8:20.129 | הם |  |
| 8:20.130 | שוים |  |
| 8:20.131 | באמונת |  |
| 8:20.132 | האלהות | , |
| 8:20.133 | ולא |  |
| 8:20.134 | ביחנו |  |
| 8:20.135 | יפה |  |
| 8:20.136 | להבין |  |
| 8:20.137 | דבריו |  |
| 8:20.138 | אשר |  |
| 8:20.139 | ידבר |  |
| 8:20.140 | במליצות |  |
| 8:20.141 | סתומות |  |
| 8:20.142 | ומושאלות |  |
| 8:20.143 | ובלעני |  |
| 8:20.144 | שפה | . |
| 8:20.145 | אמנם |  |
| 8:20.146 | רבותינו |  |
| 8:20.147 | התלמודיים |  |
| 8:20.148 | הדבקים |  |
| 8:20.149 | בעסק |  |
| 8:20.150 | המשנה |  |
| 8:20.151 | והתלמוד |  |
| 8:20.152 | ובמדרשי |  |
| 8:20.153 | רז"ל |  |
| 8:20.154 | וגם |  |
| 8:20.155 | המון |  |
| 8:20.156 | העם |  |
| 8:20.157 | אשר |  |
| 8:20.158 | לא |  |
| 8:20.159 | ידעו |  |
| 8:20.160 | דבר |  |
| 8:20.161 | מהקבלה |  |
| 8:20.162 | החדשה | , |
| 8:20.163 | הנם |  |
| 8:20.164 | עומדים |  |
| 8:20.165 | באמונתם |  |
| 8:20.166 | השלימה |  |
| 8:20.167 | וטהורה |  |
| 8:20.168 | מכל |  |
| 8:20.169 | סיג |  |
| 8:20.170 | הלאה | , |
| 8:20.171 | וכל |  |
| 8:20.172 | שכן |  |
| 8:20.173 | האנשים |  |
| 8:20.174 | המעיינים |  |
| 8:20.175 | בשום |  |
| 8:20.176 | שכל |  |
| 8:20.177 | בדברי |  |
| 8:20.178 | הרס"ג | , |
| 8:20.179 | והכוזרי | , |
| 8:20.180 | והרמב"ם | , |
| 8:20.181 | והח"ל | , |
| 8:20.182 | וכיוצא |  |
| 8:20.183 | בהם | , |
| 8:20.184 | כולם |  |
| 8:20.185 | כאחד |  |
| 8:20.186 | תקועים |  |
| 8:20.187 | באמונת |  |
| 8:20.188 | האחדות |  |
| 8:20.189 | האמתית | . |
| 8:20.190 | ועטרת |  |
| 8:20.191 | אלהינו |  |
| 8:20.192 | ברוך |  |
| 8:20.193 | הוא |  |
| 8:20.194 | עומדת |  |
| 8:20.195 | אצלם |  |
| 8:20.196 | בישנה |  |
| 8:20.197 | כבימי |  |
| 8:20.198 | אנשי |  |
| 8:20.199 | כנה"ג | . |
| 8:20.200 | אך |  |
| 8:20.201 | דא |  |
| 8:20.202 | עקא | . |
| 8:20.203 | כי |  |
| 8:20.204 | המה |  |
| 8:20.205 | מאמינים |  |
| 8:20.206 | בספרי |  |
| 8:20.207 | הקבלה |  |
| 8:20.208 | החדשה |  |
| 8:20.209 | הצבועים |  |
| 8:20.210 | למראה |  |
| 8:20.211 | עינים | . |
| 8:20.212 | כי |  |
| 8:20.213 | קדושים |  |
| 8:20.214 | המה |  |
| 8:20.215 | ושיש |  |
| 8:20.216 | בהם |  |
| 8:20.217 | דעות |  |
| 8:20.218 | נשגבות |  |
| 8:20.219 | אשר |  |
| 8:20.220 | לא |  |
| 8:20.221 | תכיל |  |
| 8:20.222 | דעתם |  |
| 8:20.223 | להשיגם | . |
| 8:20.224 | ושהמה |  |
| 8:20.225 | שלובי |  |
| 8:20.226 | יד |  |
| 8:20.227 | עם |  |
| 8:20.228 | האמונה |  |
| 8:20.229 | הטהורה | . |
| 8:20.230 | ואינם |  |
| 8:20.231 | מרגישים |  |
| 8:20.232 | במרורות |  |
| 8:20.233 | הפתנים |  |
| 8:20.234 | אשר |  |
| 8:20.235 | בקרבה | . |
| 8:20.236 | העוקרים |  |
| 8:20.237 | יסודות |  |
| 8:20.238 | תוה"ק | . |
| 8:20.239 | ועמודיה |  |
| 8:20.240 | יתפלצון | . |
| 8:20.241 | ורבים |  |
| 8:20.242 | מחכמינו |  |
| 8:20.243 | שנו |  |
| 8:20.244 | בה | . |
| 8:20.245 | ויהיו |  |
| 8:20.246 | לאבן |  |
| 8:20.247 | מכשול |  |
| 8:20.248 | לבאים |  |
| 8:20.249 | אחריהם | . |


## Milchamot ha-Shem 8:21

Hebrew (Wikisource): קיד) ורבים מהמעיינים יודעים ומעידים ומגידים בצנעה ובהסתר גדול כי זרה היא עד מאוד לאמונת תורתינו באמונת יחוד השי"ת. ולא יחפצו לגלות דעתם, מיראת החסידים והמתחסדים אוהבי ורודפי כבוד וגדולה. העושים כוונות למלכת השמים המדומה אשר שמו ברעיונותם בסתרו של עולם. שלא הועילה אצלם תפלת אנשי כנה"ג ותקנותם וכל יסודותם רק שלא לעשות צלמים ומסכות כסף וזהב ועץ ואבן מעשה ידי חרש. כמו שפתח מחבר הזהר לומר ארור האיש אשר יעשה פסל ומסכה וכו' מעשה ידי חרש. אבל במחשבתם יש הרבה צלמים ובעלים ועשתרות העובדים ומשתחוים להם. ומתפללים אליהם בטהרה. וקורים להם בכל השמות המיוחדים לשי"ת. וכוונות עושים להם לחברם כמכנס חבלים רבים ומושך גמלים. ואומרים ובולא חד. והמשכילים המאמינים באחדותו ית' על פי האמת והצדק. ומחזיקים בתורת האמת בכתב ובע"פ. יראים לנפשם להגיד ולהעלות על דל שפתם את האמת הגמור. מפני תרמית האמת אשר בפי המתחסדים פן יפגעו בהם ויתנפלו עליהם בנידוים וחרמות כתוקף ובחזק יד שלא כרצונו וכרצון יראיו, כנפילים אשר מעולם אנשי השם. ומפחד המשכילים מהם. המה יראים למסור כבודם על קדושת השם הנכבד והנורא. הוא ה' אלהינו אשר בכל הארץ משפטיו. והמה מעלימים מעיניהם מעשה אבותינו שמסרו נפשם וגופם על קדושת השם. אברהם אבינו הולשך לאור כשדים. חנניה מיקאל לועזריה מסרו עצמן והושלכו לכבשן האש. דניאל הושלך לגוב אריות, וכמה צדיקים וחסידים המסופר עליהם בגמרא ומדרשי רז"ל שמסרו עצמן על קדושת השם. יש מהן שניצולו. ויש שנתפסו ונהרגו בעון דוריהם.

| id | word | punctuation after |
|---|---|---|
| 8:21.1 | קיד | ) |
| 8:21.2 | ורבים |  |
| 8:21.3 | מהמעיינים |  |
| 8:21.4 | יודעים |  |
| 8:21.5 | ומעידים |  |
| 8:21.6 | ומגידים |  |
| 8:21.7 | בצנעה |  |
| 8:21.8 | ובהסתר |  |
| 8:21.9 | גדול |  |
| 8:21.10 | כי |  |
| 8:21.11 | זרה |  |
| 8:21.12 | היא |  |
| 8:21.13 | עד |  |
| 8:21.14 | מאוד |  |
| 8:21.15 | לאמונת |  |
| 8:21.16 | תורתינו |  |
| 8:21.17 | באמונת |  |
| 8:21.18 | יחוד |  |
| 8:21.19 | השי"ת | . |
| 8:21.20 | ולא |  |
| 8:21.21 | יחפצו |  |
| 8:21.22 | לגלות |  |
| 8:21.23 | דעתם | , |
| 8:21.24 | מיראת |  |
| 8:21.25 | החסידים |  |
| 8:21.26 | והמתחסדים |  |
| 8:21.27 | אוהבי |  |
| 8:21.28 | ורודפי |  |
| 8:21.29 | כבוד |  |
| 8:21.30 | וגדולה | . |
| 8:21.31 | העושים |  |
| 8:21.32 | כוונות |  |
| 8:21.33 | למלכת |  |
| 8:21.34 | השמים |  |
| 8:21.35 | המדומה |  |
| 8:21.36 | אשר |  |
| 8:21.37 | שמו |  |
| 8:21.38 | ברעיונותם |  |
| 8:21.39 | בסתרו |  |
| 8:21.40 | של |  |
| 8:21.41 | עולם | . |
| 8:21.42 | שלא |  |
| 8:21.43 | הועילה |  |
| 8:21.44 | אצלם |  |
| 8:21.45 | תפלת |  |
| 8:21.46 | אנשי |  |
| 8:21.47 | כנה"ג |  |
| 8:21.48 | ותקנותם |  |
| 8:21.49 | וכל |  |
| 8:21.50 | יסודותם |  |
| 8:21.51 | רק |  |
| 8:21.52 | שלא |  |
| 8:21.53 | לעשות |  |
| 8:21.54 | צלמים |  |
| 8:21.55 | ומסכות |  |
| 8:21.56 | כסף |  |
| 8:21.57 | וזהב |  |
| 8:21.58 | ועץ |  |
| 8:21.59 | ואבן |  |
| 8:21.60 | מעשה |  |
| 8:21.61 | ידי |  |
| 8:21.62 | חרש | . |
| 8:21.63 | כמו |  |
| 8:21.64 | שפתח |  |
| 8:21.65 | מחבר |  |
| 8:21.66 | הזהר |  |
| 8:21.67 | לומר |  |
| 8:21.68 | ארור |  |
| 8:21.69 | האיש |  |
| 8:21.70 | אשר |  |
| 8:21.71 | יעשה |  |
| 8:21.72 | פסל |  |
| 8:21.73 | ומסכה |  |
| 8:21.74 | וכו | ' |
| 8:21.75 | מעשה |  |
| 8:21.76 | ידי |  |
| 8:21.77 | חרש | . |
| 8:21.78 | אבל |  |
| 8:21.79 | במחשבתם |  |
| 8:21.80 | יש |  |
| 8:21.81 | הרבה |  |
| 8:21.82 | צלמים |  |
| 8:21.83 | ובעלים |  |
| 8:21.84 | ועשתרות |  |
| 8:21.85 | העובדים |  |
| 8:21.86 | ומשתחוים |  |
| 8:21.87 | להם | . |
| 8:21.88 | ומתפללים |  |
| 8:21.89 | אליהם |  |
| 8:21.90 | בטהרה | . |
| 8:21.91 | וקורים |  |
| 8:21.92 | להם |  |
| 8:21.93 | בכל |  |
| 8:21.94 | השמות |  |
| 8:21.95 | המיוחדים |  |
| 8:21.96 | לשי"ת | . |
| 8:21.97 | וכוונות |  |
| 8:21.98 | עושים |  |
| 8:21.99 | להם |  |
| 8:21.100 | לחברם |  |
| 8:21.101 | כמכנס |  |
| 8:21.102 | חבלים |  |
| 8:21.103 | רבים |  |
| 8:21.104 | ומושך |  |
| 8:21.105 | גמלים | . |
| 8:21.106 | ואומרים |  |
| 8:21.107 | ובולא |  |
| 8:21.108 | חד | . |
| 8:21.109 | והמשכילים |  |
| 8:21.110 | המאמינים |  |
| 8:21.111 | באחדותו |  |
| 8:21.112 | ית | ' |
| 8:21.113 | על |  |
| 8:21.114 | פי |  |
| 8:21.115 | האמת |  |
| 8:21.116 | והצדק | . |
| 8:21.117 | ומחזיקים |  |
| 8:21.118 | בתורת |  |
| 8:21.119 | האמת |  |
| 8:21.120 | בכתב |  |
| 8:21.121 | ובע"פ | . |
| 8:21.122 | יראים |  |
| 8:21.123 | לנפשם |  |
| 8:21.124 | להגיד |  |
| 8:21.125 | ולהעלות |  |
| 8:21.126 | על |  |
| 8:21.127 | דל |  |
| 8:21.128 | שפתם |  |
| 8:21.129 | את |  |
| 8:21.130 | האמת |  |
| 8:21.131 | הגמור | . |
| 8:21.132 | מפני |  |
| 8:21.133 | תרמית |  |
| 8:21.134 | האמת |  |
| 8:21.135 | אשר |  |
| 8:21.136 | בפי |  |
| 8:21.137 | המתחסדים |  |
| 8:21.138 | פן |  |
| 8:21.139 | יפגעו |  |
| 8:21.140 | בהם |  |
| 8:21.141 | ויתנפלו |  |
| 8:21.142 | עליהם |  |
| 8:21.143 | בנידוים |  |
| 8:21.144 | וחרמות |  |
| 8:21.145 | כתוקף |  |
| 8:21.146 | ובחזק |  |
| 8:21.147 | יד |  |
| 8:21.148 | שלא |  |
| 8:21.149 | כרצונו |  |
| 8:21.150 | וכרצון |  |
| 8:21.151 | יראיו | , |
| 8:21.152 | כנפילים |  |
| 8:21.153 | אשר |  |
| 8:21.154 | מעולם |  |
| 8:21.155 | אנשי |  |
| 8:21.156 | השם | . |
| 8:21.157 | ומפחד |  |
| 8:21.158 | המשכילים |  |
| 8:21.159 | מהם | . |
| 8:21.160 | המה |  |
| 8:21.161 | יראים |  |
| 8:21.162 | למסור |  |
| 8:21.163 | כבודם |  |
| 8:21.164 | על |  |
| 8:21.165 | קדושת |  |
| 8:21.166 | השם |  |
| 8:21.167 | הנכבד |  |
| 8:21.168 | והנורא | . |
| 8:21.169 | הוא |  |
| 8:21.170 | ה | ' |
| 8:21.171 | אלהינו |  |
| 8:21.172 | אשר |  |
| 8:21.173 | בכל |  |
| 8:21.174 | הארץ |  |
| 8:21.175 | משפטיו | . |
| 8:21.176 | והמה |  |
| 8:21.177 | מעלימים |  |
| 8:21.178 | מעיניהם |  |
| 8:21.179 | מעשה |  |
| 8:21.180 | אבותינו |  |
| 8:21.181 | שמסרו |  |
| 8:21.182 | נפשם |  |
| 8:21.183 | וגופם |  |
| 8:21.184 | על |  |
| 8:21.185 | קדושת |  |
| 8:21.186 | השם | . |
| 8:21.187 | אברהם |  |
| 8:21.188 | אבינו |  |
| 8:21.189 | הולשך |  |
| 8:21.190 | לאור |  |
| 8:21.191 | כשדים | . |
| 8:21.192 | חנניה |  |
| 8:21.193 | מיקאל |  |
| 8:21.194 | לועזריה |  |
| 8:21.195 | מסרו |  |
| 8:21.196 | עצמן |  |
| 8:21.197 | והושלכו |  |
| 8:21.198 | לכבשן |  |
| 8:21.199 | האש | . |
| 8:21.200 | דניאל |  |
| 8:21.201 | הושלך |  |
| 8:21.202 | לגוב |  |
| 8:21.203 | אריות | , |
| 8:21.204 | וכמה |  |
| 8:21.205 | צדיקים |  |
| 8:21.206 | וחסידים |  |
| 8:21.207 | המסופר |  |
| 8:21.208 | עליהם |  |
| 8:21.209 | בגמרא |  |
| 8:21.210 | ומדרשי |  |
| 8:21.211 | רז"ל |  |
| 8:21.212 | שמסרו |  |
| 8:21.213 | עצמן |  |
| 8:21.214 | על |  |
| 8:21.215 | קדושת |  |
| 8:21.216 | השם | . |
| 8:21.217 | יש |  |
| 8:21.218 | מהן |  |
| 8:21.219 | שניצולו | . |
| 8:21.220 | ויש |  |
| 8:21.221 | שנתפסו |  |
| 8:21.222 | ונהרגו |  |
| 8:21.223 | בעון |  |
| 8:21.224 | דוריהם | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

