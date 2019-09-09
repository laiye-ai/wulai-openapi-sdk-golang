package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

func Test_GetBotResponseWithText(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	text := &Text{"您好!"}
	botResp, err := wulaiClient.MSGBotResponse("userID", text, "")
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if err, ok := err.(*errors.ClientError); ok {

			t.Errorf("[Test_GetBotResponseWithText]=> %s.", err.Message())
		} else {
			t.Error("[Test_GetBotResponseWithText]=> failed.")
		}
	}
	_, _ = botResp.ToString()
}

func Test_GetBotResponseQAWithText(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	botResp, err := wulaiClient.MSGBotResponseQa("userID", "您好!", "")
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if err, ok := err.(*errors.ClientError); ok {

			t.Errorf("[Test_GetBotResponseWithText]=> %s.", err.Message())
		} else {
			t.Error("[Test_GetBotResponseWithText]=> failed.")
		}
	}
	_, _ = botResp.ToString()
}

func Test_GetBotResponseWitCustom(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)

	custom := &Custom{"您好!"}
	botResp, err := wulaiClient.MSGBotResponse("userID", custom, "")
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if _, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_GetBotResponseWitCustom]=> failed.")
		} else {
			t.Error("[Test_GetBotResponseWitCustom]=> failed.")
		}
	}
	_, _ = botResp.ToString()
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
