package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 知识点
****************/
func Test_QaKnowledgeTagList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	parentTagID := 139256 //父节点分类id，如果不传值，代表获取根节点下的知识点分类
	page := 1             //页码，代表查看第几页的数据，从1开始
	pageSize := 100       //每页的属性组数量

	resp, err := wulaiClient.QaKnowledgeTagList(parentTagID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaKnowledgeTagList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaKnowledgeTagList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaqaKnowledgeCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeTagID := 139258           //知识点分类id >=1
	standardQuestion := "GOSDK-添加知识点2" //知识点标题
	status := true                     //知识点状态:true: 已生效;false: 未生效
	respondAll := true                 //发送全部回复 true: 已生效;false: 随机一条发送
	maintained := true                 //true: 是;false: 否

	resp, err := wulaiClient.QaqaKnowledgeCreate(knowledgeTagID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaqaKnowledgeCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaqaKnowledgeCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaqaKnowledgeUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeID := 1257716                   //知识点id >=1
	standardQuestion := "GOSDK-添加知识点-update" //知识点标题
	status := true                           //知识点状态:true: 已生效;false: 未生效
	respondAll := true                       //发送全部回复 true: 已生效;false: 随机一条发送
	maintained := true                       //true: 是;false: 否

	resp, err := wulaiClient.QaqaKnowledgeUpdate(knowledgeID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaqaKnowledgeUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaqaKnowledgeUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaKnowledgeItemList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	page := 1      //页码 >=1
	pageSize := 50 //每页的知识点数量[1 .. 200]
	resp, err := wulaiClient.QaKnowledgeItemList(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaKnowledgeItemList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaKnowledgeItemList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

/****************
- 相识问
****************/

func Test_QaSimilarQuestionList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)

	knowledgeID := "12"     //知识点id
	similarQuestionID := "" //相似问id
	page := 1               //页码，代表查看第几页的数据，从1开始
	pageSize := 50          //每页的属性组数量

	resp, err := wulaiClient.QaSimilarQuestionList(knowledgeID, similarQuestionID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionList]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaSimilarQuestionCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)

	knowledgeID := "12" //知识点id
	question := "12"    //相似问
	id := "12"          //相似问id

	resp, err := wulaiClient.QaSimilarQuestionCreate(knowledgeID, question, id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaSimilarQuestionDelete(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	err := wulaiClient.QaSimilarQuestionDelete("id")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 用户属性组
****************/

func Test_QaCreateUserAttributeGroup(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupName := "SDK"
	attributeID := "101185"
	attributeName := "shzy2012"

	resp, err := wulaiClient.QaCreateUserAttributeGroup(groupName, attributeID, attributeName)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaCreateUserAttributeGroup]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaCreateUserAttributeGroup]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaListUserAttributeGroupItems(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	page := 1
	pageSize := 10
	_, err := wulaiClient.QaListUserAttributeGroupItems(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaListUserAttributeGroupItems]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaListUserAttributeGroupItems]=>%s\n", serErr.Error())
		}

		return
	}
}

func Test_QaUpdateUserAttributeGroup(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupID := "6163"
	groupName := "SDK"
	attributes := make(map[string]string)
	attributes["101185"] = "shzy2012-update"

	_, err := wulaiClient.QaUpdateUserAttributeGroup(groupID, groupName, attributes)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUpdateUserAttributeGroup]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUpdateUserAttributeGroup]=>%s\n", serErr.Error())
		}

		return
	}
}
