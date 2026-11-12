# Translation pack: milhamot-hashem perek-5-06 (Milchamot ha-Shem 5:11–5:12)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-06.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-06`.

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
- **אצילות / באצילות / בריאה**: fixed, translated, the verbs linked:
  "emanation"; באצילות → "in Emanation"; עולם האצילות → "the World of
  Emanation"; נאצל → "emanated"; המאציל → "the Emanator". The four worlds
  (אבי"ע): "Emanation, Creation, Formation, Making" (verbs: "emanated,
  created, formed, made"). נאצל and נברא stay distinct ("נברא ולא בורא" →
  "created, not creator").
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
- **קדמון / קדמות**: of God or the world → "eternal / eternity" (note: "without
  beginning"); נצחי / נצחיות → "everlasting / everlastingness", distinct.
  קדמונינו → "our forebears" (fixed); הקדמונים → "the ancients"; קדמו →
  "preceded".
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **ובפתרון הכתובים** → "the verses" when the example is from the Torah; "the
  Writings" only for the third part of the Bible.
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
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


## Milchamot ha-Shem 5:11

Hebrew (Wikisource): מד) ואחר שידעת את כל אלה בא וראה איך שגגו רבותינו המקובלים והטילו דופי בידיעת השי"ת בעתידות, גם סלקו ההשגחה מהמלך העליון ותלו ההשגחה בזעיר אנפין לבד, ולא באשר הפרצופין אבא ואמא, ואריך, ועתיק, ואדם קדמון, וכ"ש באין סוף העליון, ואמרו שאין מקפידין ומשגיחין במעשה התחתונים בין טוב לרע ח"ו כטוב כחוטא, והיה כצדיק כרשע, הנשבע כאשר שבועה ירא כל אפיא שוין. ודבר זה מבואר בכתבי האר"י הובא בספר ד' מאות שקל כסף דף צ"א ע"ב ז"ל שאלה ששאל הרב המובהק המקובל האלהי כמהר"ר אברהם מונצוץ מעיר טניפיל להאר"י זצוק"ל יש לשאול דבר אחד קשה והוא כי ראינו בספרי הקבלה והזהר ובספר המדע ובספר הקנה כי בשעה שהאל מתעסק בבריאה, אינו יודע. וכשיהיה באצילות יודע. אם כן הרי נתבאר שבעולם האצילות היא הידיעה והכריחה, א"כ קשה לדברי רז"ל על פסוק ויאמר ה' אל משה אמר אל הכוהנים בני אהרן, מלך ראשון שיעמוד לישראל ידקר בחרב. אמר לו ולי אתה שואל, אמור אל הכוהנים בני אהרן שהרי הרג נוב עיר הכהנים. ונמצא שהכריחה דעתו של שאול, באיזה אופן שיהיה אם בבריאה אם באצילות. וא"כ נמצא שהכריחה על כל פנים.

