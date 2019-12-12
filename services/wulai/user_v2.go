package wulai

import (
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

/*userCreateV2 创建用户(V2)
@userID:用户id [1~128]characters
@nickname:用户昵称 <= 128 characters
@avatarURL:用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
*/
func (x *Client) userCreateV2(userID, nickname, avatarURL string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"nickname": "%s",
		"avatar_url": "%s",
		"user_id": "%s"
	}`, nickname, avatarURL, userID)

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return resp.ResponseBodyBytes, err
}

/*userUpdateV2 更新用户信息(V2)
@userID:用户id [1~128]characters
@nickname:用户昵称 <= 128 characters
@avatarURL:用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
*/
func (x *Client) userUpdateV2(userID, nickname, avatarURL string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"nickname": "%s",
		"avatar_url": "%s",
		"user_id": "%s"
	}`, nickname, avatarURL, userID)

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return resp.ResponseBodyBytes, err
}

/*userGetV2 查询用户信息(V2)
@userID:用户id [1~128]characters
*/
func (x *Client) userGetV2(userID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user/get", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"user_id": "%s"
	}`, userID)

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return resp.ResponseBodyBytes, err
}

/*userAttributeListV2 获取用户属性列表
@filter:是否可以作为属性组属性.如果为空,则不过滤
@page:页码，代表查看第几页的数据，从1开始 >=1
@pageSize:每页的属性组数量 [1~200]
*/
func (x *Client) userAttributeListV2(filter *UserAttributeFilter, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user-attribute/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		@filter
		"page": %v,
		"page_size": %v
	}`, page, pageSize)

	//添加filter
	if filter != nil {
		filter := fmt.Sprintf(`
		"filter": {
			"use_in_user_attribute_group": %v
		},`, filter.HasAttribute)
		input = strings.Replace(input, "@filter", filter, -1)
	} else {
		input = strings.Replace(input, "@filter", "", -1)
	}

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, errors.NewServerError(resp.StatusCode, err.Error(), err)
	}

	return resp.ResponseBodyBytes, nil
}

/*UserAttributeCreate 给用户添加属性值
@userID:用户id [1~128]characters
@attrID:用户属性id
@attrValue:用户属性值 [1 ~ 128] characters
*/
func (x *Client) userAttributeCreateV2(userID, attrID, attrValue string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user/user-attribute/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"user_attribute_user_attribute_value": [
		  {
			"user_attribute": {
			  "id": "%s"
			},
			"user_attribute_value": {
			  "name": "%s"
			}
		  }
		],
		"user_id": "%s"
	}`, attrID, attrValue, userID)

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, errors.NewServerError(resp.StatusCode, err.Error(), err)
	}

	return resp.ResponseBodyBytes, nil
}

/*userAttributePairListV2 查询用户的属性值(V2)
@userID:用户id [1~128]characters
*/
func (x *Client) userAttributePairListV2(userID string) ([]byte, error) {
	//https://openapi.wul.ai/v2/user/user-attribute/pair/list
	url := fmt.Sprintf("%s/%s/user/user-attribute/pair/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"user_id": "%s"
	}`, userID)

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, errors.NewServerError(resp.StatusCode, err.Error(), err)
	}

	return resp.ResponseBodyBytes, nil
}
