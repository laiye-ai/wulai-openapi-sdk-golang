package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 场景类
****************/

//SceneList 查询场景列表
func (x *Client) SceneList() (model *SceneListResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneListV2()
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneList Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneCreate 创建场景
@name：场景名称.[1-200]characters
@description：场景描述 <= 600 characters
@intentSwitchMode：意图切换模式
@smartSlotFillingThreshold：智能填槽阈值 <= 1
*/
func (x *Client) SceneCreate(name, description string, intentSwitchMode IntentSwitchMode, smartSlotFillingThreshold int) (model *SceneResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneCreateV2(name, description, intentSwitchMode, smartSlotFillingThreshold)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneUpdate 更新场景
@scene：场景类型
*/
func (x *Client) SceneUpdate(scene *Scene) (model *SceneResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneUpdateV2(scene)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//SceneDelete 删除场景
func (x *Client) SceneDelete(id int) (err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneDeleteV2(id)

	if x.Debug {
		log.Debugf("[SceneDelete Response]:%s\n", bytes)
	}

	if err != nil {
		return err
	}

	return nil
}

/****************
- 意图
****************/

/*SceneIntentList 查询意图列表
@sceneID：意图所属的场景ID>=1
*/
func (x *Client) SceneIntentList(sceneID int) (model *SceneIntentResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentListV2(sceneID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentList Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 词槽类
****************/

/****************

- 触发器
****************/

/*SceneIntentTriggerList 查询触发器列表
@intentID：意图ID >=1
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) SceneIntentTriggerList(intentID, page, pageSize int) (model *SceneIntentTriggerResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentTriggerListV2(intentID, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentTriggerList Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentTriggerResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneIntentTriggerCreate 创建触发器
@intentID：意图ID
@text：触发文本[1-200]characters
@triggerType：触发器模式
*/
func (x *Client) SceneIntentTriggerCreate(intentID int, text string, triggerType TriggerType) (model *SceneIntentTriggerResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentTriggerCreate2(intentID, text, triggerType)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentTriggerCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentTriggerResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 单元内回复
****************/

/****************
- 任务待审核消息
****************/
