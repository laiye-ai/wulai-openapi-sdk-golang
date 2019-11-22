package wulai

import (
	"fmt"
)

/****************
- 场景类
****************/

//sceneListV2 查询场景列表
func (x *Client) sceneListV2() ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/list", x.Endpoint, x.Version)
	input := "{}"
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneCreateV2 创建场景
@name：场景名称.[1-200]characters
@description：场景描述 <= 600 characters
@intentSwitchMode：意图切换模式
@smartSlotFillingThreshold：智能填槽阈值 <= 1
*/
func (x *Client) sceneCreateV2(name, description string, intentSwitchMode IntentSwitchMode, smartSlotFillingThreshold int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"scene": {
		  "intent_switch_mode": "%s",
		  "name": "%s",
		  "smart_slot_filling_threshold": %v,
		  "description": %q
		}
	}`, intentSwitchMode, name, smartSlotFillingThreshold, description)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneUpdateV2 更新场景
@scene：场景类型
*/
func (x *Client) sceneUpdateV2(scene *Scene) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"scene": {
		  "intent_switch_mode": "%s",
		  "name": "%s",
		  "id": %v,
		  "smart_slot_filling_threshold": %v,
		  "description": %q
		}
	}`, scene.IntentSwitchMode, scene.Name, scene.ID, scene.SmartSlotFillingThreshold, scene.Description)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//sceneDeleteV2 删除场景
func (x *Client) sceneDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 意图
****************/

/*sceneIntentListV2 查询意图列表
@sceneID：意图所属的场景ID>=1
*/
func (x *Client) sceneIntentListV2(sceneID int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"scene_id": %v
	}`, sceneID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentCreate2 创建意图
@sceneID：意图所属场景ID >=1
@name：意图名称[1-200]characters
@lifespanMins：意图闲置等待时长（分钟），默认3分钟) <= 60
*/
func (x *Client) sceneIntentCreateV2(sceneID int, name string, lifespanMins int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"intent": {
		  "scene_id": %v,
		  "name": "%s",
		  "lifespan_mins": %v
		}
	}`, sceneID, name, lifespanMins)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentDeleteV2 删除意图
@id：意图ID >=1
*/
func (x *Client) sceneIntentDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentUpdateV2 更新意图
@id：意图ID >=1
@name：意图名称[1-200]characters
@lifespanMins：意图闲置等待时长（分钟），默认3分钟) <= 60
*/
func (x *Client) sceneIntentUpdateV2(id int, name string, lifespanMins int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"intent": {
		  "id": %v,
		  "name": "%s",
		  "lifespan_mins": %v
		}
	}`, id, name, lifespanMins)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentStatusUpdateV2 更新意图状态
@intentID：意图ID >=1
@status：意图状态,意图生成时默认未生效.false: 未生效,true: 已生效
@firstBlockID：该意图的第一个单元ID >=1
*/
func (x *Client) sceneIntentStatusUpdateV2(intentID int, status bool, firstBlockID int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/status/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"intent_id": %v,
		"status": %v,
		"first_block_id": %v
	}`, intentID, status, firstBlockID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 词槽类
****************/

/*sceneSlotListV2 查询词槽列表
@sceneID：词槽所属的场景ID >=1
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) sceneSlotListV2(sceneID, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"scene_id": %v,
		"page": %v,
		"page_size": %v
	}`, sceneID, page, pageSize)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneSlotCreateV2 创建词槽
@sceneID：词槽所属的场景ID >=1
@name：词槽名称 [1-200]characters
@querySlotFilling：是否允许整句填槽,默认关闭.true: 开启;false: 关闭
*/
func (x *Client) sceneSlotCreateV2(sceneID int, name string, querySlotFilling bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"slot": {
		  "scene_id": %v,
		  "name": "%s",
		  "query_slot_filling": %v
		}
	}`, sceneID, name, querySlotFilling)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneSlotGetV2 查询词槽
@id：词槽ID >=1
*/
func (x *Client) sceneSlotGetV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/get", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"id": %v
	}`, id)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneSlotUpdateV2 更新词槽
@id：词槽ID >=1
@name：词槽名称 [1-200]characters
@querySlotFilling：是否允许整句填槽,默认关闭.true: 开启;false: 关闭
*/
func (x *Client) sceneSlotUpdateV2(id int, name string, querySlotFilling bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"slot": {
		  "id": %v,
		  "name": "%s",
		  "query_slot_filling": %v
		}
	}`, id, name, querySlotFilling)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneSlotDeleteV2 删除词槽
@id：词槽ID >=1
*/
func (x *Client) sceneSlotDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 词槽数据来源
****************/

/*sceneSlotDataSourceListV2 查询词槽数据来源
@slotID：词槽ID >=1
*/
func (x *Client) sceneSlotDataSourceListV2(slotID int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/data-source/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"slot_id": %v
	}`, slotID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneSlotDataSourceCreateV2 创建词槽数据来源
