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

//QaKnowledgeTagList 获取知识点分类列表
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

//QaqaKnowledgeCreate 创建知识点
func (x *Client) QaqaKnowledgeCreate(knowledgeTagID int, standardQuestion string, status, respondAll, maintained bool) (model *QaKnowledgeTagResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaKnowledgeCreateV2(knowledgeTagID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaqaKnowledgeCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaKnowledgeTagResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//QaqaKnowledgeUpdate 更新知识点
func (x *Client) QaqaKnowledgeUpdate(knowledgeID int, standardQuestion string, status, respondAll, maintained bool) (model *QaKnowledgeResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaKnowledgeUpdateV2(knowledgeID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[QaqaKnowledgeUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaKnowledgeResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//QaKnowledgeItemList 获取知识点列表
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
		log.Debugf("[QaqaKnowledgeCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaKnowledgeItemListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 相识问
****************/

//QaSimilarQuestionList 获取相似问列表
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

//QaSimilarQuestionCreate 创建相似问
func (x *Client) QaSimilarQuestionCreate(knowledgeID, question, id string) (model *QaSimilarQuestionResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaSimilarQuestionCreateV2(knowledgeID, question, id)
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

//QaSimilarQuestionDelete 删除相似问
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

//QaListUserAttributeGroupItems 查询属性组及属性列表
func (x *Client) QaListUserAttributeGroupItems(page, pageSize int) (model *QaUserAttributeGroupItemList, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaListUserAttributeGroupItemsV2(page, pageSize)
	if err != nil {
		return nil, err
	}
	if x.Debug {
		log.Debugf("[QaListUserAttributeGroupItems Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaUserAttributeGroupItemList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//QaCreateUserAttributeGroup 创建属性组
func (x *Client) QaCreateUserAttributeGroup(groupName, attributeID, attributeName string) (model *QaUserAttributeGroupResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaCreateUserAttributeGroupV2(groupName, attributeID, attributeName)
	if err != nil {
		return nil, err
	}
	if x.Debug {
		log.Debugf("[QaCreateUserAttributeGroup Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaUserAttributeGroupResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//QaUpdateUserAttributeGroup 更新属性组
func (x *Client) QaUpdateUserAttributeGroup(groupID, groupName string, attributes map[string]string) (model *QaUserAttributeGroupResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//发起调用
	bytes, err := x.qaUpdateUserAttributeGroupV2(groupID, groupName, attributes)
	if err != nil {
		return nil, err
	}
	if x.Debug {
		log.Debugf("[QaUpdateUserAttributeGroup Response]:%s\n", bytes)
	}

	//返回结果
	model = &QaUserAttributeGroupResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}
