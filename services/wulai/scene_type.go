package wulai

//IntentSwitchMode 意图切换模式
type IntentSwitchMode string

const (
	//INTENT_SWITCH_MODE_SWITCH :优先切换到其他意图
	INTENT_SWITCH_MODE_SWITCH IntentSwitchMode = "INTENT_SWITCH_MODE_SWITCH"

	//INTENT_SWITCH_MODE_STAY :优先停留在当前意图并填槽
	INTENT_SWITCH_MODE_STAY IntentSwitchMode = "INTENT_SWITCH_MODE_STAY"
)

type TriggerType string

const (
	//TRIGGER_TYPE_ERROR :错误
	TRIGGER_TYPE_ERROR TriggerType = "TRIGGER_TYPE_ERROR"

	//TRIGGER_TYPE_EXACT_MATCH_KEYWORD :关键词完全匹配
	TRIGGER_TYPE_EXACT_MATCH_KEYWORD TriggerType = "TRIGGER_TYPE_EXACT_MATCH_KEYWORD"

	//TRIGGER_TYPE_INCLUDE_KEYWORD :关键词包含匹配
	TRIGGER_TYPE_INCLUDE_KEYWORD TriggerType = "TRIGGER_TYPE_INCLUDE_KEYWORD"

	//TRIGGER_TYPE_SENTENCE :相似说法
	TRIGGER_TYPE_SENTENCE TriggerType = "TRIGGER_TYPE_SENTENCE"
)

//ResponseType 单元回复类型
type ResponseType string

const (

	//RESPONSE_ERROR :错误
	RESPONSE_ERROR ResponseType = "RESPONSE_ERROR"

	//RESPONSE_RANDOM :随机回复
	RESPONSE_RANDOM ResponseType = "RESPONSE_RANDOM"

	//RESPONSE_ALL :全部回复
	RESPONSE_ALL ResponseType = "RESPONSE_ALL"

	//RESPONSE_LOOP :依次回复
	RESPONSE_LOOP ResponseType = "RESPONSE_LOOP"
)

//BlockType 对话单元类型
type BlockType string

const (

	//SCENE_BLOCK_TYPE_DEFAULT :错误
	SCENE_BLOCK_TYPE_DEFAULT BlockType = "SCENE_BLOCK_TYPE_DEFAULT"

	//SCENE_BLOCK_TYPE_INFORM :消息发送单元
	SCENE_BLOCK_TYPE_INFORM BlockType = "SCENE_BLOCK_TYPE_INFORM"

	//SCENE_BLOCK_TYPE_REQUEST :询问填槽单元
	SCENE_BLOCK_TYPE_REQUEST BlockType = "SCENE_BLOCK_TYPE_REQUEST"

	//SCENE_BLOCK_TYPE_END :意图终点单元
	SCENE_BLOCK_TYPE_END BlockType = "SCENE_BLOCK_TYPE_END"
)

/****************
- 场景类
****************/

//SceneListResponse 场景列表结果
type SceneListResponse struct {
	Scenes []Scene `json:"scenes"` //任务列表
}

//Scene 场景
type Scene struct {
	ID                        int              `json:"id"`                           //场景ID
	Name                      string           `json:"name"`                         //场景名称
	Description               string           `json:"description"`                  //场景描述
	IntentSwitchMode          IntentSwitchMode `json:"intent_switch_mode"`           //意图切换模式.在意图流程中，当用户消息既可以在当前意图中填槽、又可以触发其他意图时，优先选择的处理方式
	SmartSlotFillingThreshold float64          `json:"smart_slot_filling_threshold"` //智能填槽阈值
}

//SceneResponse 场景操作的结果
type SceneResponse struct {
	Scene Scene `json:"scene"`
}

/****************
- 词槽类
****************/

//SceneSlotListResponse 词槽列表的结果
type SceneSlotListResponse struct {
	Slots []Slot `json:"slots"` //词槽概况列表
}

//SceneSlotResponse 获取词槽的结果
type SceneSlotResponse struct {
	Slot Slot `json:"slot"`
}

