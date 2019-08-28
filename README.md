
<p align="center">
	<a href="https://www.laiye.com"><img src="https://www.laiye.com/static/official-website/logo.png"></a>
</p>

<h1 align="center">Wulai Openapi SDK for Go</h1>

<p align="center">

[![build status][travis-image]][travis-url]   [![codecov][cov-image]][cov-url]

[travis-image]: https://travis-ci.org/laiye-ai/wulai-openapi-sdk-golang.svg?branch=master

[travis-url]: https://travis-ci.org/laiye-ai/wulai-openapi-sdk-golang

[cov-image]: https://codecov.io/gh/laiye-ai/wulai-openapi-sdk-golang/branch/master/graph/badge.svg

[cov-url]: https://codecov.io/gh/laiye-ai/wulai-openapi-sdk-golang

</p>

欢迎使用 Wulai Openapi SDK for Go。

## 环境要求
- 您的系统需要达到 [环境要求][Requirements], 例如，安装了不低于 1.10.x 版本的 Go 环境。

## 安装
使用 `go get` 下载安装 SDK

```sh
$ go get -u github.com/laiye-ai/wulai-openapi-sdk-golang
```

如果您使用了 glide 管理依赖，您也可以使用 glide 来安装 SDK

```sh
$ glide get github.com/laiye-ai/wulai-openapi-sdk-golang
```

## 快速使用
在您开始之前，您需要注册帐户并获取您的[凭证](https://openapi.wul.ai/docs/latest/saas.openapi.v2/openapi.v2.html#section/%E9%89%B4%E6%9D%83%E8%AE%A4%E8%AF%81)。

### 创建HTTP客户端
```go
package main

import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common"
)

func main() {
	credential := common.NewCredential("secret", "pubkey")
	client := common.NewClient(credential)
	client.Post("http://openapi.wul.ai", []byte("input"))
}
```

### 创建用户,对机器人发起问答
```go
package main

import (
	log "github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/zlog"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

func main() {

	//实例化客户端
	wulaiClient := wulai.NewClient("secret", "pubkey")
	wulaiClient.Version = "v2"
	log.Infof("[wulaiClient]%+v\n", wulaiClient)

	user, err := wulaiClient.UserCreate("test_golang_api", "", "test_golang_api")
	if err != nil {
		log.Fatalf("user Create test reuslt:%s", err.Error())
	}

	log.Infof("[创建用户]: %+v\n", user)

	//消息类型[文本消息]
	textMsg := &wulai.Text{
		Content: "您好",
	}
	//发起问答
	botResp, err := wulaiClient.BotResponse(user.UserID, textMsg, "预留信息")
	if err != nil {
		log.Fatalf("user Create test reuslt:%s", err.Error())
	}
	log.Infof("[机器人回复]%s\n", botResp)
}

```


### 协议说明
```text
网络协议: HTTP

请求方式: POST

请求、响应数据格式: JSON
```