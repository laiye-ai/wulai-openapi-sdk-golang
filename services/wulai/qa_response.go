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
	KnowledgeID string `json:"knowledge_id"`
	Question    string `json:"question"`
	ID          string `json:"id"`
}

//QaKnowledgeResponse 更新知识点返回结构体
type QaKnowledgeResponse struct {
	Knowledge Knowledge `json:"knowledge"`
}

/****************
- 相识问
****************/

//QaSimilarQuestionResponse 创建返回的响应体-相识问结构体
type QaSimilarQuestionResponse struct {
	SimilarQuestion struct {
		KnowledgeID string `json:"knowledge_id"` //知识点id
		Question    string `json:"question"`     //相似问
		ID          string `json:"id"`           //相似问id
	} `json:"similar_question"`
}

//QaSimilarQuestionList 相识问列表
type QaSimilarQuestionList struct {
	SimilarQuestions []struct {
		KnowledgeID string `json:"knowledge_id"`
		Question    string `json:"question"`
		ID          string `json:"id"`
	} `json:"similar_questions"`
	PageCount int `json:"page_count"`
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
	UserAttribute      QaUserAttribute      `json:"user_attribute"`
	UserAttributeValue QaUserAttributeValue `json:"user_attribute_value"`
}

//QaUserAttributeGroup 用户属性组详情
type QaUserAttributeGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//QaUserAttribute 用户属性
type QaUserAttribute struct {
	Name                    string             `json:"name"`
	Lifespan                int                `json:"lifespan"`
	ValueType               AttributeValueType `json:"value_type"`
	UseInUserAttributeGroup bool               `json:"use_in_user_attribute_group"`
	Type                    AttributeType      `json:"type"`
	ID                      string             `json:"id"`
}

//QaUserAttributeValue 用户属性值
type QaUserAttributeValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
