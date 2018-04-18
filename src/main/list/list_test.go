package main_test

import (
	handler "MyRestservice/src/handlers"
	main "MyRestservice/src/main/list"
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
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
			description:        "Successfully revcived all items in dynamodb",
			deviceclient:       &devclient,
			expectedStatuscode: 200,
			expectedBody:       "{devices{id:id1,deviceModel:deviceModel/id1},{id:id2,deviceModel:deviceModel/id2}",
		},
	}

	for _, test := range tests {
		var ctx context.Context
		var response events.APIGatewayProxyResponse
		var request events.APIGatewayProxyRequest
		request.Body = test.expectedBody

		response, err := app.ListHandler(ctx, request)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, test.expectedBody, response.Body)
	}
}
