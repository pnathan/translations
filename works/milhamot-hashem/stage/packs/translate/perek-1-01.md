# Translation pack: milhamot-hashem perek-1-01 (Milchamot ha-Shem 1:1–1:2)

Brief: `briefs/translate.md`. Write: `works/milhamot-hashem/stage/translate/perek-1-01.json`. Check: `go run ./cmd/tr check milhamot-hashem perek-1-01`.

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
- **קצר אפים** (his calque of זעיר אנפין) → "quick to anger", and ארך אפים →
  "slow to anger" everywhere, so the pair shows; a note each time (Aramaic
  "small face"; Prov 14:17). "קצא אפים" as קצר, the spelling in `apparatus`.
- **אדם קדמון** → "Adam Kadmon" (note: "Primordial Man"); **אדם קדמאה** → "Adam
  Kadma'ah": two entities, never merged.
- **כלא חד / כולא חד / ולא אמרינן כלא חד / באגרונו**: one English at every
  occurrence; the sense awaits the owner. The English stands; the reviewer's
  alternative in `readings` / `apparatus`.
- **גשמות / הגשמה / מגשים**: "corporeality"; להגשים → "to ascribe corporeality
  to". Not technical: "לידי התגשמות" → "to realization". "גוף ולא כח בגוף" →
  "a body, or a force in a body".
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
- **מרגל לרגל** → "from festival to festival", not "foot".
- **גופי תורה** → "the essentials of the Torah" (note: "bodies").
- **יש מאין / יש מיש** → "something from nothing / from something".
- **דברי תימה** → "astonishing words".
- **לשבר את האזן** → "to accommodate the ear" (note: "break"); never "shock".
- **דרך ארץ** → "worldly occupation".
- **קורין אל** → "call upon"; **גבורה** as a name of God → "the Power".
- **ואין צריך לומר** → "needless to say".
- **אורה זו תורה** (Esth 8:16) → "'orah' means Torah".
- **סדר** with no function → "Seder" as printed (note); the other in `readings`.
- **רשע בן רשע** → "Wicked one, son of a wicked one"; **רוע לב** → "wickedness of
  heart".
- **ממש** after a comparison: attached to it ("exactly as") unless the syntax
  forbids.
- **אין לנו להחמיר** → "we may not be stringent".
- **הגיע חתור** (infinitive after an adjective of motion): the grammar as
  written; the reviewer's reading may go in `readings`.
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


## Milchamot ha-Shem 1:1

Hebrew (Wikisource): הקדמת המחבר.

| id | word | punctuation after |
|---|---|---|
| 1:1.1 | הקדמת |  |
| 1:1.2 | המחבר | . |


## Milchamot ha-Shem 1:2

