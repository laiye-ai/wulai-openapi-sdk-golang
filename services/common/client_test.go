package common

import (
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

func Test_PostWithNetError(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	client := NewClient(credential)
	client.SetHTTPTimeout(5)
	_, err := client.Request("POST", "http://aaabbbxxxyyyyzzzzzz.info", []byte(""), 1)
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.")
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}
func Test_PostWithTransport(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	client := NewClient(credential)

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 3 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second, //限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。
		MaxIdleConns:          30,              //连接池对所有host的最大连接数量，默认无穷大
		MaxConnsPerHost:       30,              //连接池对每个host的最大连接数量。
		IdleConnTimeout:       30 * time.Minute,
		DisableKeepAlives:     false,
	}

	client.SetTransport(transport)
	_, err := client.Request("POST", "http://aaabbbxxxyyyyzzzzzz.info", []byte(""), 1)
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithTransport]=> failed.")
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithTransport]=> failed.")
		}
	}
}

func Test_Request1(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	client := NewClient(credential)
	client.SetHTTPTimeout(3)
	_, err := client.Request("POST", "http://aaabbbxxxyyyyzzzzzz.info", []byte(`{"name":"request"}`), 3)
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.")
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}

func Test_Request2(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	client := NewClient(credential)
	_, err := client.Request("POST", "http://qq.com", []byte(`错误的json`), 3)
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.")
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}

func Test_IsJSONString(t *testing.T) {

	if isJSONString([]byte(`{"name":"right"}`)) {
		//正确 json string
	} else {
		t.Fail()
	}

	if isJSONString([]byte(`{"name":wrong"}`)) {
		t.Fail()
	} else {
		//错误 json string
	}
}

func Benchmark_isJSONString(t *testing.B) {
	for i := 0; i < t.N; i++ {
		if isJSONString([]byte(`{"name":"right"}`)) {
			t.Fail()
		}
	}
}

func Benchmark_Post(t *testing.B) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	client := NewClient(credential)

	for i := 0; i < t.N; i++ {
		_, err := client.Request("POST", "http://baidu.com", []byte(""), 1)
		if err != nil {
			t.Error(err)
		}
	}
}