| id | word | punctuation after |
|---|---|---|
| 5:11.1 | מד | ) |
| 5:11.2 | ואחר |  |
| 5:11.3 | שידעת |  |
| 5:11.4 | את |  |
| 5:11.5 | כל |  |
| 5:11.6 | אלה |  |
| 5:11.7 | בא |  |
| 5:11.8 | וראה |  |
| 5:11.9 | איך |  |
| 5:11.10 | שגגו |  |
| 5:11.11 | רבותינו |  |
| 5:11.12 | המקובלים |  |
| 5:11.13 | והטילו |  |
| 5:11.14 | דופי |  |
| 5:11.15 | בידיעת |  |
| 5:11.16 | השי"ת |  |
| 5:11.17 | בעתידות | , |
| 5:11.18 | גם |  |
| 5:11.19 | סלקו |  |
| 5:11.20 | ההשגחה |  |
| 5:11.21 | מהמלך |  |
| 5:11.22 | העליון |  |
| 5:11.23 | ותלו |  |
| 5:11.24 | ההשגחה |  |
| 5:11.25 | בזעיר |  |
| 5:11.26 | אנפין |  |
| 5:11.27 | לבד | , |
| 5:11.28 | ולא |  |
| 5:11.29 | באשר |  |
| 5:11.30 | הפרצופין |  |
| 5:11.31 | אבא |  |
| 5:11.32 | ואמא | , |
| 5:11.33 | ואריך | , |
| 5:11.34 | ועתיק | , |
| 5:11.35 | ואדם |  |
| 5:11.36 | קדמון | , |
| 5:11.37 | וכ"ש |  |
| 5:11.38 | באין |  |
| 5:11.39 | סוף |  |
| 5:11.40 | העליון | , |
| 5:11.41 | ואמרו |  |
| 5:11.42 | שאין |  |
| 5:11.43 | מקפידין |  |
| 5:11.44 | ומשגיחין |  |
| 5:11.45 | במעשה |  |
| 5:11.46 | התחתונים |  |
| 5:11.47 | בין |  |
| 5:11.48 | טוב |  |
| 5:11.49 | לרע |  |
| 5:11.50 | ח"ו |  |
| 5:11.51 | כטוב |  |
| 5:11.52 | כחוטא | , |
| 5:11.53 | והיה |  |
| 5:11.54 | כצדיק |  |
| 5:11.55 | כרשע | , |
| 5:11.56 | הנשבע |  |
| 5:11.57 | כאשר |  |
| 5:11.58 | שבועה |  |
| 5:11.59 | ירא |  |
| 5:11.60 | כל |  |
| 5:11.61 | אפיא |  |
| 5:11.62 | שוין | . |
| 5:11.63 | ודבר |  |
| 5:11.64 | זה |  |
| 5:11.65 | מבואר |  |
| 5:11.66 | בכתבי |  |
| 5:11.67 | האר"י |  |
| 5:11.68 | הובא |  |
| 5:11.69 | בספר |  |
| 5:11.70 | ד | ' |
| 5:11.71 | מאות |  |
| 5:11.72 | שקל |  |
| 5:11.73 | כסף |  |
| 5:11.74 | דף |  |
| 5:11.75 | צ"א |  |
| 5:11.76 | ע"ב |  |
| 5:11.77 | ז"ל |  |
| 5:11.78 | שאלה |  |
| 5:11.79 | ששאל |  |
| 5:11.80 | הרב |  |
| 5:11.81 | המובהק |  |
| 5:11.82 | המקובל |  |
| 5:11.83 | האלהי |  |
| 5:11.84 | כמהר"ר |  |
| 5:11.85 | אברהם |  |
| 5:11.86 | מונצוץ |  |
| 5:11.87 | מעיר |  |
| 5:11.88 | טניפיל |  |
| 5:11.89 | להאר"י |  |
| 5:11.90 | זצוק"ל |  |
| 5:11.91 | יש |  |
| 5:11.92 | לשאול |  |
| 5:11.93 | דבר |  |
| 5:11.94 | אחד |  |
| 5:11.95 | קשה |  |
| 5:11.96 | והוא |  |
| 5:11.97 | כי |  |
| 5:11.98 | ראינו |  |
| 5:11.99 | בספרי |  |
| 5:11.100 | הקבלה |  |
| 5:11.101 | והזהר |  |
| 5:11.102 | ובספר |  |
| 5:11.103 | המדע |  |
| 5:11.104 | ובספר |  |
| 5:11.105 | הקנה |  |
| 5:11.106 | כי |  |
| 5:11.107 | בשעה |  |
| 5:11.108 | שהאל |  |
| 5:11.109 | מתעסק |  |
| 5:11.110 | בבריאה | , |
| 5:11.111 | אינו |  |
| 5:11.112 | יודע | . |
| 5:11.113 | וכשיהיה |  |
| 5:11.114 | באצילות |  |
| 5:11.115 | יודע | . |
| 5:11.116 | אם |  |
| 5:11.117 | כן |  |
| 5:11.118 | הרי |  |
| 5:11.119 | נתבאר |  |
| 5:11.120 | שבעולם |  |
| 5:11.121 | האצילות |  |
| 5:11.122 | היא |  |
| 5:11.123 | הידיעה |  |
| 5:11.124 | והכריחה | , |
| 5:11.125 | א"כ |  |
| 5:11.126 | קשה |  |
| 5:11.127 | לדברי |  |
| 5:11.128 | רז"ל |  |
| 5:11.129 | על |  |
| 5:11.130 | פסוק |  |
| 5:11.131 | ויאמר |  |
| 5:11.132 | ה | ' |
| 5:11.133 | אל |  |
| 5:11.134 | משה |  |
| 5:11.135 | אמר |  |
| 5:11.136 | אל |  |
| 5:11.137 | הכוהנים |  |
| 5:11.138 | בני |  |
| 5:11.139 | אהרן | , |
| 5:11.140 | מלך |  |
| 5:11.141 | ראשון |  |
| 5:11.142 | שיעמוד |  |
| 5:11.143 | לישראל |  |
| 5:11.144 | ידקר |  |
| 5:11.145 | בחרב | . |
| 5:11.146 | אמר |  |
| 5:11.147 | לו |  |
| 5:11.148 | ולי |  |
| 5:11.149 | אתה |  |
| 5:11.150 | שואל | , |
| 5:11.151 | אמור |  |
| 5:11.152 | אל |  |
| 5:11.153 | הכוהנים |  |
| 5:11.154 | בני |  |
| 5:11.155 | אהרן |  |
| 5:11.156 | שהרי |  |
| 5:11.157 | הרג |  |
| 5:11.158 | נוב |  |
| 5:11.159 | עיר |  |
| 5:11.160 | הכהנים | . |
| 5:11.161 | ונמצא |  |
| 5:11.162 | שהכריחה |  |
| 5:11.163 | דעתו |  |
| 5:11.164 | של |  |
| 5:11.165 | שאול | , |
| 5:11.166 | באיזה |  |
| 5:11.167 | אופן |  |
| 5:11.168 | שיהיה |  |
| 5:11.169 | אם |  |
| 5:11.170 | בבריאה |  |
| 5:11.171 | אם |  |
| 5:11.172 | באצילות | . |
| 5:11.173 | וא"כ |  |
| 5:11.174 | נמצא |  |
| 5:11.175 | שהכריחה |  |
| 5:11.176 | על |  |
| 5:11.177 | כל |  |
| 5:11.178 | פנים | . |


