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

func Test_SceneIntentCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	sceneID := 11431   //意图所属的场景ID>=1
	name := "意图名"      //意图名称[1-200]characters
	lifespanMins := 10 //意图闲置等待时长（分钟），默认3分钟) <= 60
	resp, err := wulaiClient.SceneIntentCreate(sceneID, name, lifespanMins)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneIntentUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 45627          //意图ID >=1
	name := "意图名-update" //意图名称[1-200]characters
	lifespanMins := 10   //意图闲置等待时长（分钟），默认3分钟) <= 60
	resp, err := wulaiClient.SceneIntentUpdate(id, name, lifespanMins)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneIntentStatusUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 45677      //意图ID >=1
	status := true         //：意图状态,意图生成时默认未生效.false: 未生效,true: 已生效
	firstBlockID := 247031 //：该意图的第一个单元ID >=1

	resp, err := wulaiClient.SceneIntentStatusUpdate(intentID, status, firstBlockID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentStatusUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentStatusUpdate]=>%s\n", serErr.Error())
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

	intentID := 45627                               //触发器对应的意图ID >=1
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

func Test_SceneIntentTriggerUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	triggerID := 1419126 //触发器ID >=1
	text := "吃饭-update"  //触发文本[1-200]characters
	resp, err := wulaiClient.SceneIntentTriggerUpdate(triggerID, text)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentTriggerUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentTriggerUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneIntentTriggerDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	triggerID := 1419126 //：触发器ID >=1
	err := wulaiClient.SceneIntentTriggerDelete(triggerID)
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
- 词槽类
****************/

func Test_SceneSlotList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	sceneID := 11431 //词槽所属的场景ID >=1
	page := 1        //页码，代表查看第几页的数据，从1开始
	pageSize := 10   //每页的触发器数量[1-200]

	resp, err := wulaiClient.SceneSlotList(sceneID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneSlotList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneSlotList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneSlotCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	sceneID := 11431          //词槽所属的场景ID >=1
	name := "SDK创建词槽"         //词槽名称 [1-200]characters
	querySlotFilling := false //是否允许整句填槽,默认关闭.true: 开启;false: 关闭
	resp, err := wulaiClient.SceneSlotCreate(sceneID, name, querySlotFilling)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneSlotCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneSlotCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneSlotGet(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 71031 //词槽ID >=1

	resp, err := wulaiClient.SceneSlotGet(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneSlotGet]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneSlotGet]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneSlotUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 71031               //词槽ID >=1
	name := "SDK创建词槽-update"  //词槽名称 [1-200]characters
	querySlotFilling := false //是否允许整句填槽,默认关闭.true: 开启;false: 关闭
	resp, err := wulaiClient.SceneSlotUpdate(id, name, querySlotFilling)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneSlotGet]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneSlotGet]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneSlotDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 71031 //词槽ID >=1
	err := wulaiClient.SceneSlotDelete(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneSlotDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneSlotDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 词槽数据来源
****************/

func Test_SceneSlotDataSourceList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 71034 //词槽ID >=1
	resp, err := wulaiClient.SceneSlotDataSourceList(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneSlotDataSourceList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneSlotDataSourceList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneSlotDataSourceCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	slotID := 71034 //词槽ID >=1
	entityID := 3   //实体ID >=1
	resp, err := wulaiClient.SceneSlotDataSourceCreate(slotID, entityID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneSlotDataSourceCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneSlotDataSourceCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

/****************
- 消息发送单元
****************/

func Test_SceneBlockInformCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 45677        //所属意图ID>=1
	name := "创建消息发送单元"       //单元名称[1-200]characters
	model := RESPONSE_RANDOM //单元回复类型
	resp, err := wulaiClient.SceneBlockInformCreate(intentID, name, model)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockInformCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockInformCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockInformUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 45677              //所属意图ID>=1
	name := "消息发送单元-update"  //单元名称[1-200]characters
	model := RESPONSE_RANDOM //单元回复类型
	resp, err := wulaiClient.SceneBlockInformUpdate(id, name, model)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockInformUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockInformUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockInformGet(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	blockID := 2 //单元ID>=1
	resp, err := wulaiClient.SceneBlockInformGet(blockID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockInformGet]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockInformGet]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

/****************
- 任务待审核消息
****************/
