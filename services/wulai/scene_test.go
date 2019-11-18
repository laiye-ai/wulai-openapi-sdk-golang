package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 场景类
****************/
func Test_SceneList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	resp, err := wulaiClient.SceneList()
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	name := "测试Scene名称"                           //场景名称.[1-200]characters
	description := "这是SDK 创建的"                    //场景描述 <= 600 characters
	intentSwitchMode := INTENT_SWITCH_MODE_SWITCH //意图切换模式
	smartSlotFillingThreshold := 0                //智能填槽阈值 <= 1

	resp, err := wulaiClient.SceneCreate(name, description, intentSwitchMode, smartSlotFillingThreshold)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	scens := &Scene{
		ID:                        11424,
		Name:                      "测试Scene名称-2",
		Description:               "更新",
		IntentSwitchMode:          INTENT_SWITCH_MODE_SWITCH,
		SmartSlotFillingThreshold: 0,
	}

	resp, err := wulaiClient.SceneUpdate(scens)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 11424 //场景ID
	err := wulaiClient.SceneDelete(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 意图
****************/
func Test_SceneIntentList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	sceneID := 11431 //意图所属的场景ID>=1
	resp, err := wulaiClient.SceneIntentList(sceneID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

/****************
- 触发器
****************/

func Test_SceneIntentTriggerList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 100 //意图ID >=1
	page := 1       //页码，代表查看第几页的数据，从1开始
	pageSize := 10  //每页的触发器数量[1-200]

	resp, err := wulaiClient.SceneIntentTriggerList(intentID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentTriggerList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentTriggerList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneIntentTriggerCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 11431                               //触发器对应的意图ID >=1
	text := "吃饭"                                    //触发文本[1-200]characters
	triggerType := TRIGGER_TYPE_EXACT_MATCH_KEYWORD //触发器模式

	resp, err := wulaiClient.SceneIntentTriggerCreate(intentID, text, triggerType)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentTriggerCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentTriggerCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

/****************
- 词槽类
****************/

/****************
- 单元内回复
****************/

/****************
- 任务待审核消息
****************/
