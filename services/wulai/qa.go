package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

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
