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
	SmartSlotFillingThreshold int              `json:"smart_slot_filling_threshold"` //智能填槽阈值
}

//SceneResponse 场景操作的结果
type SceneResponse struct {
	Scene Scene `json:"scene"`
}

/****************
- 词槽类
****************/

/****************
- 触发器
****************/

//SceneIntentTriggerResponse 触发器列表结果
type SceneIntentTriggerResponse struct {
	IntentTriggers []IntentTrigger `json:"intent_triggers"`
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

//SceneIntentResponse 意图结果
type SceneIntentResponse struct {
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

/****************
- 单元内回复
****************/

/****************
- 任务待审核消息
****************/
