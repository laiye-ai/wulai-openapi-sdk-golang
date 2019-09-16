package wulai

/****************
- 通用机器人
****************/

//BotResponse 机器人回复
type BotResponse struct {
	IsDispatch        bool                `json:"is_dispatch"`        //是否转人工
	SuggestedResponse []SuggestedResponse `json:"suggested_response"` //机器人应答内容列表。机器人的响应可能会有多个内容
	MsgID             string              `json:"msg_id"`             //消息的唯一标识
}

//SuggestedResponse 机器人应答内容列表。机器人的响应可能会有多个内容
type SuggestedResponse struct {
	IsSend     bool       `json:"is_send"`     //是否发给用户。只有当召回知识点的分数最高且高于吾来SaaS上设置的阈值时，取值为 true
	Bot        Bot        `json:"bot"`         //机器人类型:问答机器人/闲聊机器人/任务机器人/关键字机器人
	Source     source     `json:"source"`      //回复的来源
	Score      int        `json:"score"`       //置信度(score<=1)
	Response   []Response `json:"response"`    //回复内容
	QuickReply []string   `json:"quick_reply"` //快捷回复
}

//Response 回复内容
type Response struct {
	MsgBody         MsgBody           `json:"msg_body"`         //任意选择一种消息类型（文本/图片/语音/视频/文件/图文/自定义消息）填充
	SimilarResponse []SimilarResponse `json:"similar_response"` //机器人平台手动或自动配置的推荐知识点，在机器人平台-知识库-知识点详情配置
	EnableEvaluate  bool              `json:"enable_evaluate"`  //是否允许点赞/点踩
	DelayTs         int               `json:"delay_ts"`         //延时发送的秒数 delay_ts<= 604800
}

//MsgBody 消息体格式，任意选择一种消息类型
type MsgBody struct {
	Text      Text      `json:"text"`
	Image     Image     `json:"image"`
	Custom    Custom    `json:"custom"`
	Video     Video     `json:"video"`
	File      File      `json:"file"`
	Voice     Voice     `json:"voice"`
	Event     Event     `json:"event"`
	ShareLink ShareLink `json:"share_link"`
}

//SimilarResponse 推荐知识点
type SimilarResponse struct {
	URL    string `json:"url"`    //相似问对应的url
	Source source `json:"source"` //回复的来源
	Detail Detail `json:"detail"` //回复的机器人:问答机器人/闲聊机器人/任务机器人/关键字机器人
}

//Qa 问答机器人回复
type Qa struct {
	KnowledgeID      int    `json:"knowledge_id"`
	StandardQuestion string `json:"standard_question"`
	Question         string `json:"question"`
}

//Chitchat 闲聊机器人
type Chitchat struct {
	Corpus string `json:"corpus"` //闲聊机器人类型
}

//Bot 机器人类型
type Bot struct {
	Qa       Qa       `json:"qa"`       //问答机器人
	Chitchat Chitchat `json:"chitchat"` //闲聊机器人
	Task     Task     `json:"task"`     //任务机器人
	Keyword  Keyword  `json:"keyword"`  //关键字机器人
}

//Detail 机器人类型
type Detail struct {
	Qa       Qa       `json:"qa"`       //问答机器人
	Chitchat Chitchat `json:"chitchat"` //闲聊机器人
	Task     Task     `json:"task"`     //任务机器人
	Keyword  Keyword  `json:"keyword"`  //关键字机器人
}

/****************
- 问答机器人
****************/

//BotResponseQa 问答机器人回复
type BotResponseQa struct {
	IsDispatch          bool                  `json:"is_dispatch"`           //是否转人工
	MsgID               string                `json:"msg_id"`                //消息的唯一标识
	QaSuggestedResponse []QaSuggestedResponse `json:"qa_suggested_response"` //机器人的响应
}

//QaSuggestedResponse 本次机器人应答内容列表。机器人的响应可能会有多个内容。
type QaSuggestedResponse struct {
	Qa         Qa         `json:"qa"`          //问答机器人
	IsSend     bool       `json:"is_send"`     //是否发给用户
	Score      int        `json:"score"`       //置信度 score<= 1
	Response   []Response `json:"response"`    //回复内容
	QuickReply []string   `json:"quick_reply"` //快捷回复
}

/****************
- 关键字机器人
****************/

//BotResponseKeyword 关键字机器人回复
type BotResponseKeyword struct {
	IsDispatch               bool                       `json:"is_dispatch"`                //是否转人工
	MsgID                    string                     `json:"msg_id"`                     //消息的唯一标识
	KeywordSuggestedResponse []KeywordSuggestedResponse `json:"keyword_suggested_response"` //机器人的响应
}

//Keyword 关键字机器人
type Keyword struct {
	KeywordID int    `json:"keyword_id"` //关键字id
	Keyword   string `json:"keyword"`    //命中的关键字
}

//KeywordSuggestedResponse 本次机器人应答内容列表。机器人的响应可能会有多个内容。
type KeywordSuggestedResponse struct {
	IsSend     bool       `json:"is_send"`     //是否发给用户
	Score      int        `json:"score"`       //置信度 score<= 1
	Response   []Response `json:"response"`    //回复内容
	Keyword    Keyword    `json:"keyword"`     //关键字机器人
	QuickReply []string   `json:"quick_reply"` //快捷回复
}

/****************
- 任务机器人
****************/

//BotResponseTask 任务机器人回复
type BotResponseTask struct {
	IsDispatch            bool                    `json:"is_dispatch"`             //是否转人工
	MsgID                 string                  `json:"msg_id"`                  //消息的唯一标识
	TaskSuggestedResponse []TaskSuggestedResponse `json:"task_suggested_response"` //机器人的响应
}

//Entities 抽取的实体列表
type Entities struct {
	IdxEnd   int    `json:"idx_end"`   //实体值原始片段在query中的结束位置
	Name     string `json:"name"`      //实体名称
	IdxStart int    `json:"idx_start"` //实体值原始片段在query中的起始位置
	Value    string `json:"value"`     //实体值
	SegValue string `json:"seg_value"` //实体值的原始片段
	Type     string `json:"type"`      //实体类型枚举:"ENTITY_TYPE_SYS" "ENTITY_TYPE_UDEFINE" "ENTITY_TYPE_CONTENT" "ENTITY_TYPE_REGEX" "ENTITY_TYPE_CITE" "ENTITY_TYPE_WEBHOOK" "ENTITY_TYPE_USERACT"
	Desc     string `json:"desc"`      //实体别名
}

//Task 任务机器人
type Task struct {
	BlockType string     `json:"block_type"`
	BlockID   int        `json:"block_id"`
	TaskID    int        `json:"task_id"`
	BlockName string     `json:"block_name"`
	Entities  []Entities `json:"entities"`
	TaskName  string     `json:"task_name"`
	RobotID   int        `json:"robot_id"`
}

//TaskSuggestedResponse 本次机器人应答内容列表。机器人的响应可能会有多个内容。
type TaskSuggestedResponse struct {
	Score      int        `json:"score"`
	IsSend     bool       `json:"is_send"`
	Task       Task       `json:"task"`
	Response   []Response `json:"response"`
	QuickReply []string   `json:"quick_reply"`
}
