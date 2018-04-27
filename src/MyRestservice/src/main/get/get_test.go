package main_test

import (
	handler "MyRestservice/src/handlers"
	main "MyRestservice/src/main/get"
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

/*
TestCreate is used to unit test the main.go function in the get package
which invokes the get lambda function
*/
func TestGet(t *testing.T) {
	var devclient handler.TestDeviceClient
	app := &main.App{Handler: devclient}

	tests := []struct {
		description        string
		deviceclient       *handler.TestDeviceClient
		expectedStatuscode int
		expectedBody       string
	}{
		{
			description:        "Server error",
			deviceclient:       &devclient,
			expectedStatuscode: 500,
			expectedBody:       "Internal Server error",
		},
		{
			description:        "Successfully retrived item with id1",
			deviceclient:       &devclient,
			expectedStatuscode: 200,
			expectedBody:       "{id:id1,deviceModel:deviceModel/id1}",
		},
		{
			description:        "Resource Not found",
			deviceclient:       &devclient,
			expectedStatuscode: 404,
			expectedBody:       "Not found",
		},
	}

	for _, test := range tests {
		var ctx context.Context
		var response events.APIGatewayProxyResponse
		var request events.APIGatewayProxyRequest
		request.Body = test.expectedBody

		response, err := app.GetHandler(ctx, request)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, test.expectedBody, response.Body)
	}
}
