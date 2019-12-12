package wulai

//User 用户
type User struct {
	UserID    string `json:"user_id"`    //用户id作为用户的唯一识别。如果调用方客户端用户没有唯一标识，尽量通过其他标识来唯一区分用户，如设备号 [1 .. 128 ] characters
	AvatarURL string `json:"avatar_url"` //用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
	Nickname  string `json:"nickname"`   //用户昵称 <= 128 characters
}

//UserAttribute 用户属性
type UserAttribute struct {
	ID                      string             `json:"id"`                          //属性id  TODO: 如果修改为int64 则抛出 unmarshal string错误
	Name                    string             `json:"name"`                        //属性名
	Lifespan                int                `json:"lifespan"`                    //属性的有效期(单位为秒)。如果为永久, 则有效期为0
	ValueType               AttributeValueType `json:"value_type"`                  //用户属性值类型.
	UseInUserAttributeGroup bool               `json:"use_in_user_attribute_group"` //是否可以作为属性组属性
	Type                    AttributeType      `json:"type"`                        //用户属性类型.
}

//UserAttributeValue 枚举值
type UserAttributeValue struct {
	ID   string `json:"id"`   //属性值id
	Name string `json:"name"` //属性值
}

//UserAttributeUserAttributeValues 属性列表[数组]]
type UserAttributeUserAttributeValues struct {
	UserAttribute      UserAttribute        `json:"user_attribute"`       //用户属性
	UserAttributeValue []UserAttributeValue `json:"user_attribute_value"` //枚举类型属性的合法枚举值
}

//UserAttributeUserAttributeValue 属性列表
type UserAttributeUserAttributeValue struct {
	UserAttribute      UserAttribute      `json:"user_attribute"`       //用户属性
	UserAttributeValue UserAttributeValue `json:"user_attribute_value"` //枚举类型属性的合法枚举值
}

//UserAttributeList 用户属性列表
type UserAttributeList struct {
	PageCount                        int                                `json:"page_count"`
	UserAttributeUserAttributeValues []UserAttributeUserAttributeValues `json:"user_attribute_user_attribute_values"` //属性列表
}

//UserAttributePairListResponse 获取用户的属性结构体
type UserAttributePairListResponse struct {
	UserAttributeUserAttributeValues []UserAttributeUserAttributeValue `json:"user_attribute_user_attribute_values"`
}
