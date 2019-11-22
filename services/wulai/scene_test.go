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

	id := 45677        //意图ID >=1
	name := "问卷调查机器人"  //意图名称[1-200]characters
	lifespanMins := 10 //意图闲置等待时长（分钟），默认3分钟) <= 60
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

func Test_SceneIntentDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 1419126 //意图ID >=1
	err := wulaiClient.SceneIntentDelete(id)
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
- 触发器
****************/

func Test_SceneIntentTriggerList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 45677 //意图ID >=1
	page := 1         //页码，代表查看第几页的数据，从1开始
	pageSize := 10    //每页的触发器数量[1-200]

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

	intentID := 45677                               //触发器对应的意图ID >=1
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
	entityID := 9   //实体ID >=1
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
- 询问填槽单元
****************/

func Test_SceneBlockRequestGet(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 250542 //单元ID >=1
	resp, err := wulaiClient.SceneBlockRequestGet(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockRequestGet]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockRequestGet]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockRequestCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	param := &BlockRequestParam{} //指针类型
	param.IntentID = 45677
	param.Name = "询问-姓名"
	param.DefaultSlotValue = "请问您的姓名？"
	param.SlotFillingWhenAsked = false
	param.SlotID = 71034
	param.Mode = RESPONSE_LOOP
	param.RequestCount = 1
	resp, err := wulaiClient.SceneBlockRequestCreate(param)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockRequestCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockRequestCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockRequestUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	param := &BlockRequestParam{} //指针类型
	param.ID = 250542
	param.IntentID = 45677
	param.Name = "询问-您的姓名是？"
	param.DefaultSlotValue = "请问您的姓名？"
	param.SlotFillingWhenAsked = false
	param.SlotID = 71034
	param.Mode = RESPONSE_LOOP
	param.RequestCount = 1
	resp, err := wulaiClient.SceneBlockRequestUpdate(param)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockRequestUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockRequestUpdate]=>%s\n", serErr.Error())
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

	intentID := 45677      //所属意图ID>=1
	name := "GO-消息发送单元-1"  //单元名称[1-200]characters
	model := RESPONSE_LOOP //单元回复类型
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

	blockID := 250475 //单元ID>=1
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

func Test_SceneIntentTriggerLearningList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	page := 1      //页码，代表查看第几页的数据，从1开始
	pageSize := 10 //每页的触发器数量[1-200]
	resp, err := wulaiClient.SceneIntentTriggerLearningList(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentTriggerLearningList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentTriggerLearningList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneIntentTriggerLearningDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 5 //待审核消息ID >=1
	err := wulaiClient.SceneIntentTriggerLearningDelete(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneIntentTriggerLearningDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneIntentTriggerLearningDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 单元内回复
****************/

func Test_SceneBlockResponseCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	blockID := 250542       //单元ID>=1
	msgBody := &Text{"您好!"} //传入指针.消息体格式，任意选择一种消息类型[文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息]
	resp, err := wulaiClient.SceneBlockResponseCreate(blockID, msgBody)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockResponseCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockResponseCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockResponseUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 1                 //回复ID >=1
	msgBody := &Text{"您好!"} //传入指针.消息体格式，任意选择一种消息类型[文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息]
	resp, err := wulaiClient.SceneBlockResponseUpdate(id, msgBody)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockResponseUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockResponseUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockResponseDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 5 //回复ID >=1
	err := wulaiClient.SceneBlockResponseDelete(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockResponseDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockResponseDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 单元类
****************/

func Test_SceneBlockList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 45677 //意图ID
	page := 1         //页码，代表查看第几页的数据，从1开始
	pageSize := 10    //每页的触发器数量[1-200]
	resp, err := wulaiClient.SceneBlockList(intentID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 250475 //单元ID >=1
	err := wulaiClient.SceneBlockDelete(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 单元关系
****************/

func Test_SceneBlockRelationCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 45677       //意图ID >=1
	fromBlockID := 250542   //当前单元ID >=1
	toBlockID := 247031     //下一个单元ID >=1
	condition := &Default{} //默认 单元跳转条件
	resp, err := wulaiClient.SceneBlockRelationCreate(intentID, fromBlockID, toBlockID, condition)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockRelationCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockRelationCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockRelationDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 325651 //单元关系ID >=1
	err := wulaiClient.SceneBlockRelationDelete(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockRelationDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockRelationDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 意图终点单元
****************/

func Test_SceneBlockEndBlockCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 45677       //所属意图ID >=1
	name := "结束意图"          //单元名称[1-200]characters
	slotMemorizing := false //是否保存词槽值(默认关闭).true: 开启;false: 关闭
	action := &End{}        //结束单元跳转方式 (指定意图 / 上个意图 / 不跳转)) 指针类型
	resp, err := wulaiClient.SceneBlockEndBlockCreate(intentID, name, slotMemorizing, action)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockEndBlockCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockEndBlockCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockEndBlockUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	intentID := 45677       //所属意图ID >=1
	id := 254445            //单元ID>=1
	name := "结束意图-for QA"   //单元名称[1-200]characters
	slotMemorizing := false //是否保存词槽值(默认关闭).true: 开启;false: 关闭
	action := &End{}        //结束单元跳转方式 (指定意图 / 上个意图 / 不跳转)) 指针类型
	resp, err := wulaiClient.SceneBlockEndBlockUpdate(intentID, id, name, slotMemorizing, action)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockEndBlockCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockEndBlockCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_SceneBlockEndBlockGet(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 254445 //单元ID>=1
	resp, err := wulaiClient.SceneBlockEndBlockGet(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_SceneBlockEndBlockGet]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_SceneBlockEndBlockGet]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}
