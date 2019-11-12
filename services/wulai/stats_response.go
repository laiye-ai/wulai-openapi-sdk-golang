package wulai

//StatsRecallDailyResponse 天粒度的的问答召回数列表
type StatsRecallDailyResponse struct {
	QaRecallDailyStats []QaRecallDailyStats `json:"qa_recall_daily_stats"` //天粒度的的问答召回数列表
}

//StatsRecallDailyKnowledgeResponse 问答召回数统计列表
type StatsRecallDailyKnowledgeResponse struct {
	QaRecallKnowledgeStats []QaRecallKnowledgeStats `json:"qa_recall_knowledge_stats"` //知识点粒度的问答满意度列表
	PageCount              int                      `json:"page_count"`                //总页数
}

//StatsDailyKnowledgeResponse 问答满意度评价统计列表
type StatsDailyKnowledgeResponse struct {
	PageCount                    int                            `json:"page_count"`
	QaSatisfactionKnowledgeStats []QaSatisfactionKnowledgeStats `json:"qa_satisfaction_knowledge_stats"`
}

//QaSatisfactionKnowledgeStats 知识点粒度的问答满意度列表
type QaSatisfactionKnowledgeStats struct {
	KnowledgeID       int               `json:"knowledge_id"`
	SatisfactionStats SatisfactionStats `json:"satisfaction_stats"`
	StandardQuestion  string            `json:"standard_question"`
}

//SatisfactionStats 满意度
type SatisfactionStats struct {
	ThumbUpCount     int `json:"thumb_up_count"`
	WrongAnswerCount int `json:"wrong_answer_count"`
	BadAnswerCount   int `json:"bad_answer_count"`
}

//QaRecallKnowledgeStats 知识点粒度的问答满意度列表
type QaRecallKnowledgeStats struct {
	KnowledgeID      int           `json:"knowledge_id"`
	StandardQuestion string        `json:"standard_question"`
	QaRecallStats    QaRecallStats `json:"qa_recall_stats"`
}

//MessageStats 消息统计
type MessageStats struct {
	ReceiveCount int `json:"receive_count"` //收到消息数
}

//QaRecallStats 问答召回数
type QaRecallStats struct {
	RecallCount int `json:"recall_count"` //召回次数
}

//QaRecallDailyStats 天粒度的的问答召回数列表
type QaRecallDailyStats struct {
	Date          string        `json:"date"`            //日期，格式如19700101
	MessageStats  MessageStats  `json:"message_stats"`   //消息统计
	QaRecallStats QaRecallStats `json:"qa_recall_stats"` //问答召回数
}
