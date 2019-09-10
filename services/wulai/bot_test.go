package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

func Test_GetBotResponse(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	text := &Text{"您好!"}
	_, err := wulaiClient.MSGBotResponse("xiao_lai", text, "")
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponse]=> %s\n", err.Error())
		}

		t.Log(err.Error())
	}
}

func Test_GetBotResponseQAWithText(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.MSGBotResponseQa("xiao_lai", "您好!", "")
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseQAWithText]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_GetBotResponseKeyword(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.MSGBotResponseKeyword("xiao_lai", "您好!", "")
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseKeyword]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_GetBotResponseTask(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.MSGBotResponseTask("xiao_lai", "您好!", "")
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseTask]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_GetBotResponseWitCustom(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	custom := &Custom{"您好!"}
	_, err := wulaiClient.MSGBotResponse("xiao_lai", custom, "")
	if err != nil {
		if _, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_GetBotResponseWitCustom]=> failed.")
		}

		t.Log(err.Error())
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