Hebrew (Wikisource): יתבונן המשכיל ויתאונן הירא והחרד אל דבר ה', על מצב עם בני ישראל בראותו אשר כל עם ועם מכל מפלגות הגוים מחבבים את לשונם ומגדילים את שפתם אשר חנכו עליה. ואשר הוחקה בה אמונתם ודתם איש ללשונו ואנחנו עם בני ישראל מעת אשר גלינו מארצנו עזבנו את לשונינו הקדושה אשר דבר בה אדם הראשון ובה ניתנה תורתינו הקדושה בהגלות כבודו ית' על הר סיני ואמר לעמו אנכי ה' אלהיך וכו' ולא עשינו כמצווה עלינו ולמדתם אותם את בניכם לדבר בם. וכמו שפי' רש״י ז״ל בפרשת עקב שמכאן אמרו כשהתינוק מתחיל לדבר אביו משיח עמו בלשון הקדש ומלמדו תורה עכ״ל, רק כל שיחתנו עם בנינו בלשון העם אשר אנו יושבים בקרבם ומלמדי תינוקות גם הם רובם לא יכינו לדבר כל צרכם בלשון הקדש אף שיודעים אותו כי לא נסו באלה ואין סדרים ללמוד תינוקות של בית רבן. וידוע הוא ומפורסם חמלת האב על בניו וכמה טרחות ויגיעות ייגעו ההורים בגידול בניהם. עם רוב השקפתם על כל צרכי ילדיהם במאכלו ומשקיו הניאותים לו ומצעו ומלבושו! אך כאשר הגיע חתור להכניסו לבית הספר ישליכהו אצל מלמד תינוקות בבתים צרים וחשוכים על שיירי מחצלאות על הארץ בלי נקיון כי לא ינקה המלמד את מעון התינוקות רק מרגל לרגל אחר. וכל היום יושבים צפופים בלי סדרים ושואפים אויר סרוח ומעופש. ושותים מים מלוחים ובגלל זה כאשר יעברו איזה חדשים מעת שהובא אל בית הספר יעלו בראשו צמחים ושחין רע. ואבותיהם לא ישגיחו עליהם להביא להם מים צלולים ומתוקים לשתות ולא מצע נקי לישב עליו וברוב הימים יושבים על האדמה מאין מחצלת! ובתקופת העת החדשה משנת התר״ן לפרט גדול ליצירה התעוררו רבים מהמשכילים ויבקשו הסתדרות בתי הספר והסתדרות ילדיו לתלמוד תורה ולימוד השפות העברית והערבית והתורקית המתהלכות! וחתרו להושיב מלמדים משכילים בתורת ה' לנהל את ילדי בית הספר בחנוך מוסרי ובנקיון מעון התלמוד בבתים מאירים ובחלונות מרווחים כדי לשאוף רוח צח ונקי! וכל יגיעם של המשכילים ההם היה לריק ולא יצאה מחשבתם לפועל ותינוקות של בית רבן נשארו כשהיו אצל מלמד בחדרים חשוכים ומעופשים יושבים ברבוביא זה לומד אלפא ביתא וזה לומד מקרא וזה לומד עם תרגום אונקלוס ותרגום רס״ג בערבי בלי הבין. גם קטני קטנים שאין ראוים ללמוד כלל כבן שתים ושלש שנים מביאים שם. ופעמים רבות התינוק נפנה שם. והמלמד שולח לאמו שתבוא ותקנח צואת בנה! ולקול הקורא במכתבי עתים בימים ההם מאת ה' אליעזר בן יהודה לתחית השפה העברית הקדושה נכספה נפשינו אך ורק בדברים ושיחות בעלמא ולא באה לידי התגשמות עד שנת תר״ע, אשר אז הואילה ממשלתינו הרוממה ממשלת תוגרמה לפתוח בתי ספר לילדי בני ישראל, ונעשה שם הסתדרות יפה לילדי בני ישראל בעליות מרווחים וחלונות פתוחים ומצעות של צמר, ויריעות עזים וגם שולחנות לילדים גדולים ללמוד ולכתוב עליהם! ואז אמרנו הנה בא העת להוציא לאור עשתונות איש ירושלים ת״ו, אשר בטרם תחיל ילדה האדם הגדול בתחית השפה העברית הוא ניהו אליעזר בן יהודה, ואפריין נמטיה על אשר פעל ועשה ובכל כחו עמל לטובת עם ישראל וישוב ארץ ישראל ותחית השפה העברית ועבודת הקדש הלזו עבד בזריזות וחריצות נפלאה ועשה והצליח להקימה מעפר, ואחרי בלותה היתה לה עדנה, ובגלל אהבתו לשפה הקדושה נתפשטו קמטיה, ותואר פניה הציץ ופרח, ותמצא חן וכבוד בעיני שוחריה יודעיה ומכיריה ויעורו משנתם להחלץ חושים להרחיבה ולהעשירה כבימי נעוריה לפני גלות עם ישראל! ובגלל מצוה זו תהי נפשו צרורה בצרור החיים את ה' אלהינו ומפרי עמלו תשבע נפשו בעדן גן אלהים!

