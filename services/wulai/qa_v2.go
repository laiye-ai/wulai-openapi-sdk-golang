package wulai

import (
	"fmt"
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
