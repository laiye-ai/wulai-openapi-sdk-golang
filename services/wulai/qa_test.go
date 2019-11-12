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
- 相似问
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

	knowledgeID := "1257716" //知识点id
	question := "我是相似问--创建"  //相似问

	resp, err := wulaiClient.QaSimilarQuestionCreate(knowledgeID, question)
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

func Test_QaSimilarQuestionUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)

	knowledgeID := "1257716" //知识点id
	question := "我是相似问--更新"  //相似问
	id := "11648029"         //相似问id

	resp, err := wulaiClient.QaSimilarQuestionUpdate(knowledgeID, question, id)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaSimilarQuestionDelete(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	err := wulaiClient.QaSimilarQuestionDelete("11648029")
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

func Test_QaUserAttributeGroupItemList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	page := 1
	pageSize := 10
	_, err := wulaiClient.QaAttributeGroupItemList(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaListUserAttributeGroupItem]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaListUserAttributeGroupItem]=>%s\n", serErr.Error())
		}

		return
	}
}

func Test_QaUserAttributeGroupCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupName := "GO 属性组-2"
	attributeID := "101669"
	attributeName := "2"

	resp, err := wulaiClient.QaUserAttributeGroupCreate(groupName, attributeID, attributeName)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaUserAttributeGroupUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupID := "6180"
	groupName := "GO 属性组-1"
	attributes := make(map[string]string)
	attributes["101669"] = "shzy2012-update"

	_, err := wulaiClient.QaUserAttributeGroupUpdate(groupID, groupName, attributes)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupUpdate]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- 属性组回复
****************/
func Test_QaUserAttributeGroupAnswerList(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeID := "1257716"
	groupID := "6180"
	page := 1
	pageSize := 10

	resp, err := wulaiClient.QaUserAttributeGroupAnswerList(knowledgeID, groupID, page, pageSize)
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

func Test_QaUserAttributeGroupAnswerCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeID := "1257716"
	groupID := "6180"
	text := &Text{"创建一个答案!"} //传入指针
	resp, err := wulaiClient.QaUserAttributeGroupAnswerCreate(knowledgeID, groupID, text)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupAnswerCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupAnswerCreate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaUserAttributeGroupAnswerUpdate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeID := "1257716"
	groupID := "6180"
	answerID := "2919012"
	text := &Text{"您好，我更新了答案!"} //传入指针
	resp, err := wulaiClient.QaUserAttributeGroupAnswerUpdate(knowledgeID, groupID, answerID, text)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupAnswerUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupAnswerUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaUserAttributeGroupAnswerDelete(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	answerID := "2919490"
	err := wulaiClient.QaUserAttributeGroupAnswerDelete(answerID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupAnswerDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupAnswerDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

/****************
- QA 统计
****************/
func Test_QaSatisCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	satisType := DEFAULT_SATISFACTION
	userID := "1"
	knowledgeID := "1"
	msgID := "1"
	err := wulaiClient.QaSatisCreate(satisType, userID, knowledgeID, msgID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSatisCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSatisCreate]=>%s\n", serErr.Error())
		}

		return
	}
}
