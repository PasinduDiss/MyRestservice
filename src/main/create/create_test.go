package main_test

import (
	handler "MyRestservice/src/handlers"
	main "MyRestservice/src/main/create"
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

/*
TestCreate is used to unit test the main.go function in the create package
which invokes the create lambda function
*/
func TestCreate(t *testing.T) {

	var devclient handler.TestDeviceClient
	app := &main.App{Handler: devclient}

	tests := []struct {
		description        string
		deviceclient       *handler.TestDeviceClient
		expectedStatuscode int
		expectedBody       string
	}{
		{
			description:        "Testing Bad request",
			deviceclient:       &devclient,
			expectedStatuscode: 400,
			expectedBody:       "Bad Request",
		},
		{
			description:        "Testing Bad request",
			deviceclient:       &devclient,
			expectedStatuscode: 201,
			expectedBody:       "Created",
		},
	}

	for _, test := range tests {
		var ctx context.Context
		var response events.APIGatewayProxyResponse
		var request events.APIGatewayProxyRequest
		request.Body = test.expectedBody

		response, err := app.CreateHandler(ctx, request)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, test.expectedBody, response.Body)
	}
}
