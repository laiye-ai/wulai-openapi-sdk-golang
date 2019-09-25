package wulai

/*****************************************************************************************/

//direction
type direction string

const (
	//BACKWARD :向旧的消息翻页，查询比传入msg_id更小的消息
	BACKWARD direction = "BACKWARD"

	//FORWARD :先新的消息翻页，查询比传入msg_id更大的消息
	FORWARD direction = "FORWARD"
)

/*****************************************************************************************/

//BotSourceEnum 机器人回复的来源
type BotSourceEnum string

const (
	//DEFAULT_ANSWER_SOURCE :默认
	DEFAULT_ANSWER_SOURCE BotSourceEnum = "DEFAULT_ANSWER_SOURCE"

	//KEYWORD_BOT :关键字机器人
	KEYWORD_BOT BotSourceEnum = "KEYWORD_BOT"

	//TASK_BOT :任务机器人
	TASK_BOT BotSourceEnum = "TASK_BOT"

	//QA_BOT :问答机器人
	QA_BOT BotSourceEnum = "QA_BOT"

	//CHITCHAT_BOT :闲聊机器人
	CHITCHAT_BOT BotSourceEnum = "CHITCHAT_BOT"
)

/*****************************************************************************************/

//MsgTypeEnum 消息类型
type MsgTypeEnum string

const (
	//TEXT 文本消息
	TEXT MsgTypeEnum = "TEXT"
	//IMAGE 图片消息
	IMAGE MsgTypeEnum = "IMAGE"
	//VOICE 音频消息
	VOICE MsgTypeEnum = "VOICE"
	//NOTICE 通知消息
	NOTICE MsgTypeEnum = "NOTICE"
	//FILE 文件消息
	FILE MsgTypeEnum = "FILE"
	//SHARELINK 分享链接消息
	SHARELINK MsgTypeEnum = "SHARELINK"
	//VIDEO 视频消息
	VIDEO MsgTypeEnum = "VIDEO"
	//CUSTOM 用户自定义消息
	CUSTOM MsgTypeEnum = "CUSTOM"
	//PUBLIC_EVENT 微信公众号的事件消息
	PUBLIC_EVENT MsgTypeEnum = "PUBLIC_EVENT"
	//NONSUPPORT 不支持的消息类型
	NONSUPPORT MsgTypeEnum = "NONSUPPORT"
	//EVENT 事件消息
	EVENT MsgTypeEnum = "TEXT"
	//CALLBACK_NOTICE 若客户是异步接入的是第三方渠道，调用第三方渠道api的报错信息
	CALLBACK_NOTICE MsgTypeEnum = "CALLBACK_NOTICE"
)
