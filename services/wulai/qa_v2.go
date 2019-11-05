package wulai

import (
	"fmt"
	"strings"
)

/****************
- 知识点
****************/

//qaKnowledgeTagListV2 知识点分类列表
func (x *Client) qaKnowledgeTagListV2(parentTagID, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge-tags/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"parent_k_tag_id": %v,
		"page": %v,
		"page_size": %v
	  }`, parentTagID, page, pageSize)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaKnowledgeCreateV2 创建知识点
func (x *Client) qaKnowledgeCreateV2(knowledgeTagID int, standardQuestion string, status, respondAll, maintained bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge-tag-knowledge/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"knowledge_tag_knowledge": {
		  "knowledge": {
			"status": %v,
			"standard_question": %q,
			"respond_all": %v,
			"maintained_by_user_attribute_group": %v
		  },
		  "knowledge_tag_id": %v
		}
	  }`, status, standardQuestion, respondAll, maintained, knowledgeTagID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaKnowledgeItemListV2 知识点列表
func (x *Client) qaKnowledgeItemListV2(page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge-items/list", x.Endpoint, x.Version)
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

//qaKnowledgeUpdateV2 更新知识点
func (x *Client) qaKnowledgeUpdateV2(knowledgeID int, standardQuestion string, status, respondAll, maintained bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"knowledge": {
		  "status": %v,
		  "standard_question": %q,
		  "respond_all": %v,
		  "id": %v,
		  "maintained_by_user_attribute_group": %v
		}
	  }`, status, standardQuestion, respondAll, knowledgeID, maintained)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 相识问
****************/

//qaSimilarQuestionListV2 相似问列表
func (x *Client) qaSimilarQuestionListV2(knowledgeID, similarQuestionID string, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/similar-question/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"knowledge_id": "%s",
		"page": %v,
		"page_size": %v,
		"similar_question_id": "%s"
	  }`, knowledgeID, page, pageSize, similarQuestionID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaSimilarQuestionCreateV2 创建相似问
func (x *Client) qaSimilarQuestionCreateV2(knowledgeID, question string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/similar-question/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"similar_question": {
		  "knowledge_id": "%s",
		  "question": %q
		}
	  }`, knowledgeID, question)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaSimilarQuestionUpdateV2 更新相似问
func (x *Client) qaSimilarQuestionUpdateV2(knowledgeID, question, id string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/similar-question/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"similar_question": {
		  "knowledge_id": "%s",
		  "question": %q,
		  "id": "%s"
		}
	  }`, knowledgeID, question, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

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

/****************
- 用户属性组
****************/

//qaUserAttributeGroupItemListV2 查询属性组及属性列表
func (x *Client) qaUserAttributeGroupItemListV2(page, pageSize int) ([]byte, error) {
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

//qaUserAttributeGroupCreateV2 创建属性组
func (x *Client) qaUserAttributeGroupCreateV2(groupName, attributeID, attributeName string) ([]byte, error) {
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

//qaUserAttributeGroupUpdateV2 更新属性组
func (x *Client) qaUserAttributeGroupUpdateV2(groupID, groupName string, attributes map[string]string) ([]byte, error) {
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

/****************
- 属性组回复
****************/

//qaUserAttributeGroupAnswerListV2 查询属性组回复列表
func (x *Client) qaUserAttributeGroupAnswerListV2(knowledgeID, groupID string, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-answers/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"filter": {
		  "knowledge_id": "%s",
		  "user_attribute_group_id": "%s"
		},
		"page": %v,
		"page_size": %v
	  }`, knowledgeID, groupID, page, pageSize)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//qaUserAttributeGroupAnswerCreateV2 创建属性组回复(v2)
func (x *Client) qaUserAttributeGroupAnswerCreateV2(knowledgeID, groupID, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-answer/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	  {
		"user_attribute_group_answer": {
		  "answer": {
			"knowledge_id": "%s",
			"msg_body": {"%s": %s}
		  },
		  "user_attribute_group_id": "%s"
		}
	  }`, knowledgeID, msgType, msgBody, groupID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//qaUserAttributeGroupAnswerUpdateV2 更新属性组回复(v2)
func (x *Client) qaUserAttributeGroupAnswerUpdateV2(knowledgeID, groupID, answerID, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-answer/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	  {
		"user_attribute_group_answer": {
		  "answer": {
			"knowledge_id": "%s",
			"msg_body": {"%s": %s},
			"id": "%s"
		  },
		  "user_attribute_group_id": "%s"
		}
	  }`, knowledgeID, msgType, msgBody, answerID, groupID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//qaUserAttributeGroupAnswerUpdateV2 删除属性组回复(v2)
func (x *Client) qaUserAttributeGroupAnswerDeleteV2(answerID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-answer/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": "%s"}`, answerID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}
