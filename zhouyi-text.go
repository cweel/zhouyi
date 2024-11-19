package zhouyi

import (
	"os"
)

// Gua type Gua
type Gua struct {
	Tuan string
	Xi   [6]string
}

// Text 定义文本
func Text() []Gua {
	var yi = make([]Gua, 64)
	//乾 经文
	yi[0].Tuan = "乾: 元，亨,利，贞。\n 彖曰：大哉乾元，万物资始，乃统天。云行雨施，品物流形。大明始终，六位时成，时 乘六龙以御天。乾道变化，各正性命，保合大和，乃利贞。首出庶物，万国咸宁。\n  象曰：天行健，君子以自强不息。潜龙勿用，阳在下也。见龙再田，德施普也。终日乾 乾，反复道也。或跃在渊，进无咎也。飞龙在天，大人造也。亢龙有悔，盈不可久也。用 九，天德不可为首也。"
	yi[0].Xi[0] = "初九 潜龙勿用。"
	yi[0].Xi[1] = "九二 见龙在田,利见大人。"
	yi[0].Xi[2] = "九三 君子终日乾乾，夕惕若，厉，无咎。"
	yi[0].Xi[3] = "九四 或跃在渊，无咎。"
	yi[0].Xi[4] = "九五 飞龙在天，利见大人。"
	yi[0].Xi[5] = "上九 亢龙有悔。用九：见群龙无首，吉。"

	//坤
	yi[1].Tuan = "坤：元，亨，利牝马之贞。君子有攸往，先迷后得主，利西南得朋，东北丧朋。安贞吉。\n 彖曰：至哉坤元，万物资生，乃顺承天。坤厚载物，德合无疆。含弘光大，品物咸亨。牝马地类，行地无疆，柔顺利贞。君子攸行，先迷失道，后顺得常。西南得朋，乃与类行，东北丧朋，乃终有庆。安贞之吉，应地无疆。用六：利永贞。 象曰：用六永贞，以大终也。"
	yi[1].Xi[0] = "初六：履霜，坚冰至。 象曰：履霜坚冰，阴始凝也。驯致其道，至坚冰也。"
	yi[1].Xi[1] = "六二：直，方，大，不习无不利。 象曰：六二之动，直以方也。不习无不利，地道光也。"
	yi[1].Xi[2] = "六三：含章可贞。或从王事，无成有终。 象曰：含章可贞，以时发也。或从王事，知光大也。"
	yi[1].Xi[3] = "六四：括囊，无咎，无誉。 象曰：括囊无咎，慎不害也。"
	yi[1].Xi[4] = "六五：黄裳，元吉。 象曰：黄裳元吉，文在中也。"
	yi[1].Xi[5] = "上六：战龙於野，其血玄黄。 象曰：战龙於野，其道穷也。"

	//屯
	yi[2].Tuan = "屯：元，亨，利，贞，勿用，有攸往，利建侯。彖曰：屯，刚柔始交而难生，动乎险中，大亨贞。雷雨之动满盈，天造草昧，宜建侯而 不宁。\n  象曰：云，雷，屯，君子以经纶。"
	yi[2].Xi[0] = "初九：磐桓，利居贞，利建侯。 象曰：虽磐桓，志 行正也。以贵下贱，大得民也。"
	yi[2].Xi[1] = "六二：屯如(辶颤)如，乘马班如。匪寇婚媾，女子贞不字，十年乃字。 象曰：六二之难，乘 刚也。十年乃字，反常也。"
	yi[2].Xi[2] = "六三：既鹿无虞，惟入于林中，君子几不如舍，往吝。 象曰：既鹿无虞，以纵禽也。君 子舍之，往吝穷也。"
	yi[2].Xi[3] = "六四：乘马班如，求婚媾，无不利。 象曰：求而往，明也。"
	yi[2].Xi[4] = "九五：屯其膏，小贞吉，大贞凶。 象曰：屯其膏，施未光也。"
	yi[2].Xi[5] = "上六：乘马班如，泣血涟如。 象曰：泣血涟如，何可长也。"

	//蒙
	yi[3].Tuan = "蒙：亨。匪我求童蒙，童蒙求我。初噬告，再三渎，渎则不告。利贞。彖曰：蒙，山下有险，险而止，蒙。蒙亨，以亨行时中也。匪我求童蒙，童蒙求我，志应也。初噬告，以刚中也。再三渎，渎则不告，渎蒙也。蒙以养正，圣功也。\n  象曰：山下出泉，蒙，君子以果行育德。"
	yi[3].Xi[0] = "初六：发蒙，利用刑人，用说桎梏，以往吝。 象曰：利用刑人， 以正法也。"
	yi[3].Xi[1] = "九二：包蒙吉，纳妇吉，子克家。 象曰：子克家，刚柔接也。"
	yi[3].Xi[2] = "六三：勿用娶女，见金夫，不有躬，无攸利。 象曰：勿用娶女，行不顺也。"
	yi[3].Xi[3] = "六四：困蒙，吝。 象曰：困蒙之吝，独远实也。"
	yi[3].Xi[4] = "六五：童蒙，吉。 象曰：童蒙之吉，顺以巽也。"
	yi[3].Xi[5] = "上九：击蒙，不利为寇，利御寇。 象曰：利用御寇，上下顺也。"

	//需
	yi[4].Tuan = "需：有孚，光亨，贞吉。利涉大川。彖曰：需，须也，险在前也。刚健而不陷，其义不困穷矣。需有孚，光亨，贞吉。位乎 天位，以正中也。利涉大川，往有功也。\n  象曰：云上於天，需，君子以饮食宴乐。"
	yi[4].Xi[0] = "初九：需 于郊。利用恒，无咎。 象曰：需于郊，不犯难行也。利用恒，无咎，未失常也。"
	yi[4].Xi[1] = "九二：需于沙。小有言，终吉。 象曰：需于沙，衍在中也。虽小有言，以终吉也。"
	yi[4].Xi[2] = "九三：需于泥，致寇至。 象曰：需于泥，灾在外也。自我致寇，敬慎不败也。"
	yi[4].Xi[3] = "六四：需于血，出自穴。 象曰：需于血，顺以听也。"
	yi[4].Xi[4] = "九五：需于酒食，贞吉。 象曰：酒食贞吉，以中正也。"
	yi[4].Xi[5] = "上六：入于穴，有不速之客三人 来，敬之终吉。 象曰：不速之客来，敬之终吉。虽不当位，未大失也。"

	//讼
	yi[5].Tuan = "讼：有孚窒惕，中吉，终凶。利见大人，不利涉大川。彖曰：讼，上刚下险，险而健讼。讼，有孚窒惕，中吉，刚来而得中也。终凶，讼不可成 也。利见大人，尚中正也。不利涉大川，入于渊也。\n  象曰：天与水违行，讼，君子以作事谋始。"
	yi[5].Xi[0] = "初六：不永所事，小有言，终吉。 象曰：不 永所事，讼不可长也。虽有小言，其辩明也。"
	yi[5].Xi[1] = "九二：不克讼，归而逋，其邑人三百户，无眚。 象曰：不克讼，归而逋也。自下讼上， 患至掇也。"
	yi[5].Xi[2] = "六三：食旧德，贞厉，终吉，或从王事，无成。 象曰：食旧德，从上吉也。"
	yi[5].Xi[3] = "九四：不克讼，复自命，渝安贞，吉。 象曰：复即命渝，安贞不失也。"
	yi[5].Xi[4] = "九五：讼元吉。 象曰：讼元吉，以中正也。"
	yi[5].Xi[5] = "上九：或锡之鞶带，终朝三褫之。 象曰：以讼受服，亦不足敬也。"

	//师
	yi[6].Tuan = "师：贞，丈人，吉无咎。彖曰：师，众也，贞正也，能以众正，可以王矣。刚中而应，行险而顺，以此毒天下， 而民从之，吉又何咎矣。\n  象曰：地中有水，师，君子以容民畜众。"
	yi[6].Xi[0] = "初六：师出以律，否臧凶。 象曰：师出以律， 失律凶也。"
	yi[6].Xi[1] = "九二：在师中，吉无咎，王三锡命。 象曰：在师中吉，承天宠也。王三锡命，怀万邦也。"
	yi[6].Xi[2] = "六三：师或舆尸，凶。 象曰：师或舆尸，大无功也。"
	yi[6].Xi[3] = "六四：师左次，无咎。 象曰：左次无咎，未失常也。"
	yi[6].Xi[4] = "六五：田有禽，利执言，无咎。长子帅师，弟子舆尸，贞凶。 象曰：长子帅师，以中行也。弟子舆师，使不当也。"
	yi[6].Xi[5] = "上六：大君有命，开国承家，小人勿用。 象曰：大君有命，以正功也。小人勿用，必乱邦也。"

	//比
	yi[7].Tuan = "比：吉。原筮元永贞，无咎。不宁方来，后夫凶。彖曰：比，吉也，比，辅也，下顺从也。原筮元永贞，无咎，以刚中也。不宁方来，上 下应也。后夫凶，其道穷也。\n  象曰：地上有水，比，先王以建万国，亲诸侯。"
	yi[7].Xi[0] = "初六：有孚比之，无咎。有孚盈缶，终 来有他，吉。 象曰：比之初六，有他吉也。"
	yi[7].Xi[1] = "六二：比自内，贞吉。 象曰：比之自内，不自失也。"
	yi[7].Xi[2] = "六三：比之匪人。 象曰：比之匪人，不亦伤乎！"
	yi[7].Xi[3] = "六四：外比之，贞吉。 象曰：外比於贤，以从上也。"
	yi[7].Xi[4] = "九五：显比，王用三驱，失前禽。邑人不诫，吉。 象曰：显比之吉，位正中也。舍逆取 顺，失前禽也。邑人不诫，上使中也。"
	yi[7].Xi[5] = "上六：比之无首，凶。 象曰：比之无首，无所终也。"

	//小蓄
	yi[8].Tuan = "小畜：亨。密云不雨，自我西郊。彖曰：小畜，柔得位，而上下应之，曰小畜。健而巽，刚中而志行，乃亨。密云不雨， 尚往也。自我西郊，施未行也。\n  象曰：风行天上，小畜，君子以懿文德。"
	yi[8].Xi[0] = "初九：复自道，何其咎，吉。 象曰：复自道， 其义吉也。"
	yi[8].Xi[1] = "九二：牵复，吉。 象曰：牵复在中，亦不自失也。"
	yi[8].Xi[2] = "九三：舆说辐，夫妻反目。 象曰：夫 妻反目，不能正室也。"
	yi[8].Xi[3] = "六四：有孚，血去惕出，无咎。 象曰：有孚惕出，上合志也。"
	yi[8].Xi[4] = "九五：有孚挛如，富以其邻。 象曰：有孚挛如，不独富也。"
	yi[8].Xi[5] = "上九：既雨既处，尚德载，妇贞厉。月几望，君子征凶。 象曰：既雨既处，德积载也。 君子征凶，有所疑也。"

	//履
	yi[9].Tuan = "履：履虎尾，不咥人，亨。彖曰：履，柔履刚也。说而应乎乾，是以履虎尾，不咥人，亨。刚中正，履帝位而不 疚，光明也。\n  象曰：上天下泽，履，君子以辨上下，安民志。"
	yi[9].Xi[0] = "初九：素履，往无咎。 象曰：素履之 往，独行愿也。"
	yi[9].Xi[1] = "九二：履道坦坦，幽人贞吉。 象曰：幽人贞吉，中不自乱也。"
	yi[9].Xi[2] = "六三：眇能视，跛能履，履虎尾，咥人，凶。武人为于大君。 象曰：眇能视，不足以 有明也。跛能履，不足以与行也。咥人之凶，位不当也。武人为于大君，志刚也。"
	yi[9].Xi[3] = "九四：履虎尾，愬愬终吉。 象曰：愬愬终吉，志行也。"
	yi[9].Xi[4] = "九五：夬履，贞厉。 象曰：夬履贞厉，位正当也。"
	yi[9].Xi[5] = "上九：视履考祥，其旋元吉。 象曰：元吉在上，大有庆也。"

	//泰
	yi[10].Tuan = "泰：小往大来，吉亨。彖曰：泰，小往大来，吉亨。则是天地交，而万物通也，上下交，而其志同也。内阳而 外阴，内健而外顺，内君子而外小人，君子道长，小人道消也。\n  象曰：天地交泰，后以财（裁）成天地之道，辅相天地之宜，以左右民。"
	yi[10].Xi[0] = "初九：拔茅 茹，以其夤，征吉。 象曰：拔茅征吉，志在外也。"
	yi[10].Xi[1] = "九二：包荒，用冯河，不遐遗，朋亡，得尚于中行。 象曰：包荒，得尚于中行，以光大 也。"
	yi[10].Xi[2] = "九三：无平不陂，无往不复，艰贞无咎。勿恤其孚，于食有福。 象曰：无往不复，天地 际也。"
	yi[10].Xi[3] = "六四：翩翩不富，以其邻，不戒以孚。 象曰：翩翩不富，皆失实也。不戒以孚，中心愿 也。"
	yi[10].Xi[4] = "六五：帝乙归妹，以祉元吉。 象曰：以祉元吉，中以行愿也。"
	yi[10].Xi[5] = "上六：城复于隍，勿用师。自邑告命，贞吝。 象曰：城复于隍，其命乱也。"

	//否
	yi[11].Tuan = "否：否之匪人，不利君子贞，大往小来。彖曰：否之匪人，不利君子贞。大往小来，则是天地不交，而万物不通也，上下不交， 而天下无邦也。内阴而外阳，内柔而外刚，内小人而外君子。小人道长，君子道消也。\n  象曰：天地不交，否，君子以俭德辟难，不可荣以禄。"
	yi[11].Xi[0] = "初六：拔茅茹，以其夤，贞吉 亨。 象曰：拔茅贞吉，志在君也。"
	yi[11].Xi[1] = "六二：包承。小人吉，大人否亨。 象曰：大人否亨，不乱群也。"
	yi[11].Xi[2] = "六三：包羞。 象曰：包羞，位不当也。"
	yi[11].Xi[3] = "九四：有命无咎，畴离祉。 象曰：有命无咎，志行也。"
	yi[11].Xi[4] = "九五：休否，大人吉。其亡其亡，系于苞桑。 象曰：大人之吉，位正当也。"
	yi[11].Xi[5] = "上九：倾否，先否后喜。 象曰：否终则倾，何可长也。"

	//同人
	yi[12].Tuan = "同人：同人于野，亨。利涉大川，利君子贞。彖曰：同人，柔得位得中，而应乎乾，曰同人。同人曰，同人于野，亨。利涉大川，乾 行也。文明以健，中正而应，君子正也。唯君子为能通天下之志。\n  象曰：天与火，同人，君子以类族辨物。"
	yi[12].Xi[0] = "初九：同人于门，无咎。 象曰：出门同人，又谁咎也。"
	yi[12].Xi[1] = "六二：同人于宗，吝。 象曰：同人于宗，吝道也。"
	yi[12].Xi[2] = "九三：伏戎于莽，升其高陵，三岁不兴。 象曰：伏戎于莽，敌刚也。三岁不兴，安行 也。"
	yi[12].Xi[3] = "九四：乘其墉，弗克攻，吉。 象曰：乘其墉，义弗克也，其吉，则困而反则也。"
	yi[12].Xi[4] = "九五：同人，先号啕而后笑。大师克相遇。 象曰：同人之先，以中直也。大师相遇，言 相克也。"
	yi[12].Xi[5] = "上九：同人于郊，无悔。 象曰：同人于郊，志未得也。"

	//大有
	yi[13].Tuan = "大有：元亨。 彖曰：大有，柔得尊位，大中而上下应之，曰大有。其德刚健而文明，应乎天而时行， 是以元亨。\n  象曰：火在天上，大有，君子以竭恶扬善，顺天休命。"
	yi[13].Xi[0] = "初九：无交害，匪咎，艰则无 咎。 象曰：大有初九，无交害也。"
	yi[13].Xi[1] = "九二：大车以载，有攸往，无咎。 象曰：大车以载，积中不败也。"
	yi[13].Xi[2] = "九三：公用亨于天子，小人弗克。 象曰：公用亨于天子，小人害也。"
	yi[13].Xi[3] = "九四：匪其彭，无咎。 象曰：匪其彭，无咎，明辨晰也。"
	yi[13].Xi[4] = "六五：厥孚交如，威如，吉。 象曰：厥孚交如，信以发志也。威如之吉，易而无备也。"
	yi[13].Xi[5] = "上九：自天佑之，吉无不利。 象曰：大有上吉，自天佑也。"

	//谦
	yi[14].Tuan = "谦：亨，君子有终。彖曰：谦，亨，天道下济而光明，地道卑而上行。天道亏盈而益谦，地道变盈而流谦， 鬼神害盈而福谦，人道恶盈而好谦。谦尊而光，卑而不可踰，君子之终也。\n  象曰：地中有山，谦，君子以裒多益寡，称物平施。"
	yi[14].Xi[0] = "初六：谦谦君子，用涉大川，吉。  象曰：谦谦君子，卑以自牧也。"
	yi[14].Xi[1] = "六二：鸣谦，贞吉。 象曰：鸣谦贞吉，中心得也。"
	yi[14].Xi[2] = "九三：劳谦君子，有终吉。 象曰：劳谦君子，万民服也。"
	yi[14].Xi[3] = "六四：无不利，㧑谦。 象曰：无不利，㧑谦，不违则也。"
	yi[14].Xi[4] = "六五：不富，以其邻，利用侵伐，无不利。 象曰：利用侵伐，征不服也。"
	yi[14].Xi[5] = "上六：鸣谦，利用行师，征邑国。 象曰：鸣谦，志未得也。可用行师，征邑国也。"

	//豫
	yi[15].Tuan = "豫：利建侯行师。彖曰：豫，刚应而志行，顺以动，豫。豫，顺以动，故天地如之，而况建侯行师乎？天 地以顺动，故日月不过，而四时不忒，圣人以顺动，则刑罚清而民服。豫之时义大矣哉！\n  象曰：雷出地奋，豫。先王以作乐崇德，殷荐之上帝，以配祖考。"
	yi[15].Xi[0] = "初六：鸣豫，凶。 象 曰：初六鸣豫，志穷凶也。"
	yi[15].Xi[1] = "六二：介于石，不终日，贞吉。 象曰：不终日，贞吉，以中正也。"
	yi[15].Xi[2] = "六三：盱豫，悔。迟有悔。 象曰：盱豫有悔，位不当也。"
	yi[15].Xi[3] = "九四：由豫，大有得。勿疑。朋盍簪。 象曰：由豫，大有得，志大行也。"
	yi[15].Xi[4] = "六五：贞疾，恒不死。 象曰：六五贞疾，乘刚也。恒不死，中未亡也。"
	yi[15].Xi[5] = "上六：冥豫，成有渝，无咎。 象曰：冥豫在上，何可长也。"

	//随
	yi[16].Tuan = "随：元亨利贞，无咎。彖曰：随，刚来而下柔，动而说，随。大亨贞，无咎，而天下随时，随之时义大矣哉！\n  象曰：泽中有雷，随，君子以向晦入宴息。"
	yi[16].Xi[0] = "初九：官有渝，贞吉。出门交有功。 象曰： 官有渝，从正吉也。出门交有功，不失也。"
	yi[16].Xi[1] = "六二：系小子，失丈夫。 象曰：系小子，弗兼与也。"
	yi[16].Xi[2] = "六三：系丈夫，失小子。随有求得，利居贞。 象曰：系丈夫，志舍下也。"
	yi[16].Xi[3] = "九四：随有获，贞凶。有孚在道，以明，何咎。 象曰：随有获，其义凶也。有孚在道， 明功也。"
	yi[16].Xi[4] = "九五：孚于嘉，吉。 象曰：孚于嘉，吉，位正中也。"
	yi[16].Xi[5] = "上六：拘系之，乃从维之。王用亨于西山。 象曰：拘系之，上穷也。"

	//盅
	yi[17].Tuan = "蛊：元亨，利涉大川。先甲三日，后甲三日。彖曰：蛊，刚上而柔下，巽而止，蛊。蛊，元亨，而天下治也。利涉大川，往有事也。 先甲三日，后甲三日，终则有始，天行也。\n  象曰：山下有风，蛊，君子以振民育德。"
	yi[17].Xi[0] = "初六：干父之蛊，有子，考无咎，厉终吉。 象 曰：干父之蛊，意承考也。"
	yi[17].Xi[1] = "九二：干母之蛊，不可贞。 象曰：干母之蛊，得中道也。"
	yi[17].Xi[2] = "九三：干父小有晦，无大咎。 象曰：干父之蛊，终无咎也。"
	yi[17].Xi[3] = "六四：裕父之蛊，往见吝。 象曰：裕父之蛊，往未得也。"
	yi[17].Xi[4] = "六五：干父之蛊，用誉。 象曰：干父之蛊，承以德也。"
	yi[17].Xi[5] = "上九：不事王侯，高尚其事。 象曰：不事王侯，志可则也。"

	//临
	yi[18].Tuan = "临：元，亨，利，贞。至于八月有凶。彖曰：临，刚浸而长。说而顺，刚中而应，大亨以正，天之道也。至于八月有凶，消不久也。\n  象曰：泽上有地，临，君子以教思无穷，容保民无疆。"
	yi[18].Xi[0] = "初九：咸临，贞吉。 象曰：咸临 贞吉，志行正也。"
	yi[18].Xi[1] = "九二：咸临，吉无不利。 象曰：咸临，吉无不利，未顺命也。"
	yi[18].Xi[2] = "六三：甘临，无攸利。既忧之，无咎。 象曰：甘临，位不当也。既忧之，咎不长也。"
	yi[18].Xi[3] = "六四：至临，无咎。 象曰：至临无咎，位当也。"
	yi[18].Xi[4] = "六五：知临，大君之宜，吉。 象曰：大君之宜，行中之谓也。"
	yi[18].Xi[5] = "上六：敦临，吉无咎。 象曰：敦临之吉，志在内也。"

	//观
	yi[19].Tuan = "观：盥而不荐，有孚□若。彖曰：大观在上，顺而巽，中正以观天下。观，盥而不荐，有孚颙若，下观而化也。观 天之神道，而四时不忒，圣人以神道设教，而天下服矣。\n  象曰：风行地上，观，先王以省方，观民设教。"
	yi[19].Xi[0] = "初六：童观，小人无咎，君子吝。 象 曰：初六童观，小人道也。"
	yi[19].Xi[1] = "六二：窥观，利女贞。 象曰：窥观女贞，亦可丑也。"
	yi[19].Xi[2] = "六三：观我生，进退。 象曰：观我生，进退，未失道也。"
	yi[19].Xi[3] = "六四：观国之光，利用宾于王。 象曰：观国之光，尚宾也。"
	yi[19].Xi[4] = "九五：观我生，君子无咎。 象曰：观我生，观民也。"
	yi[19].Xi[5] = "上九：观其生，君子无咎。 象曰：观其生，志未平也。"

	//噬嗑
	yi[20].Tuan = "噬嗑：亨。利用狱。彖曰：颐中有物，曰噬嗑，噬嗑而亨。刚柔分，动而明，雷电合而章。柔得中而上行， 虽不当位，利用狱也。\n  象曰：雷电噬嗑，先王以明罚敕法。"
	yi[20].Xi[0] = "初九：履校灭趾，无咎。 象曰：履校灭趾，不行也。"
	yi[20].Xi[1] = "六二：噬肤灭鼻，无咎。 象曰：噬肤灭鼻，乘刚也。"
	yi[20].Xi[2] = "六三：噬腊肉，遇毒，小吝，无咎。 象曰：遇毒，位不当也。"
	yi[20].Xi[3] = "九四：噬乾胏，得金矢，利艰贞，吉。 象曰：利艰贞吉，未光也。"
	yi[20].Xi[4] = "六五：噬乾肉，得黄金，贞厉，无咎。 象曰：贞厉无咎，得当也。"
	yi[20].Xi[5] = "上九：何校灭耳，凶。 象曰：何校灭耳，聪不明也。"

	//贲
	yi[21].Tuan = "贲：亨。小利有所往。彖曰：贲，亨，柔来而文刚，故亨。分刚上而文柔，故小利有攸往。天文也，文明以 止，人文也。观乎天文，以察时变，观乎人文，以化成天下。\n  象曰：山下有火，贲，君子以明庶政，无敢折狱。	"
	yi[21].Xi[0] = "初九：贲其趾，舍车而徒。 象曰：舍 车而徒，义弗乘也。"
	yi[21].Xi[1] = "六二：贲其须。 象曰：贲其须，与上兴也。"
	yi[21].Xi[2] = "九三：贲如濡如，永贞吉。 象曰：永贞之吉，终莫之陵也。"
	yi[21].Xi[3] = "六四：贲如皤如，白马翰如，匪寇婚媾。 象曰：六四，当位疑也。匪寇婚媾，终无尤也。"
	yi[21].Xi[4] = "六五：贲于丘园，束帛戋戋，吝，终吉。 象曰：六五之吉，有喜也。"
	yi[21].Xi[5] = "上九：白贲，无咎。 象曰：白贲无咎，上得志也。"

	//剥
	yi[22].Tuan = "剥：不利有攸往。彖曰：剥，剥也，柔变刚也。不利有攸往，小人长也。顺而止之，观 象也。君子尚消息 盈虚，天行也。\n  象曰：山附地上，剥，上以厚下安宅。"
	yi[22].Xi[0] = "初六：剥(爿木)以足，蔑贞凶。 象曰：剥(爿木)以足， 以灭下也。"
	yi[22].Xi[1] = "六二：剥(爿木)以辨，蔑贞凶。 象曰：剥(爿木)以辨，未有与也。"
	yi[22].Xi[2] = "六三：剥之，无咎。 象曰：剥之无咎，失上下也。"
	yi[22].Xi[3] = "六四：剥(爿木)以肤，凶。 象曰：剥(爿木)以肤，切近灾也。"
	yi[22].Xi[4] = "六五：贯鱼，以宫人宠，无不利。 象曰：以宫人宠，终无尤也。"
	yi[22].Xi[5] = "上九：硕果不食，君子得舆，小人剥庐。 象曰：君子得舆，民所载也。小人剥庐，终不可 用也。"

	//复
	yi[23].Tuan = "复：亨。出入无疾，朋来无咎。反复其道，七日来复，利有攸往。彖曰：复亨，刚反，动而以顺行，是以出入无疾，朋来无咎。反复其道，七日来复，天 行也。利有攸往，刚长也。复其见天地之心乎？\n  象曰：雷在地中，复，先王以至日闭关，商旅不行，后不省方。"
	yi[23].Xi[0] = "初九：不复远，无只 悔，元吉。 象曰：不远之复，以修身也。"
	yi[23].Xi[1] = "六二：休复，吉。 象曰：休复之吉，以下仁也。"
	yi[23].Xi[2] = "六三：频复，厉无咎。 象曰：频复之厉，义无咎也。"
	yi[23].Xi[3] = "六四：中行独复。 象曰：中行独复，以从道也。"
	yi[23].Xi[4] = "六五：敦复，无悔。 象曰：敦复无悔，中以自考也。"
	yi[23].Xi[5] = "上六：迷复，凶，有灾眚。用行师，终有大败，以其国君，凶，至于十年，不克征。 象 曰：迷复之凶，反君道也。"

	//无妄
	yi[24].Tuan = "无妄：元，亨，利，贞。其匪正有眚，不利有攸往。彖曰：无妄，刚自外来，而为主於内。动而健，刚中而应，大亨以正，天之命也。其匪 正有眚，不利有攸往。无妄之往，何之矣？天命不佑，行矣哉？\n  象曰：天下雷行，物与无妄，先王以茂对时，育万物。"
	yi[24].Xi[0] = "初九：无妄，往吉。 象曰：无妄 之往，得志也。"
	yi[24].Xi[1] = "六二：不耕获，不菑畲，则利有攸往。 象曰：不耕获，未富也。"
	yi[24].Xi[2] = "六三：无妄之灾，或系之牛，行人之得，邑人之灾。 象曰：行人得牛，邑人灾也。"
	yi[24].Xi[3] = "九四：可贞，无咎。 象曰：可贞无咎，固有之也。 "
	yi[24].Xi[4] = "九五：无妄之疾，勿药有喜。 象曰：无妄之药，不可试也。"
	yi[24].Xi[5] = "上九：无妄，行有眚，无攸利。 象曰：无妄之行，穷之灾也。"

	//大畜
	yi[25].Tuan = "大畜：利贞，不家食吉，利涉大川。彖曰：大畜，刚健笃实辉光，日新其德，刚上而尚贤。能止健，大正也。不家食吉，养 贤也。利涉大川，应乎天也。\n  象曰：天在山中，大畜，君子以多识前言往行，以畜其德。"
	yi[25].Xi[0] = "初九：有厉利已。 象曰：有 厉利已，不犯灾也。"
	yi[25].Xi[1] = "九二：舆说辐。 象曰：舆说辐，中无尤也。"
	yi[25].Xi[2] = "九三：良马逐，利艰贞。曰闲舆卫，利有攸往。 象曰：利有攸往，上合志也。"
	yi[25].Xi[3] = "六四：童豕之牿，元吉。 象曰：六四元吉，有喜也。 "
	yi[25].Xi[4] = "六五：豶豕之牙，吉。 象曰：六五之吉，有庆也。"
	yi[25].Xi[5] = "上九：何天之衢，亨。 象曰：何天之衢，道大行也。"

	//颐
	yi[26].Tuan = "颐：贞吉。观颐，自求口实。彖曰：颐贞吉，养正则吉也。观颐，观其所养也，自求口实，观其自养也。天地养万 物，圣人养贤，以及万民，颐之时义大矣哉！\n  象曰：山下有雷，颐，君子以慎言语，节饮食。"
	yi[26].Xi[0] = "初九：舍尔灵龟，观我朵颐，凶。 象 曰：观我朵颐，亦不足贵也。"
	yi[26].Xi[1] = "六二：颠颐，拂经，于丘颐，征凶。 象曰：六二征凶，行失类也。"
	yi[26].Xi[2] = "六三：拂颐，贞凶，十年勿用，无攸利。 象曰：十年勿用，道大悖也。"
	yi[26].Xi[3] = "六四：颠颐吉，虎视眈眈，其欲逐逐，无咎。 象曰：颠颐之吉，上施光也。"
	yi[26].Xi[4] = "六五：拂经，居贞吉，不可涉大川。 象曰：居贞之吉，顺以从上也。"
	yi[26].Xi[5] = "上九：由颐，厉吉，利涉大川。 象曰：由颐厉吉，大有庆也。"

	//大过
	yi[27].Tuan = "大过：栋桡，利有攸往，亨。彖曰：大过，大者过也。栋桡，本末弱也。刚过而中，巽而说行，利有攸往，乃亨。大 过之时义大矣哉！\n  象曰：泽灭木，大过，君子以独立不惧，遁世无闷。"
	yi[27].Xi[0] = "初六：藉用白茅，无咎。 象曰：藉用白茅，柔在下也。"
	yi[27].Xi[1] = "九二：枯杨生稊，老夫得其女妻，无不利。 象曰：老夫女妻，过以相与也。"
	yi[27].Xi[2] = "九三：栋桡，凶。 象曰：栋桡之凶，不可以有辅也。"
	yi[27].Xi[3] = "九四：栋隆，吉，有它吝。 象曰：栋隆之吉，不桡乎下也。"
	yi[27].Xi[4] = "九五：枯杨生华，老妇得士夫，无咎无誉。 象曰：枯杨生华，何可久也。老妇士夫，亦可丑也。"
	yi[27].Xi[5] = "上六：过涉灭顶，凶，无咎。 象曰：过涉之凶，不可咎也。"

	//坎
	yi[28].Tuan = "坎：习坎，有孚，维心亨，行有尚。彖曰：习坎，重险也。水流而不盈，行险而不失其信。维心亨，乃以刚中也。行有尚， 往有功也。天险不可升也，地险山川丘陵也，王公设险以守其国，坎之时用大矣哉！\n  象曰：水洊至，习坎，君子以常德行，习教事。"
	yi[28].Xi[0] = "初六：习坎，入于坎窞，凶。 象曰：习坎入坎，失道凶也。"
	yi[28].Xi[1] = "九二：坎有险，求小得。 象曰：求小得，未出中也。"
	yi[28].Xi[2] = "六三：来之坎坎，险且枕，入于坎窞，勿用。 象曰：来之坎坎，终无功也。"
	yi[28].Xi[3] = "六四：樽酒簋贰，用缶，纳约自牖，终无咎。 象曰：樽酒簋贰，刚柔际也。"
	yi[28].Xi[4] = "九五：坎不盈，只既平，无咎。 象曰：坎不盈，中未大也。"
	yi[28].Xi[5] = "上六：用徽繹，置于丛棘，三岁不得，凶。 象曰：上六失道，凶三岁也。"

	//离
	yi[29].Tuan = "离：利贞，亨。畜牝牛，吉。彖曰：离，丽也，日月丽乎天，百谷草木丽乎土，重明以丽乎正，乃化成天下。柔丽乎 中正，故亨，是以畜牝牛吉也。\n  象曰：明两作离，大人以继明照于四方。"
	yi[29].Xi[0] = "初九：履错然，敬之无咎。 象曰：履错之敬， 以辟咎也。"
	yi[29].Xi[1] = "六二：黄离，元吉。 象曰：黄离元吉，得中道也。"
	yi[29].Xi[2] = "九三：日昃之离，不鼓缶而歌，则大耋之嗟，凶。 象曰：日昃之离，何可久也。"
	yi[29].Xi[3] = "九四：突如其来如，焚如，死如，弃如。 象曰：突如其来如，无所容也。"
	yi[29].Xi[4] = "六五：出涕沱若，戚嗟若，吉。 象曰：六五之吉，离王公也。"
	yi[29].Xi[5] = "上九：王用出征，有嘉折首，获其匪丑，无咎。 象曰：王用出征，以正邦也。"

	//下经。。。

	//咸
	yi[30].Tuan = "咸：亨，利贞，取女吉。彖曰：咸，感也。柔上而刚下，二气感应以相与，止而说，男下女，是以亨利贞，取女 吉也。天地感而万物化生，圣人感人心而天下和平，观其所感，而天地万物之情可见矣！\n  象曰：山上有泽，咸，君子以虚受人。"
	yi[30].Xi[0] = "初六：咸其拇。 象曰：咸其拇，志在外也。"
	yi[30].Xi[1] = "六二：咸其腓，凶，居吉。 象曰：虽凶，居吉，顺不害也。"
	yi[30].Xi[2] = "九三：咸其股，执其随，往吝。 象曰：咸其股，亦不处也。志在随人，所执下也。"
	yi[30].Xi[3] = "九四：贞吉悔亡，憧憧往来，朋从尔思。 象曰：贞吉悔亡，未感害也。憧憧往来，未光 大也。"
	yi[30].Xi[4] = "九五：咸其脢，无悔。 象曰：咸其脢，志末也。"
	yi[30].Xi[5] = "上六：咸其辅，颊，舌。 象曰：咸其辅，颊，舌，滕口说也。"

	//恒
	yi[31].Tuan = "恒：亨，无咎，利贞，利有攸往。彖曰：恒，久也。刚上而柔下，雷风相与，巽而动，刚柔皆应，恒。恒亨无咎，利贞， 久於其道也，天地之道，恒久而不已也。利有攸往，终则有始也。日月得天，而能久照，四 时变化，而能久成，圣人久於其道，而天下化成，观其所恒，而天地万物之情可见矣！\n  象曰：雷风，恒，君子以立不易方。"
	yi[31].Xi[0] = "初六：浚恒，贞凶，无攸利。 象曰：浚恒之凶，始 求深也。"
	yi[31].Xi[1] = "九二：悔亡。 象曰：九二悔亡，能久中也。"
	yi[31].Xi[2] = "九三：不恒其德，或承之羞，贞吝。 象曰：不恒其德，无所容也。"
	yi[31].Xi[3] = "九四：田无禽。 象曰：久非其位，安得禽也。"
	yi[31].Xi[4] = "六五：恒其德，贞，妇人吉，夫子凶。 象曰：妇人贞吉，从一而终也。夫子制义，从妇 凶也。"
	yi[31].Xi[5] = "上六：振恒，凶。 象曰：振恒在上，大无功也。"

	//遁
	yi[32].Tuan = "遁：亨，小利贞。 彖曰：遁亨，遁而亨也。刚当位而应，与时行也。小利贞，浸而长也。遁之时义大矣 哉！\n  象曰：天下有山，遁，君子以远小人，不恶而严。"
	yi[32].Xi[0] = "初六：遁尾，厉，勿用有攸往。 象曰：遁尾之厉，不往何灾也。"
	yi[32].Xi[1] = "六二：执之用黄牛之革，莫之胜说。 象曰：执用黄牛，固志也。"
	yi[32].Xi[2] = "九三：系遁，有疾厉，畜臣妾吉。 象曰：系遁之厉，有疾惫也。畜臣妾吉，不可大事也。"
	yi[32].Xi[3] = "九四：好遁，君子吉，小人否。 象曰：君子好遁，小人否也。"
	yi[32].Xi[4] = "九五：嘉遁，贞吉。 象曰：嘉遁贞吉，以正志也。"
	yi[32].Xi[5] = "上九：肥遁，无不利。 象曰：肥遁，无不利，无所疑也。"

	//大壮
	yi[33].Tuan = "大壮：利贞。 彖曰：大壮，大者壮也。刚以动，故壮。大壮利贞，大者正也。正大而天地之情可见矣！\n  象曰：雷在天上，大壮，君子以非礼勿履。"
	yi[33].Xi[0] = "初九：壮于趾，征凶，有孚。 象曰：壮于 趾，其孚穷也。"
	yi[33].Xi[1] = "九二：贞吉。 象曰：九二贞吉，以中也。"
	yi[33].Xi[2] = "九三：小人用壮，君子用罔，贞厉。羝羊触藩，羸其角。 象曰：小人用壮，君子罔也。"
	yi[33].Xi[3] = "九四：贞吉悔亡，藩决不羸，壮于大舆之輹。 象曰：藩决不羸，尚往也。"
	yi[33].Xi[4] = "六五：丧羊于易，无悔。 象曰：丧羊于易，位不当也。"
	yi[33].Xi[5] = "上六：羝羊触藩，不能退，不能遂，无攸利，艰则吉。 象曰：不能退，不能遂，不祥 也。艰则吉，咎不长也。"

	//晋
	yi[34].Tuan = "晋：康侯用锡马蕃庶，昼日三接。彖曰：晋，进也。明出地上，顺而丽乎大明，柔进而上行。是以康侯用锡马蕃庶，昼日 三接也。\n  象曰：明出地上，晋，君子以自昭明德。"
	yi[34].Xi[0] = "初六：晋如，摧如，贞吉。罔孚，裕无咎。 象 曰：晋如，摧如，独行正也。裕无咎，未受命也。"
	yi[34].Xi[1] = "六二：晋如，愁如，贞吉。受兹介福，于其王母。 象曰：受之介福，以中正也。"
	yi[34].Xi[2] = "六三：众允，悔亡。 象曰：众允之，志上行也。"
	yi[34].Xi[3] = "九四：晋如硕鼠，贞厉。 象曰：硕鼠贞厉，位不当也。"
	yi[34].Xi[4] = "六五：悔亡，失得勿恤，往吉无不利。 象曰：失得勿恤，往有庆也。"
	yi[34].Xi[5] = "上九：晋其角，维用伐邑，厉吉无咎，贞吝。 象曰：维用伐邑，道未光也。"

	//明夷
	yi[35].Tuan = "明夷：利艰贞。 彖曰：明入地中，明夷。内文明而外柔顺，以蒙大难，文王以之。利艰贞，晦其明也， 内难而能正其志，箕子以之。\n  象曰：明入地中，明夷，君子以莅众，用晦而明。"
	yi[35].Xi[0] = "初九：明夷于飞，垂其翼。君子于 行，三日不食，有攸往，主人有言。 象曰：君子于行，义不食也。"
	yi[35].Xi[1] = "六二：明夷，夷于左股，用拯马壮，吉。 象曰：六二之吉，顺以则也。"
	yi[35].Xi[2] = "九三：明夷于南狩，得其大首，不可疾贞。 象曰：南狩之志，乃大得也。"
	yi[35].Xi[3] = "六四：入于左腹，获明夷之心，出于门庭。 象曰：入于左腹，获心意也。"
	yi[35].Xi[4] = "六五：箕子之明夷，利贞。 象曰：箕子之贞，明不可息也。"
	yi[35].Xi[5] = "上六：不明晦，初登于天，后入于地。 象曰：初登于天，照四国也。后入于地，失则也。"

	//家人
	yi[36].Tuan = "家人：利女贞。彖曰：家人，女正位乎内，男正位乎外，男女正，天地之大义也。家人有严君焉，父母之谓也。父父，子子，兄兄，弟弟，夫夫，妇妇，而家道正，正家而天下定矣。\n  象曰：风自火出，家人，君子以言有物，而行有恒。"
	yi[36].Xi[0] = "初九：闲有家，悔亡。 象曰：闲有家，志未变也。"
	yi[36].Xi[1] = "六二：无攸遂，在中馈，贞吉。 象曰：六二之吉，顺以巽也。"
	yi[36].Xi[2] = "九三：家人嗃嗃，悔厉吉，妇子嘻嘻，终吝。 象曰：家人嗃嗃，未失也，妇子嘻嘻，失家节也。"
	yi[36].Xi[3] = "六四：富家，大吉。 象曰：富家大吉，顺在位也。"
	yi[36].Xi[4] = "九五：王假有家，勿恤吉。 象曰：王假有家，交相爱也。"
	yi[36].Xi[5] = "上九：有孚威如，终吉。 象曰：威如之吉，反身之谓也。"

	//睽
	yi[37].Tuan = "睽：小事吉。彖曰：睽，火动而上，泽动而下，二女同居，其志不同行，说而丽乎明，柔进而上行， 得中而应乎刚，是以小事吉。天地睽，而其事同也，男女睽，而其志通也，万物睽，而其事 类也，睽之时用大矣哉！\n  象曰：上火下泽，睽，君子以同而异。"
	yi[37].Xi[0] = "初九：悔亡，丧马勿逐，自复，见恶人无咎。 象 曰：见恶人，以辟咎也。"
	yi[37].Xi[1] = "九二：遇主于巷，无咎。 象曰：遇主于巷，未失道也。"
	yi[37].Xi[2] = "六三：见舆曳，其牛掣，其人天且劓，无初有终。 象曰：见舆曳，位不当也。无初有 终，遇刚也。"
	yi[37].Xi[3] = "九四：睽孤，遇元夫，交孚，厉无咎。 象曰：交孚无咎，志行也。"
	yi[37].Xi[4] = "六五：悔亡，厥宗噬肤，往何咎。 象曰：厥宗噬肤，往有庆也。"
	yi[37].Xi[5] = "上九：睽孤，见豕负涂，载鬼一车，先张之弧，后说之弧，匪寇婚媾，往遇雨则吉。 象 曰：遇雨之吉，群疑亡也。"

	//蹇
	yi[38].Tuan = "蹇：利西南，不利东北，利见大人，贞吉。彖曰：蹇，难也，险在前也。见险而能止，知矣哉！蹇利西南，往得中也，不利东北， 其道穷也。利见大人，往有功也。当位贞吉，以正邦也。蹇之时用大矣哉！\n  象曰：山上有水，蹇，君子以反身修德。"
	yi[38].Xi[0] = "初六：往蹇，来誉。 象曰：往蹇来誉，宜待也。"
	yi[38].Xi[1] = "六二：王臣蹇蹇，匪躬之故。 象曰：王臣蹇蹇，终无尤也。"
	yi[38].Xi[2] = "九三：往蹇来反。 象曰：往蹇来反，内喜之也。"
	yi[38].Xi[3] = "六四：往蹇来连。 象曰：往蹇来连，当位实也。"
	yi[38].Xi[4] = "九五：大蹇朋来。 象曰：大蹇朋来，以中节也。"
	yi[38].Xi[5] = "上六：往蹇来硕，吉，利见大人。 象曰：往蹇来硕，志在内也。利见大人，以从贵也。"

	//解
	yi[39].Tuan = "解：利西南，无所往，其来复吉。有攸往，夙吉。彖曰：解，险以动，动而免乎险，解。解利西南，往得众也。其来复吉，乃得中也。有攸往夙吉，往有功也。天地解，而雷雨作，雷雨作，而百果草木皆甲坼，解之时义大矣哉！\n  象曰：雷雨作，解，君子以赦过宥罪。"
	yi[39].Xi[0] = "初六：无咎。 象曰：刚柔之际，义无咎也。"
	yi[39].Xi[1] = "九二：田获三狐，得黄矢，贞吉。 象曰：九二贞吉，得中道也。"
	yi[39].Xi[2] = "六三：负且乘，致寇至，贞吝。 象曰：负且乘，亦可丑也，自我致戎，又谁咎也。"
	yi[39].Xi[3] = "九四：解而拇，朋至斯孚。 象曰：解而拇，未当位也。"
	yi[39].Xi[4] = "六五：君子维有解，吉，有孚于小人。 象曰：君子有解，小人退也。"
	yi[39].Xi[5] = "上六：公用射隼，于高墉之上，获之，无不利。 象曰：公用射隼，以解悖也。"

	//损
	yi[40].Tuan = "损：有孚，元吉，无咎，可贞，利有攸往？曷之用，二簋可用享。彖曰：损，损下益上，其道上行。损而有孚，元吉，无咎，可贞，利有攸往。曷之用？ 二簋可用享，二簋应有时。损刚益柔有时，损益盈虚，与时偕行。\n  象曰：山下有泽，损，君子以惩忿窒欲。"
	yi[40].Xi[0] = "初九：已事遄往，无咎，酌损之。 象曰：已事 遄往，尚合志也。"
	yi[40].Xi[1] = "九二：利贞，征凶，弗损益之。 象曰：九二利贞，中以为志也。"
	yi[40].Xi[2] = "六三：三人行，则损一人，一人行，则得其友。 象曰：一人行，三则疑也。"
	yi[40].Xi[3] = "六四：损其疾，使遄有喜，无咎。 象曰：损其疾，亦可喜也。"
	yi[40].Xi[4] = "六五：或益之，十朋之龟弗克违，元吉。 象曰：六五元吉，自上佑也。"
	yi[40].Xi[5] = "上九：弗损益之，无咎，贞吉，利有攸往，得臣无家。 象曰：弗损益之，大得志也。"

	//益
	yi[41].Tuan = "益：利有攸往，利涉大川。彖曰：益，损上益下，民说无疆，自上下下，其道大光。利有攸往，中正有庆。利涉大 川，木道乃行。益动而巽，日进无疆。天施地生，其益无方。凡益之道，与时偕行。\n  象曰：风雷，益，君子以见善则迁，有过则改。"
	yi[41].Xi[0] = "初九：利用为大作，元吉，无咎。 象 曰：元吉无咎，下不厚事也。"
	yi[41].Xi[1] = "六二：或益之，十朋之龟弗克违，永贞吉。王用享于帝，吉。 象曰：或益之，自外来 也。"
	yi[41].Xi[2] = "六三：益之用凶事，无咎。有孚中行，告公用圭。 象曰：益用凶事，固有之也。"
	yi[41].Xi[3] = "六四：中行，告公从。利用为依迁国。 象曰：告公从，以益志也。"
	yi[41].Xi[4] = "九五：有孚惠心，勿问元吉。有孚惠我德。 象曰：有孚惠心，勿问之矣。惠我德，大得 志也。"
	yi[41].Xi[5] = "上九：莫益之，或击之，立心勿恒，凶。 象曰：莫益之，偏辞也。或击之，自外来也。"

	//夬
	yi[42].Tuan = "夬：扬于王庭，孚号，有厉，告自邑，不利即戎，利有攸往。彖曰：夬，决也，刚决柔也。健而说，决而和，扬于王庭，柔乘五刚也。孚号有厉， 其危乃光也。告自邑，不利即戎，所尚乃穷也。利有攸往，刚长乃终也。\n  象曰：泽上于天，夬，君子以施禄及下，居德则忌。"
	yi[42].Xi[0] = "初九：壮于前趾，往不胜为吝。  象曰：不胜而往，咎也。"
	yi[42].Xi[1] = "九二：惕号，莫夜有戎，勿恤。 象曰：莫夜有戎，得中道也。"
	yi[42].Xi[2] = "九三：壮于，有凶。君子夬夬，独行遇雨，若濡有愠，无咎。 象曰：君子夬夬，终无咎也。"
	yi[42].Xi[3] = "九四：臀无肤，其行次且。牵羊悔亡，闻言不信。 象曰：其行次且，位不当也。闻言不信，聪不明也。"
	yi[42].Xi[4] = "九五：苋陆夬夬，中行无咎。 象曰：中行无咎，中未光也。"
	yi[42].Xi[5] = "上六：无号，终有凶。 象曰：无号之凶，终不可长也。"

	//姤
	yi[43].Tuan = "姤：女壮，勿用取女。彖曰：姤，遇也，柔遇刚也。勿用取女，不可与长也。天地相遇，品物咸章也。刚遇中正，天下大行也。姤之时义大矣哉！\n  象曰：天下有风，姤，后以施命诰四方。"
	yi[43].Xi[0] = "初六：系于金柅，贞吉，有攸往，见凶， 羸豕踟躅。 象曰：系于金柅，柔道牵也。"
	yi[43].Xi[1] = "九二：包有鱼，无咎，不利宾。 象曰：包有鱼，义不及宾也。"
	yi[43].Xi[2] = "九三：臀无肤，其行次且，厉，无大咎。 象曰：其行次且，行未牵也。"
	yi[43].Xi[3] = "九四：包无鱼，起凶。 象曰：无鱼之凶，远民也。"
	yi[43].Xi[4] = "九五：以杞包瓜，含章，有陨自天。 象曰：九五含章，中正也。有陨自天，志不舍命也。"
	yi[43].Xi[5] = "上九：姤其角，吝，无咎。 象曰：姤其角，上穷吝也。"

	//萃
	yi[44].Tuan = "萃：亨。王假有庙，利见大人，亨，利贞。用大牲吉，利有攸往。彖曰：萃，聚也，顺以说，刚中而应，故聚也。王假有庙，致孝享也。利见大人亨，聚以正也。用大牲吉，利有攸往，顺天命也。观其所聚，而天地万物之情可见矣。\n  象曰：泽上於地，萃，君子以除戎器，戒不虞。"
	yi[44].Xi[0] = "初六：有孚不终，乃乱乃萃，若号一握 为笑，勿恤，往无咎。 象曰：乃乱乃萃，其志乱也。"
	yi[44].Xi[1] = "六二：引吉，无咎，孚乃利用禴。 象曰：引吉无咎，中未变也。"
	yi[44].Xi[2] = "六三：萃如，嗟如，无攸利，往无咎，小吝。 象曰：往无咎，上巽也。"
	yi[44].Xi[3] = "九四：大吉，无咎。 象曰：大吉无咎，位不当也。"
	yi[44].Xi[4] = "九五：萃有位，无咎。匪孚，元永贞，悔亡。 象曰：萃有位，志未光也。"
	yi[44].Xi[5] = "上六：赍咨涕洟，无咎。 象曰：赍咨涕洟，未安上也。"

	//升
	yi[45].Tuan = "升：元亨，用见大人，勿恤，南征吉。彖曰：柔以时升，巽而顺，刚中而应，是以大亨。用见大人，勿恤，有庆也。南征吉， 志行也。\n  象曰：地中生木，升，君子以顺德，积小以高大。"
	yi[45].Xi[0] = "初六：允升，大吉。 象曰：允升大 吉，上合志也。"
	yi[45].Xi[1] = "九二：孚乃利用禴，无咎。 象曰：九二之孚，有喜也。"
	yi[45].Xi[2] = "九三：升虚邑。 象曰：升虚邑，无所疑也。"
	yi[45].Xi[3] = "六四：王用亨于岐山，吉无咎。 象曰：王用亨于岐山，顺事也。"
	yi[45].Xi[4] = "六五：贞吉，升阶。 象曰：贞吉升阶，大得志也。"
	yi[45].Xi[5] = "上六：冥升，利于不息之贞。 象曰：冥升在上，消不富也。"

	//困
	yi[46].Tuan = "困：亨，贞，大人吉，无咎，有言不信。彖曰：困，刚掩也。险以说，困而不失其所，亨，其唯君子乎？贞大人吉，以刚中也。 有言不信，尚口乃穷也。\n  象曰：泽无水，困，君子以致命遂志。"
	yi[46].Xi[0] = "初六：臀困于株木，入于幽谷，三岁不见。 象 曰：入于幽谷，幽不明也。"
	yi[46].Xi[1] = "九二：困于酒食，朱绂方来，利用亨祀，征凶，无咎。 象曰：困于酒食，中有庆也。"
	yi[46].Xi[2] = "六三：困于石，据于蒺藜，入于其宫，不见其妻，凶。 象曰：据于蒺藜，乘刚也。入于 其宫，不见其妻，不祥也。"
	yi[46].Xi[3] = "九四：来徐徐，困于金车，吝，有终。 象曰：来徐徐，志在下也。虽不当位，有与也。"
	yi[46].Xi[4] = "九五：劓刖，困于赤绂，乃徐有说，利用祭祀。 象曰：劓刖，志未得也。乃徐有说，以 中直也。利用祭祀，受福也。"
	yi[46].Xi[5] = "上六：困于葛藟，于臲卼，曰动悔。有悔，征吉。 象曰：困于葛藟，未当也。 动悔，有悔吉，行也。"

	//井
	yi[47].Tuan = "井：改邑不改井，无丧无得，往来井井。汔至，亦未繘井，羸其瓶，凶。彖曰：巽乎水而上水，井，井养而不穷也。改邑不改井，乃以刚中也。汔至亦未繘 井，未有功也。羸其瓶，是以凶也。\n  象曰：木上有水，井，君子以劳民劝相。"
	yi[47].Xi[0] = "初六：井泥不食，旧井无禽。 象曰：井泥不 食，下也。旧井无禽，时舍也。"
	yi[47].Xi[1] = "九二：井谷射鲋，瓮敝漏。 象曰：井谷射鲋，无与也。"
	yi[47].Xi[2] = "九三：井渫不食，为我民恻，可用汲，王明，并受其福。 象曰：井渫不食，行恻也。求 王明，受福也。"
	yi[47].Xi[3] = "六四：井甃，无咎。 象曰：井甃无咎，修井也。"
	yi[47].Xi[4] = "九五：井冽，寒泉食。 象曰：寒泉之食，中正也。"
	yi[47].Xi[5] = "上六：井收勿幕，有孚无吉。 象曰：元吉在上，大成也。"

	//革：己日乃孚，元亨利贞，悔亡。
	yi[48].Tuan = "革：己日乃孚，元亨利贞，悔亡。彖曰：革，水火相息，二女同居，其志不相得，曰革。己日乃孚，革而信也。文明以 说，大亨以正，革而当，其悔乃亡。天地革而四时成，汤武革命，顺乎天而应乎人，革之时 义大矣哉！\n  象曰：泽中有火，革，君子以治历明时。"
	yi[48].Xi[0] = "初九：巩用黄牛之革。 象曰：巩用黄牛，不可 以有为也。"
	yi[48].Xi[1] = "六二：己日乃革之，征吉，无咎。 象曰：己日革之，行有嘉也。"
	yi[48].Xi[2] = "九三：征凶，贞厉，革言三就，有孚。 象曰：革言三就，又何之矣。"
	yi[48].Xi[3] = "九四：悔亡，有孚改命，吉。 象曰：改命之吉，信志也。"
	yi[48].Xi[4] = "九五：大人虎变，未占有孚。 象曰：大人虎变，其文炳也。"
	yi[48].Xi[5] = "上六：君子豹变，小人革面，征凶，居贞吉。 象曰：君子豹变，其文蔚也。小人革面， 顺以从君也。"

	//鼎：元吉，亨
	yi[49].Tuan = "鼎：元吉，亨。彖曰：鼎， 象也。以木巽火，亨饪也。圣人亨以享上帝，而大亨以养圣贤。巽而耳目聪 明，柔进而上行，得中而应乎刚，是以元亨。\n  象曰：木上有火，鼎，君子以正位凝命。"
	yi[49].Xi[0] = "初六：鼎颠趾，利出否，得妾以其子，无咎。  象曰：鼎颠趾，未悖也。利出否，以从贵也。"
	yi[49].Xi[1] = "九二：鼎有实，我仇有疾，不我能即，吉。 象曰：鼎有实，慎所之也。我仇有疾，终无尤也。"
	yi[49].Xi[2] = "九三：鼎耳革，其行塞，雉膏不食，方雨亏悔，终吉。 象曰：鼎耳革，失其义也。"
	yi[49].Xi[3] = "九四：鼎折足，覆公□，其形渥，凶。 象曰：覆公束，信如何也。"
	yi[49].Xi[4] = "六五：鼎黄耳金铉，利贞。 象曰：鼎黄耳，中以为实也。"
	yi[49].Xi[5] = "上九：鼎玉铉，大吉，无不利。 象曰：玉铉在上，刚柔节也。"

	//五十一卦震 震为雷 震上震下
	yi[50].Tuan = "震：亨。震来虩虩，笑言哑哑。震惊百里，不丧匕鬯。彖曰：震，亨。震来虩虩，恐致福也。笑言哑哑，后有则也。震惊百里，惊远而惧 迩也。出可以守宗庙社稷，以为祭主也。\n  象曰：洊雷，震，君子以恐惧修身。"
	yi[50].Xi[0] = "初九：震来虩虩，后笑言哑哑，吉。 象曰： 震来虩虩，恐致福也。笑言哑哑，后有则也。"
	yi[50].Xi[1] = "六二：震来厉，亿丧贝，跻于九陵，勿逐，七日得。 象曰：震来厉，乘刚也。"
	yi[50].Xi[2] = "六三：震苏苏，震行无眚。 象曰：震苏苏，位不当也。"
	yi[50].Xi[3] = "九四：震遂泥。 象曰：震遂泥，未光也。"
	yi[50].Xi[4] = "六五：震往来厉，亿无丧，有事。 象曰：震往来厉，危行也。其事在中，大无丧也。"
	yi[50].Xi[5] = "上六：震索索，视矍矍，征凶。震不于其躬，于其邻，无咎。婚媾有言。 象曰：震索 索，未得中也。虽凶无咎，畏邻戒也。"

	//艮：艮其背，不获其身，行其庭，不见其人，无咎。
	yi[51].Tuan = "艮：艮其背，不获其身，行其庭，不见其人，无咎。彖曰：艮，止也。时止则止，时行则行，动静不失其时，其道光明。艮其止，止其所 也。上下敌应，不相与也。是以不获其身，行其庭不见其人，无咎也。\n  象曰：兼山，艮，君子以思不出其位。	"
	yi[51].Xi[0] = "初六：艮其趾，无咎，利永贞。 象曰：艮其趾， 未失正也。"
	yi[51].Xi[1] = "六二：艮其腓，不拯其随，其心不快。 象曰：不拯其随，未退听也。"
	yi[51].Xi[2] = "九三：艮其限，列其夤，厉薰心。 象曰：艮其限，危薰心也。"
	yi[51].Xi[3] = "六四：艮其身，无咎。 象曰：艮其身，止诸躬也。"
	yi[51].Xi[4] = "六五：艮其辅，言有序，悔亡。 象曰：艮其辅，以中正也。"
	yi[51].Xi[5] = "上九：敦艮，吉。 象曰：敦艮之吉，以厚终也。"

	//渐：女归吉，利贞。
	yi[52].Tuan = "渐：女归吉，利贞。彖曰：渐之进也，女归吉也。进得位，往有功也。进以正，可以正邦也。其位刚，得中 也。止而巽，动不穷也。\n  象曰：山上有木，渐，君子以居贤德，善俗。"
	yi[52].Xi[0] = "初六：鸿渐于干，小子厉，有言，无咎。  象曰：小子之厉，义无咎也。"
	yi[52].Xi[1] = "六二：鸿渐于磐，饮食衍衍，吉。 象曰：饮食衍衍，不素饱也。"
	yi[52].Xi[2] = "九三：鸿渐于陆，夫征不复，妇孕不育，凶，利御寇。 象曰：夫征不复，离群丑也。妇 孕不育，失其道也。利用御寇，顺相保也。"
	yi[52].Xi[3] = "六四：鸿渐于木，或得其桷，无咎。 象曰：或得其桷，顺以巽也。"
	yi[52].Xi[4] = "九五：鸿渐于陵，妇三岁不孕，终莫之胜，吉。 象曰：终莫之胜，吉，得所愿也。"
	yi[52].Xi[5] = "上九：鸿渐于逵，其羽可用为仪，吉。 象曰：其羽可用为仪，吉，不可乱也。"

	//归妹：征凶，无攸利。
	yi[53].Tuan = "归妹：征凶，无攸利。彖曰：归妹，天地之大义也。天地不交，而万物不兴，归妹人之终始也。说以动，所归 妹也。征凶，位不当也。无攸利，柔乘刚也。\n  象曰：泽上有雷，归妹，君子以永终知敝。"
	yi[53].Xi[0] = "初九：归妹以娣，跛能履，征吉。 象曰：归 妹以娣，以恒也。跛能履吉，相承也。"
	yi[53].Xi[1] = "九二：眇能视，利幽人之贞。 象曰：利幽人之贞，未变常也。"
	yi[53].Xi[2] = "六三：归妹以须，反归以娣。 象曰：归妹以须，未当也。"
	yi[53].Xi[3] = "九四：归妹愆期，迟归有时。 象曰：愆期之志，有待而行也。"
	yi[53].Xi[4] = "六五：帝乙归妹，其君之袂，不如其娣之袂良，月几望，吉。 象曰：帝乙归妹，不如其娣之袂良也。其位在中，以贵行也。"
	yi[53].Xi[5] = "上六：女承筐无实，士刲羊无血，无攸利。 象曰：上六无实，承虚筐也。"

	//丰：亨，王假之，勿忧，宜日中。
	yi[54].Tuan = "丰：亨，王假之，勿忧，宜日中。彖曰：丰，大也。明以动，故丰。王假之，尚大也。勿忧宜日中，宜照天下也。日中则 昃，月盈则食，天地盈虚，与时消息，而况人於人乎？况於鬼神乎？\n  象曰：雷电皆至，丰，君子以折狱致刑。"
	yi[54].Xi[0] = "初九：遇其配主，虽旬无咎，往有尚。 象曰： 虽旬无咎，过旬灾也。"
	yi[54].Xi[1] = "六二：丰其蔀，日中见斗，往得疑疾，有孚发若，吉。 象曰：有孚发若，信以发志 也。"
	yi[54].Xi[2] = "九三：丰其沛，日中见昧，折其右肱，无咎。 象曰：丰其沛，不可大事也。折其右肱， 终不可用也。"
	yi[54].Xi[3] = "九四：丰其蔀，日中见斗，遇其夷主，吉。 象曰：丰其蔀，位不当也。日中见斗， 幽不明也。遇其夷主，吉，行也。"
	yi[54].Xi[4] = "六五：来章，有庆誉，吉。 象曰：六五之吉，有庆也。"
	yi[54].Xi[5] = "上六：丰其屋，蔀其家，窥其户，阒其无人，三岁不见，凶。 象曰：丰其屋，天际 翔也。窥其户，阒其无人，自藏也。"

	//旅：小亨，旅贞吉。
	yi[55].Tuan = "旅：小亨，旅贞吉。彖曰：旅，小亨，柔得中乎外，而顺乎刚，止而丽乎明，是以小亨，旅贞吉也。旅之时 义大矣哉！\n  象曰：山上有火，旅，君子以明慎用刑，而不留狱。"
	yi[55].Xi[0] = "初六：旅琐琐，斯其所取灾。 象 曰：旅琐琐，志穷灾也。"
	yi[55].Xi[1] = "六二：旅即次，怀其资，得童仆贞。 象曰：得童仆贞，终无尤也。"
	yi[55].Xi[2] = "九三：旅焚其次，丧其童仆，贞厉。 象曰：旅焚其次，亦以伤矣。以旅与下，其义丧也。"
	yi[55].Xi[3] = "九四：旅于处，得其资斧，我心不快。 象曰：旅于处，未得位也。得其资斧，心未快也。"
	yi[55].Xi[4] = "六五：射雉一矢亡，终以誉命。 象曰：终以誉命，上逮也。"
	yi[55].Xi[5] = "上九：鸟焚其巢，旅人先笑后号啕。丧牛于易，凶。 象曰：以旅在上，其义焚也。丧牛 于易，终莫之闻也。"

	//巽：小亨，利攸往，利见大人。
	yi[56].Tuan = "巽：小亨，利攸往，利见大人。彖曰：重巽以申命，刚巽乎中正而志行。柔皆顺乎刚，是以小亨，利有攸往，利见大人。\n  象曰：随风，巽，君子以申命行事。"
	yi[56].Xi[0] = "初六：进退，利武人之贞。 象曰：进退，志疑也。 利武人之贞，志治也。"
	yi[56].Xi[1] = "九二：巽在床下。用史巫纷若，吉无咎。 象曰：纷若之吉，得中也。"
	yi[56].Xi[2] = "九三：频巽，吝。 象曰：频巽之吝，志穷也。"
	yi[56].Xi[3] = "六四：悔亡，田获三品。 象曰：田获三品，有功也。"
	yi[56].Xi[4] = "九五：贞吉悔亡，无不利。无初有终，先庚三日，后庚三日，吉。 象曰：九五之吉，位 正中也。"
	yi[56].Xi[5] = "上九：巽在床下，丧其资斧，贞凶。 象曰：巽在□下，上穷也。丧其资斧，正乎凶也。"

	//兑：亨，利贞。
	yi[57].Tuan = "兑：亨，利贞。彖曰：兑，说也。刚中而柔外，说以利贞，是以顺乎天，而应乎人。说以先民，民忘其 劳，说以犯难，民忘其死，说之大，民劝矣哉！\n  象曰：丽泽，兑，君子以朋友讲习。"
	yi[57].Xi[0] = "初九：和兑，吉。 象曰：和兑之吉，行未疑也。"
	yi[57].Xi[1] = "九二：孚兑，吉，悔亡。 象曰：孚兑之吉，信志也。"
	yi[57].Xi[2] = "六三：来兑，凶。 象曰：来兑之凶，位不当也。"
	yi[57].Xi[3] = "九四：商兑，未宁，介疾有喜。 象曰：九四之喜，有庆也。"
	yi[57].Xi[4] = "九五：孚于剥，有厉。 象曰：孚于剥，位正当也。"
	yi[57].Xi[5] = "上六：引兑。 象曰：上六引兑，未光也。"

	//涣：亨。王假有庙，利涉大川，利贞。
	yi[58].Tuan = "涣：亨。王假有庙，利涉大川，利贞。彖曰：涣，亨。刚来而不穷，柔得位乎外而上同。王假有庙，王乃在中也。利涉大川， 乘木有功也。\n  象曰：风行水上，涣，先王以享于帝立庙。"
	yi[58].Xi[0] = "初六：用拯马壮，吉。 象曰：初六之吉，顺 也。"
	yi[58].Xi[1] = "九二：涣奔其机，悔亡。 象曰：涣奔其机，得愿也。"
	yi[58].Xi[2] = "六三：涣其躬，无悔。 象曰：涣其躬，志在外也。"
	yi[58].Xi[3] = "六四：涣其群，元吉。涣有丘，匪夷所思。 象曰：涣其群，元吉，光大也。"
	yi[58].Xi[4] = "九五：涣汗其大号，涣王居，无咎。 象曰：王居无咎，正位也。"
	yi[58].Xi[5] = "上九：涣其血，去逖出，无咎。 象曰：涣其血，远害也。"

	//节：亨。苦节不可贞。
	yi[59].Tuan = "节：亨。苦节不可贞。彖曰：节，亨，刚柔分，而刚得中。苦节不可贞，其道穷也。说以行险，当位以节，中 正以通。天地节而四时成，节以制度，不伤财，不害民。\n  象曰：泽上有水，节，君子以制数度，议德行。"
	yi[59].Xi[0] = "初九：不出户庭，无咎。 象曰：不出户 庭，知通塞也。"
	yi[59].Xi[1] = "九二：不出门庭，凶。 象曰：不出门庭，失时极也。"
	yi[59].Xi[2] = "六三：不节若，则嗟若，无咎。 象曰：不节之嗟，又谁咎也。"
	yi[59].Xi[3] = "六四：安节，亨。 象曰：安节之亨，承上道也。"
	yi[59].Xi[4] = "九五：甘节，吉，往有尚。 象曰：甘节之吉，居位中也。"
	yi[59].Xi[5] = "上六：苦节，贞凶，悔亡。 象曰：苦节贞凶，其道穷也。"

	//中孚：豚鱼吉，利涉大川，利贞。
	yi[60].Tuan = "中孚：豚鱼吉，利涉大川，利贞。彖曰：中孚，柔在内而刚得中。说而巽，孚，乃化邦也。豚鱼吉，信及豚鱼也。利涉大 川，乘木舟虚也。中孚以利贞，乃应乎天也。\n  象曰：泽上有风，中孚，君子以议狱缓死。"
	yi[60].Xi[0] = "初九：虞吉，有他不燕。 象曰：初九虞吉， 志未变也。"
	yi[60].Xi[1] = "九二：鸣鹤在阴，其子和之，我有好爵，吾与尔靡之。 象曰：其子和之，中心愿也。"
	yi[60].Xi[2] = "六三：得敌，或鼓或罢，或泣或歌。 象曰：可鼓或罢，位不当也。"
	yi[60].Xi[3] = "六四：月几望，马匹亡，无咎。 象曰：马匹亡，绝类上也。"
	yi[60].Xi[4] = "九五：有孚挛如，无咎。 象曰：有孚挛如，位正当也。"
	yi[60].Xi[5] = "上九：翰音登于天，贞凶。 象曰：翰音登于天，何可长也。"

	//小过：亨，利贞，可小事，不可大事。飞鸟遗之音，不宜上宜下，大吉。
	yi[61].Tuan = "小过：亨，利贞，可小事，不可大事。飞鸟遗之音，不宜上宜下，大吉。彖曰：小过，小者过而亨也。过以利贞，与时行也。柔得中，是以小事吉也。刚失位而 不中，是以不可大事也。有飞鸟之 象焉，有飞鸟遗之音，不宜上宜下，大吉，上逆而下顺 也。\n  象曰：山上有雷，小过，君子以行过乎恭，丧过乎哀，用过乎俭。"
	yi[61].Xi[0] = "初六：飞鸟以凶。 象 曰：飞鸟以凶，不可如何也。"
	yi[61].Xi[1] = "六二：过其祖，遇其妣，不及其君，遇其臣，无咎。 象曰：不及其君，臣不可过也。"
	yi[61].Xi[2] = "九三：弗过防之，从或戕之，凶。 象曰：从或戕之，凶如何也。"
	yi[61].Xi[3] = "九四：无咎，弗过遇之。往厉必戒，勿用永贞。 象曰：弗过遇之，位不当也。往厉必 戒，终不可长也。"
	yi[61].Xi[4] = "六五：密云不雨，自我西郊，公弋取彼在穴。 象曰：密云不雨，已上也。"
	yi[61].Xi[5] = "上六：弗遇过之，飞鸟离之，凶，是谓灾眚。 象曰：弗遇过之，已亢也。"

	//既济：亨小，利贞，初吉终乱。
	yi[62].Tuan = "既济：亨小利贞，初吉终乱。彖曰：既济，亨，小者亨也。利贞，刚柔正而位当也。初吉，柔得中也。终止则乱，其道穷也。\n  象曰：水在火上，既济，君子以思患而预防之。"
	yi[62].Xi[0] = "初九：曳其轮，濡其尾，无咎。 象曰： 曳其轮，义无咎也。"
	yi[62].Xi[1] = "六二：妇丧其茀，勿逐，七日得。 象曰：七日得，以中道也。"
	yi[62].Xi[2] = "九三：高宗伐鬼方，三年克之，小人勿用。 象曰：三年克之，惫也。"
	yi[62].Xi[3] = "六四：𦈡有衣袽，终日戒。 象曰：终日戒，有所疑也。"
	yi[62].Xi[4] = "九五：东邻杀牛，不如西邻之禴祭，实受其福。 象曰：东邻杀牛，不如西邻之时也， 实受其福，吉大来也。"
	yi[62].Xi[5] = "上六：濡其首，厉。 象曰：濡其首厉，何可久也。"

	//未济：亨，小狐汔济，濡其尾，无攸利。
	yi[63].Tuan = "未济：亨，小狐汔济，濡其尾，无攸利。彖曰：未济，亨，柔得中也。小狐汔济，未出中也。濡其尾，无攸利，不续终也。虽不 当位，刚柔应也。\n  象曰：火在水上，未济，君子以慎辨物居方。"
	yi[63].Xi[0] = "初六：濡其尾，吝。 象曰：濡其尾，亦不 知极也。"
	yi[63].Xi[1] = "九二：曳其轮，贞吉。 象曰：九二贞吉，中以行正也。"
	yi[63].Xi[2] = "六三：未济，征凶，利涉大川。 象曰：未济征凶，位不当也。"
	yi[63].Xi[3] = "九四：贞吉，悔亡，震用伐鬼方，三年有赏于大国。 象曰：贞吉悔亡，志行也。"
	yi[63].Xi[4] = "六五：贞吉，无悔，君子之光，有孚，吉。 象曰：君子之光，其晖吉也。"
	yi[63].Xi[5] = "上九：有孚于饮酒，无咎，濡其首，有孚失是。 象曰：饮酒濡首，亦不知节也"

	return yi

}

