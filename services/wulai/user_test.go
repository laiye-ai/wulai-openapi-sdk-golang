package wulai

import (
	"fmt"
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

func Test_UserCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	_, err := wulaiClient.UserCreate("xiao_lai", "nickname", "avatarURL")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_UserCreate]=>%s\n", serErr.Error())
		}
	}
}

func Test_UserUpdate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v1"
	_, err := wulaiClient.UserUpdate("xiao_lai", "nickname", "avatarURL")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserUpdate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_UserUpdate]=>%s\n", serErr.Error())
		}
	}
}

func Test_UserAttributeCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	err := wulaiClient.UserAttributeCreate("xiao_lai", "体重", "120")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserAttributeCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_UserAttributeCreate]=>%s\n", serErr.Error())
		}
	}
}

func Test_GetUserAttribute(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	_, err := wulaiClient.UserAttributeList(true, 1, 100)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetUserAttribute]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetUserAttribute]=>%s\n", serErr.Error())
		}
	}
}

func Test_GetUserInfo(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v1"
	_, err := wulaiClient.UserInfo("xiao_lai")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetUserInfo]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetUserInfo]=>%s\n", serErr.Error())
		}
	}
}

func Test_GetGroupMembers(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.GroupMembers(1, 10, "xiao_lai")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetGroupMembers]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetGroupMembers]=>%s\n", serErr.Error())
		}
	}
}
