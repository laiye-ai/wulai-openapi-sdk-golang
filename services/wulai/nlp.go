package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 自然语言处理类
****************/

/*NLPEntitiesExtract 实体抽取
@query：待实体抽取query [1-1024]characters
@referenced_system_entity：抽取被引用的预设实体,默认未生效. false: 未生效; true: 已生效
*/
func (x *Client) NLPEntitiesExtract(query string, referenced_system_entity bool) (model *NLPEntitiesExtractReponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.nlpEntitiesExtractV2(query, referenced_system_entity)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[NLPEntitiesExtract Response]:%s\n", bytes)
	}

	//返回结果
	model = &NLPEntitiesExtractReponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*NLPTokenize 分词&词性标注
@query：待实体抽取query [1-1024]characters
*/
func (x *Client) NLPTokenize(query string) (model *NLPTokenizeResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.nlpTokenizeV2(query)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[NLPTokenize Response]:%s\n", bytes)
	}

	//返回结果
	model = &NLPTokenizeResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}