// ReGuaNu 返回卦 序号
func ReGuaNu(m, n, l []uint, f *os.File, yijing []Gua) uint {
	var a uint

	switch {
	//乾卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 1

	//坤卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 2

	//屯卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 3
	//蒙卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 4

	//需卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 5

	//讼卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 6

	//师卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 7
	//比卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 8
	//小畜卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 9

		//履卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 10

		//泰卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 11

		//否卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 12
		//同人卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 13

		//大有卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 14

		//谦卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 15

		//豫卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 16

	//随卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 17
		//蛊卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 18

		//临卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 19
		//观卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 20
		//噬嗑卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 21

		//贲卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 22
		//剥卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 23
		//复卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 24

		//无妄卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 25

		//大畜卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 26
		//颐卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 27

		//大过卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 28

		//坎卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 29
		//离卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 30
		// .....下经......
		//咸卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 31
		//恒卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 32

		//遁卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 33
		//大壮卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 34
		//晋卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 35

		//明夷卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 36
		//家人卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 37
		//睽卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 38

		//蹇卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 39

		//解卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 40
		//损卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 41

		//益卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 42

		//夬卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 43

		//姤卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 1:
		a = 44
		//萃卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 45
		//升卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 0:
		a = 46

		//困卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 47
		//井卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 48

		//革卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 49

		//鼎卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 50

		//震卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 51

		//艮卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 0 && l[5] == 1:
		a = 52

		//渐卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 53

		//归妹卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 54

		//丰卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 55

		//旅卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 56

		//巽卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 57
		//兑卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 1 && l[5] == 0:
		a = 58

		//涣卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 59

		//节卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 60

		//中孚卦 判断
	case l[0] == 1 && l[1] == 1 && l[2] == 0 && l[3] == 0 && l[4] == 1 && l[5] == 1:
		a = 61

		//小过卦 判断
	case l[0] == 0 && l[1] == 0 && l[2] == 1 && l[3] == 1 && l[4] == 0 && l[5] == 0:
		a = 62

		//既济卦 判断
	case l[0] == 1 && l[1] == 0 && l[2] == 1 && l[3] == 0 && l[4] == 1 && l[5] == 0:
		a = 63
		//未济卦 判断
	case l[0] == 0 && l[1] == 1 && l[2] == 0 && l[3] == 1 && l[4] == 0 && l[5] == 1:
		a = 64
	}
	return a
}
