package wulai

import (
	"fmt"
)

//msgBotResponseV2 获取机器人回复(V2)
func (x *Client) msgBotResponseV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"extra": %q,
		"msg_body": {"%s": %s},
		"user_id": "%s"
	  }`, extra, msgType, msgBody, userID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgBotResponseQaV2 获取问答机器人回复(V2)
func (x *Client) msgBotResponseQaV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/msg/bot-response/qa", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {"%s": %s},
		"user_id": "%s",
		"extra": %q
	  }`, msgType, msgBody, userID, extra)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgBotResponseKeywordV2 获取问答机器人回复(V2)
func (x *Client) msgBotResponseKeywordV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/msg/bot-response/keyword", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {"%s": %s},
		"user_id": "%s",
		"extra": %q
	  }`, msgType, msgBody, userID, extra)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgBotResponseTaskV2 获取任务机器人回复(V2)
func (x *Client) msgBotResponseTaskV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response/task", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {"%s": %s},
		"user_id": "%s",
		"extra": %q
	  }`, msgType, msgBody, userID, extra)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgHistoryV2 查询历史消息(V2)
func (x *Client) msgHistoryV2(userID, msgID string, direction direction, num int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/history", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"direction": "%s",
		"msg_id": "%s",
		"user_id": "%s",
		"num": %v
	  }`, direction, msgID, userID, num)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgReceiveV2 接收用户发的消息(V2)
func (x *Client) msgReceiveV2(userID, thirdMsgID, extra, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/receive", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {"%s": %s},
		"third_msg_id": "%s",
		"user_id": "%s",
		"extra": %q
	  }`, msgType, msgBody, thirdMsgID, userID, extra)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//msgSyncV2 同步发给用户的消息(V2)
func (x *Client) msgSyncV2(userID string, answerID int, msgTS, extra, botType string, botBody []byte, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/sync", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"user_id": "%s",
		"extra": "%s",
		"bot": {"%s": %s},
		"msg_ts": "%s",
		"msg_body": {"%s": %s},
		"answer_id": %v
	  }`, userID, extra, botType, botBody, msgTS, msgType, msgBody, answerID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}
