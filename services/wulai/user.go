package wulai

import (
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

//User 用户
type User struct {
	nickname  string
	AvatarURL string
	UserID    string
}

//UserCreate 创建用户
func (x *Client) UserCreate(userID, nickname, avatarURL string) (*User, error) {
	if strings.ToUpper(x.Version) == "V1" {
		return x.createV1(nickname, avatarURL, userID)
	} else if strings.ToUpper(x.Version) == "V2" {
		return x.createV2(nickname, avatarURL, userID)
	}

	//set default
	return x.createV2(nickname, avatarURL, userID)
}

//UserAttributeList 获取用户属性列表
func (x *Client) UserAttributeList(isAttrGroup bool, page, pageSize int) ([]byte, error) {
	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	} else if strings.ToUpper(x.Version) == "V2" {
		return x.userAttributeListV2(isAttrGroup, page, pageSize)
	}

	//set default
	return x.userAttributeListV2(isAttrGroup, page, pageSize)
}

//UserAttributeCreate 给用户添加属性值
func (x *Client) UserAttributeCreate(userID, attrName, attrValue string) ([]byte, error) {
	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	} else if strings.ToUpper(x.Version) == "V2" {
		return x.userAttributeCreateV2(userID, attrName, attrValue)
	}

	//set default
	return x.userAttributeCreateV2(userID, attrName, attrValue)
}
