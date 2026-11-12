# Translation pack: milhamot-hashem perek-5-16 (Milchamot ha-Shem 5:24–5:24)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-5-16.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-5-16`.

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
- **אבות** → "fathers", lower case (note: "patriarchs"); never a title.
- **שם** → "there" or "the name" by context, the other in `apparatus`.
- **תינוקות של בית רבן** → "schoolchildren" (note: "the house of the master").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **והוא שלא** → "provided that".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **לדעת** + noun → "according to the understanding of" without an infinitive
  sense; the infinitive in `readings`.
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **אחד ... ואחד** of two kinds of case → "alike".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
- **דין / אם הלכה נקבל ואם לדין יש תשובה**: דין → "reasoning"; בקש → "in haste"
  (note: the conjecture "with straw").
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


## Milchamot ha-Shem 5:24

Hebrew (Wikisource): נ) וכ"כ הרמב"ם פרק א' מהל' ע"ז וז"ל כיון שנגמל איתן זה התחיל לשוטט בדעתו והוא קטן ולחשב ביום ובלילה, והיה תמיה היאך אפשר שיהיה הגלגל הזה נוהג תמיד ואין לו מנהיג, ומי יסבב אותו לפי שאי אפשר שיסבב את עצמו. ולא היה לו מלמד ולא מודיע לו דבר אלא מושקע באור כשדים בין עובדי עבודה זרה הטפשים. ואביו ואמו וכל העם עובדים ע"ז והוא היה עובד עמהם ולבו משוטט ומבין עד שהשיג דרך האמת והבין קו הצדק מדעתו הנכונה והבין וידע שיש שם אלוה אחד והוא מנהיג את הגלגל. והוא האל בורא הכל ואין בכל הנמצאים אלוה חוץ ממנו, וידע שכל העם טועים, ודבר שגרם להם טעות זו שעובדים את הכוכבים ואת הצורות עד שאבד האמת מדעתן. ובן ארבעים שנה הכיר אברהם את בוראו כיון שהכיר וידע התחיל להשיב תשובות על בני אור כשדים ולערוך דין עמהם ולומר שאין זו דרך האמת שאתם הולכים בה וכו' וכיון שהיו העם מתקבצין אליו היה מודיע לכל אחד ואחד כפי דעתו עד שיחזירוהו למוטב לדרך האמת עד שנתקבצו אליו אלפים ורבבות והם אנשי בית אברהם ושתל בלבן העיקר הגדול הזה וחיבר בו ספרים והודיע ליצחק בנו ויצחק ישב ללמד ומחזיר ומזהיר, ויצחק ליעקב ומינהו ללמד ומחזיר כל הנלוים אליו. ויעקב אבינו לימד בניו כולן, והבדיל לוי ומינהו ראש והושיבו בישיבה ללמד דרך ה' ולשמור מצות אברהם, וצוה את בניו שלא יפסיקו מבני לוי ממונה כדי שלא ישתכח הלמוד, והיה הדבר הולך ומתגבר בבני יעקב והנלוים אליהם ונעשית בעולם אומה שהיא יודעת את השם עד שארכו הימים לישראל במצרים, וחזרו ללמוד מעשיהם ולעבוד ע"ז כמותן חוץ משבט לוי שעמד במצות אבותיו ומעולם לא עבד שבט לוי ע"ז עכ"ל. ובהל' קריאת שמע כתוב כשקרבו ימי ישראל למות קבץ בניו וצוום וזריזם על יחוד ה' ועל דרך ה' שהלך בה אברהם ויצחק אביו, ואמר להם בני שמא יש בכם פסלות מי שאינו עומד עמו ביחוד אדון העולם ענו כולם ואמרו שמע ישראל ה' אלהינו ה' אחד (כלומר שמע ממנו אבינו ישראל ה' אלהינו ה' אחד) פתח הזקן ואמר ברוך שם כבוד מלכותו לעולם ועד, לפיכך וכו' ע"כ. דבר הלמד מכל האמור שאברהם אבינו ע"ה לימד לכל העולם אמונת יחוד השם וחיבר בה ספרים והודיעם לבניו ולכל הנמשכים אחריו וישמעאל בנו בכללם דהא קי"ל שישמעאל עשה תשובה. ואפשר שבשכר שעשה ישמעאל תשובה זכו בניו לחזור לאמונת יחוד ה' אחרי ימים רבים על ידי מחמד שלמד אצל ישראל כידוע, אך בזרע יעקב לא פסקה, ונשארה בזרע לוי ובשרידים משאר שבטים עד שבא מחוקקינו אבי כל הנביאים וראש לחכמים משה רבינו ע"ה ונגלו לו מלאך ה' בסנה בשמרו את השבועה שנשבע לאברהם אבינו וגאל את אבותינו ממצרים על ידו, וקירבנו לפני הר סיני, והכתיר אותנו במצות והודיענו דרך עבודתו ויחודו ומה יהיה משפט ע"ז והטועים אחריה. וכאשר קרב עתו של משה רבינו ע"ה להפטר הודיעו הקב"ה מה יהיה לנו באחרית הימים בגלותנו, ואמר לו, הנך שוכב עם אבותיך וקם העם הזה וזנה אחרי אלהי נכר הארץ אשר הוא בא שמה בקרבו וכו'. וכך אירע לנו שהטעו אותנו לסור מאחרי ה' אלהינו ולעבוד אלהים אחרים, ולהאמין באלוהות רבים על ידי ספר הזהר הנמסר לנו מיד מלך ממלכי העכו"ם היושבים במזרח, ויחסו הספר הלזה לרשב"י התנא בשקר ומן הוא והלאה נשתנית אמונתינו מעט מעט עד שנהפכה לזרא. ורוב חכמי ישראל לא נתנו לב לדעת מקורו ומבטן מי יצאו הדברים רק האמינ ולמשאות שוא ומדוחים הממונים בחובו, כאלו מרשב"י התנא הקדוש יצאו הדברים, ולכן לא שמו אל לבם למרורת הפתנים אשר בקרבו כי לקחם בחלק שפתיו והדיחם, ונעשית הקבלה המזוייפת לגבירה! והקבלה האמתית המסורה לנו איש מפי איש ממשה רבי' עד רבי ועד רב אשי ורבינא ור' יוחנן מרה דארעא דישראל נעשית לשפחה בזויה ומאוסה ח"ו ותחתיה רגזה ארץ כמבואר בזהר ויובא לקמן (סי' ס"ז וז"ח).

