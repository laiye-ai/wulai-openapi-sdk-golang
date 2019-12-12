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
	resp, err := wulaiClient.MsgBotResponse("xiao_lai", text, "")
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

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_GetBotResponseV1(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v1"

	text := &Text{"您好!"}
	resp, err := wulaiClient.MsgBotResponse("xiao_lai", text, "")
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

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_GetBotResponseQA(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := Text{"您好"}
	resp, err := wulaiClient.MsgBotResponseQa("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetBotResponseQA]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetBotResponseQA]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetBotResponseQA]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_MsgBotResponseKeyword(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	resp, err := wulaiClient.MsgBotResponseKeyword("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgBotResponseKeyword]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgBotResponseKeyword]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgBotResponseKeyword]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_BotResponseTask(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	resp, err := wulaiClient.MsgBotResponseTask("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_BotResponseTask]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_BotResponseTask]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_BotResponseTask]=>%s\n", err.Error())
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

	userID := "xiao_lai"    //用户唯一标识[1-128]characters
	msgBody := &Text{"您好!"} //消息体格式，任意选择一种消息类型（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）填充
	extra := "预留信息"         //自定义字段 <=1024 characters

	quickReply := []string{"今天是个好天气", "今天天气不错", "今天不下雨"} //快捷回复 <=5 items

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
	similarResponse := []SimilarResponseParam{sr1, sr2} //推荐知识点 <=5 items

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
	answerID := int64(0)                 //answer_id 的值从机器人的回复中获取
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
	answerID := int64(0)                 //answer_id 的值从机器人的回复中获取
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

func Test_MsgTriggerLink(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	userType := PRIVATE //默认 PRIVATE
	hashID := "zV5RxbXHtfVwErjaTZ707GTK7adQ8XrR"
	resp, err := wulaiClient.MsgTriggerLink(userType, hashID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgTriggerLink]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgTriggerLink]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgTriggerLink]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_CheckMsgTypeText(t *testing.T) {
	tp := &Text{"您好"}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeImage(t *testing.T) {
	tp := &Image{"url"}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeCustom(t *testing.T) {
	tp := &Custom{"url"}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeVideo(t *testing.T) {
	tp := &Video{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeRichText(t *testing.T) {
	tp := &RichText{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeFile(t *testing.T) {
	tp := &File{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeVoice(t *testing.T) {
	tp := &Voice{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeEvent(t *testing.T) {
	tp := &Event{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeShareLink(t *testing.T) {
	tp := &ShareLink{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeText2(t *testing.T) {
	tp := Text{"您好"}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeImage2(t *testing.T) {
	tp := Image{"url"}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeCustom2(t *testing.T) {
	tp := Custom{"url"}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeVideo2(t *testing.T) {
	tp := Video{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeRichText2(t *testing.T) {
	tp := RichText{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeFile2(t *testing.T) {
	tp := File{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeVoice2(t *testing.T) {
	tp := Voice{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeEvent2(t *testing.T) {
	tp := Event{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckMsgTypeShareLink2(t *testing.T) {
	tp := ShareLink{}
	_, ok := CheckMsgType(tp)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeQA(t *testing.T) {
	qa := &QA{}
	_, ok := CheckBotType(qa)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeQA2(t *testing.T) {
	q := QA{}
	_, ok := CheckBotType(q)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeTask(t *testing.T) {
	q := &Task{}
	_, ok := CheckBotType(q)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeTask2(t *testing.T) {
	q := Task{}
	_, ok := CheckBotType(q)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeChitchat(t *testing.T) {
	q := &Chitchat{}
	_, ok := CheckBotType(q)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeChitchat2(t *testing.T) {
	q := Chitchat{}
	_, ok := CheckBotType(q)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeKeyword(t *testing.T) {
	q := &Keyword{}
	_, ok := CheckBotType(q)
	if !ok {
		t.Fail()
	}
}

func Test_CheckBotTypeKeyword2(t *testing.T) {
	q := Keyword{}
	_, ok := CheckBotType(q)
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

func Test_MsgBotResponseByte(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"} //传入指针
	resp, err := wulaiClient.MsgBotResponseByte("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgBotResponseByte]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgBotResponseByte]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgBotResponseByte]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%s\n", resp)
}

func Test_MsgBotResponseQaByte(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := Text{"您好"}
	resp, err := wulaiClient.MsgBotResponseQaByte("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgBotResponseQaByte]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgBotResponseQaByte]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgBotResponseQaByte]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%s\n", resp)
}

func Test_MsgBotResponseKeywordByte(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	resp, err := wulaiClient.MsgBotResponseKeywordByte("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgBotResponseKeywordByte]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgBotResponseKeywordByte]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgBotResponseKeywordByte]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%s\n", resp)
}

func Test_MsgBotResponseTaskByte(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	text := &Text{"您好!"}
	resp, err := wulaiClient.MsgBotResponseTaskByte("xiao_lai", text, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_MsgBotResponseTaskByte]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_MsgBotResponseTaskByte]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_MsgBotResponseTaskByte]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%s\n", resp)
}
