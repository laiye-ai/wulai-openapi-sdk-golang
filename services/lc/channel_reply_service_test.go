package lc

import (
	"fmt"
	"testing"
)

var client *Client

func init() {
	agentID := "101593"
	channelID := "9fc3a90e-16a7-4801-96f0-e703a92dbfa6"
	endpoint := "https://poc-chatbot.laiye.com"
	client = NewClient(agentID, channelID, endpoint)
	client.Debug = true
}

func Test_ChannelReplyService(t *testing.T) {

	username := "abc"
	question := "打印刻录需要输入账号密码"
	resp, err := client.ChannelReplyService(username, question, UNSPECIFIED)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("%+v\n", resp)
}
