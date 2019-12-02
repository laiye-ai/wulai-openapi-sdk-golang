package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

func Test_StatsRecallDailyList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	startDate := "20191001" //开始日期，格式如19700101
	endDate := "20191030"   //结束日期，格式如19700101

	resp, err := wulaiClient.StatsRecallDailyList(startDate, endDate)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_StatsRecallDailyList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_StatsRecallDailyList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_StatsRecallDailyList]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_StatsRecallDailyKnowledgeList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	startDate := "20191020" //开始日期，格式如19700101
	endDate := "20191110"   //结束日期，格式如19700101
	page := 1               //页码，代表查看第几页的数据 page>=1
	pageSize := 100         //每页的知识点数量 [1~200]

	resp, err := wulaiClient.StatsRecallDailyKnowledgeList(startDate, endDate, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_StatsRecallDailyKnowledgeList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_StatsRecallDailyKnowledgeList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_StatsRecallDailyKnowledgeList]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_StatsDailyKnowledgeList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	startDate := "20191020" //开始日期，格式如19700101
	endDate := "20191110"   //结束日期，格式如19700101
	page := 1               //页码，代表查看第几页的数据 page>=1
	pageSize := 100         //每页的知识点数量 [1~200]

	resp, err := wulaiClient.StatsDailyKnowledgeList(startDate, endDate, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_StatsDailyKnowledgeList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_StatsDailyKnowledgeList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_StatsDailyKnowledgeList]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}
