package wulai

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

//Bot 机器人
type Bot struct {
}

//MSGBotResponse 获取机器人回复
func (x *Client) MSGBotResponse(userID string, msgType interface{}, extra string) ([]byte, error) {

	//检查消息类型是否合法
	typeStr, ok := checkMsgType(msgType)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, typeStr, "*"+typeStr)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	if strings.ToUpper(x.Version) == "V1" {

	} else if strings.ToUpper(x.Version) == "V2" {
		return x.msgBotResponseV2(userID, extra, typeStr, msgType)
	}

	return x.msgBotResponseV2(userID, extra, typeStr, msgType)
}

//msgBotResponseV2 获取机器人回复(V2)
func (x *Client) msgBotResponseV2(userID, extra, typeStr string, msgType interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response", x.Endpoint, x.Version)
	var err error
	msgBody, err := json.Marshal(msgType)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	input := fmt.Sprintf(`{
		"extra": "%s",
		"msg_body": {"%s": %s},
		"user_id": "%s"
	  }`, extra, typeStr, string(msgBody), userID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//MSGBotResponseQa 获取问答机器人回复
func (x *Client) MSGBotResponseQa(content, userID, extra string) ([]byte, error) {

	if strings.ToUpper(x.Version) == "V1" {

	} else if strings.ToUpper(x.Version) == "V2" {
		return x.msgBotResponseQaV2(content, userID, extra)
	}
	return x.msgBotResponseQaV2(content, userID, extra)
}

//msgBotResponseQaV2 获取问答机器人回复(V2)
func (x *Client) msgBotResponseQaV2(content, userID, extra string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response/qa", x.Endpoint, x.Version)
	input := fmt.Sprintf(`{
		"msg_body": {
		  "text": {
			"content": "%s"
		  }
		},
		"user_id": "%s",
		"extra": "%s"
	  }`, content, userID, extra)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//checkMsgType 检查消息类型
func checkMsgType(msgType interface{}) (string, bool) {

	switch msgType.(type) {
	case *Text:
		return "text", true
	case *Image:
		return "image", true
	case *Custom:
		return "custom", true
	case *Video:
		return "video", true
	case *File:
		return "file", true
	case *Voice:
		return "voice", true
	case *Event:
		return "event", true
	case *ShareLink:
		return "share_link", true
	}
	return reflect.TypeOf(msgType).String(), false
}
