package wulai

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	log "github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/zlog"
)

//Bot 机器人
type Bot struct {
}

//BotResponse 获取机器人回复
func (x *WulaiClient) BotResponse(userID string, msgType interface{}, extra string) ([]byte, error) {

	//检查消息类型是否合法
	typeStr, ok := checkMsgType(msgType)
	if !ok {
		return nil, errors.New("msgType 非法")
	}

	if strings.ToUpper(x.Version) == "V1" {

	} else if strings.ToUpper(x.Version) == "V2" {
		return x.getBotResponseV2(userID, extra, typeStr, msgType)
	}

	return x.getBotResponseV2(userID, extra, typeStr, msgType)
}

//GetBotResponseV2 获取机器人回复(V2)
func (x *WulaiClient) getBotResponseV2(userID, extra, typeStr string, msgType interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response", x.Endpoint, x.Version)
	var err error
	msgBody, err := json.Marshal(msgType)
	if err != nil {
		return nil, err
	}

	input := fmt.Sprintf(`{
		"extra": "%s",
		"msg_body": {"%s": %s},
		"user_id": "%s"
	  }`, extra, typeStr, string(msgBody), userID)

	respBytes, err := x.HTTPClient.Post(url, []byte(input))
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

//checkMsgType 检查消息类型
func checkMsgType(msgType interface{}) (string, bool) {

	switch x := msgType.(type) {
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
	default:
		log.Warnf("unexpected type %T:%v \n", x, x)
	}
	return "", false
}
