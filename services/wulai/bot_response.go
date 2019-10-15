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
	IsSend     bool          `json:"is_send"`     //是否发给用户。只有当召回知识点的分数最高且高于吾来SaaS上设置的阈值时，取值为 true
	Bot        Bot           `json:"bot"`         //机器人类型:问答机器人/闲聊机器人/任务机器人/关键字机器人
	Source     BotSourceEnum `json:"source"`      //回复的来源
	Score      float64       `json:"score"`       //置信度(score<=1)
	Response   []Response    `json:"response"`    //回复内容
	QuickReply []string      `json:"quick_reply"` //快捷回复
}

//Response 回复内容
type Response struct {
	MsgBody         MsgBody           `json:"msg_body"`         //任意选择一种消息类型（文本/图片/语音/视频/文件/图文/自定义消息）填充
	SimilarResponse []SimilarResponse `json:"similar_response"` //机器人平台手动或自动配置的推荐知识点，在机器人平台-知识库-知识点详情配置
	EnableEvaluate  bool              `json:"enable_evaluate"`  //是否允许点赞/点踩
	DelayTs         int64             `json:"delay_ts"`         //延时发送的秒数 delay_ts<= 604800
}

//MsgBody 消息体格式，任意选择一种消息类型
type MsgBody struct {
	Text      Text      `json:"text"`       //文本消息
	Image     Image     `json:"image"`      //图片消息
	Custom    Custom    `json:"custom"`     //自定义消息
	Video     Video     `json:"video"`      //视频消息
	File      File      `json:"file"`       //文件消息
	Voice     Voice     `json:"voice"`      //语音消息
	Event     Event     `json:"event"`      //事件消息
	ShareLink ShareLink `json:"share_link"` //图文消息
}

//SimilarResponse 推荐知识点
type SimilarResponse struct {
	URL    string        `json:"url"`    //相似问对应的url
	Source BotSourceEnum `json:"source"` //回复的来源
	Detail Detail        `json:"detail"` //回复的机器人:问答机器人/闲聊机器人/任务机器人/关键字机器人
}

//QA 问答机器人(标准)
type QA struct {
	KnowledgeID      int64  `json:"knowledge_id"`      //知识点id
	StandardQuestion string `json:"standard_question"` //标准问(<= 100 characters)
	Question         string `json:"question"`          //命中的相似问(<= 1024 characters)
	IsNoneIntention  bool   `json:"is_none_intention"` //是否为无意图知识点(boolean)
}

//Chitchat 闲聊机器人
type Chitchat struct {
	Corpus string `json:"corpus"` //闲聊机器人类型
}

//Bot 机器人类型
type Bot struct {
	QA       QA       `json:"qa"`       //问答机器人
	Chitchat Chitchat `json:"chitchat"` //闲聊机器人
	Task     Task     `json:"task"`     //任务机器人
	Keyword  Keyword  `json:"keyword"`  //关键字机器人
}

