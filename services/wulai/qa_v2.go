package wulai

import (
	"fmt"
	"strings"
)

//qaSimilarQuestionDeleteV2 删除相似问
func (x *Client) qaSimilarQuestionDeleteV2(id string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/similar-question/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": "%s"}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaListUserAttributeGroupItemsV2 查询属性组及属性列表
func (x *Client) qaListUserAttributeGroupItemsV2(page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-items/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"page": %v,
		"page_size": %v
	  }`, page, pageSize)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaCreateUserAttributeGroupV2 创建属性组
func (x *Client) qaCreateUserAttributeGroupV2(groupName, attributeID, attributeName string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-items/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"user_attribute_group_item": {
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
		  "user_attribute_group": {
			"name": "%s"
		  }
		}
	}`, attributeID, attributeName, groupName)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaUpdateUserAttributeGroupV2 更新属性组
func (x *Client) qaUpdateUserAttributeGroupV2(groupID, groupName string, attributes map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-items/update", x.Endpoint, x.Version)

	//拼接属性json字符串
	attrString := ""
	template := `
	{
		"user_attribute": {
		  "id": "@id"
		},
		"user_attribute_value": {
		  "name": "@name"
		}
	}`

	for k, v := range attributes {
		attrString = attrString + strings.Replace(strings.Replace(template, "@id", k, -1), "@name", v, -1)
	}

	//拼接输入参数
	input := fmt.Sprintf(`
	{
		"user_attribute_group_item": {
		  "user_attribute_user_attribute_value": [
			%s
		  ],
		  "user_attribute_group": {
			"id": "%s",
			"name": "%s"
		  }
		}
	}`, attrString, groupID, groupName)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
