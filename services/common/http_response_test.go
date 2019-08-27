package common

import "testing"

func Test_HTTPResponse(t *testing.T) {
	response := &HTTPResponse{}
	response.ResponseBodyBytes = []byte("message")
	if response.ToString() != "message" {
		t.Fail()
	}
}
