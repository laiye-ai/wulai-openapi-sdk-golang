## 消息路由
> 消息路由即 Webhook ，该 Web 服务需要开发者自己搭建，并且需要遵从开放平台预先定义好的输入输出。  
  机器人每次响应，吾来会把机器人回复内容、对话解析结果传给消息路由，
  调用方可以按需使或修改消息体内容达到影响机器人回复的目的  
  消息路由可视为：在获取机器人回复前（bot-response），对response的预处理。  
  在同步方式、异步方式中均可使用。


---

## 消息投递
> 调用方开发。如果机器人对接第三方渠道，机器人做出响应后，会调用消息投递接口，将消息投递到第三方渠道。

**注意**：
1) 此接口可保证至少投递一次（即服务质量QOS=At least once，不保证100%消息只投递一次）；
2) 此接口不能保证消息100%按照消息发送的顺序投递

**因此建议**：
1) 接入方根据msg_id做去重
2) 接入方根据msg_ts在客户端做重排序


--- 
example 代码

```go
package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

//Bot bot
type Bot struct {
}

// botMsgDelivery 消息投递 handles
func botMsgDelivery(hub *Hub, w http.ResponseWriter, r *http.Request) {

	respBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Infof("%s\n", err)
		return
	}
    log.Infof("[机器人投递的消息]]=>%s\n", respBytes)
    
	//将机器人回复的消息投递到前端
	hub.botMsgQueue <- respBytes
	w.Write([]byte("ok"))
}

// botMsgRoute 消息路由 handles
func botMsgRoute(hub *Hub, w http.ResponseWriter, r *http.Request) {

	inBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Infof("%s\n", err)
		return
	}
	log.Infof("[收到消息路由传递的消息]]=>%s\n", inBytes)

	//处理收到的消息
	msgBody := &wulai.MessageRoute{}
	if err := json.Unmarshal(inBytes, msgBody); err != nil {
		log.Errorf("%s\n", err)
	}

	respBody := &wulai.MessageRouteResponses{}
	respBody.IsDispatch = false                            //不转人工
	respBody.SuggestedResponse = msgBody.SuggestedResponse //不处理,直接将消息传回

	outBytes, _ := json.Marshal(respBody)

	log.Info("返回处理后的结果给机器人")
	w.Write(outBytes)
}

//ServeBotMsg bot message handle
func ServeBotMsg(hub *Hub, w http.ResponseWriter, r *http.Request) {

	//request log
	log.Infof("[/]=>remote=>%s host=>%s   url=>%s   method=>%s\n", r.RemoteAddr, r.Host, r.URL, r.Method)

	url := strings.ToLower(r.URL.String())
	switch {
	case url == "/bot/message_delivery":
		botMsgDelivery(hub, w, r)
	case url == "/bot/message_route":
		botMsgRoute(hub, w, r)
	default:
		w.Write([]byte("Unknown Pattern"))
	}
}

```




完整代码链接: [Bot异步定制对话](https://github.com/shzy2012/bot_msg_example)



