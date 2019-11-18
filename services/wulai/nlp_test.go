package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 自然语言处理类
****************/

func Test_NLPEntitiesExtract(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	query := "今天上海天气怎么样?" //待实体抽取query [1-1024]characters
	resp, err := wulaiClient.NLPEntitiesExtract(query)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_NLPEntitiesExtract]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_NLPEntitiesExtract]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_NLPTokenize(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	query := "今天上海天气怎么样?" //待实体抽取query [1-1024]characters
	resp, err := wulaiClient.NLPTokenize(query)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_NLPTokenize]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_NLPTokenize]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}
