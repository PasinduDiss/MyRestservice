package main_test

import (
	main "ServerlessRestAPI/src/main/create"
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	id := make(map[string]string)
	id["id"] = "id1"
	var CTX events.APIGatewayProxyRequestContext

	tests := []struct {
		request            events.APIGatewayProxyRequest
		responseBody       string
		expectedStatuscode int
		err                error
	}{
		{
			request:            events.APIGatewayProxyRequest{RequestContext: CTX, HTTPMethod: "POST", Body: ""},
			responseBody:       "Bad Request",
			expectedStatuscode: 400,
			err:                nil,
		},
	}
	app := &main.App{}
	for _, test := range tests {
		var ctx context.Context
		var response events.APIGatewayProxyResponse
		response, err := app.CreateHandler(ctx, test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedStatuscode, response.StatusCode)
	}
}
