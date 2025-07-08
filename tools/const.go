package tools

import "time"

/**
@description
@date: 04/04 17:26
@author Gk
**/

const (

	// online bot
	AppID                   = "cli_a137d87284f8900e"
	AppSecret               = "7VUCFdxin0hprvE5S5PJBgBV7DzXrD4w"
	//MybotEncryptKey         = "SFC2n1bvDn25ckUBGkv8PdlWH1cMt3L1"
	//BotEventUrlVerification = "url_verification"
	//MybotVerificationToken  = "g1hyinudvaowjqU4QapUoeSSImPTSsfp"

	//	export MybotAppId=cli_a2e085041df9d00c
	//	export MybotAppSecret=5VDs005RVGlLRZ0H7KMrihiwwfMn3DUg
	//export VerificationTokenMybot=xnq8YpZx4sSzX0FvuQDweeuvLKrn5rxo
	//export EncryptKeyMybot=toLtRf2jOvKjyu8kG7QxGfH7RVjDPMY7

	//	BotEventUrlVerification = "url_verification"
	//	MybotAppId              = "MybotAppId"
	//	MybotAppSecret          = "MybotAppSecret"
	//	MybotEncryptKey         = "EncryptKeyMybot"
	//	MybotVerificationToken  = "VerificationTokenMybot"

	// 交付助手
	// AppID                   = "cli_a2e085041df9d00c"
	// AppSecret               = "5VDs005RVGlLRZ0H7KMrihiwwfMn3DUg"
	MybotEncryptKey         = "toLtRf2jOvKjyu8kG7QxGfH7RVjDPMY7"
	BotEventUrlVerification = "url_verification"
	MybotVerificationToken  = "xnq8YpZx4sSzX0FvuQDweeuvLKrn5rxo"

	// test bot
	DevAppID     = "cli_a42d7034e0f9900b"
	DevAppSecret = "2Ribx0aKCjlOqF1AgQ6t8fCPGNokeJSs"
	//DevEncrypKey         = "SFC2n1bvDn25ckUBGkv8PdlWH1cMt3L1"
	//DevVerificationToken = "g1hyinudvaowjqU4QapUoeSSImPTSsfp"

	// meego 插件
	MeegoAppID  = "MII_64B0ED6547F88002"
	MeegoSecret = "7FD82D93B83A13C6906E3A8EED6165B8"

	MeegoLarkAppID  = "MII_62B034CA3E048107"
	MeegoLarkSecret = "026A91C2CC4196E75499021566A72A61"
	// gk
	MeegoUserKey = "7019976861685989377"
	OpenIdGk = "ou_f2141fdf4cdfc2eed335e12237bd9aea"
	UnionIdGk = "on_f8cae88b59d4eb11b86bb5b815e17a2a"
	UserIdGK = "58921e17"


	//meego  是lark空间的
	LarkProjectKey = "5e96d7bff4e7c525510f9156"
	LarkSimpleName = "larksuite"

	// 是飞书交付空间
	DeliveryProjectKey = "616e8bbb3341bffeff002bfa"
	DeliverySimpleName = "deliveryofquality"

	// 交付中心
	DeliveryCenterKey        = "639a9d890ca8c5215316c4f9"
	DeliveryCenterSimpleName = "ka-delivery"

	GroupAuto = "groupAuto"

	EventExpireTime = time.Hour * 6
	MeegoExpireTime = time.Minute * 10

	DeliveryPlatformUrl = "https://delivery.feishu.cn"

	// 消息类型说明: https://open.feishu.cn/document/server-docs/im-v1/message-content-description/create_json#3c92befd
	Interactive = "interactive"
	Text        = "text"

	// cmd
	BotCmdHelp       = "help"
	BotCmdEnv        = "env"
	BotCmdNetwork    = "network"
	BotCmdPsm        = "psm"
	BotCmdVm         = "vm"
	BotCmdSet        = "set"
	BotCmdUnit       = "unit"
	BotCmdBot        = "bot"
	BotCmdMini       = "mini"
	BotCmdCheck      = "check"
	BotCmdLog        = "log"
	BotCmdPath       = "path"
	BotCmdMysql      = "mysql"
	BotCmdRocketmq   = "rocketmq"
	BotCmdAbase      = "abase"
	BotCmdBytestore  = "bytestore"
	BotCmdRedis      = "redis"
	BotCmdTos        = "tos"
	BotCmdTenantlist = "tenantlist"
	BotCmdFg         = "fg"
	BotCmdPi         = "pi"
	BotCmdPSM        = "psm"
	BotCmdMeeGo      = "meego"
	BotCmdDbus       = "dts"
	BotCmdTcc        = "tcc"
	BotCmdeeconf     = "eeconf"
	BotCmdTlb        = "tlb"
	BotCmdDocker     = "docker"
	BotCmdSd         = "sd"
	BotCmdDatabus    = "databus"
	BotCmdNtp        = "ntp"
	BotCmdVolcEngine = "volcengine"
	BotCmdStatics    = "statics"
)