//Slot 词槽
type Slot struct {
	ID               int    `json:"id"`                 //词槽ID
	Name             string `json:"name"`               //词槽名称
	SceneID          int    `json:"scene_id"`           //词槽所属的场景ID
	QuerySlotFilling bool   `json:"query_slot_filling"` //是否允许整句填槽,默认关闭.true: 开启;false: 关闭
}

//SceneSlotDataSourceListResponse 查询词槽数据来源的结果
type SceneSlotDataSourceListResponse struct {
	DataSources []DataSource `json:"data_sources"` //词槽数据来源
}

//SceneSlotDataSourceResponse 创建词槽数据来源的结果
type SceneSlotDataSourceResponse struct {
	DataSource DataSource `json:"data_source"`
}

//DataSource 词槽数据来源
type DataSource struct {
	ID       int `json:"id"`        //词槽数据来源ID
	SlotID   int `json:"slot_id"`   //词槽ID
	EntityID int `json:"entity_id"` //实体ID
}

/****************
- 触发器
****************/

//SceneIntentTriggerListResponse 触发器列表结果
type SceneIntentTriggerListResponse struct {
	IntentTriggers []IntentTrigger `json:"intent_triggers"`
}

//SceneIntentTriggerResponse 更新触发器结果
type SceneIntentTriggerResponse struct {
	IntentTrigger IntentTrigger `json:"intent_trigger"`
}

//IntentTrigger 触发器
type IntentTrigger struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	IntentID int    `json:"intent_id"`
	Type     string `json:"type"`
}

/****************
- 意图
****************/

//SceneIntentResponse 操作意图结果
type SceneIntentResponse struct {
	Intent Intent `json:"intent"`
}

//SceneIntentListResponse 获取意图列表结果
type SceneIntentListResponse struct {
	Intents []Intent `json:"intents"`
}

//Intent 意图
type Intent struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SceneID      int    `json:"scene_id"`      //意图所属场景ID
	Status       bool   `json:"status"`        //意图发布状态 false: 未生效,true: 已生效
	LifespanMins int    `json:"lifespan_mins"` //意图闲置等待时长（分钟），默认3分钟
}

//SceneIntentStatusUpdateResponse 更新意图状态返回的结构体
type SceneIntentStatusUpdateResponse struct {
	IntentID     int    `json:"intent_id"`
	Status       bool   `json:"status"`
	FirstBlockID int    `json:"first_block_id"`
	UpdateTime   string `json:"update_time"`
}

/****************
- 单元类
****************/

//SceneBlockListResponse 查询单元列表的结果
type SceneBlockListResponse struct {
	Blocks []Block `json:"blocks"`
}

//Block 单元
type Block struct {
	ID   int       `json:"id"`   //单元ID
	Name string    `json:"name"` //单元名称
	Type BlockType `json:"type"` //对话单元类型
}

/****************
- 询问填槽单元
****************/

//BlockRequestParam 创建/更新 询问填槽单元的参数
type BlockRequestParam struct {
	ID                   int          `json:"id"`                      //单元ID >= 1
	IntentID             int          `json:"intent_id"`               //所属意图ID >= 1
	Name                 string       `json:"name"`                    //单元名称[1-200]characters
	DefaultSlotValue     string       `json:"default_slot_value"`      //默认词槽值 <= 200 characters
	SlotFillingWhenAsked bool         `json:"slot_filling_when_asked"` //是否仅询问时填槽,默认否.true: 仅在当前单元询问时才填充关联词槽;false: 即使机器人并没有询问，如果用户消息中有可以填槽的信息，也填充关联词槽
	SlotID               int          `json:"slot_id"`                 //绑定的词槽ID >= 1
	Mode                 ResponseType `json:"mode"`                    //单元回复类型
	RequestCount         int          `json:"request_count"`           //询问次数,默认3次 <= 200
}

//BlockRequestResponse 创建/更新 询问填槽单元响应结构体
type BlockRequestResponse struct {
	Block BlockRequest `json:"block"`
}

