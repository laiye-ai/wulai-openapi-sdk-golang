### 使用CommonRequest进行调用


CommonRequest特点是比较灵活,方便开发人员自主实现接口.



示例代码

```go
package main

import (
    "github.com/laiye-ai/wulai-openapi-sdk-golang/services/common"
    "github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

func main() {
    credential := NewCredential("secret", "pubkey")
	client := NewClient(credential)
	client.SetHTTPTimeout(5)
	_, err := client.Request("POST", "http://laiye.com", []byte("input"), 1)
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {
			//TODO
		} else if _, ok := err.(*errors.ClientError); ok {
			e
            //TODO
		} else {
			//TODO
		}
	}
}

```
