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

//MsgBotResponse 获取机器人回复
func (x *Client) MsgBotResponse(userID string, msgBody interface{}, extra string) (model *BotResponse, err error) {

	//检查消息类型是否合法
	msgType, ok := checkMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	var bytes []byte
	var errRes error

	if strings.ToUpper(x.Version) == "V1" {
		bytes, errRes = x.msgBotResponseV1(userID, msgType, msgBytes)
	} else {
		bytes, errRes = x.msgBotResponseV2(userID, extra, msgType, msgBytes)
	}

	if errRes != nil {
		return nil, errRes
	}

	if x.Debug {
		log.Debugf("[MSGBotResponse Response]:%s\n", bytes)
	}

	model = &BotResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//MsgBotResponseQa 获取问答机器人回复
func (x *Client) MsgBotResponseQa(userID string, msgBody interface{}, extra string) (model *BotResponseQa, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := checkMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.msgBotResponseQaV2(userID, extra, msgType, msgBytes)
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
func (x *Client) MsgBotResponseKeyword(userID string, msgBody interface{}, extra string) (model *BotResponseKeyword, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := checkMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.msgBotResponseKeywordV2(userID, extra, msgType, msgBytes)
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
func (x *Client) MsgBotResponseTask(userID string, msgBody interface{}, extra string) (model *BotResponseTask, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := checkMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.msgBotResponseTaskV2(userID, extra, msgType, msgBytes)
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

//MsgReceive 接收用户发的消息
func (x *Client) MsgReceive(userID string, msgBody interface{}, thirdMsgID, extra string) (model *MsgReceive, err error) {

	if strings.ToUpper(x.Version) == "V1" {
		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := checkMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.msgReceiveV2(userID, thirdMsgID, extra, msgType, msgBytes)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MsgReceive Response]:%s\n", bytes)
	}

	model = &MsgReceive{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

//MsgSync 同步发给用户的消息
func (x *Client) MsgSync(userID string, answerID int, msgTS, extra string, botBody, msgBody interface{}) (model *MsgSync, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := checkMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "*"+msgType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	//检查Bot类型是否合法
	botType, ok := checkBotType(botBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, botType, "*"+botType)
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, nil)
	}

	//msg bytes
	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	//bot bytes
	botBytes, err := json.Marshal(botBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	bytes, err := x.msgSyncV2(userID, answerID, msgTS, extra, botType, botBytes, msgType, msgBytes)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MsgSync Response]:%s\n", bytes)
	}

	model = &MsgSync{}
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

//checkBotType 检查Bot类型
func checkBotType(botType interface{}) (string, bool) {

	switch botType.(type) {
	case *QA:
		return "qa", true
	case *Chitchat:
		return "chitchat", true
	case *Task:
		return "task", true
	case *Keyword:
		return "keyword", true
	}
	return reflect.TypeOf(botType).String(), false
}