//BlockRequest 询问填槽单元的结构体
type BlockRequest struct {
	Name                 string           `json:"name"`                    //单元名称[1-200]characters
	DefaultSlotValue     string           `json:"default_slot_value"`      //默认词槽值
	SlotFillingWhenAsked bool             `json:"slot_filling_when_asked"` //是否仅询问时填槽
	Connections          []Connection     `json:"connections"`             //单元跳转关系
	SlotID               int              `json:"slot_id"`                 //绑定的词槽ID
	Mode                 string           `json:"mode"`                    //单元回复类型
	RequestCount         int              `json:"request_count"`           //询问次数,默认3次
	IntentID             int              `json:"intent_id"`               //所属意图ID
	ID                   int              `json:"id"`                      //单元ID
	Responses            []BlockResponses `json:"responses"`               //单元内回复
}

/****************
- 消息发送单元
****************/

//SceneBlockInformResponse 获取消息发送单元的响应
type SceneBlockInformResponse struct {
	Block BlockInform `json:"block"`
}

//BlockInform 消息发送单元
type BlockInform struct {
	IntentID   int              `json:"intent_id"`  //所属意图ID
	ID         int              `json:"id"`         //单元ID
	Name       string           `json:"name"`       //单元名称
	Responses  []BlockResponses `json:"responses"`  //单元内回复
	Connection Connection       `json:"connection"` //单元关系
	Mode       ResponseType     `json:"mode"`       //单元回复类型
}

//BlockResponses  单元内回复
type BlockResponses struct {
	Text      *Text      `json:"text"`       //文本消息
	Image     *Image     `json:"image"`      //图片消息
	Custom    *Custom    `json:"custom"`     //自定义消息
	RichText  *RichText  `json:"rich_text"`  //图文消息
	Video     *Video     `json:"video"`      //视频消息
	File      *File      `json:"file"`       //文件消息
	Voice     *Voice     `json:"voice"`      //语音消息
	Event     *Event     `json:"event"`      //事件消息
	ShareLink *ShareLink `json:"share_link"` //卡片消息
}

//Connection 单元关系
type Connection struct {
	FromBlockID int       `json:"from_block_id"` //当前单元ID
	ToBlockID   int       `json:"to_block_id"`   //下一个单元ID
	Condition   Condition `json:"condition"`     //单元跳转条件(默认 / 大于 / 大于等于 / 小于 / 小于等于 / 等于 / 不等于 / 包含 / 不包含 / 属于实体 / 不属于实体 / 符合正则 / 不符合正则)
}

//Condition 单元跳转条件
type Condition struct {
	Default              *Default              `json:"default"`                  //单元跳转条件 默认
	InEntity             *InEntity             `json:"in_entity"`                //单元跳转条件 属于(实体)
	EqualTo              *EqualTo              `json:"equal_to"`                 //单元跳转条件 等于
	NotEqualTo           *NotEqualTo           `json:"not_equal_to"`             //单元跳转条件 不等于
	LessThanOrQqualTo    *LessThanOrQqualTo    `json:"less_than_or_equal_to"`    //单元跳转条件 小于等于
	NotInEntity          *NotInEntity          `json:"not_in_entity"`            //单元跳转条件 不属于(实体)
	GreaterThanOrEqualTo *GreaterThanOrEqualTo `json:"greater_than_or_equal_to"` //单元跳转条件 大于等于
	DismatchRegex        *DismatchRegex        `json:"dismatch_regex"`           //单元跳转条件 不符合正则
	GreaterThan          *GreaterThan          `json:"greater_than"`             //单元跳转条件 大于
	Exclude              *Exclude              `json:"exclude"`                  //单元跳转条件 不包含
	LessThan             *LessThan             `json:"less_than"`                //单元跳转条件 小于
	Include              *Include              `json:"include"`                  //单元跳转条件 包含
	MatchRegex           *MatchRegex           `json:"match_regex"`              //单元跳转条件 符合正则
}

//Default 单元跳转条件 默认
type Default struct {
}

//InEntity 单元跳转条件 属于(实体)
type InEntity struct {
	ID int `json:"id"` //实体ID
}

//EqualTo 单元跳转条件 等于
type EqualTo struct {
	Value string `json:"value"` //值
}

//NotEqualTo 单元跳转条件 不等于
type NotEqualTo struct {
	Value string `json:"value"` //值
}

