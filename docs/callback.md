### 回调类接口实现


#### 消息路由
> 消息路由即 Webhook ，该 Web 服务需要开发者自己搭建，并且需要遵从开放平台预先定义好的输入输出。  
  机器人每次响应，吾来会把机器人回复内容、对话解析结果传给消息路由，
  调用方可以按需使或修改消息体内容达到影响机器人回复的目的  
  消息路由可视为：在获取机器人回复前（bot-response），对response的预处理。  
  在同步方式、异步方式中均可使用。

go code example

```go


```


#### 消息投递
> 调用方开发。如果机器人对接第三方渠道，机器人做出响应后，会调用消息投递接口，将消息投递到第三方渠道。

**注意**：
1) 此接口可保证至少投递一次（即服务质量QOS=At least once，不保证100%消息只投递一次）；
2) 此接口不能保证消息100%按照消息发送的顺序投递

**因此建议**：
1) 接入方根据msg_id做去重
2) 接入方根据msg_ts在客户端做重排序

go code example

```go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CallbackHandler 接收吾来平台回调, 返回信息给客户端
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	//request log
	log.Printf("[/]=>remote=>%s host=>%s   url=>%s   method=>%s\n", r.RemoteAddr, r.Host, r.URL, r.Method)

	respBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	log.Printf("[message delivery log]=>%s\n", respBytes)
	w.Write([]byte("success"))
}

func main() {

	var port int
	flag.IntVar(&port, "p", 8000, "端口号")
	flag.Parse()

	//设置接受机器人答案的函数
	http.Handle("/callback", http.HandlerFunc(CallbackHandler))

	//设置http服务
	ip := fmt.Sprintf(":%v", port)
	log.Printf("listen on:%s\n", ip)
	http.ListenAndServe(ip, nil)
}

```