## Milchamot ha-Shem 5:12

Hebrew (Wikisource): והתשובה שהשיב מורי זלה"ה, הוא אמת שיש באצילות ידיעה, אבל יש כח באדם לבחור זולתה כמו שאמר הכתוב ראה נתתי לפניך היום את החיים ואת הטוב ואת המות ואת הרע ובחרת בחיים למען תחיה אתה וזרעיך. הרי מכאן שהכריחה ממעשה שאול, ומן הכתוב מוכיח שהיא בחירה שנ' ובחרת בחיים וגם לרז"ל שאמרו אין רע יורד מן השמים, כי למעלה באצילות הכל פשוט, ואין אותה ידיעה יורד למטה להכריח על האדם כי אין שכר ועונש ואין בחירה ורצון למעלה, וזה סוד הלא אח עשו ליעקב נאם ה' ואהב את יעקב, וכי השי"ת אומר שכמו יעקב כן עשו. כי כבר העידה תורה כי יעקב איש תם יושב אהלים, ועשו עובד ע"ז, והשי"ת בוחר בצדיק, וא"כ מהו זה שאמר השי"ת ואהב את יעקב, זה צדיק וזה רשע, אלא הכוונה היא, כי למעלה באצילות כמו יעקב כן עשו, כי שם לא יש שכר ועונש, וזהו סוד הלא אח עשו ליעקב נאם ה', שהוא עתיקא קדישא, כמו שפירש באדרא, בי נשבעתי נאם ה' עתיקא קדישא. על כן אמר יש לכם לשמוע בקלי ולקבל מלכותי כי בחרתי ביעקב ואף שהוא כמו עשו וגם אם נאמר שיש ידיעה מכרחת, אין צרך לתורה ולא למצות, כי כבר הוכרח האדם במעשיו, וגם מאמר, ר' חנניה בן עקשיא אומר רצה הקב"ה לזכות את ישראל לפיכך הרבה להם תורה ומצות. אם נאמר שיש ידיעה מכרחת, אין צורך בהרבות תורה ומצות. אלא הכוונה היא למעלה באצילות יש ידיעה ואין ידיעת הרע של אדם יורד מלמעלה כי נשאר שם וביד האדם שהוא בעל בחירה לבחור בזולתה. על כן צותה תורה לשמור מצותיה וגם אם האדם דבק בתורה מושך עליו הטוב. כי אין רע יורד מעצמו מלמעלה אלא ימשוך האדם עליו הרע אז יורד. ועם זה יובן מאמר אחר פר' ויקרא או ראה או ידע כי שם אומר כי הנפש קודם שתרד משביע אותה הקב"ה תהי צדיק ואל תהי רשע. כי אם נאמר שיש ידיעה נמצא משביעה לשקר. שהרי הוא יודע בו שיחטא ומשביע. וגם מאמר אחר לרז"ל משביעין אותו תהי צדיק ואל תהי רשע. אם נאמר שידע נמצא משביעו לשקר. ואין הקב"ה רצה שישבעו בשמו לשקר. אלא אלו המאמרים וגם הרבה כתובים מוכיחים שבשעת הבריאה אין לו ידיעה. וממה שביארנו יש ראיה שאפי' ידיעת האצילות היא ידיעה פשוטה, הויא תלויה ברצון האדם שכבר נאמר לו כי הוא כמו קונו שהוא בעל בחירה ורצון כמו שאמר הכתוב כי יעקב בחר לו יה ישראל לסגולתו. הרי הוא מובן שיש לו בחירה ורצון. גם מה שנ' ורציתי אתכם וכו' הרי מן הכתובים שיש לו בחירה ורצון. וגם מדמה ישראל אליו שאמר והייתם קדושים כי קדוש אני כמו שאני קדוש כן אתם תהיו קדושים. אחי ורעי בני כדמותו כצלמו. כי אם יאמר שיש לו ידיעה ואינה מכרחת כמו שכתב הרמב"ם נמצא שאול שהכריחה אותו ידיעתו עליו יש לו תלונה על זה, אמנם אם נאמר כמו שאמרנו יבוא הכל באמת ובאמונה עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 5:12.1 | והתשובה |  |
| 5:12.2 | שהשיב |  |
| 5:12.3 | מורי |  |
| 5:12.4 | זלה"ה | , |
| 5:12.5 | הוא |  |
| 5:12.6 | אמת |  |
| 5:12.7 | שיש |  |
| 5:12.8 | באצילות |  |
| 5:12.9 | ידיעה | , |
| 5:12.10 | אבל |  |
| 5:12.11 | יש |  |
| 5:12.12 | כח |  |
| 5:12.13 | באדם |  |
| 5:12.14 | לבחור |  |
| 5:12.15 | זולתה |  |
| 5:12.16 | כמו |  |
| 5:12.17 | שאמר |  |
| 5:12.18 | הכתוב |  |
| 5:12.19 | ראה |  |
| 5:12.20 | נתתי |  |
| 5:12.21 | לפניך |  |
| 5:12.22 | היום |  |
| 5:12.23 | את |  |
| 5:12.24 | החיים |  |
| 5:12.25 | ואת |  |
| 5:12.26 | הטוב |  |
| 5:12.27 | ואת |  |
| 5:12.28 | המות |  |
| 5:12.29 | ואת |  |
| 5:12.30 | הרע |  |
| 5:12.31 | ובחרת |  |
| 5:12.32 | בחיים |  |
| 5:12.33 | למען |  |
| 5:12.34 | תחיה |  |
| 5:12.35 | אתה |  |
| 5:12.36 | וזרעיך | . |
| 5:12.37 | הרי |  |
| 5:12.38 | מכאן |  |
| 5:12.39 | שהכריחה |  |
| 5:12.40 | ממעשה |  |
| 5:12.41 | שאול | , |
| 5:12.42 | ומן |  |
| 5:12.43 | הכתוב |  |
| 5:12.44 | מוכיח |  |
| 5:12.45 | שהיא |  |
| 5:12.46 | בחירה |  |
| 5:12.47 | שנ | ' |
| 5:12.48 | ובחרת |  |
| 5:12.49 | בחיים |  |
| 5:12.50 | וגם |  |
| 5:12.51 | לרז"ל |  |
| 5:12.52 | שאמרו |  |
| 5:12.53 | אין |  |
| 5:12.54 | רע |  |
| 5:12.55 | יורד |  |
| 5:12.56 | מן |  |
| 5:12.57 | השמים | , |
| 5:12.58 | כי |  |
| 5:12.59 | למעלה |  |
| 5:12.60 | באצילות |  |
| 5:12.61 | הכל |  |
| 5:12.62 | פשוט | , |
| 5:12.63 | ואין |  |
| 5:12.64 | אותה |  |
| 5:12.65 | ידיעה |  |
| 5:12.66 | יורד |  |
| 5:12.67 | למטה |  |
| 5:12.68 | להכריח |  |
| 5:12.69 | על |  |
| 5:12.70 | האדם |  |
| 5:12.71 | כי |  |
| 5:12.72 | אין |  |
| 5:12.73 | שכר |  |
| 5:12.74 | ועונש |  |
| 5:12.75 | ואין |  |
| 5:12.76 | בחירה |  |
| 5:12.77 | ורצון |  |
| 5:12.78 | למעלה | , |
| 5:12.79 | וזה |  |
| 5:12.80 | סוד |  |
| 5:12.81 | הלא |  |
| 5:12.82 | אח |  |
| 5:12.83 | עשו |  |
| 5:12.84 | ליעקב |  |
| 5:12.85 | נאם |  |
| 5:12.86 | ה | ' |
| 5:12.87 | ואהב |  |
| 5:12.88 | את |  |
| 5:12.89 | יעקב | , |
| 5:12.90 | וכי |  |
| 5:12.91 | השי"ת |  |
| 5:12.92 | אומר |  |
| 5:12.93 | שכמו |  |
| 5:12.94 | יעקב |  |
| 5:12.95 | כן |  |
| 5:12.96 | עשו | . |
| 5:12.97 | כי |  |
| 5:12.98 | כבר |  |
| 5:12.99 | העידה |  |
| 5:12.100 | תורה |  |
| 5:12.101 | כי |  |
| 5:12.102 | יעקב |  |
| 5:12.103 | איש |  |
| 5:12.104 | תם |  |
| 5:12.105 | יושב |  |
| 5:12.106 | אהלים | , |
| 5:12.107 | ועשו |  |
| 5:12.108 | עובד |  |
| 5:12.109 | ע"ז | , |
| 5:12.110 | והשי"ת |  |
| 5:12.111 | בוחר |  |
| 5:12.112 | בצדיק | , |
| 5:12.113 | וא"כ |  |
| 5:12.114 | מהו |  |
| 5:12.115 | זה |  |
| 5:12.116 | שאמר |  |
| 5:12.117 | השי"ת |  |
| 5:12.118 | ואהב |  |
| 5:12.119 | את |  |
| 5:12.120 | יעקב | , |
| 5:12.121 | זה |  |
| 5:12.122 | צדיק |  |
| 5:12.123 | וזה |  |
| 5:12.124 | רשע | , |
| 5:12.125 | אלא |  |
| 5:12.126 | הכוונה |  |
| 5:12.127 | היא | , |
| 5:12.128 | כי |  |
| 5:12.129 | למעלה |  |
| 5:12.130 | באצילות |  |
| 5:12.131 | כמו |  |
| 5:12.132 | יעקב |  |
| 5:12.133 | כן |  |
| 5:12.134 | עשו | , |
| 5:12.135 | כי |  |
| 5:12.136 | שם |  |
| 5:12.137 | לא |  |
| 5:12.138 | יש |  |
| 5:12.139 | שכר |  |
| 5:12.140 | ועונש | , |
| 5:12.141 | וזהו |  |
| 5:12.142 | סוד |  |
| 5:12.143 | הלא |  |
| 5:12.144 | אח |  |
| 5:12.145 | עשו |  |
| 5:12.146 | ליעקב |  |
| 5:12.147 | נאם |  |
| 5:12.148 | ה | ', |
| 5:12.149 | שהוא |  |
| 5:12.150 | עתיקא |  |
| 5:12.151 | קדישא | , |
| 5:12.152 | כמו |  |
| 5:12.153 | שפירש |  |
| 5:12.154 | באדרא | , |
| 5:12.155 | בי |  |
| 5:12.156 | נשבעתי |  |
| 5:12.157 | נאם |  |
| 5:12.158 | ה | ' |
| 5:12.159 | עתיקא |  |
| 5:12.160 | קדישא | . |
| 5:12.161 | על |  |
| 5:12.162 | כן |  |
| 5:12.163 | אמר |  |
| 5:12.164 | יש |  |
| 5:12.165 | לכם |  |
| 5:12.166 | לשמוע |  |
| 5:12.167 | בקלי |  |
| 5:12.168 | ולקבל |  |
| 5:12.169 | מלכותי |  |
| 5:12.170 | כי |  |
| 5:12.171 | בחרתי |  |
| 5:12.172 | ביעקב |  |
| 5:12.173 | ואף |  |
| 5:12.174 | שהוא |  |
| 5:12.175 | כמו |  |
| 5:12.176 | עשו |  |
| 5:12.177 | וגם |  |
| 5:12.178 | אם |  |
| 5:12.179 | נאמר |  |
| 5:12.180 | שיש |  |
| 5:12.181 | ידיעה |  |
| 5:12.182 | מכרחת | , |
| 5:12.183 | אין |  |
| 5:12.184 | צרך |  |
| 5:12.185 | לתורה |  |
| 5:12.186 | ולא |  |
| 5:12.187 | למצות | , |
| 5:12.188 | כי |  |
| 5:12.189 | כבר |  |
| 5:12.190 | הוכרח |  |
| 5:12.191 | האדם |  |
| 5:12.192 | במעשיו | , |
| 5:12.193 | וגם |  |
| 5:12.194 | מאמר | , |
| 5:12.195 | ר | ' |
| 5:12.196 | חנניה |  |
| 5:12.197 | בן |  |
| 5:12.198 | עקשיא |  |
| 5:12.199 | אומר |  |
| 5:12.200 | רצה |  |
| 5:12.201 | הקב"ה |  |
| 5:12.202 | לזכות |  |
| 5:12.203 | את |  |
| 5:12.204 | ישראל |  |
| 5:12.205 | לפיכך |  |
| 5:12.206 | הרבה |  |
| 5:12.207 | להם |  |
| 5:12.208 | תורה |  |
| 5:12.209 | ומצות | . |
| 5:12.210 | אם |  |
| 5:12.211 | נאמר |  |
| 5:12.212 | שיש |  |
| 5:12.213 | ידיעה |  |
| 5:12.214 | מכרחת | , |
| 5:12.215 | אין |  |
| 5:12.216 | צורך |  |
| 5:12.217 | בהרבות |  |
| 5:12.218 | תורה |  |
| 5:12.219 | ומצות | . |
| 5:12.220 | אלא |  |
| 5:12.221 | הכוונה |  |
| 5:12.222 | היא |  |
| 5:12.223 | למעלה |  |
| 5:12.224 | באצילות |  |
| 5:12.225 | יש |  |
| 5:12.226 | ידיעה |  |
| 5:12.227 | ואין |  |
| 5:12.228 | ידיעת |  |
| 5:12.229 | הרע |  |
| 5:12.230 | של |  |
| 5:12.231 | אדם |  |
| 5:12.232 | יורד |  |
| 5:12.233 | מלמעלה |  |
| 5:12.234 | כי |  |
| 5:12.235 | נשאר |  |
| 5:12.236 | שם |  |
| 5:12.237 | וביד |  |
| 5:12.238 | האדם |  |
| 5:12.239 | שהוא |  |
| 5:12.240 | בעל |  |
| 5:12.241 | בחירה |  |
| 5:12.242 | לבחור |  |
| 5:12.243 | בזולתה | . |
| 5:12.244 | על |  |
| 5:12.245 | כן |  |
| 5:12.246 | צותה |  |
| 5:12.247 | תורה |  |
| 5:12.248 | לשמור |  |
| 5:12.249 | מצותיה |  |
| 5:12.250 | וגם |  |
| 5:12.251 | אם |  |
| 5:12.252 | האדם |  |
| 5:12.253 | דבק |  |
| 5:12.254 | בתורה |  |
| 5:12.255 | מושך |  |
| 5:12.256 | עליו |  |
| 5:12.257 | הטוב | . |
| 5:12.258 | כי |  |
| 5:12.259 | אין |  |
| 5:12.260 | רע |  |
| 5:12.261 | יורד |  |
| 5:12.262 | מעצמו |  |
| 5:12.263 | מלמעלה |  |
| 5:12.264 | אלא |  |
| 5:12.265 | ימשוך |  |
| 5:12.266 | האדם |  |
| 5:12.267 | עליו |  |
| 5:12.268 | הרע |  |
| 5:12.269 | אז |  |
| 5:12.270 | יורד | . |
| 5:12.271 | ועם |  |
| 5:12.272 | זה |  |
| 5:12.273 | יובן |  |
| 5:12.274 | מאמר |  |
| 5:12.275 | אחר |  |
| 5:12.276 | פר | ' |
| 5:12.277 | ויקרא |  |
| 5:12.278 | או |  |
| 5:12.279 | ראה |  |
| 5:12.280 | או |  |
| 5:12.281 | ידע |  |
| 5:12.282 | כי |  |
| 5:12.283 | שם |  |
| 5:12.284 | אומר |  |
| 5:12.285 | כי |  |
| 5:12.286 | הנפש |  |
| 5:12.287 | קודם |  |
| 5:12.288 | שתרד |  |
| 5:12.289 | משביע |  |
| 5:12.290 | אותה |  |
| 5:12.291 | הקב"ה |  |
| 5:12.292 | תהי |  |
| 5:12.293 | צדיק |  |
| 5:12.294 | ואל |  |
| 5:12.295 | תהי |  |
| 5:12.296 | רשע | . |
| 5:12.297 | כי |  |
| 5:12.298 | אם |  |
| 5:12.299 | נאמר |  |
| 5:12.300 | שיש |  |
| 5:12.301 | ידיעה |  |
| 5:12.302 | נמצא |  |
| 5:12.303 | משביעה |  |
| 5:12.304 | לשקר | . |
| 5:12.305 | שהרי |  |
| 5:12.306 | הוא |  |
| 5:12.307 | יודע |  |
| 5:12.308 | בו |  |
| 5:12.309 | שיחטא |  |
| 5:12.310 | ומשביע | . |
| 5:12.311 | וגם |  |
| 5:12.312 | מאמר |  |
| 5:12.313 | אחר |  |
| 5:12.314 | לרז"ל |  |
| 5:12.315 | משביעין |  |
| 5:12.316 | אותו |  |
| 5:12.317 | תהי |  |
| 5:12.318 | צדיק |  |
| 5:12.319 | ואל |  |
| 5:12.320 | תהי |  |
| 5:12.321 | רשע | . |
| 5:12.322 | אם |  |
| 5:12.323 | נאמר |  |
| 5:12.324 | שידע |  |
| 5:12.325 | נמצא |  |
| 5:12.326 | משביעו |  |
| 5:12.327 | לשקר | . |
| 5:12.328 | ואין |  |
| 5:12.329 | הקב"ה |  |
| 5:12.330 | רצה |  |
| 5:12.331 | שישבעו |  |
| 5:12.332 | בשמו |  |
| 5:12.333 | לשקר | . |
| 5:12.334 | אלא |  |
| 5:12.335 | אלו |  |
| 5:12.336 | המאמרים |  |
| 5:12.337 | וגם |  |
| 5:12.338 | הרבה |  |
| 5:12.339 | כתובים |  |
| 5:12.340 | מוכיחים |  |
| 5:12.341 | שבשעת |  |
| 5:12.342 | הבריאה |  |
| 5:12.343 | אין |  |
| 5:12.344 | לו |  |
| 5:12.345 | ידיעה | . |
| 5:12.346 | וממה |  |
| 5:12.347 | שביארנו |  |
| 5:12.348 | יש |  |
| 5:12.349 | ראיה |  |
| 5:12.350 | שאפי | ' |
| 5:12.351 | ידיעת |  |
| 5:12.352 | האצילות |  |
| 5:12.353 | היא |  |
| 5:12.354 | ידיעה |  |
| 5:12.355 | פשוטה | , |
| 5:12.356 | הויא |  |
| 5:12.357 | תלויה |  |
| 5:12.358 | ברצון |  |
| 5:12.359 | האדם |  |
| 5:12.360 | שכבר |  |
| 5:12.361 | נאמר |  |
| 5:12.362 | לו |  |
| 5:12.363 | כי |  |
| 5:12.364 | הוא |  |
| 5:12.365 | כמו |  |
| 5:12.366 | קונו |  |
| 5:12.367 | שהוא |  |
| 5:12.368 | בעל |  |
| 5:12.369 | בחירה |  |
| 5:12.370 | ורצון |  |
| 5:12.371 | כמו |  |
| 5:12.372 | שאמר |  |
| 5:12.373 | הכתוב |  |
| 5:12.374 | כי |  |
| 5:12.375 | יעקב |  |
| 5:12.376 | בחר |  |
| 5:12.377 | לו |  |
| 5:12.378 | יה |  |
| 5:12.379 | ישראל |  |
| 5:12.380 | לסגולתו | . |
| 5:12.381 | הרי |  |
| 5:12.382 | הוא |  |
| 5:12.383 | מובן |  |
| 5:12.384 | שיש |  |
| 5:12.385 | לו |  |
| 5:12.386 | בחירה |  |
| 5:12.387 | ורצון | . |
| 5:12.388 | גם |  |
| 5:12.389 | מה |  |
| 5:12.390 | שנ | ' |
| 5:12.391 | ורציתי |  |
| 5:12.392 | אתכם |  |
| 5:12.393 | וכו | ' |
| 5:12.394 | הרי |  |
| 5:12.395 | מן |  |
| 5:12.396 | הכתובים |  |
| 5:12.397 | שיש |  |
| 5:12.398 | לו |  |
| 5:12.399 | בחירה |  |
| 5:12.400 | ורצון | . |
| 5:12.401 | וגם |  |
| 5:12.402 | מדמה |  |
| 5:12.403 | ישראל |  |
| 5:12.404 | אליו |  |
| 5:12.405 | שאמר |  |
| 5:12.406 | והייתם |  |
| 5:12.407 | קדושים |  |
| 5:12.408 | כי |  |
| 5:12.409 | קדוש |  |
| 5:12.410 | אני |  |
| 5:12.411 | כמו |  |
| 5:12.412 | שאני |  |
| 5:12.413 | קדוש |  |
| 5:12.414 | כן |  |
| 5:12.415 | אתם |  |
| 5:12.416 | תהיו |  |
| 5:12.417 | קדושים | . |
| 5:12.418 | אחי |  |
| 5:12.419 | ורעי |  |
| 5:12.420 | בני |  |
| 5:12.421 | כדמותו |  |
| 5:12.422 | כצלמו | . |
| 5:12.423 | כי |  |
| 5:12.424 | אם |  |
| 5:12.425 | יאמר |  |
| 5:12.426 | שיש |  |
| 5:12.427 | לו |  |
| 5:12.428 | ידיעה |  |
| 5:12.429 | ואינה |  |
| 5:12.430 | מכרחת |  |
| 5:12.431 | כמו |  |
| 5:12.432 | שכתב |  |
| 5:12.433 | הרמב"ם |  |
| 5:12.434 | נמצא |  |
| 5:12.435 | שאול |  |
| 5:12.436 | שהכריחה |  |
| 5:12.437 | אותו |  |
| 5:12.438 | ידיעתו |  |
| 5:12.439 | עליו |  |
| 5:12.440 | יש |  |
| 5:12.441 | לו |  |
| 5:12.442 | תלונה |  |
| 5:12.443 | על |  |
| 5:12.444 | זה | , |
| 5:12.445 | אמנם |  |
| 5:12.446 | אם |  |
| 5:12.447 | נאמר |  |
| 5:12.448 | כמו |  |
| 5:12.449 | שאמרנו |  |
| 5:12.450 | יבוא |  |
| 5:12.451 | הכל |  |
| 5:12.452 | באמת |  |
| 5:12.453 | ובאמונה |  |
| 5:12.454 | עכ"ל | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

