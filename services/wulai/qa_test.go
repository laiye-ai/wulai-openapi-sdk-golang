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

	parentTagID := 0 //父节点分类id，如果值为0，代表获取根节点下的知识点分类
	page := 1        //页码，代表查看第几页的数据，从1开始
	pageSize := 100  //每页的属性组数量

	resp, err := wulaiClient.QaKnowledgeTagList(parentTagID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaKnowledgeTagList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaKnowledgeTagList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaKnowledgeTagList]=>%s\n", err.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaKnowledgeCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeTagID := "139258"         //知识点分类id >=1
	standardQuestion := "GOSDK-添加知识点2" //知识点标题
	status := true                     //知识点状态:true: 已生效;false: 未生效
	respondAll := true                 //发送全部回复 true: 已生效;false: 随机一条发送
	maintained := true                 //true: 是;false: 否

	resp, err := wulaiClient.QaKnowledgeCreate(knowledgeTagID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaKnowledgeCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaKnowledgeCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaKnowledgeCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaKnowledgeUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeID := int64(1257716)            //知识点id >=1
	standardQuestion := "GOSDK-添加知识点-update" //知识点标题
	status := true                           //知识点状态:true: 已生效;false: 未生效
	respondAll := true                       //发送全部回复 true: 已生效;false: 随机一条发送
	maintained := true                       //true: 是;false: 否

	resp, err := wulaiClient.QaKnowledgeUpdate(knowledgeID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaKnowledgeUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaKnowledgeUpdate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaKnowledgeUpdate]=>%s\n", err.Error())
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
		} else {
			log.Infof("[Test_QaKnowledgeItemList]=>%s\n", err.Error())
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
	wulaiClient.SetDebug(true)

	knowledgeID := "1276766" //知识点id
	similarQuestionID := "1" //相似问id
	page := 1                //页码，代表查看第几页的数据，从1开始
	pageSize := 50           //每页的属性组数量

	resp, err := wulaiClient.QaSimilarQuestionList(knowledgeID, similarQuestionID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaSimilarQuestionList]=>%s\n", err.Error())
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
		} else {
			log.Infof("[Test_QaSimilarQuestionCreate]=>%s\n", err.Error())
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
		} else {
			log.Infof("[Test_QaSimilarQuestionUpdate]=>%s\n", err.Error())
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
		} else {
			log.Infof("[Test_QaSimilarQuestionDelete]=>%s\n", err.Error())
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
	_, err := wulaiClient.QaUserAttributeGroupItemList(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaListUserAttributeGroupItem]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaListUserAttributeGroupItem]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaListUserAttributeGroupItem]=>%s\n", err.Error())
		}

		return
	}
}

func Test_QaUserAttributeGroupItemCreate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupName := "GO 属性组-2" //属性组名称 [1~128]characters
	attributeID := "101669" //用户属性ID
	attributeValue := "2"   //用户属性值 [1~128]characters
	resp, err := wulaiClient.QaUserAttributeGroupItemCreate(groupName, attributeID, attributeValue)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupItemCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupItemCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaUserAttributeGroupItemCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaUserAttributeGroupItemUpdate(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupID := "6180"
	groupName := "GO 属性组-1"
	attributes := make(map[string]string)
	attributes["101669"] = "shzy2012-update"

	_, err := wulaiClient.QaUserAttributeGroupItemUpdate(groupID, groupName, attributes)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupItemUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupItemUpdate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaUserAttributeGroupItemUpdate]=>%s\n", err.Error())
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

	knowledgeID := string(1257716) //知识点id，如=0，返回所有知识点
	groupID := string(6180)        //属性组id，如=0，返回所有属性组
	page := 1
	pageSize := 10

	resp, err := wulaiClient.QaUserAttributeGroupAnswerList(knowledgeID, groupID, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupAnswerList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupAnswerList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaUserAttributeGroupAnswerList]=>%s\n", err.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)
}

func Test_QaUserAttributeGroupAnswerCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	knowledgeID := "1257716"  //知识点id
	groupID := "6180"         //属性组ID
	msgBody := &Text{"创建答案!"} //消息（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）
	resp, err := wulaiClient.QaUserAttributeGroupAnswerCreate(knowledgeID, groupID, msgBody)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaUserAttributeGroupAnswerCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaUserAttributeGroupAnswerCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_QaUserAttributeGroupAnswerCreate]=>%s\n", err.Error())
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
		} else {
			log.Infof("[Test_QaUserAttributeGroupAnswerUpdate]=>%s\n", err.Error())
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
		} else {
			log.Infof("[Test_QaUserAttributeGroupAnswerDelete]=>%s\n", err.Error())
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
		} else {
			log.Infof("[Test_QaSatisCreate]=>%s\n", err.Error())
		}

		return
	}
}
