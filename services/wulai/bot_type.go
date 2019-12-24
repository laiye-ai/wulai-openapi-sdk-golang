package wulai

//VoiceType 语音消息类型
type VoiceType string

const (
	//AMR :AMR
	AMR VoiceType = "AMR"
	//PCM :PCM
	PCM VoiceType = "PCM"
	//WAV :WAV
	WAV VoiceType = "WAV"
	//OPUS :OPUS
	OPUS VoiceType = "OPUS"
	//SPEEX :SPEEX
	SPEEX VoiceType = "SPEEX"
	//MP3 :MP3
	MP3 VoiceType = "MP3"
)

/*****************************************************************************************/

//direction
type direction string

const (
	//BACKWARD :向旧的消息翻页，查询比传入msg_id更小的消息
	BACKWARD direction = "BACKWARD"

	//FORWARD :先新的消息翻页，查询比传入msg_id更大的消息
	FORWARD direction = "FORWARD"
)

/*****************机器人回复的来源*****************/

//BotSource 机器人回复的来源
type BotSource string

const (
	//DEFAULT_ANSWER_SOURCE :默认
	DEFAULT_ANSWER_SOURCE BotSource = "DEFAULT_ANSWER_SOURCE"

	//KEYWORD_BOT :关键字机器人
	KEYWORD_BOT BotSource = "KEYWORD_BOT"

	//TASK_BOT :任务机器人
	TASK_BOT BotSource = "TASK_BOT"

	//QA_BOT :问答机器人
	QA_BOT BotSource = "QA_BOT"

	//CHITCHAT_BOT :闲聊机器人
	CHITCHAT_BOT BotSource = "CHITCHAT_BOT"
)

/*****************消息类型*****************/

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

//EventType 事件类型
type EventType string

const (
	//EVENT_TYPE_DEFAULT :默认事件
	EVENT_TYPE_DEFAULT EventType = "EVENT_TYPE_DEFAULT"

	//CUSTOM_EVENT :自定义事件
	CUSTOM_EVENT EventType = "CUSTOM_EVENT"

	//ENTER :进入事件
	ENTER EventType = "ENTER"
)

//MsgDirection 消息方向
type MsgDirection string

const (
	//TO_USER 用户收到的消息
	TO_USER MsgDirection = "TO_USER"

	//FROM_USER 用户发出的消息
	FROM_USER MsgDirection = "FROM_USER"
)

//MsgUserType 用户的消息类型
type MsgUserType string

const (
	//PRIVATE PRIVATE
	PRIVATE MsgUserType = "PRIVATE"

	//PUBLIC PUBLIC
	PUBLIC MsgUserType = "PUBLIC"

	//PLATFORM PLATFORM
	PLATFORM MsgUserType = "PLATFORM"

	//PRIVATE_GROUP PRIVATE_GROUP
	PRIVATE_GROUP MsgUserType = "PRIVATE_GROUP"

	//U_BOT_PRIVATE U_BOT_PRIVATE
	U_BOT_PRIVATE MsgUserType = "U_BOT_PRIVATE"

	//U_BOT_GROUP U_BOT_GROUP
	U_BOT_GROUP MsgUserType = "U_BOT_GROUP"

	//WORK_WECHAT WORK_WECHAT
	WORK_WECHAT MsgUserType = "WORK_WECHAT"

	//WORK_WECHAT_GROUP WORK_WECHAT_GROUP
	WORK_WECHAT_GROUP MsgUserType = "WORK_WECHAT_GROUP"
)

//BlockType 对话单元类型
type BlockType string

const (
	//BLOCK_TYPE_MESSAGE 消息单元
	BLOCK_TYPE_MESSAGE BlockType = "BLOCK_TYPE_MESSAGE"

	//BLOCK_TYPE_ASK 询问单元
	BLOCK_TYPE_ASK BlockType = "BLOCK_TYPE_ASK"

	//BLOCK_TYPE_HIDE 隐藏单元
	BLOCK_TYPE_HIDE BlockType = "BLOCK_TYPE_HIDE"

	//BLOCK_TYPE_LINK 跳转单元
	BLOCK_TYPE_LINK BlockType = "BLOCK_TYPE_LINK"

	//BLOCK_TYPE_ADVANCE_INTERFACE 高级接口
	BLOCK_TYPE_ADVANCE_INTERFACE BlockType = "BLOCK_TYPE_ADVANCE_INTERFACE"

	//BLOCK_TYPE_INTERFACE 接口单元
	BLOCK_TYPE_INTERFACE BlockType = "BLOCK_TYPE_INTERFACE"

	//BLOCK_TYPE_CALCULATE 运算单元
	BLOCK_TYPE_CALCULATE BlockType = "BLOCK_TYPE_CALCULATE"

	//BLOCK_TYPE_COLLECT 收集单元
	BLOCK_TYPE_COLLECT BlockType = "BLOCK_TYPE_COLLECT"
)

