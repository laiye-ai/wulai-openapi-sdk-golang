package wulai

import (
	"fmt"
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

func Test_UserCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	model, err := wulaiClient.UserCreate("xiao_lai", "nickname", "avatarURL")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_UserCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserCreate]=>%s\n", err.Error())
		}

		return
	}

	if len(model.UserID) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_UserUpdate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v1"
	model, err := wulaiClient.UserUpdate("xiao_lai", "nickname", "avatarURL")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_UserUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserUpdate]=>%s\n", serErr.Error())
		}

		return
	}

	if len(model.UserID) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_UserAttributeCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	//wulaiClient.Debug = true
	//属性需要提前在平台创建完成
	//101521 从平台界面查看
	err := wulaiClient.UserAttributeCreate("xiao_lai", "101521", "120")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_UserAttributeCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserAttributeCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserAttributeCreate]=>%s\n", err.Error())
		}
	}
}

func Test_GetUserAttribute(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	model, err := wulaiClient.UserAttributeList(true, 1, 100)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_GetUserAttribute]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetUserAttribute]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetUserAttribute]=>%s\n", err.Error())
		}

		return
	}

	if len(model.UserAttributeUserAttributeValues) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_GetUserInfo(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v1"
	model, err := wulaiClient.UserInfo("xiao_lai")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_GetUserInfo]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetUserInfo]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetUserInfo]=>%s\n", err.Error())
		}

		return
	}

	if len(model.UserID) <= 0 {
		log.Warnf("result is empty. detail=>%+v\n", model)
	}
}

func Test_GetGroupMembers(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.groupMembers(1, 10, "xiao_lai")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); cliErr.ErrorCode() != errors.NetWorkErrorCode && ok {
			t.Errorf("[Test_GetGroupMembers]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetGroupMembers]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetGroupMembers]=>%s\n", err.Error())
		}

		return
	}
}