@slotID：词槽ID >=1
@entityID：实体ID >=1
*/
func (x *Client) sceneSlotDataSourceCreateV2(slotID, entityID int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/data-source/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"data_source": {
		  "entity_id": %v,
		  "slot_id": %v
		}
	}`, entityID, slotID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneSlotDataSourceDeleteV2 删除词槽数据来源
@id：词槽数据来源ID >=1
*/
func (x *Client) sceneSlotDataSourceDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/slot/data-source/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 触发器
****************/

/*sceneIntentTriggerListV2 查询触发器列表
@intentID：意图ID
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) sceneIntentTriggerListV2(intentID, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/trigger/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"intent_id": %v,
		"page": %v,
		"page_size": %v
	}`, intentID, page, pageSize)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentTriggerCreate2 创建触发器
@intentID：意图ID
@text：触发文本[1-200]characters
@triggerType：触发器模式
*/
func (x *Client) sceneIntentTriggerCreateV2(intentID int, text string, triggerType TriggerType) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/trigger/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"intent_trigger": {
		  "text": %q,
		  "intent_id": %v,
		  "type": "%s"
		}
	}`, text, intentID, triggerType)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentTriggerUpdateV2 更新触发器
@triggerID：触发器ID >=1
@text：触发文本[1-200]characters
*/
func (x *Client) sceneIntentTriggerUpdateV2(triggerID int, text string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/trigger/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"intent_trigger": {
		  "text": %q,
		  "id": %v
		}
	}`, text, triggerID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentTriggerDeleteV2 删除触发器
@id：触发器ID >=1
*/
func (x *Client) sceneIntentTriggerDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/trigger/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 单元类
****************/

/*sceneBlockListV2 查询单元列表
@intentID：意图ID
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/

func (x *Client) sceneBlockListV2(intentID, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"intent_id": %v,
		"page": %v,
		"page_size": %v
	}`, intentID, page, pageSize)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockDeleteV2 删除单元
@id：单元ID >=1
*/
func (x *Client) sceneBlockDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 询问填槽单元
****************/

/*sceneBlockRequestCreateV2 创建询问填槽单元
@param：参数 BlockRequestParam(指针类型)
*/
func (x *Client) sceneBlockRequestCreateV2(param *BlockRequestParam) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/request-block/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"block": {
		  "name": "%s",
		  "default_slot_value": "%s",
		  "slot_filling_when_asked": %v,
		  "slot_id": %v,
		  "mode": "%s",
		  "request_count": %v,
		  "intent_id": %v
		}
	}`, param.Name, param.DefaultSlotValue, param.SlotFillingWhenAsked, param.SlotID, param.Mode, param.RequestCount, param.IntentID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockRequestGetV2 查询询问填槽单元
@id：单元ID >=1
*/
func (x *Client) sceneBlockRequestGetV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/request-block/get", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"id": %v
	}`, id)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockRequestUpdateV2 更新询问填槽单元
@param：参数 BlockRequestParam(指针类型)
*/
func (x *Client) sceneBlockRequestUpdateV2(param *BlockRequestParam) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/request-block/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"block": {
		  "name": "%s",
		  "default_slot_value": "%s",
		  "slot_filling_when_asked": %v,
		  "slot_id": %v,
		  "mode": "%s",
		  "request_count": %v,
		  "intent_id": %v,
		  "id": %v
		}
	}`, param.Name, param.DefaultSlotValue, param.SlotFillingWhenAsked, param.SlotID, param.Mode, param.RequestCount, param.IntentID, param.ID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 消息发送单元
****************/

/*sceneBlockInformCreate2 创建消息发送单元
@intentID：所属意图ID>=1
@name：单元名称[1-200]characters
@mode：单元回复类型
*/
func (x *Client) sceneBlockInformCreateV2(intentID int, name string, mode ResponseType) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/inform-block/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"block": {
		  "intent_id": %v,
		  "name": "%s",
		  "mode": "%s"
		}
	}`, intentID, name, mode)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockInformUpdateV2 更新消息发送单元
