package main_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
TestGetLive runs live tests which access the deployed public REST API,
A sequence of two GET requests are implemented where we assert
for a success response from the first where the path parameter id is a valid id in the database,
while the next request is to trigger a 404 Not Found response by providing an invalid id value.
*/
func TestGetLive(t *testing.T) {

	var ENDPOINT string
	ENDPOINT = "https://xox3imgc04.execute-api.us-east-1.amazonaws.com/dev/devices"

	tests := []struct {
		request            string
		expectedStatuscode int
		expectedBody       string
		err                error
	}{
		{
			request:            ENDPOINT,
			expectedStatuscode: 200,
			expectedBody:       "",
			err:                nil,
		},
		{
			request:            ENDPOINT + "/noid",
			expectedStatuscode: 404,
			expectedBody:       "Not Found",
			err:                nil,
		},
	}

	bs := make([]byte, 99999)
	fmt.Println("Live testing Get")
	for _, test := range tests {

		response, err := http.Get(test.request)
		response.Body.Read(bs)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedStatuscode, response.StatusCode)

	}
}
