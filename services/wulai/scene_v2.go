package wulai

import "fmt"

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

/****************
- 词槽类
****************/

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
func (x *Client) sceneIntentTriggerCreate2(intentID int, text string, triggerType TriggerType) ([]byte, error) {
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

/****************
- 单元内回复
****************/

/****************
- 任务待审核消息
****************/
