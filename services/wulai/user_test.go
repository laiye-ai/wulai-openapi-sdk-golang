package wulai

import (
	"fmt"
	"os"
	"testing"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

func Test_UserCreateV2(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v2"

	userID := "xiao_lai"                                                                                      //用户id [1~128]characters
	nickname := "xiao lai"                                                                                    //用户昵称 <= 128 characters
	avatarURL := "https://laiye-im-saas.oss-cn-beijing.aliyuncs.com/rc-upload-1521637604400-2-login_logo.png" //用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
	resp, err := wulaiClient.UserCreate(userID, nickname, avatarURL)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserCreateV2]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserCreateV2]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserCreateV2]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_UserCreateV1(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v1"

	userID := "xiao_lai"                                                                                      //用户id [1~128]characters
	nickname := "xiao lai"                                                                                    //用户昵称 <= 128 characters
	avatarURL := "https://laiye-im-saas.oss-cn-beijing.aliyuncs.com/rc-upload-1521637604400-2-login_logo.png" //用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
	resp, err := wulaiClient.UserCreate(userID, nickname, avatarURL)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserCreateV1]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserCreateV1]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserCreateV1]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_UserUpdateV1(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v1"

	userID := "xiao_lai"                                                                                      //用户id [1~128]characters
	nickname := "xiao lai"                                                                                    //用户昵称 <= 128 characters
	avatarURL := "https://laiye-im-saas.oss-cn-beijing.aliyuncs.com/rc-upload-1521637604400-2-login_logo.png" //用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
	resp, err := wulaiClient.UserUpdate(userID, nickname, avatarURL)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserUpdateV1]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserUpdateV1]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserUpdateV1]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_UserUpdateV2(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v2"

	userID := "xiao_lai"                                                                                      //用户id [1~128]characters
	nickname := "xiao lai"                                                                                    //用户昵称 <= 128 characters
	avatarURL := "https://laiye-im-saas.oss-cn-beijing.aliyuncs.com/rc-upload-1521637604400-2-login_logo.png" //用户头像地址。用户头像会展示在吾来SaaS的用户列表、消息记录等任何展示用户信息的地方 <= 512 characters
	resp, err := wulaiClient.UserUpdate(userID, nickname, avatarURL)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserUpdateV2]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserUpdateV2]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserUpdateV2]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_UserGetV1(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v1"

	userID := "xiao_lai" //用户id [1~128]characters
	resp, err := wulaiClient.UserGet(userID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetUserInfoV1]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetUserInfoV1]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetUserInfoV1]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_UserGetV2(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v2"

	userID := "xiao_lai" //用户id [1~128]characters
	resp, err := wulaiClient.UserGet(userID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserGetV2]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserGetV2]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserGetV2]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_UserAttributeCreate(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)
	wulaiClient.Version = "v2"

	userID := "xiao_lai" //用户id [1~128]characters
	attrID := 10000      //用户属性id  >=1
	attrValue := "100"   //用户属性值 [1 ~ 128] characters
	err := wulaiClient.UserAttributeCreate(userID, attrID, attrValue)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserAttributeCreate]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_UserAttributeCreate]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserAttributeCreate]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infoln("创建用户成功")
}

func Test_GetUserAttribute(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	isAttrGroup := true //是否可以作为属性组属性
	page := 1           //页码，代表查看第几页的数据，从1开始 >=1
	pageSize := 10      //每页的属性组数量 [1~200]
	resp, err := wulaiClient.UserAttributeList(isAttrGroup, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetUserAttribute]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			log.Infof("[Test_GetUserAttribute]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetUserAttribute]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}

func Test_GetGroupMembers(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	_, err := wulaiClient.groupMembers(1, 10, "xiao_lai")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_GetGroupMembers]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_GetGroupMembers]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_GetGroupMembers]=>%s\n", err.Error())
		}

		return
	}
}

func Test_UserAttributePairList(t *testing.T) {
	secret, pubkey := os.Getenv("secret"), os.Getenv("pubkey")
	wulaiClient := NewClient(secret, pubkey)
	wulaiClient.SetDebug(true)

	userID := "xiao_lai" //用户id [1~128]characters
	resp, err := wulaiClient.UserAttributePairList(userID)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			t.Errorf("[Test_UserAttributePairList]=>%s\n", cliErr.Error())
		} else if serErr, ok := err.(*errors.ServerError); ok {
			fmt.Printf("[Test_UserAttributePairList]=>%s\n", serErr.Error())
		} else {
			log.Infof("[Test_UserAttributePairList]=>%s\n", err.Error())
		}

		return
	}

	log.Infoln("----------------------------------------------------------------------------------------")
	log.Infof("%+v\n", resp)
}
