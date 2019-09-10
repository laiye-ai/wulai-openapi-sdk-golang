package errors

import (
	"testing"
)

func Test_ServerError(t *testing.T) {
	err := NewServerError(500, "error message", nil)

	if err.HttpStatus() != 500 {
		t.Errorf("[Test_ServerError]=>HttpStatus() failed\n")
	}

	if _, ok := err.(*ServerError); !ok {
		t.Errorf("[Test_ClientError]=>ServerError failed\n")
	}
}
