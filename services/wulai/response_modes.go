package wulai

import (
	"encoding/json"
	"fmt"
)

//Response base
type Response struct{}

//BotResponse 机器人回复
type BotResponse struct {
	Response
	IsDispatch            bool   `json:"is_dispatch"`
	MsgID                 string `json:"msg_id"`
	TaskSuggestedResponse []struct {
		Score  int  `json:"score"`
		IsSend bool `json:"is_send"`
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
	Response
	IsDispatch          bool   `json:"is_dispatch"`
	MsgID               string `json:"msg_id"`
	QaSuggestedResponse []struct {
		Qa struct {
			KnowledgeID      int    `json:"knowledge_id"`
			StandardQuestion string `json:"standard_question"`
			Question         string `json:"question"`
		} `json:"qa"`
		IsSend   bool `json:"is_send"`
		Score    int  `json:"score"`
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
	Response
	IsDispatch               bool   `json:"is_dispatch"`
	MsgID                    string `json:"msg_id"`
	KeywordSuggestedResponse []struct {
		IsSend   bool `json:"is_send"`
		Score    int  `json:"score"`
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
	Response
	IsDispatch            bool   `json:"is_dispatch"`
	MsgID                 string `json:"msg_id"`
	TaskSuggestedResponse []struct {
		Score  int  `json:"score"`
		IsSend bool `json:"is_send"`
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

//ToByte 返回 字节
func (x *Response) ToByte() ([]byte, error) {
	return json.Marshal(x)
}

//ToString 返回 json string
func (x *Response) ToString() (string, error) {
	bytes, err := json.Marshal(x)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", bytes), nil
}