| id | word | punctuation after |
|---|---|---|
| 1:2.1 | יתבונן |  |
| 1:2.2 | המשכיל |  |
| 1:2.3 | ויתאונן |  |
| 1:2.4 | הירא |  |
| 1:2.5 | והחרד |  |
| 1:2.6 | אל |  |
| 1:2.7 | דבר |  |
| 1:2.8 | ה | ', |
| 1:2.9 | על |  |
| 1:2.10 | מצב |  |
| 1:2.11 | עם |  |
| 1:2.12 | בני |  |
| 1:2.13 | ישראל |  |
| 1:2.14 | בראותו |  |
| 1:2.15 | אשר |  |
| 1:2.16 | כל |  |
| 1:2.17 | עם |  |
| 1:2.18 | ועם |  |
| 1:2.19 | מכל |  |
| 1:2.20 | מפלגות |  |
| 1:2.21 | הגוים |  |
| 1:2.22 | מחבבים |  |
| 1:2.23 | את |  |
| 1:2.24 | לשונם |  |
| 1:2.25 | ומגדילים |  |
| 1:2.26 | את |  |
| 1:2.27 | שפתם |  |
| 1:2.28 | אשר |  |
| 1:2.29 | חנכו |  |
| 1:2.30 | עליה | . |
| 1:2.31 | ואשר |  |
| 1:2.32 | הוחקה |  |
| 1:2.33 | בה |  |
| 1:2.34 | אמונתם |  |
| 1:2.35 | ודתם |  |
| 1:2.36 | איש |  |
| 1:2.37 | ללשונו |  |
| 1:2.38 | ואנחנו |  |
| 1:2.39 | עם |  |
| 1:2.40 | בני |  |
| 1:2.41 | ישראל |  |
| 1:2.42 | מעת |  |
| 1:2.43 | אשר |  |
| 1:2.44 | גלינו |  |
| 1:2.45 | מארצנו |  |
| 1:2.46 | עזבנו |  |
| 1:2.47 | את |  |
| 1:2.48 | לשונינו |  |
| 1:2.49 | הקדושה |  |
| 1:2.50 | אשר |  |
| 1:2.51 | דבר |  |
| 1:2.52 | בה |  |
| 1:2.53 | אדם |  |
| 1:2.54 | הראשון |  |
| 1:2.55 | ובה |  |
| 1:2.56 | ניתנה |  |
| 1:2.57 | תורתינו |  |
| 1:2.58 | הקדושה |  |
| 1:2.59 | בהגלות |  |
| 1:2.60 | כבודו |  |
| 1:2.61 | ית | ' |
| 1:2.62 | על |  |
| 1:2.63 | הר |  |
| 1:2.64 | סיני |  |
| 1:2.65 | ואמר |  |
| 1:2.66 | לעמו |  |
| 1:2.67 | אנכי |  |
| 1:2.68 | ה | ' |
| 1:2.69 | אלהיך |  |
| 1:2.70 | וכו | ' |
| 1:2.71 | ולא |  |
| 1:2.72 | עשינו |  |
| 1:2.73 | כמצווה |  |
| 1:2.74 | עלינו |  |
| 1:2.75 | ולמדתם |  |
| 1:2.76 | אותם |  |
| 1:2.77 | את |  |
| 1:2.78 | בניכם |  |
| 1:2.79 | לדבר |  |
| 1:2.80 | בם | . |
| 1:2.81 | וכמו |  |
| 1:2.82 | שפי | ' |
| 1:2.83 | רש״י |  |
| 1:2.84 | ז״ל |  |
| 1:2.85 | בפרשת |  |
| 1:2.86 | עקב |  |
| 1:2.87 | שמכאן |  |
| 1:2.88 | אמרו |  |
| 1:2.89 | כשהתינוק |  |
| 1:2.90 | מתחיל |  |
| 1:2.91 | לדבר |  |
| 1:2.92 | אביו |  |
| 1:2.93 | משיח |  |
| 1:2.94 | עמו |  |
| 1:2.95 | בלשון |  |
| 1:2.96 | הקדש |  |
| 1:2.97 | ומלמדו |  |
| 1:2.98 | תורה |  |
| 1:2.99 | עכ״ל | , |
| 1:2.100 | רק |  |
| 1:2.101 | כל |  |
| 1:2.102 | שיחתנו |  |
| 1:2.103 | עם |  |
| 1:2.104 | בנינו |  |
| 1:2.105 | בלשון |  |
| 1:2.106 | העם |  |
| 1:2.107 | אשר |  |
| 1:2.108 | אנו |  |
| 1:2.109 | יושבים |  |
| 1:2.110 | בקרבם |  |
| 1:2.111 | ומלמדי |  |
| 1:2.112 | תינוקות |  |
| 1:2.113 | גם |  |
| 1:2.114 | הם |  |
| 1:2.115 | רובם |  |
| 1:2.116 | לא |  |
| 1:2.117 | יכינו |  |
| 1:2.118 | לדבר |  |
| 1:2.119 | כל |  |
| 1:2.120 | צרכם |  |
| 1:2.121 | בלשון |  |
| 1:2.122 | הקדש |  |
| 1:2.123 | אף |  |
| 1:2.124 | שיודעים |  |
| 1:2.125 | אותו |  |
| 1:2.126 | כי |  |
| 1:2.127 | לא |  |
| 1:2.128 | נסו |  |
| 1:2.129 | באלה |  |
| 1:2.130 | ואין |  |
| 1:2.131 | סדרים |  |
| 1:2.132 | ללמוד |  |
| 1:2.133 | תינוקות |  |
| 1:2.134 | של |  |
| 1:2.135 | בית |  |
| 1:2.136 | רבן | . |
| 1:2.137 | וידוע |  |
| 1:2.138 | הוא |  |
| 1:2.139 | ומפורסם |  |
| 1:2.140 | חמלת |  |
| 1:2.141 | האב |  |
| 1:2.142 | על |  |
| 1:2.143 | בניו |  |
| 1:2.144 | וכמה |  |
| 1:2.145 | טרחות |  |
| 1:2.146 | ויגיעות |  |
| 1:2.147 | ייגעו |  |
| 1:2.148 | ההורים |  |
| 1:2.149 | בגידול |  |
| 1:2.150 | בניהם | . |
| 1:2.151 | עם |  |
| 1:2.152 | רוב |  |
| 1:2.153 | השקפתם |  |
| 1:2.154 | על |  |
| 1:2.155 | כל |  |
| 1:2.156 | צרכי |  |
| 1:2.157 | ילדיהם |  |
| 1:2.158 | במאכלו |  |
| 1:2.159 | ומשקיו |  |
| 1:2.160 | הניאותים |  |
| 1:2.161 | לו |  |
| 1:2.162 | ומצעו |  |
| 1:2.163 | ומלבושו | ! |
| 1:2.164 | אך |  |
| 1:2.165 | כאשר |  |
| 1:2.166 | הגיע |  |
| 1:2.167 | חתור |  |
| 1:2.168 | להכניסו |  |
| 1:2.169 | לבית |  |
| 1:2.170 | הספר |  |
| 1:2.171 | ישליכהו |  |
| 1:2.172 | אצל |  |
| 1:2.173 | מלמד |  |
| 1:2.174 | תינוקות |  |
| 1:2.175 | בבתים |  |
| 1:2.176 | צרים |  |
| 1:2.177 | וחשוכים |  |
| 1:2.178 | על |  |
| 1:2.179 | שיירי |  |
| 1:2.180 | מחצלאות |  |
| 1:2.181 | על |  |
| 1:2.182 | הארץ |  |
| 1:2.183 | בלי |  |
| 1:2.184 | נקיון |  |
| 1:2.185 | כי |  |
| 1:2.186 | לא |  |
| 1:2.187 | ינקה |  |
| 1:2.188 | המלמד |  |
| 1:2.189 | את |  |
| 1:2.190 | מעון |  |
| 1:2.191 | התינוקות |  |
| 1:2.192 | רק |  |
| 1:2.193 | מרגל |  |
| 1:2.194 | לרגל |  |
| 1:2.195 | אחר | . |
| 1:2.196 | וכל |  |
| 1:2.197 | היום |  |
| 1:2.198 | יושבים |  |
| 1:2.199 | צפופים |  |
| 1:2.200 | בלי |  |
| 1:2.201 | סדרים |  |
| 1:2.202 | ושואפים |  |
| 1:2.203 | אויר |  |
| 1:2.204 | סרוח |  |
| 1:2.205 | ומעופש | . |
| 1:2.206 | ושותים |  |
| 1:2.207 | מים |  |
| 1:2.208 | מלוחים |  |
| 1:2.209 | ובגלל |  |
| 1:2.210 | זה |  |
| 1:2.211 | כאשר |  |
| 1:2.212 | יעברו |  |
| 1:2.213 | איזה |  |
| 1:2.214 | חדשים |  |
| 1:2.215 | מעת |  |
| 1:2.216 | שהובא |  |
| 1:2.217 | אל |  |
| 1:2.218 | בית |  |
| 1:2.219 | הספר |  |
| 1:2.220 | יעלו |  |
| 1:2.221 | בראשו |  |
| 1:2.222 | צמחים |  |
| 1:2.223 | ושחין |  |
| 1:2.224 | רע | . |
| 1:2.225 | ואבותיהם |  |
| 1:2.226 | לא |  |
| 1:2.227 | ישגיחו |  |
| 1:2.228 | עליהם |  |
| 1:2.229 | להביא |  |
| 1:2.230 | להם |  |
| 1:2.231 | מים |  |
| 1:2.232 | צלולים |  |
| 1:2.233 | ומתוקים |  |
| 1:2.234 | לשתות |  |
| 1:2.235 | ולא |  |
| 1:2.236 | מצע |  |
| 1:2.237 | נקי |  |
| 1:2.238 | לישב |  |
| 1:2.239 | עליו |  |
| 1:2.240 | וברוב |  |
| 1:2.241 | הימים |  |
| 1:2.242 | יושבים |  |
| 1:2.243 | על |  |
| 1:2.244 | האדמה |  |
| 1:2.245 | מאין |  |
| 1:2.246 | מחצלת | ! |
| 1:2.247 | ובתקופת |  |
| 1:2.248 | העת |  |
| 1:2.249 | החדשה |  |
| 1:2.250 | משנת |  |
| 1:2.251 | התר״ן |  |
| 1:2.252 | לפרט |  |
| 1:2.253 | גדול |  |
| 1:2.254 | ליצירה |  |
| 1:2.255 | התעוררו |  |
| 1:2.256 | רבים |  |
| 1:2.257 | מהמשכילים |  |
| 1:2.258 | ויבקשו |  |
| 1:2.259 | הסתדרות |  |
| 1:2.260 | בתי |  |
| 1:2.261 | הספר |  |
| 1:2.262 | והסתדרות |  |
| 1:2.263 | ילדיו |  |
| 1:2.264 | לתלמוד |  |
| 1:2.265 | תורה |  |
| 1:2.266 | ולימוד |  |
| 1:2.267 | השפות |  |
| 1:2.268 | העברית |  |
| 1:2.269 | והערבית |  |
| 1:2.270 | והתורקית |  |
| 1:2.271 | המתהלכות | ! |
| 1:2.272 | וחתרו |  |
| 1:2.273 | להושיב |  |
| 1:2.274 | מלמדים |  |
| 1:2.275 | משכילים |  |
| 1:2.276 | בתורת |  |
| 1:2.277 | ה | ' |
| 1:2.278 | לנהל |  |
| 1:2.279 | את |  |
| 1:2.280 | ילדי |  |
| 1:2.281 | בית |  |
| 1:2.282 | הספר |  |
| 1:2.283 | בחנוך |  |
| 1:2.284 | מוסרי |  |
| 1:2.285 | ובנקיון |  |
| 1:2.286 | מעון |  |
| 1:2.287 | התלמוד |  |
| 1:2.288 | בבתים |  |
| 1:2.289 | מאירים |  |
| 1:2.290 | ובחלונות |  |
| 1:2.291 | מרווחים |  |
| 1:2.292 | כדי |  |
| 1:2.293 | לשאוף |  |
| 1:2.294 | רוח |  |
| 1:2.295 | צח |  |
| 1:2.296 | ונקי | ! |
| 1:2.297 | וכל |  |
| 1:2.298 | יגיעם |  |
| 1:2.299 | של |  |
| 1:2.300 | המשכילים |  |
| 1:2.301 | ההם |  |
| 1:2.302 | היה |  |
| 1:2.303 | לריק |  |
| 1:2.304 | ולא |  |
| 1:2.305 | יצאה |  |
| 1:2.306 | מחשבתם |  |
| 1:2.307 | לפועל |  |
| 1:2.308 | ותינוקות |  |
| 1:2.309 | של |  |
| 1:2.310 | בית |  |
| 1:2.311 | רבן |  |
| 1:2.312 | נשארו |  |
| 1:2.313 | כשהיו |  |
| 1:2.314 | אצל |  |
| 1:2.315 | מלמד |  |
| 1:2.316 | בחדרים |  |
| 1:2.317 | חשוכים |  |
| 1:2.318 | ומעופשים |  |
| 1:2.319 | יושבים |  |
| 1:2.320 | ברבוביא |  |
| 1:2.321 | זה |  |
| 1:2.322 | לומד |  |
| 1:2.323 | אלפא |  |
| 1:2.324 | ביתא |  |
| 1:2.325 | וזה |  |
| 1:2.326 | לומד |  |
| 1:2.327 | מקרא |  |
| 1:2.328 | וזה |  |
| 1:2.329 | לומד |  |
| 1:2.330 | עם |  |
| 1:2.331 | תרגום |  |
| 1:2.332 | אונקלוס |  |
| 1:2.333 | ותרגום |  |
| 1:2.334 | רס״ג |  |
| 1:2.335 | בערבי |  |
| 1:2.336 | בלי |  |
| 1:2.337 | הבין | . |
| 1:2.338 | גם |  |
| 1:2.339 | קטני |  |
| 1:2.340 | קטנים |  |
| 1:2.341 | שאין |  |
| 1:2.342 | ראוים |  |
| 1:2.343 | ללמוד |  |
| 1:2.344 | כלל |  |
| 1:2.345 | כבן |  |
| 1:2.346 | שתים |  |
| 1:2.347 | ושלש |  |
| 1:2.348 | שנים |  |
| 1:2.349 | מביאים |  |
| 1:2.350 | שם | . |
| 1:2.351 | ופעמים |  |
| 1:2.352 | רבות |  |
| 1:2.353 | התינוק |  |
| 1:2.354 | נפנה |  |
| 1:2.355 | שם | . |
| 1:2.356 | והמלמד |  |
| 1:2.357 | שולח |  |
| 1:2.358 | לאמו |  |
| 1:2.359 | שתבוא |  |
| 1:2.360 | ותקנח |  |
| 1:2.361 | צואת |  |
| 1:2.362 | בנה | ! |
| 1:2.363 | ולקול |  |
| 1:2.364 | הקורא |  |
| 1:2.365 | במכתבי |  |
| 1:2.366 | עתים |  |
| 1:2.367 | בימים |  |
| 1:2.368 | ההם |  |
| 1:2.369 | מאת |  |
| 1:2.370 | ה | ' |
| 1:2.371 | אליעזר |  |
| 1:2.372 | בן |  |
| 1:2.373 | יהודה |  |
| 1:2.374 | לתחית |  |
| 1:2.375 | השפה |  |
| 1:2.376 | העברית |  |
| 1:2.377 | הקדושה |  |
| 1:2.378 | נכספה |  |
| 1:2.379 | נפשינו |  |
| 1:2.380 | אך |  |
| 1:2.381 | ורק |  |
| 1:2.382 | בדברים |  |
| 1:2.383 | ושיחות |  |
| 1:2.384 | בעלמא |  |
| 1:2.385 | ולא |  |
| 1:2.386 | באה |  |
| 1:2.387 | לידי |  |
| 1:2.388 | התגשמות |  |
| 1:2.389 | עד |  |
| 1:2.390 | שנת |  |
| 1:2.391 | תר״ע | , |
| 1:2.392 | אשר |  |
| 1:2.393 | אז |  |
| 1:2.394 | הואילה |  |
| 1:2.395 | ממשלתינו |  |
| 1:2.396 | הרוממה |  |
| 1:2.397 | ממשלת |  |
| 1:2.398 | תוגרמה |  |
| 1:2.399 | לפתוח |  |
| 1:2.400 | בתי |  |
| 1:2.401 | ספר |  |
| 1:2.402 | לילדי |  |
| 1:2.403 | בני |  |
| 1:2.404 | ישראל | , |
| 1:2.405 | ונעשה |  |
| 1:2.406 | שם |  |
| 1:2.407 | הסתדרות |  |
| 1:2.408 | יפה |  |
| 1:2.409 | לילדי |  |
| 1:2.410 | בני |  |
| 1:2.411 | ישראל |  |
| 1:2.412 | בעליות |  |
| 1:2.413 | מרווחים |  |
| 1:2.414 | וחלונות |  |
| 1:2.415 | פתוחים |  |
| 1:2.416 | ומצעות |  |
| 1:2.417 | של |  |
| 1:2.418 | צמר | , |
| 1:2.419 | ויריעות |  |
| 1:2.420 | עזים |  |
| 1:2.421 | וגם |  |
| 1:2.422 | שולחנות |  |
| 1:2.423 | לילדים |  |
| 1:2.424 | גדולים |  |
| 1:2.425 | ללמוד |  |
| 1:2.426 | ולכתוב |  |
| 1:2.427 | עליהם | ! |
| 1:2.428 | ואז |  |
| 1:2.429 | אמרנו |  |
| 1:2.430 | הנה |  |
| 1:2.431 | בא |  |
| 1:2.432 | העת |  |
| 1:2.433 | להוציא |  |
| 1:2.434 | לאור |  |
| 1:2.435 | עשתונות |  |
| 1:2.436 | איש |  |
| 1:2.437 | ירושלים |  |
| 1:2.438 | ת״ו | , |
| 1:2.439 | אשר |  |
| 1:2.440 | בטרם |  |
| 1:2.441 | תחיל |  |
| 1:2.442 | ילדה |  |
| 1:2.443 | האדם |  |
| 1:2.444 | הגדול |  |
| 1:2.445 | בתחית |  |
| 1:2.446 | השפה |  |
| 1:2.447 | העברית |  |
| 1:2.448 | הוא |  |
| 1:2.449 | ניהו |  |
| 1:2.450 | אליעזר |  |
| 1:2.451 | בן |  |
| 1:2.452 | יהודה | , |
| 1:2.453 | ואפריין |  |
| 1:2.454 | נמטיה |  |
| 1:2.455 | על |  |
| 1:2.456 | אשר |  |
| 1:2.457 | פעל |  |
| 1:2.458 | ועשה |  |
| 1:2.459 | ובכל |  |
| 1:2.460 | כחו |  |
| 1:2.461 | עמל |  |
| 1:2.462 | לטובת |  |
| 1:2.463 | עם |  |
| 1:2.464 | ישראל |  |
| 1:2.465 | וישוב |  |
| 1:2.466 | ארץ |  |
| 1:2.467 | ישראל |  |
| 1:2.468 | ותחית |  |
| 1:2.469 | השפה |  |
| 1:2.470 | העברית |  |
| 1:2.471 | ועבודת |  |
| 1:2.472 | הקדש |  |
| 1:2.473 | הלזו |  |
| 1:2.474 | עבד |  |
| 1:2.475 | בזריזות |  |
| 1:2.476 | וחריצות |  |
| 1:2.477 | נפלאה |  |
| 1:2.478 | ועשה |  |
| 1:2.479 | והצליח |  |
| 1:2.480 | להקימה |  |
| 1:2.481 | מעפר | , |
| 1:2.482 | ואחרי |  |
| 1:2.483 | בלותה |  |
| 1:2.484 | היתה |  |
| 1:2.485 | לה |  |
| 1:2.486 | עדנה | , |
| 1:2.487 | ובגלל |  |
| 1:2.488 | אהבתו |  |
| 1:2.489 | לשפה |  |
| 1:2.490 | הקדושה |  |
| 1:2.491 | נתפשטו |  |
| 1:2.492 | קמטיה | , |
| 1:2.493 | ותואר |  |
| 1:2.494 | פניה |  |
| 1:2.495 | הציץ |  |
| 1:2.496 | ופרח | , |
| 1:2.497 | ותמצא |  |
| 1:2.498 | חן |  |
| 1:2.499 | וכבוד |  |
| 1:2.500 | בעיני |  |
| 1:2.501 | שוחריה |  |
| 1:2.502 | יודעיה |  |
| 1:2.503 | ומכיריה |  |
| 1:2.504 | ויעורו |  |
| 1:2.505 | משנתם |  |
| 1:2.506 | להחלץ |  |
| 1:2.507 | חושים |  |
| 1:2.508 | להרחיבה |  |
| 1:2.509 | ולהעשירה |  |
| 1:2.510 | כבימי |  |
| 1:2.511 | נעוריה |  |
| 1:2.512 | לפני |  |
| 1:2.513 | גלות |  |
| 1:2.514 | עם |  |
| 1:2.515 | ישראל | ! |
| 1:2.516 | ובגלל |  |
| 1:2.517 | מצוה |  |
| 1:2.518 | זו |  |
| 1:2.519 | תהי |  |
| 1:2.520 | נפשו |  |
| 1:2.521 | צרורה |  |
| 1:2.522 | בצרור |  |
| 1:2.523 | החיים |  |
| 1:2.524 | את |  |
| 1:2.525 | ה | ' |
| 1:2.526 | אלהינו |  |
| 1:2.527 | ומפרי |  |
| 1:2.528 | עמלו |  |
| 1:2.529 | תשבע |  |
| 1:2.530 | נפשו |  |
| 1:2.531 | בעדן |  |
| 1:2.532 | גן |  |
| 1:2.533 | אלהים | ! |


## Lexicon

Senses are cited by entry key and path. BDB entries list senses by stem (binyan).

