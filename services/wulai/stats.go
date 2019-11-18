package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 统计类
****************/

/*StatsRecallDailyList 查询问答召回数统计列表（日报）
@startDate：开始日期，格式如19700101
@endDate：结束日期，格式如19700101 闭区间。开始时间和结束时间相距不能超过30天
*/
func (x *Client) StatsRecallDailyList(startDate, endDate string) (model *StatsRecallDailyResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.statsRecallDailyListV2(startDate, endDate)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[StatsRecallDaily Response]:%s\n", bytes)
	}

	//返回结果
	model = &StatsRecallDailyResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*StatsRecallDailyKnowledgeList 查询问答召回数统计列表（知识点粒度，日报）
@startDate：开始日期，格式如19700101
@endDate：结束日期，格式如19700101 闭区间。开始时间和结束时间相距不能超过30天
@page：页码，代表查看第几页的数据 page>=1
@pageSize：每页的知识点数量 [1~200]
*/
func (x *Client) StatsRecallDailyKnowledgeList(startDate, endDate string, page, pageSize int) (model *StatsRecallDailyKnowledgeResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.statsRecallDailyKnowledgeListV2(startDate, endDate, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[StatsRecallDailyKnowledgeList Response]:%s\n", bytes)
	}

	//返回结果
	model = &StatsRecallDailyKnowledgeResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*StatsDailyKnowledgeList 查询问答满意度评价统计列表（知识点粒度，日报）
@startDate：开始日期，格式如19700101
@endDate：结束日期，格式如19700101 闭区间。开始时间和结束时间相距不能超过30天
@page：页码，代表查看第几页的数据 page>=1
@pageSize：每页的知识点数量 [1~200]
*/
func (x *Client) StatsDailyKnowledgeList(startDate, endDate string, page, pageSize int) (model *StatsDailyKnowledgeResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.statsDailyKnowledgeListV2(startDate, endDate, page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[StatsDailyKnowledgeList Response]:%s\n", bytes)
	}

	//返回结果
	model = &StatsDailyKnowledgeResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}