| id | word | punctuation after |
|---|---|---|
| 5:24.1 | נ | ) |
| 5:24.2 | וכ"כ |  |
| 5:24.3 | הרמב"ם |  |
| 5:24.4 | פרק |  |
| 5:24.5 | א | ' |
| 5:24.6 | מהל | ' |
| 5:24.7 | ע"ז |  |
| 5:24.8 | וז"ל |  |
| 5:24.9 | כיון |  |
| 5:24.10 | שנגמל |  |
| 5:24.11 | איתן |  |
| 5:24.12 | זה |  |
| 5:24.13 | התחיל |  |
| 5:24.14 | לשוטט |  |
| 5:24.15 | בדעתו |  |
| 5:24.16 | והוא |  |
| 5:24.17 | קטן |  |
| 5:24.18 | ולחשב |  |
| 5:24.19 | ביום |  |
| 5:24.20 | ובלילה | , |
| 5:24.21 | והיה |  |
| 5:24.22 | תמיה |  |
| 5:24.23 | היאך |  |
| 5:24.24 | אפשר |  |
| 5:24.25 | שיהיה |  |
| 5:24.26 | הגלגל |  |
| 5:24.27 | הזה |  |
| 5:24.28 | נוהג |  |
| 5:24.29 | תמיד |  |
| 5:24.30 | ואין |  |
| 5:24.31 | לו |  |
| 5:24.32 | מנהיג | , |
| 5:24.33 | ומי |  |
| 5:24.34 | יסבב |  |
| 5:24.35 | אותו |  |
| 5:24.36 | לפי |  |
| 5:24.37 | שאי |  |
| 5:24.38 | אפשר |  |
| 5:24.39 | שיסבב |  |
| 5:24.40 | את |  |
| 5:24.41 | עצמו | . |
| 5:24.42 | ולא |  |
| 5:24.43 | היה |  |
| 5:24.44 | לו |  |
| 5:24.45 | מלמד |  |
| 5:24.46 | ולא |  |
| 5:24.47 | מודיע |  |
| 5:24.48 | לו |  |
| 5:24.49 | דבר |  |
| 5:24.50 | אלא |  |
| 5:24.51 | מושקע |  |
| 5:24.52 | באור |  |
| 5:24.53 | כשדים |  |
| 5:24.54 | בין |  |
| 5:24.55 | עובדי |  |
| 5:24.56 | עבודה |  |
| 5:24.57 | זרה |  |
| 5:24.58 | הטפשים | . |
| 5:24.59 | ואביו |  |
| 5:24.60 | ואמו |  |
| 5:24.61 | וכל |  |
| 5:24.62 | העם |  |
| 5:24.63 | עובדים |  |
| 5:24.64 | ע"ז |  |
| 5:24.65 | והוא |  |
| 5:24.66 | היה |  |
| 5:24.67 | עובד |  |
| 5:24.68 | עמהם |  |
| 5:24.69 | ולבו |  |
| 5:24.70 | משוטט |  |
| 5:24.71 | ומבין |  |
| 5:24.72 | עד |  |
| 5:24.73 | שהשיג |  |
| 5:24.74 | דרך |  |
| 5:24.75 | האמת |  |
| 5:24.76 | והבין |  |
| 5:24.77 | קו |  |
| 5:24.78 | הצדק |  |
| 5:24.79 | מדעתו |  |
| 5:24.80 | הנכונה |  |
| 5:24.81 | והבין |  |
| 5:24.82 | וידע |  |
| 5:24.83 | שיש |  |
| 5:24.84 | שם |  |
| 5:24.85 | אלוה |  |
| 5:24.86 | אחד |  |
| 5:24.87 | והוא |  |
| 5:24.88 | מנהיג |  |
| 5:24.89 | את |  |
| 5:24.90 | הגלגל | . |
| 5:24.91 | והוא |  |
| 5:24.92 | האל |  |
| 5:24.93 | בורא |  |
| 5:24.94 | הכל |  |
| 5:24.95 | ואין |  |
| 5:24.96 | בכל |  |
| 5:24.97 | הנמצאים |  |
| 5:24.98 | אלוה |  |
| 5:24.99 | חוץ |  |
| 5:24.100 | ממנו | , |
| 5:24.101 | וידע |  |
| 5:24.102 | שכל |  |
| 5:24.103 | העם |  |
| 5:24.104 | טועים | , |
| 5:24.105 | ודבר |  |
| 5:24.106 | שגרם |  |
| 5:24.107 | להם |  |
| 5:24.108 | טעות |  |
| 5:24.109 | זו |  |
| 5:24.110 | שעובדים |  |
| 5:24.111 | את |  |
| 5:24.112 | הכוכבים |  |
| 5:24.113 | ואת |  |
| 5:24.114 | הצורות |  |
| 5:24.115 | עד |  |
| 5:24.116 | שאבד |  |
| 5:24.117 | האמת |  |
| 5:24.118 | מדעתן | . |
| 5:24.119 | ובן |  |
| 5:24.120 | ארבעים |  |
| 5:24.121 | שנה |  |
| 5:24.122 | הכיר |  |
| 5:24.123 | אברהם |  |
| 5:24.124 | את |  |
| 5:24.125 | בוראו |  |
| 5:24.126 | כיון |  |
| 5:24.127 | שהכיר |  |
| 5:24.128 | וידע |  |
| 5:24.129 | התחיל |  |
| 5:24.130 | להשיב |  |
| 5:24.131 | תשובות |  |
| 5:24.132 | על |  |
| 5:24.133 | בני |  |
| 5:24.134 | אור |  |
| 5:24.135 | כשדים |  |
| 5:24.136 | ולערוך |  |
| 5:24.137 | דין |  |
| 5:24.138 | עמהם |  |
| 5:24.139 | ולומר |  |
| 5:24.140 | שאין |  |
| 5:24.141 | זו |  |
| 5:24.142 | דרך |  |
| 5:24.143 | האמת |  |
| 5:24.144 | שאתם |  |
| 5:24.145 | הולכים |  |
| 5:24.146 | בה |  |
| 5:24.147 | וכו | ' |
| 5:24.148 | וכיון |  |
| 5:24.149 | שהיו |  |
| 5:24.150 | העם |  |
| 5:24.151 | מתקבצין |  |
| 5:24.152 | אליו |  |
| 5:24.153 | היה |  |
| 5:24.154 | מודיע |  |
| 5:24.155 | לכל |  |
| 5:24.156 | אחד |  |
| 5:24.157 | ואחד |  |
| 5:24.158 | כפי |  |
| 5:24.159 | דעתו |  |
| 5:24.160 | עד |  |
| 5:24.161 | שיחזירוהו |  |
| 5:24.162 | למוטב |  |
| 5:24.163 | לדרך |  |
| 5:24.164 | האמת |  |
| 5:24.165 | עד |  |
| 5:24.166 | שנתקבצו |  |
| 5:24.167 | אליו |  |
| 5:24.168 | אלפים |  |
| 5:24.169 | ורבבות |  |
| 5:24.170 | והם |  |
| 5:24.171 | אנשי |  |
| 5:24.172 | בית |  |
| 5:24.173 | אברהם |  |
| 5:24.174 | ושתל |  |
| 5:24.175 | בלבן |  |
| 5:24.176 | העיקר |  |
| 5:24.177 | הגדול |  |
| 5:24.178 | הזה |  |
| 5:24.179 | וחיבר |  |
| 5:24.180 | בו |  |
| 5:24.181 | ספרים |  |
| 5:24.182 | והודיע |  |
| 5:24.183 | ליצחק |  |
| 5:24.184 | בנו |  |
| 5:24.185 | ויצחק |  |
| 5:24.186 | ישב |  |
| 5:24.187 | ללמד |  |
| 5:24.188 | ומחזיר |  |
| 5:24.189 | ומזהיר | , |
| 5:24.190 | ויצחק |  |
| 5:24.191 | ליעקב |  |
| 5:24.192 | ומינהו |  |
| 5:24.193 | ללמד |  |
| 5:24.194 | ומחזיר |  |
| 5:24.195 | כל |  |
| 5:24.196 | הנלוים |  |
| 5:24.197 | אליו | . |
| 5:24.198 | ויעקב |  |
| 5:24.199 | אבינו |  |
| 5:24.200 | לימד |  |
| 5:24.201 | בניו |  |
| 5:24.202 | כולן | , |
| 5:24.203 | והבדיל |  |
| 5:24.204 | לוי |  |
| 5:24.205 | ומינהו |  |
| 5:24.206 | ראש |  |
| 5:24.207 | והושיבו |  |
| 5:24.208 | בישיבה |  |
| 5:24.209 | ללמד |  |
| 5:24.210 | דרך |  |
| 5:24.211 | ה | ' |
| 5:24.212 | ולשמור |  |
| 5:24.213 | מצות |  |
| 5:24.214 | אברהם | , |
| 5:24.215 | וצוה |  |
| 5:24.216 | את |  |
| 5:24.217 | בניו |  |
| 5:24.218 | שלא |  |
| 5:24.219 | יפסיקו |  |
| 5:24.220 | מבני |  |
| 5:24.221 | לוי |  |
| 5:24.222 | ממונה |  |
| 5:24.223 | כדי |  |
| 5:24.224 | שלא |  |
| 5:24.225 | ישתכח |  |
| 5:24.226 | הלמוד | , |
| 5:24.227 | והיה |  |
| 5:24.228 | הדבר |  |
| 5:24.229 | הולך |  |
| 5:24.230 | ומתגבר |  |
| 5:24.231 | בבני |  |
| 5:24.232 | יעקב |  |
| 5:24.233 | והנלוים |  |
| 5:24.234 | אליהם |  |
| 5:24.235 | ונעשית |  |
| 5:24.236 | בעולם |  |
| 5:24.237 | אומה |  |
| 5:24.238 | שהיא |  |
| 5:24.239 | יודעת |  |
| 5:24.240 | את |  |
| 5:24.241 | השם |  |
| 5:24.242 | עד |  |
| 5:24.243 | שארכו |  |
| 5:24.244 | הימים |  |
| 5:24.245 | לישראל |  |
| 5:24.246 | במצרים | , |
| 5:24.247 | וחזרו |  |
| 5:24.248 | ללמוד |  |
| 5:24.249 | מעשיהם |  |
| 5:24.250 | ולעבוד |  |
| 5:24.251 | ע"ז |  |
| 5:24.252 | כמותן |  |
| 5:24.253 | חוץ |  |
| 5:24.254 | משבט |  |
| 5:24.255 | לוי |  |
| 5:24.256 | שעמד |  |
| 5:24.257 | במצות |  |
| 5:24.258 | אבותיו |  |
| 5:24.259 | ומעולם |  |
| 5:24.260 | לא |  |
| 5:24.261 | עבד |  |
| 5:24.262 | שבט |  |
| 5:24.263 | לוי |  |
| 5:24.264 | ע"ז |  |
| 5:24.265 | עכ"ל | . |
| 5:24.266 | ובהל | ' |
| 5:24.267 | קריאת |  |
| 5:24.268 | שמע |  |
| 5:24.269 | כתוב |  |
| 5:24.270 | כשקרבו |  |
| 5:24.271 | ימי |  |
| 5:24.272 | ישראל |  |
| 5:24.273 | למות |  |
| 5:24.274 | קבץ |  |
| 5:24.275 | בניו |  |
| 5:24.276 | וצוום |  |
| 5:24.277 | וזריזם |  |
| 5:24.278 | על |  |
| 5:24.279 | יחוד |  |
| 5:24.280 | ה | ' |
| 5:24.281 | ועל |  |
| 5:24.282 | דרך |  |
| 5:24.283 | ה | ' |
| 5:24.284 | שהלך |  |
| 5:24.285 | בה |  |
| 5:24.286 | אברהם |  |
| 5:24.287 | ויצחק |  |
| 5:24.288 | אביו | , |
| 5:24.289 | ואמר |  |
| 5:24.290 | להם |  |
| 5:24.291 | בני |  |
| 5:24.292 | שמא |  |
| 5:24.293 | יש |  |
| 5:24.294 | בכם |  |
| 5:24.295 | פסלות |  |
| 5:24.296 | מי |  |
| 5:24.297 | שאינו |  |
| 5:24.298 | עומד |  |
| 5:24.299 | עמו |  |
| 5:24.300 | ביחוד |  |
| 5:24.301 | אדון |  |
| 5:24.302 | העולם |  |
| 5:24.303 | ענו |  |
| 5:24.304 | כולם |  |
| 5:24.305 | ואמרו |  |
| 5:24.306 | שמע |  |
| 5:24.307 | ישראל |  |
| 5:24.308 | ה | ' |
| 5:24.309 | אלהינו |  |
| 5:24.310 | ה | ' |
| 5:24.311 | אחד |  |
| 5:24.312 | כלומר |  |
| 5:24.313 | שמע |  |
| 5:24.314 | ממנו |  |
| 5:24.315 | אבינו |  |
| 5:24.316 | ישראל |  |
| 5:24.317 | ה | ' |
| 5:24.318 | אלהינו |  |
| 5:24.319 | ה | ' |
| 5:24.320 | אחד | ) |
| 5:24.321 | פתח |  |
| 5:24.322 | הזקן |  |
| 5:24.323 | ואמר |  |
| 5:24.324 | ברוך |  |
| 5:24.325 | שם |  |
| 5:24.326 | כבוד |  |
| 5:24.327 | מלכותו |  |
| 5:24.328 | לעולם |  |
| 5:24.329 | ועד | , |
| 5:24.330 | לפיכך |  |
| 5:24.331 | וכו | ' |
| 5:24.332 | ע"כ | . |
| 5:24.333 | דבר |  |
| 5:24.334 | הלמד |  |
| 5:24.335 | מכל |  |
| 5:24.336 | האמור |  |
| 5:24.337 | שאברהם |  |
| 5:24.338 | אבינו |  |
| 5:24.339 | ע"ה |  |
| 5:24.340 | לימד |  |
| 5:24.341 | לכל |  |
| 5:24.342 | העולם |  |
| 5:24.343 | אמונת |  |
| 5:24.344 | יחוד |  |
| 5:24.345 | השם |  |
| 5:24.346 | וחיבר |  |
| 5:24.347 | בה |  |
| 5:24.348 | ספרים |  |
| 5:24.349 | והודיעם |  |
| 5:24.350 | לבניו |  |
| 5:24.351 | ולכל |  |
| 5:24.352 | הנמשכים |  |
| 5:24.353 | אחריו |  |
| 5:24.354 | וישמעאל |  |
| 5:24.355 | בנו |  |
| 5:24.356 | בכללם |  |
| 5:24.357 | דהא |  |
| 5:24.358 | קי"ל |  |
| 5:24.359 | שישמעאל |  |
| 5:24.360 | עשה |  |
| 5:24.361 | תשובה | . |
| 5:24.362 | ואפשר |  |
| 5:24.363 | שבשכר |  |
| 5:24.364 | שעשה |  |
| 5:24.365 | ישמעאל |  |
| 5:24.366 | תשובה |  |
| 5:24.367 | זכו |  |
| 5:24.368 | בניו |  |
| 5:24.369 | לחזור |  |
| 5:24.370 | לאמונת |  |
| 5:24.371 | יחוד |  |
| 5:24.372 | ה | ' |
| 5:24.373 | אחרי |  |
| 5:24.374 | ימים |  |
| 5:24.375 | רבים |  |
| 5:24.376 | על |  |
| 5:24.377 | ידי |  |
| 5:24.378 | מחמד |  |
| 5:24.379 | שלמד |  |
| 5:24.380 | אצל |  |
| 5:24.381 | ישראל |  |
| 5:24.382 | כידוע | , |
| 5:24.383 | אך |  |
| 5:24.384 | בזרע |  |
| 5:24.385 | יעקב |  |
| 5:24.386 | לא |  |
| 5:24.387 | פסקה | , |
| 5:24.388 | ונשארה |  |
| 5:24.389 | בזרע |  |
| 5:24.390 | לוי |  |
| 5:24.391 | ובשרידים |  |
| 5:24.392 | משאר |  |
| 5:24.393 | שבטים |  |
| 5:24.394 | עד |  |
| 5:24.395 | שבא |  |
| 5:24.396 | מחוקקינו |  |
| 5:24.397 | אבי |  |
| 5:24.398 | כל |  |
| 5:24.399 | הנביאים |  |
| 5:24.400 | וראש |  |
| 5:24.401 | לחכמים |  |
| 5:24.402 | משה |  |
| 5:24.403 | רבינו |  |
| 5:24.404 | ע"ה |  |
| 5:24.405 | ונגלו |  |
| 5:24.406 | לו |  |
| 5:24.407 | מלאך |  |
| 5:24.408 | ה | ' |
| 5:24.409 | בסנה |  |
| 5:24.410 | בשמרו |  |
| 5:24.411 | את |  |
| 5:24.412 | השבועה |  |
| 5:24.413 | שנשבע |  |
| 5:24.414 | לאברהם |  |
| 5:24.415 | אבינו |  |
| 5:24.416 | וגאל |  |
| 5:24.417 | את |  |
| 5:24.418 | אבותינו |  |
| 5:24.419 | ממצרים |  |
| 5:24.420 | על |  |
| 5:24.421 | ידו | , |
| 5:24.422 | וקירבנו |  |
| 5:24.423 | לפני |  |
| 5:24.424 | הר |  |
| 5:24.425 | סיני | , |
| 5:24.426 | והכתיר |  |
| 5:24.427 | אותנו |  |
| 5:24.428 | במצות |  |
| 5:24.429 | והודיענו |  |
| 5:24.430 | דרך |  |
| 5:24.431 | עבודתו |  |
| 5:24.432 | ויחודו |  |
| 5:24.433 | ומה |  |
| 5:24.434 | יהיה |  |
| 5:24.435 | משפט |  |
| 5:24.436 | ע"ז |  |
| 5:24.437 | והטועים |  |
| 5:24.438 | אחריה | . |
| 5:24.439 | וכאשר |  |
| 5:24.440 | קרב |  |
| 5:24.441 | עתו |  |
| 5:24.442 | של |  |
| 5:24.443 | משה |  |
| 5:24.444 | רבינו |  |
| 5:24.445 | ע"ה |  |
| 5:24.446 | להפטר |  |
| 5:24.447 | הודיעו |  |
| 5:24.448 | הקב"ה |  |
| 5:24.449 | מה |  |
| 5:24.450 | יהיה |  |
| 5:24.451 | לנו |  |
| 5:24.452 | באחרית |  |
| 5:24.453 | הימים |  |
| 5:24.454 | בגלותנו | , |
| 5:24.455 | ואמר |  |
| 5:24.456 | לו | , |
| 5:24.457 | הנך |  |
| 5:24.458 | שוכב |  |
| 5:24.459 | עם |  |
| 5:24.460 | אבותיך |  |
| 5:24.461 | וקם |  |
| 5:24.462 | העם |  |
| 5:24.463 | הזה |  |
| 5:24.464 | וזנה |  |
| 5:24.465 | אחרי |  |
| 5:24.466 | אלהי |  |
| 5:24.467 | נכר |  |
| 5:24.468 | הארץ |  |
| 5:24.469 | אשר |  |
| 5:24.470 | הוא |  |
| 5:24.471 | בא |  |
| 5:24.472 | שמה |  |
| 5:24.473 | בקרבו |  |
| 5:24.474 | וכו | '. |
| 5:24.475 | וכך |  |
| 5:24.476 | אירע |  |
| 5:24.477 | לנו |  |
| 5:24.478 | שהטעו |  |
| 5:24.479 | אותנו |  |
| 5:24.480 | לסור |  |
| 5:24.481 | מאחרי |  |
| 5:24.482 | ה | ' |
| 5:24.483 | אלהינו |  |
| 5:24.484 | ולעבוד |  |
| 5:24.485 | אלהים |  |
| 5:24.486 | אחרים | , |
| 5:24.487 | ולהאמין |  |
| 5:24.488 | באלוהות |  |
| 5:24.489 | רבים |  |
| 5:24.490 | על |  |
| 5:24.491 | ידי |  |
| 5:24.492 | ספר |  |
| 5:24.493 | הזהר |  |
| 5:24.494 | הנמסר |  |
| 5:24.495 | לנו |  |
| 5:24.496 | מיד |  |
| 5:24.497 | מלך |  |
| 5:24.498 | ממלכי |  |
| 5:24.499 | העכו"ם |  |
| 5:24.500 | היושבים |  |
| 5:24.501 | במזרח | , |
| 5:24.502 | ויחסו |  |
| 5:24.503 | הספר |  |
| 5:24.504 | הלזה |  |
| 5:24.505 | לרשב"י |  |
| 5:24.506 | התנא |  |
| 5:24.507 | בשקר |  |
| 5:24.508 | ומן |  |
| 5:24.509 | הוא |  |
| 5:24.510 | והלאה |  |
| 5:24.511 | נשתנית |  |
| 5:24.512 | אמונתינו |  |
| 5:24.513 | מעט |  |
| 5:24.514 | מעט |  |
| 5:24.515 | עד |  |
| 5:24.516 | שנהפכה |  |
| 5:24.517 | לזרא | . |
| 5:24.518 | ורוב |  |
| 5:24.519 | חכמי |  |
| 5:24.520 | ישראל |  |
| 5:24.521 | לא |  |
| 5:24.522 | נתנו |  |
| 5:24.523 | לב |  |
| 5:24.524 | לדעת |  |
| 5:24.525 | מקורו |  |
| 5:24.526 | ומבטן |  |
| 5:24.527 | מי |  |
| 5:24.528 | יצאו |  |
| 5:24.529 | הדברים |  |
| 5:24.530 | רק |  |
| 5:24.531 | האמינ |  |
| 5:24.532 | ולמשאות |  |
| 5:24.533 | שוא |  |
| 5:24.534 | ומדוחים |  |
| 5:24.535 | הממונים |  |
| 5:24.536 | בחובו | , |
| 5:24.537 | כאלו |  |
| 5:24.538 | מרשב"י |  |
| 5:24.539 | התנא |  |
| 5:24.540 | הקדוש |  |
| 5:24.541 | יצאו |  |
| 5:24.542 | הדברים | , |
| 5:24.543 | ולכן |  |
| 5:24.544 | לא |  |
| 5:24.545 | שמו |  |
| 5:24.546 | אל |  |
| 5:24.547 | לבם |  |
| 5:24.548 | למרורת |  |
| 5:24.549 | הפתנים |  |
| 5:24.550 | אשר |  |
| 5:24.551 | בקרבו |  |
| 5:24.552 | כי |  |
| 5:24.553 | לקחם |  |
| 5:24.554 | בחלק |  |
| 5:24.555 | שפתיו |  |
| 5:24.556 | והדיחם | , |
| 5:24.557 | ונעשית |  |
| 5:24.558 | הקבלה |  |
| 5:24.559 | המזוייפת |  |
| 5:24.560 | לגבירה | ! |
| 5:24.561 | והקבלה |  |
| 5:24.562 | האמתית |  |
| 5:24.563 | המסורה |  |
| 5:24.564 | לנו |  |
| 5:24.565 | איש |  |
| 5:24.566 | מפי |  |
| 5:24.567 | איש |  |
| 5:24.568 | ממשה |  |
| 5:24.569 | רבי | ' |
| 5:24.570 | עד |  |
| 5:24.571 | רבי |  |
| 5:24.572 | ועד |  |
| 5:24.573 | רב |  |
| 5:24.574 | אשי |  |
| 5:24.575 | ורבינא |  |
| 5:24.576 | ור | ' |
| 5:24.577 | יוחנן |  |
| 5:24.578 | מרה |  |
| 5:24.579 | דארעא |  |
| 5:24.580 | דישראל |  |
| 5:24.581 | נעשית |  |
| 5:24.582 | לשפחה |  |
| 5:24.583 | בזויה |  |
| 5:24.584 | ומאוסה |  |
| 5:24.585 | ח"ו |  |
| 5:24.586 | ותחתיה |  |
| 5:24.587 | רגזה |  |
| 5:24.588 | ארץ |  |
| 5:24.589 | כמבואר |  |
| 5:24.590 | בזהר |  |
| 5:24.591 | ויובא |  |
| 5:24.592 | לקמן |  |
| 5:24.593 | סי | ' |
| 5:24.594 | ס"ז |  |
| 5:24.595 | וז"ח | ). |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

