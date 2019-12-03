package wulai

import (
	"os"
	"testing"
	"time"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

func Test_GetBotResponse(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"} //传入指针
	model, err := wulaiClient.MsgBotResponse("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponse]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponse]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetBotResponse]=>%s\n", err.Error())
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
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v1"

	text := &Text{"您好!"}
	model, err := wulaiClient.MsgBotResponse("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponse]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponse]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetBotResponse]=>%s\n", err.Error())
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
	wulaiClient.SetDebug(true)

	text := &Text{"您好"}
	model, err := wulaiClient.MsgBotResponseQa("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseQAWithText]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseQAWithText]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetBotResponseQAWithText]=>%s\n", err.Error())
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
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	model, err := wulaiClient.MsgBotResponseKeyword("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseKeyword]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseKeyword]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetBotResponseKeyword]=>%s\n", err.Error())
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
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	resp, err := wulaiClient.MsgBotResponseTask("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseTask]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseTask]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetBotResponseTask]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_GetBotResponseWitCustom(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	custom := &Custom{"您好!"}
	resp, err := wulaiClient.MsgBotResponse("xiao_lai", custom, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseWitCustom]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseWitCustom]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetBotResponseWitCustom]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgSend(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	userID := "xiao_lai"
	msgBody := &Text{"您好!"}
	extra := "预留信息"
	quickReply := []string{"今天是个好天气", "今天天气不错", "今天不下雨"}

	sr1 := SimilarResponseParam{
		Source: DEFAULT_ANSWER_SOURCE,
		Detail: &QA{
			KnowledgeID: 100,
		},
	}
	sr2 := SimilarResponseParam{
		Source: DEFAULT_ANSWER_SOURCE,
		Detail: &Task{
			BlockType: BLOCK_TYPE_MESSAGE,
			BlockID:   100,
			BlockName: "测试单元",
		},
	}
	similarResponse := []SimilarResponseParam{sr1, sr2}

	resp, err := wulaiClient.MsgSend(userID, msgBody, extra, quickReply, similarResponse)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgSend]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgSend]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgSend]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgReceive(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	resp, err := wulaiClient.MsgReceive("xiao_lai", text, "third_msg_id_xxxx1", "预留信息")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgReceive]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgReceive]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgReceive]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgSync(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	bot := &QA{}
	text := &Text{"您好!"}
	answerID := 0                        //answer_id 的值从机器人的回复中获取
	msgTS := time.Now().UnixNano() / 1e6 //当前消息时间戳(毫秒级)
	resp, err := wulaiClient.MsgSync("xiao_lai", answerID, msgTS, "预留信息", bot, text)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgSync]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgSync]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgSync]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgSyncWithoutBot(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	answerID := 0                        //answer_id 的值从机器人的回复中获取
	msgTS := time.Now().UnixNano() / 1e6 //当前消息时间戳(毫秒级)
	resp, err := wulaiClient.MsgSync("xiao_lai", answerID, msgTS, "预留信息", nil, text)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgSync]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgSync]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgSync]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgHistory(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	resp, err := wulaiClient.MsgHistory("xiao_lai", "msg_id", BACKWARD, 10)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgHistory]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgHistory]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgHistory]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgSuggestion(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	userID := "xiao_lai" //用户唯一标识 [1-128]characters
	query := "您好"        //用户输入 [1-128]characters
	resp, err := wulaiClient.MsgSuggestion(userID, query)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgSuggestion]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgSuggestion]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgSuggestion]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgLinked(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	userType := PRIVATE //默认 PRIVATE
	hashID := "zV5RxbXHtfVwErjaTZ707GTK7adQ8XrR"
	resp, err := wulaiClient.MsgTriggerLink(userType, hashID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgLinked]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgLinked]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgLinked]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_CheckMsgType(t *testing.T) {
	text := &Text{"您好"}
	_, ok := CheckMsgType(text)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotType(t *testing.T) {
	qa := &QA{}
	_, ok := CheckBotType(qa)
	if !ok {
		t.Fail()
	}
}

func Benchmark_CheckMsgType(t *testing.B) {
	for i := 0; i < t.N; i++ {
		text := &Text{"toenxt"}
		CheckMsgType(text)
	}
}

func Benchmark_CheckBotType(t *testing.B) {
	for i := 0; i < t.N; i++ {
		qa := &QA{}
		CheckBotType(qa)
	}
}

func Benchmark_CheckNilType(t *testing.B) {
	for i := 0; i < t.N; i++ {
		var niltype interface{}
		CheckBotType(niltype)
	}
}
