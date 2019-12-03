package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 词库管理类
****************/

func Test_DicEntityList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	page := 1       //页码，代表查看第几页的数据，从1开始
	pageSize := 100 //每页的属性组数量

	resp, err := wulaiClient.DicEntityList(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicList]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicEntityGet(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	entityID := 31646 //实体ID
	resp, err := wulaiClient.DicEntityGet(entityID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityGet]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityGet]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityGet]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicEntityDelete(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	entityID := 31646 //实体ID
	err := wulaiClient.DicEntityDelete(entityID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityDelete]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityDelete]=>%s\n", err.Error())
		}

		return
	}
}

func Test_DicEntityEnumCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	name := "枚举实体测试" //实体名称
	resp, err := wulaiClient.DicEntityEnumCreate(name)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityEnumCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityEnumCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityEnumCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicEntityEnumValueCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	entityID := 31843                               //实体ID
	standardValue := "标准问题测试"                       //标准值
	synonyms := []string{"相识问-1", "相识问-2", "相识问-3"} //相似说法
	resp, err := wulaiClient.DicEntityEnumValueCreate(entityID, standardValue, synonyms)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityEnumValueCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityEnumValueCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityEnumValueCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicEntityEnumValueDelete(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	entityID := 31843             //实体ID
	standardValue := "标准问题测试"     //标准值
	synonyms := []string{"相识问-3"} //相似说法

	err := wulaiClient.DicEntityEnumValueDelete(entityID, standardValue, synonyms)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityEnumValueDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityEnumValueDelete]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityEnumValueDelete]=>%s\n", err.Error())
		}

		return
	}
}

func Test_DicEntityIntentCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	name := "意图测试"            //实体名称
	standardValue := "意图问题测试" //标准值
	resp, err := wulaiClient.DicEntityIntentCreate(name, standardValue)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityIntentCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityIntentCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityIntentCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicEntityIntentValueCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	entityID := 31787                               //意图实体ID
	synonyms := []string{"相识问-1", "相识问-2", "相识问-3"} //相似说法
	resp, err := wulaiClient.DicEntityIntentValueCreate(entityID, synonyms)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityIntentValueCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityIntentValueCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityIntentValueCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicEntityIntentValueDelete(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	entityID := 31787                      //意图实体ID
	synonyms := []string{"相识问-1", "相识问-2"} //相似说法
	err := wulaiClient.DicEntityIntentValueDelete(entityID, synonyms)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicEntityIntentValueDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicEntityIntentValueDelete]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicEntityIntentValueDelete]=>%s\n", err.Error())
		}

		return
	}
}

/****************
- 专有词汇
****************/

func Test_DicTermList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	page := 1       //页码，代表查看第几页的数据，从1开始
	pageSize := 100 //每页的属性组数量

	resp, err := wulaiClient.DicTermList(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicTermList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicTermList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicTermList]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicTermCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	name := "专有词汇测试"                                //专有词汇的名称
	synonyms := []string{"相识问-1", "相识问-2", "相识问-3"} //相似说法

	resp, err := wulaiClient.DicTermCreate(name, synonyms)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicTermCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicTermCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicTermCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicTermUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 25077                            //专有词汇id
	name := "专有词汇测试"                       //专有词汇的名称
	synonyms := []string{"相识问-1", "相识问-2"} //相似说法

	resp, err := wulaiClient.DicTermUpdate(id, name, synonyms)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicTermUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicTermUpdate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicTermUpdate]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_DicTermDelete(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	id := 25077 //专有词汇id

	err := wulaiClient.DicTermDelete(id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_DicTermDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_DicTermDelete]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_DicTermDelete]=>%s\n", err.Error())
		}

		return
	}
}
