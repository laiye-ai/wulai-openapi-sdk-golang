package wulai

import (
	"encoding/json"
	"fmt"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//msgBotResponseV2 获取机器人回复(V2)
func (x *Client) msgBotResponseV2(userID, extra, typeStr string, msgType interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response", x.Endpoint, x.Version)
	var err error
	msgBody, err := json.Marshal(msgType)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	input := fmt.Sprintf(`{
		"extra": %q,
		"msg_body": {"%s": %s},
		"user_id": "%s"
	  }`, extra, typeStr, string(msgBody), userID)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n [Input]: %s\n", url, input)
	}

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgBotResponseQaV2 获取问答机器人回复(V2)
func (x *Client) msgBotResponseQaV2(userID, content, extra string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response/qa", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {
		  "text": {
			"content": %q
		  }
		},
		"user_id": "%s",
		"extra": %q
	  }`, content, userID, extra)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n%s\n", url, input)
	}

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgBotResponseKeywordV2 获取问答机器人回复(V2)
func (x *Client) msgBotResponseKeywordV2(userID, content, extra string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response/keyword", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {
		  "text": {
			"content": %q
		  }
		},
		"user_id": "%s",
		"extra": %q
	  }`, content, userID, extra)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n%s\n", url, input)
	}

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgBotResponseTaskV2 获取任务机器人回复(V2)
func (x *Client) msgBotResponseTaskV2(userID, content, extra string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response/task", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {
		  "text": {
			"content": %q
		  }
		},
		"user_id": "%s",
		"extra": %q
	  }`, content, userID, extra)

	if x.Debug {
		log.Debugf("[Request URL]:%s\n%s\n", url, input)
	}

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}
