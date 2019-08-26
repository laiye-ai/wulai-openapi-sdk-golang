package common

import (
	"os"
	"testing"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/zlog"
)

func Test_Post(t *testing.T) {
	secret, pubkey:= os.Getenv("secret"),os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	client := NewClient(credential)

	_, err := client.Post("http://google.com.com", []byte(""))
	if err != nil {
		zlog.Error(err)
	}

	//zlog.Infof("%+v", resp)
}

func Benchmark_Post(t *testing.B) {
	secret, pubkey:= os.Getenv("secret"),os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	client := NewClient(credential)

	for i := 0; i < t.N; i++ {
		resp, err := client.Post("http://baidu.com", []byte(""))
		if err != nil {
			zlog.Error(err)
		}
		zlog.Infof("%+v", resp)
	}
}
