package wulai

import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common"
)

//WulaiClient 吾来平台接口统一调用入口
type WulaiClient struct {
	HTTPClient *common.Client
	Version    string
	Endpoint   string
	Debug      bool
}

//NewWulaiClient 创建 NewWulaiClient
func NewWulaiClient(secret, pubkey string) *WulaiClient {
	wulai := &WulaiClient{}
	//实例化凭证
	credential := common.NewCredential(secret, pubkey)
	//实例化http client
	client := common.NewClient(credential)
	wulai.HTTPClient = client

	wulai.Version = "v2"
	wulai.Endpoint = "https://openapi.wul.ai"
	return wulai
}

//SetDebug 设置debug
func (x *WulaiClient) SetDebug(debug bool) {
	x.Debug = debug
	x.HTTPClient.Debug = debug
}
