package wulai

/*****************用户属性类型*****************/

//AttributeType 用户属性值类型
type AttributeType string

const (
	//USER_ATTRIBUTE_TYPE_DEFAULT :默认类型
	USER_ATTRIBUTE_TYPE_DEFAULT AttributeType = "USER_ATTRIBUTE_TYPE_DEFAULT"

	//USER_ATTRIBUTE_VALUE_TYPE_STRING :系统属性
	USER_ATTRIBUTE_TYPE_SYSTEM AttributeType = "USER_ATTRIBUTE_TYPE_SYSTEM"

	//USER_ATTRIBUTE_TYPE_TEMP :临时属性
	USER_ATTRIBUTE_TYPE_TEMP AttributeType = "USER_ATTRIBUTE_TYPE_TEMP"
)

/*****************用户属性值类型*****************/

//AttributeValueType 用户属性值类型
type AttributeValueType string

const (
	//USER_ATTRIBUTE_VALUE_TYPE_DEFAULT :默认类型
	USER_ATTRIBUTE_VALUE_TYPE_DEFAULT AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_DEFAULT"

	//USER_ATTRIBUTE_VALUE_TYPE_STRING :字符串类型
	USER_ATTRIBUTE_VALUE_TYPE_STRING AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_STRING"

	//USER_ATTRIBUTE_VALUE_TYPE_ENUM :枚举类型
	USER_ATTRIBUTE_VALUE_TYPE_ENUM AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_ENUM"

	//USER_ATTRIBUTE_VALUE_TYPE_TIME :时间类型
	USER_ATTRIBUTE_VALUE_TYPE_TIME AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_TIME"

	//USER_ATTRIBUTE_VALUE_TYPE_CITY :城市类型
	USER_ATTRIBUTE_VALUE_TYPE_CITY AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_CITY"

	//USER_ATTRIBUTE_VALUE_TYPE_MULTILINE_STRING :多行文本类型
	USER_ATTRIBUTE_VALUE_TYPE_MULTILINE_STRING AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_MULTILINE_STRING"

	//USER_ATTRIBUTE_VALUE_TYPE_INT :整型类型
	USER_ATTRIBUTE_VALUE_TYPE_INT AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_INT"
)

/*****************QA 统计*****************/

//SatisType 用户满意度评价类型
type SatisType string

const (
	//DEFAULT_SATISFACTION [默认]
	DEFAULT_SATISFACTION SatisType = "DEFAULT_SATISFACTION"

	//THUMB_UP [点赞]
	THUMB_UP SatisType = "THUMB_UP"

	//THUMB_UP [回答了我的问题，但答案不够好]
	BAD_ANSWER SatisType = "BAD_ANSWER"

	//WRONG_ANSWER [没有回答我的问题]
	WRONG_ANSWER SatisType = "WRONG_ANSWER"

	//REPORT [举报]
	REPORT SatisType = "REPORT"
)

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
