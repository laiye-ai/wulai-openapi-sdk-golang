package wulai

import (
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
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserCreate]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_UserUpdate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v1"
	_, err := wulaiClient.UserUpdate("xiao_lai", "nickname", "avatarURL")
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserUpdate]=> %s\n.", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_UserAttributeCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	err := wulaiClient.UserAttributeCreate("xiao_lai", "体重", "120")
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserAttributeCreate]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_GetUserAttribute(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	_, err := wulaiClient.UserAttributeList(true, 1, 100)
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetUserAttribute]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_GetUserInfo(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v1"
	_, err := wulaiClient.UserInfo("xiao_lai")
	if err != nil {
		if err, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetUserInfo]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}

func Test_GetGroupMembers(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.GroupMembers(1, 10, "xiao_lai")
	if err != nil {
		if err, ok := err.(*errors.ServerError); ok {
			t.Errorf("[Test_GetGroupMembers]=> %s\n", err.Message())
		}

		t.Log(err.Error())
	}
}
