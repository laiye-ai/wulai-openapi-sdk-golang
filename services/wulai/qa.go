package wulai

import (
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//QaSimilarQuestionDelete 删除相似问
func (x *Client) QaSimilarQuestionDelete(id string) (err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.qaSimilarQuestionDeleteV2(id)

	if x.Debug {
		log.Debugf("[QaSimilarQuestionDelete Response]:%s\n", bytes)
	}

	if err != nil {
		return err
	}

	return nil
}