//Detail 机器人类型
type Detail struct {
	QA       QA       `json:"qa"`       //问答机器人
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
	QA         QA         `json:"qa"`          //问答机器人
	IsSend     bool       `json:"is_send"`     //是否发给用户
	Score      float64    `json:"score"`       //置信度 score<= 1
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
	Score      float64    `json:"score"`       //置信度 score<= 1
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

//Task 任务机器人
type Task struct {
	BlockType string   `json:"block_type"` //对话单元类型.
	BlockID   int64    `json:"block_id"`   //任务型机器人对话单元id
	TaskID    int64    `json:"task_id"`    //任务id
	BlockName string   `json:"block_name"` //任务机器人对话单元名
	Entities  []Entity `json:"entities"`   //抽取的实体列表
	TaskName  string   `json:"task_name"`  //任务机器人任务名
	RobotID   int64    `json:"robot_id"`   //机器人id
}

//Entity 抽取的实体列表
type Entity struct {
	IdxEnd   int64  `json:"idx_end"`   //实体值原始片段在query中的结束位置
	Name     string `json:"name"`      //实体名称
	IdxStart int64  `json:"idx_start"` //实体值原始片段在query中的起始位置
	Value    string `json:"value"`     //实体值
	SegValue string `json:"seg_value"` //实体值的原始片段
	Type     string `json:"type"`      //实体类型枚举:"ENTITY_TYPE_SYS" "ENTITY_TYPE_UDEFINE" "ENTITY_TYPE_CONTENT" "ENTITY_TYPE_REGEX" "ENTITY_TYPE_CITE" "ENTITY_TYPE_WEBHOOK" "ENTITY_TYPE_USERACT"
	Desc     string `json:"desc"`      //实体别名
}

//TaskSuggestedResponse 本次机器人应答内容列表。机器人的响应可能会有多个内容。
type TaskSuggestedResponse struct {
	Score      float64    `json:"score"`
	IsSend     bool       `json:"is_send"`
	Task       Task       `json:"task"`
	Response   []Response `json:"response"`
	QuickReply []string   `json:"quick_reply"`
}

//MsgHistory 历史消息
type MsgHistory struct {
	Msg     []Msg `json:"msg"`      //返回的消息列表
	HasMore bool  `json:"has_more"` //是否有更多消息:false：已经翻到最后一页;true：还有消息没有展示出来，可以继续翻页.
}

//SenderInfo 消息发出者信息，可以是人工或机器人
type SenderInfo struct {
	AvatarURL string `json:"avatar_url"` //头像链接。如果是机器人，显示机器人默认头像 (<= 128 characters)
	Nickname  string `json:"nickname"`   //昵称。如果是机器人，显示AI (<= 32 characters)
	RealName  string `json:"real_name"`  //真实姓名。如果是机器人，显示AI ([ 1 .. 32 ] characters)
}

//UserInfo 用户信息
type UserInfo struct {
	AvatarURL string `json:"avatar_url"` //用户头像地址(<= 512 characters)
	Nickname  string `json:"nickname"`   //用户昵称(<= 128 characters)
}

//Msg 消息列表
type Msg struct {
	Direction  string      `json:"direction"`   //消息方向:TO_USER: 用户收到的消息;FROM_USER: 用户发出的消息.//TODO:枚举类型
	SenderInfo SenderInfo  `json:"sender_info"` //消息发出者信息
	MsgType    MsgTypeEnum `json:"msg_type"`    //消息类型:默认TEXT //TODO:枚举类型
	Extra      string      `json:"extra"`       //自定义字段(<= 1024 characters)
	MsgID      string      `json:"msg_id"`      //消息的唯一标识(<= 18 characters)
	MsgTs      string      `json:"msg_ts"`      //消息毫秒级时间戳(>= 1)
	UserInfo   UserInfo    `json:"user_info"`   //用户信息
	MsgBody    MsgBody     `json:"msg_body"`    //消息体
}

//MsgReceive 接收用户发的消息
type MsgReceive struct {
	MsgID string `json:"msg_id"`
}

//MsgSync 同步发给用户的消息
type MsgSync struct {
	MsgID string `json:"msg_id"`
}

/****************
- 消息投递
****************/

//MessageDelivery "消息投递"返回的数据结构体
type MessageDelivery struct {
	UserID          string            `json:"user_id"`
	SenderInfo      SenderInfo        `json:"sender_info"`
	MsgType         MsgTypeEnum       `json:"msg_type"`
	Extra           string            `json:"extra"`
	MsgID           string            `json:"msg_id"`
	Bot             Bot               `json:"bot"`
	MsgTs           string            `json:"msg_ts"`
	Source          string            `json:"source"`
	MsgBody         MsgBody           `json:"msg_body"`
	SimilarResponse []SimilarResponse `json:"similar_response"`
	EnableEvaluate  bool              `json:"enable_evaluate"`
	QuickReply      []string          `json:"quick_reply"`
}

/****************
- 消息路由
****************/

//MessageRoute "消息路由"返回的数据结构体
type MessageRoute struct {
	SuggestedResponse []SuggestedResponse `json:"suggested_response"`
	UserID            string              `json:"user_id"`
	Extra             string              `json:"extra"`
	MsgID             string              `json:"msg_id"`
	IsDispatch        bool                `json:"is_dispatch"`
	MsgTs             string              `json:"msg_ts"`
	MsgBody           MsgBody             `json:"msg_body"`
	Nickname          string              `json:"nickname"`
}

//MessageRouteResponses 处理"消息路由"后的相应数据体
type MessageRouteResponses struct {
	IsDispatch        bool                `json:"is_dispatch"`
	SuggestedResponse []SuggestedResponse `json:"suggested_response"`
	Extra             string              `json:"extra"`
}
