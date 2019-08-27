package errors

import (
	"testing"
)

func Test_ClientError(t *testing.T) {
	err := NewClientError("error code", "error message", nil)

	if err.Error() != "[error code] error message" {
		t.Errorf("[Test_ClientError]=>Error() failed\n")
	}

	if err.HttpStatus() != 400 {
		t.Errorf("[Test_ClientError]=>HttpStatus() failed\n")
	}

	if err.ErrorCode() != "error code" {
		t.Errorf("[Test_ClientError]=>ErrorCode() failed\n")
	}

	if err.Message() != "error message" {
		t.Errorf("[Test_ClientError]=>Message() failed \n")
	}

	if err.String() != "[error code] error message" {
		t.Errorf("[Test_ClientError]=>String() failed\n")
	}

	if _, ok := err.(*ClientError); !ok {
		t.Errorf("[Test_ClientError]=>ClientError failed\n")
	}
}
