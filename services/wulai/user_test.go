package wulai

import (
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

func Test_CreateUser(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	wulaiClient.Debug = true
	_, err := wulaiClient.UserCreate("userID", "nickname", "avatarURL")
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if _, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_CreateUser]=> failed.")
		} else {
			t.Error("[Test_CreateUser]=> failed.")
		}
	}
}

func Test_UserUpdate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	err := wulaiClient.UserUpdate("userID", "nickname", "avatarURL")
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if _, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_UserUpdate]=> failed.")
		} else {
			t.Error("[Test_UserUpdate]=> failed.")
		}
	}
}

func Test_GetUserAttribute(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.Version = "v2"
	_, err := wulaiClient.UserAttributeList(true, 1, 100)
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if _, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_GetUserAttribute]=> failed.")
		} else {
			t.Error("[Test_GetUserAttribute]=> failed.")
		}
	}
}

func Test_GetUserInfo(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.UserInfo("userID")
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if _, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_GetUserInfo]=> failed.")
		} else {
			t.Error("[Test_GetUserInfo]=> failed.")
		}
	}
}

func Test_GetGroupMembers(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.GroupMembers(1, 10, "userID")
	if err != nil {
		if _, ok := err.(*errors.ServerError); ok {

		} else if _, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_GetGroupMembers]=> failed.")
		} else {
			t.Error("[Test_GetGroupMembers]=> failed.")
		}
	}
}
