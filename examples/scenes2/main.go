package main

import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

func main() {

	//实例化客户端
	wulaiClient := wulai.NewClient("secret", "pubkey")
	wulaiClient.Version = "v2"
	log.Infof("[wulaiClient]%+v\n", wulaiClient)

	//1:创建用户
	user, err := wulaiClient.UserCreate("xiao_lai", "nice_name", "")
	if err != nil {
		log.Fatalf("user Create test reuslt:%s", err.Error())
	}

	log.Infof("[创建用户]: %+v\n", user)

	//2:创建属性
	err = wulaiClient.UserAttributeCreate(user.UserID, "体重", "120")
	if err != nil {
		log.Errorf("[创建用户属性失败]: %s\n", err.Error())
	}

	//消息类型[文本消息]
	textMsg := &wulai.Text{
		Content: "您好",
	}

	//3:发起问答
	botResp, err := wulaiClient.MsgBotResponse(user.UserID, textMsg, "预留信息")
	if err != nil {
		log.Fatalf("bot response reuslt:%s", err.Error())
	}
	log.Infof("[机器人回复]%v\n", botResp)
}
