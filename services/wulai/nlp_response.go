package wulai

//ExtractType 抽取实体类型
type ExtractType string

const (
	//EXTRACT_ENTITIES_TYPE_ERROR 错误
	EXTRACT_ENTITIES_TYPE_ERROR ExtractType = "EXTRACT_ENTITIES_TYPE_ERROR"
	//EXTRACT_ENTITIES_TYPE_SYSTEM 预设实体
	EXTRACT_ENTITIES_TYPE_SYSTEM ExtractType = "EXTRACT_ENTITIES_TYPE_SYSTEM"
	//EXTRACT_ENTITIES_TYPE_ENUMERATE 枚举实体
	EXTRACT_ENTITIES_TYPE_ENUMERATE ExtractType = "EXTRACT_ENTITIES_TYPE_ENUMERATE"
	//EXTRACT_ENTITIES_TYPE_REGULAR 正则实体
	EXTRACT_ENTITIES_TYPE_REGULAR ExtractType = "EXTRACT_ENTITIES_TYPE_REGULAR"
)

//NLPEntitiesExtractReponse 实体抽取结果
type NLPEntitiesExtractReponse struct {
	Entities []Entities `json:"entities"` //实体抽取的返回值
}

//Entities 实体抽取值
type Entities struct {
	Type     ExtractType `json:"type"`      //实体类型
	IdxStart int         `json:"idx_start"` //起始位置，文本的开始字符位置
	Entity   NLPEntity   `json:"entity"`    //系统、枚举、正则 三选一
}

//NLPEntity NLP实体
type NLPEntity struct {
	EnumerateEntity EnumerateEntity `json:"enumerate_entity"` //枚举实体
	SystemEntity    SystemEntity    `json:"system_entity"`    //系统实体
	RegularEntity   RegularEntity   `json:"regular_entity"`   //正则实体
}

//SystemEntity 系统实体
type SystemEntity struct {
	Text          string `json:"text"`           //文本
	StandardValue string `json:"standard_value"` //标准值(归一化值)
	Name          string `json:"name"`           //实体名称
}

//EnumerateEntity 枚举实体
type EnumerateEntity struct {
	Text          string   `json:"text"`           //文本
	Synonyms      []string `json:"synonyms"`       //标准值的相似说法,不含text以及standard_value
	StandardValue string   `json:"standard_value"` //标准值(归一化值)
	Name          string   `json:"name"`           //实体名称
}

//RegularEntity 正则实体
type RegularEntity struct {
	Text string `json:"text"` //文本
	Name string `json:"name"` //实体名称
}

//NLPTokenizeResponse 分词&词性标注结果
type NLPTokenizeResponse struct {
	Tokens []Tokens `json:"tokens"`
}

//Tokens 文本&词性
type Tokens struct {
	Text     string `json:"text"`
	Pos      string `json:"pos"`
	IdxStart int    `json:"idx_start"`
}
