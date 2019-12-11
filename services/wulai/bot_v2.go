package wulai

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*msgBotResponseV2 获取机器人回复(V2)
@userID:用户唯一标识[1-128]characters
@extra:自定义字段 <=1024 characters
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
*/
func (x *Client) msgBotResponseV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*msgBotResponseQaV2 获取问答机器人回复(V2)
@userID:用户唯一标识[1-128]characters
@extra:自定义字段 <=1024 characters
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
*/
func (x *Client) msgBotResponseQaV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/msg/bot-response/qa", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*msgBotResponseKeywordV2 获取关键字机器人回复(V2)
@userID:用户唯一标识[1-128]characters
@extra:自定义字段 <=1024 characters
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
*/
func (x *Client) msgBotResponseKeywordV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/msg/bot-response/keyword", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*msgBotResponseTaskV2 获取任务机器人回复(V2)
@userID:用户唯一标识[1-128]characters
@extra:自定义字段 <=1024 characters
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
*/
func (x *Client) msgBotResponseTaskV2(userID, extra, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/bot-response/task", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*msgHistoryV2 查询历史消息(V2)
@userID:用户唯一标识[1-128]characters
@msgID:从这个msg_id开始查询（结果包含此条消息）；为空时查询最新的消息 <= 18 characters
@direction:翻页方向.BACKWARD: 向旧的消息翻页，查询比传入msg_id更小的消息.FORWARD: 先新的消息翻页，查询比传入msg_id更大的消息
@num:一次获取消息的数目[1-50]
*/
func (x *Client) msgHistoryV2(userID, msgID string, direction direction, num int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/history", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*msgSendV2 给用户发消息(V2)
@userID:用户唯一标识[1-128]characters
@quickReply:快捷回复 <=5 items
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
@extra:自定义字段 <=1024 characters
@similarResponse:推荐知识点 <=5 items
*/
func (x *Client) msgSendV2(userID string, quickReply []string, msgType string, msgBody []byte, extra string, similarResponse []SimilarResponseParam) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/send", x.Endpoint, x.Version)
	srBytes, _ := json.Marshal(similarResponse)
	input := fmt.Sprintf(`
	{
		"similar_response": %s,
		"msg_body": {"%s": %s},
		"quick_reply": [%s],
		"user_id": "%s",
		"extra": "%s"
	}`, srBytes, msgType, msgBody, x.join(quickReply, ","), userID, extra)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

/*msgReceiveV2 接收用户发的消息(V2)
@userID:用户唯一标识[1-128]characters
@thirdMsgID:接入方唯一msg_id，保证1分钟内的幂等性 <= 64 characters
@extra:自定义字段 <=1024 characters
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
*/
func (x *Client) msgReceiveV2(userID, thirdMsgID, extra, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/receive", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
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

/*msgSyncV2 同步发给用户的消息(V2)
@userID:用户唯一标识[1-128]characters
@answerID:答案id
@msgTS:消息毫秒级时间戳
@extra:自定义字段 <=1024 characters
@botType:机器人类型.如果机器人回复兜底内容，则bot为空
@botBody:qa / chitchat / task / keyword
@msgType:消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
@msgBody:消息内容
*/
func (x *Client) msgSyncV2(userID string, answerID, msgTS int64, extra, botType string, botBody []byte, msgType string, msgBody []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/sync", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"user_id": "%s",
		"extra": %q,
		"bot": {},
		"msg_ts": %v,
		"msg_body": {"%s": %s},
		"answer_id": %v
	}`, userID, extra, msgTS, msgType, msgBody, answerID)

	if botType != "" {
		input = strings.Replace(input, `"bot": {}`, fmt.Sprintf(`"bot": {"%s": %s}`, botType, botBody), -1)
	}

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

/*msgSuggestionV2 获取用户输入联想(V2)
@userID:用户唯一标识 [1-128]characters
@query:用户输入 [1-128]characters
*/
func (x *Client) msgSuggestionV2(userID, query string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/user-suggestion/get", x.Endpoint, x.Version)
	input := fmt.Sprintf(`
	{
		"query": "%s",
		"user_id": "%s"
	}`, query, userID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}

/*msgTriggerLinkV2 触发链接消息(V2)
@userType:用户的消息类型 PRIVATE / PUBLIC / PLATFORM / PRIVATE_GROUP / U_BOT_PRIVATE / U_BOT_GROUP / WORK_WECHAT / WORK_WECHAT_GROUP
@hashID:随机字符串 [1-128]characters
*/
func (x *Client) msgTriggerLinkV2(userType MsgUserType, hashID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/msg/t/%s", x.Endpoint, x.Version, hashID)
	input := fmt.Sprintf(`
	{
		"user_type": "%s",
		"hash_id": "%s"
	}`, userType, hashID)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}
	return respBytes.ResponseBodyBytes, nil
}
