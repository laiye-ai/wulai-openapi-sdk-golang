package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//Bot 机器人

/*MsgBotResponse 获取机器人回复
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponse(userID string, msgBody interface{}, extra string) (model *BotResponse, err error) {

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
		log.Debugf("[MSGBotResponse Response]:%s\n", bytes)
	}

	model = &BotResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*MsgBotResponseQa 获取问答机器人回复
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponseQa(userID string, msgBody interface{}, extra string) (model *BotResponseQa, err error) {

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
		log.Debugf("[MSGBotResponseQa Response]:%s\n", bytes)
	}

	model = &BotResponseQa{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*MsgBotResponseKeyword 获取关键字机器人回复
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponseKeyword(userID string, msgBody interface{}, extra string) (model *BotResponseKeyword, err error) {

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
		log.Debugf("[MSGBotResponseKeyword Response]:%s\n", bytes)
	}

	model = &BotResponseKeyword{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*MsgBotResponseTask 获取任务机器人回复
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgBotResponseTask(userID string, msgBody interface{}, extra string) (model *BotResponseTask, err error) {

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
		log.Debugf("[MSGBotResponseTask Response]:%s\n", bytes)
	}

	model = &BotResponseTask{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*MsgHistory 查询历史消息
@userID:用户唯一标识[1-128]characters
@msgID:从这个msg_id开始查询（结果包含此条消息）；为空时查询最新的消息 <= 18 characters
@direction:翻页方向.BACKWARD: 向旧的消息翻页，查询比传入msg_id更小的消息.FORWARD: 先新的消息翻页，查询比传入msg_id更大的消息
@num:一次获取消息的数目[1-50]
*/
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

/*MsgSend 给用户发消息
@userID:用户唯一标识[1-128]characters
@quickReply:快捷回复 <=5 items
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
@extra:自定义字段 <=1024 characters
@similarResponse:推荐知识点 <=5 items
*/
func (x *Client) MsgSend(userID string, msgBody interface{}, extra string, quickReply []string, similarResponse []SimilarResponseParam) (model *MsgSend, err error) {

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

	//check quickReply
	if len(quickReply) > 5 {
		errorMsg := fmt.Sprintf(errors.LimitExceededErrorMessage, "quickReply")
		return nil, errors.NewClientError(errors.LimitExceededErrorCode, errorMsg, err)
	}

	//check similarResponse
	if len(similarResponse) > 5 {
		errorMsg := fmt.Sprintf(errors.LimitExceededErrorMessage, "similarResponse")
		return nil, errors.NewClientError(errors.LimitExceededErrorCode, errorMsg, err)
	}

	bytes, err := x.msgSendV2(userID, quickReply, msgType, msgBytes, extra, similarResponse)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MsgSend Response]:%s\n", bytes)
	}

	model = &MsgSend{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*MsgReceive 接收用户发的消息
@userID:用户唯一标识[1-128]characters
@msgBody:消息体格式(指针类型)，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@thirdMsgID:接入方唯一msg_id，保证1分钟内的幂等性 <= 64 characters
@extra:自定义字段 <=1024 characters
*/
func (x *Client) MsgReceive(userID string, msgBody interface{}, thirdMsgID, extra string) (model *MsgReceive, err error) {

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

/*MsgSync 同步发给用户的消息
@userID:用户唯一标识[1-128]characters
@answerID:答案id
@msgTS:消息毫秒级时间戳 比如:当前消息时间戳(毫秒级) = time.Now().UnixNano() / 1e6
@extra:自定义字段 <=1024 characters
@botBody:机器人类型qa / chitchat / task / keyword.如果机器人回复兜底内容，则bot为空
@msgBody:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
*/
func (x *Client) MsgSync(userID string, answerID, msgTS int64, extra string, botBody, msgBody interface{}) (model *MsgSync, err error) {

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

	//msg bytes
	msgBytes, err := json.Marshal(msgBody)
	if err != nil {
		return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	//检查Bot类型是否合法
	botBytes := make([]byte, 0)
	botType, ok := CheckBotType(botBody)
	if ok {
		//传递bot
		botBytes, err = json.Marshal(botBody)
		if err != nil {
			return nil, errors.NewClientError(errors.JsonMarshalErrorCode, errors.JsonMarshalErrorMessage, err)
		}
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

/*MsgSuggestion 获取用户输入联想
@userID:用户唯一标识 [1-128]characters
@query:用户输入 [1-128]characters
*/
func (x *Client) MsgSuggestion(userID, query string) (model *MsgSuggestionResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.msgSuggestionV2(userID, query)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MsgSuggestion Response]:%s\n", bytes)
	}

	model = &MsgSuggestionResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*MsgTriggerLink 触发链接消息
@userType:用户的消息类型 PRIVATE / PUBLIC / PLATFORM / PRIVATE_GROUP / U_BOT_PRIVATE / U_BOT_GROUP / WORK_WECHAT / WORK_WECHAT_GROUP
@hashID:随机字符串 [1-128]characters
*/
func (x *Client) MsgTriggerLink(userType MsgUserType, hashID string) (model *MsgTriggerLink, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.msgTriggerLinkV2(userType, hashID)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[MsgTriggerLink Response]:%s\n", bytes)
	}

	model = &MsgTriggerLink{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/****************
- 类型检查方法
****************/

//CheckMsgType 检查消息类型
func CheckMsgType(msgType interface{}) (string, bool) {

	switch msgType.(type) {
	case *Text:
		return "text", true
	case Text:
		return "text", true
	case *RichText:
		return "rich_text", true
	case RichText:
		return "rich_text", true
	case *Image:
		return "image", true
	case Image:
		return "image", true
	case *Custom:
		return "custom", true
	case Custom:
		return "custom", true
	case *Video:
		return "video", true
	case Video:
		return "video", true
	case *File:
		return "file", true
	case File:
		return "file", true
	case *Voice:
		return "voice", true
	case Voice:
		return "voice", true
	case *Event:
		return "event", true
	case Event:
		return "event", true
	case *ShareLink:
		return "share_link", true
	case ShareLink:
		return "share_link", true
	}
	return "", false
}

//CheckBotType 检查Bot类型
func CheckBotType(botType interface{}) (string, bool) {

	switch botType.(type) {
	case *QA:
		return "qa", true
	case QA:
		return "qa", true
	case *Chitchat:
		return "chitchat", true
	case Chitchat:
		return "chitchat", true
	case *Task:
		return "task", true
	case Task:
		return "task", true
	case *Keyword:
		return "keyword", true
	case Keyword:
		return "keyword", true
	}
	return "", false
}
