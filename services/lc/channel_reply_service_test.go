package lc

import (
	"fmt"
	"testing"
)

var client *Client

func init() {
	secret := "xxx"
	pubkey := "xxx"
	endpoint := "https://lc.laiye.com"
	client = NewClient(secret, pubkey, endpoint)
	client.Debug = true
}

func Test_ChannelReplyService(t *testing.T) {

	agentID := "154"
	channelID := "6754ad84-28b5-42da-97bf-7df295aee17a"
	username := "abc"
	question := "知识库怎么创建"
	resp, err := client.ChannelReplyService(agentID, channelID, username, question, UNSPECIFIED)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("%+v\n", resp)
}
