package main_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListLive(t *testing.T) {

	tests := []struct {
		request            string
		expectedStatuscode int
		expectedBody       string
		err                error
	}{
		{
			request:            "https://xox3imgc04.execute-api.us-east-1.amazonaws.com/dev/devices",
			expectedStatuscode: 200,
			expectedBody:       "",
			err:                nil,
		},
	}

	bs := make([]byte, 99999)
	fmt.Println("Live testing List")
	for _, test := range tests {

		response, err := http.Get(test.request)
		response.Body.Read(bs)

		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedStatuscode, response.StatusCode)
	}
}
