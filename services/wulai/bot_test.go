package wulai

import (
	"fmt"
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

func Test_GetBotResponse(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	text := &Text{"您好!"}
	_, err := wulaiClient.MsgBotResponse("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponse]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetBotResponse]=>%s\n", serErr.Error())
		}
	}
}

func Test_GetBotResponseQAWithText(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	text := &Text{"您好!"}
	_, err := wulaiClient.MsgBotResponseQa("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseQAWithText]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetBotResponseQAWithText]=>%s\n", serErr.Error())
		}
	}
}

func Test_GetBotResponseKeyword(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	text := &Text{"您好!"}
	_, err := wulaiClient.MsgBotResponseKeyword("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseKeyword]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetBotResponseKeyword]=>%s\n", serErr.Error())
		}
	}
}

func Test_GetBotResponseTask(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	text := &Text{"您好!"}
	_, err := wulaiClient.MsgBotResponseTask("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseTask]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetBotResponseTask]=>%s\n", serErr.Error())
		}
	}
}

func Test_GetBotResponseWitCustom(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	custom := &Custom{"您好!"}
	_, err := wulaiClient.MsgBotResponse("xiao_lai", custom, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseWitCustom]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetBotResponseWitCustom]=>%s\n", serErr.Error())
		}
	}
}

func Test_MsgHistory(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.MsgHistory("xiao_lai", "", BACKWARD, 1)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgHistory]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_MsgHistory]=>%s\n", serErr.Error())
		}
	}
}

func Test_MsgReceive(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Debug = true
	text := &Text{"您好!"}
	_, err := wulaiClient.MsgReceive("xiao_lai", text, "", "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgReceive]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_MsgReceive]=>%s\n", serErr.Error())
		}
	}
}

func Test_MsgSync(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Debug = true
	text := &Text{"您好!"}
	_, err := wulaiClient.MsgSync("xiao_lai", text, "", "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgSync]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_MsgSync]=>%s\n", serErr.Error())
		}
	}
}

func Test_CheckMsgType(t *testing.T) {
	text := &Text{"您好"}
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
