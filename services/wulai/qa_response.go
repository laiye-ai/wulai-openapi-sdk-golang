package wulai

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
