# Translation pack: milhamot-hashem perek-6-25 (Milchamot ha-Shem 6:38–6:38)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-6-25.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-6-25`.

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
- **קצר אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
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
- **מורא** → "fear".
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **העתיק** → "rendered"; "copied" in a note when it could mean copying.
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **דמי** → "blood", or "price/value" in a sale (דמי השור); the other in a note.
- **דמיון** → "analogy" for an argument, "likeness" for resemblance.
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
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


## Milchamot ha-Shem 6:38

Hebrew (Wikisource): עז) דבר הלמד מדברי הר' הנז' כי דבר הגורם לכל אלה השגיאות ששגגו רבותינו בעלי הקבלה החדשה להאמין באלהים רבים, הוא כי האמינו לדברי נביא השקר המסית מחבר ספר הזהר. ולא נתנו אל לבם אזהרת תורתינו הקדושה באמרה ונשמרתם מאד לנפשותיכם כי לא ראיתם כל תמונה. ולא לאזהרת הנביא שאמר ואל מי תדמיון אל. ואל מי תדמיוני ואשוה. ובקשו חשבונות רבים בכמה מקראות שנאמרו על דרך ההשאלה והרחבת לשון בדרך המליצה לפרשם על דרך הנחתם הראשונה. כגון בצלמינו כדמותינו ושאר לשונות כנז' לעיל שלא על דרך שפירשו רז"ל והגאונים החכמים. עד שגרם להם שבאו לכלל טעות גדולה לעבוד לנברא שהוא קצר אפים (זעיר אנפין) או מלכא קדישא דכל קדישין לדעת עז לאלהים. והמציאו כמה רמזים ורמזי רמזים שאין להם שחר או במספר האותיות, או בחילופיהן וחילופי חלופיהן ומלואיהן. ושבוש סדריהן ובלבולן. כמו מה תצעק אלי מאצ"ל העתי"ק. כדי למצוא רמז שיתפלל לז"א ויכלול עמו פרצוף עתיק יומין לשתף יחד כל החמשה פרצופין עם העתיק. כמו שכ' כסא אליהו דבר שאסרה תורה לכלול ולשתף עם אלוהינו שום דבר. כמו שאמר רשב"י ע"ה כל המשתף שם שמים ודבר אחר נעקר מן העולם שנ' בלתי לה' לבדו. וכל זה נמשך מחמת מיעוט התבוננות בדברי רז"ל. וגם בעבור שנביא השקר מחבר הזהר שהמציא רעיונות אלו ילדי נכרים. תלה אותם ברשב"י. התנא החסיד בשקר. והם לתומם לא חשו לבדוק אחריו. והאמינו באמונת אומן שמפי רשב"י וחביריו יצאו כבושים. וכבר ביארתי לעיל שחלילה לשום תנא או אמורא לומר כן. עוד כתב הרב שבילי אמונה למעלה מזה המאמר, כי הבורא ית' לא יאויך. ר"ל שאינו בעל איכות, ומהנמנע להיות מציאות איכות, אלא במאוייך, מפני שהאיכות מקרה, והמקרה אין לו עמידה אלא בעצם. ולכן הבורא ית' אינו נושא האיכות. גם בעל האיכות הוא אשר לו התאר בעצמו ולו ית' אין לו תואר וכו' ועיין במורה נבוכים פרק נ"ב לראשון.

