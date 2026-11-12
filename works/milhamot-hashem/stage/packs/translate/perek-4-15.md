# Translation pack: milhamot-hashem perek-4-15 (Milchamot ha-Shem 4:24–4:25)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-4-15.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-4-15`.

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

- **הקדמת המחבר / הקדמת המו"ל / המו"ל** → "The Author's Introduction" / "The
  Publisher's Introduction" / "the publisher" (1:10–1:13 are his voice).
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
- **Peoples and places**: תימן → "Yemen"; צנעא → "Sanaa"; ממשלת תוגרמה → "the
  government of Turkey" (note: Togarmah); הישמעאלים → "the Ishmaelites" (note:
  Muslims); המושלמנים → "the Muslims"; הנוצרים / הנצרים → "the Christians";
  השונים → "the Dualists"; האמגושים → "the Magians"; מאמיני השלוש → "the
  believers in the Trinity"; מערב beside Ashkenaz, Sefarad, Yemen → "the
  Maghreb" (note: "the West"); בא"י → "in the Land of Israel".

## Terms

- **ספירה / ספירות / העשיריות**: fixed, "sefirah / sefirot" (not "emanations",
  "spheres", "numbers"); העשיריות → "the tens" (note). Names transliterated,
  glossed first in a unit: Keter (Crown), Chokhmah, Binah, Chesed, Gevurah,
  Tif'eret, Netzach, Hod, Yesod, Malkhut.
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
- **פרצוף / פרצופים**: the Kabbalistic entities → "partzuf / partzufim";
  generic, with צורה or תמונה → "face"; a play on both, a note. Never "configuration" or
  "countenance-structure".
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
- **ב"ד** → "court" (ב"ד הגדול → "the Great Court"); "Beit Din" only as a
  title he gives in transliteration.
- **עמי הארץ** → "the unlearned" (as Avot's עם הארץ), the Hebrew in a note; not
  "common folk" or "the ignorant".
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **שלשה פנים** → "three aspects" (note: "faces").
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
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


## Milchamot ha-Shem 4:24

Hebrew (Wikisource): לג) כל משכיל מבין המעיין מעט בספרי הזהר והמקובלים, ידע ויכיר כי הספירות והפרצופים הנזכרים בדבריהם קדמו לבריאת שמים וארץ וכל צבאם כנראה מדברי מהרח"ו בספר עץ חיים שער הכללים ושער עיגולים ויושר. ומבואר עוד באר היטב שער שבירת הכלים פ"ג וז"ל ונבאר סדר יציאת מלכים ונתחיל מן הראשון שהוא הדעת אשר זה יצא ראשונה. וכאשר לא היה יכול הכלי לסבול כנזכר לעיל נשבר וירד למטה בעולם הבריאה. רצוני לומר במקום שיהיה עתיד להיות עולם הבריאה אחר כך כי הרי עדיין לא היה נברא עולם הבריאה, ונפל הכלי הזה במקום שעתיד הדעת דבריאה להיות עכ"ל.