/****************
- 消息类型
****************/

//Text 消息类型:文本消息
type Text struct {
	Content string `json:"content"` //文本消息体 [ 1 .. 2048 ] characters
}

//Image 消息类型:图片消息
type Image struct {
	ResourceURL string `json:"resource_url"` //图片的资源链接,必填 [ 1 .. 512 ] characters
}

//Custom 消息类型:自定义消息
type Custom struct {
	Content string `json:"content"` //消息内容
}

//Video 消息类型:视频消息
type Video struct {
	ResourceURL string `json:"resource_url"` //视频的资源链接,必填 [ 1 .. 512 ] characters
	Thumb       string `json:"thumb"`        //视频的缩略图 <= 512 characters
	Description string `json:"description"`  //视频描述 <= 512 characters
	Title       string `json:"title"`        //视频标题 <= 128 characters
}

//File 消息类型:文件消息
type File struct {
	FileName    string `json:"file_name"`    //文件名 <= 128 characters
	ResourceURL string `json:"resource_url"` //视频的资源链接,必填 [ 1 .. 512 ] characters
}

//RichText 消息类型:图文消息
type RichText struct {
	ResourceURL string `json:"resource_url"` //视频的资源链接,必填 [ 1 .. 512 ] characters
}

//Voice 消息类型:语音消息
type Voice struct {
	ResourceURL string    `json:"resource_url"` //语音的资源链接 [ 1 .. 512 ] characters
	Type        VoiceType `json:"type"`         //语音消息类型 "AMR" "PCM" "WAV" "OPUS" "SPEEX" "MP3"
	Recognition string    `json:"recognition"`  //语音识别文本结果，发消息时非必填 <= 1024 characters
}

//Event 消息类型:事件消息
type Event struct {
	Fields    interface{} `json:"fields"`     //protobufStruct
	EventType EventType   `json:"event_type"` //事件消息
}

//ShareLink 消息类型:图文消息
type ShareLink struct {
	Description    string `json:"description"`     //视频描述 <= 512 characters
	DestinationURL string `json:"destination_url"` //链接目标地址,必填 [ 1 .. 512 ] characters
	CoverURL       string `json:"cover_url"`       //封面图片地址,必填 [ 1 .. 512 ] characters
	Title          string `json:"title"`           //链接的文字标题,必填 [ 1 .. 128 ] characters
}

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
	Source     BotSource  `json:"source"`      //回复的来源
	Score      float64    `json:"score"`       //置信度(score<=1)
	Response   []Response `json:"response"`    //回复内容
	QuickReply []string   `json:"quick_reply"` //快捷回复
}

//Response 回复内容
type Response struct {
	MsgBody         MsgBody           `json:"msg_body"`         //任意选择一种消息类型（文本/图片/语音/视频/文件/图文/自定义消息）填充
	SimilarResponse []SimilarResponse `json:"similar_response"` //机器人平台手动或自动配置的推荐知识点，在机器人平台-知识库-知识点详情配置
	EnableEvaluate  bool              `json:"enable_evaluate"`  //是否允许点赞/点踩
	DelayTs         int64             `json:"delay_ts"`         //延时发送的秒数 delay_ts<= 604800
	AnswerID        int64             `json:"answer_id"`        //答案Id
}

//MsgBody 消息体格式，任意选择一种消息类型
type MsgBody struct {
	Text      *Text      `json:"text"`       //文本消息
	Image     *Image     `json:"image"`      //图片消息
	Custom    *Custom    `json:"custom"`     //自定义消息
	RichText  *RichText  `json:"rich_text"`  //图文消息
	Video     *Video     `json:"video"`      //视频消息
	File      *File      `json:"file"`       //文件消息
	Voice     *Voice     `json:"voice"`      //语音消息
	Event     *Event     `json:"event"`      //事件消息
	ShareLink *ShareLink `json:"share_link"` //图文消息
}

