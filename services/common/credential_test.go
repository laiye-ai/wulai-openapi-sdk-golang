package common

import (
	"os"
	"testing"
)

func Test_GetHeaders(t *testing.T) {
	secret, pubkey:= os.Getenv("secret"),os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	headers := credential.GetHeaders()

	if headers["Api-Auth-pubkey"] == "" {
		t.Error("pubkey is missed.")
	}

	if headers["Api-Auth-sign"] == "" {
		t.Error("sign is missed.")
	}
}

func Test_GetRandomString(t *testing.T) {
	secret, pubkey:= os.Getenv("secret"),os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)
	str := credential.GetRandomString(15)
	if len(str) != 15 {
		t.Error("random strint len is wrong.")
	}
}

func Benchmark_GetHeaders(t *testing.B) {
	secret, pubkey:= os.Getenv("secret"),os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)

	for i := 0; i < t.N; i++ {
		headers := credential.GetHeaders()
		if headers["Api-Auth-pubkey"] == "" {
			t.Error("pubkey is missed.")
		}
	}
}

func Benchmark_GetRandomString(t *testing.B) {
	secret, pubkey:= os.Getenv("secret"),os.Getenv("pubkey")
	credential := NewCredential(secret, pubkey)

	for i := 0; i < t.N; i++ {
		str := credential.GetRandomString(15)
		if len(str) != 15 {
			t.Error("random strint len is wrong.")
		}
	}
}
