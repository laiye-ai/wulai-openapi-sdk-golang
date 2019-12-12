package main

import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

func main() {

	//实例化客户端
	wulaiClient := wulai.NewClient("secret", "pubkey")
	wulaiClient.Version = "v2" //默认v2
	wulaiClient.SetDebug(true)

	//1:创建用户
	user, err := wulaiClient.UserCreate("xiao_lai", "nice_name", "")
	if err != nil {
		log.Fatalf("user Create test reuslt:%s", err.Error())
	}

	log.Infof("[创建用户]: %+v\n", user)

	//2:创建属性
	err = wulaiClient.UserAttributeCreate(user.UserID, "101521", "120") //101521(属性ID):在平台查看
	if err != nil {
		log.Errorf("[更新用户属性失败]: %s\n", err.Error())
	}

	//消息类型[文本消息]
	textMsg := &wulai.Text{
		Content: "您好",
	}

	//3:将消息传给机器人
	botResp, err := wulaiClient.MsgReceive(user.UserID, textMsg, "third_msg_id_xxxx1", "预留信息")
	if err != nil {
		log.Fatalf("bot response reuslt:%s", err.Error())
	}
	log.Infof("[机器人回复]%v\n", botResp)
}
