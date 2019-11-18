package wulai

import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common"
)

//Client 吾来平台接口统一调用入口
type Client struct {
	HTTPClient *common.Client
	Version    string
	Endpoint   string
	Debug      bool
}

//NewClient 创建 NewWulaiClient
func NewClient(secret, pubkey string) *Client {
	wulai := &Client{}
	//实例化凭证
	credential := common.NewCredential(secret, pubkey)
	//实例化http client
	client := common.NewClient(credential)
	wulai.HTTPClient = client

	//默认版本:V2
	wulai.Version = "v2"
	wulai.Endpoint = "https://openapi.wul.ai"
	return wulai
}

//SetDebug 设置debug
func (x *Client) SetDebug(debug bool) {
	x.Debug = debug
	x.HTTPClient.Debug = debug
}
