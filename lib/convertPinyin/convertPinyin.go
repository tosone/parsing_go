package convertPinyin

import (
	"strings"
)

var chars map[string]string

func init() {
	chars = map[string]string{
		"a":      "啊阿",
		"ai":     "埃挨哎唉哀皑癌蔼矮艾碍爱隘",
		"an":     "鞍氨安俺按暗岸胺案",
		"ang":    "肮昂盎",
		"ao":     "凹敖熬翱袄傲奥懊澳",
		"ba":     "芭捌扒叭吧笆八疤巴拔跋靶把耙坝霸罢爸",
		"bai":    "白柏百摆佰败拜稗",
		"ban":    "斑班搬扳般颁板版扮拌伴瓣半办绊",
		"bang":   "邦帮梆榜膀绑棒磅蚌镑傍谤",
		"bao":    "苞胞包褒剥盄盇盉盋盌盓盕盙盚盜盝盞盠盡盢監盤盦盧盨盩盪盫盬盭盰盳盵盶盷盺盻盽盿眀眂眃眅眆眊県眎眏眐眑眒眓眔眕眖眗眘眛眜眝眞眡眣眤眥眧眪眫眬眮眰眱眲眳眴眹眻眽眾眿睂睄睅睆睈睉睊睋睌睍睎睏睒睓睔睕睖睗睘睙睜薄雹保堡饱宝抱报暴豹鲍爆",
		"bei":    "杯碑悲卑北辈背贝钡倍狈备惫焙被",
		"ben":    "奔苯本笨",
		"beng":   "崩绷甭泵蹦迸",
		"bi":     "逼鼻比鄙笔彼碧蓖蔽毕毙毖币庇痹闭敝弊必辟壁臂避陛",
		"bian":   "鞭边编贬扁便变卞辨辩辫遍",
		"biao":   "标彪膘表",
		"bie":    "鳖憋别瘪",
		"bin":    "彬斌濒滨宾摈",
		"bing":   "兵冰柄丙秉饼炳　睝睞睠睤睧睩睪睭睮睯睰睱睲睳睴睵睶睷睸睺睻睼瞁瞂瞃瞆瞇瞈瞉瞊瞋瞏瞐瞓瞔瞕瞖瞗瞘瞙瞚瞛瞜瞝瞞瞡瞣瞤瞦瞨瞫瞭瞮瞯瞱瞲瞴瞶瞷瞸瞹瞺瞼瞾矀矁矂矃矄矅矆矇矈矉矊矋矌矎矏矐矑矒矓矔矕矖矘矙矚矝矞矟矠矡矤病并",
		"bo":     "玻菠播拨钵波博勃搏铂箔伯帛舶脖膊渤泊驳捕卜",
		"bu":     "哺补埠不布步簿部怖",
		"ca":     "擦",
		"cai":    "猜裁材才财睬踩采彩菜蔡",
		"can":    "餐参蚕残惭惨灿",
		"cang":   "苍舱仓沧藏",
		"cao":    "操糙槽曹草",
		"ce":     "厕策侧册测",
		"ceng":   "层蹭",
		"cha":    "插叉茬茶查碴搽察岔差诧",
		"chai":   "拆柴豺",
		"chan":   "搀掺蝉馋谗缠铲产阐颤",
		"chang":  "昌猖　矦矨矪矰矱矲矴矵矷矹矺矻矼砃砄砅砆砇砈砊砋砎砏砐砓砕砙砛砞砠砡砢砤砨砪砫砮砯砱砲砳砵砶砽砿硁硂硃硄硆硈硉硊硋硍硏硑硓硔硘硙硚硛硜硞硟硠硡硢硣硤硥硦硧硨硩硯硰硱硲硳硴硵硶硸硹硺硻硽硾硿碀碁碂碃场尝常长偿肠厂敞畅唱倡",
		"chao":   "超抄钞朝嘲潮巢吵炒",
		"che":    "车扯撤掣彻澈",
		"chen":   "郴臣辰尘晨忱沉陈趁衬",
		"cheng":  "撑称城橙成呈乘程惩澄诚承逞骋秤",
		"chi":    "吃痴持匙池迟弛驰耻齿侈尺赤翅斥炽",
		"chong":  "充冲虫崇宠",
		"chou":   "抽酬畴踌稠愁筹仇绸瞅丑臭",
		"chu":    "初出橱厨躇锄雏滁除楚碄碅碆碈碊碋碏碐碒碔碕碖碙碝碞碠碢碤碦碨碩碪碫碬碭碮碯碵碶碷碸確碻碼碽碿磀磂磃磄磆磇磈磌磍磎磏磑磒磓磖磗磘磚磛磜磝磞磟磠磡磢磣磤磥磦磧磩磪磫磭磮磯磰磱磳磵磶磸磹磻磼磽磾磿礀礂礃礄礆礇礈礉礊礋礌础储矗搐触处",
		"chuai":  "揣",
		"chuan":  "川穿椽传船喘串",
		"chuang": "疮窗幢床闯创",
		"chui":   "吹炊捶锤垂",
		"chun":   "春椿醇唇淳纯蠢",
		"chuo":   "戳绰",
		"ci":     "疵茨磁雌辞慈瓷词此刺赐次",
		"cong":   "聪葱囱匆从丛",
		"cou":    "凑",
		"cu":     "粗醋簇促",
		"cuan":   "蹿篡窜",
		"cui":    "摧崔催脆瘁粹淬翠",
		"cun":    "村存寸",
		"cuo":    "磋撮搓措挫错",
		"da":     "搭达答瘩打大",
		"dai":    "呆歹傣戴带殆代贷袋待逮礍礎礏礐礑礒礔礕礖礗礘礙礚礛礜礝礟礠礡礢礣礥礦礧礨礩礪礫礬礭礮礯礰礱礲礳礵礶礷礸礹礽礿祂祃祄祅祇祊祋祌祍祎祏祐祑祒祔祕祘祙祡祣祤祦祩祪祫祬祮祰祱祲祳祴祵祶祹祻祼祽祾祿禂禃禆禇禈禉禋禌禍禎禐禑禒怠",
		"dan":    "耽担丹单郸掸胆旦氮但惮淡诞弹蛋",
		"dang":   "当挡党荡档",
		"dao":    "刀捣蹈倒岛祷导到稻悼道盗",
		"de":     "德得的",
		"deng":   "蹬灯登等瞪凳邓",
		"di":     "堤低滴迪敌笛狄涤翟嫡抵底地蒂第帝弟递缔",
		"dian":   "颠掂滇碘点典靛垫电佃甸店惦奠淀殿",
		"diao":   "碉叼雕凋刁掉吊钓调",
		"die":    "跌爹碟蝶迭谍叠禓禔禕禖禗禘禙禛禜禝禞禟禠禡禢禣禤禥禦禨禩禪禫禬禭禮禯禰禱禲禴禵禶禷禸禼禿秂秄秅秇秈秊秌秎秏秐秓秔秖秗秙秚秛秜秝秞秠秡秢秥秨秪秬秮秱秲秳秴秵秶秷秹秺秼秾秿稁稄稅稇稈稉稊稌稏稐稑稒稓稕稖稘稙稛稜",
		"ding":   "丁盯叮钉顶鼎锭定订",
		"diu":    "丢",
		"dong":   "东冬董懂动栋侗恫冻洞",
		"dou":    "兜抖斗陡豆逗痘",
		"du":     "都督毒犊独读堵睹赌杜镀肚度渡妒",
		"duan":   "端短锻段断缎",
		"dui":    "堆兑队对",
		"dun":    "墩吨蹲敦顿囤钝盾遁",
		"duo":    "掇哆多夺垛躲朵跺舵剁惰堕",
		"e":      "蛾峨鹅俄额讹娥恶厄扼遏鄂饿",
		"en":     "恩",
		"er":     "而儿耳尔饵洱二稝稟稡稢稤稥稦稧稨稩稪稫稬稭種稯稰稱稲稴稵稶稸稺稾穀穁穂穃穄穅穇穈穉穊穋穌積穎穏穐穒穓穔穕穖穘穙穚穛穜穝穞穟穠穡穢穣穤穥穦穧穨穩穪穫穬穭穮穯穱穲穳穵穻穼穽穾窂窅窇窉窊窋窌窎窏窐窓窔窙窚窛窞窡窢贰",
		"fa":     "发罚筏伐乏阀法珐",
		"fan":    "藩帆番翻樊矾钒繁凡烦反返范贩犯饭泛",
		"fang":   "坊芳方肪房防妨仿访纺放",
		"fei":    "菲非啡飞肥匪诽吠肺废沸费",
		"fen":    "芬酚吩氛分纷坟焚汾粉奋份忿愤粪",
		"feng":   "丰封枫蜂峰锋风疯烽逢冯缝讽奉凤",
		"fo":     "佛",
		"fou":    "否",
		"fu":     "夫敷肤孵扶拂辐幅氟符伏俘服窣窤窧窩窪窫窮窯窰窱窲窴窵窶窷窸窹窺窻窼窽窾竀竁竂竃竄竅竆竇竈竉竊竌竍竎竏竐竑竒竓竔竕竗竘竚竛竜竝竡竢竤竧竨竩竪竫竬竮竰竱竲竳竴竵競竷竸竻竼竾笀笁笂笅笇笉笌笍笎笐笒笓笖笗笘笚笜笝笟笡笢笣笧笩笭浮涪福袱弗甫抚辅俯釜斧脯腑府腐赴副覆赋复傅付阜父腹负富讣附妇缚咐",
		"ga":     "噶嘎",
		"gai":    "该改概钙盖溉",
		"gan":    "干甘杆柑竿肝赶感秆敢赣",
		"gang":   "冈刚钢缸肛纲岗港杠",
		"gao":    "篙皋高膏羔糕搞镐稿告",
		"ge":     "哥歌搁戈鸽胳疙割革葛格蛤阁隔铬个各",
		"gei":    "给",
		"gen":    "根跟",
		"geng":   "耕更庚羹笯笰笲笴笵笶笷笹笻笽笿筀筁筂筃筄筆筈筊筍筎筓筕筗筙筜筞筟筡筣筤筥筦筧筨筩筪筫筬筭筯筰筳筴筶筸筺筼筽筿箁箂箃箄箆箇箈箉箊箋箌箎箏箑箒箓箖箘箙箚箛箞箟箠箣箤箥箮箯箰箲箳箵箶箷箹箺箻箼箽箾箿節篂篃範埂耿梗",
		"gong":   "工攻功恭龚供躬公宫弓巩汞拱贡共",
		"gou":    "钩勾沟苟狗垢构购够",
		"gu":     "辜菇咕箍估沽孤姑鼓古蛊骨谷股故顾固雇",
		"gua":    "刮瓜剐寡挂褂",
		"guai":   "乖拐怪",
		"guan":   "棺关官冠观管馆罐惯灌贯",
		"guang":  "光广逛",
		"gui":    "瑰规圭硅归龟闺轨鬼诡癸桂柜跪贵刽",
		"gun":    "辊滚棍",
		"guo":    "锅郭国果裹过",
		"ha":     "哈篅篈築篊篋篍篎篏篐篒篔篕篖篗篘篛篜篞篟篠篢篣篤篧篨篩篫篬篭篯篰篲篳篴篵篶篸篹篺篻篽篿簀簁簂簃簄簅簆簈簉簊簍簎簐簑簒簓簔簕簗簘簙簚簛簜簝簞簠簡簢簣簤簥簨簩簫簬簭簮簯簰簱簲簳簴簵簶簷簹簺簻簼簽簾籂",
		"hai":    "骸孩海氦亥害骇",
		"han":    "酣憨邯韩含涵寒函喊罕翰撼捍旱憾悍焊汗汉",
		"hang":   "夯杭航",
		"hao":    "壕嚎豪毫郝好耗号浩",
		"he":     "呵喝荷菏核禾和何合盒貉阂河涸赫褐鹤贺",
		"hei":    "嘿黑",
		"hen":    "痕很狠恨",
		"heng":   "哼亨横衡恒",
		"hong":   "轰哄烘虹鸿洪宏弘红",
		"hou":    "喉侯猴吼厚候后",
		"hu":     "呼乎忽瑚壶葫胡蝴狐糊湖籃籄籅籆籇籈籉籊籋籌籎籏籐籑籒籓籔籕籖籗籘籙籚籛籜籝籞籟籠籡籢籣籤籥籦籧籨籩籪籫籬籭籮籯籰籱籲籵籶籷籸籹籺籾籿粀粁粂粃粄粅粆粇粈粊粋粌粍粎粏粐粓粔粖粙粚粛粠粡粣粦粧粨粩粫粬粭粯粰粴粵粶粷粸粺粻弧虎唬护互沪户",
		"hua":    "花哗华猾滑画划化话",
		"huai":   "槐徊怀淮坏",
		"huan":   "欢环桓还缓换患唤痪豢焕涣宦幻",
		"huang":  "荒慌黄磺蝗簧皇凰惶煌晃幌恍谎",
		"hui":    "灰挥辉徽恢蛔回毁悔慧卉惠晦贿秽会烩汇讳诲绘",
		"hun":    "荤昏婚魂浑混",
		"huo":    "豁活伙火获或惑霍货祸",
		"ji":     "击圾基机畸稽积箕粿糀糂糃糄糆糉糋糎糏糐糑糒糓糔糘糚糛糝糞糡糢糣糤糥糦糧糩糪糫糬糭糮糰糱糲糳糴糵糶糷糹糺糼糽糾糿紀紁紂紃約紅紆紇紈紉紋紌納紎紏紐紑紒紓純紕紖紗紘紙級紛紜紝紞紟紡紣紤紥紦紨紩紪紬紭紮細紱紲紳紴紵紶肌饥迹激讥鸡姬绩缉吉极棘辑籍集及急疾汲即嫉级挤几脊己蓟技冀季伎祭剂悸济寄寂计记既忌际妓继纪",
		"jia":    "嘉枷夹佳家加荚颊贾甲钾假稼价架驾嫁",
		"jian":   "歼监坚尖笺间煎兼肩艰奸缄茧检柬碱硷拣捡简俭剪减荐槛鉴践贱见键箭件　紷紸紹紺紻紼紾紿絀絁終絃組絅絆絇絈絉絊絋経絍絎絏結絑絒絓絔絕絖絗絘絙絚絛絜絝絞絟絠絡絢絣絤絥給絧絨絩絪絫絬絭絯絰統絲絳絴絵絶絸絹絺絻絼絽絾絿綀綁綂綃綄綅綆綇綈綉綊綋綌綍綎綏綐綑綒經綔綕綖綗綘健舰剑饯渐溅涧建",
		"jiang":  "僵姜将浆江疆蒋桨奖讲匠酱降",
		"jiao":   "蕉椒礁焦胶交郊浇骄娇嚼搅铰矫侥脚狡角饺缴绞剿教酵轿较叫窖",
		"jie":    "揭接皆秸街阶截劫节桔杰捷睫竭洁结解姐戒藉芥界借介疥诫届",
		"jin":    "巾筋斤金今津襟紧锦仅谨进靳晋禁近烬浸継続綛綜綝綞綟綠綡綢綣綤綥綧綨綩綪綫綬維綯綰綱網綳綴綵綶綷綸綹綺綻綼綽綾綿緀緁緂緃緄緅緆緇緈緉緊緋緌緍緎総緐緑緒緓緔緕緖緗緘緙線緛緜緝緞緟締緡緢緣緤緥緦緧編緩緪緫緬緭緮緯緰緱緲緳練緵緶緷緸緹緺尽劲",
		"jing":   "荆兢茎睛晶鲸京惊精粳经井警景颈静境敬镜径痉靖竟竞净",
		"jiong":  "炯窘",
		"jiu":    "揪究纠玖韭久灸九酒厩救旧臼舅咎就疚",
		"ju":     "鞠拘狙疽居驹菊局咀矩举沮聚拒据巨具距踞锯俱句惧炬剧",
		"juan":   "捐鹃娟倦眷卷绢",
		"jue":    "撅攫抉掘倔爵觉决诀绝",
		"jun":    "均菌钧军君峻緻緼緽緾緿縀縁縂縃縄縅縆縇縈縉縊縋縌縍縎縏縐縑縒縓縔縕縖縗縘縙縚縛縜縝縞縟縠縡縢縣縤縥縦縧縨縩縪縫縬縭縮縯縰縱縲縳縴縵縶縷縸縹縺縼總績縿繀繂繃繄繅繆繈繉繊繋繌繍繎繏繐繑繒繓織繕繖繗繘繙繚繛繜繝俊竣浚郡骏",
		"ka":     "喀咖卡咯",
		"kai":    "开揩楷凯慨",
		"kan":    "刊堪勘坎砍看",
		"kang":   "康慷糠扛抗亢炕",
		"kao":    "考拷烤靠",
		"ke":     "坷苛柯棵磕颗科壳咳可渴克刻客课",
		"ken":    "肯啃垦恳",
		"keng":   "坑吭",
		"kong":   "空恐孔控",
		"kou":    "抠口扣寇",
		"ku":     "枯哭窟苦酷库裤",
		"kua":    "夸垮挎跨胯",
		"kuai":   "块筷侩快",
		"kuan":   "宽款",
		"kuang":  "匡筐狂框矿眶旷况",
		"kui":    "亏盔岿窥葵奎魁傀繞繟繠繡繢繣繤繥繦繧繨繩繪繫繬繭繮繯繰繱繲繳繴繵繶繷繸繹繺繻繼繽繾繿纀纁纃纄纅纆纇纈纉纊纋續纍纎纏纐纑纒纓纔纕纖纗纘纙纚纜纝纞纮纴纻纼绖绤绬绹缊缐缞缷缹缻缼缽缾缿罀罁罃罆罇罈罉罊罋罌罍罎罏罒罓馈愧溃",
		"kun":    "坤昆捆困",
		"kuo":    "括扩廓阔",
		"la":     "垃拉喇蜡腊辣啦",
		"lai":    "莱来赖",
		"lan":    "蓝婪栏拦篮阑兰澜谰揽览懒缆烂滥",
		"lang":   "琅榔狼廊郎朗浪",
		"lao":    "捞劳牢老佬姥酪烙涝",
		"le":     "勒乐了",
		"lei":    "雷镭蕾磊累儡垒擂肋类泪",
		"leng":   "棱楞冷",
		"li":     "厘梨犁黎篱狸离漓理李里鲤礼莉荔吏栗丽厉励砾历利傈例俐罖罙罛罜罝罞罠罣罤罥罦罧罫罬罭罯罰罳罵罶罷罸罺罻罼罽罿羀羂羃羄羅羆羇羈羉羋羍羏羐羑羒羓羕羖羗羘羙羛羜羠羢羣羥羦羨義羪羫羬羭羮羱羳羴羵羶羷羺羻羾翀翂翃翄翆翇翈翉翋翍翏翐翑習翓翖翗翙翚翛翜翝翞翢翣痢立粒沥隶力璃哩",
		"lia":    "俩",
		"lian":   "联莲连镰廉怜涟帘敛脸链恋炼练",
		"liang":  "粮凉梁粱良两辆量晾亮谅",
		"liao":   "撩聊僚疗燎寥辽潦撂镣廖料",
		"lie":    "列裂烈劣猎",
		"lin":    "琳林磷霖临邻鳞淋凛赁吝拎",
		"ling":   "玲菱零龄铃伶羚凌灵陵岭领另令",
		"liu":    "溜琉榴硫馏留刘瘤流柳六",
		"long":   "龙聋咙笼窿　翤翧翨翪翫翬翭翲翴翵翶翷翸翹翺翽翾翿耂耇耈耉耊耎耏耑耓耚耛耝耞耟耡耣耤耫耬耭耮耯耰耲耴耹耺耼耾聀聁聄聅聇聈聉聎聏聐聑聓聕聖聗聙聛聜聝聞聟聠聡聢聣聤聥聦聧聨聫聬聭聮聯聰聲聳聴聵聶職聸聹聺聻聼聽隆垄拢陇",
		"lou":    "楼娄搂篓漏陋",
		"lu":     "芦卢颅庐炉掳卤虏鲁麓碌露路赂鹿潞禄录陆戮",
		"lv":     "驴吕铝侣旅履屡缕虑氯律率滤绿",
		"luan":   "峦挛孪滦卵乱",
		"lue":    "掠略",
		"lun":    "抡轮伦仑沦纶论",
		"luo":    "萝螺罗逻锣箩骡裸落洛骆络",
		"ma":     "妈麻玛码蚂马骂嘛吗",
		"mai":    "埋买麦卖迈脉",
		"man":    "瞒馒蛮满蔓曼慢漫聾肁肂肅肈肊肍肎肏肐肑肒肔肕肗肙肞肣肦肧肨肬肰肳肵肶肸肹肻胅胇胈胉胊胋胏胐胑胒胓胔胕胘胟胠胢胣胦胮胵胷胹胻胾胿脀脁脃脄脅脇脈脋脌脕脗脙脛脜脝脟脠脡脢脣脤脥脦脧脨脩脪脫脭脮脰脳脴脵脷脹脺脻脼脽脿谩",
		"mang":   "芒茫盲氓忙莽",
		"mao":    "猫茅锚毛矛铆卯茂冒帽貌贸",
		"me":     "么",
		"mei":    "玫枚梅酶霉煤没眉媒镁每美昧寐妹媚",
		"men":    "门闷们",
		"meng":   "萌蒙檬盟锰猛梦孟",
		"mi":     "眯醚靡糜迷谜弥米秘觅泌蜜密幂",
		"mian":   "棉眠绵冕免勉娩缅面",
		"miao":   "苗描瞄藐秒渺庙妙",
		"mie":    "蔑灭",
		"min":    "民抿皿敏悯闽",
		"ming":   "明螟鸣铭名命",
		"miu":    "谬",
		"mo":     "摸腀腁腂腃腄腅腇腉腍腎腏腒腖腗腘腛腜腝腞腟腡腢腣腤腦腨腪腫腬腯腲腳腵腶腷腸膁膃膄膅膆膇膉膋膌膍膎膐膒膓膔膕膖膗膙膚膞膟膠膡膢膤膥膧膩膫膬膭膮膯膰膱膲膴膵膶膷膸膹膼膽膾膿臄臅臇臈臉臋臍臎臏臐臑臒臓摹蘑模膜磨摩魔抹末莫墨默沫漠寞陌",
		"mou":    "谋牟某",
		"mu":     "拇牡亩姆母墓暮幕募慕木目睦牧穆",
		"na":     "拿哪呐钠那娜纳",
		"nai":    "氖乃奶耐奈",
		"nan":    "南男难",
		"nang":   "囊",
		"nao":    "挠脑恼闹淖",
		"ne":     "呢",
		"nei":    "馁内",
		"nen":    "嫩",
		"neng":   "能",
		"ni":     "妮霓倪泥尼拟你匿腻逆溺",
		"nian":   "蔫拈年碾撵捻念",
		"niang":  "娘酿",
		"niao":   "鸟尿",
		"nie":    "捏聂孽啮镊镍涅",
		"nin":    "您",
		"ning":   "柠狞凝宁　臔臕臖臗臘臙臛臜臝臞臟臠臡臢臤臥臦臨臩臫臮臯臰臱臲臵臶臷臸臹臺臽臿舃與興舉舊舋舎舏舑舓舕舖舗舘舙舚舝舠舤舥舦舧舩舮舲舺舼舽舿艀艁艂艃艅艆艈艊艌艍艎艐艑艒艓艔艕艖艗艙艛艜艝艞艠艡艢艣艤艥艦艧艩拧泞",
		"niu":    "牛扭钮纽",
		"nong":   "脓浓农弄",
		"nu":     "奴努怒",
		"nv":     "女",
		"nuan":   "暖",
		"nue":    "虐疟",
		"nuo":    "挪懦糯诺",
		"o":      "哦",
		"ou":     "欧鸥殴藕呕偶沤",
		"pa":     "啪趴爬帕怕琶",
		"pai":    "拍排牌徘湃派",
		"pan":    "攀潘盘磐盼畔判叛",
		"pang":   "乓庞旁耪胖",
		"pao":    "抛咆刨炮袍跑泡",
		"pei":    "呸胚培裴赔陪配佩沛",
		"pen":    "喷盆",
		"peng":   "砰抨烹澎彭蓬棚硼篷膨朋鹏捧碰",
		"pi":     "坯砒霹批披劈琵毗艪艫艬艭艱艵艶艷艸艻艼芀芁芃芅芆芇芉芌芐芓芔芕芖芚芛芞芠芢芣芧芲芵芶芺芻芼芿苀苂苃苅苆苉苐苖苙苚苝苢苧苨苩苪苬苭苮苰苲苳苵苶苸苺苼苽苾苿茀茊茋茍茐茒茓茖茘茙茝茞茟茠茡茢茣茤茥茦茩茪茮茰茲茷茻茽啤脾疲皮匹痞僻屁譬",
		"pian":   "篇偏片骗",
		"piao":   "飘漂瓢票",
		"pie":    "撇瞥",
		"pin":    "拼频贫品聘",
		"ping":   "乒坪苹萍平凭瓶评屏",
		"po":     "坡泼颇婆破魄迫粕剖",
		"pu":     "扑铺仆莆葡菩蒲埔朴圃普浦谱曝瀑",
		"qi":     "期欺栖戚妻七凄漆柒沏其棋奇歧畦崎脐齐旗祈祁骑起岂乞企启契砌器气迄弃汽泣讫",
		"qia":    "掐茾茿荁荂荄荅荈荊荋荌荍荎荓荕荖荗荘荙荝荢荰荱荲荳荴荵荶荹荺荾荿莀莁莂莃莄莇莈莊莋莌莍莏莐莑莔莕莖莗莙莚莝莟莡莢莣莤莥莦莧莬莭莮莯莵莻莾莿菂菃菄菆菈菉菋菍菎菐菑菒菓菕菗菙菚菛菞菢菣菤菦菧菨菫菬菭恰洽",
		"qian":   "牵扦钎铅千迁签仟谦乾黔钱钳前潜遣浅谴堑嵌欠歉",
		"qiang":  "枪呛腔羌墙蔷强抢",
		"qiao":   "橇锹敲悄桥瞧乔侨巧鞘撬翘峭俏窍",
		"qie":    "切茄且怯窃",
		"qin":    "钦侵亲秦琴勤芹擒禽寝沁",
		"qing":   "青轻氢倾卿清擎晴氰情顷请庆",
		"qiong":  "琼穷",
		"qiu":    "秋丘邱球求囚酋泅",
		"qu":     "趋区蛆曲躯屈驱渠菮華菳菴菵菶菷菺菻菼菾菿萀萂萅萇萈萉萊萐萒萓萔萕萖萗萙萚萛萞萟萠萡萢萣萩萪萫萬萭萮萯萰萲萳萴萵萶萷萹萺萻萾萿葀葁葂葃葄葅葇葈葉葊葋葌葍葎葏葐葒葓葔葕葖葘葝葞葟葠葢葤葥葦葧葨葪葮葯葰葲葴葷葹葻葼取娶龋趣去",
		"quan":   "圈颧权醛泉全痊拳犬券劝",
		"que":    "缺炔瘸却鹊榷确雀",
		"qun":    "裙群",
		"ran":    "然燃冉染",
		"rang":   "瓤壤攘嚷让",
		"rao":    "饶扰绕",
		"re":     "惹热",
		"ren":    "壬仁人忍韧任认刃妊纫",
		"reng":   "扔仍",
		"ri":     "日",
		"rong":   "戎茸蓉荣融熔溶容绒冗",
		"rou":    "揉柔肉",
		"ru":     "茹蠕儒孺如辱乳汝入褥",
		"ruan":   "软阮",
		"rui":    "蕊瑞锐",
		"run":    "闰润",
		"ruo":    "若弱",
		"sa":     "撒洒萨",
		"sai":    "腮鳃塞赛",
		"san":    "三叁葽葾葿蒀蒁蒃蒄蒅蒆蒊蒍蒏蒐蒑蒒蒓蒔蒕蒖蒘蒚蒛蒝蒞蒟蒠蒢蒣蒤蒥蒦蒧蒨蒩蒪蒫蒬蒭蒮蒰蒱蒳蒵蒶蒷蒻蒼蒾蓀蓂蓃蓅蓆蓇蓈蓋蓌蓎蓏蓒蓔蓕蓗蓘蓙蓚蓛蓜蓞蓡蓢蓤蓧蓨蓩蓪蓫蓭蓮蓯蓱蓲蓳蓴蓵蓶蓷蓸蓹蓺蓻蓽蓾蔀蔁蔂伞散",
		"sang":   "桑嗓丧",
		"sao":    "搔骚扫嫂",
		"se":     "瑟色涩",
		"sen":    "森",
		"seng":   "僧",
		"sha":    "莎砂杀刹沙纱傻啥煞",
		"shai":   "筛晒",
		"shan":   "珊苫杉山删煽衫闪陕擅赡膳善汕扇缮",
		"shang":  "墒伤商赏晌上尚裳",
		"shao":   "梢捎稍烧芍勺韶少哨邵绍",
		"she":    "奢赊蛇舌舍赦摄射慑涉社设",
		"shen":   "砷申呻伸身深娠绅神沈审婶甚肾慎渗",
		"sheng":  "声生甥牲升绳　蔃蔄蔆蔇蔈蔉蔊蔋蔍蔎蔏蔐蔒蔔蔕蔖蔘蔙蔛蔜蔝蔞蔠蔢蔣蔤蔥蔦蔧蔨蔩蔪蔭蔮蔯蔰蔱蔲蔳蔴蔵蔶蔾蔿蕀蕁蕂蕄蕅蕆蕇蕋蕌蕍蕎蕏蕐蕑蕒蕓蕔蕕蕗蕘蕚蕛蕜蕝蕟蕠蕡蕢蕣蕥蕦蕧蕩蕪蕫蕬蕭蕮蕯蕰蕱蕳蕵蕶蕷蕸蕼蕽蕿薀薁省盛剩胜圣",
		"shi":    "师失狮施湿诗尸虱十石拾时什食蚀实识史矢使屎驶始式示士世柿事拭誓逝势是嗜噬适仕侍释饰氏市恃室视试",
		"shou":   "收手首守寿授售受瘦兽",
		"shu":    "蔬枢梳殊抒输叔舒淑疏书赎孰熟薯暑曙署蜀黍鼠属术述树束戍竖墅庶数漱薂薃薆薈薉薊薋薌薍薎薐薑薒薓薔薕薖薗薘薙薚薝薞薟薠薡薢薣薥薦薧薩薫薬薭薱薲薳薴薵薶薸薺薻薼薽薾薿藀藂藃藄藅藆藇藈藊藋藌藍藎藑藒藔藖藗藘藙藚藛藝藞藟藠藡藢藣藥藦藧藨藪藫藬藭藮藯藰藱藲藳藴藵藶藷藸恕",
		"shua":   "刷耍",
		"shuai":  "摔衰甩帅",
		"shuan":  "栓拴",
		"shuang": "霜双爽",
		"shui":   "谁水睡税",
		"shun":   "吮瞬顺舜",
		"shuo":   "说硕朔烁",
		"si":     "斯撕嘶思私司丝死肆寺嗣四伺似饲巳",
		"song":   "松耸怂颂送宋讼诵",
		"sou":    "搜艘擞",
		"su":     "嗽苏酥俗素速粟僳塑溯宿诉肃",
		"suan":   "酸蒜算",
		"sui":    "虽隋随绥髓碎岁穗遂隧祟",
		"sun":    "孙损笋",
		"suo":    "蓑梭唆缩琐索锁所",
		"ta":     "塌他它她塔藹藺藼藽藾蘀蘁蘂蘃蘄蘆蘇蘈蘉蘊蘋蘌蘍蘎蘏蘐蘒蘓蘔蘕蘗蘘蘙蘚蘛蘜蘝蘞蘟蘠蘡蘢蘣蘤蘥蘦蘨蘪蘫蘬蘭蘮蘯蘰蘱蘲蘳蘴蘵蘶蘷蘹蘺蘻蘽蘾蘿虀虁虂虃虄虅虆虇虈虉虊虋虌虒虓處虖虗虘虙虛虜虝號虠虡虣虤虥虦虧虨虩虪獭挞蹋踏",
		"tai":    "胎苔抬台泰酞太态汰",
		"tan":    "坍摊贪瘫滩坛檀痰潭谭谈坦毯袒碳探叹炭",
		"tang":   "汤塘搪堂棠膛唐糖倘躺淌趟烫",
		"tao":    "掏涛滔绦萄桃逃淘陶讨套",
		"te":     "特",
		"teng":   "藤腾疼誊",
		"ti":     "梯剔踢锑提题蹄啼体替嚏惕涕剃屉",
		"tian":   "天添填田甜恬舔腆",
		"tiao":   "挑条迢眺跳",
		"tie":    "贴铁帖",
		"ting":   "厅听烃虭虯虰虲虳虴虵虷虸蚃蚄蚅蚆蚇蚈蚉蚎蚏蚐蚑蚒蚔蚖蚗蚘蚙蚚蚛蚞蚟蚠蚡蚢蚥蚦蚫蚭蚮蚲蚳蚷蚸蚹蚻蚼蚽蚾蚿蛁蛂蛃蛅蛈蛌蛍蛒蛓蛕蛖蛗蛚蛜蛝蛠蛡蛢蛣蛥蛦蛧蛨蛪蛫蛬蛯蛵蛶蛷蛺蛻蛼蛽蛿蜁蜄蜅蜆蜋蜌蜎蜏蜐蜑蜔蜖汀廷停亭庭挺艇",
		"tong":   "通桐酮瞳同铜彤童桶捅筒统痛",
		"tou":    "偷投头透",
		"tu":     "凸秃突图徒途涂屠土吐兔",
		"tuan":   "湍团",
		"tui":    "推颓腿蜕褪退",
		"tun":    "吞屯臀",
		"tuo":    "拖托脱鸵陀驮驼椭妥拓唾",
		"wa":     "挖哇蛙洼娃瓦袜",
		"wai":    "歪外",
		"wan":    "豌弯湾玩顽丸烷完碗挽晚皖惋宛婉万腕",
		"wang":   "汪王亡枉网往旺望忘妄",
		"wei":    "威蜙蜛蜝蜟蜠蜤蜦蜧蜨蜪蜫蜬蜭蜯蜰蜲蜳蜵蜶蜸蜹蜺蜼蜽蝀蝁蝂蝃蝄蝅蝆蝊蝋蝍蝏蝐蝑蝒蝔蝕蝖蝘蝚蝛蝜蝝蝞蝟蝡蝢蝦蝧蝨蝩蝪蝫蝬蝭蝯蝱蝲蝳蝵蝷蝸蝹蝺蝿螀螁螄螆螇螉螊螌螎螏螐螑螒螔螕螖螘螙螚螛螜螝螞螠螡螢螣螤巍微危韦违桅围唯惟为潍维苇萎委伟伪尾纬未蔚味畏胃喂魏位渭谓尉慰卫",
		"wen":    "瘟温蚊文闻纹吻稳紊问",
		"weng":   "嗡翁瓮",
		"wo":     "挝蜗涡窝我斡卧握沃",
		"wu":     "巫呜钨乌污诬屋无芜梧吾吴毋武五捂午舞伍侮坞戊雾晤物勿务悟误",
		"xi":     "昔熙析西硒矽晰嘻吸锡牺螥螦螧螩螪螮螰螱螲螴螶螷螸螹螻螼螾螿蟁蟂蟃蟄蟅蟇蟈蟉蟌蟍蟎蟏蟐蟔蟕蟖蟗蟘蟙蟚蟜蟝蟞蟟蟡蟢蟣蟤蟦蟧蟨蟩蟫蟬蟭蟯蟰蟱蟲蟳蟴蟵蟶蟷蟸蟺蟻蟼蟽蟿蠀蠁蠂蠄蠅蠆蠇蠈蠉蠋蠌蠍蠎蠏蠐蠑蠒蠔蠗蠘蠙蠚蠜蠝蠞蠟蠠蠣稀息希悉膝夕惜熄烯溪汐犀檄袭席习媳喜铣洗系隙戏细",
		"xia":    "瞎虾匣霞辖暇峡侠狭下厦夏吓",
		"xian":   "掀锨先仙鲜纤咸贤衔舷闲涎弦嫌显险现献县腺馅羡宪陷限线",
		"xiang":  "相厢镶香箱襄湘乡翔祥详想响享项巷橡像向象",
		"xiao":   "萧硝霄削哮嚣销消宵淆晓蠤蠥蠧蠨蠩蠪蠫蠬蠭蠮蠯蠰蠱蠳蠴蠵蠶蠷蠸蠺蠻蠽蠾蠿衁衂衃衆衇衈衉衊衋衎衏衐衑衒術衕衖衘衚衛衜衝衞衟衠衦衧衪衭衯衱衳衴衵衶衸衹衺衻衼袀袃袆袇袉袊袌袎袏袐袑袓袔袕袗袘袙袚袛袝袞袟袠袡袣袥袦袧袨袩袪小孝校肖啸笑效",
		"xie":    "楔些歇蝎鞋协挟携邪斜胁谐写械卸蟹懈泄泻谢屑",
		"xin":    "薪芯锌欣辛新忻心信衅",
		"xing":   "星腥猩惺兴刑型形邢行醒幸杏性姓",
		"xiong":  "兄凶胸匈汹雄熊",
		"xiu":    "休修羞朽嗅锈秀袖绣",
		"xu":     "墟戌需虚嘘须徐许蓄酗叙旭序畜恤絮婿绪续",
		"xuan":   "轩喧宣悬旋玄　袬袮袯袰袲袴袵袶袸袹袺袻袽袾袿裀裃裄裇裈裊裋裌裍裏裐裑裓裖裗裚裛補裝裞裠裡裦裧裩裪裫裬裭裮裯裲裵裶裷裺裻製裿褀褁褃褄褅褆複褈褉褋褌褍褎褏褑褔褕褖褗褘褜褝褞褟褠褢褣褤褦褧褨褩褬褭褮褯褱褲褳褵褷选癣眩绚",
		"xue":    "靴薛学穴雪血",
		"xun":    "勋熏循旬询寻驯巡殉汛训讯逊迅",
		"ya":     "压押鸦鸭呀丫芽牙蚜崖衙涯雅哑亚讶",
		"yan":    "焉咽阉烟淹盐严研蜒岩延言颜阎炎沿奄掩眼衍演艳堰燕厌砚雁唁彦焰宴谚验",
		"yang":   "殃央鸯秧杨扬佯疡羊洋阳氧仰痒养样漾",
		"yao":    "邀腰妖瑶褸褹褺褻褼褽褾褿襀襂襃襅襆襇襈襉襊襋襌襍襎襏襐襑襒襓襔襕襖襗襘襙襚襛襜襝襠襡襢襣襤襥襧襨襩襪襫襬襭襮襯襰襱襲襳襴襵襶襷襸襹襺襼襽襾覀覂覄覅覇覈覉覊見覌覍覎規覐覑覒覓覔覕視覗覘覙覚覛覜覝覞覟覠覡摇尧遥窑谣姚咬舀药要耀",
		"ye":     "椰噎耶爷野冶也页掖业叶曳腋夜液",
		"yi":     "一壹医揖铱依伊衣颐夷遗移仪胰疑沂宜姨彝椅蚁倚已乙矣以艺抑易邑屹亿役臆逸肄疫亦裔意毅忆义益溢诣议谊译异翼翌绎",
		"yin":    "茵荫因殷音阴姻吟银淫寅饮尹引隐覢覣覤覥覦覧覨覩親覫覬覭覮覯覰覱覲観覴覵覶覷覸覹覺覻覼覽覾覿觀觃觍觓觔觕觗觘觙觛觝觟觠觡觢觤觧觨觩觪觬觭觮觰觱觲觴觵觶觷觸觹觺觻觼觽觾觿訁訂訃訄訅訆計訉訊訋訌訍討訏訐訑訒訓訔訕訖託記訙訚訛訜訝印",
		"ying":   "英樱婴鹰应缨莹萤营荧蝇迎赢盈影颖硬映",
		"yo":     "哟",
		"yong":   "拥佣臃痈庸雍踊蛹咏泳涌永恿勇用",
		"you":    "幽优悠忧尤由邮铀犹油游酉有友右佑釉诱又幼迂",
		"yu":     "淤于盂榆虞愚舆余俞逾鱼愉渝渔隅予娱雨与屿禹宇语羽玉域芋郁吁遇喻峪御愈欲狱育誉訞訟訠訡訢訣訤訥訦訧訨訩訪訫訬設訮訯訰許訲訳訴訵訶訷訸訹診註証訽訿詀詁詂詃詄詅詆詇詉詊詋詌詍詎詏詐詑詒詓詔評詖詗詘詙詚詛詜詝詞詟詠詡詢詣詤詥試詧詨詩詪詫詬詭詮詯詰話該詳詴詵詶詷詸詺詻詼詽詾詿誀浴寓裕预豫驭",
		"yuan":   "鸳渊冤元垣袁原援辕园员圆猿源缘远苑愿怨院",
		"yue":    "曰约越跃钥岳粤月悦阅",
		"yun":    "耘云郧匀陨允运蕴酝晕韵孕",
		"za":     "匝砸杂",
		"zai":    "栽哉灾宰载再在",
		"zan":    "咱攒暂赞",
		"zang":   "赃脏葬",
		"zao":    "遭糟凿藻枣早澡蚤躁噪造皂灶燥",
		"ze":     "责择则泽",
		"zei":    "贼",
		"zen":    "怎",
		"zeng":   "增憎曾赠",
		"zha":    "扎喳渣札轧誁誂誃誄誅誆誇誈誋誌認誎誏誐誑誒誔誕誖誗誘誙誚誛誜誝語誟誠誡誢誣誤誥誦誧誨誩說誫説読誮誯誰誱課誳誴誵誶誷誸誹誺誻誼誽誾調諀諁諂諃諄諅諆談諈諉諊請諌諍諎諏諐諑諒諓諔諕論諗諘諙諚諛諜諝諞諟諠諡諢諣铡闸眨栅榨咋乍炸诈",
		"zhai":   "摘斋宅窄债寨",
		"zhan":   "瞻毡詹粘沾盏斩辗崭展蘸栈占战站湛绽",
		"zhang":  "樟章彰漳张掌涨杖丈帐账仗胀瘴障",
		"zhao":   "招昭找沼赵照罩兆肇召",
		"zhe":    "遮折哲蛰辙者锗蔗这浙",
		"zhen":   "珍斟真甄砧臻贞针侦枕疹诊震振镇阵",
		"zheng":  "蒸挣睁征狰争怔整拯正政諤諥諦諧諨諩諪諫諬諭諮諯諰諱諲諳諴諵諶諷諸諹諺諻諼諽諾諿謀謁謂謃謄謅謆謈謉謊謋謌謍謎謏謐謑謒謓謔謕謖謗謘謙謚講謜謝謞謟謠謡謢謣謤謥謧謨謩謪謫謬謭謮謯謰謱謲謳謴謵謶謷謸謹謺謻謼謽謾謿譀譁譂譃譄譅帧症郑证",
		"zhi":    "芝枝支吱蜘知肢脂汁之织职直植殖执值侄址指止趾只旨纸志挚掷至致置帜峙制智秩稚质炙痔滞治窒",
		"zhong":  "中盅忠钟衷终种肿重仲众",
		"zhou":   "舟周州洲诌粥轴肘帚咒皱宙昼骤",
		"zhu":    "珠株蛛朱猪诸诛逐竹烛煮拄瞩嘱主著柱助蛀贮铸筑譆譇譈證譊譋譌譍譎譏譐譑譒譓譔譕譖譗識譙譚譛譜譝譞譟譠譡譢譣譤譥譧譨譩譪譫譭譮譯議譱譲譳譴譵譶護譸譹譺譻譼譽譾譿讀讁讂讃讄讅讆讇讈讉變讋讌讍讎讏讐讑讒讓讔讕讖讗讘讙讚讛讜讝讞讟讬讱讻诇诐诪谉谞住注祝驻",
		"zhua":   "抓爪",
		"zhuai":  "拽",
		"zhuan":  "专砖转撰赚篆",
		"zhuang": "桩庄装妆撞壮状",
		"zhui":   "椎锥追赘坠缀",
		"zhun":   "谆准",
		"zhuo":   "捉拙卓桌琢茁酌啄着灼浊",
		"zi":     "兹咨资姿滋淄孜紫仔籽滓子自渍字",
		"zong":   "鬃棕踪宗综总纵",
		"zou":    "邹走奏揍",
		"zu":     "租足卒族祖诅阻组",
		"zuan":   "钻纂",
		"zui":    "嘴醉最罪",
		"zun":    "尊遵",
		"zuo":    "昨左佐柞做作坐座",
	}
}

// Convert return a zh-cn string to its pinyin
func Convert(str string) (res []string) {
	strArr := strings.Split(str, "")
	for _, char := range strArr {
		for k, v := range chars {
			if strings.Index(v, char) != -1 {
				res = append(res, k)
			}
		}
	}
	return res
}
