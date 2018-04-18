package main_test

import (
	main "ServerlessRestAPI/src/main/delete"
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
			responseBody:       "Internal server error",
			expectedStatuscode: 500,
			err:                nil,
		},
	}
	app := &main.App{}
	for _, test := range tests {
		var ctx context.Context
		var response events.APIGatewayProxyResponse
		response, err := app.DeleteHandler(ctx, test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedStatuscode, response.StatusCode)
	}
}
