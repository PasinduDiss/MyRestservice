package main_test

import (
	main "ServerlessRestAPI/src/main/get"
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	id := make(map[string]string)
	id["id"] = "id1"
	tests := []struct {
		request            events.APIGatewayProxyRequest
		expectedStatuscode int
		err                error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request:            events.APIGatewayProxyRequest{Body: ""},
			expectedStatuscode: 500,
			err:                nil,
		},
	}
	app := &main.App{}
	for _, test := range tests {
		var ctx context.Context
		var response events.APIGatewayProxyResponse
		response, err := app.GetHandler(ctx, test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedStatuscode, response.StatusCode)
	}
}
