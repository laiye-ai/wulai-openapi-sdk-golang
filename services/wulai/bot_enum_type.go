package wulai

//direction
type direction string

const (
	//BACKWARD :向旧的消息翻页，查询比传入msg_id更小的消息
	BACKWARD direction = "BACKWARD"

	//FORWARD :先新的消息翻页，查询比传入msg_id更大的消息
	FORWARD direction = "FORWARD"
)

//source 机器人回复的来源
type source string

const (
	//DEFAULT_ANSWER_SOURCE :默认
	DEFAULT_ANSWER_SOURCE source = "DEFAULT_ANSWER_SOURCE"

	//KEYWORD_BOT :关键字机器人
	KEYWORD_BOT source = "KEYWORD_BOT"

	//TASK_BOT :任务机器人
	TASK_BOT source = "TASK_BOT"

	//QA_BOT :问答机器人
	QA_BOT source = "QA_BOT"

	//CHITCHAT_BOT :闲聊机器人
	CHITCHAT_BOT source = "CHITCHAT_BOT"
)
