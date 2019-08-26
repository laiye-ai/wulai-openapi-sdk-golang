package wulai

import (
	"fmt"
)

//CreateV1 创建用户
func (x *WulaiClient) createV1(nickname, avatarURL, userID string) (*User, error) {
	url := x.Endpoint + "/v1/user/create"
	input := fmt.Sprintf(`{
		"username": "%s",
		"nickname": "%s",
		"imgurl": "%s"
	  }`, userID, nickname, avatarURL)
	_, err := x.HTTPClient.Post(url, []byte(input))
	if err != nil {
		return nil, err
	}
	return &User{nickname, avatarURL, userID}, nil
}

//UserUpdate 更新用户信息
func (x *WulaiClient) UserUpdate(nickname, avatarURL, userID string) error {
	url := x.Endpoint + "/v1/user/update"
	input := fmt.Sprintf(`{
		"username": "%s",
		"nickname": "%s",
		"imgurl": "%s"
	  }`, userID, nickname, avatarURL)
	_, err := x.HTTPClient.Post(url, []byte(input))
	if err != nil {
		return err
	}
	return nil
}

//GroupMembersGet 获取群成员列表信息
func (x *WulaiClient) GroupMembersGet(pageIndex, pageSize int, userID string) ([]byte, error) {
	url := x.Endpoint + "/v1/user/group/members"
	input := fmt.Sprintf(`{
		"page_index": %v,
		"page_size": %v,
		"user_id": "%s"
	 }`, pageIndex, pageSize, userID)

	respBytes, err := x.HTTPClient.Post(url, []byte(input))
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//UserGet 查询用户信息
func (x *WulaiClient) UserGet(userID string) ([]byte, error) {
	url := x.Endpoint + "/v1/user/info/get"
	input := fmt.Sprintf(`{
		"user_id": "%s"
		}`, userID)
	respBytes, err := x.HTTPClient.Post(url, []byte(input))
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
