### zlog 是基于go log 的简单封装

例子
```go

import (
	log "wulai-openapi-sdk-golang/services/common/zlog"
)

func main() {
	log.Info("test info")
	log.Infof("%s", "test infof")
}
```

运行性能测试
```go
go test -benchmem  -bench=.  -run=^$
```
运行结果
```go
  100000             13042 ns/op             200 B/op          4 allocs/op
  100000             12683 ns/op             200 B/op          4 allocs/op
  100000             12803 ns/op             200 B/op          4 allocs/op
PASS
ok      wulai-openapi-sdk-golang/services/common/zlog   10.441s
```