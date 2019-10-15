package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//User 用户
type User struct {
	UserID    string `json:"user_id"`
	AvatarURL string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
}

//UserCreate 创建用户
func (x *Client) UserCreate(userID, nickname, avatarURL string) (model *User, err error) {
	var bytes []byte
	if strings.ToUpper(x.Version) == "V1" {
		bytes, err = x.createV1(nickname, avatarURL, userID)
	} else {
		bytes, err = x.createV2(nickname, avatarURL, userID)
	}

	if x.Debug {
		log.Debugf("[UserCreate Response]:%s\n", bytes)
	}

	if err != nil {
		return nil, err
	}

	return &User{userID, avatarURL, nickname}, nil
}

//UserUpdate 更新用户信息
func (x *Client) UserUpdate(userID, nickname, avatarURL string) (model *User, err error) {
	var bytes []byte
	if strings.ToUpper(x.Version) == "V1" {
		bytes, err = x.userUpdateV1(nickname, avatarURL, userID)
	} else {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V2", "V1")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	if x.Debug {
		log.Debugf("[UserUpdate Response]:%s\n", bytes)
	}

	if err != nil {
		return nil, err
	}

	return &User{nickname, avatarURL, userID}, nil
}

//UserInfo 查询用户信息
func (x *Client) UserInfo(userID string) (model *User, err error) {
	var bytes []byte
	if strings.ToUpper(x.Version) == "V1" {
		bytes, err = x.userInfoV1(userID)
	} else {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V2", "V1")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	if x.Debug {
		log.Debugf("[UserInfo Response]:%s\n", bytes)
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

//UserAttributeList 获取用户属性列表
func (x *Client) UserAttributeList(isAttrGroup bool, page, pageSize int) (model *UserAttributeList, err error) {
	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.userAttributeListV2(isAttrGroup, page, pageSize)
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

//UserAttributeCreate 给用户添加属性值
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
