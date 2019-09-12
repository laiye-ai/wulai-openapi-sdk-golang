package wulai

//BotResponse 机器人回复
type BotResponse struct {
	IsDispatch            bool   `json:"is_dispatch"`
	MsgID                 string `json:"msg_id"`
	TaskSuggestedResponse []struct {
		Score  float64 `json:"score"`
		IsSend bool    `json:"is_send"`
		Task   struct {
			BlockType string `json:"block_type"`
			BlockID   int    `json:"block_id"`
			TaskID    int    `json:"task_id"`
			BlockName string `json:"block_name"`
			Entities  []struct {
				IdxEnd   int    `json:"idx_end"`
				Name     string `json:"name"`
				IdxStart int    `json:"idx_start"`
				Value    string `json:"value"`
				SegValue string `json:"seg_value"`
				Type     string `json:"type"`
				Desc     string `json:"desc"`
			} `json:"entities"`
			TaskName string `json:"task_name"`
			RobotID  int    `json:"robot_id"`
		} `json:"task"`
		Response []struct {
			MsgBody struct {
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"msg_body"`
			SimilarResponse []struct {
				URL    string `json:"url"`
				Source string `json:"source"`
				Detail struct {
					Qa struct {
						KnowledgeID      int    `json:"knowledge_id"`
						StandardQuestion string `json:"standard_question"`
						Question         string `json:"question"`
					} `json:"qa"`
				} `json:"detail"`
			} `json:"similar_response"`
			EnableEvaluate bool `json:"enable_evaluate"`
			DelayTs        int  `json:"delay_ts"`
		} `json:"response"`
		QuickReply []string `json:"quick_reply"`
	} `json:"task_suggested_response"`
}

//BotResponseQa 问答机器人回复
type BotResponseQa struct {
	IsDispatch          bool   `json:"is_dispatch"`
	MsgID               string `json:"msg_id"`
	QaSuggestedResponse []struct {
		Qa struct {
			KnowledgeID      int    `json:"knowledge_id"`
			StandardQuestion string `json:"standard_question"`
			Question         string `json:"question"`
		} `json:"qa"`
		IsSend   bool    `json:"is_send"`
		Score    float64 `json:"score"`
		Response []struct {
			MsgBody struct {
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"msg_body"`
			SimilarResponse []struct {
				URL    string `json:"url"`
				Source string `json:"source"`
				Detail struct {
					Qa struct {
						KnowledgeID      int    `json:"knowledge_id"`
						StandardQuestion string `json:"standard_question"`
						Question         string `json:"question"`
					} `json:"qa"`
				} `json:"detail"`
			} `json:"similar_response"`
			EnableEvaluate bool `json:"enable_evaluate"`
			DelayTs        int  `json:"delay_ts"`
		} `json:"response"`
		QuickReply []string `json:"quick_reply"`
	} `json:"qa_suggested_response"`
}

//BotResponseKeyword 关键字机器人回复
type BotResponseKeyword struct {
	IsDispatch               bool   `json:"is_dispatch"`
	MsgID                    string `json:"msg_id"`
	KeywordSuggestedResponse []struct {
		IsSend   bool    `json:"is_send"`
		Score    float64 `json:"score"`
		Response []struct {
			MsgBody struct {
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"msg_body"`
			SimilarResponse []struct {
				URL    string `json:"url"`
				Source string `json:"source"`
				Detail struct {
					Qa struct {
						KnowledgeID      int    `json:"knowledge_id"`
						StandardQuestion string `json:"standard_question"`
						Question         string `json:"question"`
					} `json:"qa"`
				} `json:"detail"`
			} `json:"similar_response"`
			EnableEvaluate bool `json:"enable_evaluate"`
			DelayTs        int  `json:"delay_ts"`
		} `json:"response"`
		Keyword struct {
			KeywordID int    `json:"keyword_id"`
			Keyword   string `json:"keyword"`
		} `json:"keyword"`
		QuickReply []string `json:"quick_reply"`
	} `json:"keyword_suggested_response"`
}

//BotResponseTask 任务机器人回复
type BotResponseTask struct {
	IsDispatch            bool   `json:"is_dispatch"`
	MsgID                 string `json:"msg_id"`
	TaskSuggestedResponse []struct {
		Score  float64 `json:"score"`
		IsSend bool    `json:"is_send"`
		Task   struct {
			BlockType string `json:"block_type"`
			BlockID   int    `json:"block_id"`
			TaskID    int    `json:"task_id"`
			BlockName string `json:"block_name"`
			Entities  []struct {
				IdxEnd   int    `json:"idx_end"`
				Name     string `json:"name"`
				IdxStart int    `json:"idx_start"`
				Value    string `json:"value"`
				SegValue string `json:"seg_value"`
				Type     string `json:"type"`
				Desc     string `json:"desc"`
			} `json:"entities"`
			TaskName string `json:"task_name"`
			RobotID  int    `json:"robot_id"`
		} `json:"task"`
		Response []struct {
			MsgBody struct {
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"msg_body"`
			SimilarResponse []struct {
				URL    string `json:"url"`
				Source string `json:"source"`
				Detail struct {
					Qa struct {
						KnowledgeID      int    `json:"knowledge_id"`
						StandardQuestion string `json:"standard_question"`
						Question         string `json:"question"`
					} `json:"qa"`
				} `json:"detail"`
			} `json:"similar_response"`
			EnableEvaluate bool `json:"enable_evaluate"`
			DelayTs        int  `json:"delay_ts"`
		} `json:"response"`
		QuickReply []string `json:"quick_reply"`
	} `json:"task_suggested_response"`
}

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

//direction
type direction string

const (
	//BACKWARD :向旧的消息翻页，查询比传入msg_id更小的消息
	BACKWARD direction = "BACKWARD"

	//FORWARD :先新的消息翻页，查询比传入msg_id更大的消息
	FORWARD direction = "FORWARD"
)
