package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 获取机器人回复,返回[]byte
****************/

/*MsgBotResponseByte 获取机器人回复(byte)
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponseByte(userID string, msgBody interface{}, extra string) ([]byte, error) {

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "Text/Image/Custom/Video//File/RichText/Voice/Event/ShareLink")
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
		log.Debugf("[MsgBotResponseByte Response]:%s\n", bytes)
	}

	return bytes, nil
}

/*MsgBotResponseQaByte 获取问答机器人回复(byte)
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponseQaByte(userID string, msgBody interface{}, extra string) ([]byte, error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "Text/Image/Custom/Video//File/RichText/Voice/Event/ShareLink")
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
		log.Debugf("[MsgBotResponseQaByte Response]:%s\n", bytes)
	}

	return bytes, nil
}

/*MsgBotResponseKeywordByte 获取关键字机器人回复(byte)
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponseKeywordByte(userID string, msgBody interface{}, extra string) ([]byte, error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "Text/Image/Custom/Video//File/RichText/Voice/Event/ShareLink")
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
		log.Debugf("[MsgBotResponseKeywordByte Response]:%s\n", bytes)
	}

	return bytes, nil
}

/*MsgBotResponseTaskByte 获取任务机器人回复(byte)
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponseTaskByte(userID string, msgBody interface{}, extra string) ([]byte, error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	//检查消息类型是否合法
	msgType, ok := CheckMsgType(msgBody)
	if !ok {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, msgType, "Text/Image/Custom/Video//File/RichText/Voice/Event/ShareLink")
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
		log.Debugf("[MsgBotResponseTaskByte Response]:%s\n", bytes)
	}

	return bytes, nil
}