@id：单元ID>=1
@name：单元名称[1-200]characters
@mode：单元回复类型
*/
func (x *Client) sceneBlockInformUpdateV2(id int, name string, mode ResponseType) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/inform-block/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"block": {
		  "id": %v,
		  "name": "%s",
		  "mode": "%s"
		}
	}`, id, name, mode)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockInformGetV2 查询消息发送单元
@blockID：单元ID>=1
*/
func (x *Client) sceneBlockInformGetV2(blockID int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/inform-block/get", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"id": %v
	}`, blockID)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 任务待审核消息
****************/

/*sceneIntentTriggerLearningListV2 查询任务待审核消息列表
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) sceneIntentTriggerLearningListV2(page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/trigger-learning/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"page": %v,
		"page_size": %v
	}`, page, pageSize)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneIntentTriggerLearningDeleteV2 删除任务待审核消息
@id：待审核消息ID >=1
*/
func (x *Client) sceneIntentTriggerLearningDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/intent/trigger-learning/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 单元关系
****************/

/*sceneBlockRelationCreateV2 创建单元关系
@intentID：意图ID >=1
@fromBlockID：当前单元ID >=1
@toBlockID：下一个单元ID >=1
@msgType：条件类型
@msgBody：条件值
*/
func (x *Client) sceneBlockRelationCreateV2(intentID, fromBlockID, toBlockID int, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/relation/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"relation": {
		  "connection": {
			"from_block_id": %v,
			"to_block_id": %v,
			"condition": {"%s": %s}
		  },
		  "intent_id": %v
		}
	}`, fromBlockID, toBlockID, msgType, msgBody, intentID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockRelationDeleteV2 删除单元关系
@id：单元关系ID >=1
*/
func (x *Client) sceneBlockRelationDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/relation/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 单元内回复
****************/

/*sceneBlockResponseCreateV2 创建单元内回复
@blockID：单元ID>=1
@msgType：消息类型
@msgBody：消息体格式，任意选择一种消息类型[文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息]
*/
func (x *Client) sceneBlockResponseCreateV2(blockID int, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/response/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"response": {
		  "block_id": %v,
		  "response": {"%s": %s}
		}
	}`, blockID, msgType, msgBody)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockResponseUpdateV2 更新单元内回复
@id：回复ID >=1
@msgType：消息类型
@msgBody：消息体格式，任意选择一种消息类型[文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息]
*/
func (x *Client) sceneBlockResponseUpdateV2(id int, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/response/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"response": {
		  "id": %v,
		  "response": {"%s": %s}
		}
	}`, id, msgType, msgBody)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockResponseDeleteV2 删除单元内回复
@id：回复ID >=1
*/
func (x *Client) sceneBlockResponseDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/response/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": %v}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 意图终点单元
****************/

/*sceneBlockEndBlockCreateV2 创建意图终点单元
@intentID：所属意图ID>=1
@name：单元名称[1-200]characters
@slot_memorizing：是否保存词槽值(默认关闭).true: 开启;false: 关闭
@actionType：结束单元跳转方式
@actionBody：结束单元跳转方式 (指定意图 / 上个意图 / 不跳转))
*/
func (x *Client) sceneBlockEndBlockCreateV2(intentID int, name string, slotMemorizing bool, actionType string, actionBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/end-block/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"block": {
		  "action": {"%s": %s},
		  "intent_id": %v,
		  "name": "%s",
		  "slot_memorizing": %v
		}
	}`, actionType, actionBody, intentID, name, slotMemorizing)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockEndBlockCreateV2 更新意图终点单元
@intentID：所属意图ID>=1
@id：单元ID>=1
@name：单元名称[1-200]characters
@slot_memorizing：是否保存词槽值(默认关闭).true: 开启;false: 关闭
@actionType：结束单元跳转方式
@actionBody：结束单元跳转方式 (指定意图 / 上个意图 / 不跳转))
*/
func (x *Client) sceneBlockEndBlockUpdateV2(intentID int, id int, name string, slotMemorizing bool, actionType string, actionBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/end-block/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"block": {
		  "action": {"%s": %s},
		  "intent_id": %v,
		  "name": "%s",
		  "id": %v,
		  "slot_memorizing": %v
		}
	}`, actionType, actionBody, intentID, name, id, slotMemorizing)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*sceneBlockEndBlockGetV2 查询意图终点单元
@id：单元ID>=1
*/
func (x *Client) sceneBlockEndBlockGetV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/scene/block/end-block/get", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"id": %v
	}`, id)
	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
