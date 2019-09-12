### log 说明


log 使用示例

```go
import (
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

func main() {

    //打印日志
  	log.Info("测试 log")
	log.Debug("测试 log")
	log.Error("测试 log")
    
    //将日志写入文件
    f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.Instance.SetOutput(f)
}
 
```