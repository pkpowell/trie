package trie

var (
	voluspa = `
	Vǫluspá
‘Hljóðs bið ek allar kindir,
meiri ok minni, mǫgu Heimdallar! Vildu at ek, Valfǫðr, vel fyrtelja forn spjǫll fira, þau er fremst um man.
‘Ek man jǫtna, ár um borna, þá er forðum mik fœdda hǫfðu; níu man ek heima, níu íviðjur, mjǫtvið mæran, fyr mold neðan.
‘Ár var alda, þar er Ymir bygði; vara sandr né sær né svalar unnir; jǫrð fannsk æva né upphiminn, gap var ginnunga, en gras hvergi.
‘Áðr Burs synir bjǫðum um ypðu, þeir er Miðgarð mæran skópu; Sól skein sunnan á salar steina, þá var grund gróin grœnum lauki.
‘Sól varp sunnan, sinni Mána, hendi inni hœgri um himinjódýr;
Sól þat né vissi hvar hon sali átti, stjǫrnur þat né vissu hvar þær staði áttu, Máni þat né vissi hvat hann megins átti.
‘Þá gengu regin ǫll á rǫkstóla, ginnheilǫg goð, ok um þat gættusk: nótt ok niðjum nǫfn um gáfu, morgin hétu ok miðjan dag, undorn ok aptan, árum at telja.
‘Hittusk Æsir á Iðavelli, þeir er hǫrg ok hof hátimbruðu; afla lǫgðu, auð smíðuðu, tangir skópu ok tól gørðu.
‘Teflðu í túni, teitir váru,
var þeim vættergis vant ór gulli, unz þrjár kvómu þursa meyjar, ámátkar mjǫk, ór Jǫtunheimum.
‘Þá gengu regin ǫll á rǫkstóla, ginnheilǫg goð, ok um þat gættusk, hverr skyldi dverga dróttin skepja ór Brimis blóði ok ór blám leggjum.
‘Þar var Mótsognir mæztr um orðinn dverga allra, en Durinn annarr;
þeir manlíkun dvergar, ór jǫrðu,
‘Nýi ok Niði, Austri ok Vestri, Bívǫrr, Bávǫrr,
Án ok Ánarr,
mǫrg um gørðu, sem Durinn sagði.
Norðri ok Suðri, Alþjófr, Dvalinn, Bǫmburr, Nóri,
Ái, Mjǫðvitnir,
‘Veigr ok Gandálfr, Vindálfr, Þráinn, Þekkr ok Þorinn, Þrór, Vitr ok Litr, Nár ok Nýráðr — nú hefi ek dverga — Reginn ok Ráðsviðr — rétt um talða.
‘Fíli, Kíli, Fundinn, Náli, Hepti, Víli, Hánarr, Svíurr, Frár, Hornbori, Frægr ok Lóni, Aurvangr, Jari, Eikinskjaldi.
‘Mál er dverga í Dvalins liði ljóna kindum til Lofars telja: þeir er sóttu frá Salarsteini
Aurvanga sjǫt
‘Þar var Draupnir Hár, Haugspori,
Skirvir, Virvir,
‘Álfr ok Yngvi,
Fjalarr ok Frosti, Finnr ok Ginnarr;
þat mun uppi, meðan ǫld lifir, langniðja tal Lofars hafat.
‘Unz þrír kvómu ór því liði, ǫflgir ok ástgir, Æsir, at húsi; fundu á landi, lítt megandi, Ask ok Emblu, ørlǫglausa.

‘Ǫnd þau né áttu, óð þau né hǫfðu, lá né læti né litu góða;
ǫnd gaf Óðinn, óð gaf Hœnir,
lá gaf Lóðurr ok litu góða.
‘Ask veit ek standa, heitir Yggdrasill, hár baðmr ausinn hvíta auri; þaðan koma dǫggvar, þærs í dala falla, stendr æ yfir grœnn Urðar brunni.
‘Þaðan koma meyjar, margs vitandi, þrjár, ór þeim sæ er und þolli stendr; Urð hétu eina, aðra Verðandi — skáru á skíði — Skuld ina þriðju; þær lǫg lǫgðu, þær líf kuru alda bǫrnum, ørlǫg seggja.
‘Þat man hon fólkvíg fyrst í heimi, er Gullveigu geirum studdu, ok í hǫll Hárs hana brendu; þrysvar brendu þrysvar borna, opt, ósjaldan, þó hon enn lifir.
‘Heiði hana hétu, hvars til húsa kom, vǫlu velspá, vitti hon ganda; seið hon kunni, seið hon leikin,
æ var hon angan illrar brúðar.
‘Þá gengu regin ǫll á rǫkstóla, ginnheilǫg goð, ok um þat gættusk, hvárt skyldu Æsir afráð gjalda eða skyldu goðin ǫll gildi eiga.
‘Fleygði Óðinn ok í fólk um skaut — þat var enn fólkvíg fyrst í heimi; brotinn var borðvegr borgar Ása, knáttu Vanir vígspá vǫllu sporna.
‘Þá gengu regin ǫll á rǫkstóla, ginnheilǫg goð, ok um þat gættusk: hverr hefði lopt allt lævi blandit eða ætt jǫtuns Óðs mey gefna.
‘Þórr einn þar var, þrunginn móði, hann sjaldan sitr er hann slíkt um fregn; á gengusk eiðar, orð ok sœri,
mál ǫll meginlig er á meðal fóru.
‘Veit hon Heimdallar hljóð um fólgit undir heiðvǫnum helgum baðmi;
á sér hon ausask aurgum forsi
af veði Valfǫðrs. Vituð ér enn, eða hvat?
‘Ein sat hon úti, þá er inn aldni kom, Yggjungr Ása, ok í augu leit: “Hvers fregnið mik? Hví freistið mín? Allt veit ek, Óðinn, hvar þú auga falt,
í inum mæra Mímis brunni; drekkr mjǫð Mímir morgin hverjan
af veði Valfǫðrs!” Vituð ér enn, eða hvat?
‘Valði henni Herfǫðr hringa ok men, fé, spjǫll spaklig ok spáganda;
sá hon vítt ok um vítt of verǫld hverja.
‘Sá hon valkyrjur, vítt um komnar, gǫrvar at ríða til goðþjóðar; Skuld helt skildi, en Skǫgul ǫnnur, Gunnr, Hildr, Gǫndul ok Geirskǫgul; nú eru talðar nǫnnur Herjans, gǫrvar at ríða grund, valkyrjur.
‘Ek sá Baldri, blóðgum tívur, Óðins barni, ørlǫg fólgin; stóð um vaxinn, vǫllum hæri, mjór ok mjǫk fagr, mistilteinn.
‘Varð af þeim meiði, er mær sýndisk, harmflaug hættlig; Hǫðr nam skjóta; Baldrs bróðir var of borinn snemma, sá nam Óðins sonr einnættr vega.
‘Þó hann æva hendr né hǫfuð kembði, áðr á bál um bar Baldrs andskota; en Frigg um grét í Fensǫlum
vá Valhallar. Vituð ér enn, eða hvat?

‘Hapt sá hon liggja undir Hveralundi, lægjarns líki Loka áþekkjan;
þar sitr Sigyn, þeygi um sínum ver velglýjuð. Vituð ér enn, eða hvat?
‘Á fellr austan um eitrdala, sǫxum ok sverðum, Slíðr heitir sú.
‘Stóð fyr norðan á Niðavǫllum salr ór gulli Sindra ættar;
en annarr stóð á Ókólni, bjórsalr jǫtuns, en sá Brimir heitir.
‘Sal sá hon standa sólu fjarri, Nástrǫndu á, norðr horfa dyrr;
fellu eitrdropar sá er undinn salr
‘Sá hon þar vaða menn meinsvara
inn um ljóra, orma hryggjum.
þunga strauma ok morðvarga,
ok þanns annars glepr eyrarúnu; þar saug Niðhǫggr nái framgengna, sleit vargr vera. Vituð ér enn, eða hvat?
‘Austr sat in aldna í Járnviði
ok fœddi þar Fenris kindir; verðr af þeim ǫllum einna nǫkkurr tungls tjúgari í trolls hami.
‘Fyllisk fjǫrvi feigra manna,
rýðr ragna sjǫt rauðum dreyra; svǫrt var ða sólskin of sumur eptir, veðr ǫll válynd. Vituð ér enn, eða hvat?
‘Sat þar á haugi ok sló hǫrpu gýgjar hirðir, glaðr Eggþér; gól um honum í gaglviði fagrrauðr hani, sá er Fjalarr heitir.
‘Gól um Ásum Gullinkambi, sá vekr hǫlða at Herjafǫðrs; en annarr gelr fyr jǫrð neðan, sótrauðr hani, at sǫlum Heljar.

‘Geyr Garmr mjǫk fyr Gnipahelli, festr mun slitna en freki renna; fjǫlð veit hon frœða, fram sé ek lengra, um ragna rǫk rǫmm, sigtíva.
‘Brœðr munu berjask ok at bǫnum verða, munu systrungar sifjum spilla;
hart er í heimi, hórdómr mikill; skeggǫld, skálmǫld — skildir ru klofnir — vindǫld, vargǫld, áðr verǫld steypisk; mun engi maðr ǫðrum þyrma.
‘Leika Míms synir, en mjǫtuðr kyndisk at inu galla Gjallarhorni;
hátt blæss Heimdallr — horn er á lopti — mælir Óðinn við Míms hǫfuð.
‘Ymr it aldna tré, en jǫtunn losnar; skelfr Yggdrasils askr standandi.
‘Geyr nú Garmr mjǫk fyr Gnipahelli, festr mun slitna en freki renna; fjǫlð veit hon frœða, fram sé ek lengra, um ragna rǫk rǫmm, sigtíva.
‘Hrymr ekr austan, hefisk lind fyrir, snýsk Jǫrmungandr í jǫtunmóði, ormr knýr unnir, en ari hlakkar, slítr nái neffǫlr, Naglfar losnar.
‘Kjóll ferr austan, koma munu Muspells um lǫg lýðir, en Loki stýrir;
fara fífls megir með freka allir, þeim er bróðir Býleipts í fǫr.
‘Hvat er með Ásum? Hvat er með álfum? Gnýr allr Jǫtunheimr, Æsir ru á þingi; stynja dvergar fyr steindurum, veggbergs vísir. Vituð ér enn, eða hvat?
‘Surtr ferr sunnan með sviga lævi, skínn af sverði sól valtíva; grjótbjǫrg gnata en gífr rata, troða halir Helveg, en himinn klofnar.

‘Þá kømr Hlínar harmr annarr fram, er Óðinn ferr við úlf vega,
en bani Belja bjartr at Surti;
þá mun Friggjar falla Angantýr.
‘Þá kømr inn mikli mǫgr Sigfǫður, Víðarr, vega at valdýri;
lætr hann megi Hveðrungs mund um standa hjǫr til hjarta; þá er hefnt fǫður.
‘Þá kømr inn mæri mǫgr Hlóðynjar, gengr Óðins sonr við úlf vega; drepr hann af móði Miðgarðs véur; munu halir allir heimstǫð ryðja; gengr fet níu Fjǫrgynjar burr, neppr, frá naðri niðs ókvíðnum.
‘Sól tér sortna, sígr fold í mar, hverfa af himni heiðar stjǫrnur; geisar eimi við aldnara, leikr hár hiti við himin sjálfan.
‘Geyr nú Garmr mjǫk fyr Gnipahelli, festr mun slitna en freki renna; fjǫlð veit hon frœða, fram sé ek lengra, um ragna rǫk rǫmm, sigtíva.
‘Sér hon upp koma ǫðru sinni jǫrð ór ægi, iðjagrœna; falla forsar, flýgr ǫrn yfir, sá er á fjalli fiska veiðir.
‘Finnask Æsir á Iðavelli
ok um moldþinur mátkan dœma, ok á Fimbultýs fornar rúnar.
‘Þar munu eptir undrsamligar gullnar tǫflur í grasi finnask, þærs í árdaga áttar hǫfðu.
‘Munu ósánir akrar vaxa,
bǫls mun alls batna; Baldr mun koma; búa þeir Hǫðr ok Baldr Hropts sigtóptir, vel, valtívar. Vituð ér enn, eða hvat?
	`
)
