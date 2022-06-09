package lc

import (
	"encoding/json"
	"fmt"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/*ChannelReplyService 获取机器人回复
@agentID:针对指定机器人，机器人id可在机器人设置中查看
@channelID:渠道ID
@username:用户唯一标识[1-128]characters,如果不存在,则由系统自动创建
@question:消息内容
@environment:UNSPECIFIED,SKETCH,PRODUCT
参考: https://poc-chatbot.laiye.com/chatbot-openapi/swagger-ui/#/ChannelReplyService/ChannelReplyService_GetReply
*/
func (x *Client) ChannelReplyService(agentID, channelID, username, question string, environment Environment) (*ChannelReplyServiceResponse, error) {
	if environment == "" {
		environment = UNSPECIFIED
	}

	model := &ChannelReplyServiceResponse{}
	url := fmt.Sprintf("%s/chatbot/v1alpha1/agents/%s/channels/%s/getReply", x.Endpoint, agentID, channelID)
	input := fmt.Sprintf(`{
		"username": "%s",
		"query": {
		  "text": {
			"content": "%s"
		  }
		},
		"environment": "%s"
	  }`, username, question, environment)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return model, err
	}

	if x.Debug {
		log.Debugf("[ChannelReplyService]=>%s\n", respBytes.ResponseBodyBytes)
	}

	if err = json.Unmarshal(respBytes.ResponseBodyBytes, model); err != nil {
		return model, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

type Environment string

const (
	UNSPECIFIED Environment = "UNSPECIFIED"
	SKETCH      Environment = "SKETCH"
	PRODUCT     Environment = "PRODUCT"
)

type ChannelReplyServiceResponse struct {
	MsgID              string          `json:"msgId"`
	Responses          []Responses     `json:"responses"`
	QuickReply         QuickReply      `json:"quickReply"`
	Slots              []Slot          `json:"slots"`
	ResponseConfidence int64           `json:"responseConfidence"`
	Intent             Intent          `json:"intent"`
	IntentRanking      []IntentRanking `json:"intentRanking"`
	Entities           []Entity        `json:"entities"`
	Metadata           Metadata        `json:"metadata"`
	DebugInfo          DebugInfo       `json:"debugInfo"`
	ResponseTime       string          `json:"responseTime"`
}
type Text struct {
	Content string `json:"content"`
}
type Image struct {
	ResourceURL string `json:"resourceUrl"`
}
type Voice struct {
	ResourceURL string `json:"resourceUrl"`
	Recognition string `json:"recognition"`
}
type Video struct {
	ResourceURL string `json:"resourceUrl"`
}
type File struct {
	ResourceURL string `json:"resourceUrl"`
}
type RichText struct {
	Content string `json:"content"`
}

type Custom struct {
	Content interface{} `json:"content"`
}
type MsgBody struct {
	Text     Text     `json:"text,omitempty"`
	Image    Image    `json:"image,omitempty"`
	Voice    Voice    `json:"voice,omitempty"`
	Video    Video    `json:"video,omitempty"`
	File     File     `json:"file,omitempty"`
	RichText RichText `json:"richText,omitempty"`
	Custom   Custom   `json:"custom,omitempty"`
}
type SugIntents struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}
type Responses struct {
	MsgBody    MsgBody      `json:"msgBody"`
	SugIntents []SugIntents `json:"sugIntents"`
}
type QuickReply struct {
	Prompt       string   `json:"prompt"`
	QuickReplies []string `json:"quickReplies"`
}
type Slot struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Origin string `json:"origin"`
}
type Intent struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}
type Skill struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}
type Trigger struct {
	Type    string `json:"type"`
	Example string `json:"example"`
}
type IntentRanking struct {
	Intent  Intent  `json:"intent"`
	Skill   Skill   `json:"skill"`
	Trigger Trigger `json:"trigger"`
}
type Entity struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	Origin     string `json:"origin"`
	Start      int64  `json:"start"`
	End        int64  `json:"end"`
	Tag        string `json:"tag"`
	Role       string `json:"role"`
	Group      int64  `json:"group"`
	Confidence int64  `json:"confidence"`
}
type Metadata struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
type DebugInfo struct {
}
