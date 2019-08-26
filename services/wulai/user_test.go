package wulai

import (
	"os"
	"testing"
	log "github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/zlog"
)

func Test_CreateUser(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewWulaiClient(secret, pubkey)
	wulaiClient.Version = "v2"
	wulaiClient.Debug = true
	//credential.GetRandomString(15)
	resp, err := wulaiClient.UserCreate("test_golang_api", "", "test_golang_api")
	if err != nil {
		t.Fatalf("user Create test reuslt:%s", err.Error())
	}
	log.Infof("%+v\n", resp)
}

func Test_GetUserAttribute(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewWulaiClient(secret, pubkey)
	wulaiClient.Version = "v2"
	resp, err := wulaiClient.UserAttributeGet()
	if err != nil {
		t.Fatalf("user Create test reuslt:%s", err.Error())
	}
	log.Infof("%+v\n", resp)
}
