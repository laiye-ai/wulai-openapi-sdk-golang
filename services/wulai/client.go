package wulai

import (
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//Client 吾来平台接口统一调用入口
type Client struct {
	HTTPClient *common.Client
	Version    string
	Endpoint   string
	Debug      bool
}

/*NewClient 创建 NewWulaiClient
@secret 和 @pubkey: 从吾来平台获取. 每个开放平台渠道有一组 pubkey 和 Secret
*/
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

//SetVersion 设置version
func (x *Client) SetVersion(version string) {
	x.Version = strings.ToLower(version)
}

/****************
- tools
****************/

//join 返回包含引号("")的字符串
func (x *Client) join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return "\"" + a[0] + "\""
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i]) + 2
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString("\"" + a[0] + "\"")
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString("\"" + s + "\"")
	}

	log.Infoln(b.String())
	return b.String()
}
