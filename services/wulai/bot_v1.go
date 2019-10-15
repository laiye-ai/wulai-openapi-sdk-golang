package wulai

import (
	"fmt"
)

//msgBotResponseV1 获取机器人回复(V1)
func (x *Client) msgBotResponseV1(userID, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {"%s": %s},
		"user_id": "%s"
	  }`, msgType, msgBody, userID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}