| id | word | punctuation after |
|---|---|---|
| 6:38.1 | עז | ) |
| 6:38.2 | דבר |  |
| 6:38.3 | הלמד |  |
| 6:38.4 | מדברי |  |
| 6:38.5 | הר | ' |
| 6:38.6 | הנז | ' |
| 6:38.7 | כי |  |
| 6:38.8 | דבר |  |
| 6:38.9 | הגורם |  |
| 6:38.10 | לכל |  |
| 6:38.11 | אלה |  |
| 6:38.12 | השגיאות |  |
| 6:38.13 | ששגגו |  |
| 6:38.14 | רבותינו |  |
| 6:38.15 | בעלי |  |
| 6:38.16 | הקבלה |  |
| 6:38.17 | החדשה |  |
| 6:38.18 | להאמין |  |
| 6:38.19 | באלהים |  |
| 6:38.20 | רבים | , |
| 6:38.21 | הוא |  |
| 6:38.22 | כי |  |
| 6:38.23 | האמינו |  |
| 6:38.24 | לדברי |  |
| 6:38.25 | נביא |  |
| 6:38.26 | השקר |  |
| 6:38.27 | המסית |  |
| 6:38.28 | מחבר |  |
| 6:38.29 | ספר |  |
| 6:38.30 | הזהר | . |
| 6:38.31 | ולא |  |
| 6:38.32 | נתנו |  |
| 6:38.33 | אל |  |
| 6:38.34 | לבם |  |
| 6:38.35 | אזהרת |  |
| 6:38.36 | תורתינו |  |
| 6:38.37 | הקדושה |  |
| 6:38.38 | באמרה |  |
| 6:38.39 | ונשמרתם |  |
| 6:38.40 | מאד |  |
| 6:38.41 | לנפשותיכם |  |
| 6:38.42 | כי |  |
| 6:38.43 | לא |  |
| 6:38.44 | ראיתם |  |
| 6:38.45 | כל |  |
| 6:38.46 | תמונה | . |
| 6:38.47 | ולא |  |
| 6:38.48 | לאזהרת |  |
| 6:38.49 | הנביא |  |
| 6:38.50 | שאמר |  |
| 6:38.51 | ואל |  |
| 6:38.52 | מי |  |
| 6:38.53 | תדמיון |  |
| 6:38.54 | אל | . |
| 6:38.55 | ואל |  |
| 6:38.56 | מי |  |
| 6:38.57 | תדמיוני |  |
| 6:38.58 | ואשוה | . |
| 6:38.59 | ובקשו |  |
| 6:38.60 | חשבונות |  |
| 6:38.61 | רבים |  |
| 6:38.62 | בכמה |  |
| 6:38.63 | מקראות |  |
| 6:38.64 | שנאמרו |  |
| 6:38.65 | על |  |
| 6:38.66 | דרך |  |
| 6:38.67 | ההשאלה |  |
| 6:38.68 | והרחבת |  |
| 6:38.69 | לשון |  |
| 6:38.70 | בדרך |  |
| 6:38.71 | המליצה |  |
| 6:38.72 | לפרשם |  |
| 6:38.73 | על |  |
| 6:38.74 | דרך |  |
| 6:38.75 | הנחתם |  |
| 6:38.76 | הראשונה | . |
| 6:38.77 | כגון |  |
| 6:38.78 | בצלמינו |  |
| 6:38.79 | כדמותינו |  |
| 6:38.80 | ושאר |  |
| 6:38.81 | לשונות |  |
| 6:38.82 | כנז | ' |
| 6:38.83 | לעיל |  |
| 6:38.84 | שלא |  |
| 6:38.85 | על |  |
| 6:38.86 | דרך |  |
| 6:38.87 | שפירשו |  |
| 6:38.88 | רז"ל |  |
| 6:38.89 | והגאונים |  |
| 6:38.90 | החכמים | . |
| 6:38.91 | עד |  |
| 6:38.92 | שגרם |  |
| 6:38.93 | להם |  |
| 6:38.94 | שבאו |  |
| 6:38.95 | לכלל |  |
| 6:38.96 | טעות |  |
| 6:38.97 | גדולה |  |
| 6:38.98 | לעבוד |  |
| 6:38.99 | לנברא |  |
| 6:38.100 | שהוא |  |
| 6:38.101 | קצר |  |
| 6:38.102 | אפים |  |
| 6:38.103 | זעיר |  |
| 6:38.104 | אנפין | ) |
| 6:38.105 | או |  |
| 6:38.106 | מלכא |  |
| 6:38.107 | קדישא |  |
| 6:38.108 | דכל |  |
| 6:38.109 | קדישין |  |
| 6:38.110 | לדעת |  |
| 6:38.111 | עז |  |
| 6:38.112 | לאלהים | . |
| 6:38.113 | והמציאו |  |
| 6:38.114 | כמה |  |
| 6:38.115 | רמזים |  |
| 6:38.116 | ורמזי |  |
| 6:38.117 | רמזים |  |
| 6:38.118 | שאין |  |
| 6:38.119 | להם |  |
| 6:38.120 | שחר |  |
| 6:38.121 | או |  |
| 6:38.122 | במספר |  |
| 6:38.123 | האותיות | , |
| 6:38.124 | או |  |
| 6:38.125 | בחילופיהן |  |
| 6:38.126 | וחילופי |  |
| 6:38.127 | חלופיהן |  |
| 6:38.128 | ומלואיהן | . |
| 6:38.129 | ושבוש |  |
| 6:38.130 | סדריהן |  |
| 6:38.131 | ובלבולן | . |
| 6:38.132 | כמו |  |
| 6:38.133 | מה |  |
| 6:38.134 | תצעק |  |
| 6:38.135 | אלי |  |
| 6:38.136 | מאצ"ל |  |
| 6:38.137 | העתי"ק | . |
| 6:38.138 | כדי |  |
| 6:38.139 | למצוא |  |
| 6:38.140 | רמז |  |
| 6:38.141 | שיתפלל |  |
| 6:38.142 | לז"א |  |
| 6:38.143 | ויכלול |  |
| 6:38.144 | עמו |  |
| 6:38.145 | פרצוף |  |
| 6:38.146 | עתיק |  |
| 6:38.147 | יומין |  |
| 6:38.148 | לשתף |  |
| 6:38.149 | יחד |  |
| 6:38.150 | כל |  |
| 6:38.151 | החמשה |  |
| 6:38.152 | פרצופין |  |
| 6:38.153 | עם |  |
| 6:38.154 | העתיק | . |
| 6:38.155 | כמו |  |
| 6:38.156 | שכ | ' |
| 6:38.157 | כסא |  |
| 6:38.158 | אליהו |  |
| 6:38.159 | דבר |  |
| 6:38.160 | שאסרה |  |
| 6:38.161 | תורה |  |
| 6:38.162 | לכלול |  |
| 6:38.163 | ולשתף |  |
| 6:38.164 | עם |  |
| 6:38.165 | אלוהינו |  |
| 6:38.166 | שום |  |
| 6:38.167 | דבר | . |
| 6:38.168 | כמו |  |
| 6:38.169 | שאמר |  |
| 6:38.170 | רשב"י |  |
| 6:38.171 | ע"ה |  |
| 6:38.172 | כל |  |
| 6:38.173 | המשתף |  |
| 6:38.174 | שם |  |
| 6:38.175 | שמים |  |
| 6:38.176 | ודבר |  |
| 6:38.177 | אחר |  |
| 6:38.178 | נעקר |  |
| 6:38.179 | מן |  |
| 6:38.180 | העולם |  |
| 6:38.181 | שנ | ' |
| 6:38.182 | בלתי |  |
| 6:38.183 | לה | ' |
| 6:38.184 | לבדו | . |
| 6:38.185 | וכל |  |
| 6:38.186 | זה |  |
| 6:38.187 | נמשך |  |
| 6:38.188 | מחמת |  |
| 6:38.189 | מיעוט |  |
| 6:38.190 | התבוננות |  |
| 6:38.191 | בדברי |  |
| 6:38.192 | רז"ל | . |
| 6:38.193 | וגם |  |
| 6:38.194 | בעבור |  |
| 6:38.195 | שנביא |  |
| 6:38.196 | השקר |  |
| 6:38.197 | מחבר |  |
| 6:38.198 | הזהר |  |
| 6:38.199 | שהמציא |  |
| 6:38.200 | רעיונות |  |
| 6:38.201 | אלו |  |
| 6:38.202 | ילדי |  |
| 6:38.203 | נכרים | . |
| 6:38.204 | תלה |  |
| 6:38.205 | אותם |  |
| 6:38.206 | ברשב"י | . |
| 6:38.207 | התנא |  |
| 6:38.208 | החסיד |  |
| 6:38.209 | בשקר | . |
| 6:38.210 | והם |  |
| 6:38.211 | לתומם |  |
| 6:38.212 | לא |  |
| 6:38.213 | חשו |  |
| 6:38.214 | לבדוק |  |
| 6:38.215 | אחריו | . |
| 6:38.216 | והאמינו |  |
| 6:38.217 | באמונת |  |
| 6:38.218 | אומן |  |
| 6:38.219 | שמפי |  |
| 6:38.220 | רשב"י |  |
| 6:38.221 | וחביריו |  |
| 6:38.222 | יצאו |  |
| 6:38.223 | כבושים | . |
| 6:38.224 | וכבר |  |
| 6:38.225 | ביארתי |  |
| 6:38.226 | לעיל |  |
| 6:38.227 | שחלילה |  |
| 6:38.228 | לשום |  |
| 6:38.229 | תנא |  |
| 6:38.230 | או |  |
| 6:38.231 | אמורא |  |
| 6:38.232 | לומר |  |
| 6:38.233 | כן | . |
| 6:38.234 | עוד |  |
| 6:38.235 | כתב |  |
| 6:38.236 | הרב |  |
| 6:38.237 | שבילי |  |
| 6:38.238 | אמונה |  |
| 6:38.239 | למעלה |  |
| 6:38.240 | מזה |  |
| 6:38.241 | המאמר | , |
| 6:38.242 | כי |  |
| 6:38.243 | הבורא |  |
| 6:38.244 | ית | ' |
| 6:38.245 | לא |  |
| 6:38.246 | יאויך | . |
| 6:38.247 | ר"ל |  |
| 6:38.248 | שאינו |  |
| 6:38.249 | בעל |  |
| 6:38.250 | איכות | , |
| 6:38.251 | ומהנמנע |  |
| 6:38.252 | להיות |  |
| 6:38.253 | מציאות |  |
| 6:38.254 | איכות | , |
| 6:38.255 | אלא |  |
| 6:38.256 | במאוייך | , |
| 6:38.257 | מפני |  |
| 6:38.258 | שהאיכות |  |
| 6:38.259 | מקרה | , |
| 6:38.260 | והמקרה |  |
| 6:38.261 | אין |  |
| 6:38.262 | לו |  |
| 6:38.263 | עמידה |  |
| 6:38.264 | אלא |  |
| 6:38.265 | בעצם | . |
| 6:38.266 | ולכן |  |
| 6:38.267 | הבורא |  |
| 6:38.268 | ית | ' |
| 6:38.269 | אינו |  |
| 6:38.270 | נושא |  |
| 6:38.271 | האיכות | . |
| 6:38.272 | גם |  |
| 6:38.273 | בעל |  |
| 6:38.274 | האיכות |  |
| 6:38.275 | הוא |  |
| 6:38.276 | אשר |  |
| 6:38.277 | לו |  |
| 6:38.278 | התאר |  |
| 6:38.279 | בעצמו |  |
| 6:38.280 | ולו |  |
| 6:38.281 | ית | ' |
| 6:38.282 | אין |  |
| 6:38.283 | לו |  |
| 6:38.284 | תואר |  |
| 6:38.285 | וכו | ' |
| 6:38.286 | ועיין |  |
| 6:38.287 | במורה |  |
| 6:38.288 | נבוכים |  |
| 6:38.289 | פרק |  |
| 6:38.290 | נ"ב |  |
| 6:38.291 | לראשון | . |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