| id | word | punctuation after |
|---|---|---|
| 4:24.1 | לג | ) |
| 4:24.2 | כל |  |
| 4:24.3 | משכיל |  |
| 4:24.4 | מבין |  |
| 4:24.5 | המעיין |  |
| 4:24.6 | מעט |  |
| 4:24.7 | בספרי |  |
| 4:24.8 | הזהר |  |
| 4:24.9 | והמקובלים | , |
| 4:24.10 | ידע |  |
| 4:24.11 | ויכיר |  |
| 4:24.12 | כי |  |
| 4:24.13 | הספירות |  |
| 4:24.14 | והפרצופים |  |
| 4:24.15 | הנזכרים |  |
| 4:24.16 | בדבריהם |  |
| 4:24.17 | קדמו |  |
| 4:24.18 | לבריאת |  |
| 4:24.19 | שמים |  |
| 4:24.20 | וארץ |  |
| 4:24.21 | וכל |  |
| 4:24.22 | צבאם |  |
| 4:24.23 | כנראה |  |
| 4:24.24 | מדברי |  |
| 4:24.25 | מהרח"ו |  |
| 4:24.26 | בספר |  |
| 4:24.27 | עץ |  |
| 4:24.28 | חיים |  |
| 4:24.29 | שער |  |
| 4:24.30 | הכללים |  |
| 4:24.31 | ושער |  |
| 4:24.32 | עיגולים |  |
| 4:24.33 | ויושר | . |
| 4:24.34 | ומבואר |  |
| 4:24.35 | עוד |  |
| 4:24.36 | באר |  |
| 4:24.37 | היטב |  |
| 4:24.38 | שער |  |
| 4:24.39 | שבירת |  |
| 4:24.40 | הכלים |  |
| 4:24.41 | פ"ג |  |
| 4:24.42 | וז"ל |  |
| 4:24.43 | ונבאר |  |
| 4:24.44 | סדר |  |
| 4:24.45 | יציאת |  |
| 4:24.46 | מלכים |  |
| 4:24.47 | ונתחיל |  |
| 4:24.48 | מן |  |
| 4:24.49 | הראשון |  |
| 4:24.50 | שהוא |  |
| 4:24.51 | הדעת |  |
| 4:24.52 | אשר |  |
| 4:24.53 | זה |  |
| 4:24.54 | יצא |  |
| 4:24.55 | ראשונה | . |
| 4:24.56 | וכאשר |  |
| 4:24.57 | לא |  |
| 4:24.58 | היה |  |
| 4:24.59 | יכול |  |
| 4:24.60 | הכלי |  |
| 4:24.61 | לסבול |  |
| 4:24.62 | כנזכר |  |
| 4:24.63 | לעיל |  |
| 4:24.64 | נשבר |  |
| 4:24.65 | וירד |  |
| 4:24.66 | למטה |  |
| 4:24.67 | בעולם |  |
| 4:24.68 | הבריאה | . |
| 4:24.69 | רצוני |  |
| 4:24.70 | לומר |  |
| 4:24.71 | במקום |  |
| 4:24.72 | שיהיה |  |
| 4:24.73 | עתיד |  |
| 4:24.74 | להיות |  |
| 4:24.75 | עולם |  |
| 4:24.76 | הבריאה |  |
| 4:24.77 | אחר |  |
| 4:24.78 | כך |  |
| 4:24.79 | כי |  |
| 4:24.80 | הרי |  |
| 4:24.81 | עדיין |  |
| 4:24.82 | לא |  |
| 4:24.83 | היה |  |
| 4:24.84 | נברא |  |
| 4:24.85 | עולם |  |
| 4:24.86 | הבריאה | , |
| 4:24.87 | ונפל |  |
| 4:24.88 | הכלי |  |
| 4:24.89 | הזה |  |
| 4:24.90 | במקום |  |
| 4:24.91 | שעתיד |  |
| 4:24.92 | הדעת |  |
| 4:24.93 | דבריאה |  |
| 4:24.94 | להיות |  |
| 4:24.95 | עכ"ל | . |


## Milchamot ha-Shem 4:25

