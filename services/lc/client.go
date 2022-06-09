package lc

import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common"
)

//Client Laiye Chatbot(LC)平台接口统一调用入口
type Client struct {
	HTTPClient *common.Client
	Endpoint   string
	Debug      bool
}

/*NewClient 创建 NewLClient
@secret 和 @pubkey: 从吾来平台获取. 每个开放平台渠道有一组 pubkey 和 Secret
*/
func NewClient(secret, pubkey, endpoint string) *Client {
	lc := &Client{}
	//实例化凭证
	credential := common.NewCredential(secret, pubkey)
	//实例化http client
	client := common.NewClient(credential)
	lc.HTTPClient = client

	lc.Endpoint = endpoint
	if endpoint == "" {
		lc.Endpoint = "https://newtestchatbot.wul.ai"
	}

	return lc
}

//SetDebug 设置debug
func (x *Client) SetDebug(debug bool) {
	x.Debug = debug
	x.HTTPClient.Debug = debug
}
