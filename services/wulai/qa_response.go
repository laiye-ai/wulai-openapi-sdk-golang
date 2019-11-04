package wulai

/****************
- 知识点
****************/

//QaKnowledgeTagList 知识点分类列表
type QaKnowledgeTagList struct {
	KnowledgeTags []KnowledgeTag `json:"knowledge_tags"`
	PageCount     int            `json:"page_count"`
}

//KnowledgeTag 知识点分类
type KnowledgeTag struct {
	ParentKnowledgeTagID string `json:"parent_knowledge_tag_id"`
	ID                   string `json:"id"`
	Name                 string `json:"name"`
}

//QaKnowledgeTagResponse 创建知识点返回的结构体
type QaKnowledgeTagResponse struct {
	KnowledgeTagKnowledge KnowledgeTagKnowledge `json:"knowledge_tag_knowledge"`
}

//KnowledgeTagKnowledge 知识点
type KnowledgeTagKnowledge struct {
	Knowledge      Knowledge `json:"knowledge"`
	KnowledgeTagID string    `json:"knowledge_tag_id"`
}

//Knowledge 知识点详情
type Knowledge struct {
	Status                         bool   `json:"status"`
	UpdateTime                     string `json:"update_time"`
	MaintainedByUserAttributeGroup bool   `json:"maintained_by_user_attribute_group"`
	StandardQuestion               string `json:"standard_question"`
	CreateTime                     string `json:"create_time"`
	RespondAll                     bool   `json:"respond_all"`
	ID                             string `json:"id"`
}

//QaKnowledgeItemListResponse 知识点列表
type QaKnowledgeItemListResponse struct {
	PageCount      int             `json:"page_count"`
	KnowledgeItems []KnowledgeItem `json:"knowledge_items"`
}

//KnowledgeItem 知识点
type KnowledgeItem struct {
	KnowledgeTag     KnowledgeTag      `json:"knowledge_tag"`
	SimilarQuestions []SimilarQuestion `json:"similar_questions"`
	Knowledge        Knowledge         `json:"knowledge"`
}

//SimilarQuestion 相似问
type SimilarQuestion struct {
	KnowledgeID string `json:"knowledge_id"` //知识点id
	Question    string `json:"question"`     //相似问
	ID          string `json:"id"`           //相似问id
}

//QaKnowledgeResponse 更新知识点返回结构体
type QaKnowledgeResponse struct {
	Knowledge Knowledge `json:"knowledge"`
}

/****************
- 相似问
****************/

//QaSimilarQuestionResponse 相似问结构体
type QaSimilarQuestionResponse struct {
	SimilarQuestion SimilarQuestion `json:"similar_question"`
}

//QaSimilarQuestionList 相识问列表
type QaSimilarQuestionList struct {
	SimilarQuestions []SimilarQuestion `json:"similar_questions"`
	PageCount        int               `json:"page_count"`
}

/****************
- 用户属性组
****************/

//QaUserAttributeGroupItemList 列表
type QaUserAttributeGroupItemList struct {
	PageCount               int                        `json:"page_count"`
	UserAttributeGroupItems []QaUserAttributeGroupItem `json:"user_attribute_group_items"`
}

//QaUserAttributeGroupResponse 创建属性返回的结果
type QaUserAttributeGroupResponse struct {
	UserAttributeGroupItem QaUserAttributeGroupItem `json:"user_attribute_group_item"`
}

//QaUserAttributeGroupItem 用户属性组及属性
type QaUserAttributeGroupItem struct {
	UserAttributeUserAttributeValues []QaUserAttributeUserAttributeValues `json:"user_attribute_user_attribute_values"`
	UserAttributeGroup               QaUserAttributeGroup                 `json:"user_attribute_group"`
}

//QaUserAttributeUserAttributeValues 属性
type QaUserAttributeUserAttributeValues struct {
	UserAttribute      UserAttribute      `json:"user_attribute"`
	UserAttributeValue UserAttributeValue `json:"user_attribute_value"`
}

//QaUserAttributeGroup 用户属性组详情
type QaUserAttributeGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/****************
- 属性组回复
****************/

//QaUserAttributeGroupAnswerList 属性组回复列表
type QaUserAttributeGroupAnswerList struct {
	UserAttributeGroupAnswers []UserAttributeGroupAnswer `json:"user_attribute_group_answers"`
	PageCount                 int                        `json:"page_count"`
}

//Answer 回复
type Answer struct {
	KnowledgeID string  `json:"knowledge_id"` //知识点id
	MsgBody     MsgBody `json:"msg_body"`     //消息体格式
	ID          string  `json:"id"`           //回复id
}

//UserAttributeGroupAnswer 属性组id
type UserAttributeGroupAnswer struct {
	Answer               Answer `json:"answer"`
	UserAttributeGroupID string `json:"user_attribute_group_id"`
}

//QaUserAttributeGroupAnswer 属性组
type QaUserAttributeGroupAnswer struct {
	UserAttributeGroupAnswer UserAttributeGroupAnswer `json:"user_attribute_group_answer"`
}
