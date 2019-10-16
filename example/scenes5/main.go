package main

import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

func main() {

	//实例化客户端
	wulaiClient := wulai.NewClient("secret", "pubkey")
	wulaiClient.Version = "v2" //默认v2
	wulaiClient.Debug = true
	log.Infof("[wulaiClient]%+v\n", wulaiClient)

	//1:创建用户
	user, err := wulaiClient.UserCreate("xiao_lai", "nice_name", "")
	if err != nil {
		log.Fatalf("user Create test reuslt:%s", err.Error())
	}
	log.Infof("[创建用户]: %+v\n", user)

	//bot 类型
	botType := &wulai.QA{}

	//消息类型[文本消息]
	textMsg := &wulai.Text{
		Content: "您好",
	}

	//2:同步消息给吾来平台
	answerID := 0 //answer_id 的值从机器人的回复中获取
	msgSync, _ := wulaiClient.MsgSync("xiao_lai", answerID, "0", "预留信息", botType, textMsg)

	//3:获取历史消息
	botResp, err := wulaiClient.MsgHistory(user.UserID, msgSync.MsgID, wulai.BACKWARD, 10)
	if err != nil {
		log.Fatalf("bot response reuslt:%s", err.Error())
	}
	log.Infof("[机器人回复]%v\n", botResp)
}
