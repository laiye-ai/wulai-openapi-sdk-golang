package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 知识点
****************/

/*QaKnowledgeTagList 获取知识点分类列表
@parentTagID：父节点分类id，如果值为0，代表获取根节点下的知识点分类
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) QaKnowledgeTagList(parentTagID, page, pageSize int) (model *QaKnowledgeTagList, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaKnowledgeTagListV2(parentTagID, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaKnowledgeTagList Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaKnowledgeTagList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaKnowledgeCreate 创建知识点
@knowledgeTagID：知识点分类id >=1
@standardQuestion：知识点标题 <= 100 characters
@status：知识点状态. [true:生效 false:未生效]
@respondAll：发送全部回复. [true:发送全部回复 false:随机一条发送]
@maintained：该属性是否被用于定义属性组。True：该属性是定义属性组所使用的属性之一。False：该属性不被用于定义属性组。[true:是 false:否]
*/
func (x *Client) QaKnowledgeCreate(knowledgeTagID, standardQuestion string, status, respondAll, maintained bool) (model *QaKnowledgeTagResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaKnowledgeCreateV2(knowledgeTagID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaKnowledgeCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaKnowledgeTagResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaKnowledgeUpdate 更新知识点
@knowledgeID：知识点id
@standardQuestion：知识点标题 <= 100 characters
@status：知识点状态. [true:生效 false:未生效]
@respondAll：发送全部回复. [true:发送全部回复 false:随机一条发送]
@maintained：该属性是否被用于定义属性组。True：该属性是定义属性组所使用的属性之一。False：该属性不被用于定义属性组。[true:是 false:否]
*/
func (x *Client) QaKnowledgeUpdate(knowledgeID int64, standardQuestion string, status, respondAll, maintained bool) (model *QaKnowledgeResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaKnowledgeUpdateV2(knowledgeID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaKnowledgeUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaKnowledgeResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaKnowledgeItemList 获取知识点列表
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) QaKnowledgeItemList(page, pageSize int) (model *QaKnowledgeItemListResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaKnowledgeItemListV2(page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaKnowledgeItemList Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaKnowledgeItemListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 相似问
****************/

/*QaSimilarQuestionList 获取相似问列表
@knowledgeID：知识点id
@similarQuestionID：相似问id
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) QaSimilarQuestionList(knowledgeID, similarQuestionID string, page, pageSize int) (model *QaSimilarQuestionList, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaSimilarQuestionListV2(knowledgeID, similarQuestionID, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaSimilarQuestionList Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaSimilarQuestionList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaSimilarQuestionCreate 创建相似问
@knowledgeID：知识点id
@question：相似问 <=100 characters
*/
func (x *Client) QaSimilarQuestionCreate(knowledgeID, question string) (model *QaSimilarQuestionResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaSimilarQuestionCreateV2(knowledgeID, question)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaSimilarQuestionCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaSimilarQuestionResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaSimilarQuestionUpdate 更新相似问
@knowledgeID：知识点id
@question：相似问 <=100 characters
@id：相识问id
*/
func (x *Client) QaSimilarQuestionUpdate(knowledgeID, question, id string) (model *QaSimilarQuestionResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaSimilarQuestionUpdateV2(knowledgeID, question, id)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaSimilarQuestionUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaSimilarQuestionResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaSimilarQuestionDelete 删除相似问
@id：相似问id
*/
func (x *Client) QaSimilarQuestionDelete(id string) (err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaSimilarQuestionDeleteV2(id)

	if x.Debug {
		log.Debugf("[QaSimilarQuestionDelete Response]:%s\n", bytes)
	}

	if err != nil {
		return err
	}

	return nil
}

/****************
- 用户属性组
****************/

/*QaUserAttributeGroupItemList 查询属性组及属性列表
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) QaUserAttributeGroupItemList(page, pageSize int) (model *QaUserAttributeGroupItemList, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaUserAttributeGroupItemListV2(page, pageSize)
	if err != nil {
		return nil, err
	}
	if x.Debug {
		log.Debugf("[QaUserAttributeGroupItemList Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaUserAttributeGroupItemList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaUserAttributeGroupItemCreate 创建属性组
@groupName:属性组名称 [1~128]characters
@attributeID:用户属性 用户属性ID
@attributeValue:用户属性值 [1~128]characters
*/
func (x *Client) QaUserAttributeGroupItemCreate(groupName, attributeID, attributeValue string) (model *QaUserAttributeGroupResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaUserAttributeGroupItemCreateV2(groupName, attributeID, attributeValue)
	if err != nil {
		return nil, err
	}
	if x.Debug {
		log.Debugf("[QaUserAttributeGroupItemCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaUserAttributeGroupResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaUserAttributeGroupItemUpdate 更新属性组
@groupID：属性组id
@groupName：属性组名称 [1~128] characters
@attributes：用户属性组及属性
*/
func (x *Client) QaUserAttributeGroupItemUpdate(groupID, groupName string, attributes map[string]string) (model *QaUserAttributeGroupResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaUserAttributeGroupItemUpdateV2(groupID, groupName, attributes)
	if err != nil {
		return nil, err
	}
	if x.Debug {
		log.Debugf("[QaUserAttributeGroupItemUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaUserAttributeGroupResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 属性组回复
****************/

/*QaUserAttributeGroupAnswerList 查询属性组回复列表
@knowledgeID：知识点id，如不为空，返回所有知识点
@groupID：属性组id，如不为空，返回所有属性组
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) QaUserAttributeGroupAnswerList(knowledgeID, groupID string, page, pageSize int) (model *QaUserAttributeGroupAnswerList, err error) {

	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaUserAttributeGroupAnswerListV2(knowledgeID, groupID, page, pageSize)
	if err != nil {
		return nil, err
	}
	if x.Debug {
		log.Debugf("[QaUserAttributeGroupAnswerList Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaUserAttributeGroupAnswerList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaUserAttributeGroupAnswerCreate 创建属性组回复
@knowledgeID:知识点id
@groupID:属性组ID
@msgBody:消息（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充)
*/
func (x *Client) QaUserAttributeGroupAnswerCreate(knowledgeID, groupID string, msgBody interface{}) (model *QaUserAttributeGroupAnswer, err error) {

	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "Text/Image/Custom/Video//File/RichText/Voice/Event/ShareLink")
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.qaUserAttributeGroupAnswerCreateV2(knowledgeID, groupID, msgType, msgBytes)

	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaUserAttributeGroupAnswerCreate Response]:%s\n", bytes)
	}

	model = &QaUserAttributeGroupAnswer{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaUserAttributeGroupAnswerUpdate 更新属性组回复
@knowledgeID：知识点id
@groupID：属性组ID
@answerID：回复id
@msgBody：消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）
*/
func (x *Client) QaUserAttributeGroupAnswerUpdate(knowledgeID, groupID string, answerID string, msgBody interface{}) (model *QaUserAttributeGroupAnswer, err error) {

	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "Text/Image/Custom/Video//File/RichText/Voice/Event/ShareLink")
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.qaUserAttributeGroupAnswerUpdateV2(knowledgeID, groupID, answerID, msgType, msgBytes)

	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaUserAttributeGroupAnswerUpdate Response]:%s\n", bytes)
	}

	model = &QaUserAttributeGroupAnswer{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*QaUserAttributeGroupAnswerDelete 删除属性组回复
@answerID：属性组回复ID
*/
func (x *Client) QaUserAttributeGroupAnswerDelete(answerID string) error {

	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaUserAttributeGroupAnswerDeleteV2(answerID)
	if err != nil {
		return err
	}
	if x.Debug {
		log.Debugf("[QaUserAttributeGroupAnswerDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- QA 统计
****************/

/*QaSatisCreate 添加用户满意度评价
@satisType：满意度枚举类型
@userID：用户id
@knowledgeID：知识点id
@msgID：机器人回复的消息id
*/
func (x *Client) QaSatisCreate(satisType SatisType, userID, knowledgeID, msgID string) error {

	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaSatisCreateV2(satisType, userID, knowledgeID, msgID)
	if err != nil {
		return err
	}
	if x.Debug {
		log.Debugf("[QaSatisCreate Response]:%s\n", bytes)
	}

	return nil
}
