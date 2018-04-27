package main_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
TestCreateLive runs live tests which access the deployed public REST API,
A sequence of POST request following a GET request is implemented where we assert
for a success response from the POST request response but a 404 'Not Found' from the following
GET request response.
*/
func TestCreateLive(t *testing.T) {
	var ENDPOINT string
	ENDPOINT = "https://xox3imgc04.execute-api.us-east-1.amazonaws.com/dev/devices"

	tests := []struct {
		request            string
		Url                string
		expectedStatuscode int
		expectedBody       string
		err                error
	}{
		{
			request:            "POST",
			Url:                ENDPOINT,
			expectedStatuscode: 201,
			expectedBody:       "",
			err:                nil,
		},
		{
			request:            "GET",
			Url:                ENDPOINT + "/id13",
			expectedStatuscode: 200,
			expectedBody:       "",
			err:                nil,
		},
	}

	fmt.Println("Live testing Create")
	client := &http.Client{}

	for _, test := range tests {
		if test.request == "POST" {
			var newDevice = []byte(`{
				"id": "/devices/id13",
				"deviceModel": "/deviceModel/id13", "name": "Sensor13",
				"note": "Testing a sensor13",
				"serial": "A020000113"
				}`)

			request, err := http.NewRequest(test.request, test.Url, bytes.NewBuffer(newDevice))
			response, err := client.Do(request)
			assert.IsType(t, test.err, err)
			assert.Equal(t, test.expectedStatuscode, response.StatusCode)

		} else {
			request, err := http.NewRequest(test.request, test.Url, nil)
			response, err := client.Do(request)
			assert.IsType(t, test.err, err)
			assert.Equal(t, test.expectedStatuscode, response.StatusCode)
		}
	}
}
