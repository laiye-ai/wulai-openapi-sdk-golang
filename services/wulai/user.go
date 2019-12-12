package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/*UserCreate 创建用户
@userID:用户id [1~128]characters
@nickname:用户昵称 <= 128 characters
@avatarURL:用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
*/
func (x *Client) UserCreate(userID, nickname, avatarURL string) (model *User, err error) {
	var bytes []byte
	if strings.ToLower(x.Version) == "v1" {
		bytes, err = x.userCreateV1(userID, nickname, avatarURL)
	} else {
		bytes, err = x.userCreateV2(userID, nickname, avatarURL)
	}

	if x.Debug {
		log.Debugf("[UserCreate Response]:%s\n", bytes)
	}

	if err != nil {
		return nil, err
	}

	return &User{userID, avatarURL, nickname}, nil
}

/*UserUpdate 更新用户信息
@userID:用户id [1~128]characters
@nickname:用户昵称 <= 128 characters
@avatarURL:用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
*/
func (x *Client) UserUpdate(userID, nickname, avatarURL string) (model *User, err error) {
	var bytes []byte
	if strings.ToUpper(x.Version) == "V1" {
		bytes, err = x.userUpdateV1(userID, nickname, avatarURL)
	} else {
		bytes, err = x.userUpdateV2(userID, nickname, avatarURL)
	}

	if x.Debug {
		log.Debugf("[UserUpdate Response]:%s\n", bytes)
	}

	if err != nil {
		return nil, err
	}

	return &User{userID, nickname, avatarURL}, nil
}

/*UserGet 查询用户信息
@userID:用户id [1~128]characters
*/
func (x *Client) UserGet(userID string) (model *User, err error) {
	var bytes []byte
	if strings.ToUpper(x.Version) == "V1" {
		bytes, err = x.userGetV1(userID)
	} else {
		bytes, err = x.userGetV2(userID)
	}

	if x.Debug {
		log.Debugf("[UserGet Response]:%s\n", bytes)
	}

	if err != nil {
		return nil, err
	}

	model = &User{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	model.UserID = userID
	return model, nil
}

/*UserAttributeList 获取用户属性列表
@UserAttributeFilter:用户属性过滤条件。如果填写，代表需要过滤；反之不过滤
@page:页码，代表查看第几页的数据，从1开始 >=1
@pageSize:每页的属性组数量 [1~200]
*/
func (x *Client) UserAttributeList(filter *UserAttributeFilter, page, pageSize int) (model *UserAttributeList, err error) {
	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.userAttributeListV2(filter, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[UserAttributeList Response]:%s\n", bytes)
	}

	model = &UserAttributeList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*UserAttributeCreate 给用户添加属性值
@userID:用户id [1~128]characters
@attrID:用户属性id  >=1
@attrValue:用户属性值 [1 ~ 128] characters
*/
func (x *Client) UserAttributeCreate(userID, attrID, attrValue string) error {

	var bytes []byte

	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.userAttributeCreateV2(userID, attrID, attrValue)

	if x.Debug {
		log.Debugf("[UserAttributeCreate Response]:%s\n", bytes)
	}

	return err
}

/*UserAttributePairList 查询用户的属性值
@userID:用户id [1~128]characters
*/
func (x *Client) UserAttributePairList(userID string) (model *UserAttributePairListResponse, err error) {
	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.userAttributePairListV2(userID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[UserAttributePairList Response]:%s\n", bytes)
	}

	model = &UserAttributePairListResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}
