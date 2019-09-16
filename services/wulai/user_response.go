package wulai

//UserAttributeList 用户属性列表
type UserAttributeList struct {
	PageCount                        int `json:"page_count"`
	UserAttributeUserAttributeValues []struct {
		UserAttribute struct {
			ValueType               string `json:"value_type"`
			UseInUserAttributeGroup bool   `json:"use_in_user_attribute_group"`
			Type                    string `json:"type"`
			ID                      string `json:"id"`
			Name                    string `json:"name"`
		} `json:"user_attribute"`
		UserAttributeValue []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"user_attribute_value"`
	} `json:"user_attribute_user_attribute_values"`
}

//MsgHistory 历史消息
type MsgHistory struct {
	Msg []struct {
		Direction  string `json:"direction"`
		SenderInfo struct {
			AvatarURL string `json:"avatar_url"`
			Nickname  string `json:"nickname"`
			RealName  string `json:"real_name"`
		} `json:"sender_info"`
		MsgType  string `json:"msg_type"`
		Extra    string `json:"extra"`
		MsgID    string `json:"msg_id"`
		MsgTs    string `json:"msg_ts"`
		UserInfo struct {
			AvatarURL string `json:"avatar_url"`
			Nickname  string `json:"nickname"`
		} `json:"user_info"`
		MsgBody struct {
			Text struct {
				Content string `json:"content"`
			} `json:"text"`
		} `json:"msg_body"`
	} `json:"msg"`
	HasMore bool `json:"has_more"`
}

//MsgReceive 接收用户发的消息
type MsgReceive struct {
	MsgID string `json:"msg_id"`
}

//MsgSync 同步发给用户的消息
type MsgSync struct {
	MsgID string `json:"msg_id"`
}
