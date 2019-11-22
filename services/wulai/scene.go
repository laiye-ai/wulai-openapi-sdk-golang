package wulai

import (
	"encoding/json"
	"fmt"
	"reflect"
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
func (x *Client) SceneIntentList(sceneID int) (model *SceneIntentListResponse, err error) {

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
	model = &SceneIntentListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneIntentCreate 创建意图
@intentID：意图所属场景ID >=1
@name：意图名称[1-200]characters
@lifespanMins：意图闲置等待时长（分钟），默认3分钟) <= 60
*/
func (x *Client) SceneIntentCreate(sceneID int, name string, lifespanMins int) (model *SceneIntentResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentCreateV2(sceneID, name, lifespanMins)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneIntentUpdate 更新意图
@id：意图ID >=1
@name：意图名称[1-200]characters
@lifespanMins：意图闲置等待时长（分钟），默认3分钟) <= 60
*/
func (x *Client) SceneIntentUpdate(id int, name string, lifespanMins int) (model *SceneIntentResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentUpdateV2(id, name, lifespanMins)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneIntentDelete 删除意图
@id：意图ID >=1
*/
func (x *Client) SceneIntentDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneIntentDelete Response]:%s\n", bytes)
	}

	return nil
}

/*SceneIntentStatusUpdate 更新意图状态
@intentID：意图ID >=1
@status：意图状态,意图生成时默认未生效.false: 未生效,true: 已生效
@firstBlockID：该意图的第一个单元ID >=1
*/
func (x *Client) SceneIntentStatusUpdate(intentID int, status bool, firstBlockID int) (model *SceneIntentStatusUpdateResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentStatusUpdateV2(intentID, status, firstBlockID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentStatusUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentStatusUpdateResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 词槽类
****************/

/*SceneSlotList 查询词槽列表
@sceneID：词槽所属的场景ID >=1
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) SceneSlotList(sceneID, page, pageSize int) (model *SceneSlotListResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotListV2(sceneID, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneSlotList Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneSlotListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneSlotCreate 创建词槽
@sceneID：词槽所属的场景ID >=1
@name：词槽名称 [1-200]characters
@querySlotFilling：是否允许整句填槽,默认关闭.true: 开启;false: 关闭
*/
func (x *Client) SceneSlotCreate(sceneID int, name string, querySlotFilling bool) (model *SceneSlotResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotCreateV2(sceneID, name, querySlotFilling)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneSlotCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneSlotResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneSlotGet 查询词槽
@id：词槽ID >=1
*/
func (x *Client) SceneSlotGet(id int) (model *SceneSlotResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotGetV2(id)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneSlotGet Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneSlotResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneSlotUpdate 更新词槽
@id：词槽ID >=1
@name：词槽名称 [1-200]characters
@querySlotFilling：是否允许整句填槽,默认关闭.true: 开启;false: 关闭
*/
func (x *Client) SceneSlotUpdate(id int, name string, querySlotFilling bool) (model *SceneSlotResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotUpdateV2(id, name, querySlotFilling)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneSlotUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneSlotResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneSlotDelete 删除词槽
@id：词槽ID >=1
*/
func (x *Client) SceneSlotDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneSlotDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- 词槽数据来源
****************/

/*SceneSlotDataSourceList 查询词槽数据来源
@slotID：词槽ID >=1
*/
func (x *Client) SceneSlotDataSourceList(slotID int) (model *SceneSlotDataSourceListResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotDataSourceListV2(slotID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneSlotUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneSlotDataSourceListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneSlotDataSourceCreate 创建词槽数据来源
@slotID：词槽ID >=1
@entityID：实体ID >=1
*/
func (x *Client) SceneSlotDataSourceCreate(slotID, entityID int) (model *SceneSlotDataSourceResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotDataSourceCreateV2(slotID, entityID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneSlotDataSourceCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneSlotDataSourceResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneSlotDataSourceDelete 删除词槽数据来源
@id：词槽数据来源ID >=1
*/
func (x *Client) SceneSlotDataSourceDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneSlotDataSourceDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneSlotDataSourceDelete Response]:%s\n", bytes)
	}

	return nil
}

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
func (x *Client) SceneIntentTriggerCreate(intentID int, text string, triggerType TriggerType) (model *SceneIntentTriggerListResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentTriggerCreateV2(intentID, text, triggerType)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentTriggerCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentTriggerListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneIntentTriggerUpdate 更新触发器
@triggerID：触发器ID >=1
@text：触发文本[1-200]characters
*/
func (x *Client) SceneIntentTriggerUpdate(triggerID int, text string) (model *SceneIntentTriggerResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentTriggerUpdateV2(triggerID, text)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentTriggerUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentTriggerResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneIntentTriggerDelete 删除触发器
@triggerID：触发器ID >=1
*/
func (x *Client) SceneIntentTriggerDelete(triggerID int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentTriggerDeleteV2(triggerID)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneIntentTriggerDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- 单元类
****************/

/*SceneBlockList 查询单元列表
@intentID：意图ID
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) SceneBlockList(intentID, page, pageSize int) (model *SceneBlockListResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockListV2(intentID, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockList Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockDelete 删除单元
@id：单元ID >=1
*/
func (x *Client) SceneBlockDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneBlockDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- 询问填槽单元
****************/

/*SceneBlockRequestGet 查询询问填槽单元
@id：单元ID >=1
*/
func (x *Client) SceneBlockRequestGet(id int) (model *BlockRequestResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockRequestGetV2(id)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockRequestGet Response]:%s\n", bytes)
	}

	//返回结果
	model = &BlockRequestResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockRequestCreate 创建询问填槽单元
@param：参数 BlockRequest(指针类型)
*/
func (x *Client) SceneBlockRequestCreate(param *BlockRequestParam) (model *BlockRequestResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockRequestCreateV2(param)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockRequestCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &BlockRequestResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockRequestUpdate 更新询问填槽单元
@param：参数 BlockRequest(指针类型)
*/
func (x *Client) SceneBlockRequestUpdate(param *BlockRequestParam) (model *BlockRequestResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockRequestUpdateV2(param)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockRequestUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &BlockRequestResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 消息发送单元
****************/

/*SceneBlockInformGet 查询消息发送单元
@blockID：单元ID>=1
*/
func (x *Client) SceneBlockInformGet(blockID int) (model *SceneBlockInformResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockInformGetV2(blockID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[sceneBlockInformGet Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockInformResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockInformCreate 创建消息发送单元
@intentID：所属意图ID>=1
@name：单元名称[1-200]characters
@mode：单元回复类型
*/
func (x *Client) SceneBlockInformCreate(intentID int, name string, mode ResponseType) (model *SceneBlockInformResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockInformCreateV2(intentID, name, mode)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockInformCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockInformResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockInformUpdate 更新消息发送单元
@intentID：所属意图ID>=1
@name：单元名称[1-200]characters
@mode：单元回复类型
*/
func (x *Client) SceneBlockInformUpdate(id int, name string, mode ResponseType) (model *SceneBlockInformResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockInformUpdateV2(id, name, mode)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockInformUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockInformResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 任务待审核消息
****************/

/*SceneIntentTriggerLearningList 查询任务待审核消息列表
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) SceneIntentTriggerLearningList(page, pageSize int) (model *SceneIntentTriggerLearningListResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentTriggerLearningListV2(page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneIntentTriggerLearningList Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneIntentTriggerLearningListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneIntentTriggerLearningDelete 删除任务待审核消息
@id：待审核消息ID >=1
*/
func (x *Client) SceneIntentTriggerLearningDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneIntentTriggerLearningDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneIntentTriggerLearningDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- 单元内回复
****************/

/*SceneBlockResponseCreate 创建单元内回复
@blockID：单元ID>=1
@msgBody：消息体格式，任意选择一种消息类型[文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息]
*/
func (x *Client) SceneBlockResponseCreate(blockID int, msgBody interface{}) (model *SceneBlockCreateResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	//检查 json format
	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.sceneBlockResponseCreateV2(blockID, msgType, msgBytes)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockResponseCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockCreateResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockResponseUpdate 更新单元内回复
@id：回复ID >=1
@msgBody：消息体格式，任意选择一种消息类型[文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息]
*/
func (x *Client) SceneBlockResponseUpdate(id int, msgBody interface{}) (model *SceneBlockUpdateResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	//检查 json format
	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.sceneBlockResponseUpdateV2(id, msgType, msgBytes)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockResponseUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockUpdateResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockResponseDelete 删除单元内回复
@id：回复ID >=1
*/
func (x *Client) SceneBlockResponseDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockResponseDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneBlockResponseDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- 单元关系
****************/

/*SceneBlockRelationCreate 创建单元关系
@intentID：意图ID >=1
@fromBlockID：当前单元ID >=1
@toBlockID：下一个单元ID >=1
@msgType：条件类型
@msgBody：条件值
*/
func (x *Client) SceneBlockRelationCreate(intentID, fromBlockID, toBlockID int, condition interface{}) (model *SceneBlockRelationResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckConditionType(condition)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	//msg bytes
	msgBytes, err := json.Marshal(condition)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.sceneBlockRelationCreateV2(intentID, fromBlockID, toBlockID, msgType, msgBytes)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockRelationCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockRelationResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockRelationDelete 删除单元关系
@id：单元关系ID >=1
*/
func (x *Client) SceneBlockRelationDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockRelationDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[SceneBlockRelationDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- 意图终点单元
****************/

/*SceneBlockEndBlockCreate 创建意图终点单元
@intentID：所属意图ID>=1
@name：单元名称[1-200]characters
@slot_memorizing：是否保存词槽值(默认关闭).true: 开启;false: 关闭
@action：结束单元跳转方式 (指定意图 / 上个意图 / 不跳转))
*/
func (x *Client) SceneBlockEndBlockCreate(intentID int, name string, slotMemorizing bool, action interface{}) (model *SceneBlockEndBlockResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查类型是否合法
	actionType, ok := CheckActionType(action)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, actionType, "*"+actionType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	//msg bytes
	actionBytes, err := json.Marshal(action)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.sceneBlockEndBlockCreateV2(intentID, name, slotMemorizing, actionType, actionBytes)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockEndBlockCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockEndBlockResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockEndBlockUpdate 更新意图终点单元
@intentID：所属意图ID>=1
@id：单元ID>=1
@name：单元名称[1-200]characters
@slot_memorizing：是否保存词槽值(默认关闭).true: 开启;false: 关闭
@action：结束单元跳转方式 (指定意图 / 上个意图 / 不跳转))
*/
func (x *Client) SceneBlockEndBlockUpdate(intentID int, id int, name string, slotMemorizing bool, action interface{}) (model *SceneBlockEndBlockResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查类型是否合法
	actionType, ok := CheckActionType(action)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, actionType, "*"+actionType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	//msg bytes
	actionBytes, err := json.Marshal(action)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.sceneBlockEndBlockUpdateV2(intentID, id, name, slotMemorizing, actionType, actionBytes)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockEndBlockUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockEndBlockResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*SceneBlockEndBlockGet 查询意图终点单元
@id：单元ID>=1
*/
func (x *Client) SceneBlockEndBlockGet(blockID int) (model *SceneBlockEndBlockResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.sceneBlockEndBlockGetV2(blockID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[SceneBlockEndBlockGet Response]:%s\n", bytes)
	}

	//返回结果
	model = &SceneBlockEndBlockResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 类型检查
****************/

//CheckConditionType 检查消息类型
func CheckConditionType(msgType interface{}) (string, bool) {

	switch msgType.(type) {
	case *Default:
		return "default", true
	case *InEntity:
		return "in_entity", true
	case *EqualTo:
		return "equal_to", true
	case *NotEqualTo:
		return "not_equal_to", true
	case *LessThanOrQqualTo:
		return "less_than_or_equal_to", true
	case *NotInEntity:
		return "not_in_entity", true
	case *GreaterThanOrEqualTo:
		return "greater_than_or_equal_to", true
	case *DismatchRegex:
		return "dismatch_regex", true
	case *GreaterThan:
		return "greater_than", true
	case *Exclude:
		return "exclude", true
	case *LessThan:
		return "less_than", true
	case *Include:
		return "include", true
	case *MatchRegex:
		return "match_regex", true
	}
	return reflect.TypeOf(msgType).String(), false
}

//CheckActionType 检查结束单元跳转方式
func CheckActionType(msgType interface{}) (string, bool) {

	switch msgType.(type) {
	case *Last:
		return "last", true
	case *End:
		return "end", true
	case *Specified:
		return "specified", true
	}
	return reflect.TypeOf(msgType).String(), false
}
