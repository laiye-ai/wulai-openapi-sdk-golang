package wulai

import (
	"fmt"
	"strings"
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

//UserAttribute 获取用户属性
func (x *Client) UserAttribute() ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user-attribute/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"filter": {
		  "use_in_user_attribute_group": true
		},
		"page": 1,
		"page_size": 1
	  }`)
	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return resp.ResponseBodyBytes, nil
}

//createV2 创建用户(版本2)
func (x *Client) createV2(nickname, avatarURL, userID string) (*User, error) {
	url := fmt.Sprintf("%s/%s/user/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"nickname": "%s",
		"avatar_url": "%s",
		"user_id": "%s"
	  }`, nickname, avatarURL, userID)
	_, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return &User{nickname, avatarURL, userID}, nil
}
