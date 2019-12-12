package main

import (
	"os"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

func main() {

	//实例化客户端
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := wulai.NewClient(secret, pubkey)
	wulaiClient.Version = "v2" //默认v2
	wulaiClient.SetDebug(true)

	/* ---------------------------------------------------------------
		1. 创建知识点
			a. 调用 QaKnowledgeTagList 方法，查询可用的知识点分类.
			b. 调用 QaKnowledgeCreate 方法，创建知识点.
	   ---------------------------------------------------------------*/

	//1->a:获取知识点分类
	parentTagID := 0 //父节点分类id，如果不传值(或 0)，代表获取根节点下的知识点分类
	page := 1        //页码，代表查看第几页的数据，从1开始
	pageSize := 100  //每页的属性组数量
	tags, err := wulaiClient.QaKnowledgeTagList(parentTagID, page, pageSize)
	if err != nil {
		log.Errorf("[获取知识点分类]=>%s\n", err.Error())
		return
	}

	//1->b:创建知识点
	knowledgeTagID := tags.KnowledgeTags[0].ID //知识点分类id >=1
	standardQuestion := "通用话术"                 //知识点标题
	status := true                             //知识点状态:true: 已生效;false: 未生效
	respondAll := true                         //发送全部回复 true: 已生效;false: 随机一条发送
	maintained := true                         //true: 是;false: 否
	_, err = wulaiClient.QaKnowledgeCreate(knowledgeTagID, standardQuestion, status, respondAll, maintained)
	if err != nil {
		log.Errorf("[创建知识点]=>%s\n", err.Error())
		return
	}
	/* ---------------------------------------------------------------
		2. 创建属性组
			a.调用 QaUserAttributeGroupItemList 方法，查询可用于定义属性组的用户属性.
			b.调用 QaUserAttributeGroupItemCreate 接口，创建属性组.
	   ---------------------------------------------------------------*/

	//2->a:查询可用于定义属性组的用户属性
	groupItem, err := wulaiClient.QaUserAttributeGroupItemList(page, pageSize)
	if err != nil {
		log.Errorf("[用户属性]=>%s\n", err.Error())
		return
	}

	//2->b:创建属性组
	groupName := "GO属性组-1"                                                    //属性组名称 [1~128]characters
	attributeID := groupItem.UserAttributeGroupItems[0].UserAttributeGroup.ID //用户属性ID
	attributeValue := "32"                                                    //用户属性值 [1~128]characters
	_, err = wulaiClient.QaUserAttributeGroupItemCreate(groupName, attributeID, attributeValue)
	if err != nil {
		log.Errorf("[创建属性组]=>%s\n", err.Error())
		return
	}
	/* ---------------------------------------------------------------
	3. 创建属性组回复
		a.调用 QaUserAttributeGroupItemList 方法,查询属性组.
		b.调用 QaKnowledgeItemList 方法,查询知识点.
		c.调用 QaUserAttributeGroupAnswerCreate 方法,创建属性组回复.
	---------------------------------------------------------------*/

	//3->a:查询属性组
	_, err = wulaiClient.QaUserAttributeGroupItemList(page, pageSize)
	if err != nil {
		log.Errorf("[查询属性组]=>%s\n", err.Error())
		return
	}
	//3->b:查询知识点
	_, err = wulaiClient.QaKnowledgeItemList(page, pageSize)
	if err != nil {
		log.Errorf("[查询知识点]=>%s\n", err.Error())
		return
	}
	//3->c:创建属性组回复
	knowledgeID := "1257716"                 //知识点id
	groupID := "6180"                        //属性组ID
	msgBody := &wulai.Text{Content: "创建答案!"} //消息（文本 / 图片 / 语音 / 视频 / 文件 / 图文 / 自定义消息）
	_, err = wulaiClient.QaUserAttributeGroupAnswerCreate(knowledgeID, groupID, msgBody)
	if err != nil {
		log.Errorf("[创建属性组回复]=>%s\n", err.Error())
		return
	}
}
