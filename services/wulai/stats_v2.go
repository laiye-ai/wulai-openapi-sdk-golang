package wulai

import (
	"fmt"
)

/****************
- 统计类
****************/

/*statsRecallDailyListV2 查询问答召回数统计列表（日报）
@startDate：开始日期，格式如19700101
@endDate：结束日期，格式如19700101 闭区间。开始时间和结束时间相距不能超过30天
*/
func (x *Client) statsRecallDailyListV2(startDate, endDate string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/stats/qa/recall/daily/list", x.Endpoint, x.Version)

	input := fmt.Sprintf(`
	{
		"start_date": "%s",
		"end_date": "%s"
	}`, startDate, endDate)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*statsRecallDailyKnowledgeListV2 查询问答召回数统计列表（知识点粒度，日报）
@startDate：开始日期，格式如19700101
@endDate：结束日期，格式如19700101 闭区间。开始时间和结束时间相距不能超过30天
@page：页码，代表查看第几页的数据 page>=1
@pageSize：每页的知识点数量 [1~200]
*/
func (x *Client) statsRecallDailyKnowledgeListV2(startDate, endDate string, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/stats/qa/recall/daily/knowledge/list", x.Endpoint, x.Version)

	input := fmt.Sprintf(`
	{
		"start_date": "%s",
		"end_date": "%s",
		"page": %v,
		"page_size": %v
	}`, startDate, endDate, page, pageSize)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}

/*statsDailyKnowledgeListV2 查询问答满意度评价统计列表（知识点粒度，日报）
@startDate：开始日期，格式如19700101
@endDate：结束日期，格式如19700101 闭区间。开始时间和结束时间相距不能超过30天
@page：页码，代表查看第几页的数据 page>=1
@pageSize：每页的知识点数量 [1~200]
*/
func (x *Client) statsDailyKnowledgeListV2(startDate, endDate string, page, pageSize int) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/stats/qa/satisfaction/daily/knowledge/list", x.Endpoint, x.Version)

	input := fmt.Sprintf(`
	{
		"start_date": "%s",
		"end_date": "%s",
		"page": %v,
		"page_size": %v
	}`, startDate, endDate, page, pageSize)

	respBytes, err := x.HTTPClient.Request("POST", url, []byte(input), 1)
	if err != nil {
		return nil, err
	}

	return respBytes.ResponseBodyBytes, nil
}
