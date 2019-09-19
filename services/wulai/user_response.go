package wulai

//UserAttributeList 用户属性列表
type UserAttributeList struct {
	PageCount                        int64                             `json:"page_count"`
	UserAttributeUserAttributeValues []UserAttributeUserAttributeValue `json:"user_attribute_user_attribute_values"` //属性列表
}

//UserAttribute 用户属性
type UserAttribute struct {
	ValueType               string `json:"value_type"`                  //用户属性值类型.//TODO:枚举类型
	UseInUserAttributeGroup bool   `json:"use_in_user_attribute_group"` //是否可以作为属性组属性
	Type                    string `json:"type"`                        //用户属性类型. //TODO:枚举类型
	ID                      string `json:"id"`                          //属性id
	Name                    string `json:"name"`                        //属性名
}

//UserAttributeValue 枚举值
type UserAttributeValue struct {
	ID   string `json:"id"`   //属性值id
	Name string `json:"name"` //属性值
}

//UserAttributeUserAttributeValue 属性列表
type UserAttributeUserAttributeValue struct {
	UserAttribute      UserAttribute        `json:"user_attribute"`       //用户属性
	UserAttributeValue []UserAttributeValue `json:"user_attribute_value"` //枚举类型属性的合法枚举值
}
