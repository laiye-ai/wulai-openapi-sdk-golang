package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

//Credential 认证
type Credential struct {
	secret string
	pubkey string
}

//NewCredential 创建Credential
func NewCredential(secret, pubkey string) *Credential {
	return &Credential{
		secret: secret,
		pubkey: pubkey,
	}
}

//GetRandomString 获取随机字符串
func (a *Credential) GetRandomString(length int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(int(len(bytes)))])
	}
	return string(result)
}

//GetHeaders 获取鉴权http header
func (a *Credential) GetHeaders() map[string]string {
	nonce := a.GetRandomString(32)
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	s := fmt.Sprintf("%s%s%s", nonce, timestamp, a.secret)
	t := sha1.New()
	io.WriteString(t, s)
	sign := fmt.Sprintf("%x", t.Sum(nil))
	return map[string]string{
		"Api-Auth-pubkey":    a.pubkey,
		"Api-Auth-nonce":     nonce,
		"Api-Auth-timestamp": timestamp,
		"Api-Auth-sign":      sign,
	}
}
