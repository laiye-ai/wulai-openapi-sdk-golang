package wulai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

/****************
- 词库管理类
****************/

//DicEntityList 查询全部实体概要
func (x *Client) DicEntityList(page, pageSize int) (model *DicList, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityListV2(page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicEntityList Response]:%s\n", bytes)
	}

	//返回结果
	model = &DicList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicEntityGet 查询一个实体详情
 *@id：实体ID
 */
func (x *Client) DicEntityGet(id int) (model *EntityReponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityGetV2(id)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicEntityGet Response]:%s\n", bytes)
	}

	//返回结果
	model = &EntityReponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicEntityDelete 删除实体
 *@id：实体ID
 */
func (x *Client) DicEntityDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[DicEntityDelete Response]:%s\n", bytes)
	}

	return nil
}

/*DicEntityEnumCreate 创建枚举实体
 *@name：实体名称[1-200]字符
 */
func (x *Client) DicEntityEnumCreate(name string) (model *EntityEnumReponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityEnumCreateV2(name)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicEntityEnumCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &EntityEnumReponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicEntityEnumValueCreate 创建枚举实体值
 *@entityID：实体ID
 *@standardValue：标准值(归一化值)
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) DicEntityEnumValueCreate(entityID int, standardValue string, synonyms []string) (model *EntityEnumReponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityEnumValueCreateV2(entityID, standardValue, synonyms)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicEntityEnumValueCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &EntityEnumReponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicEntityEnumValueDelete 删除枚举实体值
 *@entityID：实体ID
 *@standardValue：标准值(归一化值)
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) DicEntityEnumValueDelete(entityID int, standardValue string, synonyms []string) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityEnumValueDeleteV2(entityID, standardValue, synonyms)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[DicEntityEnumValueDelete Response]:%s\n", bytes)
	}

	return nil
}

/*DicEntityIntentCreate 创建意图实体
@name：实体名称
@standardValue：标准值(归一化值)
*/
func (x *Client) DicEntityIntentCreate(name, standardValue string) (model *IntentEntityResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityIntentCreateV2(name, standardValue)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicEntityIntentCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &IntentEntityResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicEntityIntentValueCreate 创建意图实体值相似说法
 *@entityID：实体ID
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) DicEntityIntentValueCreate(entityID int, synonyms []string) (model *IntentEntityResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityIntentValueCreateV2(entityID, synonyms)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicEntityIntentValueCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &IntentEntityResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicEntityIntentValueDelete 删除意图实体值相似说法
 *@entityID：实体ID
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) DicEntityIntentValueDelete(entityID int, synonyms []string) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicEntityIntentValueDeleteV2(entityID, synonyms)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[DicEntityIntentValueDelete Response]:%s\n", bytes)
	}

	return nil
}

/****************
- 专有词汇
****************/

//DicTermList 查询专有词汇列表
func (x *Client) DicTermList(page, pageSize int) (model *TermList, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicTermListV2(page, pageSize)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicTermList Response]:%s\n", bytes)
	}

	//返回结果
	model = &TermList{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicTermCreate 创建专有词汇
 *@name：专有词汇的名称 [1~64]characters
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) DicTermCreate(name string, synonyms []string) (model *TermResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicTermCreateV2(name, synonyms)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicTermCreate Response]:%s\n", bytes)
	}

	//返回结果
	model = &TermResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicTermUpdate 更新专有词汇
 *@id：专有词汇id
 *@name：专有词汇的名称 [1~64]characters
 *@synonyms：标准值的相似说法 []string
 */
func (x *Client) DicTermUpdate(id int, name string, synonyms []string) (model *TermResponse, err error) {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return nil, errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicTermUpdateV2(id, name, synonyms)
	if err != nil {
		return nil, err
	}

	if x.Debug {
		log.Debugf("[DicTermUpdate Response]:%s\n", bytes)
	}

	//返回结果
	model = &TermResponse{}
	if err = json.Unmarshal(bytes, model); err != nil {
		return nil, errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonMarshalErrorMessage, err)
	}

	return model, nil
}

/*DicTermDelete 删除专有词汇
 *@id：专有词汇id
 */
func (x *Client) DicTermDelete(id int) error {

	if strings.ToUpper(x.Version) == "V1" {

		errMsg := fmt.Sprintf(errors.UnsupportedMethodErrorMessage, "V1", "V2")
		return errors.NewClientError(errors.UnsupportedMethodErrorCode, errMsg, nil)
	}

	bytes, err := x.dicTermDeleteV2(id)
	if err != nil {
		return err
	}

	if x.Debug {
		log.Debugf("[DicTermDelete Response]:%s\n", bytes)
	}

	return nil
}
