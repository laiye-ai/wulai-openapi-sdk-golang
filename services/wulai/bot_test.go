package wulai

import (
	"os"
	"strings"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

func Test_GetBotResponse(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	text := &Text{"您好!"}
	model, err := wulaiClient.MsgBotResponse("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponse]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponse]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.SuggestedResponse) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_GetBotResponseV1(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Debug = true
	wulaiClient.Version = "v1"
	text := &Text{"您好!"}
	model, err := wulaiClient.MsgBotResponse("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponse]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponse]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.SuggestedResponse) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_GetBotResponseQAWithText(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	text := &Text{"您好!"}
	model, err := wulaiClient.MsgBotResponseQa("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseQAWithText]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseQAWithText]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.QaSuggestedResponse) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_GetBotResponseKeyword(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	text := &Text{"您好!"}
	model, err := wulaiClient.MsgBotResponseKeyword("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseKeyword]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseKeyword]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.KeywordSuggestedResponse) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_GetBotResponseTask(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	text := &Text{"您好!"}
	model, err := wulaiClient.MsgBotResponseTask("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseTask]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseTask]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.TaskSuggestedResponse) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_GetBotResponseWitCustom(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Debug = true
	custom := &Custom{"您好!"}
	model, err := wulaiClient.MsgBotResponse("xiao_lai", custom, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseWitCustom]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseWitCustom]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.SuggestedResponse) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_MsgReceive(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Debug = true
	text := &Text{"您好!"}
	model, err := wulaiClient.MsgReceive("xiao_lai", text, "third_msg_id_xxxx1", "预留信息")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgReceive]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgReceive]=>%s\n", serErr.Error())
		}

		return
	}

	if strings.TrimSpace(model.MsgID) == "" {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_MsgSync(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Debug = true

	bot := &QA{}
	text := &Text{"您好!"}
	answerID := 0 //answer_id 的值从机器人的回复中获取
	model, err := wulaiClient.MsgSync("xiao_lai", answerID, "", "预留信息", bot, text)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgSync]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgSync]=>%s\n", serErr.Error())
		}

		return
	}
	if strings.TrimSpace(model.MsgID) == "" {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_MsgHistory(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Debug = true

	model, err := wulaiClient.MsgHistory("xiao_lai", "msg_id", BACKWARD, 10)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgHistory]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgHistory]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.Msg) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
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
