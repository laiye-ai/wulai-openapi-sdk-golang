package wulai

import (
	"encoding/json"
	"fmt"
)

/****************
- 词库管理类
****************/

//dicEntityListV2 查询全部实体概要
func (x *Client) dicEntityListV2(page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/list", x.Endpoint, x.Version)
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

/*dicEntityGetV2 查询一个实体详情
 *@id：实体ID
 */
func (x *Client) dicEntityGetV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/get", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"id": %v
	}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicEntityDeleteV2 删除实体
 *@id：实体ID
 */
func (x *Client) dicEntityDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"id": %v
	}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

//dicEntityEnumCreateV2 创建枚举实体
//name:实体名称
func (x *Client) dicEntityEnumCreateV2(name string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/enumeration/create", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"enum_entity": {
		  "name": %q
		}
	}`, name)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicEntityEnumValueCreateV2 创建枚举实体值
 *@entityID：实体ID
 *@standardValue：标准值(归一化值)
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) dicEntityEnumValueCreateV2(entityID int, standardValue string, synonyms []string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/enumeration/value/create", x.Endpoint, x.Version)

	tmpSynonym, err := json.Marshal(synonyms)
	if err != nil {
		return nil, err
	}

	input := fmt.Sprintf(`
	{
		"entity_id": %v,
		"value": {
		  "synonyms": %s,
		  "standard_value": %q
		}
	}`, entityID, tmpSynonym, standardValue)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicEntityEnumValueDeleteV2 删除枚举实体值
 *@entityID：实体ID
 *@standardValue：标准值(归一化值)
 *@synonyms：标准值的相似说法 []string
 */
 func (x *Client) dicEntityEnumValueDeleteV2(entityID int, standardValue string, synonyms []string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/enumeration/value/delete", x.Endpoint, x.Version)

	tmpSynonym, err := json.Marshal(synonyms)
	if err != nil {
		return nil, err
	}

	input := fmt.Sprintf(`
	{
		"entity_id": %v,
		"value": {
		  "synonyms": %s,
		  "standard_value": %q
		}
	}`, entityID, tmpSynonym, standardValue)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicEntityIntentCreateV2 创建意图实体
 *@name：实体名称 len:[1~200]
 *@standardValue：标准值(归一化值) len:[1~200]
 */
func (x *Client) dicEntityIntentCreateV2(name, standardValue string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/intent/create", x.Endpoint, x.Version)

	input := fmt.Sprintf(`
	{
		"intent_entity": {
		  "standard_value": "%s",
		  "name": "%s"
		}
	}`, standardValue, name)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicEntityIntentValueCreateV2 创建意图实体值相似说法
 *@entityID：实体ID
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) dicEntityIntentValueCreateV2(entityID int, synonyms []string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/intent/value/create", x.Endpoint, x.Version)

	tmpSynonym, err := json.Marshal(synonyms)
	if err != nil {
		return nil, err
	}

	input := fmt.Sprintf(`
	{
		"entity_id": %v,
		"synonyms": %s
	}`, entityID, tmpSynonym)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicEntityIntentValueDeleteV2 删除意图实体值相似说法
 *@entityID：实体ID
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) dicEntityIntentValueDeleteV2(entityID int, synonyms []string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/entity/intent/value/delete", x.Endpoint, x.Version)

	tmpSynonym, err := json.Marshal(synonyms)
	if err != nil {
		return nil, err
	}

	input := fmt.Sprintf(`
	{
		"entity_id": %v,
		"synonyms": %s
	}`, entityID, tmpSynonym)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/****************
- 专有词汇
****************/

//dicTermListV2 查询专有词汇列表
func (x *Client) dicTermListV2(page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/term/list", x.Endpoint, x.Version)
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

/*dicTermCreateV2 创建专有词汇
 *@name：专有词汇的名称 [1~64]characters
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) dicTermCreateV2(name string, synonyms []string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/term/create", x.Endpoint, x.Version)
	tmpSynonym, err := json.Marshal(synonyms)
	if err != nil {
		return nil, err
	}

	input := fmt.Sprintf(`
	{
		"term_item": {
		  "term": {
			"name": "%s"
		  },
		  "synonyms": %s
		}
	}`, name, tmpSynonym)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicTermUpdateV2 更新专有词汇
 *@id：专有词汇id
 *@name：专有词汇的名称 [1~64]characters
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) dicTermUpdateV2(id int, name string, synonyms []string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/term/update", x.Endpoint, x.Version)
	tmpSynonym, err := json.Marshal(synonyms)
	if err != nil {
		return nil, err
	}

	input := fmt.Sprintf(`
	{
		"term_item": {
		  "term": {
			"id": %v,
			"name": "%s"
		  },
		  "synonyms": %s
		}
	}`, id, name, tmpSynonym)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*dicTermDeleteV2 删除专有词汇
 *@id：专有词汇id
 */
func (x *Client) dicTermDeleteV2(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/dictionary/term/delete", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"id": %v
	}`, id)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
