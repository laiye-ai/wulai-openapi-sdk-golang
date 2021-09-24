package wulai

import "fmt"

/****************
- 自然语言处理类
****************/

/*nlpEntitiesExtractV2 实体抽取
@query：待实体抽取query [1-1024]characters
@referenced_system_entity：抽取被引用的预设实体,默认未生效. false: 未生效; true: 已生效
*/
func (x *Client) nlpEntitiesExtractV2(query string, referenced_system_entity bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/nlp/entities/extract", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"query": %q,"referenced_system_entity_only": %v}`, query, referenced_system_entity)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*nlpEntitiesExtractV2 实体抽取
@query：待实体抽取query [1-1024]characters
*/
func (x *Client) nlpTokenizeV2(query string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/nlp/tokenize", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"query": %q}`, query)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