Hebrew (Wikisource): והיוצא לנו מזה כי רבותינו המקובלים החדשים שגגו לשאול ולמצוא פתרונים במה שלמעלה ומה למטה מה לפנים מה לאחור. דבר שקבלו חז"ל בעלי התלמוד בבלי וירושלמי ומדרש רבה ותנחומא וזולתם ממדרשי רז"ל לאסור. כדגר' בחגיגה פרק אין דורשין. כי שאל נא לימים ראשונים. יכול ישאל אדם קדם שנברא העולם ת"ל למן היום אשר ברא אלהים אדם על הארץ. יכול לא ישאל אדם מששת ימי בראשית ת"ל לימים ראשונים (פי' רש"י מיום ראשון) יכול ישאל אדם מה למעלה מה למטה ומה לפנים ומה לאחור תלמוד לומר ולמקצה השמים ועד קצה השמים. מקצה השמים ועוד קצה השמים אתה שואל. ואי אתה שואל מה למעלה ומה למטה ומה לפנים ומה לאחור עכ"ל הגמרא. ועיין במהרש"א בחדושי אגדות במעשה דר' אלעזר שאמר לר' יוחנן תא ואגמרך מעשה מרכבה ואמר ליה אכתי לא קשאי וכד קש אמר ליה ר' אסי תא ואגמרך מעשה מרכבה. אמר ליה אי זכאי הוה גמירנא מר' יוחנן רבך ונמנע ולא רצה ללמוד. לא בילדותו ולא בזקנותו. ושם הרגיש הר' מהרש"א שדברי הקבלה החדשה היא מדברת למעלה ממעשה מרכבה וראוי יותר להסתירה ולא לדרוש בה ברבים ע"ש. ואני אומר דמדינא אסור לדרוש בה כלל לא ביחיד ולא ברבים. שהרי במעשה מרכבה לא התירו חכמים לדרוש בהם ולחכם ומבין ראשי פרקים אלא עד החשמל. ואיכא דאמרי עד וארא ולא יותר דמשם ולמעלה אסור לדבר כלל, כפי קבלת חכמים ואזהרתם. והעובר על דברי חכמים חייב מיתה. ומכל זה למדנו ששגגה גדולה שגגו רבותינו בעלי הקבלה החדשה והגיסו דעתם ומלאם לבם לדרוש במה שלא הותר והורשה להם כמו שיתבאר לקמן. ועלתה בלבם מחשבה זרה שבהתקרב זמן הגאולה תבטל הלכה זו ויהיה מותר לדרוש במה שלא הותר להם מוקדם. כמו שמבואר בפסק שעשה ר' יצחק דלמאש המודפס בתחלת ס' הזהר. וטעות הוא בידו לפסוק הלכה מדברי הזהר היפך המשנה והתלמוד והפוסקים ואין לסמוך עליו כלל. ולא שמו המקובלים החדשים אל לבם מאמר רז"ל שכל הלכות של תורה שבעל פה לא יבטלו כמו שכ' הרמב"ם ז"ל בסוף הלכות מגלה. ובעברם על דברי חז"ל אירע להם בשגגת זו להאמין רבוי באלהות ולהאמין במציאות אלהים אחרים, ולשתף שם שמים ודבר אחר שעליו אמר ר' שמעון כל המשתף שם שמים ודבר אחר נעקר מן העולם. ובשלהי פסחים גרסינן מאי למכסה עתיק, זה המכסה דברים שכיסה עתיק יומין, ומאי נינהו סתרי תורה (פ' רש"י מעשה מרכבה ומע' בראשית) ואיכא דאמרי זה המגלה דברים שכיסה עתיק יומין, ומאי נינהו טעמי תורה ע"כ, הרי לך שהמגלה מעשה מרכבה לא יזכה לטוב הצפון לצדיקים, כ"ש וקל וחומר דברים שהם למעלה ממעשה מרכבה שאמור לדבר בהם כלל, ולא ניתן רשות לגלות אלא מטעמי תורה בלבד, אבל סתרי תורה לא ניתן רשות לגלותן, וכל שכן לדרשן ברבים בהדיא בבית המדרש אפי' לתלמידי חכמים, כל שכן בין ההמון אשר לא ידעו ולא יבינו, כי יביאם לידי מינות באמונת אלוהות רבים כאשר ראו עינינו, וחלילה לחשוב על ר' שמעון ב"י התנא שיעבור על איסור זה לדבר במה שלמעלה ממעשה מרכבה ובמה שהיה לפנים קודם בריאת העולם, וגם לא לשום תנא או אמורא לעשות כן ולומר שנגלה הקב"ה והתיר לו את האסור לדרוש בו, כי זו נביאות שקר היא (ומיתתו בחנק כמו שכתב הרמב"ם בהקדמת סדר זרעים) שאין נביא רשאי לחדש דבר מעתה כי לא בשמים היא.

| id | word | punctuation after |
|---|---|---|
| 4:25.1 | והיוצא |  |
| 4:25.2 | לנו |  |
| 4:25.3 | מזה |  |
| 4:25.4 | כי |  |
| 4:25.5 | רבותינו |  |
| 4:25.6 | המקובלים |  |
| 4:25.7 | החדשים |  |
| 4:25.8 | שגגו |  |
| 4:25.9 | לשאול |  |
| 4:25.10 | ולמצוא |  |
| 4:25.11 | פתרונים |  |
| 4:25.12 | במה |  |
| 4:25.13 | שלמעלה |  |
| 4:25.14 | ומה |  |
| 4:25.15 | למטה |  |
| 4:25.16 | מה |  |
| 4:25.17 | לפנים |  |
| 4:25.18 | מה |  |
| 4:25.19 | לאחור | . |
| 4:25.20 | דבר |  |
| 4:25.21 | שקבלו |  |
| 4:25.22 | חז"ל |  |
| 4:25.23 | בעלי |  |
| 4:25.24 | התלמוד |  |
| 4:25.25 | בבלי |  |
| 4:25.26 | וירושלמי |  |
| 4:25.27 | ומדרש |  |
| 4:25.28 | רבה |  |
| 4:25.29 | ותנחומא |  |
| 4:25.30 | וזולתם |  |
| 4:25.31 | ממדרשי |  |
| 4:25.32 | רז"ל |  |
| 4:25.33 | לאסור | . |
| 4:25.34 | כדגר | ' |
| 4:25.35 | בחגיגה |  |
| 4:25.36 | פרק |  |
| 4:25.37 | אין |  |
| 4:25.38 | דורשין | . |
| 4:25.39 | כי |  |
| 4:25.40 | שאל |  |
| 4:25.41 | נא |  |
| 4:25.42 | לימים |  |
| 4:25.43 | ראשונים | . |
| 4:25.44 | יכול |  |
| 4:25.45 | ישאל |  |
| 4:25.46 | אדם |  |
| 4:25.47 | קדם |  |
| 4:25.48 | שנברא |  |
| 4:25.49 | העולם |  |
| 4:25.50 | ת"ל |  |
| 4:25.51 | למן |  |
| 4:25.52 | היום |  |
| 4:25.53 | אשר |  |
| 4:25.54 | ברא |  |
| 4:25.55 | אלהים |  |
| 4:25.56 | אדם |  |
| 4:25.57 | על |  |
| 4:25.58 | הארץ | . |
| 4:25.59 | יכול |  |
| 4:25.60 | לא |  |
| 4:25.61 | ישאל |  |
| 4:25.62 | אדם |  |
| 4:25.63 | מששת |  |
| 4:25.64 | ימי |  |
| 4:25.65 | בראשית |  |
| 4:25.66 | ת"ל |  |
| 4:25.67 | לימים |  |
| 4:25.68 | ראשונים |  |
| 4:25.69 | פי | ' |
| 4:25.70 | רש"י |  |
| 4:25.71 | מיום |  |
| 4:25.72 | ראשון | ) |
| 4:25.73 | יכול |  |
| 4:25.74 | ישאל |  |
| 4:25.75 | אדם |  |
| 4:25.76 | מה |  |
| 4:25.77 | למעלה |  |
| 4:25.78 | מה |  |
| 4:25.79 | למטה |  |
| 4:25.80 | ומה |  |
| 4:25.81 | לפנים |  |
| 4:25.82 | ומה |  |
| 4:25.83 | לאחור |  |
| 4:25.84 | תלמוד |  |
| 4:25.85 | לומר |  |
| 4:25.86 | ולמקצה |  |
| 4:25.87 | השמים |  |
| 4:25.88 | ועד |  |
| 4:25.89 | קצה |  |
| 4:25.90 | השמים | . |
| 4:25.91 | מקצה |  |
| 4:25.92 | השמים |  |
| 4:25.93 | ועוד |  |
| 4:25.94 | קצה |  |
| 4:25.95 | השמים |  |
| 4:25.96 | אתה |  |
| 4:25.97 | שואל | . |
| 4:25.98 | ואי |  |
| 4:25.99 | אתה |  |
| 4:25.100 | שואל |  |
| 4:25.101 | מה |  |
| 4:25.102 | למעלה |  |
| 4:25.103 | ומה |  |
| 4:25.104 | למטה |  |
| 4:25.105 | ומה |  |
| 4:25.106 | לפנים |  |
| 4:25.107 | ומה |  |
| 4:25.108 | לאחור |  |
| 4:25.109 | עכ"ל |  |
| 4:25.110 | הגמרא | . |
| 4:25.111 | ועיין |  |
| 4:25.112 | במהרש"א |  |
| 4:25.113 | בחדושי |  |
| 4:25.114 | אגדות |  |
| 4:25.115 | במעשה |  |
| 4:25.116 | דר | ' |
| 4:25.117 | אלעזר |  |
| 4:25.118 | שאמר |  |
| 4:25.119 | לר | ' |
| 4:25.120 | יוחנן |  |
| 4:25.121 | תא |  |
| 4:25.122 | ואגמרך |  |
| 4:25.123 | מעשה |  |
| 4:25.124 | מרכבה |  |
| 4:25.125 | ואמר |  |
| 4:25.126 | ליה |  |
| 4:25.127 | אכתי |  |
| 4:25.128 | לא |  |
| 4:25.129 | קשאי |  |
| 4:25.130 | וכד |  |
| 4:25.131 | קש |  |
| 4:25.132 | אמר |  |
| 4:25.133 | ליה |  |
| 4:25.134 | ר | ' |
| 4:25.135 | אסי |  |
| 4:25.136 | תא |  |
| 4:25.137 | ואגמרך |  |
| 4:25.138 | מעשה |  |
| 4:25.139 | מרכבה | . |
| 4:25.140 | אמר |  |
| 4:25.141 | ליה |  |
| 4:25.142 | אי |  |
| 4:25.143 | זכאי |  |
| 4:25.144 | הוה |  |
| 4:25.145 | גמירנא |  |
| 4:25.146 | מר | ' |
| 4:25.147 | יוחנן |  |
| 4:25.148 | רבך |  |
| 4:25.149 | ונמנע |  |
| 4:25.150 | ולא |  |
| 4:25.151 | רצה |  |
| 4:25.152 | ללמוד | . |
| 4:25.153 | לא |  |
| 4:25.154 | בילדותו |  |
| 4:25.155 | ולא |  |
| 4:25.156 | בזקנותו | . |
| 4:25.157 | ושם |  |
| 4:25.158 | הרגיש |  |
| 4:25.159 | הר | ' |
| 4:25.160 | מהרש"א |  |
| 4:25.161 | שדברי |  |
| 4:25.162 | הקבלה |  |
| 4:25.163 | החדשה |  |
| 4:25.164 | היא |  |
| 4:25.165 | מדברת |  |
| 4:25.166 | למעלה |  |
| 4:25.167 | ממעשה |  |
| 4:25.168 | מרכבה |  |
| 4:25.169 | וראוי |  |
| 4:25.170 | יותר |  |
| 4:25.171 | להסתירה |  |
| 4:25.172 | ולא |  |
| 4:25.173 | לדרוש |  |
| 4:25.174 | בה |  |
| 4:25.175 | ברבים |  |
| 4:25.176 | ע"ש | . |
| 4:25.177 | ואני |  |
| 4:25.178 | אומר |  |
| 4:25.179 | דמדינא |  |
| 4:25.180 | אסור |  |
| 4:25.181 | לדרוש |  |
| 4:25.182 | בה |  |
| 4:25.183 | כלל |  |
| 4:25.184 | לא |  |
| 4:25.185 | ביחיד |  |
| 4:25.186 | ולא |  |
| 4:25.187 | ברבים | . |
| 4:25.188 | שהרי |  |
| 4:25.189 | במעשה |  |
| 4:25.190 | מרכבה |  |
| 4:25.191 | לא |  |
| 4:25.192 | התירו |  |
| 4:25.193 | חכמים |  |
| 4:25.194 | לדרוש |  |
| 4:25.195 | בהם |  |
| 4:25.196 | ולחכם |  |
| 4:25.197 | ומבין |  |
| 4:25.198 | ראשי |  |
| 4:25.199 | פרקים |  |
| 4:25.200 | אלא |  |
| 4:25.201 | עד |  |
| 4:25.202 | החשמל | . |
| 4:25.203 | ואיכא |  |
| 4:25.204 | דאמרי |  |
| 4:25.205 | עד |  |
| 4:25.206 | וארא |  |
| 4:25.207 | ולא |  |
| 4:25.208 | יותר |  |
| 4:25.209 | דמשם |  |
| 4:25.210 | ולמעלה |  |
| 4:25.211 | אסור |  |
| 4:25.212 | לדבר |  |
| 4:25.213 | כלל | , |
| 4:25.214 | כפי |  |
| 4:25.215 | קבלת |  |
| 4:25.216 | חכמים |  |
| 4:25.217 | ואזהרתם | . |
| 4:25.218 | והעובר |  |
| 4:25.219 | על |  |
| 4:25.220 | דברי |  |
| 4:25.221 | חכמים |  |
| 4:25.222 | חייב |  |
| 4:25.223 | מיתה | . |
| 4:25.224 | ומכל |  |
| 4:25.225 | זה |  |
| 4:25.226 | למדנו |  |
| 4:25.227 | ששגגה |  |
| 4:25.228 | גדולה |  |
| 4:25.229 | שגגו |  |
| 4:25.230 | רבותינו |  |
| 4:25.231 | בעלי |  |
| 4:25.232 | הקבלה |  |
| 4:25.233 | החדשה |  |
| 4:25.234 | והגיסו |  |
| 4:25.235 | דעתם |  |
| 4:25.236 | ומלאם |  |
| 4:25.237 | לבם |  |
| 4:25.238 | לדרוש |  |
| 4:25.239 | במה |  |
| 4:25.240 | שלא |  |
| 4:25.241 | הותר |  |
| 4:25.242 | והורשה |  |
| 4:25.243 | להם |  |
| 4:25.244 | כמו |  |
| 4:25.245 | שיתבאר |  |
| 4:25.246 | לקמן | . |
| 4:25.247 | ועלתה |  |
| 4:25.248 | בלבם |  |
| 4:25.249 | מחשבה |  |
| 4:25.250 | זרה |  |
| 4:25.251 | שבהתקרב |  |
| 4:25.252 | זמן |  |
| 4:25.253 | הגאולה |  |
| 4:25.254 | תבטל |  |
| 4:25.255 | הלכה |  |
| 4:25.256 | זו |  |
| 4:25.257 | ויהיה |  |
| 4:25.258 | מותר |  |
| 4:25.259 | לדרוש |  |
| 4:25.260 | במה |  |
| 4:25.261 | שלא |  |
| 4:25.262 | הותר |  |
| 4:25.263 | להם |  |
| 4:25.264 | מוקדם | . |
| 4:25.265 | כמו |  |
| 4:25.266 | שמבואר |  |
| 4:25.267 | בפסק |  |
| 4:25.268 | שעשה |  |
| 4:25.269 | ר | ' |
| 4:25.270 | יצחק |  |
| 4:25.271 | דלמאש |  |
| 4:25.272 | המודפס |  |
| 4:25.273 | בתחלת |  |
| 4:25.274 | ס | ' |
| 4:25.275 | הזהר | . |
| 4:25.276 | וטעות |  |
| 4:25.277 | הוא |  |
| 4:25.278 | בידו |  |
| 4:25.279 | לפסוק |  |
| 4:25.280 | הלכה |  |
| 4:25.281 | מדברי |  |
| 4:25.282 | הזהר |  |
| 4:25.283 | היפך |  |
| 4:25.284 | המשנה |  |
| 4:25.285 | והתלמוד |  |
| 4:25.286 | והפוסקים |  |
| 4:25.287 | ואין |  |
| 4:25.288 | לסמוך |  |
| 4:25.289 | עליו |  |
| 4:25.290 | כלל | . |
| 4:25.291 | ולא |  |
| 4:25.292 | שמו |  |
| 4:25.293 | המקובלים |  |
| 4:25.294 | החדשים |  |
| 4:25.295 | אל |  |
| 4:25.296 | לבם |  |
| 4:25.297 | מאמר |  |
| 4:25.298 | רז"ל |  |
| 4:25.299 | שכל |  |
| 4:25.300 | הלכות |  |
| 4:25.301 | של |  |
| 4:25.302 | תורה |  |
| 4:25.303 | שבעל |  |
| 4:25.304 | פה |  |
| 4:25.305 | לא |  |
| 4:25.306 | יבטלו |  |
| 4:25.307 | כמו |  |
| 4:25.308 | שכ | ' |
| 4:25.309 | הרמב"ם |  |
| 4:25.310 | ז"ל |  |
| 4:25.311 | בסוף |  |
| 4:25.312 | הלכות |  |
| 4:25.313 | מגלה | . |
| 4:25.314 | ובעברם |  |
| 4:25.315 | על |  |
| 4:25.316 | דברי |  |
| 4:25.317 | חז"ל |  |
| 4:25.318 | אירע |  |
| 4:25.319 | להם |  |
| 4:25.320 | בשגגת |  |
| 4:25.321 | זו |  |
| 4:25.322 | להאמין |  |
| 4:25.323 | רבוי |  |
| 4:25.324 | באלהות |  |
| 4:25.325 | ולהאמין |  |
| 4:25.326 | במציאות |  |
| 4:25.327 | אלהים |  |
| 4:25.328 | אחרים | , |
| 4:25.329 | ולשתף |  |
| 4:25.330 | שם |  |
| 4:25.331 | שמים |  |
| 4:25.332 | ודבר |  |
| 4:25.333 | אחר |  |
| 4:25.334 | שעליו |  |
| 4:25.335 | אמר |  |
| 4:25.336 | ר | ' |
| 4:25.337 | שמעון |  |
| 4:25.338 | כל |  |
| 4:25.339 | המשתף |  |
| 4:25.340 | שם |  |
| 4:25.341 | שמים |  |
| 4:25.342 | ודבר |  |
| 4:25.343 | אחר |  |
| 4:25.344 | נעקר |  |
| 4:25.345 | מן |  |
| 4:25.346 | העולם | . |
| 4:25.347 | ובשלהי |  |
| 4:25.348 | פסחים |  |
| 4:25.349 | גרסינן |  |
| 4:25.350 | מאי |  |
| 4:25.351 | למכסה |  |
| 4:25.352 | עתיק | , |
| 4:25.353 | זה |  |
| 4:25.354 | המכסה |  |
| 4:25.355 | דברים |  |
| 4:25.356 | שכיסה |  |
| 4:25.357 | עתיק |  |
| 4:25.358 | יומין | , |
| 4:25.359 | ומאי |  |
| 4:25.360 | נינהו |  |
| 4:25.361 | סתרי |  |
| 4:25.362 | תורה |  |
| 4:25.363 | פ | ' |
| 4:25.364 | רש"י |  |
| 4:25.365 | מעשה |  |
| 4:25.366 | מרכבה |  |
| 4:25.367 | ומע | ' |
| 4:25.368 | בראשית | ) |
| 4:25.369 | ואיכא |  |
| 4:25.370 | דאמרי |  |
| 4:25.371 | זה |  |
| 4:25.372 | המגלה |  |
| 4:25.373 | דברים |  |
| 4:25.374 | שכיסה |  |
| 4:25.375 | עתיק |  |
| 4:25.376 | יומין | , |
| 4:25.377 | ומאי |  |
| 4:25.378 | נינהו |  |
| 4:25.379 | טעמי |  |
| 4:25.380 | תורה |  |
| 4:25.381 | ע"כ | , |
| 4:25.382 | הרי |  |
| 4:25.383 | לך |  |
| 4:25.384 | שהמגלה |  |
| 4:25.385 | מעשה |  |
| 4:25.386 | מרכבה |  |
| 4:25.387 | לא |  |
| 4:25.388 | יזכה |  |
| 4:25.389 | לטוב |  |
| 4:25.390 | הצפון |  |
| 4:25.391 | לצדיקים | , |
| 4:25.392 | כ"ש |  |
| 4:25.393 | וקל |  |
| 4:25.394 | וחומר |  |
| 4:25.395 | דברים |  |
| 4:25.396 | שהם |  |
| 4:25.397 | למעלה |  |
| 4:25.398 | ממעשה |  |
| 4:25.399 | מרכבה |  |
| 4:25.400 | שאמור |  |
| 4:25.401 | לדבר |  |
| 4:25.402 | בהם |  |
| 4:25.403 | כלל | , |
| 4:25.404 | ולא |  |
| 4:25.405 | ניתן |  |
| 4:25.406 | רשות |  |
| 4:25.407 | לגלות |  |
| 4:25.408 | אלא |  |
| 4:25.409 | מטעמי |  |
| 4:25.410 | תורה |  |
| 4:25.411 | בלבד | , |
| 4:25.412 | אבל |  |
| 4:25.413 | סתרי |  |
| 4:25.414 | תורה |  |
| 4:25.415 | לא |  |
| 4:25.416 | ניתן |  |
| 4:25.417 | רשות |  |
| 4:25.418 | לגלותן | , |
| 4:25.419 | וכל |  |
| 4:25.420 | שכן |  |
| 4:25.421 | לדרשן |  |
| 4:25.422 | ברבים |  |
| 4:25.423 | בהדיא |  |
| 4:25.424 | בבית |  |
| 4:25.425 | המדרש |  |
| 4:25.426 | אפי | ' |
| 4:25.427 | לתלמידי |  |
| 4:25.428 | חכמים | , |
| 4:25.429 | כל |  |
| 4:25.430 | שכן |  |
| 4:25.431 | בין |  |
| 4:25.432 | ההמון |  |
| 4:25.433 | אשר |  |
| 4:25.434 | לא |  |
| 4:25.435 | ידעו |  |
| 4:25.436 | ולא |  |
| 4:25.437 | יבינו | , |
| 4:25.438 | כי |  |
| 4:25.439 | יביאם |  |
| 4:25.440 | לידי |  |
| 4:25.441 | מינות |  |
| 4:25.442 | באמונת |  |
| 4:25.443 | אלוהות |  |
| 4:25.444 | רבים |  |
| 4:25.445 | כאשר |  |
| 4:25.446 | ראו |  |
| 4:25.447 | עינינו | , |
| 4:25.448 | וחלילה |  |
| 4:25.449 | לחשוב |  |
| 4:25.450 | על |  |
| 4:25.451 | ר | ' |
| 4:25.452 | שמעון |  |
| 4:25.453 | ב"י |  |
| 4:25.454 | התנא |  |
| 4:25.455 | שיעבור |  |
| 4:25.456 | על |  |
| 4:25.457 | איסור |  |
| 4:25.458 | זה |  |
| 4:25.459 | לדבר |  |
| 4:25.460 | במה |  |
| 4:25.461 | שלמעלה |  |
| 4:25.462 | ממעשה |  |
| 4:25.463 | מרכבה |  |
| 4:25.464 | ובמה |  |
| 4:25.465 | שהיה |  |
| 4:25.466 | לפנים |  |
| 4:25.467 | קודם |  |
| 4:25.468 | בריאת |  |
| 4:25.469 | העולם | , |
| 4:25.470 | וגם |  |
| 4:25.471 | לא |  |
| 4:25.472 | לשום |  |
| 4:25.473 | תנא |  |
| 4:25.474 | או |  |
| 4:25.475 | אמורא |  |
| 4:25.476 | לעשות |  |
| 4:25.477 | כן |  |
| 4:25.478 | ולומר |  |
| 4:25.479 | שנגלה |  |
| 4:25.480 | הקב"ה |  |
| 4:25.481 | והתיר |  |
| 4:25.482 | לו |  |
| 4:25.483 | את |  |
| 4:25.484 | האסור |  |
| 4:25.485 | לדרוש |  |
| 4:25.486 | בו | , |
| 4:25.487 | כי |  |
| 4:25.488 | זו |  |
| 4:25.489 | נביאות |  |
| 4:25.490 | שקר |  |
| 4:25.491 | היא |  |
| 4:25.492 | ומיתתו |  |
| 4:25.493 | בחנק |  |
| 4:25.494 | כמו |  |
| 4:25.495 | שכתב |  |
| 4:25.496 | הרמב"ם |  |
| 4:25.497 | בהקדמת |  |
| 4:25.498 | סדר |  |
| 4:25.499 | זרעים | ) |
| 4:25.500 | שאין |  |
| 4:25.501 | נביא |  |
| 4:25.502 | רשאי |  |
| 4:25.503 | לחדש |  |
| 4:25.504 | דבר |  |
| 4:25.505 | מעתה |  |
| 4:25.506 | כי |  |
| 4:25.507 | לא |  |
| 4:25.508 | בשמים |  |
| 4:25.509 | היא | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

