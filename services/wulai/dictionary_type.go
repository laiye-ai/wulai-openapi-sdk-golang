package wulai

/*****************词库实体类型*****************/

//EntityType 词库实体类型
type EntityType string

const (
	//LIST_ENTITY_TYPE_ERROR :默认类型
	LIST_ENTITY_TYPE_ERROR EntityType = "LIST_ENTITY_TYPE_ERROR"

	//LIST_ENTITY_TYPE_SYSTEM :预设实体
	LIST_ENTITY_TYPE_SYSTEM EntityType = "LIST_ENTITY_TYPE_SYSTEM"

	//LIST_ENTITY_TYPE_ENUMERATION :枚举实体
	LIST_ENTITY_TYPE_ENUMERATION EntityType = "LIST_ENTITY_TYPE_ENUMERATION"

	//LIST_ENTITY_TYPE_REGEX :正则实体
	LIST_ENTITY_TYPE_REGEX EntityType = "LIST_ENTITY_TYPE_REGEX"

	//LIST_ENTITY_TYPE_INTENT :意图实体
	LIST_ENTITY_TYPE_INTENT EntityType = "LIST_ENTITY_TYPE_INTENT"
)

//DicList 全部词库实体
type DicList struct {
	Entities []DicEntity `json:"entities"`
}

//EntityEnumReponse 创建枚举实体详情
type EntityEnumReponse struct {
	EnumEntity EnumEntity `json:"enum_entity"`
}

//EnumEntity 枚举实体详情
type EnumEntity struct {
	Values []EntityValues `json:"values"`
	ID     int64          `json:"id"`
	Name   string         `json:"name"`
}

//EntityValues 实体值
type EntityValues struct {
	Synonyms      []string `json:"synonyms"`
	StandardValue string   `json:"standard_value"`
}

//EntityReponse 获取实体详情
type EntityReponse struct {
	Entity DicEntity `json:"entity"`
}

//DicEntity 词库实体
type DicEntity struct {
	Type  EntityType  `json:"type"`
	ID    int64       `json:"id"`
	Name  string      `json:"name"`
	Value EntityValue `json:"value"`
}

//EntityValue 实体值详情: 预设实体 / 枚举实体 / 正则实体 / 意图实体
type EntityValue struct {
	IntentEntityValue      IntentEntityValue      `json:"intent_entity_value"`      //意图实体值
	SystemEntityValue      SystemEntityValue      `json:"system_entity_value"`      //预设实体
	RegexEntityValue       RegexEntityValue       `json:"regex_entity_value"`       //正则实体
	EnumerationEntityValue EnumerationEntityValue `json:"enumeration_entity_value"` //枚举实体所有实体值
}

/****************
- 意图实体
****************/

//IntentEntity  意图实体
type IntentEntity struct {
	ID    int64             `json:"id"`
	Value IntentEntityValue `json:"value"`
	Name  string            `json:"name"`
}

//IntentEntityValue 意图实体值
type IntentEntityValue struct {
	Synonyms      []string `json:"synonyms"`       //标准值的相似说法
	StandardValue string   `json:"standard_value"` //标准值
}

//SystemEntityValue 预设实体
type SystemEntityValue struct {
	Description string `json:"description"` //描述
}

//RegexEntityValue 正则实体
type RegexEntityValue struct {
	Regex string `json:"regex"` //正则表达式
}

//EnumerationEntityValue 枚举实体所有实体值
type EnumerationEntityValue struct {
	Values IntentEntityValue `json:"values"` //枚举实体值
}

//IntentEntityResponse 创建意图实体响应
type IntentEntityResponse struct {
	IntentEntity IntentEntity `json:"intent_entity"`
}

/****************
- 专有词汇
****************/

//TermList  专有词汇列表
type TermList struct {
	TermItem  []TermItem `json:"term_item"`
	PageCount int        `json:"page_count"`
}

//TermResponse 操作Term响应
type TermResponse struct {
	TermItem TermItem `json:"term_item"`
}

//TermItem 专有词汇列表
type TermItem struct {
	Term     Term     `json:"term"`     //专有词汇详情
	Synonyms []string `json:"synonyms"` //专有词汇的同义词
}

//Term 专有词汇
type Term struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