//LessThanOrQqualTo 小于等于
type LessThanOrQqualTo struct {
	Value float32 `json:"value"` //值
}

//NotInEntity 单元跳转条件 不属于(实体)
type NotInEntity struct {
	ID int `json:"id"` //实体ID >=1
}

//GreaterThanOrEqualTo 小于等于
type GreaterThanOrEqualTo struct {
	Value float32 `json:"value"` //值
}

//DismatchRegex 单元跳转条件 不符合正则
type DismatchRegex struct {
	Regex string `json:"regex"` //正则表达式
}

//MatchRegex 单元跳转条件 符合正则
type MatchRegex struct {
	Regex string `json:"regex"` //正则表达式
}

//GreaterThan 大于
type GreaterThan struct {
	Value float32 `json:"value"` //值
}

//Exclude 单元跳转条件 不等于
type Exclude struct {
	Value string `json:"value"` //值
}

//Include 单元跳转条件 包含
type Include struct {
	Value string `json:"value"` //值
}

//LessThan 小于
type LessThan struct {
	Value float32 `json:"value"` //值
}

/****************
- 单元关系
****************/

//SceneBlockRelationResponse 创建单元关系响应值
type SceneBlockRelationResponse struct {
	Relation Relation `json:"relation"`
}

//Relation 单元关系
type Relation struct {
	Connection Connection `json:"connection"` //单元关系
	IntentID   int        `json:"intent_id"`  //意图ID
	ID         int        `json:"id"`         //单元关系ID
}

//Relation2 单元关系参数
type Relation2 struct {
	Connection Connection `json:"connection"` //单元关系
	IntentID   int        `json:"intent_id"`  //意图ID
}

/****************
- 任务待审核消息
****************/

//SceneIntentTriggerLearningListResponse 查询任务待审核消息列表
type SceneIntentTriggerLearningListResponse struct {
	QueryItems []QueryItem `json:"query_items"` //待审核问题列表
}

//QueryItem 待审核问题
type QueryItem struct {
	Content         string          `json:"content"`          //用户消息内容
	ID              int             `json:"id"`               //待审核消息ID
	RecommendIntent RecommendIntent `json:"recommend_intent"` //推荐意图
}

//RecommendIntent 推荐意图
type RecommendIntent struct {
	IntentID   int     `json:"intent_id"`   //意图ID
	Score      float64 `json:"score"`       //置信度
	IntentName string  `json:"intent_name"` //意图名称
}

/****************
- 单元内回复
****************/

//SceneBlockCreateResponse 创建单元内回复
type SceneBlockCreateResponse struct {
	Response SceneBlockResponse `json:"response"` //单元内回复
}

//SceneBlockUpdateResponse 更新单元内回复
type SceneBlockUpdateResponse struct {
	Response SceneBlockResponse `json:"response"` //单元内回复
}

//SceneBlockResponse 单元内回复
type SceneBlockResponse struct {
	Response BlockResponses `json:"response"` //消息体格式
	ID       int            `json:"id"`       //回复ID
	BlockID  int            `json:"block_id"` //单元ID
}

/****************
- 意图终点单元
****************/

//SceneBlockEndBlockParam 创建意图终点单元的参数
type SceneBlockEndBlockParam struct {
	Block BlockEnd `json:"block"`
}

//SceneBlockEndBlockResponse 创建/更新 意图终点单元
type SceneBlockEndBlockResponse struct {
	Block BlockEnd `json:"block"`
}

//BlockEnd 意图终点单元
type BlockEnd struct {
	IntentID       int    `json:"intent_id"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SlotMemorizing bool   `json:"slot_memorizing"`
	Action         Action `json:"action"`
}

//Action 结束单元跳转方式
type Action struct {
	Last      Last      `json:"last"`      //跳转上个意图
	End       End       `json:"end"`       //不跳转
	Specified Specified `json:"specified"` //跳转指定意图
}

//Last 意图终点单元跳转上个意图
type Last struct {
}

//End 意图终点单元不跳转
type End struct {
}

//Specified 意图终点单元跳转指定意图
type Specified struct {
	ID int `json:"id"` //意图ID >=1
}