//SimilarResponse 推荐知识点
type SimilarResponse struct {
	URL    string    `json:"url"`    //相似问对应的url
	Source BotSource `json:"source"` //回复的来源
	Detail Detail    `json:"detail"` //回复的机器人:问答机器人/闲聊机器人/任务机器人/关键字机器人
}

//SimilarResponseParam 推荐知识点(给用户发消息参数)
type SimilarResponseParam struct {
	Source BotSource   `json:"source"` //回复的来源(必填项)
	Detail interface{} `json:"detail"` //指针类型 qa / chitchat / task / keyword  (如果机器人回复兜底内容，则bot为空)(必填项)
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
	QA       *QA       `json:"qa"`       //问答机器人
	Chitchat *Chitchat `json:"chitchat"` //闲聊机器人
	Task     *Task     `json:"task"`     //任务机器人
	Keyword  *Keyword  `json:"keyword"`  //关键字机器人
}

//Detail 机器人类型
type Detail struct {
	QA       *QA       `json:"qa"`       //问答机器人
	Chitchat *Chitchat `json:"chitchat"` //闲聊机器人
	Task     *Task     `json:"task"`     //任务机器人
	Keyword  *Keyword  `json:"keyword"`  //关键字机器人
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
	KeywordID int64  `json:"keyword_id"` //关键字id
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
	BlockType BlockType `json:"block_type"` //对话单元类型.
	BlockID   int64     `json:"block_id"`   //任务型机器人对话单元id
	TaskID    int64     `json:"task_id"`    //任务id
	BlockName string    `json:"block_name"` //任务机器人对话单元名
	Entities  []Entity  `json:"entities"`   //抽取的实体列表
	TaskName  string    `json:"task_name"`  //任务机器人任务名
	RobotID   int64     `json:"robot_id"`   //机器人id
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
	Direction  MsgDirection `json:"direction"`   //消息方向:TO_USER: 用户收到的消息;FROM_USER: 用户发出的消息.
	SenderInfo SenderInfo   `json:"sender_info"` //消息发出者信息
	MsgType    MsgTypeEnum  `json:"msg_type"`    //消息类型:默认TEXT
	Extra      string       `json:"extra"`       //自定义字段(<= 1024 characters)
	MsgID      string       `json:"msg_id"`      //消息的唯一标识(<= 18 characters)
	MsgTs      string       `json:"msg_ts"`      //消息毫秒级时间戳(>= 1)
	UserInfo   UserInfo     `json:"user_info"`   //用户信息
	MsgBody    MsgBody      `json:"msg_body"`    //消息体
}

//MsgReceive 接收用户发的消息
type MsgReceive struct {
	MsgID string `json:"msg_id"`
}

//MsgSend 给用户发消息
type MsgSend struct {
	MsgID string `json:"msg_id"`
}

//MsgSync 同步发给用户的消息
type MsgSync struct {
	MsgID string `json:"msg_id"`
}

//MsgTriggerLink 触发链接消息
type MsgTriggerLink struct {
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
	IsDispatch        bool                `json:"is_dispatch"`        //是否转人工
	SuggestedResponse []SuggestedResponse `json:"suggested_response"` //本次机器人应答内容列表
	Extra             string              `json:"extra"`              //自定义字段
}

//MsgSuggestionResponse 获取用户输入联想响应
type MsgSuggestionResponse struct {
	UserSuggestions []UserSuggestion `json:"user_suggestions"` //在相同对话类型中，如有多条联想内容，按照置信度从高到低排序；在不同对话类型中，如有任务对话的联想，则不返回问答对话的联想内容
}

//UserSuggestion 联想内容
type UserSuggestion struct {
	Suggestion string `json:"suggestion"`
}

//UserAttributeFilter 用户属性过滤条件。如果填写，代表需要过滤；反之不过滤。
type UserAttributeFilter struct {
	HasAttribute bool `json:"use_in_user_attribute_group"` //是否可以作为属性组属性
}
