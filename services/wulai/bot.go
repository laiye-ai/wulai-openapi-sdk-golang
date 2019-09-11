package wulai

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//Bot 机器人
type Bot struct {
}

//MsgBotResponse 获取机器人回复
func (x *Client) MsgBotResponse(userID string, msgType interface{}, extra string) (model *BotResponseKeyword, err error) {

	//检查消息类型是否合法
	typeStr, ok := checkMsgType(msgType)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, typeStr, "*"+typeStr)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.msgBotResponseV2(userID, extra, typeStr, msgType)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MSGBotResponse Response]:%s\n", bytes)
	}

	model = &BotResponseKeyword{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//MsgBotResponseQa 获取问答机器人回复
func (x *Client) MsgBotResponseQa(userID, content, extra string) (model *BotResponseQa, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.msgBotResponseQaV2(userID, content, extra)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MSGBotResponseQa Response]:%s\n", bytes)
	}

	model = &BotResponseQa{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//MsgBotResponseKeyword 获取关键字机器人回复
func (x *Client) MsgBotResponseKeyword(userID, content, extra string) (model *BotResponseKeyword, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.msgBotResponseKeywordV2(userID, content, extra)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MSGBotResponseKeyword Response]:%s\n", bytes)
	}

	model = &BotResponseKeyword{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//MsgBotResponseTask 获取任务机器人回复
func (x *Client) MsgBotResponseTask(userID, content, extra string) (model *BotResponseTask, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.msgBotResponseTaskV2(userID, content, extra)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MSGBotResponseTask Response]:%s\n", bytes)
	}

	model = &BotResponseTask{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//MsgHistory 查询历史消息
func (x *Client) MsgHistory(userID, msgID string, direction direction, num int) (model *MsgHistory, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//设置默认值为 BACKWARD
	if strings.TrimSpace(string(direction)) == "" {
		direction = BACKWARD
	}

	bytes, err := x.msgHistoryV2(userID, msgID, direction, num)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MsgHistory Response]:%s\n", bytes)
	}

	model = &MsgHistory{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
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
