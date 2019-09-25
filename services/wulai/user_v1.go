package wulai

import (
	"fmt"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//CreateV1 创建用户
func (x *Client) createV1(nickname, avatarURL, userID string) ([]byte, error) {
	url := x.Endpoint + "/v1/user/create"
	input := fmt.Sprintf(`{
		"username": "%s",
		"nickname": "%s",
		"imgurl": "%s"
	  }`, userID, nickname, avatarURL)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n%s\n", url, input)
	}

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return resp.ResponseBodyBytes, err
}

//UserUpdate 更新用户信息
func (x *Client) userUpdateV1(userID, nickname, avatarURL string) ([]byte, error) {
	url := x.Endpoint + "/v1/user/update"
	input := fmt.Sprintf(`{
		"username": "%s",
		"nickname": "%s",
		"imgurl": "%s"
	  }`, userID, nickname, avatarURL)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n%s\n", url, input)
	}

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return resp.ResponseBodyBytes, err
}

//GroupMembers 获取群成员列表信息
func (x *Client) groupMembers(pageIndex, pageSize int, userID string) ([]byte, error) {
	url := x.Endpoint + "/v1/user/group/members"
	input := fmt.Sprintf(`{
		"page_index": %v,
		"page_size": %v,
		"user_id": "%s"
	 }`, pageIndex, pageSize, userID)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n%s\n", url, input)
	}

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//UserInfo 查询用户信息
func (x *Client) userInfoV1(userID string) ([]byte, error) {
	url := x.Endpoint + "/v1/user/info/get"
	input := fmt.Sprintf(`{
		"user_id": "%s"
		}`, userID)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n%s\n", url, input)
	}

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
