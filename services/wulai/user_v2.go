package wulai

import (
	"fmt"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

//参考 https://openapi.wul.ai/docs/latest/saas.openapi.v2/openapi.v2.html#operation/CreateUser
//createV2 创建用户(版本2)
func (x *Client) createV2(nickname, avatarURL, userID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
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

//参考 https://openapi.wul.ai/docs/latest/saas.openapi.v2/openapi.v2.html#operation/ListUserAttribute
//UserAttributeList 获取用户属性列表
func (x *Client) userAttributeListV2(isAttrGroup bool, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user-attribute/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"filter": {
		  "use_in_user_attribute_group": %v
		},
		"page": %v,
		"page_size": %v
	  }`, isAttrGroup, page, pageSize)

	resp, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, errors.NewServerError(resp.StatusCode, err.Error(), err)
	}

	return resp.ResponseBodyBytes, nil
}

//参考 https://openapi.wul.ai/docs/latest/saas.openapi.v2/openapi.v2.html#operation/ListUserAttribute
//userAttributeCreate 给用户添加属性值
func (x *Client) userAttributeCreateV2(userID, attrID, attrValue string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/user/user-attribute/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
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
