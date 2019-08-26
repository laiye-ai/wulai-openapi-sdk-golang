package wulai

import (
	"os"
	"testing"
)

func Test_GetBotResponseWithText(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewWulaiClient(secret, pubkey)
	wulaiClient.Version = "v2"
	wulaiClient.Debug = true

	text := &Text{"您好!"}
	_, err := wulaiClient.BotResponse("test_golang_api", text, "")
	if err != nil {
		t.Errorf("user Create test reuslt:%s", err.Error())
	}
}

func Test_GetBotResponseWitCustom(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewWulaiClient(secret, pubkey)
	wulaiClient.Version = "v2"
	wulaiClient.Debug = true

	custom := &Custom{"您好!"}
	_, err := wulaiClient.BotResponse("test_golang_api", custom, "")
	if err != nil {
		t.Errorf("user Create test reuslt:%s", err.Error())
	}
}

func Test_CheckMsgType(t *testing.T) {
	text := &Text{"ceshi"}
	_, ok := checkMsgType(text)
	if !ok {
		t.Fail()
	}
}

func Benchmark_CheckMsgType(t *testing.B) {
	for i := 0; i < t.N; i++ {
		text := &Text{"toenxt"}
		checkMsgType(text)
	}
}
