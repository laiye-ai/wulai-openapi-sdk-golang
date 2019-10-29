package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

func Test_QaSimilarQuestionDelete(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	err := wulaiClient.QaSimilarQuestionDelete("id")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionDelete]=>%s\n", serErr.Error())
		}

		return
	}
}

func Test_QaCreateUserAttributeGroup(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupName := "SDK"
	attributeID := "101185"
	attributeName := "shzy2012"

	resp, err := wulaiClient.QaCreateUserAttributeGroup(groupName, attributeID, attributeName)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionDelete]=>%s\n", serErr.Error())
		}

		return
	}

	log.Infof("%+v\n", resp)

	/*result example
	{
		"user_attribute_group_item":{
			"user_attribute_group":{
				"id":"6163",
				"name":"SDK"
			},
			"user_attribute_user_attribute_values":[
				{
					"user_attribute":{
						"id":"101185",
						"name":"公众号",
						"type":"USER_ATTRIBUTE_TYPE_SYSTEM",
						"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
						"use_in_user_attribute_group":true,
						"lifespan":0
					},
					"user_attribute_value":{
						"id":"5089503",
						"name":"shzy2012"
					}
				},
				{
					"user_attribute":{
						"id":"101317",
						"name":"年龄",
						"type":"USER_ATTRIBUTE_TYPE_TEMP",
						"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
						"use_in_user_attribute_group":true,
						"lifespan":0
					},
					"user_attribute_value":{
						"id":"0",
						"name":""
					}
				},
				{
					"user_attribute":{
						"id":"101321",
						"name":"企业部门",
						"type":"USER_ATTRIBUTE_TYPE_TEMP",
						"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
						"use_in_user_attribute_group":true,
						"lifespan":0
					},
					"user_attribute_value":{
						"id":"0",
						"name":""
					}
				},
				{
					"user_attribute":{
						"id":"101324",
						"name":"籍贯",
						"type":"USER_ATTRIBUTE_TYPE_TEMP",
						"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
						"use_in_user_attribute_group":true,
						"lifespan":0
					},
					"user_attribute_value":{
						"id":"0",
						"name":""
					}
				}
			]
		}
	}

	*/
}

func Test_QaListUserAttributeGroupItems(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	page := 1
	pageSize := 10
	_, err := wulaiClient.QaListUserAttributeGroupItems(page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionDelete]=>%s\n", serErr.Error())
		}

		return
	}

	/*result example
	{
		"user_attribute_group_items":[
			{
				"user_attribute_group":{
					"id":"6154",
					"name":"SDK"
				},
				"user_attribute_user_attribute_values":[
					{
						"user_attribute":{
							"id":"101185",
							"name":"公众号",
							"type":"USER_ATTRIBUTE_TYPE_DEFAULT",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_DEFAULT",
							"use_in_user_attribute_group":false,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"5087333",
							"name":"Hello"
						}
					},
					{
						"user_attribute":{
							"id":"101317",
							"name":"年龄",
							"type":"USER_ATTRIBUTE_TYPE_DEFAULT",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_DEFAULT",
							"use_in_user_attribute_group":false,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"0",
							"name":""
						}
					},
					{
						"user_attribute":{
							"id":"101321",
							"name":"企业部门",
							"type":"USER_ATTRIBUTE_TYPE_DEFAULT",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_DEFAULT",
							"use_in_user_attribute_group":false,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"0",
							"name":""
						}
					},
					{
						"user_attribute":{
							"id":"101324",
							"name":"籍贯",
							"type":"USER_ATTRIBUTE_TYPE_DEFAULT",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_DEFAULT",
							"use_in_user_attribute_group":false,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"0",
							"name":""
						}
					}
				]
			}
		],
		"page_count":1
	}
	*/
}

func Test_QaUpdateUserAttributeGroup(t *testing.T) {

	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	groupID := "6163"
	groupName := "SDK"
	attributes := make(map[string]string)
	attributes["101185"] = "shzy2012-update"

	_, err := wulaiClient.QaUpdateUserAttributeGroup(groupID, groupName, attributes)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_QaSimilarQuestionDelete]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_QaSimilarQuestionDelete]=>%s\n", serErr.Error())
		}

		return
	}

	/*
		{
			"user_attribute_group_item":{
				"user_attribute_group":{
					"id":"6163",
					"name":"SDK"
				},
				"user_attribute_user_attribute_values":[
					{
						"user_attribute":{
							"id":"101185",
							"name":"公众号",
							"type":"USER_ATTRIBUTE_TYPE_SYSTEM",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
							"use_in_user_attribute_group":true,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"5089551",
							"name":"shzy2012-update"
						}
					},
					{
						"user_attribute":{
							"id":"101317",
							"name":"年龄",
							"type":"USER_ATTRIBUTE_TYPE_TEMP",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
							"use_in_user_attribute_group":true,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"0",
							"name":""
						}
					},
					{
						"user_attribute":{
							"id":"101321",
							"name":"企业部门",
							"type":"USER_ATTRIBUTE_TYPE_TEMP",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
							"use_in_user_attribute_group":true,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"0",
							"name":""
						}
					},
					{
						"user_attribute":{
							"id":"101324",
							"name":"籍贯",
							"type":"USER_ATTRIBUTE_TYPE_TEMP",
							"value_type":"USER_ATTRIBUTE_VALUE_TYPE_STRING",
							"use_in_user_attribute_group":true,
							"lifespan":0
						},
						"user_attribute_value":{
							"id":"0",
							"name":""
						}
					}
				]
			}
		}
	*/
}
