package wulai

import (
	"fmt"
	"strings"
)

/****************
- 知识点
****************/

/*qaKnowledgeTagListV2 知识点分类列表
@parentTagID：父节点分类id，如果值为0，代表获取根节点下的知识点分类
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) qaKnowledgeTagListV2(parentTagID, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge-tags/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*qaKnowledgeCreateV2 创建知识点
@knowledgeTagID：知识点分类id >=1
@standardQuestion：知识点标题 <= 100 characters
@status：知识点状态. [true:生效 false:未生效]
@respondAll：发送全部回复. [true:发送全部回复 false:随机一条发送]
@maintained：该属性是否被用于定义属性组。True：该属性是定义属性组所使用的属性之一。False：该属性不被用于定义属性组。[true:是 false:否]
*/
func (x *Client) qaKnowledgeCreateV2(knowledgeTagID, standardQuestion string, status, respondAll, maintained bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge-tag-knowledge/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"knowledge_tag_knowledge": {
		  "knowledge": {
			"status": %v,
			"standard_question": %q,
			"respond_all": %v,
			"maintained_by_user_attribute_group": %v
		  },
		  "knowledge_tag_id": "%s"
		}
	}`, status, standardQuestion, respondAll, maintained, knowledgeTagID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*qaKnowledgeItemListV2 知识点列表
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) qaKnowledgeItemListV2(page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge-items/list", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"page": %v,
		"page_size": %v
	}`, page, pageSize)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*qaKnowledgeUpdateV2 更新知识点
@knowledgeID：知识点id
@standardQuestion：知识点标题 <= 100 characters
@status：知识点状态. [true:生效 false:未生效]
@respondAll：发送全部回复. [true:发送全部回复 false:随机一条发送]
@maintained：该属性是否被用于定义属性组。True：该属性是定义属性组所使用的属性之一。False：该属性不被用于定义属性组。[true:是 false:否]
*/
func (x *Client) qaKnowledgeUpdateV2(knowledgeID int64, standardQuestion string, status, respondAll, maintained bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/knowledge/update", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*qaSimilarQuestionListV2 相似问列表
@knowledgeID：知识点id（为空：查询所有）
@similarQuestionID：相似问id（为空：查询所有）
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
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

/*qaSimilarQuestionCreateV2 创建相似问
@knowledgeID：知识点id
@question：相似问 <=100 characters
*/
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

/*qaSimilarQuestionUpdateV2 更新相似问
@knowledgeID：知识点id
@question：相似问 <=100 characters
@id：相识问id
*/
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

/*qaSimilarQuestionDeleteV2 删除相似问
@id：相似问id
*/
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

/*qaUserAttributeGroupItemListV2 查询属性组及属性列表
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
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

/*qaUserAttributeGroupCreateV2 创建属性组
@groupName:属性组名称 [1~128]characters
@attributeID:用户属性ID
@attributeValue:用户属性值 [1~128]characters
*/
func (x *Client) qaUserAttributeGroupItemCreateV2(groupName, attributeID, attributeValue string) ([]byte, error) {
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
	}`, attributeID, attributeValue, groupName)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*qaUserAttributeGroupUpdateV2 更新属性组
@groupID：属性组id
@groupName：属性组名称 [1~128] characters
@attributes：用户属性组及属性
*/
func (x *Client) qaUserAttributeGroupItemUpdateV2(groupID, groupName string, attributes map[string]string) ([]byte, error) {
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

/*qaUserAttributeGroupAnswerListV2 查询属性组回复列表(v2)
@knowledgeID：知识点id，如不为空，返回所有知识点
@groupID：属性组id，如不为空，返回所有属性组
@page：页码，代表查看第几页的数据，从1开始
@pageSize：每页的触发器数量[1-200]
*/
func (x *Client) qaUserAttributeGroupAnswerListV2(knowledgeID, groupID string, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-answers/list", x.Endpoint, x.Version)
	//@filter 可选
	input := fmt.Sprintf(`
	 {
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

/*qaUserAttributeGroupAnswerCreateV2 创建属性组回复(v2)
@knowledgeID：知识点id
@groupID：属性组ID
@msgType：消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）
@msgBody：消息
*/
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

/*qaUserAttributeGroupAnswerUpdateV2 更新属性组回复(v2)
@knowledgeID：知识点id
@groupID：属性组ID
@answerID：回复id
@msgType：消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）
@msgBody：消息
*/
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

/*qaUserAttributeGroupAnswerUpdateV2 删除属性组回复(v2)
@answerID：属性组回复ID
*/
func (x *Client) qaUserAttributeGroupAnswerDeleteV2(answerID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/user-attribute-group-answer/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"id": "%s"}`, answerID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

/****************
- QA 统计
****************/

/*qaSatisCreateV2 添加用户满意度评价
@satisType:满意度枚举类型
@userID:用户id
@knowledgeID:知识点id
@msgID:机器人回复的消息id
*/
func (x *Client) qaSatisCreateV2(satisType SatisType, userID, knowledgeID, msgID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/qa/satisfaction/create", x.Endpoint, x.Version)

	input := fmt.Sprintf(`
	{
		"satisfaction": "%s",
		"msg_id": "%s",
		"user_id": "%s",
		"bot_id": {
		  "knowledge_id": "%s"
		}
	}`, satisType, msgID, userID, knowledgeID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
