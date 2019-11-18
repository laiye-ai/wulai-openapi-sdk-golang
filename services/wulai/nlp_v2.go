package wulai

import "fmt"

/****************
- 自然语言处理类
****************/

/*nlpEntitiesExtractV2 实体抽取
@query：待实体抽取query [1-1024]characters
*/
func (x *Client) nlpEntitiesExtractV2(query string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/nlp/entities/extract", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{"query": %q}`, query)

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
