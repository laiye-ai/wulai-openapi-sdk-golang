package errors

import (
	"testing"
)

func Test_ServerError(t *testing.T) {
	err := NewServerError(500, "error message", nil)

	if err.Error() != "[error code] error message" {
		t.Errorf("[Test_ServerError]=>Error() failed\n")
	}

	if err.HttpStatus() != 500 {
		t.Errorf("[Test_ServerError]=>HttpStatus() failed\n")
	}

	if err.ErrorCode() != "error code" {
		t.Errorf("[Test_ServerError]=>ErrorCode() failed\n")
	}

	if err.Message() != "error message" {
		t.Errorf("[Test_ServerError]=>Message() failed\n")
	}

	if err.Error() != "[error code] error message" {
		t.Errorf("[Test_ServerError]=>Error() failed\n")
	}

	if err.String() != "[error code] error message" {
		t.Errorf("[Test_ServerError]=>String() failed\n")
	}

	if _, ok := err.(*ServerError); !ok {
		t.Errorf("[Test_ClientError]=>ServerError failed\n")
	}
}
